package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Function to initialize Maven settings
	initMavenSettings()

	// Construct the Maven command
	mvnCommand := "mvn -B"

	// Set default goals if not provided
	if os.Getenv("PLUGIN_GOALS") == "" {
		fmt.Println("No $PLUGIN_GOALS defined, using default goals -DskipTests clean install")
		os.Setenv("PLUGIN_GOALS", "-DskipTests clean install")
	} else {
		os.Setenv("PLUGIN_GOALS", strings.ReplaceAll(os.Getenv("PLUGIN_GOALS"), ",", " "))
	}
	mvnCommand += " " + os.Getenv("PLUGIN_GOALS")

	// Add Maven modules if provided
	if os.Getenv("PLUGIN_MAVEN_MODULES") != "" {
		mvnCommand += " -pl " + os.Getenv("PLUGIN_MAVEN_MODULES")
	}

	// Add context directory if provided
	if os.Getenv("PLUGIN_CONTEXT_DIR") != "" {
		mvnCommand += " -f " + os.Getenv("PLUGIN_CONTEXT_DIR")
	}

	// Print the Maven command if debug mode is enabled
	if os.Getenv("PLUGIN_LOG_LEVEL") == "debug" {
		fmt.Println("Running command:", mvnCommand)
	}

	// Execute the Maven command
	cmd := exec.Command("bash", "-c", mvnCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// Function to initialize Maven settings
func initMavenSettings() {
	// Ensure MAVEN_CONFIG is set to a default value if not provided
	if os.Getenv("MAVEN_CONFIG") == "" {
		os.Setenv("MAVEN_CONFIG", "/root/.m2")
	}
}
