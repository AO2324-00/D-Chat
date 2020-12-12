package gateway

type UserListInput struct {
	MembersId  string   `json:"members_id"`
}

func (userList *UserListInput) GetUserList() domain.UserList {
	return domain.UserList{
		MembersId:   userList.MembersId,
	}
}