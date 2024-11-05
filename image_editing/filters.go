package image_editing

import (
	"image"
	"image/color"
	"math"
	"sort"
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
func FilterContrast(img image.Image, contrast float64) image.Image {
	contrastImage := image.NewRGBA(img.Bounds())

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, a := img.At(x, y).RGBA()

			newR := uint8(getContrastedChannelVal(r, contrast))
			newG := uint8(getContrastedChannelVal(g, contrast))
			newB := uint8(getContrastedChannelVal(b, contrast))

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

func getContrastedChannelVal(x uint32, contrast float64) float64 {
	temp := (float64(x) * contrast) / 256
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
func FilterMedianBlur(img image.Image, filterSize int) image.Image {
	resultImg := image.NewRGBA(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			pixel := medianFilterPixel(img, x, y, filterSize)
			resultImg.Set(x, y, pixel)
		}
	}
	return resultImg
}

func medianFilterPixel(img image.Image, row, column, filterSize int) color.Color {
	var values []color.Color
	for i := -filterSize / 2; i <= filterSize/2; i++ {
		for j := -filterSize / 2; j <= filterSize/2; j++ {
			x := row + i
			y := column + j

			if x >= 0 && x < img.Bounds().Dx() && y >= 0 && y < img.Bounds().Dy() {
				values = append(values, img.At(x, y))
			}
		}
	}

	middle := len(values) / 2
	// std::nth_element(values.begin(), values.begin() + middle, values.end(), comparePixelsByRGB);
	sort.Slice(values, func(i, j int) bool {
		r1, g1, b1, _ := values[i].RGBA()
		sumA := r1 + g1 + b1
		r2, g2, b2, _ := values[j].RGBA()
		sumB := r2 + g2 + b2
		return sumA < sumB
	})

	return values[middle]
}

// FilterGaussianBlur :
// Applies the Gaussian blur filter to the image.
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

func limit(x, minX, maxX int) int {
	if x < minX {
		return minX
	} else if x > maxX {
		return maxX
	}
	return x
}

// FilterSobelBorderDetection applies Sobel edge detection filter to the image
func FilterSobelBorderDetection(img image.Image, threshold float64) image.Image {
	grayImg := FilterGrayScale(img)
	resultImage := image.NewGray(img.Bounds())
	xKernel := [][]int{
		{1, 0, -1},
		{2, 0, -2},
		{1, 0, -1},
	}
	yKernel := [][]int{
		{1, 2, 1},
		{0, 0, 0},
		{-1, -2, -1},
	}
	widthM1 := grayImg.Bounds().Dx() - 1
	heightM1 := grayImg.Bounds().Dy() - 1
	for x := 1; x < heightM1; x++ {
		for y := 1; y < widthM1; y++ {
			gx := 0
			gy := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					r, g, b, _ := grayImg.At(x+(i-1), y+(j-1)).RGBA()
					r1 := int(r) * xKernel[i][j]
					g1 := int(g) * xKernel[i][j]
					b1 := int(b) * xKernel[i][j]

					gx += r1 + g1 + b1

					r2 := int(r) * yKernel[i][j]
					g2 := int(g) * yKernel[i][j]
					b2 := int(b) * yKernel[i][j]

					gy += r2 + g2 + b2
				}
			}
			g := math.Sqrt(math.Pow(float64(gx), 2) + math.Pow(float64(gy), 2))
			p := 0

			if (g / 256) > threshold {
				p = 255
			}

			resultImage.Set(x, y, color.Gray{Y: uint8(p)})
		}
	}
	return resultImage
}
