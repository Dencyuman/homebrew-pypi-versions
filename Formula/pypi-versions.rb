# Formula/pypi-versions.rb
class PypiVersions < Formula
  desc "PyPI Versions is a CLI tool to fetch package versions from PyPI."
  homepage "https://github.com/Dencyuman/pypi-versions"
  url "https://github.com/Dencyuman/pypi-versions/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "a7ecde589d3fe202bc6bc19ddaa6c0cde118ef2c4833aad5d61c8ea6e4ea605c"
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