package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func FilesGET(c *gin.Context) {
	// Validate method.
	switch c.Request.Method {
	case http.MethodGet:
	case http.MethodHead:
	default:
		c.Status(http.StatusMethodNotAllowed)
		return
	}

	// Parse path.
	path := c.Request.URL.Path
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	path = strings.ReplaceAll(path, "..", "")
	if path == "" {
		path = "/index.html"
	}

	fmt.Println(path)
	// Read file.
	data, err := os.ReadFile("html" + path)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	// Set content type.
	fileParts := strings.Split(path, ".")
	switch fileParts[len(fileParts)-1] {
	case "html":
		c.Header("Content-Type", "text/html; charset=utf-8")
	case "css":
		c.Header("Content-Type", "text/css; charset=utf-8")
	case "js":
		c.Header("Content-Type", "application/javascript; charset=utf-8")
	case "json":
		c.Header("Content-Type", "application/json; charset=utf-8")
	case "png":
		c.Header("Content-Type", "image/png")
	case "jpg", "jpeg":
		c.Header("Content-Type", "image/jpeg")
	case "gif":
		c.Header("Content-Type", "image/gif")
	case "svg":
		c.Header("Content-Type", "image/svg+xml")
	case "ico":
		c.Header("Content-Type", "image/x-icon")
	case "woff":
		c.Header("Content-Type", "font/woff")
	case "woff2":
		c.Header("Content-Type", "font/woff2")
	case "ttf":
		c.Header("Content-Type", "font/ttf")
	case "otf":
		c.Header("Content-Type", "font/otf")
	case "eot":
		c.Header("Content-Type", "font/eot")
	case "map":
		c.Header("Content-Type", "application/json")
	case "mp3":
		c.Header("Content-Type", "audio/mpeg")
	default:
		c.Header("Content-Type", "application/octet-stream")
	}

	// Set content length.
	c.Header("Content-Length", strconv.Itoa(len(data)))

	// Write data.
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(data)
}
