package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MindmatterSolutions/go-this/pkg/artefact"
)

const commandName = "this"

var usageText = fmt.Sprintf(`Usage:
  %s [path]

Examples:
  %s
  %s ./path/to/artefact

When no path is provided the current working directory is used.
`, commandName, commandName, commandName)

// thisCommand contains the data necessary to execute the `this` command.
type thisCommand struct {
	targetPath string
}

// main is the entry point for the `this` command binary.
func main() {
	if err := run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

// run parses CLI arguments and dispatches the artefact resolver.
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

// resolveThis returns a populated thisCommand from the provided args slice.
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

// dispatch calls the file-type specific logic for the resolved path.
func dispatch(this thisCommand) error {
	info, err := os.Stat(this.targetPath)
	if err != nil {
		return fmt.Errorf("stat target path %q: %w", this.targetPath, err)
	}

	if info.IsDir() {
		entries, err := artefact.HandleDirectory(this.targetPath)
		if err != nil {
			return err
		}
		fmt.Printf("Directory %s contains %d entries\n", this.targetPath, len(entries))
		return nil
	}

	if info.Mode().IsRegular() {
		data, err := artefact.HandleFile(this.targetPath)
		if err != nil {
			return err
		}
		fmt.Printf("File %s has %d bytes\n", this.targetPath, len(data))
		return nil
	}

	return artefact.HandleUnknown(this.targetPath)
}
