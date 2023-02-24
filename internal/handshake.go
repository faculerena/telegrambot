package internal

import (
	"fmt"
	gg "github.com/fogleman/gg"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func Handshake(a string, b string, c string) error {
	// Open the image file
	imgFile, err := os.Open("/assets/handshake.jpg")
	if err != nil {
		log.Println("Error opening image file:", err)
		return err
	}
	defer imgFile.Close()

	// Decode the image file into an image.Image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Error decoding image file:", err)
		return err
	}

	// Create a new image context to draw onto
	dc := gg.NewContextForImage(img)

	// Load the font file and create a FontFace
	fontFile := "/assets/impact.ttf" // Replace with the path to your own TrueType font file
	fontFace, err := gg.LoadFontFace(fontFile, 48)
	if err != nil {
		log.Println("Error loading font file:", err)
		return err
	}
	dc.SetFontFace(fontFace)

	// Set the color of the text
	textColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	dc.SetColor(textColor)

	// Set the position of the text
	x, y := 400.0, 250.0
	xc, yc := 450.0, 100.0
	maxWidth := 200.0

	// Write the text onto the image
	dc.DrawStringWrapped(a, x, y, 2, 0.0, maxWidth, 1.2, gg.AlignCenter)
	dc.DrawStringWrapped(b, x, y, -1.5, 0.0, maxWidth, 1.2, gg.AlignCenter)
	dc.DrawStringWrapped(c, xc, yc, 0.7, 0.0, maxWidth, 1.2, gg.AlignCenter)

	// Save the edited image to a new file
	outputImgFile, err := os.Create("/assets/output.jpg")
	if err != nil {
		log.Println("Error creating output file:", err)
		return err
	}
	defer outputImgFile.Close()

	// Encode the edited image as JPEG and write it to the output file
	jpeg.Encode(outputImgFile, dc.Image(), nil)
	fmt.Println("Text added to image successfully!")
	return nil
}
