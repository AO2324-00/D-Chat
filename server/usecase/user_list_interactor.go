package usecase

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type UserListInteractor struct {
	UserListRepository UserListRepository
}

func NewUserListInteractor(userListRepository UserListRepository) *UserListInteractor {
	userListInteractor := UserListInteractor{
		UserListRepository: userListRepository,
	}
	return &userListInteractor
}

func (interactor *UserListInteractor) GetUserList() ([]domain.UserList, error) {
	userList, err := interactor.UserListRepository.GetList()
	if err != nil {
		return nil, err
	}
	return userList, nil
}
