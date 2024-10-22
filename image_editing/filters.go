package image_editing

import (
	"image"
	"image/color"
)

// Filters -------------------------------------------------------------------------------------------------------------

// FilterGrayScale :
// Removes the color channels of an image and returns an image with only shades of gray.
func FilterGrayScale(img image.Image) image.Image {
	grayImage := image.NewGray(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			lum := 0.2125*float64(r) + 0.7154*float64(g) + 0.0721*float64(b)
			pixel := color.Gray{Y: uint8(lum / 256)}
			grayImage.Set(x, y, pixel)
		}
	}
	return grayImage
}

// FilterContrast :
// Applies an amount of contrast to the image.
func FilterContrast(img image.Image, factor float64) image.Image {
	if factor == 0 {
		return img
	}

	contrastImage := image.NewRGBA(img.Bounds())
	contrast := factor / 50
	if contrast < 0 {
		contrast *= -1
	}
	removedValue := 0.0
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, a := img.At(x, y).RGBA()

			newR := uint8(getContrastedChannelVal(r, contrast, removedValue))
			newG := uint8(getContrastedChannelVal(g, contrast, removedValue))
			newB := uint8(getContrastedChannelVal(b, contrast, removedValue))

			pixel := color.RGBA{
				R: newR,
				G: newG,
				B: newB,
				A: uint8(a),
			}
			contrastImage.Set(x, y, pixel)
		}
	}
	return contrastImage
}

func getContrastedChannelVal(x uint32, contrast float64, removedValue float64) float64 {
	temp := (float64(x)*contrast - removedValue) / 256
	if temp > 255 {
		temp = 255
	} else if temp < 0 {
		temp = 0
	}
	return temp
}

// FilterBrightness :
// Applies an amount of brightness to the image.
func FilterBrightness(img image.Image, brightness int64) image.Image {
	brightnessImage := image.NewRGBA(img.Bounds())

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, a := img.At(x, y).RGBA()

			newR := getBrightnessChannelVal(r, brightness)
			newG := getBrightnessChannelVal(g, brightness)
			newB := getBrightnessChannelVal(b, brightness)

			pixel := color.RGBA{
				R: uint8(newR),
				G: uint8(newG),
				B: uint8(newB),
				A: uint8(a),
			}

			brightnessImage.Set(x, y, pixel)
		}
	}
	return brightnessImage
}

func getBrightnessChannelVal(x uint32, brightness int64) int64 {
	temp := int64(x/255) + brightness
	if temp > 255 {
		temp = 255
	} else if temp < 0 {
		temp = 0
	}
	return temp
}

// FilterThreshold :
// Applies the threshold filter to the image.
func FilterThreshold(img image.Image, threshold uint32) image.Image {
	grayImage := FilterGrayScale(img)
	resultImage := image.NewGray(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			var pixel color.RGBA
			r, g, b, a := grayImage.At(x, y).RGBA()
			pixelValue := (r/256 + g/256 + b/256) / 3
			if pixelValue > threshold {
				pixel = color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: uint8(a),
				}
			} else {
				pixel = color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: uint8(a),
				}
			}
			resultImage.Set(x, y, pixel)
		}
	}
	return resultImage
}

// FilterMedianBlur :
// Applies the median blur filter to the image.
func FilterMedianBlur(img image.Image) image.Image {
	resultImg := image.NewRGBA(img.Bounds())
	bounds := img.Bounds()
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			neighbours := [][]color.RGBA{
				{color.RGBA{}, color.RGBA{}, color.RGBA{}},
				{color.RGBA{}, color.RGBA{}, color.RGBA{}},
				{color.RGBA{}, color.RGBA{}, color.RGBA{}},
			}
			for lx := 0; lx < 3; lx++ {
				for ly := 0; ly < 3; ly++ {
					ix := limit(x+lx-1, 0, bounds.Dx())
					iy := limit(y+ly-1, 0, bounds.Dy())

					r, g, b, a := img.At(ix, iy).RGBA()

					newR := uint8(r / 256)
					newG := uint8(g / 256)
					newB := uint8(b / 256)

					pixel := color.RGBA{
						R: newR,
						G: newG,
						B: newB,
						A: uint8(a),
					}
					neighbours[lx][ly] = pixel
				}
			}
			newValue := computeCenter(neighbours)
			r, g, b, a := newValue.RGBA()
			newR := uint8(r / 256)
			newG := uint8(g / 256)
			newB := uint8(b / 256)
			pixel := color.RGBA{
				R: newR,
				G: newG,
				B: newB,
				A: uint8(a),
			}
			resultImg.Set(x, y, pixel)
		}
	}
	return resultImg
}

func limit(x, minX, maxX int) int {
	if x < minX {
		return minX
	} else if x > maxX {
		return maxX
	}
	return x
}

func computeCenter(neighbours [][]color.RGBA) color.RGBA {
	values := make([]color.RGBA, len(neighbours)*len(neighbours[0]))
	for x := 0; x < len(neighbours); x++ {
		for y := 0; y < len(neighbours[x]); y++ {
			value := neighbours[x][y]
			values = insert(values, value)
		}
	}
	return values[len(values)/2]
}

func insert(values []color.RGBA, newValue color.RGBA) []color.RGBA {
	for i := 0; i < len(values); i++ {
		r, g, b, _ := values[i].RGBA()
		originChannelsSum := r + g + b

		r, g, b, _ = newValue.RGBA()
		newChannelsSum := r + g + b

		if i == len(values)-1 || newChannelsSum > originChannelsSum {
			for j := len(values) - 2; j >= i; j-- {
				values[j+1] = values[j]
			}
			values[i] = newValue
		}
	}
	return values
}

func FilterGaussianBlur(img image.Image, sigma float64, maskSize int) image.Image {
	kernel := [][]int{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}
	resultImage := image.NewRGBA(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			processKernel(kernel, img, resultImage, x, y, sigma, maskSize)
		}
	}
	return resultImage
}

func processKernel(kernel [][]int, sourceImg image.Image, resultImage *image.RGBA, x int, y int, sigma float64, maskSize int) {
	sumValue := 0.0
	valueKernel := 0.0
	kernelSize := len(kernel)
	bounds := sourceImg.Bounds()
	for i := 0; i < kernelSize; i++ {
		for j := 0; j < kernelSize; j++ {
			xPos := limit(x+(i-1), 0, bounds.Dx())
			yPos := limit(y+(j-1), 0, bounds.Dy())
			r, _, _, _ := sourceImg.At(xPos, yPos).RGBA()
			pixelValue := float64(r >> 8)
			sumValue += pixelValue * float64(kernel[i][j])
			valueKernel += float64(kernel[i][j])
		}
	}
	if valueKernel > 0 {
		sumValue /= valueKernel
	}
	value := limit(int(sumValue), 0, 255)
	resultImage.Set(x, y, color.RGBA{
		R: uint8(value),
		G: uint8(value),
		B: uint8(value),
		A: 255,
	})
}
