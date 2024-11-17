package image_editing

import (
	"image"
	"image/color"
)

func MathMorpDilation(img image.Image) image.Image {
	resultImg := image.NewRGBA(img.Bounds())
	kernel := [][]int{
		{0, 10, 0},
		{10, 10, 10},
		{0, 10, 0},
	}

	// iterate over each element in the image
	for x := 1; x < resultImg.Bounds().Dx(); x++ {
		for y := 1; y < resultImg.Bounds().Dy(); y++ {
			temp := [][][]uint8{
				{
					{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
				},
				{
					{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
				},
				{
					{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
				},
			}

			// ITERAR SOBRE OS VALORES AO REDOR DO PIXEL ATUAL QUE ENGLOBAM O KERNEL
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					r, g, b, _ := img.At(x+i, y+j).RGBA()

					r /= 256
					r += uint32(kernel[1-i][1-j])
					if r >= 255 {
						r = 255
					}

					g /= 256
					g += uint32(kernel[1-i][1-j])
					if g >= 255 {
						g = 255
					}

					b /= 256
					b += uint32(kernel[1-i][1-j])
					if b >= 255 {
						b = 255
					}

					temp[1-i][1-j][0] = uint8(r)
					temp[1-i][1-j][1] = uint8(g)
					temp[1-i][1-j][2] = uint8(b)
				}
			}

			// checar qual o maior valor e manter ele
			greatest := uint32(255)
			index := []int{0, 0}
			for i := 0; i < len(temp); i++ {
				for j := 0; j < len(temp[i]); j++ {
					r, g, b, _ := img.At(x+i, y+j).RGBA()
					temp := (r / 256) + (g / 256) + (b / 256)
					if temp < greatest {
						greatest = temp
						index = []int{i, j}
					}
				}
			}

			// definir todos os valores ao redor do pixel como greatest
			r, g, b, a := img.At(x+index[0], y+index[1]).RGBA()
			pixel := color.NRGBA{
				R: uint8(r / 256),
				G: uint8(g / 256),
				B: uint8(b / 256),
				A: uint8(a / 256),
			}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					resultImg.Set(x+i, y+j, pixel)
				}
			}
		}
	}

	return resultImg
}

func MathMorpErosion(img image.Image) image.Image {
	resultImg := image.NewRGBA(img.Bounds())
	kernel := [][]int{
		{0, 10, 0},
		{10, 10, 10},
		{0, 10, 0},
	}

	// iterate over each element in the image
	for x := 1; x < resultImg.Bounds().Dx(); x++ {
		for y := 1; y < resultImg.Bounds().Dy(); y++ {
			temp := [][][]uint8{
				{
					{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
				},
				{
					{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
				},
				{
					{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
				},
			}

			// ITERAR SOBRE OS VALORES AO REDOR DO PIXEL ATUAL QUE ENGLOBAM O KERNEL
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					r, g, b, _ := img.At(x+i, y+j).RGBA()

					r /= 256
					r += uint32(kernel[1-i][1-j])
					if r >= 255 {
						r = 255
					}

					g /= 256
					g += uint32(kernel[1-i][1-j])
					if g >= 255 {
						g = 255
					}

					b /= 256
					b += uint32(kernel[1-i][1-j])
					if b >= 255 {
						b = 255
					}

					temp[1-i][1-j][0] = uint8(r)
					temp[1-i][1-j][1] = uint8(g)
					temp[1-i][1-j][2] = uint8(b)
				}
			}

			// checar qual o maior valor e manter ele
			greatest := uint32(0)
			index := []int{0, 0}
			for i := 0; i < len(temp); i++ {
				for j := 0; j < len(temp[i]); j++ {
					r, g, b, _ := img.At(x+i, y+j).RGBA()
					temp := (r / 256) + (g / 256) + (b / 256)
					if temp > greatest {
						greatest = temp
						index = []int{i, j}
					}
				}
			}

			// definir todos os valores ao redor do pixel como greatest
			r, g, b, a := img.At(x+index[0], y+index[1]).RGBA()
			pixel := color.NRGBA{
				R: uint8(r / 256),
				G: uint8(g / 256),
				B: uint8(b / 256),
				A: uint8(a / 256),
			}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					resultImg.Set(x+i, y+j, pixel)
				}
			}
		}
	}

	return resultImg
}
