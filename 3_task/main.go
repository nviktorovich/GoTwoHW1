package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const Path = "hw1/3_task/test/"

func FileCreate(name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err, file.Name())
		os.Exit(1)
	}
	fmt.Println(file.Name(), "was successfully created")
	defer FileClose(file)
}

func FileClose(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println(err, file.Name())
		os.Exit(1)
	}
	fmt.Println(file.Name(), "was successfully closed")
}

func main() {
	for i := 0; i < 1000000; i++ {
		fileNum := strconv.Itoa(i)
		fileName := Path + fileNum + ".txt"
		FileCreate(fileName)
	}
}
