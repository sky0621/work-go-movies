package client

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// SetupLog ...
func SetupLog(outputDir string) (*os.File, error) {
	logfile, err := os.OpenFile(filepath.Join(outputDir, "movies-client.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("[%s]のログファイル「movies-client.log」オープンに失敗しました。 [ERROR]%s\n", outputDir, err)
		return nil, err
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	return logfile, nil
}
