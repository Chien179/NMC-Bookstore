package models

type DislikeResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	ReviewId  int64  `json:"review_id"`
	IsDislike bool   `json:"is_dislike"`
}

type DislikeRequest struct {
	Username string `json:"username" binding:"required"`
	ReviewId int64  `json:"review_id" binding:"required"`
}
