package models

import (
	"fmt"
	"image"
	"image-manipulation-tool/fyne_project/file_handling"
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
	p.currentVersion--
	if p.currentVersion > 0 {
		p.originalImage = p.previewImage[p.currentVersion-1]
	} else {
		p.originalImage = p.previewImage[0]
	}
	return p.previewImage[p.currentVersion]
}

// Updating Values -----------------------------------------------------------------------------------------------------

func (p *Project) AddPreviewImage(img image.Image) {
	if p.currentVersion < p.versions {
		p.UpdatePreviewImage(img)
		p.versions = p.currentVersion + 1
	} else {
		p.previewImage = append(p.previewImage, img)
		p.versions++
	}
	p.originalImage = p.PreviousPreviewImage()
}

func (p *Project) UpdatePreviewImage(img image.Image) {
	p.previewImage = append(p.previewImage[:p.currentVersion], img)
}

// Constructor ---------------------------------------------------------------------------------------------------------

func NewProject() *Project {
	initialImg, err := file_handling.LoadImage("fyne_project/cat.jpeg")

	if err != nil {
		panic(err)
	}

	return &Project{
		versions:       1,
		currentVersion: 0,
		originalImage:  initialImg,
		previewImage:   []image.Image{initialImg},
	}
}
