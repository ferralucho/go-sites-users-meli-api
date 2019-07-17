package miapi

import (
	"../../utils/apierrors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Sites []struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

const urlSites = "https://api.mercadolibre.com/sites/"

func (site *Site) Get() *apierrors.ApiError {
	if site.Id == "" {
		return &apierrors.ApiError{
			Message: "El site está vacío",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", urlSites, site.Id)
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
	res, err := http.Get(urlSites)
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
