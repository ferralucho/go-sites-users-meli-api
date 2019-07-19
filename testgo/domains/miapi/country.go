package miapi

import (
	"../../utils/apierrors"
	"../../utils/constants"
	"../../utils/requests"
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Locale     string `json:"locale"`
	CurrencyId string `json:"currency_id"`
}

type Countries []struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Locale     string `json:"locale"`
	CurrencyId string `json:"currency_id"`
}

func (country *Country) Get() *apierrors.ApiError {
	if country.Id == "" {
		return &apierrors.ApiError{
			Message: "El country está vacío",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", constants.UrlCountries, country.Id)
	data, err := requests.Get(url)

	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country);
		err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil

}
