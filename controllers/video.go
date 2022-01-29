package controllers

import (
	"github.com/astaxie/beego"
	"go-dgVideoApi/models"
)

type VideoController struct {
	beego.Controller
}

//频道页-获取顶部广告
// @router /channel/advert [*]
func (v *VideoController) ChannelAdvert() {
	channelId, _ := v.GetInt("channelId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelAdvert(channelId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试")
		v.ServeJSON()
	}
}

// 频道页-获取正在热播
// @router /channel/hot [*]
func (v *VideoController) ChannelHotList() {
	channelId, _ := v.GetInt("channelId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelHotList(channelId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// 频道页-根据频道地区ID获取推荐视频
// @router /channel/recommend/region [*]
func (v *VideoController) ChannelRecommendRegionList() {
	channelId, _ := v.GetInt("channelId")
	regionId, _ := v.GetInt("regionId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}
	if regionId == 0 {
		v.Data["json"] = ReturnError(4002, "必须指定频道地区")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// 频道页-根据频道类型获取推荐视频
// @router /channel/recommend/type [*]
func (v *VideoController) GetChannelRecommendTypeList() {
	channelId, _ := v.GetInt("channelId")
	typeId, _ := v.GetInt("typeId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}
	if typeId == 0 {
		v.Data["json"] = ReturnError(4002, "必须指定频道类型")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// 根据传入参数获取视频列表
// @router /channel/video [*]
func (v *VideoController) ChannelVideo() {
	// 获取频道ID
	channelId, _ := v.GetInt("channelId")
	// 获取频道地区ID
	regionId, _ := v.GetInt("regionId")
	//获取频道类型ID
	typeId, _ := v.GetInt("typeId")
	//获取状态
	end := v.GetString("end")
	//获取排序
	sort := v.GetString("sort")
	//获取页码信息
	limit, _ := v.GetInt("limit")
	offset, _ := v.GetInt("offset")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

//获取视频详情
// @router /video/info [*]
func (v *VideoController) VideoInfo() {
	videoId, _ := v.GetInt("videoId")
	if videoId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定视频ID")
		v.ServeJSON()
	}

	video, err := models.GetVideoInfo(videoId)
	// 使用redis版
	//video, err := models.RedisGetVideoInfo(videoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", video, 1)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试")
		v.ServeJSON()
	}
}

//获取视频剧集列表
// @router /video/episodes/list [*]
func (v *VideoController) VideoEpisodesList() {
	videoId, _ := v.GetInt("videoId")
	if videoId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定视频ID")
		v.ServeJSON()
	}

	num, episodes, err := models.GetVideoEpisodesList(videoId)
	//使用redis版本
	//num, episodes, err := models.RedisGetVideoEpisodesList(videoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", episodes, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试")
		v.ServeJSON()
	}
}

//我的视频管理
// @router /user/video [*]
func (v *VideoController) UserVideo() {
	uid, _ := v.GetInt("uid")
	if uid == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定用户")
		v.ServeJSON()
	}

	num, videos, err := models.GetUserVideo(uid)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "成功", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}
