package forms

import (
	"github.com/msajad79/pricing_engine/config"
	"github.com/msajad79/pricing_engine/models"
)

type RouteForm struct {
	Origin      string `json:"origin" validate:"is-city"`
	Destination string `json:"destination"`
}

type RuleForm struct {
	Routes      []RouteForm `json:"routes"   binding:"required"`
	Airelines   []string    `json:"airlines" binding:"required"`
	Agencies    []string    `json:"agencies" binding:"required"`
	Suppliers   []string    `json:"suppliers" binding:"required"`
	AmountType  string      `json:"amountType" binding:"required"`
	AmountValue float64     `json:"amountValue" binding:"required"`
}

func CheckValidateSliceRouteForm(routes []RouteForm) error {
	for _, route := range routes {
		err := config.Validate.Struct(route)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckValidateSliceCitiesForm(slice_forms []models.City) error {

	for _, form := range slice_forms {
		err := config.Validate.Struct(form)
		if err != nil {
			return err
		}
	}
	return nil
}
func CheckValidateSliceAgenciesForm(slice_forms []models.Agency) error {
	for _, form := range slice_forms {
		err := config.Validate.Struct(form)
		if err != nil {
			return err
		}
	}
	return nil
}
func CheckValidateSliceAirlinesForm(slice_forms []models.Airline) error {
	for _, form := range slice_forms {
		err := config.Validate.Struct(form)
		if err != nil {
			return err
		}
	}
	return nil
}
func CheckValidateSliceSuppliersForm(slice_forms []models.Supplier) error {
	for _, form := range slice_forms {
		err := config.Validate.Struct(form)
		if err != nil {
			return err
		}
	}
	return nil
}
