package usecase

import "github.com/kindai-csg/D-Chat/domain"

type UserListRepository interface {
	GetUserList(domain.UserList) (domain.UserList, error)
}
