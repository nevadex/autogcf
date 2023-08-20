package main

import (
	"fmt"
	pb "github.com/schollz/progressbar/v3"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func Deploy() error {
	allDirs := append(RootFiles, RootDirs...)
	var wg sync.WaitGroup
	bar := pb.NewOptions(len(allDirs), pb.OptionSetItsString("deployed"), pb.OptionSetDescription("Deploying all functions"), pb.OptionSetRenderBlankState(true))
	for i := range allDirs {
		wg.Add(1)
		i := i
		go func() {
			defer bar.Add(1)
			defer wg.Done()
			deployAsync(allDirs[i])
		}()
	}

	wg.Wait()

	for true {
		if bar.IsFinished() {
			fmt.Println()
			break
		} else {
			time.Sleep(time.Second)
		}
	}

	cmd := exec.Command("gcloud", "functions", "describe", allDirs[0], "--region="+Region, "--format=value(url)")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %v", err, string(output))
	}
	url := strings.Split(string(output), "/"+allDirs[0])
	fmt.Println()
	fmt.Println("Your website is available here:", url[0])
	fmt.Println("Create a CNAME record on your custom domain's DNS that maps to the above URL.")

	return nil
}

func deployAsync(source string) {
	args := []string{
		"functions",
		"deploy",
		source,
		"--gen2",
		"--region=" + Region,
		"--runtime=go120",
		"--source=" + source,
		"--entry-point=" + source,
		"--trigger-http",
		"--memory=" + Memory,
		"--max-instances=" + MaxInstances,
		"--allow-unauthenticated",
	}
	cmd := exec.Command("gcloud", args...)
	_, _ = cmd.CombinedOutput()
}

// gcloud functions deploy js --gen2 --region us-east4 --runtime go120 --source . --entry-point js --trigger-http --memory 128Mi --max-instances 10 --allow-unauthenticated
