package common

import (
	"github.com/YuanJey/wpsApi/pkg/api_resp"
	"strconv"
)

type User struct {
	api_resp.Member
}

func (u *User) GetDeptId() string {
	return strconv.Itoa(u.DeptNum)
}
