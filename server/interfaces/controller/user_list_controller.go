package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type UserListController struct {
	logger     interfaces.Logger
	interactor usecase.UserListInteractor
}

func NewUserListController(logger interfaces.Logger, mongoHandler database.MongoHandler) *UserListController {
	userListController := UserListController{
		logger:     logger,
		interactor: *usecase.NewUserListInteractor(database.NewUserListRepository(mongoHandler)),
	}
	return &userListController
}

func (controller *UserListController) getUserList(c Context) {
	input := gateway.UserListInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(400, gateway.UserListOutput{false, "Parameter is invalid."})
		return
	}
	userList, err = controller.interactor.GetUserList(input.GetUserList())
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(500, gateway.UserListOutput{false, "Failed to get user list."})
		return
	}
	c.JSON(200, gateway.UserListOutput{true, "success!"})
}