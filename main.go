package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// Helper function to serve an image file
func ServeImage(c *gin.Context, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		// Return a 404 if the file is not found
		c.Status(http.StatusNotFound)
		return
	}
	defer file.Close()

	// Determine the content type based on the file extension
	ext := filepath.Ext(filePath)
	var contentType string
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	default:
		// Return 415 if the file type is not supported
		c.Status(http.StatusUnsupportedMediaType)
		return
	}

	// Set the appropriate content type and serve the file
	c.Header("Content-Type", contentType)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		// Return 500 if there's an error while serving the file
		c.Status(http.StatusInternalServerError)
		return
	}
}

// Route handler for fetching images
func FetchCDN(c *gin.Context) {
	// Extract the filename from the URL parameter
	imageName := c.Param("filename")
	if imageName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image name"})
		return
	}

	// Clean up the image path, avoid malicious paths
	imageName = strings.Replace(imageName, "/cdn", "", -1)
	fmt.Printf("Image name: %s\n", imageName)

	// Construct the file path from the "images" directory
	filePath := filepath.Join("images", imageName)
	if _, err := os.Stat(filePath); err == nil {
		ServeImage(c, filePath)
		return
	}

	// Return a 404 if the file is not found
	c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Route to serve image files
	r.GET("/cdn/:filename", FetchCDN)

	// Run the server
	r.Run(":8282")
}
