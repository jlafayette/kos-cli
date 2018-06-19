package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/jlafayette/kos-cli/deploy"
)

const help = `kosync continuously syncs a source folder to a destination folder.

Usage: kosync [options] SOURCE DEST

Options:
	-f --filter   Glob string for files to be synced.
	              If not specified, defaults to '*.ks' which matches KerboScript files.
`

func main() {
	filterPtr := flag.String("filter", "*.ks", "Glob string for files to be synced.")
	fPtr := flag.String("f", "", "(Alias for filter) Glob string for files to be synced.")

	flag.Parse()

	if len(flag.Args()) != 2 {
		// fmt.Fprintln(os.Stderr, "BAD!")
		fmt.Fprintln(os.Stderr, help)
		os.Exit(2)
	}

	filter := *filterPtr
	if *fPtr != "" && filter == "*.ks" {
		filter = *fPtr
	}
	src := flag.Args()[0]
	dst := flag.Args()[1]

	// Initial sync
	err := deploy.CopyFiles(src, dst, filter, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
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
						err := deploy.CopyFile(event.Name, replaceFirst(event.Name, src, dst), false)
						if err != nil {
							fmt.Printf("ERROR copying file: %v\n", err)
							// ToDo: occasionally gets error:
							// 	     open src/filname.ext: The process cannot access the file
							// 	     because it is being used by another process.
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
		os.Exit(2)
	}

	<-done
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
