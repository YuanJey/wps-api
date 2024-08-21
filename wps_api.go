package wpsApi

import (
	"wpsApi/avatar"
	"wpsApi/dept"
	"wpsApi/pkg/api_resp"
	"wpsApi/pkg/config"
	"wpsApi/pkg/sign"
	"wpsApi/user"
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
