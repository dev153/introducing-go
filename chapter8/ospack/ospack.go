package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fileName := "test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File", fileName, "not found.")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Could not obtain file", file, "statistics.")
	} else {
		fmt.Println("File", stat.Name(), "size is", stat.Size())
	}

	// reading the whole file in a slice of bytes
	bs := make([]byte, stat.Size())
	// err here is not considered to be a new variable.
	// because the underscore (don't care variable) is used
	// thus the previous err variable is consideted to be valid.
	_, err = file.Read(bs)

	if err != nil {
		fmt.Println("Could not read the content of file", stat.Name())
	} else {
		fileContents := string(bs)
		fmt.Println(fileContents)
	}

	// Now let's write some content to a file.
	fh, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Could not create file", "\"output.txt\"")
		return
	}
	defer fh.Close()
	written, err := fh.WriteString("This is a line.")
	if err != nil {
		fmt.Println("Could not write content to file.")
		return
	}

	fmt.Println("Number of bytes written:", written)

	// Let's read the contents of a directory
	dirfh, err := os.Open(".")
	if err != nil {
		fmt.Println("Could not open directory . for reading it's contents")
	}
	defer dirfh.Close()

	fileInfos, err := dirfh.Readdir(-1)
	if err != nil {
		fmt.Println("Could not read the contents of the directory", dirfh.Name())
	}

	for _, fileInfo := range fileInfos {
		fmt.Println("File Name:", fileInfo.Name(),
			"Size:", fileInfo.Size(), "Is directory:", fileInfo.IsDir())
	}

	// Let's show the contents of a directory recursively.
	walkFuncToCall := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("Path:", path)
		fmt.Println("Is directory:", fileInfo.IsDir())
		if fileInfo.IsDir() == false {
			fmt.Println("Size:", fileInfo.Size())
		}
		return nil
	}
	filepath.Walk("/home/nthomos/go", walkFuncToCall)
}
