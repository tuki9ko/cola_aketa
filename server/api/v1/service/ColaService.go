package service

import (
	"context"
	"fmt"
	"time"

	"github.com/tuki9ko/cola_aketa/api/v1/dto"
	"github.com/tuki9ko/cola_aketa/database"
	"github.com/tuki9ko/cola_aketa/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ColaService struct{}

func (s ColaService) GetCola(userId int, cola_id int) ([]dto.ColasDTO, error) {
	var dto []dto.ColasDTO

	models.ColaResults(
		qm.Select("cola_results.id as id, users.name as user_name, manufacturers.name as manufacturer_name, cola_types.name as cola_name, cola_types.amount as amount, packages.name as package_name, cola_types.calories as calories"),
		qm.InnerJoin("users on users.id = cola_results.user_id"),
		qm.InnerJoin("cola_types on cola_types.id = cola_results.cola_id"),
		qm.InnerJoin("manufacturers on manufacturers.id = cola_types.manufacturer_id"),
		qm.InnerJoin("packages on packages.id = cola_types.package_id"),
		qm.Where("cola_results.user_id = ?", userId),
		qm.Where("cola_results.id = ?", cola_id),
	).Bind(context.Background(), database.Db, &dto)

	fmt.Printf("userId: %d\n", userId)
	fmt.Printf("cola_id: %d\n", cola_id)

	return dto, nil
}

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

func (s ColaService) GetColas(userId int) ([]dto.ColasDTO, error) {
	var dto []dto.ColasDTO

	models.ColaResults(
		qm.Select("cola_results.id as id, users.name as user_name, manufacturers.name as manufacturer_name, cola_types.name as cola_name, cola_types.amount as amount, packages.name as package_name, cola_types.calories as calories"),
		qm.InnerJoin("users on users.id = cola_results.user_id"),
		qm.InnerJoin("cola_types on cola_types.id = cola_results.cola_id"),
		qm.InnerJoin("manufacturers on manufacturers.id = cola_types.manufacturer_id"),
		qm.InnerJoin("packages on packages.id = cola_types.package_id"),
		qm.Where("user_id = ?", userId),
		qm.OrderBy("id ASC"),
	).BindP(context.Background(), database.Db, &dto)

	return dto, nil
}
