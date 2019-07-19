package miapi

import (
	"../../utils/apierrors"
	"../../utils/constants"
	"../../utils/requests"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CountryID string `json:"country_id"`
	SalesFreeMode string `json:"sales_free_mode"`
	MercadoPagoVersion string `json:"mercado_pago_version"`
	DefaultCurrencyId string `json:"default_currency_id"`
	InmediatePayment string `json:"inmediate_payment"`
	PaymentMethods []interface{} `json:"payment_methods"`
	Settings struct{
		IdentificationTypes []interface{} `json:"identification_types"`
		TaxpayerTypes []interface{}
	} `json:"settings"`
	Currencies []interface{} `json:"currencies"`
	Categories []interface{} `json:"categories"`


}

type Sites[] struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CountryID string `json:"country_id"`
	SalesFreeMode string `json:"sales_free_mode"`
	MercadoPagoVersion string `json:"mercado_pago_version"`
	DefaultCurrencyId string `json:"default_currency_id"`
	InmediatePayment string `json:"inmediate_payment"`
	PaymentMethods []interface{} `json:"payment_methods"`
	Settings struct{
		IdentificationTypes []interface{} `json:"identification_types"`
		TaxpayerTypes []interface{}
	} `json:"settings"`
	Currencies []interface{} `json:"currencies"`
	Categories []interface{} `json:"categories"`


}

func (site *Site) Get() *apierrors.ApiError {
	if site.ID == "" {
		return &apierrors.ApiError{
			Message: "El site está vacío",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", constants.MockUrlSites, site.ID)
	data, err := requests.Get(url)

	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &site);
		err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil

}

/*
func GetAllSites() *apierrors.ApiError {

	url := fmt.Sprintf("%s", urlSites)
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

	if err := json.Unmarshal(data, &site);
		err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil

}
*/

func GetAllSites(sites *Sites) *apierrors.ApiError {
	fmt.Println("About to get all sites from MELI...")
	res, err := http.Get(constants.UrlSites)
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

	if err := json.Unmarshal(data, &sites);
		err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}
