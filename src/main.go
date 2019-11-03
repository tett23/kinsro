package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/commands"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/syscalls"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

func init() {
	config.Initialize()
}

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	command := os.Args[1]
	flagSet.Parse(os.Args[2:])
	log.SetOutput(os.Stdout)

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
	case "rebuild":
		commandFunc = rebuild
	case "encode":
		commandFunc = encodeTS
	case "encode-all":
		commandFunc = encodeTSAll
	default:
		panic(fmt.Sprintf("Unexpected command. command=%v", command))
	}

	err := commandFunc(config.GetConfig(), flagSet)
	if err != nil {
		log.Printf("%+v\n", err)
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
	fs := filesystem.GetFs()
	vindex, err := commands.ListIndex(fs, conf.VIndexPath)
	if err != nil {
		return err
	}

	for i := range vindex {
		item := vindex[i]
		log.Printf("%v %v %v %v\n", item.HexDigest(), item.Storage, item.Date, filepath.Base(item.Filename))
	}

	return nil
}

func append(conf *config.Config, flagSet *flag.FlagSet) error {
	args := flagSet.Args()
	path := args[0]
	if path == "" {
		return errors.Errorf("Invalid filename")
	}

	vindexItem, err := vindexdata.ParseFullFilepath(conf.StoragePaths, path)
	if err != nil {
		return errors.Errorf("ParseFilepath failed. path=%v", path)
	}

	fs := filesystem.GetFs()
	err = commands.AppendToIndex(conf, fs, vindexItem)
	if err != nil {
		return err
	}

	log.Println(vindexItem.HexDigest())

	return nil
}

func symlink(conf *config.Config, flagSet *flag.FlagSet) error {
	args := flagSet.Args()
	if len(args) == 0 {
		return symlinkAll(conf, flagSet)
	}

	digest := args[0]
	fs := filesystem.GetFs()
	err := commands.Symlink(conf, syscalls.NewOSSyscalls(), fs, digest)
	if err != nil {
		return err
	}

	return nil
}

func symlinkAll(conf *config.Config, flagSet *flag.FlagSet) error {
	fs := filesystem.GetFs()
	err := commands.SymlinkAll(conf, syscalls.NewOSSyscalls(), fs)
	if err != nil {
		return err
	}

	return nil
}

func rebuild(conf *config.Config, flagSet *flag.FlagSet) error {
	err := build(conf, flagSet)
	if err != nil {
		errors.Cause(err)
	}

	err = symlinkAll(conf, flagSet)
	if err != nil {
		errors.Cause(err)
	}

	return nil
}

func encodeTS(conf *config.Config, flagSet *flag.FlagSet) error {
	args := flagSet.Args()
	tsPath := args[0]
	fs := filesystem.GetFs()
	options := commands.EncodeOptions{
		RemoveTS: false,
	}
	err := commands.EncodeTS(conf, syscalls.NewOSSyscalls(), fs, tsPath, options)
	if err != nil {
		return err
	}

	return nil
}

func encodeTSAll(conf *config.Config, flagSet *flag.FlagSet) error {
	args := flagSet.Args()
	videoTmpPath := args[0]
	fs := filesystem.GetFs()
	options := commands.EncodeOptions{
		RemoveTS: false,
	}
	err := commands.EncodeTSAll(conf, syscalls.NewOSSyscalls(), fs, videoTmpPath, options)
	if err != nil {
		return err
	}

	return nil
}
