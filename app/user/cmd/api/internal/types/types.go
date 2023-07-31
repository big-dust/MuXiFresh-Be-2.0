// Code generated by goctl. DO NOT EDIT.
package types

type GetUserInfoReq struct {
	Authorization string `header:"Authorization"`
	UserId        string `path:"id"`
}

type GetUserInfoResp struct {
	Avatar    string `json:"avatar"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Group     string `json:"group"`
	StudentID string `json:"student_id"`
}

type SetUserInfoReq struct {
	Authorization string `header:"Authorization"`
	Avatar        string `json:"avatar"`
	NickName      string `json:"nickname"`
}

type SetUserInfoResp struct {
	Flag bool `json:"flag"`
}

type SetUserTypeReq struct {
	Authorization string `header:"Authorization"`
	Email         string `json:"email"`
	UserType      string `json:"user_type,options=[freshman,normal,admin,super_admin]"`
}

type SetUserTypeResp struct {
	Flag bool `json:"flag"`
}

type GetAdminListReq struct {
	Authorization string `header:"Authorization"`
	UserType      string `form:"user_type,options=[super_admin,admin]"`
	Page          int64  `form:"page"`
}

type One struct {
	UserId   string `json:"user_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type GetAdminListResp struct {
	List []One `json:"list"`
}
