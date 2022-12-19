
DIR=build
if [ -d "$DIR" ];
then
    echo "$DIR directory exists."
else
    mkdir $DI
	echo "$DIR directory created."
fi

cd build
echo "Configure"
cmake ..
echo "Compile c++ layer"
ninja
ninja install
cd ..
echo "Import golang dependency"
go mod tidy
echo "Build golang layer"
go build
echo "Done!"