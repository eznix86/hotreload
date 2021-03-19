package args

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

const ReloadTime = 1500
const DefaultPort = 9023

type pathList []string

func (p *pathList) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *pathList) Set(value string) error {
	*p = append(*p, value)
	return nil
}

type Args struct {
	Paths pathList
	ReloadTimeInMilliseconds time.Duration
	ReloadPort int
	Verbose bool
}

func Get() (*Args, error) {

	var config = Args{}

	flag.Var(&config.Paths, "path", "Paths to watch file changes\n--path /to/folder [--path /to/folder]")
	var reload = flag.Int("duration", ReloadTime, "Duration after which page will reload in ms")
	var reloadPort = flag.Int("port", DefaultPort, "Default port to listen for reload requests")
	var verbose = flag.Bool("verbose", false, "verbose")

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		return nil, errors.New("no parameter found")
	}

	flag.Parse()

	if *reload < 100 {
		return nil, errors.New("reload time cannot be less than 100 ms")
	}


	config.ReloadTimeInMilliseconds = time.Duration(*reload)
	config.ReloadPort = *reloadPort
	config.Verbose = *verbose

	return &config, nil
}
