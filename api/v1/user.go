package v1

import (
	"go-skeleton/serializer"
	"go-skeleton/utils/log"
	"go-skeleton/utils/typedef"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req typedef.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json when login, err: %v", err)
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	token, err := svc.Login(req.LoginName, req.Password)
	if err != nil {
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	loginResp := typedef.LoginResp{
		Token: token,
	}
	res := serializer.ReplySuccess(loginResp)
	c.JSON(http.StatusOK, res)
}

func AddUser(c *gin.Context) {
	var req typedef.AddUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json when add user, err: %v", err)
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	err = svc.AddUser(req.Username, req.LoginName, req.Password)
	if err != nil {
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	res := serializer.ReplySuccess(nil)
	c.JSON(http.StatusOK, res)
}

func GetUserInfo(c *gin.Context) {
	var req typedef.GetUserInfoReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json when get user info, err: %v", err)
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	userInfo, err := svc.GetUserInfo(req.Id, req.LoginName)
	if err != nil {
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	getUserInfoResp := typedef.GetUserInfoResp{
		Id:        userInfo.Id,
		LoginName: userInfo.LoginName,
		Username:  userInfo.Username,
		Status:    userInfo.Status,
	}
	res := serializer.ReplySuccess(getUserInfoResp)
	c.JSON(http.StatusOK, res)
}

func GetUserList(c *gin.Context) {
	var req typedef.GetUserListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json when get user list, err: %v", err)
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	userList, count, err := svc.GetUserList(req.Status, req.PageNo, req.PageSize)
	if err != nil {
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	totalPage := math.Ceil(float64(count) / float64(req.PageSize))
	getUserListResp := typedef.GetUserListResp{
		TotalCount: count,
		TotalPage:  int64(totalPage),
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
		Data:       userList,
	}
	res := serializer.ReplySuccess(getUserListResp)
	c.JSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context) {
	var req typedef.UpdateUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json when update user, err: %v", err)
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	err = svc.UpdateUser(req.Id, req.Username, req.Password)
	if err != nil {
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}
	res := serializer.ReplySuccess(nil)
	c.JSON(http.StatusOK, res)
}

func DeleteUsers(c *gin.Context) {
	var req typedef.DeleteUsersReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json when delete users, err: %v", err)
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	err = svc.DeleteUsers(req.IdList)
	if err != nil {
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}
	res := serializer.ReplySuccess(nil)
	c.JSON(http.StatusOK, res)
}

func BatchUpdateUserStatus(c *gin.Context) {
	var req typedef.BatchUpdateUserStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("failed to bind json when batch update user status, err: %v", err)
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}

	err = svc.BatchUpdateUserStatus(req.IdList, req.Status)
	if err != nil {
		res := serializer.ReplyError(err)
		c.JSON(http.StatusOK, res)
		return
	}
	res := serializer.ReplySuccess(nil)
	c.JSON(http.StatusOK, res)
}
