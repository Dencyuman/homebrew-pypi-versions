// cmd/metadata.go
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// metadataCmd represents the metadata command
var metadataCmd = &cobra.Command{
	Use:   "metadata [packages...]",
	Short: "Display detailed metadata of specified Python packages from PyPI.",
	Long: `metadata is a sub-command of ppv that allows you to fetch
detailed metadata information of specified Python packages from PyPI.`,
	Example: `  # Display basic metadata of pandas (without description)
  ppv metadata pandas

  # Display metadata of pandas with description
  ppv metadata pandas --description

  # Display metadata of multiple packages with description
  ppv metadata pandas requests numpy --description

  # Display metadata in JSON format
  ppv metadata pandas --json

  # Display metadata with description in JSON format
  ppv metadata pandas --json --description`,
	Args: cobra.MinimumNArgs(1), // Require at least one argument
	Run: func(cmd *cobra.Command, args []string) {
		runMetadata(cmd, args)
	},
}

// Flags
var (
	outputJSONMetadata bool
	showDescription    bool
)

func init() {
	// Register the metadata command as a subcommand of rootCmd
	rootCmd.AddCommand(metadataCmd)

	// Define flags specific to metadata command
	metadataCmd.Flags().BoolVarP(&outputJSONMetadata, "json", "j", false, "Output in JSON format")
	metadataCmd.Flags().BoolVarP(&showDescription, "description", "d", false, "Include package description in the output")
}

// runMetadata handles fetching and displaying the metadata information
func runMetadata(cmd *cobra.Command, args []string) {
	for _, packageName := range args {
		url := fmt.Sprintf("https://pypi.org/pypi/%s/json", packageName)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching API for package %s: %v\n", packageName, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("Package '%s' not found. HTTP Status: %d\n", packageName, resp.StatusCode)
			continue
		}

		var pypiResp PyPIResponse
		if err := json.NewDecoder(resp.Body).Decode(&pypiResp); err != nil {
			fmt.Printf("Error decoding JSON for package %s: %v\n", packageName, err)
			continue
		}

		if outputJSONMetadata {
			// Prepare a map to hold the metadata
			metadataMap := map[string]interface{}{
				"Name":         pypiResp.Info.Name,
				"Version":      pypiResp.Info.Version,
				"Summary":      pypiResp.Info.Summary,
				"Author":       pypiResp.Info.Author,
				"Author Email": pypiResp.Info.AuthorEmail,
				"License":      pypiResp.Info.License,
				"Home Page":    pypiResp.Info.HomePage,
			}

			// Conditionally add Description if flag is set
			if showDescription {
				metadataMap["Description"] = pypiResp.Info.Description
			}

			// Conditionally add Repository URL if available
			if pypiResp.Info.ProjectURL != "" {
				metadataMap["Repository URL"] = pypiResp.Info.ProjectURL
			}

			// Marshal the map to JSON
			metadataJSON, err := json.MarshalIndent(metadataMap, "", "  ")
			if err != nil {
				fmt.Printf("Error marshaling JSON for package %s: %v\n", packageName, err)
				os.Exit(1)
			}
			fmt.Println(string(metadataJSON))
		} else {
			// Human-readable output
			displayMetadata(pypiResp.Info, packageName, showDescription)

			// If Description is not shown, inform the user about the flag
			if !showDescription {
				fmt.Printf("To include the description in the output, use the '--description' flag.\n\n")
			}
		}
	}
}

// displayMetadata prints the metadata in a human-readable format
func displayMetadata(info PackageInfo, packageName string, showDescription bool) {
	fmt.Printf("Metadata for %s:\n", packageName)
	fmt.Printf("  Name: %s\n", info.Name)
	fmt.Printf("  Version: %s\n", info.Version)
	fmt.Printf("  Summary: %s\n", info.Summary)
	fmt.Printf("  Author: %s\n", info.Author)
	fmt.Printf("  Author Email: %s\n", info.AuthorEmail)
	fmt.Printf("  License: %s\n", info.License)
	fmt.Printf("  Home Page: %s\n", info.HomePage)
	fmt.Printf("  Repository URL: %s\n", info.ProjectURL)

	if showDescription {
		fmt.Printf("  Description: %s\n", info.Description)
	}

	fmt.Println()
}
