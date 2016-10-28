package update

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// SetupLog ...
func SetupLog(logDir string) (*os.File, error) {
	logfile, err := os.OpenFile(filepath.Join(logDir, "movies-update.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("[%s]のログファイル「update.log」オープンに失敗しました。 [ERROR]%s\n", logDir, err)
		return nil, err
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	return logfile, nil
}
