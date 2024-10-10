# Image Manipulation Tool

## About

This is project is being developed as tool to manipulate some aspects of an image. When finished, this project is
expected to contain a GUI where the user will be able to apply filters and manipulate images.

### TODO

- [X] Basic GUI
- [X] Separate the build process into its own file
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
            - [X] Increase
            - [X] Decrease
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

### Known issues

- I noticed there seems to be a problem with the function that opens and saves images when running the program o
  Windows, but since I am currently developing the application in a Linux environment, it may take a bit of time for me
  to address what exactly is causing the error
- I found a behavior that does not allow images with the same file extension to be loaded, but I am still figuring out
  what is the root cause of this

## Requirements

### Go

The machine should have Go version 1.23 or higher installed.

### Wails

This is the official Wails Vanilla template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
