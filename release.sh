#!/bin/bash

# エラーが発生した場合にスクリプトを終了する
set -e

# 使用方法の表示
usage() {
  echo "使用方法: $0 <バージョン>"
  echo "例: $0 1.1.2"
  exit 1
}

# 引数のチェック
if [ "$#" -ne 1 ]; then
  usage
fi

VERSION=$1
TAG="v$VERSION"
REPO_OWNER="Dencyuman"
REPO_NAME="pypi-versions" # 実際のリポジトリ名に変更してください
FORMULA_FILE="Formula/pypi-versions.rb"
BREW_TAP_DIR=~/development/private/homebrew-pypi-versions # Homebrewタップのパスを適宜変更

# 必要なツールの確認
command -v git >/dev/null 2>&1 || { echo "gitがインストールされていません。インストールしてから再試行してください。"; exit 1; }
command -v curl >/dev/null 2>&1 || { echo "curlがインストールされていません。インストールしてから再試行してください。"; exit 1; }
command -v shasum >/dev/null 2>&1 || { echo "shasumがインストールされていません。インストールしてから再試行してください。"; exit 1; }

# 現在のブランチの確認
CURRENT_BRANCH=$(git branch --show-current)
echo "現在のブランチ: ${CURRENT_BRANCH}"

# タグの存在確認
if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "エラー: タグ '${TAG}' は既に存在します。別のバージョン番号を使用してください。"
  exit 1
fi

# 1. タグの作成
echo "Gitタグ ${TAG} を作成しています..."
git tag "$TAG"

# タグ作成の確認
if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "タグ '${TAG}' が正常に作成されました。"
else
  echo "エラー: タグ '${TAG}' の作成に失敗しました。"
  exit 1
fi

# 2. ソースコードのダウンロード
ARCHIVE_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/archive/refs/tags/${TAG}.tar.gz"
ARCHIVE_NAME="${REPO_NAME}-${VERSION}.tar.gz"

echo "ソースアーカイブをダウンロードしています: ${ARCHIVE_URL}"
curl -L "$ARCHIVE_URL" -o "$ARCHIVE_NAME"

# 3. SHA256チェックサムの計算
echo "SHA256チェックサムを計算しています..."
CHECKSUM=$(shasum -a 256 "$ARCHIVE_NAME" | awk '{print $1}')

echo "SHA256チェックサム: $CHECKSUM"

# 4. Homebrewフォーミュラの更新
echo "Homebrewフォーミュラを更新しています: ${FORMULA_FILE}"

# フォーミュラファイルが存在することを確認
if [ ! -f "${BREW_TAP_DIR}/${FORMULA_FILE}" ]; then
  echo "エラー: フォーミュラファイル ${FORMULA_FILE} が ${BREW_TAP_DIR} に存在しません。"
  exit 1
fi

# URLとSHA256の更新
sed -i '' "s|url \"https://github.com/${REPO_OWNER}/${REPO_NAME}/archive/refs/tags/.*\.tar\.gz\"|url \"https://github.com/${REPO_OWNER}/${REPO_NAME}/archive/refs/tags/${TAG}.tar.gz\"|g" "${BREW_TAP_DIR}/${FORMULA_FILE}"
sed -i '' "s|sha256 \".*\"|sha256 \"${CHECKSUM}\"|g" "${BREW_TAP_DIR}/${FORMULA_FILE}"

# バージョン確認のテスト部分の更新
sed -i '' "s|ppv version .*|ppv version ${VERSION}|g" "${BREW_TAP_DIR}/${FORMULA_FILE}"

# 5. フォーミュラの変更をコミット
echo "フォーミュラの変更をコミットしています..."
cd "$BREW_TAP_DIR"

git add "$FORMULA_FILE"
git commit -m "pypi-versions ${VERSION} リリース: URLとSHA256を更新"

echo "フォーミュラの変更がコミットされました。"

# プッシュは手動で行うため、スクリプトはここで終了
echo "リモートリポジトリにプッシュするには、以下のコマンドを実行してください："
echo "  git push origin ${CURRENT_BRANCH}"
echo "  git push origin ${TAG}"

# クリーンアップ
rm "$ARCHIVE_NAME"

echo "リリースプロセスのタグ作成とフォーミュラの更新が完了しました。"

exit 0
