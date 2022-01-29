package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"go-dgVideoApi/models"
	"net/http"
)

type BarrageController struct {
	beego.Controller
}

type WsData struct {
	CurrentTime int
	EpisodesId int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//获取弹幕
// @router /barrage/ws [*]
func (b *BarrageController) BarrageWs() {
	var (
		conn *websocket.Conn
		err error
		data []byte
		barrages []models.BarrageData
	)
	if conn, err = upgrader.Upgrade(b.Ctx.ResponseWriter, b.Ctx.Request, nil); err != nil {
		goto ERR
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		var wsData WsData
		json.Unmarshal([]byte(data), &wsData)
		endTime := wsData.CurrentTime + 60
		//获取弹幕数据
		_, barrages, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err == nil {
			if err := conn.WriteJSON(barrages); err != nil {
				goto ERR
			}
		}
	}

ERR:
	conn.Close()
}

//保存弹幕
// @router /barrage/save [*]
func (b *BarrageController) Save() {
	uid, _ := b.GetInt("uid")
	content := b.GetString("content")
	currentTime, _ := b.GetInt("currentTime")
	episodesId, _ := b.GetInt("episodesId")
	videoId, _ := b.GetInt("videoId")
	if content == "" {
		b.Data["json"] = ReturnError(4001, "弹幕不能为空")
		b.ServeJSON()
	}
	if uid == 0 {
		b.Data["json"] = ReturnError(4002, "请先登录")
		b.ServeJSON()
	}
	if episodesId == 0 {
		b.Data["json"] = ReturnError(4003, "必须指定剧集ID")
		b.ServeJSON()
	}
	if videoId == 0 {
		b.Data["json"] = ReturnError(4005, "必须指定视频ID")
		b.ServeJSON()
	}
	if currentTime == 0 {
		b.Data["json"] = ReturnError(4006, "必须指定视频播放时间")
		b.ServeJSON()
	}

	err := models.SaveBarrage(episodesId, videoId, currentTime, uid, content)
	if err == nil {
		b.Data["json"] = ReturnSuccess(0, "成功", "", 1)
		b.ServeJSON()
	} else {
		b.Data["json"] = ReturnError(5000, err)
		b.ServeJSON()
	}
}

