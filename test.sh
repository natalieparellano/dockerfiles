# clean directory
echo "cleaning directory..."
rm -rf ./acceptance/in

mkdir ./acceptance/in

# build binary in container
echo "building image..."
docker build -t test-kaniko .

# execute binary in container
echo "executing image..."
docker run \
			--rm \
			-v $PWD/acceptance/testdata:/workspace \
			test-kaniko
