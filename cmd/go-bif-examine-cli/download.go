package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func (cli *Cli) buildDownloadCmds() {
	cli.downloadCmd = &cobra.Command{
		Use: "download",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.downloadCmd.Usage()
		},
	}
	cli.downloadResourceCmd = &cobra.Command{
		Use:     "resource",
		Short:   "Downloads provided resources by id",
		PreRunE: cli.grpcConfigure,
		RunE:    cli.downloadResource,
	}
	cli.downloadResourceCmd.Flags().StringVar(&cli.downloadResourceDir, "output-dir", ".", "Where to save downloaded resources to")
	cli.rootCmd.AddCommand(cli.downloadCmd)
	cli.downloadCmd.AddCommand(cli.downloadResourceCmd)
}

func (cli *Cli) downloadResource(cmd *cobra.Command, args []string) error {
	// Assume each argument is the DB ID of a resource
	for _, idStr := range args {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			log.Error().
				Err(err).
				Str("given", idStr).
				Msg("Failed to convert given ID to an int; continuing with list")
			continue
		}

		req := &pb.DownloadResourceRequest{
			ResourceId: uint32(id),
		}

		resp, err := cli.grpcClient.DownloadResource(cmd.Context(), req)
		if err != nil {
			log.Error().
				Uint64("resourceId", id).
				Msg("Failed to download resource; continuing with list")
		}
		saveTo := filepath.Join(cli.downloadResourceDir, resp.GetName())
		err = os.WriteFile(saveTo, resp.GetContent(), 0666)
		if err != nil {
			log.Error().Err(err).Msg("Failed to write content to file; stopping")
			return err
		}
		log.Info().
			Str("name", resp.GetName()).
			Uint32("listedSize", resp.GetSize()).
			Int("downloadedSize", len(resp.GetContent())).
			Str("path", saveTo).
			Msg("Downloaded resource")
	}
	return nil
}
