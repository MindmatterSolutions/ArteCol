package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const commandName = "this"

var usageText = fmt.Sprintf(`Usage:
  %s [path]

Examples:
  %s
  %s ./path/to/artefact

When no path is provided the current working directory is used.
`, commandName, commandName, commandName)

type thisCommand struct {
	targetPath string
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func run(args []string) error {
	this, err := resolveThis(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageText)
		return err
	}

	if err := dispatch(this); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	return nil
}

func resolveThis(args []string) (thisCommand, error) {
	switch len(args) {
	case 0:
		cwd, err := os.Getwd()
		if err != nil {
			return thisCommand{}, fmt.Errorf("determine working directory: %w", err)
		}
		return thisCommand{targetPath: cwd}, nil
	case 1:
		abs, err := filepath.Abs(args[0])
		if err != nil {
			return thisCommand{}, fmt.Errorf("resolve target path %q: %w", args[0], err)
		}
		return thisCommand{targetPath: abs}, nil
	default:
		return thisCommand{}, fmt.Errorf("invalid invocation: expected 0 or 1 arguments, got %d", len(args))
	}
}

func dispatch(this thisCommand) error {
	fmt.Printf("Dispatching artefact resources for %s\n", this.targetPath)
	return nil
}
