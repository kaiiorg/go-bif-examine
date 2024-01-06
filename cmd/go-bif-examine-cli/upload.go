package main

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	ErrBifPathNotProvided = errors.New("bif path not provided")
)

func (cli *Cli) buildUploadCmds() {
	cli.uploadCmd = &cobra.Command{
		Use: "upload",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.uploadCmd.Usage()
		},
	}
	cli.uploadKeyCmd = &cobra.Command{
		Use:     "key",
		Short:   "Uploads a key and creates a new project",
		PreRunE: cli.grpcConfigure,
		RunE:    cli.uploadKey,
	}
	cli.uploadBifCmd = &cobra.Command{
		Use:     "bif",
		Short:   "Uploads a bif to a project",
		PreRunE: cli.grpcConfigure,
		RunE:    cli.uploadBif,
	}
	cli.uploadAutoCmd = &cobra.Command{
		Use:     "auto",
		Short:   "Uploads a key to create a new project, then tries to automatically find and upload the bifs defined within",
		PreRunE: cli.grpcConfigure,
		RunE:    cli.uploadAuto,
	}
	cli.uploadKeyCmd.Flags().StringVar(&cli.uploadKeyName, "project-name", "", "Name of the project; will use the name of the keyfile if not provided")
	cli.uploadBifCmd.Flags().UintVar(&cli.uploadBifProjectId, "project-id", 0, "Project ID to upload bif to")
	cli.uploadBifCmd.MarkFlagRequired("project-id")
	cli.uploadBifCmd.Flags().StringVar(&cli.uploadBifNameInKey, "name-in-key", "", "Exact name used in the key to link to this file")
	cli.uploadBifCmd.MarkFlagRequired("name-in-key")
	cli.uploadAutoCmd.Flags().StringVar(&cli.uploadKeyName, "project-name", "", "Name of the project; will use the name of the keyfile if not provided")
	cli.rootCmd.AddCommand(cli.uploadCmd)
	cli.uploadCmd.AddCommand(cli.uploadKeyCmd, cli.uploadBifCmd, cli.uploadAutoCmd)
}

func (cli *Cli) uploadKey(cmd *cobra.Command, args []string) error {
	// Assume all the following args are paths to keys; attempt to upload them
	for _, keyPath := range args {
		_, err := cli.uploadSingleKey(cmd.Context(), keyPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cli *Cli) uploadSingleKey(ctx context.Context, keyPath string) (*pb.UploadKeyResponse, error) {
	if cli.uploadKeyName == "" {
		cli.uploadKeyName = filepath.Base(keyPath)
	}
	req := &pb.UploadKeyRequest{
		ProjectName: cli.uploadKeyName,
		FileName:    filepath.Base(keyPath),
	}
	var err error
	req.Contents, err = os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	resp, err := cli.grpcClient.UploadKey(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Info().
		Uint32("projectId", resp.GetProjectId()).
		Str("parsedVersion", resp.GetKey().GetParsedVersion()).
		Str("parsedSignature", resp.GetKey().GetParsedSignature()).
		Uint32("resourceEntryCount", resp.GetKey().GetResourceEntryCount()).
		Uint32("resourcesWithAudio", resp.GetKey().GetResourcesWithAudio()).
		Int("bifFilesContainingAudioCount", len(resp.GetKey().GetBifFilesContainingAudio())).
		Strs("bifFilesContainingAudio", resp.GetKey().GetBifFilesContainingAudio()).
		Msg("Uploaded key")

	return resp, nil
}

func (cli *Cli) uploadBif(cmd *cobra.Command, args []string) error {
	// Assume the first string in the args is the path
	if len(args) == 0 {
		return ErrBifPathNotProvided
	}
	return cli.uploadSingleBif(cmd.Context(), args[0], uint32(cli.uploadBifProjectId))
}

func (cli *Cli) uploadSingleBif(ctx context.Context, bifPath string, projectId uint32) error {
	req := &pb.UploadBifRequest{
		ProjectId: projectId,
		FileName:  strings.ToLower(filepath.Base(bifPath)),
		NameInKey: cli.uploadBifNameInKey,
	}
	file, err := os.Open(bifPath)
	if err != nil {
		return err
	}
	defer file.Close()

	stream, err := cli.grpcClient.UploadBif(ctx)
	if err != nil {
		return err
	}

	buf := make([]byte, 2097152) // 2MB
	offset := int64(0)
	for {
		bytesRead, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		req.Offset = offset
		req.Contents = buf[:bytesRead]
		err = stream.Send(req)
		if err != nil {
			return err
		}
		offset += int64(bytesRead)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Info().
		Uint32("resourcesFound", resp.GetResourcesFound()).
		Uint32("resourcesNotFound", resp.GetResourcesNotFound()).
		Msg("Uploaded bif")
	return nil
}

func (cli *Cli) uploadAuto(cmd *cobra.Command, args []string) error {
	// Assume all the following args are paths to keys; attempt to upload them
	for _, keyPath := range args {
		resp, err := cli.uploadSingleKey(cmd.Context(), keyPath)
		if err != nil {
			return err
		}
		keyDir := filepath.Dir(keyPath)

		for _, keyBifPath := range resp.GetKey().GetBifFilesContainingAudio() {
			err = cli.uploadSingleBif(cmd.Context(), filepath.Join(keyDir, keyBifPath), resp.GetProjectId())
			if err != nil {
				return err
			}
		}

	}
	return nil
}
