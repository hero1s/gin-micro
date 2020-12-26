package controls

import (
	"gin-micro/api/clients"
	"gin-micro/protos/user"
	"github.com/hero1s/golib/helpers/status"
	"github.com/hero1s/golib/log"
	"github.com/hero1s/golib/web/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserControl struct {
	response.Control
	response.Router
}

func (c *UserControl) InitRoutes(r *gin.Engine) {
	user := r.Group("user")
	user.POST("login", c.Login)
	user.POST("register", c.SignIn)
	user.POST("say-hello", c.SayHello)
}

// @Summary 用户登陆接口
// @Tags UserControl
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param code     query string false "验证码"
// @Success 200 {object} response.JsonObject
// @Router /user/login [post]
func (c *UserControl) Login(ctx *gin.Context) {
	log.Debug("call login")
	params := &user.LoginParams{}
	if err := ctx.Bind(params); err == nil {
		log.Debug("params:%", *params)
		result, err := clients.UserService.Login(ctx, params)
		if err == nil {
			c.SuccessContent(ctx, status.StatusText(result.Status), result.Map)
		} else {
			log.Error(zap.Any("call micro svr err:", err))
			c.InternalError(ctx, err.Error())
		}
	} else {
		log.Debug("bind err : ", err)
		c.BindingError(ctx)
	}
	ctx.Abort()
}

// @Summary 用户注册接口
// @Tags UserControl
// @Accept json
// @Produce json
// @Success 200 {object} response.JsonObject
// @Router /user/register [post]
func (c *UserControl) SignIn(ctx *gin.Context) {
	result, err := clients.UserService.SignIn(ctx, &user.UserInfo{})
	if err == nil {
		c.SuccessContent(ctx, status.StatusText(result.Status), result.Map)
	} else {
		c.InternalError(ctx, err.Error())
	}
	ctx.Abort()
}

type Resp struct {
	Name string `json:"name" desc:"名字"`
	Age  int    `json:"age" desc:"年龄"`
}

// @Summary 用户hello
// @Tags UserControl
// @Accept json
// @Produce json
// @Success 200 {object} controls.Resp
// @Router /user/say-hello [post]
func (c *UserControl) SayHello(ctx *gin.Context) {
	result, err := clients.UserService.SayHello(ctx, &user.NoneParam{})
	if err == nil {
		c.SuccessContent(ctx, status.StatusText(result.Status), result.Map)
	} else {
		c.InternalError(ctx, err.Error())
	}
	ctx.Abort()
}
