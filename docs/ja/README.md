# PyPI Versions (ppv)

PyPI Versions (ppv) は、Python Package Index (PyPI) と連携する強力なコマンドラインインターフェース（CLI）ツールです。指定したPythonパッケージのバージョン、詳細なメタデータ、依存関係をターミナルから直接取得できます。

## Docs
- [English Docs](../../README.md)

## 特徴
- **利用可能なバージョンの取得**: 一つ以上のPythonパッケージのすべてのバージョンをPyPIから取得。
- **最新バージョンの表示**: パッケージの最新の安定バージョンのみを表示。
- **プレリリースバージョンの含有**: オプションでプレリリースバージョンを出力に含める。
- **JSON出力**: JSON形式で結果を出力し、他のツールやスクリプトとの連携を容易に。
- **パッケージメタデータの取得**: サマリー、著者、ライセンスなど、詳細なメタデータを取得。
- **依存関係の表示**: 特定バージョンのパッケージの依存関係を表示。

## インストール
### 必要条件
- **Go**: Go 1.18以上がインストールされていること。 ダウンロードはこちら。

### go install を使用してインストール
```bash
go install github.com/Dencyuman/pypi-versions@latest
```

このコマンドにより、ppv バイナリが $GOPATH/bin にインストールされます。このディレクトリがシステムの PATH に含まれていることを確認してください。

### 事前コンパイル済みバイナリのダウンロード
リポジトリの Releases ページから事前コンパイル済みバイナリをダウンロードできます。ダウンロード後、バイナリをシステムの PATH に含まれるディレクトリに配置してください。

### Homebrew を使用してインストール（macOSおよびLinux）
Homebrewを使用している場合、以下のコマンドで ppv をインストールできます。

```bash
brew tap Dencyuman/pypi-versions
brew install pypi-versions
```
注意: Homebrew tapが利用可能であることを確認してください。利用できない場合は、Homebrewフォーミュラの作成が必要です。

## 使い方
一般的な ppv の使用方法は以下の通りです。

```bash
ppv [コマンド] [フラグ] [引数]
```

### グローバルフラグ
- `--prerelease`, `-p`: プレリリースバージョンを含める。
- `--latest`, `-l`: 最新の安定バージョンのみを表示。
- `--json`, `-j`: 結果をJSON形式で出力。

### コマンド
`versions`
指定したPyPIパッケージの利用可能なバージョンを表示します。

#### 使用方法:

```bash
ppv versions [パッケージ...] [フラグ]
```

`metadata`
指定したPythonパッケージの詳細なメタデータを表示します。

#### 使用方法:

```bash
ppv metadata [パッケージ...] [フラグ]
```

`deps`
指定したバージョンのPythonパッケージの依存関係を表示します。

#### 使用方法:

```bash
ppv deps [パッケージ] [バージョン] ... [フラグ]
```

## 例
### パッケージのすべての利用可能なバージョンを表示
```bash
ppv versions pandas
```
#### 出力:

```bash
Available versions for pandas:
1.0.0
1.1.0
1.2.0
...
```

### 最新の安定バージョンのみを表示
```bash
ppv versions pandas --latest
```
#### 出力:


```bash
Latest version of pandas: 1.5.3
```

### プレリリースバージョンを含める
```bash
ppv versions pandas --prerelease
```

#### 出力:

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

### JSON形式でバージョンを出力
```bash
ppv versions pandas --json
```
#### 出力:

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

###  パッケージのメタデータを表示
```bash
ppv metadata pandas
```
#### 出力:

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

### 特定バージョンの依存関係を表示
```bash
ppv deps pandas latest --json
```
#### 出力:

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

## ライセンス
このプロジェクトは MIT License のもとでライセンスされています。詳細については [LICENSE](../../LICENSE) ファイルをご覧ください。