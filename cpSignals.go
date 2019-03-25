package main

import (
	"fmt"
	"io"
	"os"
)

var BUFFERSIZE int64
var FILESIZE int64
var BYTESWEITTEN int64

func Copy(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	FILESIZE = sourceFileStat.Size()

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err != nil {
		return fmt.Errorf("File %s already exists", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[ :n]); err != nil {
			return err
		}
		BYTESWEITTEN = BYTESWEITTEN + int64(n)
	}
	return err
}

func progressInfo() {
	progress := float64(BYTESWEITTEN) / float64(FILESIZE) * 100
	fmt.Printf("Progress: %.2f%%\n", progress)
}

func main() {

}
