package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/sky0621/work-go-movies/update"
)

// 【update概要】
// 指定のパスにある、新規にアップロードされた動画の情報が記載されたファイル（watch.list）を読み取り、
// persistenceが参照するJSONに足しこむ。
// 足しこみが終わったら watch.list を削除する。
func main() {
	log.Println("[main]START")
	var watchListDir string
	var storageAddr string
	var sleep time.Duration
	var logDir string

	flag.StringVar(&watchListDir, "w", "list", "watch.list格納ディレクトリ")
	flag.StringVar(&storageAddr, "s", "movies-persistence.json", "ストレージサーバ接続先アドレス")
	flag.DurationVar(&sleep, "t", 600, "監視間隔（秒）")
	flag.StringVar(&logDir, "l", ".", "ログ出力先ディレクトリ")
	flag.Parse()
	log.Println("[main]flag parse fin.")

	logfile, err := update.SetupLog(logDir)
	if err != nil {
		os.Exit(1)
	}
	defer logfile.Close()
	log.Println("[main]setup log fin.")

	log.Println("[main]args")
	log.Printf("[main] -> watch.list格納ディレクトリ「%s」\n", watchListDir)
	log.Printf("[main] -> ストレージサーバ接続先アドレス「%s」\n", storageAddr)
	log.Printf("[main] -> 監視間隔（秒）「%d」\n", sleep)
	log.Printf("[main] -> ログ出力先ディレクトリ「%s」\n", logDir)

	log.Println("[main]go loop!!!")
	for {
		if update.IsRunning {
			log.Println("[main]update is running... ... ...")
		} else {
			go update.Run(watchListDir, storageAddr)
		}

		time.Sleep(sleep * time.Second)
	}
}
