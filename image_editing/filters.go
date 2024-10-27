package image_editing

import (
	"image"
	"image/color"
	"math"
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
	kernel := generateGaussianKernel(sigma, maskSize)
	bounds := img.Bounds()
	resultImg := image.NewNRGBA(bounds)

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			sumR, sumG, sumB := 0.0, 0.0, 0.0
			for i := 0; i < maskSize; i++ {
				xPos := limit(x+i-maskSize/2, bounds.Min.X, bounds.Max.X-1)
				yPos := limit(y+i-maskSize/2, bounds.Min.Y, bounds.Max.Y-1)
				r, g, b, _ := img.At(xPos, yPos).RGBA()
				weight := kernel[i]
				sumR += float64(r) / 256 * weight
				sumG += float64(g) / 256 * weight
				sumB += float64(b) / 256 * weight
			}
			valueR, valueG, valueB := limit(int(sumR), 0, 255), limit(int(sumG), 0, 255), limit(int(sumB), 0, 255)
			resultImg.Set(x, y, color.NRGBA{R: uint8(valueR), G: uint8(valueG), B: uint8(valueB), A: 255})
		}
	}
	return resultImg
}

// generateGaussianKernel generates a 1D Gaussian kernel based on sigma
func generateGaussianKernel(sigma float64, maskSize int) []float64 {
	kernel := make([]float64, maskSize)
	center := maskSize / 2
	sum := 0.0
	for i := 0; i < maskSize; i++ {
		x := i - center
		exp := math.Exp(-float64(x*x) / (2 * sigma * sigma))
		kernel[i] = exp
		sum += exp
	}
	// Normalize the kernel
	for i := range kernel {
		kernel[i] /= sum
	}
	return kernel
}

// FilterSobelBorderDetection applies Sobel edge detection filter to the image
func FilterSobelBorderDetection(img image.Image) image.Image {
	grayImg := FilterGrayScale(img)
	resultImage := image.NewGray(img.Bounds())

	sobelX := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	sobelY := [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	for x := 1; x < img.Bounds().Dx()-1; x++ {
		for y := 1; y < img.Bounds().Dy()-1; y++ {
			pixelX := 0

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					gx := x + j - 1
					gy := y + i - 1
					r, g, b, _ := grayImg.At(gx, gy).RGBA()
					pixelX += sobelX[i][j] * int((r+g+b)/3)
				}
			}

			pixelY := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					gx := x + j - 1
					gy := y + i - 1
					r, g, b, _ := grayImg.At(gx, gy).RGBA()
					pixelY += sobelY[i][j] * int((r+g+b)/3)
				}
			}

			val := math.Max(0, math.Floor(math.Sqrt(float64(pixelX*pixelX+pixelY*pixelY))))
			if val > 255 {
				val = 255
			}
			resultImage.Set(x, y, color.Gray{Y: uint8(val)})
		}
	}

	return resultImage
}
