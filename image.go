package main

import (
	"image/jpeg"
	"log"
	"os"
	"fmt"
	"math"
)

func FromTwoD(x int, y int, YMax int) (int, error) {
	if ( y >= YMax ) {
		return 0, fmt.Errorf("Matrix Index y out of image bounds %d", y)
	}

	return x*YMax+y, nil
}

func ToTwoD(n int, YMax int) [2]int {
	CalcY := math.Mod(float64(n),float64(YMax))
	CalcX := (n-int(CalcY))/YMax

	return [2]int{int(CalcX), int(CalcY)}
}

func img_to_array( image_path string ) []int {
	file, err := os.Open(image_path)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	XImgSize := img.Bounds().Max.X - img.Bounds().Min.X
	YImgSize := img.Bounds().Max.Y - img.Bounds().Min.Y

	Size := XImgSize*YImgSize

	var LinearImage = make([]int, Size)

	for x:=0; x < XImgSize; x++ {
		for y:=0; y < YImgSize; y++ {
			colour := img.At(img.Bounds().Min.X+x,img.Bounds().Min.Y+y)
			r,g,b,_ := colour.RGBA()
			pos, err := FromTwoD(x,y,YImgSize)
			if err != nil {
				log.Fatal(err)
			}
			LinearImage[pos] = int((r+g+b)/3)
		}
	}

	return LinearImage
}

func main() {
	in_image_path := os.Args[1]
	ImageMatrix := img_to_array(in_image_path)
	fmt.Println("%d", ImageMatrix[70])
}
