#!/bin/sh

declare -a arr=("linux_386" "linux_amd64" "linux_arm" "linux_arm64" "darwin_amd64" "freebsd_386" "freebsd_amd64" "freebsd_arm" "netbsd_386" "netbsd_amd64" "netbsd_arm" "openbsd_386" "openbsd_amd64" "openbsd_arm" "openbsd_arm64" "windows_386" "windows_amd64" "plan9_386" "plan9_amd64" "plan9_arm")
for i in "${arr[@]}"
do
  export GOOS=$(echo $i | cut -f1 -d_)
  export GOARCH=$(echo $i | cut -f2 -d_)

  echo ""
  echo "=== Building ${GOOS} ${GOARCH} ==="

  go build -v .

  tar -czf ./uveira_${GOOS}_${GOARCH}.tar.gz uveira
  go clean
done
