/*
Application that renames images

*** MORE DOCUMENTATION NEEDE ***
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

var debug bool
var dir string
var dry bool

func init() {
	const (
		dryrun    = "Show what renimg would do, but don't do it."
		dryrundef = false
	)
	// Flags
	flag.StringVar(&dir, "dir", "", "Image directory")

	flag.BoolVar(&dry, "dry-run", dryrundef, dryrun)
	flag.BoolVar(&dry, "d", dryrundef, dryrun+" (shorthand)")

	flag.BoolVar(&debug, "debug", false, "Output debugging messages")
}

func main() {
	flag.Parse()

	// If dir is empty, use current
	if dir == "" {
		dir, _ = os.Getwd()
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("Bad path (%s) specified: %v", dir, err)
	}
	debugmsg(fmt.Sprintf("renimg running against %s\r", dir))

	// Compile first as complex and possibly for performance
	pattern := "[[:digit:]]{4}-[[:digit:]]{2}-[[:digit:]]{2}-[[:digit:]]{6}-.*\\.[[:alpha:]]{3,4}"
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Error in regular expression check: %v", err)
	}

	// Execute this function on every file
	callback := func(path string, fi os.FileInfo, err error) error {

		// Only deal with JPEGs
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".jpg" && ext != ".jpeg" {
			return nil
		}

		// Only deal with non-changed files
		if match := r.MatchString(fi.Name()); match {
			return nil
		}
		debugmsg("No match")

		// Get new name
		newname, err := getNewName(fi.Name())
		if err != nil {
			return filepath.SkipDir
		}

		debugmsg(newname)
		debugmsg(fi.Name())

		// Name change
		if dry {
			log.Printf("Rename %s to %s\r", fi.Name(), newname)
		} else {
			os.Rename(path, newname)
			debugmsg(fmt.Sprintf("DEBUG: Rename %s to %s\r", fi.Name(), newname))
		}

		return nil
	}

	// Walk tree under dir
	err = filepath.Walk(dir, callback)
	if err != nil {
		log.Fatalf("Error processing files: %v", err)
	}
}

func getNewName(file string) (string, error) {
	// open file and Decode
	f, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening file %s: %v\r", file, err)
		return "", err
	}
	defer f.Close()

	ex, err := exif.Decode(f)
	if err != nil {
		log.Printf("Error decoding file %s: %v\r", file, err)
		return "", err
	}

	// get exif fields
	time, err := ex.DateTime()
	if err != nil {
		log.Printf("Error reading file %s' creation time: %v\r", file, err)
		return "", err
	}

	// build new Name
	return time.Format("2006-01-02-150405-") + file, nil
}

// Checks the debug flag and outputs if true
func debugmsg(message string) {
	if debug {
		log.Println(message)
	}
}
