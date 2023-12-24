package main

import (
	"errors"
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
	cli.uploadKeyCmd.Flags().StringVar(&cli.uploadKeyName, "project-name", "", "Name of the project; will use the name of the keyfile if not provided")
	cli.uploadBifCmd.Flags().UintVar(&cli.uploadBifProjectId, "project-id", 0, "Project ID to upload bif to")
	cli.uploadBifCmd.MarkFlagRequired("project-id")
	cli.uploadBifCmd.Flags().StringVar(&cli.uploadBifNameInKey, "name-in-key", "", "Exact name used in the key to link to this file")
	cli.uploadBifCmd.MarkFlagRequired("name-in-key")
	cli.rootCmd.AddCommand(cli.uploadCmd)
	cli.uploadCmd.AddCommand(cli.uploadKeyCmd, cli.uploadBifCmd)
}

func (cli *Cli) uploadKey(cmd *cobra.Command, args []string) error {
	// Assume all the following args are paths to keys; attempt to upload them
	for _, keyPath := range args {
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
			return err
		}

		resp, err := cli.grpcClient.UploadKey(cmd.Context(), req)
		if err != nil {
			return err
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
	}
	return nil
}

func (cli *Cli) uploadBif(cmd *cobra.Command, args []string) error {
	// Assume the first string in the args is the path
	if len(args) == 0 {
		return ErrBifPathNotProvided
	}
	req := &pb.UploadBifRequest{
		ProjectId: uint32(cli.uploadBifProjectId),
		FileName:  strings.ToLower(filepath.Base(args[0])),
		NameInKey: cli.uploadBifNameInKey,
	}
	var err error
	req.Contents, err = os.ReadFile(args[0])
	if err != nil {
		return err
	}

	resp, err := cli.grpcClient.UploadBif(cmd.Context(), req)
	if err != nil {
		return err
	}
	log.Info().
		Uint32("resourcesFound", resp.GetResourcesFound()).
		Uint32("resourcesNotFound", resp.GetResourcesNotFound()).
		Msg("Uploaded bif")
	return nil
}
