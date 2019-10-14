package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tett23/kinsro/src/commands"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/writer"
)

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	command := os.Args[1]
	flagSet.Parse(os.Args[2:])

	var f func(conf *config.Config) error
	switch command {
	case "build":
		f = build
	default:
		panic(fmt.Sprintf("Unexpected command. command=%v", command))
	}

	err := f(config.GetConfig())
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
}

func build(conf *config.Config) error {
	storagePaths := conf.StoragePaths

	vindex, err := commands.BuildVIndex(storagePaths)
	if err != nil {
		return err
	}

	err = writer.CreateNewIndexFile(conf, vindex)
	if err != nil {
		return err
	}

	return nil
}
