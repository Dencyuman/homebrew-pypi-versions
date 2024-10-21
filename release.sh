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
REPO_NAME="homebrew-pypi-versions" # 実際のリポジトリ名に変更してください
REPO_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}.git"

# 必要なツールの確認
command -v git >/dev/null 2>&1 || { echo "gitがインストールされていません。インストールしてから再試行してください。"; exit 1; }

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

# プッシュは手動で行うため、スクリプトはここで終了
echo "タグ '${TAG}' がローカルリポジトリに作成されました。"
echo "リモートリポジトリにプッシュするには、以下のコマンドを実行してください："
echo "  git push origin ${TAG}"

exit 0
