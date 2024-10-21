class PypiVersions < Formula
  desc "CLI tool to fetch package versions and metadata from PyPI."
  homepage "https://github.com/Dencyuman/ppv"
  url "https://github.com/Dencyuman/ppv/archive/refs/tags/v1.1.3.tar.gz"
  sha256 "d5558cd419c8d46bdc958064cb97f963d1ea793866414c025906ec15033512ed"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    # バージョン確認のテスト
    assert_match "ppv version 1.1.7
  end
end
