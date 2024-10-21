// cmd/deps.go
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// depsCmd represents the deps command
var depsCmd = &cobra.Command{
	Use:   "deps [package] [version] ...",
	Short: "Display dependencies of specific versions of Python packages from PyPI.",
	Long: `deps is a sub-command of ppv that allows you to fetch
dependencies information of specified versions of Python packages from PyPI.

You can specify 'latest' as the version to fetch dependencies for the latest version.

If the version is omitted, 'latest' is assumed.`,
	Example: `  # Display dependencies of pandas version 1.5.3
    ppv deps pandas 1.5.3

  # Display dependencies of pandas latest version
    ppv deps pandas

  # Display dependencies of multiple packages with mixed version specifications
    ppv deps pandas 1.5.3 requests numpy latest

  # Display dependencies in JSON format
    ppv deps pandas latest --json

  # Mixed usage with latest and specific versions
    ppv deps pandas latest requests 2.31.0`,
	Args: cobra.MinimumNArgs(1), // Require at least one argument (package name)
	Run:  runDeps,
}

// Flags
var (
	outputJSONDeps bool
)

func init() {
	// Register the deps command as a subcommand of rootCmd
	rootCmd.AddCommand(depsCmd)

	// Define flags specific to deps command
	depsCmd.Flags().BoolVarP(&outputJSONDeps, "json", "j", false, "Output dependencies in JSON format")
}

// runDeps handles fetching and displaying the dependencies information
func runDeps(cmd *cobra.Command, args []string) {
	// Check if the number of arguments is valid (pairs of [package] [version])
	if len(args) < 1 {
		fmt.Println("Error: Please provide at least one package name.")
		fmt.Println("Usage: ppv deps [package] [version] ...")
		os.Exit(1)
	}

	// Initialize an index to iterate through the arguments
	i := 0
	for i < len(args) {
		packageName := args[i]
		i++

		var version string

		// Check if the next argument exists and is not a package name
		// For simplicity, we'll assume that if the next argument does not contain
		// a dot and is not 'latest', it's likely a package name. Otherwise, treat it as a version.
		if i < len(args) {
			nextArg := args[i]
			if !isPackageName(nextArg) {
				version = nextArg
				i++
			} else {
				version = "latest"
			}
		} else {
			// If there's no next argument, default to 'latest'
			version = "latest"
		}

		// If the version is explicitly set to 'latest', fetch the latest version
		if strings.ToLower(version) == "latest" {
			latestVer, err := getLatestVersion(packageName)
			if err != nil {
				fmt.Printf("Error fetching latest version for package '%s': %v\n", packageName, err)
				continue
			}
			version = latestVer
			fmt.Printf("Fetching dependencies for package '%s' version '%s' (latest).\n", packageName, version)
		}

		// Fetch package info for the specified version
		pypiResp, err := fetchPackageInfo(packageName, version)
		if err != nil {
			fmt.Println(err)
			continue
		}

		dependencies := pypiResp.Info.RequiresDist

		if dependencies == nil || len(dependencies) == 0 {
			fmt.Printf("No dependencies found for package '%s' version '%s'.\n", packageName, version)
			continue
		}

		if outputJSONDeps {
			// Prepare a map to hold the dependencies
			depList := []string{}
			for _, dep := range dependencies {
				depList = append(depList, dep)
			}
			output := map[string]interface{}{
				"package":      packageName,
				"version":      version,
				"dependencies": depList,
			}
			depJSON, err := json.MarshalIndent(output, "", "  ")
			if err != nil {
				fmt.Printf("Error marshaling JSON for package '%s' version '%s': %v\n", packageName, version, err)
				os.Exit(1)
			}
			fmt.Println(string(depJSON))
		} else {
			// Human-readable output
			displayDeps(packageName, version, dependencies)
		}
	}
}

// isPackageName is a helper function to determine if the argument is a package name
// For simplicity, this function assumes that a package name does not contain dots and
// does not resemble a version string. This can be enhanced based on specific requirements.
func isPackageName(arg string) bool {
	// Common heuristic: package names usually do not contain '=' or other version specifiers.
	// Also, they might contain '-' or '_', but not necessarily dots.
	// Here, if the arg contains digits and dots in a way that's typical for versions, treat differently.
	// This is a simplistic check and may need improvements.
	if strings.Contains(arg, "=") || strings.Contains(arg, ">") || strings.Contains(arg, "<") {
		return false
	}
	// Another heuristic: if the arg can be converted to a semantic version, it's a version
	_, err := semver.NewVersion(arg)
	return err != nil
}

// getLatestVersion fetches the latest version of a package from PyPI
func getLatestVersion(packageName string) (string, error) {
	url := fmt.Sprintf("https://pypi.org/pypi/%s/json", packageName)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("package '%s' not found. HTTP Status: %d", packageName, resp.StatusCode)
	}

	var pypiResp PyPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&pypiResp); err != nil {
		return "", fmt.Errorf("error decoding JSON: %w", err)
	}

	if pypiResp.Info.Version == "" {
		return "", fmt.Errorf("latest version not found for package '%s'", packageName)
	}

	return pypiResp.Info.Version, nil
}

// fetchPackageInfo fetches the package info from PyPI for a given package and version
func fetchPackageInfo(packageName, version string) (*PyPIResponse, error) {
	url := fmt.Sprintf("https://pypi.org/pypi/%s/%s/json", packageName, version)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching API for package '%s' version '%s': %w", packageName, version, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Package '%s' version '%s' not found. HTTP Status: %d", packageName, version, resp.StatusCode)
	}

	var pypiResp PyPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&pypiResp); err != nil {
		return nil, fmt.Errorf("error decoding JSON for package '%s' version '%s': %w", packageName, version, err)
	}

	return &pypiResp, nil
}

// displayDeps prints the dependencies in a human-readable format
func displayDeps(packageName, version string, dependencies []string) {
	fmt.Printf("Dependencies for %s version %s:\n", packageName, version)
	for _, dep := range dependencies {
		fmt.Printf("  - %s\n", dep)
	}
	fmt.Println()
}
