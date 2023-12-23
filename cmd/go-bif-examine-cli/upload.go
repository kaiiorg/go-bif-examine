package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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

	cli.rootCmd.AddCommand(cli.uploadCmd)
	cli.uploadCmd.AddCommand(cli.uploadKeyCmd, cli.uploadBifCmd)
}

func (cli *Cli) uploadKey(cmd *cobra.Command, args []string) error {
	log.Info().Msg("uploadKey")
	return nil
}

func (cli *Cli) uploadBif(cmd *cobra.Command, args []string) error {
	log.Info().Msg("uploadBif")
	return nil
}
