// cmd/versions.go
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions [packages...]",
	Short: "Display available versions of specified PyPI packages.",
	Long:  `Fetch and display all available versions of one or more specified PyPI packages.`,
	Example: `  # Display all available versions of pandas
  ppv versions pandas

  # Display only the latest stable version of pandas
  ppv versions pandas --latest

  # Include pre-release versions when displaying all versions
  ppv versions pandas --prerelease

  # Display the latest stable version in JSON format
  ppv versions pandas --latest --json

  # Display available versions for multiple packages
  ppv versions pandas requests numpy`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		includePreRelease, _ := cmd.Flags().GetBool("prerelease")
		showLatest, _ := cmd.Flags().GetBool("latest")
		outputJSON, _ := cmd.Flags().GetBool("json")

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

			var versions []*semver.Version
			for version := range pypiResp.Releases {
				v, err := semver.NewVersion(version)
				if err != nil {
					continue
				}
				if !includePreRelease && v.Prerelease() != "" {
					continue
				}
				versions = append(versions, v)
			}

			if len(versions) == 0 {
				fmt.Printf("No versions found for package %s.\n", packageName)
				continue
			}

			// Sort versions
			sort.Sort(semver.Collection(versions))

			if showLatest {
				latest := versions[len(versions)-1]
				if outputJSON {
					latestJSON, err := json.Marshal(struct {
						Package string `json:"package"`
						Latest  string `json:"latest"`
					}{
						Package: packageName,
						Latest:  latest.String(),
					})
					if err != nil {
						fmt.Printf("Error marshaling JSON for package %s: %v\n", packageName, err)
						os.Exit(1)
					}
					fmt.Println(string(latestJSON))
				} else {
					fmt.Printf("Latest version of %s: %s\n", packageName, latest.String())
				}
				continue
			}

			if outputJSON {
				versionStrs := []string{}
				for _, v := range versions {
					versionStrs = append(versionStrs, v.String())
				}
				versionsJSON, err := json.Marshal(struct {
					Package  string   `json:"package"`
					Versions []string `json:"versions"`
				}{
					Package:  packageName,
					Versions: versionStrs,
				})
				if err != nil {
					fmt.Printf("Error marshaling JSON for package %s: %v\n", packageName, err)
					os.Exit(1)
				}
				fmt.Println(string(versionsJSON))
			} else {
				fmt.Printf("Available versions for %s:\n", packageName)
				for _, v := range versions {
					fmt.Println(v.String())
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
	versionsCmd.Flags().BoolP("prerelease", "p", false, "Include pre-release versions")
	versionsCmd.Flags().BoolP("latest", "l", false, "Show only the latest stable version")
	versionsCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
}
