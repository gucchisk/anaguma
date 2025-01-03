package main

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gucchisk/anaguma/common"
)

func main() {
	v, err := common.GetVersionNumber()
	if err != nil {
		fmt.Println(err)
		return
	}
	dirname := fmt.Sprintf("anaguma-%s", v)
	filename := fmt.Sprintf("%s.%s.zip", dirname, runtime.GOOS)
	fmt.Printf("file: %s\n", filename)
	zipFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close()
	writer := zip.NewWriter(zipFile)
	defer writer.Close()
	files := []string{"bin/anaguma"}
	for _, file := range files {
		if err := addFileToZip(writer, file, dirname); err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("zip archived")
}

func addFileToZip(w *zip.Writer, path string, zipBaseDir string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	addFile := filepath.Join(zipBaseDir, path)
	header.Name = addFile
	header.Method = zip.Deflate
	writer, err := w.CreateHeader(header)
	dat, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = writer.Write(dat)
	return err
}
