package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

//Storage is the main entry point into our .boom file.
type Storage struct {
	path  string
	Lists []map[string][]map[string]string `json:"lists"`
}

//FindList looks for the list specified and returns it if found.
func (s *Storage) FindList(name string) []map[string]string {
	for _, m := range s.Lists {
		for key, value := range m {
			if key == name {
				return value
			}
		}
	}
	return nil
}

//FindValueByEntry looks to see if there is any entry
//with the specified name and returns it if so.
func (s *Storage) FindValueByEntry(name string) string {
	for _, list := range s.Lists {
		for _, entries := range list {
			for _, entry := range entries {
				for key, value := range entry {
					if key == name {
						return value
					}
				}
			}
		}
	}
	return ""
}

//Save persists changes back to the .boom file.
func (s *Storage) Save() {

	data, error := json.Marshal(s)

	if error != nil {
		panic(error)
	}

	ioutil.WriteFile(s.path, data, 0644)
}

//NewDefaultListStorage creates a Storage object
//that provides access to all stored entries in the
//default ~/.boom location.
func NewDefaultListStorage() Storage {

	user, error := user.Current()

	if error != nil {
		panic(error)
	}

	defaultBoomPath := user.HomeDir + "/.boom"

	return NewCustomListStorage(defaultBoomPath)
}

//NewCustomListStorage creates a Storage object
//that provides access to all stored entries in
//the specified location.
func NewCustomListStorage(path string) Storage {

	path, _ = filepath.Abs(path)
	file, error := ioutil.ReadFile(path)

	if error != nil {
		panic(error)
	}

	storage := new(Storage)

	json.Unmarshal(file, &storage)

	storage.path = path

	return *storage

}
