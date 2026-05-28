package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type Index struct {
	TotalPages int      `json:"total_pages"`
	Pages      []string `json:"pages"`
}

type Photo struct {
	Filename    string `json:"filename"`
	Description string `json:"description"`
}

const (
	IndexPath  = "web/index/"
	PhotoPath  = "web/photos/"
	MaxPerPage = 15
)

func main() {
	fmt.Println("Starting gallery generation...")

	err := cleanIndexDirectory()
	if err != nil {
		log.Fatalf("Error clearing index: %v", err)
	}

	photos, err := loadPhotos()
	if err != nil {
		log.Fatalf("Error loading photos: %v", err)
	}

	photos, err = readDescriptions(photos)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("description.json not found, continuing without descriptions.")
		} else {
			log.Printf("Warning: could not read descriptions: %v. Continuing without them.", err)
		}
	}

	err = generateIndex(photos)
	if err != nil {
		log.Fatalf("Error generating index: %v", err)
	}

	err = generatePages(photos)
	if err != nil {
		log.Fatalf("Error generating pages: %v", err)
	}

	fmt.Println("Gallery generation complete!")
}

func generatePages(photos map[string]string) error {
	fmt.Println("Creating page files...")
	var filenames []string
	for filename := range photos {
		filenames = append(filenames, filename)
	}
	sort.Strings(filenames)

	var pages []Photo
	for _, filename := range filenames {
		pages = append(pages, Photo{
			Filename:    filename,
			Description: photos[filename],
		})
	}

	totalPages := (len(pages) + MaxPerPage - 1) / MaxPerPage
	fmt.Printf("Creating %d page files...", totalPages)

	for pageNum := 1; pageNum <= totalPages; pageNum++ {
		startIndex := (pageNum - 1) * MaxPerPage
		endIndex := min(startIndex+MaxPerPage, len(pages))

		pagePhotos := pages[startIndex:endIndex]

		jsonData, err := json.Marshal(pagePhotos)
		if err != nil {
			return err
		}

		pageFilename := filepath.Join(IndexPath, fmt.Sprintf("page%d.json", pageNum))
		err = os.WriteFile(pageFilename, jsonData, 0644)
		if err != nil {
			return err
		}
	}

	fmt.Println("Page files created successfully!")
	return nil
}

func generateIndex(photos map[string]string) error {
	fmt.Println("Creating index file...")
	totalImages := len(photos)
	totalPages := (totalImages + MaxPerPage - 1) / MaxPerPage

	var pages []string
	for i := 1; i <= totalPages; i++ {
		pages = append(pages, fmt.Sprintf("page%d.json", i))
	}

	index := Index{
		TotalPages: totalPages,
		Pages:      pages,
	}

	data, err := json.Marshal(index)
	if err != nil {
		return err
	}

	fmt.Printf("Index created with %d pages and %d images", totalPages, totalImages)
	return os.WriteFile(filepath.Join(IndexPath, "index.json"), data, 0644)
}

func readDescriptions(photos map[string]string) (map[string]string, error) {
	fmt.Println("Loading descriptions...")
	descriptionPath := filepath.Join(PhotoPath, "description.json")
	data, err := os.ReadFile(descriptionPath)
	if err != nil {
		return photos, err
	}

	var descriptions map[string]string
	err = json.Unmarshal(data, &descriptions)
	if err != nil {
		return photos, err
	}

	for filename, desc := range descriptions {
		if _, exists := photos[filename]; exists {
			photos[filename] = desc
		}
	}

	fmt.Printf("Applied descriptions to %d photos", len(descriptions))
	return photos, nil
}

func loadPhotos() (map[string]string, error) {
	fmt.Println("Scanning photos...")
	files, err := os.ReadDir(PhotoPath)
	if err != nil {
		return nil, err
	}

	photos := make(map[string]string)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			continue
		}
		photos[file.Name()] = ""
	}

	fmt.Printf("Found %d photos", len(photos))
	return photos, nil
}

func cleanIndexDirectory() error {
	fmt.Println("Cleaning index directory...")
	files, err := os.ReadDir(IndexPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			err := os.Remove(filepath.Join(IndexPath, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	fmt.Println("Index directory cleaned")
	return nil
}
