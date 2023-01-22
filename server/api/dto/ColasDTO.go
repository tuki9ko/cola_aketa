package dto

type ColasDTO struct {
	ID               int    `boil:"id" json:"id"`
	UserName         string `boil:"user_name" json:"user_name"`
	ManufacturerName string `boil:"manufacturer_name" json:"manufacturer_name"`
	ColaName         string `boil:"cola_name" json:"cola_name"`
	Amount           int    `boil:"amount" json:"amount"`
	PackageName      string `boil:"package_name" json:"package_name"`
	Calories         int    `boil:"calories" json:"calories"`
}
