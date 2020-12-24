package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ToolCollectionMgr struct {
	*_BaseMgr
}

// ToolCollectionMgr open func
func ToolCollectionMgr(db *gorm.DB) *_ToolCollectionMgr {
	if db == nil {
		panic(fmt.Errorf("ToolCollectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ToolCollectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tool_collection"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ToolCollectionMgr) GetTableName() string {
	return "tool_collection"
}

// Get 获取
func (obj *_ToolCollectionMgr) Get() (result ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ToolCollectionMgr) Gets() (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ToolCollectionMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithToolID tool_id获取 工具编号
func (obj *_ToolCollectionMgr) WithToolID(toolID uint64) Option {
	return optionFunc(func(o *options) { o.query["tool_id"] = toolID })
}

// WithUserID user_id获取 会员MID
func (obj *_ToolCollectionMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStatus status获取
func (obj *_ToolCollectionMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ToolCollectionMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_ToolCollectionMgr) GetByOption(opts ...Option) (result ToolCollection, err error) {
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
func (obj *_ToolCollectionMgr) GetByOptions(opts ...Option) (results []*ToolCollection, err error) {
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
func (obj *_ToolCollectionMgr) GetFromID(id uint64) (result ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ToolCollectionMgr) GetBatchFromID(ids []uint64) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromToolID 通过tool_id获取内容 工具编号
func (obj *_ToolCollectionMgr) GetFromToolID(toolID uint64) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tool_id = ?", toolID).Find(&results).Error

	return
}

// GetBatchFromToolID 批量唯一主键查找 工具编号
func (obj *_ToolCollectionMgr) GetBatchFromToolID(toolIDs []uint64) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tool_id IN (?)", toolIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 会员MID
func (obj *_ToolCollectionMgr) GetFromUserID(userID int64) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 会员MID
func (obj *_ToolCollectionMgr) GetBatchFromUserID(userIDs []int64) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ToolCollectionMgr) GetFromStatus(status bool) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ToolCollectionMgr) GetBatchFromStatus(statuss []bool) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ToolCollectionMgr) GetFromCreateAt(createAt time.Time) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ToolCollectionMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ToolCollectionMgr) FetchByPrimaryKey(id uint64) (result ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleLikeCid  获取多个内容
func (obj *_ToolCollectionMgr) FetchIndexByIDxDataArticleLikeCid(toolID uint64) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tool_id = ?", toolID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleLikeMid  获取多个内容
func (obj *_ToolCollectionMgr) FetchIndexByIDxDataArticleLikeMid(userID int64) (results []*ToolCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
