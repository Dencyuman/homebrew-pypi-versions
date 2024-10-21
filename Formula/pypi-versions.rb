class PypiVersions < Formula
  desc "CLI tool to fetch package versions and metadata from PyPI."
  homepage "https://github.com/Dencyuman/ppv"
  url "https://github.com/Dencyuman/pypi-versions/archive/refs/tags/v1.1.11.tar.gz"
  sha256 "d5558cd419c8d46bdc958064cb97f963d1ea793866414c025906ec15033512ed"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end
end
