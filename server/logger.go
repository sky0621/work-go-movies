package server

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// SetupLog ...
func SetupLog(outputDir string) (*os.File, error) {
	logfile, err := os.OpenFile(filepath.Join(outputDir, "movies-server.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("[%s]のログファイル「movies-server.log」オープンに失敗しました。 [ERROR]%s\n", outputDir, err)
		return nil, err
	}

	// [MEMO]内容に応じて出力するファイルを切り替える場合はどうするんだ・・・？
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	return logfile, nil
}

// [MEMO]セットアップ時に「log.Lshortfile」セットしたかったけど、このファイルでログ出力を担うようにすると全ログの「log.Lshortfile」結果が「logger.go」になるので諦め
// TODO ロギングフレームワークを採用する！　logrusあたりがメジャーらしい。最近では（ベータ版だけど）zapがいいらしい。
type logger struct {
	isDebugEnable bool
}

func (l *logger) debug(fname string, v ...interface{}) {
	if l.isDebugEnable {
		log.Println("[DEBUG]["+fname+"]", v)
	}
}

func (l *logger) debugf(fname string, formatStr string, v ...interface{}) {
	if l.isDebugEnable {
		log.Printf("[DEBUG]["+fname+"] "+formatStr, v)
	}
}

func (l *logger) info(fname string, v ...interface{}) {
	log.Println("[INFO]["+fname+"]", v)
}

func (l *logger) infof(fname string, formatStr string, v ...interface{}) {
	log.Printf("[INFO]["+fname+"] "+formatStr, v)
}

func (l *logger) error(fname string, v ...interface{}) {
	log.Println("[ERROR]["+fname+"]", v)
}

func (l *logger) errorf(fname string, formatStr string, v ...interface{}) {
	log.Printf("[ERROR]["+fname+"] "+formatStr, v)
}
