package logics

import (
	pb "article"
	"article/internal/models"
	"article/internal/utils"
)

func GetArticleDetail(article_id int) (*pb.Detail, error) {
	mgr := models.ArticleMgr(utils.DB())
	record, err := mgr.GetByOption(mgr.WithID(uint64(article_id)))
	if err != nil {
		return nil, err
	}

	return &pb.Detail{
		ArticleId:      utils.EncodeID(int(record.ID)),
		CategoryId:     utils.EncodeID(int(record.CategoryID)),
		CategoryName:   "xxxx",
		Title:          record.Title,
		Type:           record.Type,
		Logo:           record.Logo,
		Keywords:       record.Keywords,
		Remark:         record.Remark,
		Content:        record.Content,
		NumberLikes:    int32(record.NumberLikes),
		NumberComments: int32(record.NumberComment),
		NumberReads:    int32(record.NumberReads),
		NumberCollects: int32(record.NumberCollection),
		Status:         int32(record.Status),
		EditorType:     int32(record.EditorType),
		Author:         record.Author,
		Link:           record.Link,
		IsTop:          int32(record.IsTop),
		CreateAt:       "",
		User:           nil,
	}, nil
}
