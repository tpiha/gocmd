package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type CliParser struct {
	Command string
}

func (c *CliParser) parseCli() {
	help := flag.Bool("h", false, "")
	helpLong := flag.Bool("help", false, "")

	dev := flag.Bool("d", false, "")
	devLong := flag.Bool("dev", false, "")

	flag.Parse()

	if *help || *helpLong {
		c.printHelpMessage()
		os.Exit(0)
	}

	if *dev || *devLong {
		conf.Dev = true
	}

	for _, arg := range os.Args {
		if arg == "default" || arg == "worker" || arg == "server" {
			c.Command = arg
		}

		if strings.Contains(arg, "-d") || strings.Contains(arg, "--dev") {
			if !*dev && !*devLong {
				conf.Dev = false
			}
		}
	}
}

func (c *CliParser) printHelpMessage() {
	helpMessage := `gocmd ` + VERSION + ` - (C) 2017 Tihomir Piha

Usage: gocmd [COMMAND] [FLAGS]

Commands:

    default  - run in server mode together with worker
    worker   - run worker without server
    server   - run in server mode without worker

Flags:

    -d,   --dev     run in development mode (more logging)
    -h,   --help    print this message
`
	fmt.Println(helpMessage)
}

func initCli() *CliParser {
	clip := &CliParser{Command: "default"}
	clip.parseCli()
	return clip
}
