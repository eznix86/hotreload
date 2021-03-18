package notifier

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"time"
)

type Listener interface {
	Reload()
}

type Notifier struct {
	ReloadTime time.Duration
	WatchPaths []string
	Listener  Listener
}

func New(reloadTime time.Duration, watchPaths []string) *Notifier {
	fmt.Printf("Reload time set to %v\n", reloadTime)
	fmt.Printf("Watch Paths are %v\n", watchPaths)
	return &Notifier {
		reloadTime,
		watchPaths,
		nil,
	}
}

func (n *Notifier) AddListener(l Listener)  {
	n.Listener = l
}

func (n *Notifier) Start() {
	w := watcher.New()
	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)
	w.IgnoreHiddenFiles(true)
	// Only notify rename and move events.
	w.FilterOps(watcher.Write, watcher.Create, watcher.Rename)

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
				if n.Listener != nil {
					n.Listener.Reload()
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch folders recursively for changes.
	for _, path := range n.WatchPaths {
		if err := w.AddRecursive(path); err != nil {
			log.Fatalln(err)
		}
	}


	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()


	go func() {
		// Start the watching process - it'll check for changes every time
		if err := w.Start(time.Millisecond * n.ReloadTime); err != nil {
			log.Fatalln(err)
		}
	}()
}