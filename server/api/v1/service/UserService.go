package service

import (
	"context"

	"github.com/tuki9ko/cola_aketa/database"
	"github.com/tuki9ko/cola_aketa/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
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

func (s UserService) UpdateUser(id int, name string, email string, password string, role int, biography string, user_icon string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		ID:             id,
		Name:           name,
		Email:          email,
		HashedPassword: string(hashedPassword),
		RoleID:         role,
		Biography:      biography,
		UserIcon:       user_icon,
	}

	user.Update(context.Background(), database.Db, boil.Infer())
}
