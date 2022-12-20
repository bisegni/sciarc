#!/bin/bash
DIR=build
if [ -d "$DIR" ];
then
    echo "$DIR directory exists."
else
    mkdir $DIR
	echo "$DIR directory created."
fi

cd $DIR
echo "Configure"
pwd
cmake -GNinja ..
RESULT=$?
if [ $RESULT -eq 1 ]; then
  echo failed to Configure
  return 1
fi
echo "Compile c++ layer"
ninja
ninja install
cd ..

echo "Import golang dependency"
go mod tidy
echo "Build golang layer"
go build
echo "Done!"