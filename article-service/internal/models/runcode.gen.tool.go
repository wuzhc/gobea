package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ToolMgr struct {
	*_BaseMgr
}

// ToolMgr open func
func ToolMgr(db *gorm.DB) *_ToolMgr {
	if db == nil {
		panic(fmt.Errorf("ToolMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ToolMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tool"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ToolMgr) GetTableName() string {
	return "tool"
}

// Get 获取
func (obj *_ToolMgr) Get() (result Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ToolMgr) Gets() (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ToolMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 1在线工具，2运维工具，3调试工具，4设计工具
func (obj *_ToolMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithUserID user_id获取
func (obj *_ToolMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithTitle title获取
func (obj *_ToolMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithLogo logo获取
func (obj *_ToolMgr) WithLogo(logo string) Option {
	return optionFunc(func(o *options) { o.query["logo"] = logo })
}

// WithURL url获取
func (obj *_ToolMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithRemark remark获取
func (obj *_ToolMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateAt create_at获取
func (obj *_ToolMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithStatus status获取
func (obj *_ToolMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithScope scope获取 是否公开，1公开，2私人
func (obj *_ToolMgr) WithScope(scope int8) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithSort sort获取 排序，越大排越前
func (obj *_ToolMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *_ToolMgr) GetByOption(opts ...Option) (result Tool, err error) {
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
func (obj *_ToolMgr) GetByOptions(opts ...Option) (results []*Tool, err error) {
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
func (obj *_ToolMgr) GetFromID(id int) (result Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromID(ids []int) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 1在线工具，2运维工具，3调试工具，4设计工具
func (obj *_ToolMgr) GetFromType(_type int8) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 1在线工具，2运维工具，3调试工具，4设计工具
func (obj *_ToolMgr) GetBatchFromType(_types []int8) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_ToolMgr) GetFromUserID(userID int) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromUserID(userIDs []int) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_ToolMgr) GetFromTitle(title string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromTitle(titles []string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromLogo 通过logo获取内容
func (obj *_ToolMgr) GetFromLogo(logo string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo = ?", logo).Find(&results).Error

	return
}

// GetBatchFromLogo 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromLogo(logos []string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo IN (?)", logos).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_ToolMgr) GetFromURL(url string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromURL(urls []string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容
func (obj *_ToolMgr) GetFromRemark(remark string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromRemark(remarks []string) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_ToolMgr) GetFromCreateAt(createAt time.Time) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ToolMgr) GetFromStatus(status int8) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ToolMgr) GetBatchFromStatus(statuss []int8) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容 是否公开，1公开，2私人
func (obj *_ToolMgr) GetFromScope(scope int8) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scope = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量唯一主键查找 是否公开，1公开，2私人
func (obj *_ToolMgr) GetBatchFromScope(scopes []int8) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scope IN (?)", scopes).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序，越大排越前
func (obj *_ToolMgr) GetFromSort(sort int) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序，越大排越前
func (obj *_ToolMgr) GetBatchFromSort(sorts []int) (results []*Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ToolMgr) FetchByPrimaryKey(id int) (result Tool, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
