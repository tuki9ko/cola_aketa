package service

import (
	"context"

	"github.com/tuki9ko/cola_aketa/database"
	"github.com/tuki9ko/cola_aketa/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserService struct{}

func (s UserService) GetUserById(userId int) (*models.User, error) {
	user, err := models.Users(
		qm.Where("id = ?", userId),
	).One(context.Background(), database.Db)

	if err != nil {
		return nil, err
	}

	return user, nil
}
