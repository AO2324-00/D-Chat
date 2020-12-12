package database

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/kindai-csg/D-Chat/domain"
	"github.com/kindai-csg/D-Chat/domain/enum"
)

type UserListRepository struct {
	mongoHandler   MongoHandler
	collectionName string
}

func NewUserListRepository(mongoHandler MongoHandler) *UserListRepository {
	userListRepository := UserListRepository{
		mongoHandler:   mongoHandler,
		collectionName: "Users",
	}
	//userListRepository.createIndex()
	return &userListRepository
}
/*
func (repository *UserListRepository) createIndex() {
	repository.mongoHandler.CreateIndex(repository.collectionName, []KV{{"user_id", 1}}, []KV{{"unique", true}})
	repository.mongoHandler.CreateIndex(repository.collectionName, []KV{{"mail", 1}}, []KV{{"unique", true}})
}
*/
func (repository *UserListRepository) GetUserList(userList domain.UserList) (domain.UserList, error) {
	query := []KV{}
	for i := range userList.MembersId {
		query = append(query, {"user_id", i})
	}
	raw, err := repository.mongoHandler.Find(repository.collectionName, query)
	if err != nil {
		return nil, err
	}
	users := []domain.User{}
	for _, kv := range raw {
		for _, u := range kv {
			user := domain.User{}
			switch u.Key {
			case "_id":
				
				user.Id = u.Value.(string)
			case "user_id":
				user.UserId = u.Value.(string)
			case "password":
				user.Password = u.Value.(string)
			case "mail":
				user.Mail = u.Value.(string)
			case "bio":
				user.Bio = u.Value.(string)
			case "status":
				user.Status = u.Value.(string)
			case "status_text":
				user.StatusText = u.Value.(string)
			case "auth":
				user.Auth = u.Value.(domain.AuthType)
			}
			users = append(users, user)
		}
	}
	return users, nil
}
