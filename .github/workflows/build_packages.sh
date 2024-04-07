#!/bin/bash

# Welcome message
echo "Pipeline build packages"

# Get args
if [ -z "$1" ]; then
    release_version="1.0.0"
else
    release_version=$1
fi

if [ -z "$2" ]; then
    build_dir="../.build"
else
    build_dir=$2
fi

if [ -z "$3" ]; then
    release_name="Helm-Registry"
else
    release_name=$3
fi

# Release prefix
release_prefix="${release_name}-${release_version}"
echo "Release prefix: ${release_prefix}"

# Set location to the backend file 
cd backend

# Build for Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -ldflags "${release_ldflags}" -o "${build_dir}/${release_prefix}_windows.exe"
if [ $? -eq 0 ]; then
    echo "Build for Windows successful"
    echo "> Artifact build in    ${build_dir}/${release_prefix}_windows.exe"
else
    echo "Failed to build for Windows"
    exit 1
fi

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -ldflags "${release_ldflags}" -o "${build_dir}/${release_prefix}"
if [ $? -eq 0 ]; then
    echo "Build for Linux successful"
    echo "> Artifact build in    ${build_dir}/${release_prefix}"

    # Zip
    tar -czvf "${build_dir}/${release_prefix}_linux.tar.gz" -C "${build_dir}" "${release_prefix}"
else
    echo "Failed to build for Linux"
    exit 1
fi

# Build for Docker
echo "Building Docker image..."
docker build -t "${release_name,,}:${release_version}" .
if [ $? -eq 0 ]; then
    echo "Docker image build successful"

    # Save Docker image
    echo "Save Docker image..."
    docker save "${release_name,,}:${release_version}" -o "${build_dir}/${release_prefix}_docker.tar"

    if [ $? -eq 0 ]; then
        echo "Docker image saved"
        echo "> Artifact build in    ${build_dir}/${release_prefix}_docker.tar"
        exit 0
    else
        echo "Failed to save Docker image"
        exit 1
    fi

else
    echo "Failed to build Docker image"
    exit 1
fi

echo "All builds completed successfully"
