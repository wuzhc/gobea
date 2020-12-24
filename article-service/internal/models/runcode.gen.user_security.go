package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UserSecurityMgr struct {
	*_BaseMgr
}

// UserSecurityMgr open func
func UserSecurityMgr(db *gorm.DB) *_UserSecurityMgr {
	if db == nil {
		panic(fmt.Errorf("UserSecurityMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserSecurityMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_security"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserSecurityMgr) GetTableName() string {
	return "user_security"
}

// Get 获取
func (obj *_UserSecurityMgr) Get() (result UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserSecurityMgr) Gets() (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserSecurityMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_UserSecurityMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithQuestion question获取 密保问题
func (obj *_UserSecurityMgr) WithQuestion(question string) Option {
	return optionFunc(func(o *options) { o.query["question"] = question })
}

// WithAnswer answer获取 密保答案
func (obj *_UserSecurityMgr) WithAnswer(answer string) Option {
	return optionFunc(func(o *options) { o.query["answer"] = answer })
}

// GetByOption 功能选项模式获取
func (obj *_UserSecurityMgr) GetByOption(opts ...Option) (result UserSecurity, err error) {
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
func (obj *_UserSecurityMgr) GetByOptions(opts ...Option) (results []*UserSecurity, err error) {
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
func (obj *_UserSecurityMgr) GetFromID(id int) (result UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserSecurityMgr) GetBatchFromID(ids []int) (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_UserSecurityMgr) GetFromUserID(userID int) (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_UserSecurityMgr) GetBatchFromUserID(userIDs []int) (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromQuestion 通过question获取内容 密保问题
func (obj *_UserSecurityMgr) GetFromQuestion(question string) (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("question = ?", question).Find(&results).Error

	return
}

// GetBatchFromQuestion 批量唯一主键查找 密保问题
func (obj *_UserSecurityMgr) GetBatchFromQuestion(questions []string) (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("question IN (?)", questions).Find(&results).Error

	return
}

// GetFromAnswer 通过answer获取内容 密保答案
func (obj *_UserSecurityMgr) GetFromAnswer(answer string) (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("answer = ?", answer).Find(&results).Error

	return
}

// GetBatchFromAnswer 批量唯一主键查找 密保答案
func (obj *_UserSecurityMgr) GetBatchFromAnswer(answers []string) (results []*UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("answer IN (?)", answers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserSecurityMgr) FetchByPrimaryKey(id int) (result UserSecurity, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
