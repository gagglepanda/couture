# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Couture < Formula
  desc "Allows for tailing multiple event sources."
  homepage ""
  version "0.0.9"
  bottle :unneeded

  if OS.mac? && Hardware::CPU.intel?
    url "https://github.com/gaggle-net/couture/releases/download/v0.0.9/couture_0.0.9_Darwin_x86_64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
    sha256 "05da100dc74ef6709ec4120939ea4b101a2dc2c44a86e67c34bb9be45333780d"
  end
  if OS.mac? && Hardware::CPU.arm?
    url "https://github.com/gaggle-net/couture/releases/download/v0.0.9/couture_0.0.9_Darwin_arm64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
    sha256 "54c231f2d598984782c4c9c01e846d51ac9e566231c0e4ef4462ac82fbe4a10b"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/gaggle-net/couture/releases/download/v0.0.9/couture_0.0.9_Linux_x86_64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
    sha256 "a89457d90d78d142c841f78299b8d5b5611b0699f890656e3f87a90e05fa018b"
  end
  if OS.linux? && Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
    url "https://github.com/gaggle-net/couture/releases/download/v0.0.9/couture_0.0.9_Linux_arm64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
    sha256 "5dcb5dcdb9e262df5c33aa5473d90083adb5ffa6b10bf02e22f068ba8723531e"
  end

  def install
    bin.install "couture"
  end
end
