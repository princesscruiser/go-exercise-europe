package services

import (
	"encoding/json"
	"go-exercise-europe/clients"
	"go-exercise-europe/server/model"
	repository "go-exercise-europe/server/repositories"
)

type TranslateService interface {
	GetTranslation(text string, token string) (string, error)
}

type GTranslateService struct {
	GoogleClient clients.GoogleClient
}

func (gts *GTranslateService) GetTranslation(text string, token string) (string, error) {
	return "", nil
}

type QService interface {
	Find(userId string) ([]model.Question, error)
	Create(userId string, questions []model.Question) ([]model.Question, error)
}

type QServImpl struct {
	QRepo repository.QRepository
}

func (qsi *QServImpl) Find(userId string) ([]model.Question, error) {
	byteValue, _ := qsi.QRepo.Read()
	result := make([]model.Question, 0, 2)
	json.Unmarshal([]byte(byteValue), &result)
	return result, nil
}

func (qsi *QServImpl) Create(userId string, questions []model.Question) ([]model.Question, error) {
	return nil, nil
}
