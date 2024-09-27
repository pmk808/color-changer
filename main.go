package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Open the original image
	file, err := os.Open("jersey.png") // Change to your image file
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Create a new image with the same dimensions
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// Define the brown color
	brown := color.RGBA{139, 69, 19, 255} // RGB for brown

	// Define the jersey color range (example values)
	const (
		jerseyRMin = 0
		jerseyRMax = 255
		jerseyGMin = 0
		jerseyGMax = 100
		jerseyBMin = 0
		jerseyBMax = 100
	)

	// Change the color of each pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the original pixel color
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()

			// Check if the pixel color is within the jersey color range
			if int(r) >= jerseyRMin && int(r) <= jerseyRMax &&
				int(g) >= jerseyGMin && int(g) <= jerseyGMax &&
				int(b) >= jerseyBMin && int(b) <= jerseyBMax {
				newImg.Set(x, y, brown) // Change to brown
			} else {
				newImg.Set(x, y, originalColor) // Keep the original color
			}
		}
	}

	// Save the new image
	outFile, err := os.Create("jersey_brown.png") // Output file
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Encode the new image to PNG
	err = png.Encode(outFile, newImg)
	if err != nil {
		panic(err)
	}

	println("Image color changed to brown and saved as jersey_brown.png")
}
