package main

import (
	"image/jpeg"
	"log"
	"os"
	"fmt"
	"bytes"
)

func open_image( img_file string ) []byte {
	file, err := os.Open(img_file)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func main() {
	in_image_path := os.Args[1]
	image_array := open_image(in_image_path)
	
}
