class PypiVersions < Formula
  desc "CLI tool to fetch package versions and metadata from PyPI."
  homepage "https://github.com/Dencyuman/ppv"
  url "https://github.com/Dencyuman/pypi-versions/archive/refs/tags/v1.1.11.tar.gz"
  sha256 "cd8cd0ac042abbe5f6a8bce7581df3cc85ad666858fd82748f7dd5c12376b86b"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end
end
