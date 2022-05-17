/**
 *
Class MsgBoard
 * @package
@version 1.164
@time 2022/5/17 14:30
*/
package model

type MsgBoard struct {
	BaseModel
	Title string
	Content string
	Uid uint
}

func (MsgBoard) TableName() string {
	return "message_board"
}