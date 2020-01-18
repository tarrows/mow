package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	images, err := listImages("data")
	if err != nil {
		log.Fatalln(err)
	}
	for _, image := range images {
		log.Println(image)
		processImage(image, "data", "dest")
	}
}

func listImages(root string) ([]string, error) {
	acceptableExts := map[string]struct{}{
		".png":  struct{}{},
		".jpg":  struct{}{},
		".jpeg": struct{}{},
	}

	entries, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var images []string

	for _, entry := range entries {
		ext := filepath.Ext(entry.Name())

		if _, ok := acceptableExts[ext]; !entry.IsDir() && ok {
			images = append(images, entry.Name())
		}
	}

	return images, nil
}

func processImage(name, srcDir, dstDir string) error {
	file, err := os.Open(filepath.Join(srcDir, name))
	defer file.Close()

	if err != nil {
		return err
	}

	img, imageFormat, err := image.Decode(file)
	if err != nil {
		return err
	}

	rect := img.Bounds()
	log.Println("Format:", imageFormat)
	log.Printf("(width, height) = (%v, %v)", rect.Dx(), rect.Dy())

	return nil
}
