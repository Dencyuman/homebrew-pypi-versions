# Formula/ppv.rb
class Ppv < Formula
  desc "CLI tool to fetch package versions and metadata from PyPI."
  homepage "https://github.com/Dencyuman/ppv"
  url "https://github.com/Dencyuman/ppv/archive/refs/tags/v1.1.0.tar.gz"
  sha256 "5fcd5055074d22c3d23a756bb3aaab1abc6851585393feaf1489cb84ee9c2f4d"
  license "MIT" # プロジェクトのライセンスに応じて変更

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    # バージョン確認のテスト
    assert_match "ppv version 1.1.0", shell_output("#{bin}/ppv --version")
  end
end