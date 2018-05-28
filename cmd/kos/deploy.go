package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type copyInstructions struct {
	src   string
	dest  string
	match string
}

func deploy(kspsrc, kspscript string) error {
	fmt.Println("running deploy function...")
	copyInfo := []copyInstructions{
		{filepath.Join(kspsrc, "library"), kspscript, "*.ks"},
		{filepath.Join(kspsrc, "missions"), filepath.Join(kspscript, "missions"), "*.ks"},
		{filepath.Join(kspsrc, "boot"), filepath.Join(kspscript, "boot"), "*.ks"},
	}
	for _, info := range copyInfo {
		err := os.MkdirAll(info.dest, os.ModePerm)
		if err != nil {
			return err
		}
		err = cpFiles(info.src, info.dest, info.match)
		if err != nil {
			return err
		}
	}
	return nil
}

func cpFiles(src, dest, match string) error {
	fmt.Printf("\nCopy .ks files from '%s' to '%s'\n", src, dest)
	// files, err := ioutil.ReadDir(src)
	files, err := filepath.Glob(filepath.Join(src, match))
	if err != nil {
		return err
	}
	for _, f := range files {
		// fmt.Println(f.Name())
		// fmt.Println(f)
		err = cpFile(f, filepath.Join(dest, filepath.Base(f)))
		if err != nil {
			return err
		}
	}
	return nil
}

// cpFile copies the contents of the file named src to the file named by dest.
// The file will be created if it does not already exist. If the destination
// file exists, all it's contents will be replaced.
func cpFile(src, dest string) error {
	fmt.Printf("Copy %s to %s\n", src, dest)
	infile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer infile.Close()
	outfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer outfile.Close()
	_, err = io.Copy(outfile, infile)
	if err != nil {
		return err
	}
	return nil
}
