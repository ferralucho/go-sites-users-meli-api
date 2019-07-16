package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetCategoryFromApi(categoryID string) (*miapi.Category, *apierrors.ApiError) {
	category := &miapi.Category{
		Id: categoryID,
	}
	if err := category.Get(); err != nil {
		return nil, err
	}
	return category, nil
}
