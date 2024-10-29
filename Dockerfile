# Pull container image for building applications with Go
FROM docker.io/golang:1-bullseye AS build
LABEL authors="solaire"
LABEL description="A tool for digital image processing. Supports filters, geometric transformations and mathematical morphology manipulation."

# creates a dir called /app and sets it as the CWD
WORKDIR /app

# clone all the files from the project to the container

ENTRYPOINT ["top", "-b"]