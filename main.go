package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type BuildTarget struct {
	GOOS   string
	GOARCH string
}

var DefaultTargets = []BuildTarget{
	{"linux", "amd64"},
	{"linux", "386"},
	{"linux", "arm"},
	{"linux", "arm64"},
	{"darwin", "amd64"},
	{"darwin", "arm64"},
	{"windows", "amd64"},
	{"windows", "386"},
}

func GetModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", fmt.Errorf("could not open go.mod: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading go.mod: %v", err)
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

func BuildSingle(goos, goarch, moduleName string) error {
	color.Green("Building for OS: %s, ARCH: %s\n", goos, goarch)

	// Set the output directory to a single "builds" directory
	outputDir := "builds"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_%s_%s", moduleName, goos, goarch))
	if goos == "windows" {
		outputFile += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", outputFile)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("build failed for %s/%s: %v\nOutput:\n%s", goos, goarch, err, output)
	}

	color.Green("Successfully built %s\n", outputFile)
	return nil
}

func BuildAll() {
	moduleName, err := GetModuleName()
	if err != nil {
		color.Red("Error getting module name: %v\n", err)
		return
	}

	for _, target := range DefaultTargets {
		if err := BuildSingle(target.GOOS, target.GOARCH, moduleName); err != nil {
			color.Red("Failed to build for %s/%s: %v\n", target.GOOS, target.GOARCH, err)
		}
	}
}

func main() {
	moduleName, err := GetModuleName()
	if err != nil {
		color.Red("Error: %v\n", err)
		return
	}

	if len(os.Args) == 3 {
		goos, goarch := os.Args[1], os.Args[2]
		if err := BuildSingle(goos, goarch, moduleName); err != nil {
			color.Red("Error: %v\n", err)
		}
	} else {
		BuildAll()
	}
}
