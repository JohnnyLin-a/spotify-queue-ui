#!/bin/bash
docker buildx build --push -t ghcr.io/johnnylin-a/spotify-queue-ui --platform=linux/arm64,linux/arm/v7,linux/amd64 .