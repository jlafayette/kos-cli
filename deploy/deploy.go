package deploy

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

// Deploy .ks file from development path to KSP install path
func Deploy(kspsrc, kspscript string, verbose bool) error {
	copyInfo := []copyInstructions{
		{filepath.Join(kspsrc, "library"), kspscript, "*.ks"},
		{filepath.Join(kspsrc, "missions"), filepath.Join(kspscript, "missions"), "*.ks"},
		{filepath.Join(kspsrc, "boot"), filepath.Join(kspscript, "boot"), "*.ks"},
		{filepath.Join(kspsrc, "working", "missions"), filepath.Join(kspscript, "missions"), "*.ks"},
		{filepath.Join(kspsrc, "working", "boot"), filepath.Join(kspscript, "boot"), "*.ks"},
	}
	for _, info := range copyInfo {
		err := os.MkdirAll(info.dest, os.ModePerm)
		if err != nil {
			return err
		}
		err = CopyFiles(info.src, info.dest, info.match, verbose)
		if err != nil {
			return err
		}
	}
	return nil
}

// CopyFiles copies the files matching a pattern string in a source folder to
// a destination folder.
func CopyFiles(src, dest, match string, verbose bool) error {
	if verbose {
		fmt.Printf("\nCopy .ks files from '%s' to '%s'\n", src, dest)
	}
	files, err := filepath.Glob(filepath.Join(src, match))
	if err != nil {
		return err
	}
	for _, f := range files {
		err = CopyFile(f, filepath.Join(dest, filepath.Base(f)), verbose)
		if err != nil {
			return err
		}
	}
	return nil
}

// CopyFile copies the contents of the file named src to the file named by dest.
// The file will be created if it does not already exist. If the destination
// file exists, all it's contents will be replaced.
func CopyFile(src, dest string, verbose bool) error {
	if verbose {
		fmt.Printf("Copy %s to %s\n", src, dest)
	}
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
