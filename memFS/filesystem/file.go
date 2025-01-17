package filesystem

import (
	"errors"
	"os"
)

type File struct {
	Name       string
	Content    []byte
	Size       int
	Permission string
}

func CreateFile(name string) *File {
	return &File{Name: name, Content: []byte{}, Size: 0}
}

func (f *File) Write(content []byte) error {
	//this overwrites the contents of the file
	//to append, use the append function
	if f.Permission != "WRITE" {
		return errors.New("cannot write to file: permission denied")
	}
	f.Content = content
	f.Size = len(f.Content)
	return nil
}

func (f *File) Append(content []byte) error {
	//this adds data to the end of the byte slice of the file, appending to it
	if f.Permission != "WRITE" {
		return errors.New("cannot append to file: permission denied")
	}
	f.Content = append(f.Content, content...)
	f.Size = len(f.Content)
	return nil
}

func (f File) ReadFile() ([]byte, error) {
	if f.Permission != "READ" {
		return nil, errors.New("cannot read from file: permission denied")
	}
	return f.Content, nil
}

func (f *File) DeleteAllContent() error {
	if f.Permission != "WRITE" {
		return errors.New("cannot delete contents: permission denied")
	}
	f.Content = []byte{}
	return nil
}

func (f File) Persist(dirPath string) error {
	//This allows the user to save the file from memory to disk
	if err := os.Chdir(dirPath); err != nil {
		os.Mkdir(dirPath, 0777)
		err = os.Chdir(dirPath)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(f.Name)
	if err != nil {
		return err
	}
	_, err = file.Write(f.Content)
	if err != nil {
		return err
	}
	return nil
}
