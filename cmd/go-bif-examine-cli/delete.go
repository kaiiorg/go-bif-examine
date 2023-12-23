package main

import (
	"strconv"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func (cli *Cli) buildDeleteCmds() {
	cli.deleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm", "del"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.deleteCmd.Usage()
		},
	}
	cli.deleteProjectCmd = &cobra.Command{
		Use:     "project",
		Short:   "Deletes projects by ID",
		PreRunE: cli.grpcConfigure,
		RunE:    cli.deleteProject,
	}

	cli.rootCmd.AddCommand(cli.deleteCmd)
	cli.deleteCmd.AddCommand(cli.deleteProjectCmd)
}

func (cli *Cli) deleteProject(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cli.deleteProjectCmd.Usage()
	}
	for _, arg := range args {
		projectId, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		_, err = cli.grpcClient.DeleteProject(
			cli.deleteProjectCmd.Context(),
			&pb.DeleteProjectRequest{
				ProjectId: uint32(projectId),
			},
		)
		if err != nil {
			log.Error().Int("projectId", projectId).Msg("Failed to delete project")
			return err
		}
		log.Info().Int("projectId", projectId).Msg("Deleted project")
	}
	return nil
}
