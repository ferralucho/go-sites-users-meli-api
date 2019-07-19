package miapi

import (
	"../../utils/apierrors"
	"../../utils/constants"
	"../../utils/requests"
	"encoding/json"
	"fmt"
	"net/http"
)

type Category struct {
	Name 				string `json:"name"`
	Id 			string `json:"id"`
}



func (category *Category) Get() *apierrors.ApiError {

	if category.Id == "" {
		return &apierrors.ApiError{
			Message: "La category está vacío",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", constants.UrlCategories , category.Id)
	data, err := requests.Get(url)

	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &category);
		err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	return nil

}