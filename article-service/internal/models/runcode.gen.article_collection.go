package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticleCollectionMgr struct {
	*_BaseMgr
}

// ArticleCollectionMgr open func
func ArticleCollectionMgr(db *gorm.DB) *_ArticleCollectionMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleCollectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleCollectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article_collection"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleCollectionMgr) GetTableName() string {
	return "article_collection"
}

// Get 获取
func (obj *_ArticleCollectionMgr) Get() (result ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleCollectionMgr) Gets() (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticleCollectionMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章编号
func (obj *_ArticleCollectionMgr) WithArticleID(articleID int64) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithUserID user_id获取 会员MID
func (obj *_ArticleCollectionMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStatus status获取
func (obj *_ArticleCollectionMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ArticleCollectionMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleCollectionMgr) GetByOption(opts ...Option) (result ArticleCollection, err error) {
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
func (obj *_ArticleCollectionMgr) GetByOptions(opts ...Option) (results []*ArticleCollection, err error) {
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
func (obj *_ArticleCollectionMgr) GetFromID(id uint64) (result ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticleCollectionMgr) GetBatchFromID(ids []uint64) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章编号
func (obj *_ArticleCollectionMgr) GetFromArticleID(articleID int64) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找 文章编号
func (obj *_ArticleCollectionMgr) GetBatchFromArticleID(articleIDs []int64) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 会员MID
func (obj *_ArticleCollectionMgr) GetFromUserID(userID uint64) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 会员MID
func (obj *_ArticleCollectionMgr) GetBatchFromUserID(userIDs []uint64) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ArticleCollectionMgr) GetFromStatus(status bool) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ArticleCollectionMgr) GetBatchFromStatus(statuss []bool) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ArticleCollectionMgr) GetFromCreateAt(createAt time.Time) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ArticleCollectionMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticleCollectionMgr) FetchByPrimaryKey(id uint64) (result ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleCollectionCid  获取多个内容
func (obj *_ArticleCollectionMgr) FetchIndexByIDxDataArticleCollectionCid(articleID int64) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleCollectionMid  获取多个内容
func (obj *_ArticleCollectionMgr) FetchIndexByIDxDataArticleCollectionMid(userID uint64) (results []*ArticleCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
