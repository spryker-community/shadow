#!/bin/sh
# Usage: [sudo] [BIN_DIR=/usr/local/bin] ./install.sh [<BIN_DIR>]
#
# Example:
#     1. sudo ./install.sh /usr/local/bin
#     2. sudo ./install.sh /usr/bin
#     3. ./install.sh $HOME/usr/bin
#     4. BIN_DIR=$HOME/usr/bin ./install.sh
#
# Default BIN_DIR=/usr/bin

set -euf

if [ -n "${DEBUG-}" ]; then
  set -x
fi

: "${BIN_DIR:=/usr/bin}"

if [ $# -gt 0 ]; then
  BIN_DIR=$1
fi

_can_install() {
  if [ ! -d "${BIN_DIR}" ]; then
    mkdir -p "${BIN_DIR}" 2>/dev/null
  fi
  [ -d "${BIN_DIR}" ] && [ -w "${BIN_DIR}" ]
}

if ! _can_install && [ "$(id -u)" != 0 ]; then
  printf "Run script as sudo\n"
  exit 1
fi

if ! _can_install; then
  printf -- "Can't install to %s\n" "${BIN_DIR}"
  exit 1
fi

kernel=$(uname -s 2>/dev/null || /usr/bin/uname -s)
case ${kernel} in
"Linux" | "linux")
  kernel="linux"
  ;;
"Darwin" | "darwin")
  kernel="darwin"
  ;;
*)
  printf -- "OS '%s' not supported\n" "${kernel}"
  exit 1
  ;;
esac

machine=$(uname -m 2>/dev/null || /usr/bin/uname -m)
case ${machine} in
arm | armv7*)
  machine="arm"
  ;;
aarch64* | armv8* | arm64)
  machine="arm64"
  ;;
i[36]86)
  machine="386"
  ;;
x86_64)
  machine="amd64"
  ;;
*)
  printf -- "Your architecture '%s' is not supported\n" "${machine}"
  exit 1
  ;;
esac

if [ "darwin" = "${kernel}" ]; then
    machine="all"
fi

tmpFolder="/tmp/shadow-$(date +"%s")"

mkdir -p "${tmpFolder}" 2> /dev/null

printf -- "Downloading shadow_%s_%s.tar.gz\n" "${kernel}" "${machine}"
curl -sL "https://github.com/andreaspenz/shadow/releases/latest/download/shadow_${kernel}_${machine}.tar.gz" | tar -C "${tmpFolder}" -xzf -

printf -- "Installing...\n"
install -m755 "${tmpFolder}/shadow" "${BIN_DIR}/shadow"

printf "Cleaning up temp files\n"
rm -rf "${tmpFolder}"

printf -- "Successfully installed shadow into %s/\n" "${BIN_DIR}"