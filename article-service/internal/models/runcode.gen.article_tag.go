package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ArticleTagMgr struct {
	*_BaseMgr
}

// ArticleTagMgr open func
func ArticleTagMgr(db *gorm.DB) *_ArticleTagMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleTagMgr) GetTableName() string {
	return "article_tag"
}

// Get 获取
func (obj *_ArticleTagMgr) Get() (result ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleTagMgr) Gets() (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticleTagMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章编号
func (obj *_ArticleTagMgr) WithArticleID(articleID uint64) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithTagID tag_id获取 标签ID
func (obj *_ArticleTagMgr) WithTagID(tagID uint64) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithStatus status获取
func (obj *_ArticleTagMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleTagMgr) GetByOption(opts ...Option) (result ArticleTag, err error) {
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
func (obj *_ArticleTagMgr) GetByOptions(opts ...Option) (results []*ArticleTag, err error) {
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
func (obj *_ArticleTagMgr) GetFromID(id int64) (result ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticleTagMgr) GetBatchFromID(ids []int64) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章编号
func (obj *_ArticleTagMgr) GetFromArticleID(articleID uint64) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找 文章编号
func (obj *_ArticleTagMgr) GetBatchFromArticleID(articleIDs []uint64) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromTagID 通过tag_id获取内容 标签ID
func (obj *_ArticleTagMgr) GetFromTagID(tagID uint64) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id = ?", tagID).Find(&results).Error

	return
}

// GetBatchFromTagID 批量唯一主键查找 标签ID
func (obj *_ArticleTagMgr) GetBatchFromTagID(tagIDs []uint64) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id IN (?)", tagIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ArticleTagMgr) GetFromStatus(status bool) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ArticleTagMgr) GetBatchFromStatus(statuss []bool) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticleTagMgr) FetchByPrimaryKey(id int64) (result ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleLikeCid  获取多个内容
func (obj *_ArticleTagMgr) FetchIndexByIDxDataArticleLikeCid(articleID uint64) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleLikeMid  获取多个内容
func (obj *_ArticleTagMgr) FetchIndexByIDxDataArticleLikeMid(tagID uint64) (results []*ArticleTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id = ?", tagID).Find(&results).Error

	return
}
