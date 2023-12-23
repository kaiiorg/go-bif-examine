package main

import (
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

	cli.rootCmd.AddCommand(cli.downloadCmd)
	cli.downloadCmd.AddCommand(cli.downloadResourceCmd)
}

func (cli *Cli) downloadResource(cmd *cobra.Command, args []string) error {
	log.Info().Msg("downloadResource")
	return nil
}
