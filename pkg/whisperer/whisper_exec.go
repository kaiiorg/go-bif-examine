package whisperer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func (w *Whisperer) runWhisperer(audioFile string) (string, error) {
	args := []string{
		w.whisperPath,
		audioFile,
		"--model", whisper_model,
		"--output_format", "json",
		"--output_dir", w.tempDir,
		"--language", "en",
	}
	log.Debug().Strs("args", args).Msg("Attempting to run whisper")

	cmd := exec.CommandContext(w.ctx, w.whisperPath, args...)

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	wavName := filepath.Base(audioFile)
	name := strings.TrimSuffix(wavName, filepath.Ext(wavName))
	jsonName := fmt.Sprintf("%s.json", name)

	return filepath.Join(w.tempDir, jsonName), nil
}

func (w *Whisperer) extractOutput(outputFile string) (*WhisperOutput, []byte, error) {
	contents, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, nil, err
	}

	whisperOutput := &WhisperOutput{Segments: []WhisperOutputSegments{}}

	err = json.Unmarshal(contents, whisperOutput)
	if err != nil {
		return nil, nil, err
	}

	return whisperOutput, contents, nil
}
