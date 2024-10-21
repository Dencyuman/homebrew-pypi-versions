# Formula/pypi-versions.rb
class PypiVersions < Formula
  desc "PyPI Versions is a CLI tool to fetch package versions from PyPI."
  homepage "https://github.com/Dencyuman/pypi-versions"
  url "https://github.com/Dencyuman/pypi-versions/archive/refs/tags/v1.1.13.tar.gz"
  sha256 "fbe37224c5016346b91fd0bfa17e17e23fb446b65405822019c1835cb0e7fac1"
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