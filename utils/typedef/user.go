package typedef

type UserInfo struct {
	Id        int64  `json:"id"`
	LoginName string `json:"loginName"`
	Username  string `json:"username"`
	Status    int64  `json:"status"`
}

// #############################################################################

type LoginReq struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type AddUserReq struct {
	Username  string `json:"username"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type AddUserResp struct {
}

type GetUserInfoReq struct {
	Id        int64  `json:"id"`
	LoginName string `json:"loginName"`
}

type GetUserInfoResp struct {
	Id        int64  `json:"id"`
	LoginName string `json:"loginName"`
	Username  string `json:"username"`
	Status    int64  `json:"status"`
}

type GetUserListReq struct {
	Status   int64 `json:"status"`
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type GetUserListResp struct {
	TotalCount int64       `json:"totalCount"`
	TotalPage  int64       `json:"totalPage"`
	PageNo     int64       `json:"pageNo"`
	PageSize   int64       `json:"pageSize"`
	Data       []*UserInfo `json:"data"`
}

type UpdateUserReq struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserResp struct {
}

type DeleteUsersReq struct {
	IdList []int64 `json:"idList"`
}

type DeleteUserResp struct {
}

type BatchUpdateUserStatusReq struct {
	IdList []int64 `json:"idList"`
	Status int64   `json:"status"`
}

type BatchUpdateUserStatusResp struct {
}
