# Formula/pypi-versions.rb
class PypiVersions < Formula
  desc "PyPI Versions is a CLI tool to fetch package versions from PyPI."
  homepage "https://github.com/Dencyuman/pypi-versions"
  url "https://github.com/Dencyuman/pypi-versions/archive/refs/tags/v1.1.0.tar.gz"
  sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5"
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