package wpsApi

import (
	"github.com/yuanjie/wps-api/avatar"
	"github.com/yuanjie/wps-api/dept"
	"github.com/yuanjie/wps-api/pkg/api_resp"
	"github.com/yuanjie/wps-api/pkg/config"
	"github.com/yuanjie/wps-api/pkg/sign"
	"github.com/yuanjie/wps-api/user"
)

type AllWpsUser struct {
	WidUserList map[string]*api_resp.Member
	TidUserList map[string]*api_resp.Member
}
type sdk struct {
	Dept   *dept.Dept
	User   *user.User
	Avatar *avatar.Avatar
}

func newSDK(dept *dept.Dept, user *user.User, avatar *avatar.Avatar) *sdk {
	return &sdk{Dept: dept, User: user, Avatar: avatar}
}

var Sdk sdk

func init() {
	wps4Sign := sign.NewWps4Sign(config.Config.Env.AK, config.Config.Env.SK)
	Sdk = *newSDK(dept.NewDept(wps4Sign, config.Config.WPS.Addr), user.NewUser(wps4Sign, config.Config.WPS.Addr), avatar.NewAvatar(config.Config.WPS.Addr, wps4Sign))
}
