package models

// TagMd 标签表
type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`
	Articles []ArticleModel `gorm:"many2many:article_tag_models" json:"-"` // 关联的文章列表
}
