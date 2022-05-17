package msgBoard

import (
	"gim/serializer"
	"gim/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MsgList(c *gin.Context) {
	var listMsg service.MsgBoardListService
	if err := c.ShouldBind(&listMsg); err == nil {
		if list, err, totalNum, totalPageNum := listMsg.List(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildMsgBoardListResponse(list,totalNum,totalPageNum)
			c.JSON(200, res)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func InsertMsg(c *gin.Context) {
	var insertMsg service.MsgBoardSaveService
	if err := c.ShouldBind(&insertMsg); err == nil {
		if _, err := insertMsg.Save(c); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, gin.H{
				"msg": "新增成功",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
