package serializer

import "gim/model"

// MsgBoard 用户序列化器
type MsgBoard struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// MsgBoardResponse 单个用户序列化
type MsgBoardListResponse struct {
	List []MsgBoard `json:"list"`
	TotalCount int `json:"total_count"`
	TotalPageCount float64 `json:"total_page_count"`
}

// MsgBoard 序列化用户
func BuildMsgBoard(msgBoard model.MsgBoard) MsgBoard {
	return MsgBoard{
		ID:      msgBoard.ID,
		Title:   msgBoard.Title,
		Content: msgBoard.Content,
	}
}

// BuildMsgBoardListResponse 序列化用户响应
func BuildMsgBoardListResponse(msgBoard []MsgBoard,TotalCount int,TotalPageCount float64) MsgBoardListResponse {
	return MsgBoardListResponse{
		List: msgBoard,
		TotalCount: TotalCount,
		TotalPageCount: TotalPageCount,
	}
}

// BuildMsgBoards 序列化留言板列表
func BuildMsgBoards(msgBoards []model.MsgBoard) (formatMsgBoards []MsgBoard) {
	for _, item := range msgBoards {
		msgBoard := BuildMsgBoard(item)
		formatMsgBoards = append(formatMsgBoards, msgBoard)
	}
	return formatMsgBoards
}
