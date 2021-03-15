package repository

import (
	"fmt"
	"io/ioutil"
	"os"
)

type QRepository interface {
	Write() error
	Read() ([]byte, error)
}

type JsonRepo struct {
	Path string
}

func (jw *JsonRepo) Write() error {
	return nil
}

func (jw *JsonRepo) Read() ([]byte, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("../../resources/questions.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue, nil
}
