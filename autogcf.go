package main

import (
	"fmt"
	"os"
)

func AutoGCF() error {
	/*f := templates.GenMultiFile("dir", "\"site.js\"", "\"site.js\"", "\"text/javascript\"")
	fmt.Println(f.Render(os.Stdout))
	os.Exit(0)*/

	fmt.Println("[ Initializing ]")
	if err := Initialize(); err != nil {
		return err
	}

	fmt.Println("[ Reading ]")
	if err := Read(); err != nil {
		return err
	}

	fmt.Println("[ Writing ]")
	if err := Write(); err != nil {
		return err
	}

	fmt.Println("[ Generating ]")
	if err := Generate(); err != nil {
		return err
	}

	fmt.Println("[ Deploying ]")
	if err := Deploy(); err != nil {
		return err
	}

	if DeleteGenFiles {
		err := os.RemoveAll(AutoGenDir)
		if err != nil {
			return err
		}
		fmt.Println("Deleted auto-generated directory")
	}

	return nil
}
