package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	RootFiles []string
	RootDirs  []string
	DirFiles  = make(map[string][]string)
)

func Read() error {
	var err error
	RootFiles, RootDirs, err = ListRootFiles()
	if err != nil {
		return err
	}

	fmt.Println("Listing root files:")
	for i := range RootFiles {
		fmt.Println(RootFiles[i])
	}
	fmt.Println("Listing root directories:")
	for i := range RootDirs {
		fmt.Println(RootDirs[i])
	}

	err = WaitForConfirm("Are these the correct files and directories?")
	if err != nil {
		return err
	}

	for i := range RootDirs {
		f, er := GetFilesInPath(RootDirs[i])
		if er != nil {
			return er
		}
		DirFiles[RootDirs[i]] = f
	}

	return os.Chdir(RootWorkingDir)
}

func GetFilesInPath(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}

	var good []string

	for i := range files {
		for x := range extensions {
			if files[i][len(files[i])-len(extensions[x]):] == extensions[x] {
				good = append(good, files[i])
				break
			}
		}
	}

	return good, nil
}

func ListRootFiles() ([]string, []string, error) {
	en, err := os.ReadDir(".")
	if err != nil {
		return []string{}, []string{}, err
	}

	var files []string
	var dir []string

	for i := range en {
		if en[i].IsDir() {
			dir = append(dir, en[i].Name())
		} else {
			for x := range extensions {
				if en[i].Name()[len(en[i].Name())-len(extensions[x]):] == extensions[x] {
					files = append(files, en[i].Name())
					break
				}
			}
		}
	}

	return files, dir, nil
}
