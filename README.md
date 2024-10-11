# Image Manipulation Tool

## About

This is project is being developed as tool to manipulate some aspects of an image. When finished, this project is
expected to contain a GUI where the user will be able to apply filters and manipulate images.

### TODO

- [X] Basic GUI
- [ ] Theme customization
- Tools:
    - [X] Open image
    - [X] Save image
    - [X] Geometric transformations
        - [X] Translate
        - [X] Rotate
        - [X] Horizontal mirroring
        - [X] Vertical Mirroring
        - [X] Resize
    - [ ] Filters
        - [ ] Grayscale
        - [ ] Low fade
        - [ ] High fade
        - [ ] Threshold
    - [ ] Mathematical Morphology
        - [ ] Dilation
        - [ ] Erosion
        - [ ] Opening
        - [ ] Closing
    - [ ] Feature extraction

## Requirements

- [go](https://go.dev/dl/) : `1.23.0+`
- [fyne](https://fyne.io/) : `2.5.1+`

## Supported image file types

* JPEG\JPG
* PNG

## Supported image file types

* JPEG\JPG
* PNG

## Running the project

To run the project, first clone the repository, then follow these steps:

1. `cd path/to/the/project`
2. `fyne build` or simply run `go build .`
3. `./path/to/the/binary/file`
