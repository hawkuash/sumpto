package misc

import (
	"fmt"
	"os"
	"path/filepath"
)

func glob(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		//fmt.Println(path)
		//fmt.Println(filepath.Ext(path))
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

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
		if fi.IsDir() {
			paths = append(paths, val)
		} else {
			files = append(files, val)
		}
	}
	return paths, files
}

func ListFiles(p string) ([]string, error) {

	fi, err := os.Stat(p)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {

		li, err := glob(p, ".png")
		//li, err := filepath.Glob(p + "/*.[jp]*g")
		if err != nil {
			return nil, err
		}
		//fmt.Println(len(li))
		for _, val := range li {
			fmt.Println(val)
		}
	} else {
		fmt.Println("Not a directory")
	}
}
