package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	fileHandle, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fileHandle.Close()

	fileHash := crc32.NewIEEE()
	_, err = io.Copy(fileHash, fileHandle)

	if err != nil {
		return 0, err
	}
	return fileHash.Sum32(), nil
}

func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)

	// shahash := sha1.New()
	// shahash.Write([]byte("test"))
	// bs := shahash.Sum([]byte{})
	// fmt.Println(bs)
}
