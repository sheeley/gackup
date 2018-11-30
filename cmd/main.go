package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sheeley/gackup"
)

var files = []string{
	".mongorc.js",
	".gitident-work",
	".zshrc",
	".ssh/config",
	".gitignore",
	".gitconfig",
	".zshenv",
	// ".vscode",
	"Library/Preferences/com.surteesstudios.Bartender.plist",
	"Library/Preferences/com.googlecode.iterm2.plist",
	"Library/Preferences/info.marcel-dierkes.KeepingYouAwake.plist",
	// "Library/KeyBindings/DefaultKeyBinding.dict",
	// "Library/Services",
	// "Library/Speech/Speakable Items",
	// "Library/Scripts",
	// "Library/Workflows",
	// "Library/PDF Services",
	"Library/Preferences/com.apple.symbolichotkeys.plist",
	"Library/Preferences/org.shiftitapp.ShiftIt.plist",
	"Library/Application Support/Code/User/settings.json",
}

func main() {
	config := gackup.DefaultConfig

	flag.StringVar(&config.BaseDir, "base", config.BaseDir, "Set base directory")
	flag.StringVar(&config.ConfigDir, "configDir", config.ConfigDir, "Set directory to store synced files in")
	flag.BoolVar(&config.Verbose, "verbose", config.Verbose, "")
	flag.BoolVar(&config.ForceRelink, "relink", config.ForceRelink, "Force re-linking of all files")
	flag.Parse()

	if config.Verbose {
		fmt.Printf("%+v\n", config)
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
