package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserMessageMgr struct {
	*_BaseMgr
}

// UserMessageMgr open func
func UserMessageMgr(db *gorm.DB) *_UserMessageMgr {
	if db == nil {
		panic(fmt.Errorf("UserMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMessageMgr) GetTableName() string {
	return "user_message"
}

// Get 获取
func (obj *_UserMessageMgr) Get() (result UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMessageMgr) Gets() (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserMessageMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 1系统，2留言，3评论
func (obj *_UserMessageMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithFromUserID from_user_id获取 发送者ID
func (obj *_UserMessageMgr) WithFromUserID(fromUserID int) Option {
	return optionFunc(func(o *options) { o.query["from_user_id"] = fromUserID })
}

// WithToUserID to_user_id获取 接受者ID
func (obj *_UserMessageMgr) WithToUserID(toUserID int) Option {
	return optionFunc(func(o *options) { o.query["to_user_id"] = toUserID })
}

// WithContent content获取 消息内容
func (obj *_UserMessageMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithParams params获取 评论参数，json格式
func (obj *_UserMessageMgr) WithParams(params string) Option {
	return optionFunc(func(o *options) { o.query["params"] = params })
}

// WithObjectID object_id获取 对象ID，可以article_id,user_id,tool_id
func (obj *_UserMessageMgr) WithObjectID(objectID int) Option {
	return optionFunc(func(o *options) { o.query["object_id"] = objectID })
}

// WithObjectType object_type获取
func (obj *_UserMessageMgr) WithObjectType(objectType string) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithURL url获取
func (obj *_UserMessageMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithStatus status获取 0未读，1已读
func (obj *_UserMessageMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取
func (obj *_UserMessageMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserMessageMgr) GetByOption(opts ...Option) (result UserMessage, err error) {
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
func (obj *_UserMessageMgr) GetByOptions(opts ...Option) (results []*UserMessage, err error) {
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
func (obj *_UserMessageMgr) GetFromID(id int) (result UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserMessageMgr) GetBatchFromID(ids []int) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 1系统，2留言，3评论
func (obj *_UserMessageMgr) GetFromType(_type int8) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 1系统，2留言，3评论
func (obj *_UserMessageMgr) GetBatchFromType(_types []int8) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromFromUserID 通过from_user_id获取内容 发送者ID
func (obj *_UserMessageMgr) GetFromFromUserID(fromUserID int) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_user_id = ?", fromUserID).Find(&results).Error

	return
}

// GetBatchFromFromUserID 批量唯一主键查找 发送者ID
func (obj *_UserMessageMgr) GetBatchFromFromUserID(fromUserIDs []int) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_user_id IN (?)", fromUserIDs).Find(&results).Error

	return
}

// GetFromToUserID 通过to_user_id获取内容 接受者ID
func (obj *_UserMessageMgr) GetFromToUserID(toUserID int) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to_user_id = ?", toUserID).Find(&results).Error

	return
}

// GetBatchFromToUserID 批量唯一主键查找 接受者ID
func (obj *_UserMessageMgr) GetBatchFromToUserID(toUserIDs []int) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to_user_id IN (?)", toUserIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 消息内容
func (obj *_UserMessageMgr) GetFromContent(content string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 消息内容
func (obj *_UserMessageMgr) GetBatchFromContent(contents []string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromParams 通过params获取内容 评论参数，json格式
func (obj *_UserMessageMgr) GetFromParams(params string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("params = ?", params).Find(&results).Error

	return
}

// GetBatchFromParams 批量唯一主键查找 评论参数，json格式
func (obj *_UserMessageMgr) GetBatchFromParams(paramss []string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("params IN (?)", paramss).Find(&results).Error

	return
}

// GetFromObjectID 通过object_id获取内容 对象ID，可以article_id,user_id,tool_id
func (obj *_UserMessageMgr) GetFromObjectID(objectID int) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_id = ?", objectID).Find(&results).Error

	return
}

// GetBatchFromObjectID 批量唯一主键查找 对象ID，可以article_id,user_id,tool_id
func (obj *_UserMessageMgr) GetBatchFromObjectID(objectIDs []int) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_id IN (?)", objectIDs).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_UserMessageMgr) GetFromObjectType(objectType string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_type = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量唯一主键查找
func (obj *_UserMessageMgr) GetBatchFromObjectType(objectTypes []string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_type IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_UserMessageMgr) GetFromURL(url string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找
func (obj *_UserMessageMgr) GetBatchFromURL(urls []string) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 0未读，1已读
func (obj *_UserMessageMgr) GetFromStatus(status int8) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 0未读，1已读
func (obj *_UserMessageMgr) GetBatchFromStatus(statuss []int8) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_UserMessageMgr) GetFromCreateAt(createAt time.Time) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_UserMessageMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserMessageMgr) FetchByPrimaryKey(id int) (result UserMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
