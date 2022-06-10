#!/bin/bash

# BEFORE RUN:
# sudo apt-get update && sudo apt-get upgrade
# sudo apt-get install clang libstdc++6
# sudo update-alternatives --config c++
# sudo update-alternatives --config cc
# ldconfig

set +x
set -e

# see https://root.cern.ch/cling-build-instructions
# see https://github.com/alandefreitas/find_package_online/blob/master/Modules/ExternalProjectCling.cmake

# mkdir from scripts/
mkdir -p $1
cd $1

git clone http://root.cern.ch/git/llvm.git src
cd src
git checkout cling-patches
cd tools
git clone http://root.cern.ch/git/cling.git
#cd cling
#git checkout master
#cd ..
git clone http://root.cern.ch/git/clang.git
cd clang
git checkout cling-patches
