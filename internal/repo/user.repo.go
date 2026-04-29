package repo

import (
	"errors"
	"strconv"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser(id string) (string, error) {
	// id không phải số thì trả về lỗi
	_, err := strconv.Atoi(id)
	if err != nil {
		return "", errors.New("id must be a number")
	}

	return "Hai Dang Phuong, " + id, nil
}
