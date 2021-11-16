# clean directory
rm -rf ./acceptance/in

mkdir ./acceptance/in

# build binary
GOOS=linux go build -o ./acceptance/testdata/dockerfiles .

# execute binary in container
docker run \
			--rm \
			--env PATH=/usr/local/bin:/kaniko \
			--env HOME=/root \
			--env USER=root \
			--env SSL_CERT_DIR=/kaniko/ssl/certs \
			--env DOCKER_CONFIG=/kaniko/.docker/ \
			--env DOCKER_CREDENTIAL_GCR_CONFIG=/kaniko/.config/gcloud/docker_credential_gcr_config.json \
			-v $PWD/acceptance/testdata:/workspace \
			-v $PWD/acceptance/in:/kaniko \
			golang \
			/workspace/dockerfiles /workspace/Dockerfile tarball