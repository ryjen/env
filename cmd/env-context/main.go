package main

import (
	"flag"
	"os"
	"os/exec"

	"github.com/ryjen/env/context"
	"github.com/ryjen/env/files"
)

var (
	directory = flag.String("dir", "", "the directory to load environment files from")
	overload  = flag.Bool("overload", false, "overload existing OS environment variables")
	envFile   = flag.String("env-file", "", "specify environment file name")
	localExt  = flag.String("local-ext", "", "specify local extension")
	contextId = flag.String("context", "", "specify the context as a string or name of a environment variable")
)

func main() {

	flag.Parse()

	// flags to configuration

	if envFile != nil && len(*envFile) > 0 {
		files.SetEnvFileName(*envFile)
	}

	if localExt != nil && len(*localExt) > 0 {
		files.SetLocalExt(*localExt)
	}

	if contextId != nil && len(*contextId) > 0 {
		context.Set(*contextId)
	}

	if directory != nil && len(*directory) > 0 {
		if overload != nil && *overload {
			files.OverloadFromDir(*directory)
		} else {
			files.LoadFromDir(*directory)
		}
	} else {
		if overload != nil && *overload {
			files.Overload()
		} else {
			files.Load()
		}
	}

	if flag.NArg() == 0 {
		os.Exit(0)
	}

	// if a command is specified run it

	args := flag.Args()
	prog, args := args[0], append(args[:0], args[1:]...)

	cmd := exec.Command(prog)
	cmd.Args = args
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
