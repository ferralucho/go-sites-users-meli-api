package miapi

import (
	"../../utils/apierrors"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

const urlCountries = "https://api.mercadolibre.com/countries/"

func (country *Country) Get() *apierrors.ApiError {
	if country.Id == "" {
		return &apierrors.ApiError{
			Message: "El country está vacío",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", urlCountries, country.Id)
	res, err := http.Get(url)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
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
