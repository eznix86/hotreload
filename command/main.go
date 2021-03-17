package main

import (
	"fmt"
	"github.com/theArtechnology/hotreload/args"
	"github.com/theArtechnology/hotreload/notifier"
	"github.com/theArtechnology/hotreload/websocket"
	"os"
)




func main() {
	var config, err = args.Get()

	if err != nil {
		fmt.Println()
		fmt.Println(err)
		os.Exit(1)
	}

	var changeNotifier = notifier.New(config.ReloadTimeInMilliseconds, config.Paths)

	changeNotifier.Start()
	hotReloadHandler := websocket.HotReloadHandler{}

	changeNotifier.AddListener(&hotReloadHandler)
	hotReloadHandler.Serve()
}
