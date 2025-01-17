package filesystem

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

type Directory struct {
	Name   string
	Files  []*File      //im not sure if this field should be of type []File or []*File
	SubDir []*Directory //im not sure if this field should be of type []Directory or []*Directory
}

func CreateDir(name string) *Directory {
	return &Directory{Name: name}
}

func (d Directory) GetFiles() []*File {
	//returns the list of files in the directory
	return d.Files
}

func (d Directory) GetSubDirs() []*Directory {
	//returns the list of subdirectories
	return d.SubDir
}

func (d *Directory) DeleteFile(name string) error {
	num_of_files := len(d.Files)
	slices.DeleteFunc(d.Files, func(f *File) bool {
		return f.Name == name
	})
	if num_of_files == len(d.Files) {
		//no file had the same name to be deleted
		return errors.New(fmt.Sprintf("No such file in this directory with the name: %s", name))
	}
	return nil
}

func (d *Directory) DeleteSubDir(name string) error {
	num_of_dirs := len(d.SubDir)
	slices.DeleteFunc(d.SubDir, func(d *Directory) bool {
		return d.Name == name
	})
	if len(d.SubDir) == num_of_dirs {
		return errors.New(fmt.Sprintf("No such subdirectory in this directory with the name: %s", name))
	}
	return nil
}

func (d Directory) ListAll() {
	fmt.Println("================FILES============")
	for _, file := range d.Files {
		fmt.Println(file.Name)
	}
	fmt.Println("================DIR==============")
	for _, sd := range d.SubDir {
		fmt.Println(sd.Name)
	}
}

func (d Directory) Persist(mode os.FileMode) {
	//this should be called recursively for all subdirectories in the outer directory
}
