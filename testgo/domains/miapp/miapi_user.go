package miapp

import (
	"../../utils/apierrors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID               int64  `json:"id"`
	NickName         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	Email            string `json:"email"`
}

const urlUsers = "https://api.mercadolibre.com/users/"

func (user *User) Get() *apierrors.ApiError {
	if user.ID == 0 {
		return &apierrors.ApiError{
			Message: "El usuario está vacío",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%d", urlUsers, user.ID)
	res, err := http.Get(url)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &user);
	err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	return nil

}
