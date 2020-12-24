package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TagCategoryMgr struct {
	*_BaseMgr
}

// TagCategoryMgr open func
func TagCategoryMgr(db *gorm.DB) *_TagCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("TagCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TagCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tag_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TagCategoryMgr) GetTableName() string {
	return "tag_category"
}

// Get 获取
func (obj *_TagCategoryMgr) Get() (result TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagCategoryMgr) Gets() (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TagCategoryMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCategoryID category_id获取 分类ID
func (obj *_TagCategoryMgr) WithCategoryID(categoryID uint32) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithTagID tag_id获取 标签ID
func (obj *_TagCategoryMgr) WithTagID(tagID int64) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithStatus status获取
func (obj *_TagCategoryMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_TagCategoryMgr) GetByOption(opts ...Option) (result TagCategory, err error) {
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
func (obj *_TagCategoryMgr) GetByOptions(opts ...Option) (results []*TagCategory, err error) {
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
func (obj *_TagCategoryMgr) GetFromID(id uint64) (result TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_TagCategoryMgr) GetBatchFromID(ids []uint64) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类ID
func (obj *_TagCategoryMgr) GetFromCategoryID(categoryID uint32) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量唯一主键查找 分类ID
func (obj *_TagCategoryMgr) GetBatchFromCategoryID(categoryIDs []uint32) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromTagID 通过tag_id获取内容 标签ID
func (obj *_TagCategoryMgr) GetFromTagID(tagID int64) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id = ?", tagID).Find(&results).Error

	return
}

// GetBatchFromTagID 批量唯一主键查找 标签ID
func (obj *_TagCategoryMgr) GetBatchFromTagID(tagIDs []int64) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id IN (?)", tagIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_TagCategoryMgr) GetFromStatus(status bool) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_TagCategoryMgr) GetBatchFromStatus(statuss []bool) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TagCategoryMgr) FetchByPrimaryKey(id uint64) (result TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleLikeCid  获取多个内容
func (obj *_TagCategoryMgr) FetchIndexByIDxDataArticleLikeCid(categoryID uint32) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id = ?", categoryID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleLikeMid  获取多个内容
func (obj *_TagCategoryMgr) FetchIndexByIDxDataArticleLikeMid(tagID int64) (results []*TagCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id = ?", tagID).Find(&results).Error

	return
}
