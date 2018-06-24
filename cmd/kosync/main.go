package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jlafayette/kos-cli/deploy"
)

const help = `kosync continuously syncs a dev source folder to KSP Ships\Script folder.

Usage: kosync [options]

Options:
	-s --source   Development root folder to sync from. If not specified,
	              defaults to $KSPSRC environment variable.
	-d --dest     KSP Ships/Script folder to sync to. If not specified,
	              defaults to $KSPSCRIPT environment variable.
`

func main() {
	sourcePtr := flag.String("source", "", "Development root folder to sync from.")
	sPtr := flag.String("s", os.Getenv("KSPSRC"), "(Alias for source) Development root folder to sync from.")
	destPtr := flag.String("dest", "", "KSP Ships/Script folder to sync to.")
	dPtr := flag.String("d", os.Getenv("KSPSCRIPT"), "(Alias for dest) KSP Ships/Script folder to sync to.")
	flag.Parse()

	if len(flag.Args()) != 0 {
		fmt.Printf("ERROR: Too many arguments: %v\n\n", flag.Args())
		fmt.Fprintln(os.Stderr, help)
		os.Exit(2)
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
	fmt.Printf("match: %v\n", deploy.KsMatch)
	fmt.Printf("  src: %v\n", src)
	fmt.Printf("  dst: %v\n", dst)

	done := make(chan bool)
	deployInfo := deploy.GetInstructions(src, dst, deploy.KsMatch)
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

func sync(src, dst, match string) error {
	fmt.Printf("Starting sync %v -> %v\n", src, dst)

	// Initial sync
	err := deploy.CopyFiles(src, dst, match, false)
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
				m, _ := filepath.Match(match, file)
				if m {
					switch event.Op {
					case 1: // Create Op  = 1 << iota
						newFile, err := os.Create(replaceFirst(event.Name, src, dst))
						if err != nil {
							fmt.Printf("ERROR creating file: %v\n", err)
						}
						newFile.Close()
					case 2: // Write --> copy
						err := retryCopyFile(event.Name, replaceFirst(event.Name, src, dst), 10)
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

// retry will retry a function given number of times with a given
// delay, but doubling the delay after each failed try.
func retry(tries int, delay time.Duration, fn func() error) error {
	err := fn()
	if err != nil {
		// If there are still more tries then try again
		if tries--; tries > 0 {
			time.Sleep(delay)
			return retry(tries, delay*2, fn)
		}
		// Out of tries
		return err
	}
	return nil
}

// retryCopyFile retries the CopyFile function for given number of times.
func retryCopyFile(src, dst string, tries int) error {
	return retry(tries, time.Millisecond, func() error {
		return deploy.CopyFile(src, dst, false)
	})
}
