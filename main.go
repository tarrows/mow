package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

func main() {
	images, err := listImages("data")
	if err != nil {
		log.Fatalln(err)
	}
	for _, image := range images {
		log.Println(image)
		err = processImage(image, "data", "dest")
		if err != nil {
			log.Fatalln(err)
		}
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
	srcFile, err := os.Open(filepath.Join(srcDir, name))
	defer srcFile.Close()

	if err != nil {
		return err
	}

	img, imageFormat, err := image.Decode(srcFile)
	if err != nil {
		return err
	}

	rect := img.Bounds()
	log.Println("Format:", imageFormat)
	log.Printf("(width, height) = (%v, %v)", rect.Dx(), rect.Dy())

	target := image.NewRGBA(image.Rect(0, 0, rect.Dx()/2, rect.Dy()/2))

	draw.CatmullRom.Scale(target, target.Bounds(), img, rect, draw.Over, nil)

	dstFile, err := os.Create(filepath.Join(dstDir, name))
	defer dstFile.Close()

	if err != nil {
		return err
	}

	switch imageFormat {
	case "jpeg":
		if err := jpeg.Encode(dstFile, target, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(dstFile, target); err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("Format not supported: %v", imageFormat))
	}

	log.Println("image converted:", name)

	return nil
}
