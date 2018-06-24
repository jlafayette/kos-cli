package deploy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// KsMatch matches .ks files when passed to filepath.Glob.
const KsMatch = "*.ks"

// CopyInstructions stores info required to copy or sync files for deployment.
type CopyInstructions struct {
	Src   string
	Dst   string
	Match string
}

// GetInstructions returns a slice of instructions for deployment based on source and destination root.
func GetInstructions(src, dst, match string) []CopyInstructions {
	copyInfo := []CopyInstructions{
		{filepath.Join(src, "library"), dst, match},
		{filepath.Join(src, "missions"), filepath.Join(dst, "missions"), match},
		{filepath.Join(src, "boot"), filepath.Join(dst, "boot"), match},
	}
	return copyInfo
}

// Deploy .ks file from development path to KSP install path
func Deploy(kspsrc, kspscript string, verbose bool) error {
	copyInfo := GetInstructions(kspsrc, kspscript, KsMatch)
	for _, info := range copyInfo {
		err := os.MkdirAll(info.Dst, os.ModePerm)
		if err != nil {
			return err
		}
		err = CopyFiles(info.Src, info.Dst, info.Match, verbose)
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
