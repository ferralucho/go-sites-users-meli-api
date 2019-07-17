package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
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

func GetUserAsyncFromApi(userID int64, wg *sync.WaitGroup) (*miapi.User, *apierrors.ApiError) {
	user := &miapi.User{
		ID:userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	wg.Done()
	return user, nil
}
