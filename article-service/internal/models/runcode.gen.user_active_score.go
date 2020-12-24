package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserActiveScoreMgr struct {
	*_BaseMgr
}

// UserActiveScoreMgr open func
func UserActiveScoreMgr(db *gorm.DB) *_UserActiveScoreMgr {
	if db == nil {
		panic(fmt.Errorf("UserActiveScoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserActiveScoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_active_score"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserActiveScoreMgr) GetTableName() string {
	return "user_active_score"
}

// Get 获取
func (obj *_UserActiveScoreMgr) Get() (result UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserActiveScoreMgr) Gets() (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserActiveScoreMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_UserActiveScoreMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithScene scene获取 活跃度场景
func (obj *_UserActiveScoreMgr) WithScene(scene int8) Option {
	return optionFunc(func(o *options) { o.query["scene"] = scene })
}

// WithScore score获取 活跃度
func (obj *_UserActiveScoreMgr) WithScore(score int16) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithCreateAt create_at获取
func (obj *_UserActiveScoreMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserActiveScoreMgr) GetByOption(opts ...Option) (result UserActiveScore, err error) {
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
func (obj *_UserActiveScoreMgr) GetByOptions(opts ...Option) (results []*UserActiveScore, err error) {
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
func (obj *_UserActiveScoreMgr) GetFromID(id int) (result UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserActiveScoreMgr) GetBatchFromID(ids []int) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_UserActiveScoreMgr) GetFromUserID(userID int) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_UserActiveScoreMgr) GetBatchFromUserID(userIDs []int) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromScene 通过scene获取内容 活跃度场景
func (obj *_UserActiveScoreMgr) GetFromScene(scene int8) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scene = ?", scene).Find(&results).Error

	return
}

// GetBatchFromScene 批量唯一主键查找 活跃度场景
func (obj *_UserActiveScoreMgr) GetBatchFromScene(scenes []int8) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scene IN (?)", scenes).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 活跃度
func (obj *_UserActiveScoreMgr) GetFromScore(score int16) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("score = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量唯一主键查找 活跃度
func (obj *_UserActiveScoreMgr) GetBatchFromScore(scores []int16) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("score IN (?)", scores).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_UserActiveScoreMgr) GetFromCreateAt(createAt time.Time) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_UserActiveScoreMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserActiveScoreMgr) FetchByPrimaryKey(id int) (result UserActiveScore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
