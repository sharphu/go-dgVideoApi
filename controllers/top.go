package controllers

import (
	"github.com/astaxie/beego"
	"go-dgVideoApi/models"
)

type TopController struct {
	beego.Controller
}

//根据频道获取排行榜
// @router /channel/top [*]
func (t *TopController) ChannelTop() {
	//获取频道ID
	channelId, _ := t.GetInt("channelId")
	if channelId == 0 {
		t.Data["json"] = ReturnError(4001, "必须指定频道")
		t.ServeJSON()
	}

	num, videos, err := models.GetChannelTop(channelId)
	//使用redis版本
	//num, videos, err := models.RedisGetChannelTop(channelId)
	if err == nil {
		t.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		t.ServeJSON()
	} else {
		t.Data["json"] = ReturnError(4004, "没有相关内容")
		t.ServeJSON()
	}
}

//根据类型获取排行榜
// @router /type/top [*]
func (t *TopController) TypeTop() {
	typeId, _ := t.GetInt("typeId")
	if typeId == 0 {
		t.Data["json"] = ReturnError(4001, "必须指定类型")
		t.ServeJSON()
	}

	num, videos, err := models.GetTypeTop(typeId)
	//使用redis版本
	//num, videos, err := models.RedisGetTypeTop(typeId)
	if err == nil {
		t.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		t.ServeJSON()
	} else {
		t.Data["json"] = ReturnError(4004, "没有相关内容")
		t.ServeJSON()
	}
}
