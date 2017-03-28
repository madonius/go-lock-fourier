package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"fmt"
	"math"
)

func From2D(x int, y int, YMax int) (int, error) {
	if ( y >= YMax ) {
		return 0, fmt.Errorf("Matrix Index y out of image bounds %d", y)
	}

	return x*YMax+y, nil
}

func To2D(n int, YMax int) [2]int {
	CalcY := math.Mod(float64(n),float64(YMax))
	CalcX := (n-int(CalcY))/YMax

	return [2]int{int(CalcX), int(CalcY)}
}

func GetImage( image_path string ) image.Image {
	file, err := os.Open(image_path)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return img
}

func img_to_array( img image.Image ) []uint32 {
	XImgSize := img.Bounds().Max.X - img.Bounds().Min.X
	YImgSize := img.Bounds().Max.Y - img.Bounds().Min.Y
	Size := XImgSize*YImgSize

	var LinearImage = make([]uint32, Size)

	for x:=0; x < XImgSize; x++ {
		for y:=0; y < YImgSize; y++ {
			colour := img.At(img.Bounds().Min.X+x,img.Bounds().Min.Y+y)
			r,b,g,_ := colour.RGBA()
			pos, err := From2D(x,y,YImgSize)
			if err != nil {
				log.Fatal(err)
			}
			LinearImage[pos] = uint32(r/3+b/3+g/3)
		}
	}
	return LinearImage
}

func main() {
	in_image_path := os.Args[1]
	img := GetImage(in_image_path)
	ImageMatrix := img_to_array(img)
	for i:=0; i<len(ImageMatrix); i++ {
		fmt.Println(To2D(i, img.Bounds().Max.Y))
	}
}
