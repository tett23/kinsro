package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tett23/kinsro/src/commands"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/writer"
)

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	command := os.Args[1]
	flagSet.Parse(os.Args[2:])

	var commandFunc func(conf *config.Config) error
	switch command {
	case "build":
		commandFunc = build
	case "ls":
		commandFunc = ls
	default:
		panic(fmt.Sprintf("Unexpected command. command=%v", command))
	}

	err := commandFunc(config.GetConfig())
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

func ls(conf *config.Config) error {
	vindex, err := commands.ListIndex(conf.VIndexPath)
	if err != nil {
		return err
	}

	for i := range vindex {
		item := vindex[i]
		fmt.Printf("%v\t%v\t%v\n", item.Storage, item.Date, filepath.Base(item.Filename))
	}

	return nil
}
