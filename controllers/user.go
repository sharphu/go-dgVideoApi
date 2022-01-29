package controllers

import (
	"go-dgVideoApi/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// 用户注册功能
// @router /register/save [post]
func (u *UserController) SaveRegister()  {
	var (
		mobile string
		password string
		err error
	)
	mobile = u.GetString("mobile")
	password = u.GetString("password")
	if mobile == "" {
		u.Data["json"] = ReturnError(4001, "手机号不能为空")
		u.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		u.Data["json"] = ReturnError(4002, "手机号格式不正确")
		u.ServeJSON()
	}
	if password == "" {
		u.Data["json"] = ReturnError(4003, "密码不能为空")
		u.ServeJSON()
	}

	// 判断手机号是否已经注册
	status := models.IsUserMobile(mobile)
	if status {
		u.Data["json"] = ReturnError(4005, "此手机号已经注册")
		u.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err == nil {
			u.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			u.ServeJSON()
		} else {
			u.Data["json"] = ReturnError(5000, err)
			u.ServeJSON()
		}
	}
}

// 用户登录
// @router /login/do [*]
func (u *UserController) LoginDo() {
	mobile := u.GetString("mobile")
	password := u.GetString("password")

	if mobile == "" {
		u.Data["json"] = ReturnError(4001, "手机号不能为空")
		u.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		u.Data["json"] = ReturnError(4002, "手机号格式不正确")
		u.ServeJSON()
	}
	if password == "" {
		u.Data["json"] = ReturnError(4003, "密码不能为空")
		u.ServeJSON()
	}
	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		u.Data["json"] = ReturnSuccess(0, "登陆成功", map[string]interface{}{"uid": uid, "username": name}, 1)
		u.ServeJSON()
	} else {
		u.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		u.ServeJSON()
	}
}

//批量发送通知消息
// @router /send/message [*]
func (u *UserController) SendMessageDo() {
	uids := u.GetString("uids")
	content := u.GetString("content")
	if uids == "" {
		u.Data["json"] = ReturnError(4001, "请填写接收人")
		u.ServeJSON()
	}
	if content == "" {
		u.Data["json"] = ReturnError(4002, "请填写发送内容")
		u.ServeJSON()
	}

	messageId, err := models.SendMessageDo(content)
	if err == nil {
		uidConfig := strings.Split(uids, ",")
		for _, v := range uidConfig {
			userId, _ := strconv.Atoi(v)
			models.SendMessageUser(userId, messageId)
		}
		u.Data["json"] = ReturnSuccess(0, "发送成功", "", 1)
		u.ServeJSON()
	} else {
		u.Data["json"] = ReturnError(5000, "发送失败，请联系客服")
		u.ServeJSON()
	}
}
