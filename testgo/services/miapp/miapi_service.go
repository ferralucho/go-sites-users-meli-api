package miapp

import (
	"../../domains/miapp"
	"../../utils/apierrors"
)

func GetUserFromApi(userID int64) (*miapp.User, *apierrors.ApiError) {
	user := &miapp.User{
		ID:userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}