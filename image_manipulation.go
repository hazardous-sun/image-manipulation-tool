package main

import "image"

type DisplayedImages struct {
	originPath    string
	originalImage image.Image
	previewPath   string
	previewImage  image.Image
}
