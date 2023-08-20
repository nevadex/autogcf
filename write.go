package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Write() error {
	_, err := os.Stat(AutoGenDir)
	if err == nil {
		er := WaitForConfirm(AutoGenDir + " already exists. Delete?")
		if er != nil {
			return er
		}
		er = os.RemoveAll(AutoGenDir)
		if er != nil {
			return er
		}
		fmt.Println("Deleted old auto-generated directory")
	}

	err = os.Mkdir(AutoGenDir, DefaultPerm)
	if err != nil {
		return err
	}
	err = os.Chdir(AutoGenDir)
	if err != nil {
		return err
	}
	fmt.Println("Created auto-generated directory")

	for i := range RootFiles {
		er := os.Mkdir(RootFiles[i], DefaultPerm)
		if er != nil {
			return er
		}
		er = WriteGoModFile(filepath.Join(RootFiles[i], "go.mod"), "file/"+RootFiles[i])
		if er != nil {
			return er
		}

		ogFile, er := os.Open(filepath.Join(RootWorkingDir, SourceDir, RootFiles[i]))
		if er != nil {
			return er
		}
		newFile, er := os.Create(filepath.Join(RootFiles[i], RootFiles[i]))
		if er != nil {
			return er
		}
		_, er = io.Copy(newFile, ogFile)
		if er != nil {
			return er
		}
	}
	fmt.Println("Created single-file directories")

	for i := range RootDirs {
		er := os.Mkdir(RootDirs[i], DefaultPerm)
		if er != nil {
			return er
		}
		er = WriteGoModFile(RootDirs[i]+"/go.mod", "dir/"+RootDirs[i])
		if er != nil {
			return er
		}

		for x := range DirFiles[RootDirs[i]] {
			ogFile, e := os.Open(filepath.Join(RootWorkingDir, SourceDir, DirFiles[RootDirs[i]][x]))
			if e != nil {
				return e
			}
			newFile, e := os.Create(filepath.Join(RootDirs[i], strings.ReplaceAll(DirFiles[RootDirs[i]][x][len(RootDirs[i])+1:], string(os.PathSeparator), "_")))
			if e != nil {
				return e
			}
			_, e = io.Copy(newFile, ogFile)
			if e != nil {
				return e
			}
		}
	}
	fmt.Println("Created multi-file directories")
	fmt.Println("Copied all website source files")
	fmt.Println("Created all go.mod files")

	return WaitForConfirm("Start code generation?")
}

const DefaultPerm = 0666

func WriteGoModFile(fileName string, modName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	ffv, err := GetLatestFFVersionNumber()
	if err != nil {
		return err
	}

	_, err = f.WriteString(fmt.Sprintf(`module %v

require (
  github.com/GoogleCloudPlatform/functions-framework-go %v
)
`, ModulePath+modName, ffv))

	if err != nil {
		return err
	}
	return f.Close()
}

func GetLatestFFVersionNumber() (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", "https://github.com/GoogleCloudPlatform/functions-framework-go/releases/latest", nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	return strings.Split(resp.Header.Get("Location"), "/releases/tag/")[1], resp.Body.Close()
}
