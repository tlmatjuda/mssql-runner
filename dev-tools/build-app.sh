#!/bin/bash


# BASE CONFIGURATIONS
# =====================================================================================================================
if [ -z "${BASH_VERSINFO+x}" ]; then
  BUILD_WINDOWS_SCRIPT_PATH=${0:a:h}
else
  BUILD_WINDOWS_SCRIPT_PATH=$(cd "$(dirname ${BASH_SOURCE[0]})" && pwd)
fi

# Getting the Root Path of this project.
# This is so that we can go there and compile our Go Module from the root path.
PROJECT_ROOT_PATH=$(dirname ${BUILD_WINDOWS_SCRIPT_PATH})
cd ${PROJECT_ROOT_PATH}


echo "Currently in path : ${PROJECT_ROOT_PATH}"

# Now compiling for windows first
echo "Compiling for Windows OS x64"
env GOOS=windows GOARCH=amd64 go build .

# And then for Mac OS M Series
echo "Compiling for Mac OS M-Series Silicon"
env GOOS=darwin GOARCH=arm64 go build -o .

echo "Compilation done"
