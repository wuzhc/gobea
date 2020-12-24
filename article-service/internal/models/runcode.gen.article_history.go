package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticleHistoryMgr struct {
	*_BaseMgr
}

// ArticleHistoryMgr open func
func ArticleHistoryMgr(db *gorm.DB) *_ArticleHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleHistoryMgr) GetTableName() string {
	return "article_history"
}

// Get 获取
func (obj *_ArticleHistoryMgr) Get() (result ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleHistoryMgr) Gets() (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticleHistoryMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章编号
func (obj *_ArticleHistoryMgr) WithArticleID(articleID uint64) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithUserID user_id获取 会员MID
func (obj *_ArticleHistoryMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ArticleHistoryMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleHistoryMgr) GetByOption(opts ...Option) (result ArticleHistory, err error) {
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
func (obj *_ArticleHistoryMgr) GetByOptions(opts ...Option) (results []*ArticleHistory, err error) {
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
func (obj *_ArticleHistoryMgr) GetFromID(id uint64) (result ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticleHistoryMgr) GetBatchFromID(ids []uint64) (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章编号
func (obj *_ArticleHistoryMgr) GetFromArticleID(articleID uint64) (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找 文章编号
func (obj *_ArticleHistoryMgr) GetBatchFromArticleID(articleIDs []uint64) (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 会员MID
func (obj *_ArticleHistoryMgr) GetFromUserID(userID uint64) (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 会员MID
func (obj *_ArticleHistoryMgr) GetBatchFromUserID(userIDs []uint64) (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ArticleHistoryMgr) GetFromCreateAt(createAt time.Time) (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ArticleHistoryMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticleHistoryMgr) FetchByPrimaryKey(id uint64) (result ArticleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
