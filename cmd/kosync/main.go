package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jlafayette/kos-cli/deploy"
)

const help = `kosync continuously syncs a source folder to a destination folder.

Usage: kosync [options]

Options:
	-f --filter   Glob string for files to be synced. If not specified,
	              defaults to '*.ks' which matches KerboScript files.
	-s --source   Development root folder to sync from. If not specified,
	              defaults to $KSPSRC environment variable.
	-d --dest     KSP Ships/Script folder to sync to. If not specified,
	              defaults to $KSPSCRIPT environment variable.
`

func main() {
	filterPtr := flag.String("filter", deploy.KsMatch, "Glob string for files to be synced.")
	fPtr := flag.String("f", "", "(Alias for filter) Glob string for files to be synced.")
	sourcePtr := flag.String("source", "", "Development root folder to sync from.")
	sPtr := flag.String("s", os.Getenv("KSPSRC"), "(Alias for source) Development root folder to sync from.")
	destPtr := flag.String("dest", "", "KSP Ships/Script folder to sync to.")
	dPtr := flag.String("d", os.Getenv("KSPSCRIPT"), "(Alias for dest) KSP Ships/Script folder to sync to.")
	flag.Parse()

	if len(flag.Args()) != 0 {
		fmt.Printf("args: %v", flag.Args())
		fmt.Fprintln(os.Stderr, help)
		os.Exit(2)
	}

	filter := *filterPtr
	if *fPtr != "" && filter == deploy.KsMatch {
		filter = *fPtr
	}
	var src string
	if *sourcePtr != "" {
		src = *sourcePtr
	} else if *sPtr != "" {
		src = *sPtr
	} else {
		fmt.Fprintf(os.Stderr, "ERROR: Missing required flag: --source\n\n")
		fmt.Fprintln(os.Stderr, help)
		os.Exit(2)
	}
	var dst string
	if *destPtr != "" {
		dst = *destPtr
	} else if *dPtr != "" {
		dst = *dPtr
	} else {
		fmt.Fprintf(os.Stderr, "ERROR: Missing required flag: --dest\n\n")
		fmt.Fprintln(os.Stderr, help)
		os.Exit(2)
	}
	fmt.Printf("filter: %v\n", filter)
	fmt.Printf("   src: %v\n", src)
	fmt.Printf("   dst: %v\n", dst)

	done := make(chan bool)
	deployInfo := deploy.GetInstructions(src, dst, filter)
	for _, info := range deployInfo {
		go func(info deploy.CopyInstructions) {
			err := sync(info.Src, info.Dst, info.Match)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}(info)
	}
	<-done
}

func sync(src, dst, filter string) error {
	fmt.Printf("Starting sync %v -> %v\n", src, dst)

	// Initial sync
	err := deploy.CopyFiles(src, dst, filter, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
		return err
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				_, file := filepath.Split(event.Name)
				m, _ := filepath.Match(filter, file)
				if m {
					switch event.Op {
					case 1: // Create Op  = 1 << iota
						newFile, err := os.Create(replaceFirst(event.Name, src, dst))
						if err != nil {
							fmt.Printf("ERROR creating file: %v\n", err)
						}
						newFile.Close()
					case 2: // Write --> copy
						err := copyFile(event.Name, replaceFirst(event.Name, src, dst), 500)
						if err != nil {
							fmt.Printf("ERROR copying file: %v\n", err)
						}
					case 4: // Remove --> remove
						err := os.Remove(replaceFirst(event.Name, src, dst))
						if err != nil {
							fmt.Printf("ERROR removing file: %v\n", err)
						}
					case 8: // Rename --> remove
						err := os.Remove(replaceFirst(event.Name, src, dst))
						if err != nil {
							fmt.Printf("ERROR removing file: %v\n", err)
						}
					case 16: // Chmod --> ignore
						// put something here if Chmod needs to be handled.
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("WatcherError:", err)
				// ToDo: Handle 'short read in readEvents()' error.
				// This happens when the event buffer overflows and it means that some
				// events have been missed and the folders and now out of sync. Need
				// to re-sync and then start handling events again.
			}
		}
	}()

	if err := watcher.Add(src); err != nil {
		fmt.Println("ERROR", err)
		return err
	}

	<-done

	return nil
}

// Same result for this use case and slightly better performance than the
// standard libraries strings.Replace(s, old, new, 1)
func replaceFirst(s, old, new string) string {
	ind := strings.Index(s, old)
	if ind == -1 {
		return s
	}
	ind += len(old)
	return new + s[ind:]
}

// copyFile is the same as deploy.CopyFile, but it takes a timeout and will
// try to copy the file until the timeout has expired. If it does not copy
// the file in that time frame it will return an error.
func copyFile(src, dest string, timeout int) error {
	ch := make(chan os.File, 1)

	go func() {
		defer close(ch)
		var f *os.File
		var e error
		t := 0
		for t < timeout {
			f, e = os.Open(src)
			if e != nil {
				time.Sleep(1 * time.Millisecond)
				t++
			} else {
				ch <- *f
				break
			}
		}
	}()

	for infile := range ch {
		defer infile.Close()
		outfile, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer outfile.Close()
		_, err = io.Copy(outfile, &infile)
		if err != nil {
			return err
		}
	}
	return nil
}
