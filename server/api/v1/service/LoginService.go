package service

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/database"
	"github.com/tuki9ko/cola_aketa/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct{}

func (s LoginService) Login(c *gin.Context, userId string, password string) (*models.User, error) {
	user, err := models.Users(
		qm.Where("email = ?", userId),
	).One(c, database.Db)

	if err != nil {
		return nil, err
	}

	hashedPassword := "$2a$10$xxxxxxxxxxxxxxxxxxxxx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	if user != nil {
		hashedPassword = user.HashedPassword
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return nil, errors.New("Incorrect username or password")
	}

	session := sessions.Default(c)
	session.Set("userId", user.ID)
	session.Save()

	return user, nil
}

func (s LoginService) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
