package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	images, err := listImages("data")
	if err != nil {
		log.Fatalln(err)
	}
	for _, image := range images {
		log.Println(image)
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
