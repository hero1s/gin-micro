package services

import (
	"context"
	"gin-micro/protos/user"
	"github.com/hero1s/golib/helpers/status"
	token2 "github.com/hero1s/golib/helpers/token"
	"github.com/hero1s/golib/log"
)

type UserService struct{}

func (userService *UserService) SignIn(ctx context.Context, userinfo *user.UserInfo, result *user.Result) error {
	result.Status = status.SaveStatusOK
	return nil
}

// 用户方法
func (userService *UserService) Login(ctx context.Context, loginParams *user.LoginParams, result *user.Result) error {
	log.Debug("call login:", loginParams)
	username := loginParams.GetAccount()
	password := loginParams.GetPassword()
	if username == "admin" && password == "123456" {
		result.Status = status.LoginStatusOK
		infos := map[string]string{"account": "admin", "nickname": "皮卡丘"}
		astoken, _ := token2.GenerateToken(infos)
		infos["ACCESS_TOKEN"] = astoken
		result.Map = infos
	} else {
		result.Status = status.LoginStatusErr
	}
	return nil
}

func (userService *UserService) SayHello(ctx context.Context, userinfo *user.NoneParam, result *user.Result) error {
	result.Status = status.SaveStatusOK
	return nil
}
