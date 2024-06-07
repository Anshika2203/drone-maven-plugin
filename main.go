package main

import (
	"fmt"
	"os/exec"
)

func main() {
	buildCmd := exec.Command("mvn", "clean install")
	buildOutput, buildErr := buildCmd.CombinedOutput()
	if buildErr != nil {
		fmt.Println("Error running 'mvn clean install':", buildErr)
		fmt.Println(string(buildOutput))
		return
	}
	fmt.Println("Output of 'mvn clean install':", string(buildOutput))

	versionCmd := exec.Command("mvn", "--version")
	versionOutput, versionErr := versionCmd.CombinedOutput()
	if versionErr != nil {
		fmt.Println("Error running 'maven -version':", versionErr)
		fmt.Println(string(versionOutput))
		return
	}
	fmt.Println("Output of 'maven -version':", string(versionOutput))
}
