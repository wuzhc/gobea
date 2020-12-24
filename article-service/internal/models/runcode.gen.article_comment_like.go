package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticleCommentLikeMgr struct {
	*_BaseMgr
}

// ArticleCommentLikeMgr open func
func ArticleCommentLikeMgr(db *gorm.DB) *_ArticleCommentLikeMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleCommentLikeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleCommentLikeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article_comment_like"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleCommentLikeMgr) GetTableName() string {
	return "article_comment_like"
}

// Get 获取
func (obj *_ArticleCommentLikeMgr) Get() (result ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleCommentLikeMgr) Gets() (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticleCommentLikeMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取
func (obj *_ArticleCommentLikeMgr) WithArticleID(articleID int) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithCommentID comment_id获取 评论编号
func (obj *_ArticleCommentLikeMgr) WithCommentID(commentID int64) Option {
	return optionFunc(func(o *options) { o.query["comment_id"] = commentID })
}

// WithUserID user_id获取 会员MID
func (obj *_ArticleCommentLikeMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ArticleCommentLikeMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithStatus status获取
func (obj *_ArticleCommentLikeMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleCommentLikeMgr) GetByOption(opts ...Option) (result ArticleCommentLike, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ArticleCommentLikeMgr) GetByOptions(opts ...Option) (results []*ArticleCommentLike, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ArticleCommentLikeMgr) GetFromID(id uint64) (result ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticleCommentLikeMgr) GetBatchFromID(ids []uint64) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容
func (obj *_ArticleCommentLikeMgr) GetFromArticleID(articleID int) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找
func (obj *_ArticleCommentLikeMgr) GetBatchFromArticleID(articleIDs []int) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromCommentID 通过comment_id获取内容 评论编号
func (obj *_ArticleCommentLikeMgr) GetFromCommentID(commentID int64) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("comment_id = ?", commentID).Find(&results).Error

	return
}

// GetBatchFromCommentID 批量唯一主键查找 评论编号
func (obj *_ArticleCommentLikeMgr) GetBatchFromCommentID(commentIDs []int64) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("comment_id IN (?)", commentIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 会员MID
func (obj *_ArticleCommentLikeMgr) GetFromUserID(userID uint64) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 会员MID
func (obj *_ArticleCommentLikeMgr) GetBatchFromUserID(userIDs []uint64) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ArticleCommentLikeMgr) GetFromCreateAt(createAt time.Time) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ArticleCommentLikeMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ArticleCommentLikeMgr) GetFromStatus(status bool) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ArticleCommentLikeMgr) GetBatchFromStatus(statuss []bool) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticleCommentLikeMgr) FetchByPrimaryKey(id uint64) (result ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleLikeCid  获取多个内容
func (obj *_ArticleCommentLikeMgr) FetchIndexByIDxDataArticleLikeCid(commentID int64) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("comment_id = ?", commentID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleLikeMid  获取多个内容
func (obj *_ArticleCommentLikeMgr) FetchIndexByIDxDataArticleLikeMid(userID uint64) (results []*ArticleCommentLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
