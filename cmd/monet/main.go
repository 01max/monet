package main

import (
	"fmt"
	"os"

	"github.com/01max/monet/internal/recipe"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("hello monet")
		return
	}

	switch os.Args[1] {
	case "ls":
		runList()
	case "debug":
		runDebug()
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func runList() {
	recipes, err := recipe.LoadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	for _, r := range recipes {
		fmt.Printf("%s\t%s\n", r.Name, r.Description)
	}
}

func runDebug() {
	fmt.Fprintf(os.Stdout, "Debug info for monet\n")
	fmt.Fprintf(os.Stdout, "Arguments: %v\n", os.Args)
	fmt.Fprintf(os.Stdout, "Config dir (XDG_CONFIG_HOME): %s\n", os.Getenv("XDG_CONFIG_HOME"))
	var confDir, confDirErr = os.UserConfigDir()
	if confDirErr != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", confDirErr)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Config dir (os.UserConfigDir()): %s\n", confDir)
	recipe.DebugRecipesDir()
}
