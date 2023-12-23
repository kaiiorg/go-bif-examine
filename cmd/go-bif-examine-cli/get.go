package main

import (
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func (cli *Cli) buildGetCmds() {
	cli.getCmd = &cobra.Command{
		Use: "get",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.getCmd.Usage()
		},
	}
	cli.getProjectCmd = &cobra.Command{
		Use:     "project",
		Short:   "Gets all projects",
		PreRunE: cli.grpcConfigure,
		RunE:    cli.getProject,
	}

	cli.rootCmd.AddCommand(cli.getCmd)
	cli.getCmd.AddCommand(cli.getProjectCmd)
}

func (cli *Cli) getProject(cmd *cobra.Command, args []string) error {
	allProjects, err := cli.grpcClient.GetAllProjects(
		cli.getProjectCmd.Context(),
		&pb.GetAllProjectsRequest{},
	)
	if err != nil {
		return err
	}
	log.Info().Int("count", len(allProjects.Projects)).Msg("Got all projects")
	for _, project := range allProjects.GetProjects() {
		log.Info().
			Uint32("id", project.GetId()).
			Str("name", project.GetName()).
			Str("originalKeyFileName", project.GetOriginalKeyFileName()).
			Send()
	}
	return nil
}
