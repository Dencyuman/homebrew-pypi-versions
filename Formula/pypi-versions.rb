# Formula/pypi-versions.rb
class PypiVersions < Formula
  desc "PyPI Versions is a CLI tool to fetch package versions from PyPI."
  homepage "https://github.com/Dencyuman/pypi-versions"
  url "https://github.com/Dencyuman/pypi-versions/archive/refs/tags/v1.1.14.tar.gz"
  sha256 "5f7d330eeb9bbcfffa175dae542447724ba2e4741dae95bf05ee68a79452aea6"
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