package repository

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
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
	jsonFile, err := os.Open("../resources/questions.json")
	if err != nil {
		fmt.Println(err)
	}
	log.Info().Msg("Successfully Opened questions.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue, nil
}
