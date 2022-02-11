package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func saveFileHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("upload/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "CI Loader"})

	files, err := ioutil.ReadDir("upload")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		filelist := file.Name()
		c.JSON(http.StatusOK, gin.H{"filelist": filelist})
	}
}

func status(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func build(c *gin.Context) {
	cmd := exec.Command("docker", "build", ".")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed: %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if len(errStr) > 1 {
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	}
	fmt.Printf(outStr)
	c.JSON(200, gin.H{
		"build": "OK",
	})
}

func routes() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", home)
	r.GET("/ping", status)
	r.POST("/upload", saveFileHandler)
	r.GET("/run", build)
	r.StaticFS("/file", http.Dir("upload"))
	r.Run("localhost:8080")
}

func main() {
	routes()
}
