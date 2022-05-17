package service

import (
	"gim/model"
	"gim/serializer"
	"math"
)

type MsgBoardListService struct {
	Page     int    `form:"page" json:"page" binding:"required,min=0"`
	PageSize int    `form:"page_size" json:"page_size" binding:"required,min=0"`
	Keyword  string `form:"keyword" json:"keyword"`
}

/**
留言版列表
*/
func (msg *MsgBoardListService) List() ([]serializer.MsgBoard, *serializer.Response, int, float64) {
	var msgBoard []model.MsgBoard
	var totalNum int
	model.DB.Model(&msgBoard).Count(&totalNum)
	selectModel := model.DB.Limit(msg.PageSize).Offset((msg.Page - 1) * msg.PageSize).Order("id desc")
	if msg.Keyword != "" {
		selectModel = selectModel.Where("title like ?", "%"+msg.Keyword+"%").
			Or("content like ?", "%"+msg.Keyword+"%")
	}
	err := selectModel.Find(&msgBoard).Error
	if err != nil {
		return serializer.BuildMsgBoards(msgBoard), &serializer.Response{
			Status: 40001,
			Msg:    "失败",
		}, 0, 0
	}
	totalPageNum := math.Ceil(float64(totalNum) / float64(msg.PageSize))
	return serializer.BuildMsgBoards(msgBoard), nil, totalNum, totalPageNum
}
