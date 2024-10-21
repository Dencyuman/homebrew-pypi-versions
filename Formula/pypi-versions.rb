# Formula/ppv.rb
class Ppv < Formula
  desc "CLI tool to fetch package versions and metadata from PyPI."
  homepage "https://github.com/Dencyuman/ppv"
  url "https://github.com/Dencyuman/ppv/archive/refs/tags/v1.1.2.tar.gz"
  sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5"
  license "MIT" # プロジェクトのライセンスに応じて変更

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    # バージョン確認のテスト
    assert_match "ppv version 1.1.2", shell_output("#{bin}/ppv --version")
  end
end