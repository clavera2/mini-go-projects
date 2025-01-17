package filesystem

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

type Directory struct {
	Name string
	//a list of pointers to file structs respresenting files
	Files []*File
	//a list of pointers to directories representing sub-directories
	SubDir []*Directory
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

func (d *Directory) AddFile(file *File) error {
	if file == nil {
		return errors.New("file is nil")
	}
	if isPresent := slices.ContainsFunc(d.Files, func(f *File) bool {
		return f == file
	}); isPresent {
		return errors.New("file already exists in directory")
	}
	d.Files = append(d.Files, file)
	return nil
}

func (d *Directory) AddSubDir(dir *Directory) error {
	if dir == nil {
		return errors.New("subdirectory is nil")
	}
	if isPresent := slices.ContainsFunc(d.SubDir, func(dslice *Directory) bool {
		return dslice == dir
	}); isPresent {
		return errors.New("subdirectory is already in directory")
	}
	d.SubDir = append(d.SubDir, dir)
	return nil
}

func (d Directory) Persist(mode os.FileMode) {
	//this should be called recursively for all subdirectories in the outer directory
}
