# Formula/pypi-versions.rb
class PypiVersions < Formula
  desc "PyPI Versions is a CLI tool to fetch package versions from PyPI."
  homepage "https://github.com/Dencyuman/pypi-versions"
  url "https://github.com/Dencyuman/pypi-versions/archive/refs/tags/v1.1.12.tar.gz"
  sha256 "bc4796dad5de47c65203d1b7d0b42390a20025f4d6e45a6a8c270026320ac984"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "ppv"
    bin.install "ppv"
  end

  test do
    assert_match "Available versions for", shell_output("#{bin}/ppv requests 2>&1", 0)
  end
end