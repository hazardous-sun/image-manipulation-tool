# Image Manipulation Tool

## About

This is project is being developed as tool to manipulate some aspects of an image. When finished, this project is 
expected to contain a GUI where the user will be able to apply filters and manipulate images.

### TODO

- [X] Basic GUI
- [X] Open image
- [ ] Save image
- [ ] Geometric transformations
  - [ ] Translate
  - [ ] Rotate
  - [ ] Horizontal mirroring
  - [ ] Vertical Mirroring
  - [ ] Resize
  - [ ] Increase
  - [ ] Decrease
- [ ] Filters
  - [ ] Grayscale
  - [ ] Low fade
  - [ ] High fade
  - [ ] Threshold
- [ ] Mathematical Morfology
  - [ ] Dilation
  - [ ] Erosion
  - [ ] Opening
  - [ ] Closing
- [ ] Feature extraction

## Wails

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
