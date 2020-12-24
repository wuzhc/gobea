package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _EbookMgr struct {
	*_BaseMgr
}

// EbookMgr open func
func EbookMgr(db *gorm.DB) *_EbookMgr {
	if db == nil {
		panic(fmt.Errorf("EbookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EbookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ebook"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EbookMgr) GetTableName() string {
	return "ebook"
}

// Get 获取
func (obj *_EbookMgr) Get() (result Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EbookMgr) Gets() (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EbookMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取
func (obj *_EbookMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithArticleID article_id获取
func (obj *_EbookMgr) WithArticleID(articleID int) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithParentID parent_id获取
func (obj *_EbookMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithLevel level获取
func (obj *_EbookMgr) WithLevel(level int8) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithPath path获取
func (obj *_EbookMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithAnchor anchor获取
func (obj *_EbookMgr) WithAnchor(anchor string) Option {
	return optionFunc(func(o *options) { o.query["anchor"] = anchor })
}

// GetByOption 功能选项模式获取
func (obj *_EbookMgr) GetByOption(opts ...Option) (result Ebook, err error) {
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
func (obj *_EbookMgr) GetByOptions(opts ...Option) (results []*Ebook, err error) {
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
func (obj *_EbookMgr) GetFromID(id int) (result Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_EbookMgr) GetBatchFromID(ids []int) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_EbookMgr) GetFromTitle(title string) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_EbookMgr) GetBatchFromTitle(titles []string) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容
func (obj *_EbookMgr) GetFromArticleID(articleID int) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找
func (obj *_EbookMgr) GetBatchFromArticleID(articleIDs []int) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_EbookMgr) GetFromParentID(parentID int) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量唯一主键查找
func (obj *_EbookMgr) GetBatchFromParentID(parentIDs []int) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容
func (obj *_EbookMgr) GetFromLevel(level int8) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量唯一主键查找
func (obj *_EbookMgr) GetBatchFromLevel(levels []int8) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level IN (?)", levels).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容
func (obj *_EbookMgr) GetFromPath(path string) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("path = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量唯一主键查找
func (obj *_EbookMgr) GetBatchFromPath(paths []string) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("path IN (?)", paths).Find(&results).Error

	return
}

// GetFromAnchor 通过anchor获取内容
func (obj *_EbookMgr) GetFromAnchor(anchor string) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("anchor = ?", anchor).Find(&results).Error

	return
}

// GetBatchFromAnchor 批量唯一主键查找
func (obj *_EbookMgr) GetBatchFromAnchor(anchors []string) (results []*Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("anchor IN (?)", anchors).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EbookMgr) FetchByPrimaryKey(id int) (result Ebook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
