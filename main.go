package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sort"
)

var (
	d     = flag.String("d", ".", "Directory to process")
	a     = flag.Bool("a", false, "Print all info")
	h     = flag.Bool("h", false, "Print storage info")
	msort = flag.Bool("sort", false, "Sort for modification date")
)

/*func sort(files []os.FileInfo) {
	for j := 0; j < len(files); j++ {
		for i := 0; i < len(files)-j-1; i++ {
			if files[i].ModTime().Before(files[i+1].ModTime()) {
				files[i], files[i+1] = files[i+1], files[i]
			}
		}
	}
}*/

type FilesTimeSort struct {
	files []os.FileInfo
}

func (fts *FilesTimeSort) Len() int{
	return len(fts.files)
}

func (fts *FilesTimeSort) Less(i, j int) bool{
	return fts.files[i].ModTime().Before(fts.files[j].ModTime())
}

func (fts *FilesTimeSort) Swap(i, j int){
	fts.files[i], fts.files[j] = fts.files[j], fts.files[i]
}

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
	sortedFiles :=FilesTimeSort{files}
	if *msort == true {
		sort.Sort(&sortedFiles)
	}
	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
