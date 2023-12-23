package whisperer

import (
	"os/exec"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (w *Whisperer) checkWhisperAvailability(programNameInPath string) error {
	var err error
	w.whisperPath, err = exec.LookPath(programNameInPath)
	if err != nil {
		w.log.Error().Err(err).Str("github", whisper_github).Msg("Unable to validate that whisper is installed. See its github page for install instructions")
		return err
	}
	w.log.Info().Str("path", w.whisperPath).Msg("Found whisper install!")
	return nil
}

func (w *Whisperer) dialGRpc(gRpcServer string) error {
	// TODO use a secure connection method
	conn, err := grpc.Dial(
		gRpcServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	w.gRpcClient = pb.NewWhispererClient(conn)
	return nil
}
