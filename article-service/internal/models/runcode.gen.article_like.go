package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticleLikeMgr struct {
	*_BaseMgr
}

// ArticleLikeMgr open func
func ArticleLikeMgr(db *gorm.DB) *_ArticleLikeMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleLikeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleLikeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article_like"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleLikeMgr) GetTableName() string {
	return "article_like"
}

// Get 获取
func (obj *_ArticleLikeMgr) Get() (result ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleLikeMgr) Gets() (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticleLikeMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章编号
func (obj *_ArticleLikeMgr) WithArticleID(articleID int64) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithUserID user_id获取 会员MID
func (obj *_ArticleLikeMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStatus status获取
func (obj *_ArticleLikeMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ArticleLikeMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleLikeMgr) GetByOption(opts ...Option) (result ArticleLike, err error) {
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
func (obj *_ArticleLikeMgr) GetByOptions(opts ...Option) (results []*ArticleLike, err error) {
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
func (obj *_ArticleLikeMgr) GetFromID(id uint64) (result ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticleLikeMgr) GetBatchFromID(ids []uint64) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章编号
func (obj *_ArticleLikeMgr) GetFromArticleID(articleID int64) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找 文章编号
func (obj *_ArticleLikeMgr) GetBatchFromArticleID(articleIDs []int64) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 会员MID
func (obj *_ArticleLikeMgr) GetFromUserID(userID int64) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 会员MID
func (obj *_ArticleLikeMgr) GetBatchFromUserID(userIDs []int64) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ArticleLikeMgr) GetFromStatus(status bool) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ArticleLikeMgr) GetBatchFromStatus(statuss []bool) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ArticleLikeMgr) GetFromCreateAt(createAt time.Time) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ArticleLikeMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticleLikeMgr) FetchByPrimaryKey(id uint64) (result ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleLikeCid  获取多个内容
func (obj *_ArticleLikeMgr) FetchIndexByIDxDataArticleLikeCid(articleID int64) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleLikeMid  获取多个内容
func (obj *_ArticleLikeMgr) FetchIndexByIDxDataArticleLikeMid(userID int64) (results []*ArticleLike, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
