package main

import (
	"context"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"strconv"

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
	cli.getBifsMissingContentCmd = &cobra.Command{
		Use:     "bif",
		Short:   "Gets the name of bif files still missing for the given project ID",
		PreRunE: cli.grpcConfigure,
		RunE:    cli.getBifsMissingContent,
	}

	cli.rootCmd.AddCommand(cli.getCmd)
	cli.getCmd.AddCommand(cli.getProjectCmd, cli.getBifsMissingContentCmd)
}

func (cli *Cli) getProject(cmd *cobra.Command, args []string) error {
	var err error
	// If there are no additional args provided, assume that we've been asked to get all projects
	// Otherwise, assume each is a project ID we've been asked to get
	if len(args) == 0 {
		err = cli.getAllProjects(cmd.Context())
	} else {
		err = cli.getProjectByIds(cmd.Context(), args)
	}
	return err
}

func (cli *Cli) getProjectByIds(ctx context.Context, projectIdStrs []string) error {
	for _, projectIdStr := range projectIdStrs {
		projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
		if err != nil {
			return err
		}

		project, err := cli.grpcClient.GetProjectById(
			ctx,
			&pb.GetProjectByIdRequest{
				Id: uint32(projectId),
			},
		)
		if err != nil {
			return err
		}

		log.Info().
			Uint32("id", project.GetProject().GetId()).
			Str("name", project.GetProject().GetName()).
			Str("originalKeyFileName", project.GetProject().GetOriginalKeyFileName()).
			Send()
	}
	return nil
}

func (cli *Cli) getAllProjects(ctx context.Context) error {
	allProjects, err := cli.grpcClient.GetAllProjects(
		ctx,
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

func (cli *Cli) getBifsMissingContent(cmd *cobra.Command, args []string) error {
	for _, projectIdStr := range args {
		projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
		if err != nil {
			return err
		}

		bifs, err := cli.grpcClient.GetBifsMissingContents(
			cmd.Context(),
			&pb.GetBifsMissingContentsRequest{
				ProjectId: uint32(projectId),
			},
		)
		if err != nil {
			return err
		}
		log.Info().
			Str("projectId", projectIdStr).
			Int("count", len(bifs.NameInKey)).
			Msg("Got bifs missing content for project")
		for _, bif := range bifs.GetNameInKey() {
			log.Info().Str("bifMissingContent", bif).Send()
		}

	}
	return nil
}
