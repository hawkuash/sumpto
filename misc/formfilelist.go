package misc

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func deduplicate(paths []string) []string {
	slices.Sort(paths)

	var new []string
	i, j := 0, 0
	for i < len(paths) {
		if i == j {
			new = append(new, paths[i])
		}
		if strings.Contains(paths[j], paths[i]) {
			j = j + 1
		} else {
			i = j
		}
		if j >= len(paths) {
			break
		}
	}
	return new
}

func globRec(dir string, ext []string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if slices.Contains(ext, filepath.Ext(path)) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func glob(dir string, ext []string) ([]string, error) {

	files := []string{}
	li, err := filepath.Glob(dir + "/*.*")
	if err != nil {
		return nil, err
	}
	for _, path := range li {
		if slices.Contains(ext, filepath.Ext(path)) {
			files = append(files, path)
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

func FormFileList(in string, rec bool, ext []string) []string {
	var files, paths []string
	paths, files = ParseInput(in)
	if rec {
		paths = deduplicate(paths)
		for _, path := range paths {
			found, err := globRec(path, ext)
			if err == nil {
				files = append(files, found...)
			}
		}
	}
	if !rec {
		for _, path := range paths {
			found, err := glob(path, ext)
			if err == nil {
				files = append(files, found...)
			}
		}
	}
	slices.Sort(files)
	files = slices.Compact(files)
	return files
}
