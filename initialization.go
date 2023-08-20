package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	extensions     []string
	RootWorkingDir string
)

func Initialize() error {
	email, err := GetConfigValue("account")
	if err != nil {
		return err
	}
	fmt.Println("Using account:", email)

	projectId, err := GetConfigValue("project")
	if err != nil {
		return err
	}
	fmt.Println("Using project id:", projectId)

	fmt.Println("Reading from directory:", SourceDir)

	exts := ParseExtensions()
	fmt.Println("Allowing extensions:", exts)

	err = WaitForConfirm("Is this correct?")
	if err != nil {
		return err
	}

	RootWorkingDir, err = os.Getwd()
	if err != nil {
		return err
	}

	return os.Chdir(SourceDir)
}

func ParseExtensions() string {
	s := strings.Split(FileExtensions, ",")
	var ext string
	for i := range s {
		if s[i] != "" {
			extensions = append(extensions, "."+s[i])
			ext += "." + s[i] + " "
		}
	}
	return ext
}
