package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticlePanShareMgr struct {
	*_BaseMgr
}

// ArticlePanShareMgr open func
func ArticlePanShareMgr(db *gorm.DB) *_ArticlePanShareMgr {
	if db == nil {
		panic(fmt.Errorf("ArticlePanShareMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticlePanShareMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article_pan_share"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticlePanShareMgr) GetTableName() string {
	return "article_pan_share"
}

// Get 获取
func (obj *_ArticlePanShareMgr) Get() (result ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticlePanShareMgr) Gets() (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticlePanShareMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章编号
func (obj *_ArticlePanShareMgr) WithArticleID(articleID int64) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithPanShareID pan_share_id获取 分享ID
func (obj *_ArticlePanShareMgr) WithPanShareID(panShareID uint64) Option {
	return optionFunc(func(o *options) { o.query["pan_share_id"] = panShareID })
}

// WithStatus status获取
func (obj *_ArticlePanShareMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ArticlePanShareMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_ArticlePanShareMgr) GetByOption(opts ...Option) (result ArticlePanShare, err error) {
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
func (obj *_ArticlePanShareMgr) GetByOptions(opts ...Option) (results []*ArticlePanShare, err error) {
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
func (obj *_ArticlePanShareMgr) GetFromID(id uint64) (result ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticlePanShareMgr) GetBatchFromID(ids []uint64) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章编号
func (obj *_ArticlePanShareMgr) GetFromArticleID(articleID int64) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找 文章编号
func (obj *_ArticlePanShareMgr) GetBatchFromArticleID(articleIDs []int64) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromPanShareID 通过pan_share_id获取内容 分享ID
func (obj *_ArticlePanShareMgr) GetFromPanShareID(panShareID uint64) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pan_share_id = ?", panShareID).Find(&results).Error

	return
}

// GetBatchFromPanShareID 批量唯一主键查找 分享ID
func (obj *_ArticlePanShareMgr) GetBatchFromPanShareID(panShareIDs []uint64) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pan_share_id IN (?)", panShareIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ArticlePanShareMgr) GetFromStatus(status bool) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ArticlePanShareMgr) GetBatchFromStatus(statuss []bool) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ArticlePanShareMgr) GetFromCreateAt(createAt time.Time) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ArticlePanShareMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticlePanShareMgr) FetchByPrimaryKey(id uint64) (result ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleLikeCid  获取多个内容
func (obj *_ArticlePanShareMgr) FetchIndexByIDxDataArticleLikeCid(articleID int64) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleLikeMid  获取多个内容
func (obj *_ArticlePanShareMgr) FetchIndexByIDxDataArticleLikeMid(panShareID uint64) (results []*ArticlePanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pan_share_id = ?", panShareID).Find(&results).Error

	return
}
