package main

import (
	"github.com/theArtechnology/hotreload/args"
	"github.com/theArtechnology/hotreload/notifier"
	"github.com/theArtechnology/hotreload/websocket"
	"log"
	"os"
)

func main() {
	var config, err = args.Get()

	if err != nil {
		log.Println()
		log.Println(err)
		os.Exit(1)
	}

	var changeNotifier = notifier.New(config.ReloadTimeInMilliseconds, config.Paths)

	hotReloadHandler := websocket.HotReloadHandler{}
	changeNotifier.AddListener(&hotReloadHandler)

	changeNotifier.Verbose = config.Verbose
	hotReloadHandler.ServerPort = config.ReloadPort

	changeNotifier.Start()
	hotReloadHandler.Serve()
}
