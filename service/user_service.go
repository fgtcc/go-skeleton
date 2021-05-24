package service

import (
	"go-skeleton/middleware"
	"go-skeleton/model"
	"go-skeleton/utils"
	"go-skeleton/utils/errors"
	"go-skeleton/utils/typedef"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(loginName, password string) (string, error)
	AddUser(username, loginName, password string) error
	GetUserInfo(id int64, loginName string) (*typedef.UserInfo, error)
	GetUserList(status, pageNo, pageSize int64) ([]*typedef.UserInfo, int64, error)
	GetWholeUserList() ([]*typedef.UserInfo, int64, error)
	UpdateUser(id int64, username, password string) error
	DeleteUsers(idList []int64) error
	BatchUpdateUserStatus(idList []int64, status int64) error
}

func (svc *service) Login(loginName, password string) (string, error) {
	encodePwd, err := model.User.GetEncodePwd(loginName)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(password))
	if err != nil {
		return "", errors.New(errors.ERROR_PASSWORD_WRONG, errors.GetErrMsg(errors.ERROR_PASSWORD_WRONG))
	}

	token, err := middleware.SetToken(loginName)
	return token, err
}

func (svc *service) AddUser(username, loginName, password string) error {
	isExist, err := model.User.IsExist(0, loginName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(errors.ERROR_USER_EXIST, errors.GetErrMsg(errors.ERROR_USER_EXIST))
	}

	// encrypt password
	originPwd := []byte(password)
	hashPwd, _ := bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost)
	encodePwd := string(hashPwd)

	err = model.User.Create(username, loginName, encodePwd)
	return err
}

func (svc *service) GetUserInfo(id int64, loginName string) (*typedef.UserInfo, error) {
	userInfo, err := model.User.GetInfo(id, loginName)
	return userInfo, err
}

func (svc *service) GetUserList(status, pageNo, pageSize int64) ([]*typedef.UserInfo, int64, error) {
	userList, count, err := model.User.GetList(status, pageNo, pageSize)
	return userList, count, err
}

func (svc *service) GetWholeUserList() ([]*typedef.UserInfo, int64, error) {
	userList, count, err := model.User.GetWholeList()
	return userList, count, err
}

func (svc *service) UpdateUser(id int64, username, password string) error {
	isExist, err := model.User.IsExist(id, "")
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(errors.ERROR_USER_NOT_EXIST, errors.GetErrMsg(errors.ERROR_USER_NOT_EXIST))
	}

	// encrypt password
	encodePwd := ""
	if password != "" {
		originPwd := []byte(password)
		hashPwd, _ := bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost)
		encodePwd = string(hashPwd)
	}

	err = model.User.Update(id, username, encodePwd)
	return err
}

func (svc *service) DeleteUsers(idList []int64) error {
	// adminInfo, err := model.User.GetAdminInfo()
	// if err != nil {
	// 	return err
	// }
	isContain := utils.IsContainInt64(idList, typedef.ADMIN_USER_ID)
	if isContain {
		return errors.New(errors.ERROR_FORBIDDEN_EDIT_ADMIN, errors.GetErrMsg(errors.ERROR_FORBIDDEN_EDIT_ADMIN))
	}

	err := model.User.Delete(idList)
	return err
}

func (svc *service) BatchUpdateUserStatus(idList []int64, status int64) error {
	isContain := utils.IsContainInt64(idList, typedef.ADMIN_USER_ID)
	if isContain {
		return errors.New(errors.ERROR_FORBIDDEN_EDIT_ADMIN, errors.GetErrMsg(errors.ERROR_FORBIDDEN_EDIT_ADMIN))
	}

	err := model.User.BatchUpdateStatus(idList, status)
	return err
}
