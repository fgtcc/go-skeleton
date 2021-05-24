package model

import (
	"fmt"
	"go-skeleton/utils/errors"
	"go-skeleton/utils/log"
	"go-skeleton/utils/typedef"
)

var User = UserInfo{}

const userTableName = "user"

type UserInfo struct {
	Id        int64 `gorm:"primary_key"`
	Username  string
	LoginName string
	Password  string
	Status    int64
}

func (this *UserInfo) IsExist(id int64, loginName string) (bool, error) {
	var user UserInfo
	condition := assembleUserCondition(id, loginName)

	err := db.Table(userTableName).Model(this).Where(condition).Find(&user)
	if err.Error != nil {
		if err.RecordNotFound() {
			log.Debugf("user record not found, id: %v, loginName: %v", id, loginName)
			return false, nil
		}
		log.Errorf("failed to find user info, id: %v, loginName: %v, err: %v", id, loginName, err)
		return false, err.Error
	}

	return true, err.Error
}

func (this *UserInfo) Create(username, loginName, password string) error {
	user := &UserInfo{
		Username:  username,
		LoginName: loginName,
		Password:  password,
	}

	err := db.Table(userTableName).Model(this).Create(user).Error
	if err != nil {
		log.Errorf("failed to create user, username: %v, loginName: %v, err: %v", username, loginName, err)
		return err
	}
	return nil
}

func (this *UserInfo) GetInfo(id int64, loginName string) (*typedef.UserInfo, error) {
	log.Debugf("get user info, id: %v, loginName: %v", id, loginName)
	var user typedef.UserInfo
	condition := assembleUserCondition(id, loginName)

	err := db.Table(userTableName).Model(this).
		Select("id, login_name, username, status").
		Where(condition).Scan(&user)
	if err.RecordNotFound() {
		log.Debugf("user is not exist, id: %v, loginName: %v", id, loginName)
		return &user, errors.New(errors.ERROR_USER_NOT_EXIST, errors.GetErrMsg(errors.ERROR_USER_NOT_EXIST))
	}
	if err.Error != nil {
		log.Errorf("failed to get user info, id: %v, loginName: %v, err: %v", id, loginName, err)
		return &user, errors.New(errors.ERROR, err.Error.Error())
	}
	return &user, nil
}

func (this *UserInfo) GetList(status, pageNo, pageSize int64) ([]*typedef.UserInfo, int64, error) {
	var userList []*typedef.UserInfo
	var count int64
	limit := pageSize
	offset := (pageNo - 1) * pageSize

	whereCondition := ""
	if status != -1 {
		whereCondition = fmt.Sprintf("user.status = %d", status)
	}

	err := db.Table(userTableName).Model(this).
		Select("id, login_name, username, status").
		Where(whereCondition).
		Count(&count).
		Offset(offset).
		Limit(limit).
		Scan(&userList).Error

	if err != nil {
		log.Errorf("failed to get user list, pageNo: %v, pageSize: %v, err: %v", pageNo, pageSize, err)
		return userList, 0, err
	}
	return userList, count, nil
}

func (this *UserInfo) GetWholeList() ([]*typedef.UserInfo, int64, error) {
	var userList []*typedef.UserInfo
	var count int64

	err := db.Table(userTableName).Model(this).
		Select("id, login_name, username, status").
		Count(&count).
		Scan(&userList).Error

	if err != nil {
		log.Errorf("failed to get whole user list, err: %v", err)
		return userList, 0, err
	}
	return userList, count, nil
}

func (this *UserInfo) Update(id int64, username, password string) error {
	var maps = make(map[string]interface{})
	if username != "" {
		maps["username"] = username
	}
	if password != "" {
		maps["password"] = password
	}

	err := db.Table(userTableName).Model(this).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		log.Errorf("failed to update user, id: %v, username: %v, err: %v", id, username, err)
		return err
	}
	return err
}

func (this *UserInfo) Delete(idList []int64) error {
	var user UserInfo
	err := db.Table(userTableName).Model(this).Where("id IN (?)", idList).Delete(&user).Error
	if err != nil {
		log.Errorf("failed to delete users, idList: %v, err: %v", idList, err)
		return err
	}
	return err
}

func (this *UserInfo) GetEncodePwd(loginName string) (string, error) {
	var user UserInfo
	err := db.Table(userTableName).Model(this).Where("login_name = ?", loginName).First(&user)
	if err.RecordNotFound() {
		return "", errors.New(errors.ERROR_USER_NOT_EXIST, errors.GetErrMsg(errors.ERROR_USER_NOT_EXIST))
	}
	if err.Error != nil {
		log.Errorf("failed to get user encode password, loginName: %v, err: %v", loginName, err)
		return "", err.Error
	}

	return user.Password, nil
}

func (this *UserInfo) GetAdminInfo() (*typedef.UserInfo, error) {
	var user typedef.UserInfo

	err := db.Table(userTableName).Model(this).
		Select("id, login_name, username, status").
		Where("login_name = ?", "admin").Scan(&user).Error
	if err != nil {
		log.Errorf("failed to get admin info, err: %v", err)
		return &user, err
	}
	return &user, err
}

func (this *UserInfo) BatchUpdateStatus(idList []int64, status int64) error {
	var maps = make(map[string]interface{})
	maps["status"] = status

	err := db.Table(userTableName).Model(this).Where("id IN (?)", idList).Updates(maps).Error
	if err != nil {
		log.Errorf("failed to batch update user status, idList: %v, status: %v, err: %v", idList, status, err)
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////
//
//      Internal function
//
////////////////////////////////////////////////////////////////////////////////////

func assembleUserCondition(id int64, loginName string) string {
	condition := ""
	if id > 0 {
		condition = fmt.Sprintf("id = %d", id)
	}

	if loginName != "" {
		if condition != "" {
			condition = fmt.Sprintf("%s AND login_name = '%s'", condition, loginName)
		} else {
			condition = fmt.Sprintf("login_name ='%s'", loginName)
		}
	}

	// log.Debugf("assemble condition: %s\n", condition)
	return condition
}
