package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/clavera2/mini-go-projects/memFS/filesystem"
)

func main() {
	fmt.Println("WELCOME TO MEMFS: IN MEMORY FILESYSTEM")
	fmt.Println("choose an option below:\n1) Create file\n2) Create Directory\n3) List Directory contents\n4) Delete file\n5)Exit")
	for {
		fmt.Print(":")
		reader := bufio.NewReader(os.Stdin)
		c, err := reader.ReadString('\n')
		//fmt.Println(c)
		if err != nil {
			log.Fatal("cannot read string")
		}
		choice, err := strconv.Atoi(strings.Trim(c, "\n"))
		if err != nil {
			log.Fatal("cannot convert to integer")
		}
		switch choice {
		case 1:
			fmt.Print("Enter filename: ")
			filename, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("an error occured")
			}
			filesystem.CreateFile(filename)
			fmt.Println("file successfully created")
			fmt.Print(":")
		case 2:
			fmt.Print("enter directory name: ")
			dirname, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("an error occured")
			}
			filesystem.CreateDir(dirname)
			fmt.Println("directory successfully created")
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
		default:
			fmt.Println("must choose an option from 1-5")
		}
	}
}
