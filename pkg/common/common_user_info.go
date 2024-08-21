package common

import (
	"strconv"
	"wpsApi/pkg/api_resp"
)

type User struct {
	api_resp.Member
}

func (u *User) GetDeptId() string {
	return strconv.Itoa(u.DeptNum)
}
