package db

type NewsArticleContent struct {
	Id      int    `gorm:"id" json:"id"`           // 文章id
	Content string `gorm:"content" json:"content"` // 文章内容
}
