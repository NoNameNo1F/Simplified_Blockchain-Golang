package utils

import (
	"fmt"
	"path/filepath"
)

// IntToHex converts an integer to its hexadecimal representation.
//
// Parameters:
// - data: the integer to be converted.
//
// Returns:
// - []byte: the hexadecimal representation of the integer.
func IntToHex(data int64) []byte {
	return []byte(fmt.Sprintf("%x", data))
}

// GenerateSpacing generates a string of spaces with the specified size.
//
// Parameters:
// - size: the number of spaces to generate.
//
// Returns:
// - string: a string of spaces with the specified size.
func GenerateSpacing(size int) string {
	spacings := ""
	for i := 0; i < size; i++ {
		spacings += " "
	}
	return spacings
}

// GetPath returns the absolute path of the given path.
//
// Parameters:
// - path: the path to be converted to an absolute path.
//
// Returns:
// - string: the absolute path of the given path. If an error occurs, an empty string is returned.
func GetPath(path string) string {
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}
	return path
}

// GetFileByExtens retrieves the file with the specified extension in the given folder.
//
// Parameters:
// - folderName: the name of the folder to search for files.
// - ext: the extension of the file to be retrieved.
// Returns a string representing the file path.
func GetFileByExtens(folderName string, ext string) string {
	files, err := filepath.Glob(
		fmt.Sprintf("%s/*.%s", folderName, ext))
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}

	return files[0]
}

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
