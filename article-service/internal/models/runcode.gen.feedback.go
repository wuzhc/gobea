package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FeedbackMgr struct {
	*_BaseMgr
}

// FeedbackMgr open func
func FeedbackMgr(db *gorm.DB) *_FeedbackMgr {
	if db == nil {
		panic(fmt.Errorf("FeedbackMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FeedbackMgr{_BaseMgr: &_BaseMgr{DB: db.Table("feedback"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FeedbackMgr) GetTableName() string {
	return "feedback"
}

// Get 获取
func (obj *_FeedbackMgr) Get() (result Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FeedbackMgr) Gets() (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FeedbackMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContent content获取
func (obj *_FeedbackMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateTime create_time获取
func (obj *_FeedbackMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUserID user_id获取
func (obj *_FeedbackMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUa ua获取
func (obj *_FeedbackMgr) WithUa(ua string) Option {
	return optionFunc(func(o *options) { o.query["ua"] = ua })
}

// WithToken token获取
func (obj *_FeedbackMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// GetByOption 功能选项模式获取
func (obj *_FeedbackMgr) GetByOption(opts ...Option) (result Feedback, err error) {
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
func (obj *_FeedbackMgr) GetByOptions(opts ...Option) (results []*Feedback, err error) {
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
func (obj *_FeedbackMgr) GetFromID(id int) (result Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_FeedbackMgr) GetBatchFromID(ids []int) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_FeedbackMgr) GetFromContent(content string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_FeedbackMgr) GetBatchFromContent(contents []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_FeedbackMgr) GetFromCreateTime(createTime time.Time) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找
func (obj *_FeedbackMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_FeedbackMgr) GetFromUserID(userID int) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_FeedbackMgr) GetBatchFromUserID(userIDs []int) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUa 通过ua获取内容
func (obj *_FeedbackMgr) GetFromUa(ua string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ua = ?", ua).Find(&results).Error

	return
}

// GetBatchFromUa 批量唯一主键查找
func (obj *_FeedbackMgr) GetBatchFromUa(uas []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ua IN (?)", uas).Find(&results).Error

	return
}

// GetFromToken 通过token获取内容
func (obj *_FeedbackMgr) GetFromToken(token string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error

	return
}

// GetBatchFromToken 批量唯一主键查找
func (obj *_FeedbackMgr) GetBatchFromToken(tokens []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_FeedbackMgr) FetchByPrimaryKey(id int) (result Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
