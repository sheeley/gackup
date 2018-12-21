package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sheeley/gackup"
)

func main() {
	config := gackup.DefaultConfig

	flag.StringVar(&config.SourceDir, "source", config.SourceDir, "Set source directory")
	flag.StringVar(&config.TargetDir, "target", config.TargetDir, "Set directory to store synced files in")
	flag.BoolVar(&config.Verbose, "verbose", config.Verbose, "")
	flag.BoolVar(&config.ForceRelink, "relink", config.ForceRelink, "Force re-linking of all files")
	flag.Parse()

	if config.Verbose {
		fmt.Printf("%+v\n", config)
	}

	files, err := gackup.LoadFileList(config)
	if err != nil {
		panic(err)
	}

	b, err := gackup.New(files, config)
	if err != nil {
		panic(err)
	}

	proposed, err := b.Proposed()
	if err != nil {
		fmt.Println(proposed)
		panic(err)
	}

	if len(proposed) == 0 {
		os.Exit(0)
	}

	fmt.Printf("%s\nConfirm [y/N]:", proposed)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(input)) == "y" {
		_, err := b.Move()
		if err != nil {
			panic(err)
		}
	}
}
