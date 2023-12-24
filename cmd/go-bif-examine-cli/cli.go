package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"github.com/kaiiorg/go-bif-examine/pkg/util"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	applicationName        = "go-bif-examine-cli"
	applicationDescription = "Quick and dirty CLI tool, for now."
)

/*
 * go-bif-examine-cli
 * 		get
 * 			project
 * 		delete
 *			project
 * 		upload
 *			key ${path}
 *			bif ${path}
 *		download
 *			resource ${ID}
 */

type Cli struct {
	grpcClient pb.BifExamineClient

	// Global variables
	logLevel   string
	grpcServer string

	// Command specific variables
	// TODO split out commands into their own struct so they aren't storing their variables together
	uploadKeyName       string
	uploadBifProjectId  uint
	uploadBifNameInKey  string
	downloadResourceDir string

	rootCmd             *cobra.Command
	getCmd              *cobra.Command
	getProjectCmd       *cobra.Command
	deleteCmd           *cobra.Command
	deleteProjectCmd    *cobra.Command
	uploadCmd           *cobra.Command
	uploadKeyCmd        *cobra.Command
	uploadBifCmd        *cobra.Command
	downloadCmd         *cobra.Command
	downloadResourceCmd *cobra.Command
}

func NewCli() *Cli {
	return &Cli{}
}

func (cli *Cli) Execute() error {
	cli.buildCmds()
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()
	errChan := make(chan error, 1)
	go func() {
		errChan <- cli.rootCmd.ExecuteContext(ctx)
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case err := <-errChan:
		return err
	case <-signalChan:
		log.Warn().Msg("Exiting early")
		return nil
	}
}

func (cli *Cli) buildCmds() {
	cli.rootCmd = &cobra.Command{
		Use:               "go-bif-examine-cli",
		Short:             "An overcomplicated way of extracting audio from BIF files",
		PersistentPreRunE: cli.globalConfigure,
	}
	cli.rootCmd.PersistentFlags().StringVarP(&cli.logLevel, "log-level", "l", "info", "Zerolog log level to use; trace, debug, info, warn, error, panic, etc")
	cli.rootCmd.PersistentFlags().StringVarP(&cli.grpcServer, "grpc", "g", "localhost:50051", "IP:Port of gRPC server")
	cli.buildGetCmds()
	cli.buildDeleteCmds()
	cli.buildUploadCmds()
	cli.buildDownloadCmds()
}

func (cli *Cli) globalConfigure(cmd *cobra.Command, args []string) error {
	util.ConfigureLogging(cli.logLevel, applicationName, applicationDescription)
	return nil
}

func (cli *Cli) grpcConfigure(cmd *cobra.Command, args []string) error {
	conn, err := grpc.Dial(
		cli.grpcServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	cli.grpcClient = pb.NewBifExamineClient(conn)
	return nil
}
