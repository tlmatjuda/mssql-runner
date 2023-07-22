#!/bin/bash


# BASE CONFIGURATIONS
# =====================================================================================================================
if [ -z "${BASH_VERSINFO+x}" ]; then
  BUILD_MAC_SILICON_SCRIPT_PATH=${0:a:h}
else
  BUILD_MAC_SILICON_SCRIPT_PATH=$(cd "$(dirname ${BASH_SOURCE[0]})" && pwd)
fi

# Getting the Root Path of this project.
# This is so that we can go there and compile our Go Module from the root path.
PROJECT_ROOT_PATH=$(dirname ${BUILD_MAC_SILICON_SCRIPT_PATH})
cd ${PROJECT_ROOT_PATH}

# Now compiling for windows
echo "Compiling for Mac Os Silicon while in directory : ${PROJECT_ROOT_PATH}"
env GOOS=darwin GOARCH=arm64 go build -o .

echo "Compilation done and ready for distribution"
