package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetUserFromApi(userID int64) (*miapi.User, *apierrors.ApiError) {
	user := &miapi.User{
		ID:userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserAsyncFromApi(userID int64, values chan *miapi.Result, errors chan *apierrors.ApiError) {
	user := &miapi.User{
		ID:userID,
	}
	if err := user.Get();
	err != nil {
		errors <- err
	} else{
		errors <- nil
	}

	result := &miapi.Result{
		User: user,
	}

	values <- result
}
