package service

import (
	"context"
	"time"

	"github.com/tuki9ko/cola_aketa/database"
	"github.com/tuki9ko/cola_aketa/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ColaService struct{}

func (s ColaService) AddResult(userId int, cola_id int, result_date int64, note string) {
	formattedDate := time.Unix(result_date, 0)

	newRecord := models.ColaResult{
		UserID:     userId,
		ColaID:     cola_id,
		ResultDate: formattedDate,
		Note:       note,
	}

	newRecord.Insert(context.Background(), database.Db, boil.Infer())
}