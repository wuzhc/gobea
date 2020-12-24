package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _EbookContentMgr struct {
	*_BaseMgr
}

// EbookContentMgr open func
func EbookContentMgr(db *gorm.DB) *_EbookContentMgr {
	if db == nil {
		panic(fmt.Errorf("EbookContentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EbookContentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ebook_content"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EbookContentMgr) GetTableName() string {
	return "ebook_content"
}

// Get 获取
func (obj *_EbookContentMgr) Get() (result EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EbookContentMgr) Gets() (results []*EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EbookContentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithEbookID ebook_id获取
func (obj *_EbookContentMgr) WithEbookID(ebookID int) Option {
	return optionFunc(func(o *options) { o.query["ebook_id"] = ebookID })
}

// WithContent content获取
func (obj *_EbookContentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// GetByOption 功能选项模式获取
func (obj *_EbookContentMgr) GetByOption(opts ...Option) (result EbookContent, err error) {
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
func (obj *_EbookContentMgr) GetByOptions(opts ...Option) (results []*EbookContent, err error) {
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
func (obj *_EbookContentMgr) GetFromID(id int) (result EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_EbookContentMgr) GetBatchFromID(ids []int) (results []*EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromEbookID 通过ebook_id获取内容
func (obj *_EbookContentMgr) GetFromEbookID(ebookID int) (results []*EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ebook_id = ?", ebookID).Find(&results).Error

	return
}

// GetBatchFromEbookID 批量唯一主键查找
func (obj *_EbookContentMgr) GetBatchFromEbookID(ebookIDs []int) (results []*EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ebook_id IN (?)", ebookIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_EbookContentMgr) GetFromContent(content string) (results []*EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_EbookContentMgr) GetBatchFromContent(contents []string) (results []*EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EbookContentMgr) FetchByPrimaryKey(id int) (result EbookContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
