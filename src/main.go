package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/commands"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	command := os.Args[1]
	flagSet.Parse(os.Args[2:])

	var commandFunc func(conf *config.Config, flagSet *flag.FlagSet) error
	switch command {
	case "build":
		commandFunc = build
	case "ls":
		commandFunc = ls
	case "append":
		commandFunc = append
	case "symlink":
		commandFunc = symlink
	default:
		panic(fmt.Sprintf("Unexpected command. command=%v", command))
	}

	err := commandFunc(config.GetConfig(), flagSet)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
}

func build(conf *config.Config, flagSet *flag.FlagSet) error {
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

func ls(conf *config.Config, flagSet *flag.FlagSet) error {
	vindex, err := commands.ListIndex(conf.VIndexPath)
	if err != nil {
		return err
	}

	for i := range vindex {
		item := vindex[i]
		fmt.Printf("%v %v %v %v\n", item.HexDigest(), item.Storage, item.Date, filepath.Base(item.Filename))
	}

	return nil
}

func append(conf *config.Config, flagSet *flag.FlagSet) error {
	args := flagSet.Args()
	path := args[0]
	if path == "" {
		return errors.Errorf("Invalid filename")
	}

	vindexItem, err := vindexdata.ParseFilepath(conf.StoragePaths, path)
	if err != nil {
		return errors.Errorf("ParseFilepath failed. path=%v", path)
	}

	err = commands.Append(conf, vindexItem)
	if err != nil {
		return err
	}

	fmt.Println(vindexItem.HexDigest())

	return nil
}

func symlink(conf *config.Config, flagSet *flag.FlagSet) error {
	args := flagSet.Args()
	digest := args[0]
	if digest == "" {
		return errors.Errorf("Invalid digest")
	}

	err := commands.Symlink(conf, digest)
	if err != nil {
		return err
	}

	return nil
}
