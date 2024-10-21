# PyPI Versions (ppv)

PyPI Versions (ppv) is a command-line interface (CLI) tool designed to interact with the Python Package Index (PyPI). It allows users to fetch package versions, detailed metadata, and dependencies of specified Python packages directly from the terminal.

## Docs
- [日本語版ドキュメント](./docs/ja/README.md)

## Features
- **Fetch Available Versions**: Retrieve all available versions of one or more Python packages from PyPI.
- **Latest Version**: Display only the latest stable version of a package.
- **Pre-release Versions**: Optionally include pre-release versions in the output.
- **JSON Output**: Get outputs in JSON format for easy integration with other tools and scripts.
- **Package Metadata**: Access detailed metadata of Python packages, including summaries, authors, licenses, and more.
- **Dependencies**: View dependencies of specific versions of Python packages.

## Installation
### Prerequisites
- **Go**: Ensure you have Go installed (version 1.18 or higher).

### Install via go install
```bash
go install github.com/Dencyuman/pypi-versions@latest
```

This command will install the ppv binary to your $GOPATH/bin. Make sure this directory is added to your system's PATH to access ppv from anywhere in your terminal.

### Using Homebrew (macOS and Linux)
If you are using Homebrew, you can install ppv with:

```bash
brew tap Dencyuman/pypi-versions
brew install pypi-versions
```
Note: Ensure that the Homebrew tap is available. If not, you might need to create a Homebrew formula for ppv.

## Usage
The general syntax for using ppv is:

```bash
ppv [command] [flags] [arguments]
```
### Global Flags
- `--prerelease`, `-p`: Include pre-release versions in the output.
- `--latest`, `-l`: Show only the latest stable version.
- `--json`, `-j`: Output the results in JSON format.
- `--help`, `-h`: Display help for specific commands.

### Commands
`versions`
Display available versions of specified PyPI packages.

#### Usage:

```bash
ppv versions [packages...] [flags]
```

`metadata`
Display detailed metadata of specified Python packages from PyPI.

#### Usage:

```bash
ppv metadata [packages...] [flags]
```

`deps`
Display dependencies of specific versions of Python packages from PyPI.

#### Usage:

```bash
ppv deps [package] [version] ... [flags]
```

`help`
Display overall help information.

```bash
ppv help
```

## Examples
### Display All Available Versions of a Package
```bash
ppv versions pandas
```
#### Output:

```bash
Available versions for pandas:
1.0.0
1.1.0
1.2.0
...
```

### Display Only the Latest Stable Version
```bash
ppv versions pandas --latest
```
#### Output:

```bash
Latest version of pandas: 1.5.3
```

### Include Pre-release Versions
```bash
ppv versions pandas --prerelease
```
#### Output:

```bash
Available versions for pandas:
1.0.0
1.1.0
1.2.0
1.3.0-beta
1.4.0
1.5.0-rc1
1.5.3
```

### Output Versions in JSON Format
```bash
ppv versions pandas --json
```
#### Output:

```json
{
    "package": "pandas",
    "versions": [
        "1.0.0",
        "1.1.0",
        "1.2.0",
        "1.3.0",
        "1.4.0",
        "1.5.3"
    ]
}
```

### Display Metadata of a Package
```bash
ppv metadata pandas
```
#### Output:

```bash
Metadata for pandas:
Name: pandas
Version: 1.5.3
Summary: Powerful data structures for data analysis, time series, and statistics
Author: The pandas development team
Author Email: pandas-dev@python.org
License: BSD License
Home Page: https://pandas.pydata.org/
Repository URL: https://github.com/pandas-dev/pandas

To include the description in the output, use the '--description' flag.
```

### Display Dependencies of a Specific Version
```bash
ppv deps pandas 1.5.3
```
#### Output:

```bash
Dependencies for pandas version 1.5.3:
- numpy>=1.20.0
- python-dateutil>=2.7.3
- pytz>=2017.3
```

### Display Dependencies of the Latest Version in JSON Format
```bash
ppv deps pandas latest --json
```
#### Output:

```bash
{
    "package": "pandas",
    "version": "1.5.3",
    "dependencies": [
        "numpy>=1.20.0",
        "python-dateutil>=2.7.3",
        "pytz>=2017.3"
    ]
}
```

## License
This project is licensed under the [MIT License](../../LICENSE).
© 2024 [Dencyuman](https://github.com/Dencyuman)