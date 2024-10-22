// cmd/root.go
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

// バージョン情報を定義
var (
	version = "v1.1.14"
)

// rootCmdはベースとなるコマンドです
var rootCmd = &cobra.Command{
	Use:   "ppv [packages...]",
	Short: "PyPI Versions is a CLI tool to fetch package versions and metadata from PyPI.",
	Long: `ppv is a command-line interface tool that allows you to fetch
available versions and detailed metadata of specified Python packages from PyPI.`,
	Example: `  # Display all available versions of pandas
    ppv pandas

    # Display only the latest stable version of pandas
    ppv pandas --latest

    # Include pre-release versions when displaying all versions
    ppv pandas --prerelease

    # Display the latest stable version in JSON format
    ppv pandas --latest --json

    # Display available versions for multiple packages
    ppv pandas requests numpy

    # Display metadata of pandas
    ppv metadata pandas`,
	Args: cobra.ArbitraryArgs, // 任意の引数を許可
	Run: func(cmd *cobra.Command, args []string) {
		// バージョンフラグが設定されている場合はバージョンを表示して終了
		if showVersion {
			fmt.Printf("PyPi-Versions(ppv): %s\n", version)
			return
		}

		if len(args) > 0 {
			// 引数が提供され、サブコマンドが使用されていない場合はversionsコマンドとして扱う
			runVersions(cmd, args)
		} else {
			// それ以外の場合はヘルプを表示
			cmd.Help()
		}
	},
}

// Flags
var (
	includePreRelease bool
	showLatest        bool
	outputJSON        bool
	showVersion       bool // バージョン表示用フラグを追加
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// フラグを定義
	rootCmd.Flags().BoolVarP(&includePreRelease, "prerelease", "p", false, "Include pre-release versions")
	rootCmd.Flags().BoolVarP(&showLatest, "latest", "l", false, "Show only the latest stable version")
	rootCmd.Flags().BoolVarP(&outputJSON, "json", "j", false, "Output in JSON format")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show ppv version")
}

// runVersionsはバージョン取得ロジックを処理します
func runVersions(cmd *cobra.Command, args []string) {
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

		// バージョンをソート
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
}
