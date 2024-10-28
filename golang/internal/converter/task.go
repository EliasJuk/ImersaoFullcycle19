package converter

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type VideoConverter struct{}

func NewVideoConverter() *VideoConverter {
	return &VideoConverter{}
}

// {"video_id": 1, "path": "media/uploads/1"}
type VideoTask struct {
	VideoID int    `json:"video_id"`
	Path    string `json:"path"`
}

func (vc *VideoConverter) Handle(msg []byte) {
	var task VideoTask
	err := json.Unmarshal(msg, &task)
	if err != nil {
		vc.logError(task, "failed to unmasrshal task", err)
		return
	}

	err = vc.processVideo(&task)
	if err != nil {
		vc.logError(task, "failed to unmarshal task", err)
		return
	}
}

func (vc *VideoConverter) processVideo(task *VideoTask) error {
	mergedFile := filepath.Join(task.Path, "merged.mp4")
	mpegDashPath := filepath.Join(task.Path, "mpeg-dash")

	slog.Info("Mergin chunks", slog.String("path", task.Path))
	err := vc.mergeChunks(task.Path, mergedFile)
	if err != nil {
		vc.logError(*task, "failed to unmasrshal task", err)
		//return fmt.Errorf("failed to merge chunks: %v", err)
		return err
	}

	// CRIA O DIRETORIO MPEG-DASH
	slog.Info("Creating mpeg-dash", slog.String("path", task.Path))
	err = os.MkdirAll(mpegDashPath, os.ModePerm)
	if err != nil {
		vc.logError(*task, "failed to create mpeg-dash directory", err)
		return err
	}

	// CONVERTE O VIDEO
	slog.Info("Converting video to mpeg-dash", slog.String("path", task.Path))
	ffmpegCmd := exec.Command(
		"ffmpeg", "-i", mergedFile,
		"-f", "dash",
		filepath.Join(mpegDashPath, "output.mpd"),
	)

	// TRATAR ERROS
	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
		vc.logError(*task, "failed to convert to mpeg-dash, output: "+string(output), err)
		return err
	}
	slog.Info("Video converted to mpeg-dash", slog.String("path", mpegDashPath))

	// REMOVE O ARQUIVO APOS CONVERTER
	slog.Info("Removing merged file", slog.String("path", mergedFile))
	err = os.Remove((mergedFile))
	if err != nil {
		vc.logError(*task, "failed to remove merged file", err)
	}

	return nil
}

// PADRONIZA OUTPUTS DE ERROS
func (vc *VideoConverter) logError(task VideoTask, message string, err error) {
	errorData := map[string]any{
		"video_id": task.VideoID,
		"error":    message,
		"details":  err.Error(),
		"time":     time.Now(),
	}
	serializedError, _ := json.Marshal(errorData)
	slog.Error("processing error", slog.String("error_details", string(serializedError)))

	// Todo: register error on database

}

func (vc *VideoConverter) extractNumber(fileName string) int {
	re := regexp.MustCompile(`\d+`)
	numStr := re.FindString(filepath.Base(fileName))
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}
	return num
}

func (vc *VideoConverter) mergeChunks(inputDir, outputFile string) error {
	chunks, err := filepath.Glob(filepath.Join(inputDir, "*.chunk"))
	if err != nil {
		return fmt.Errorf("fail tp find chunks %v", err)
	}

	sort.Slice(chunks, func(i, j int) bool {
		return vc.extractNumber(chunks[i]) < vc.extractNumber(chunks[j])
	})

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	for _, chunk := range chunks {
		input, err := os.Open(chunk)
		if err != nil {
			return fmt.Errorf("failed to open chunk: %v", err)
		}

		_, err = output.ReadFrom(input)
		if err != nil {
			return fmt.Errorf("failed to read chunk %s to merge file: %v", chunk, err)
		}
		input.Close()
	}

	return nil
}
