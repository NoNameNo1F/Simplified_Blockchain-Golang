package utils

import (
	"fmt"
	"path/filepath"
)

func IntToHex(data int64) []byte {
	return []byte(fmt.Sprintf("%x", data))
}

func GenerateSpacing(size int) string {
	spacings := ""
	for i := 0; i < size; i++ {
		spacings += " "
	}
	return spacings
}

func GetPath(path string) string {
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}
	return path
}

func GetFileByExtens(folderName string, ext string) string {
	files, err := filepath.Glob(
		fmt.Sprintf("%s/*.%s", folderName, ext))
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}

	return files[0]
}
