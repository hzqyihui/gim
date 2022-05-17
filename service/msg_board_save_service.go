/**
 *
Class msg_board_save_service
 * @package
@version 1.164
@time 2022/5/17 12:28
*/
package service

import (
	"gim/model"
	"gim/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MsgBoardSaveService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=45"`
	Content string `form:"content" json:"content" binding:"required,min=1,max=1900"`
}

/**
留言
*/
func (msg *MsgBoardSaveService) Save(ctx *gin.Context) (model.MsgBoard, *serializer.Response) {
	session := sessions.Default(ctx)
	msgBoard := model.MsgBoard{
		Uid:     session.Get("user_id").(uint),
		//Uid:     1,
		Title:   msg.Title,
		Content: msg.Content,
	}
	if err := model.DB.Create(&msgBoard).Error; err != nil {
		return msgBoard, &serializer.Response{
			Status: 40001,
			Msg:    "留言失败",
		}
	}
	return msgBoard, nil
}
