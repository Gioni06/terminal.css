#!/bin/bash

# Define the app name and the list of platforms and architectures
APP_NAME="terminalcss-builder"
PLATFORMS=("windows" "linux" "darwin") # darwin is for macOS
ARCHITECTURES=("amd64" "arm64" "386") # 386 is for x86 (32-bit)

# Create a directory for the builds if it doesn't exist
mkdir -p builds

# Loop through platforms and architectures and build for each
for PLATFORM in "${PLATFORMS[@]}"; do
    for ARCH in "${ARCHITECTURES[@]}"; do
        # If Platform is darwin, then ARCH must be amd64 or arm64
        if [ "$PLATFORM" == "darwin" ] && [ "$ARCH" != "amd64" ] && [ "$ARCH" != "arm64" ]; then
            continue
        fi
        # Set the output binary name, including .exe suffix for Windows
        OUTPUT_NAME="./builds/$APP_NAME-$PLATFORM-$ARCH"
        if [ "$PLATFORM" == "windows" ]; then
            OUTPUT_NAME+=".exe"
        fi

        # Set environment variables for cross-compilation
        GOOS=$PLATFORM GOARCH=$ARCH go build -o $OUTPUT_NAME

        # Verify if build was successful
        if [ $? -ne 0 ]; then
            echo "An error has occurred! Aborting the script execution..."
            exit 1
        fi

        echo "Built $OUTPUT_NAME"
    done
done

echo "Build process completed."
