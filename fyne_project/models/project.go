package models

import "image"

type Project struct {
	versions       int
	currentVersion int
	originalImage  image.Image
	previewImage   []image.Image
}

// Collecting Values ---------------------------------------------------------------------------------------------------

func (p *Project) GetPreview() image.Image {
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
	return &Project{
		versions:       0,
		currentVersion: 0,
		originalImage:  nil,
		previewImage:   []image.Image{},
	}
}
