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

func (interactor *UserListInteractor) GetUserList(userList domain.UserList) (domain.UserList, error) {
	userList, err := interactor.UserListRepository.GetUserList(userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
