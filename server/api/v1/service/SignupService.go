package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/database"
	"github.com/tuki9ko/cola_aketa/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type SignupService struct{}

func (s SignupService) Signup(c *gin.Context, name string, email string, password string, role_id int, biography string, user_icon string) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	newUser := models.User{
		Name:           name,
		Email:          email,
		HashedPassword: string(hashedPassword),
		RoleID:         role_id,
		Biography:      biography,
		UserIcon:       user_icon,
	}

	newUser.Insert(context.Background(), database.Db, boil.Infer())
}
