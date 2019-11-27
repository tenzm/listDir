package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "Print storage info")
)

func hrSize(fsize int64) string {
	if fsize/1048576 > 0 {
		if fsize%1048576 > 0 {
			return strconv.Itoa(int(fsize/1048576)+1) + "MB"
		}
		return strconv.Itoa(int(fsize/1048576)) + "MB"
	}
	if fsize%1024 > 0 {
		return strconv.Itoa(int(fsize/1024)+1) + "KB"
	}
	return strconv.Itoa(int(fsize/1024)) + "KB"
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 06 15:4")
	fSize := strconv.Itoa(int(file.Size()))
	if *h == true {
		fSize = hrSize(file.Size())
	}
	fmt.Printf("%s %s %s \n", fSize, time, file.Name())
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)

	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
