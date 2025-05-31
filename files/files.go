package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func NewFilename(filename string, suffix string) string {
	ext := filepath.Ext(filename)
	return fmt.Sprintf("%s-%s%s", strings.TrimSuffix(filename, ext), suffix, ext)
}

func deduplicate(paths []string) []string {
	slices.Sort(paths)
	n := len(paths)

	var deduplicated []string
	for i, j := 0, 0; i < n && j < n; {
		if i == j {
			deduplicated = append(deduplicated, paths[i])
		}

		if strings.Contains(paths[j], paths[i]) {
			j += 1
		} else {
			i = j
		}
	}

	return deduplicated
}

func glob(dir string, ext []string, rec bool) ([]string, error) {
	var (
		files []string
		err   error
	)
	if rec {
		err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if slices.Contains(ext, filepath.Ext(path)) {
				files = append(files, path)
			}
			return nil
		})
	} else {
		paths, err := filepath.Glob(dir + "/*.*")
		if err != nil {
			return nil, err
		}
		for _, path := range paths {
			if slices.Contains(ext, filepath.Ext(path)) {
				files = append(files, path)
			}
		}
	}
	return files, err
}

func ParseInput(s string) ([]string, []string) {
	var files, paths []string
	l := filepath.SplitList(s)
	for _, val := range l {
		fi, err := os.Stat(val)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if fi.IsDir() && !slices.Contains(paths, val) {
			paths = append(paths, val)
		}
		if !fi.IsDir() && !slices.Contains(files, val) {
			files = append(files, val)
		}
	}
	return paths, files
}

func GenerateFiles(in string, rec bool, ext []string) []string {
	paths, files := ParseInput(in)
	if rec {
		paths = deduplicate(paths)
	}

	for _, path := range paths {
		found, err := glob(path, ext, rec)
		if err != nil {
			log.Printf("An error occured while searching for files at %s: %s \n", path, err)
			continue
		}
		files = append(files, found...)

	}

	slices.Sort(files)
	return slices.Compact(files)
}
