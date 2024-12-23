package models

import (
	"fmt"
	"image"
)

// Project :
// An abstraction of a project structure. It contains the current preview image that will receive updates and an image
// of how the preview was before applying the changes.
type Project struct {
	versions       int
	originalImage  image.Image
	previewImage   image.Image
	previousStates ChangesStack
	nextStates     ChangesStack
}

func (p *Project) String() string {
	return fmt.Sprintf(
		"Project{versions: %d, originalImage: %v, previewImage: %v}",
		p.versions, p.originalImage, p.previewImage,
	)
}

// Collecting Values ---------------------------------------------------------------------------------------------------

// GetOriginal :
// Returns the current original image.
func (p *Project) GetOriginal() image.Image {
	return p.originalImage
}

// GetPreview :
// Returns the current preview image.
func (p *Project) GetPreview() image.Image {
	return p.previewImage
}

// PreviousPreviewImage :
// Cycles back into the project, returning the previous version of p.previewImage, if it exists.
func (p *Project) PreviousPreviewImage() (image.Image, error) {
	// get the previous preview image
	previousPreview := p.previousStates.Pop()

	// check if there is a previous preview image
	if previousPreview == nil {
		return p.previewImage, fmt.Errorf("no previous state")
	}

	// get the original image for the previous preview
	previousOriginal := p.previousStates.Pop()

	// if no original image exist, it means the original should be set the same as the preview
	if previousOriginal == nil {
		p.originalImage = previousPreview.(image.Image)
	} else {
		// if there IS an original image for the preview, pass it to p.originalImage
		p.originalImage = previousOriginal.(image.Image)

		// and then send it back to p.previousStates
		p.previousStates.Push(previousOriginal)
	}

	// pass the current image of p.previewImage to the stack of p.nextStates
	p.nextStates.Push(p.previewImage)
	p.previewImage = previousPreview.(image.Image)
	return p.previewImage, nil
}

func (p *Project) NextPreviewImage() (image.Image, error) {
	nextPreview := p.nextStates.Pop()

	if nextPreview == nil {
		return p.previewImage, fmt.Errorf("no next state")
	}

	p.originalImage = p.previewImage
	p.previousStates.Push(p.previewImage)
	p.previewImage = nextPreview.(image.Image)
	return p.previewImage, nil
}

// Updating Values -----------------------------------------------------------------------------------------------------

// AddPreviewImage :
// Adds a new preview image to the project and sets it as the current preview image. The Project attribute originalImage
// receives the previous preview image value, and then the previous image is pushed into p.previousStates.
//
// WARNING: p.nextStates will be cleared if it is not empty.
func (p *Project) AddPreviewImage(img image.Image) {
	// if p.nextStates is empty, it means that there is no work to reset
	if p.nextStates.Empty() {
		// set p.originalImage to p.previewImage
		p.originalImage = p.previewImage
		// push p.previewImage to p.previousStates
		p.previousStates.Push(p.previewImage)
		// set p.previewImage to img
		p.previewImage = img
		// increment p.versions
		p.versions++
	} else {
		// if the stack is not empty, p.nextStates should be cleared, because the work that had been done needs to be
		// overwritten

		// collect the amount of elements in p.nextStates to calculate the new versions amount
		previousVersions := p.nextStates.Length()
		// clear p.nextStates stack
		p.nextStates.Clear()
		// set p.originalImage to p.previewImage
		p.originalImage = p.previewImage
		// push p.previewImage to p.previousStates
		p.previousStates.Push(p.previewImage)
		// set p.previewImage to img
		p.previewImage = img
		// calculate the new p.versions value
		p.versions -= previousVersions + 1
	}
}

// LoadNewImage :
// Resets the project when called, setting p.versions to 1, clearing the stacks and setting the new image as both the
// original and previews.
func (p *Project) LoadNewImage(img image.Image) {
	// reset the stacks and the versions count
	p.versions = 1
	p.previousStates.Clear()
	p.nextStates.Clear()

	// pass img to the project
	p.originalImage = img
	p.previewImage = img
}

// Constructor ---------------------------------------------------------------------------------------------------------

// NewProject :
// Returns an instance of models.Project with a base image, used to set the size of the images container in Fyne.
func NewProject() *Project {
	initialImg := image.NewRGBA(image.Rect(0, 0, 512, 512))
	return &Project{
		versions:       1,
		originalImage:  initialImg,
		previewImage:   initialImg,
		previousStates: ChangesStack{},
		nextStates:     ChangesStack{},
	}
}
