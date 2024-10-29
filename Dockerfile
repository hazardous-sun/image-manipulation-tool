# Pull container image for building applications with Go
FROM docker.io/golang:1-bullseye AS build
LABEL authors="solaire"
LABEL description="A tool for digital image processing. Supports filters, geometric transformations and mathematical morphology manipulation."

# update the system
RUN apt-get update && apt-get upgrade -y && apt-get install libx11-dev libgl1-mesa-glx mesa-common-devz libgl1-mesa-dev xorg-dev -y

# creates a dir called /app and sets it as the CWD
WORKDIR /app

# clone all the files from the project to the container
COPY . .

# build the project
RUN go build -o image-manipulation-tool .

# Pull a smaller image
FROM docker.io/alpine

# create an "app" dir in the new image
WORKDIR /app

# Copy the binary from the build environment to the smaller environment
COPY --from=build /app/image-manipulation-tool .

CMD [ "./image-manipulation-tool" ]