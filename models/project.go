package models

import (
	"fmt"
	"image"
)

type Project struct {
	versions       int
	currentVersion int
	originalImage  image.Image
	previewImage   []image.Image
}

func (p *Project) String() string {
	return fmt.Sprintf(
		"Project{versions: %d, currentVersion: %d, originalImage: %v, previewImage: %v}",
		p.versions, p.currentVersion, p.originalImage, p.previewImage,
	)
}

// Collecting Values ---------------------------------------------------------------------------------------------------

func (p *Project) GetOriginal() image.Image {
	return p.originalImage
}

func (p *Project) GetPreview() image.Image {
	if len(p.previewImage) == 0 {
		return nil
	}
	return p.previewImage[p.currentVersion]
}

func (p *Project) PreviousPreviewImage() image.Image {
	if p.currentVersion == 0 {
		return p.previewImage[0]
	}

	index := p.currentVersion - 1

	if index > 0 {
		p.originalImage = p.previewImage[index-1]
	} else {
		p.originalImage = p.previewImage[index]
	}

	return p.previewImage[index]
}

// Updating Values -----------------------------------------------------------------------------------------------------

func (p *Project) AddPreviewImage(img image.Image) {
	if p.currentVersion < p.versions {
		p.previewImage = append(p.previewImage[:p.currentVersion], img)
		p.currentVersion++
		p.versions = p.currentVersion
	} else {
		p.previewImage = append(p.previewImage, img)
		p.currentVersion++
		p.versions++
	}
	p.originalImage = p.PreviousPreviewImage()
}

func (p *Project) LoadNewImage(img image.Image) {
	p.versions = 0
	p.currentVersion = 0
	p.originalImage = img
	p.previewImage = []image.Image{img}
}

// Constructor ---------------------------------------------------------------------------------------------------------

func NewProject() *Project {
	initialImg := image.NewRGBA(image.Rect(0, 0, 512, 512))
	return &Project{
		versions:       1,
		currentVersion: 0,
		originalImage:  initialImg,
		previewImage:   []image.Image{initialImg},
	}
}
