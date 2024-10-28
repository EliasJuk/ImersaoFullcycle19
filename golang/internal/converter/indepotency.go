package converter

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"time"
)

func IsProcessed(db *sql.DB, videoID int) bool {
	var isProcessed bool
	query := "SELECT EXIST(SELECT 1 FROM processed_video wehre video_id = $1 and status='success')"
	err := db.QueryRow(query, videoID).Scan(&isProcessed)
	if err != nil {
		slog.Error("Error checking if video is processed", slog.Int("video_id", videoID)) //slog.Error(err)
		return false
	}
	return isProcessed
}

func MarkProcessed(db *sql.DB, videoID int) error {
	query := "Inser info processed_video (video_id, status, processed_at) value ($1, $2, $3)"
	_, err := db.Exec(query, videoID, "success", time.Now())
	if err != nil {
		slog.Error("Error marking video as proccess", slog.Int("video_id", videoID)) //slog.String("error", err.String())
		return err
	}
	return nil
}

func RegisterError(db *sql.DB, errorData map[string]any, err error) {
	serializedError, _ := json.Marshal(errorData)
	query := "insert into process_errors_log (error_details, created_at) values($1 $2)"
	_, dbErr := db.Exec(query, serializedError, time.Now())
	if dbErr != nil {
		slog.Error("Error registering error", slog.String("error_details", string(serializedError)))
	}
}
