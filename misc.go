package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetConfigValue(name string) (string, error) {
	cmd := exec.Command("gcloud", "config", "get-value", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}

func WaitForConfirm(msg string) error {
	if !SkipInteraction {
		fmt.Println("[!]", msg, "(Press ENTER to continue or CTRL+C to cancel)")
		r := bufio.NewReader(os.Stdin)
		_, err := r.ReadString('\n')
		return err
	} else {
		fmt.Println("[*]", msg, "Automatically continuing...")
		fmt.Println()
		return nil
	}
}
