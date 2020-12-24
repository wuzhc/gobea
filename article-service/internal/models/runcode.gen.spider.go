package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SpiderMgr struct {
	*_BaseMgr
}

// SpiderMgr open func
func SpiderMgr(db *gorm.DB) *_SpiderMgr {
	if db == nil {
		panic(fmt.Errorf("SpiderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpiderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("spider"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SpiderMgr) GetTableName() string {
	return "spider"
}

// Get 获取
func (obj *_SpiderMgr) Get() (result Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SpiderMgr) Gets() (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SpiderMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFrom from获取
func (obj *_SpiderMgr) WithFrom(from string) Option {
	return optionFunc(func(o *options) { o.query["from"] = from })
}

// WithTitle title获取
func (obj *_SpiderMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithWeight weight获取
func (obj *_SpiderMgr) WithWeight(weight int) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithAuthor author获取
func (obj *_SpiderMgr) WithAuthor(author string) Option {
	return optionFunc(func(o *options) { o.query["author"] = author })
}

// WithRemark remark获取
func (obj *_SpiderMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithURL url获取
func (obj *_SpiderMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithStatus status获取 0未同步，1已同步，2同步失败
func (obj *_SpiderMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTipMsg tip_msg获取
func (obj *_SpiderMgr) WithTipMsg(tipMsg string) Option {
	return optionFunc(func(o *options) { o.query["tip_msg"] = tipMsg })
}

// WithCreateAt create_at获取
func (obj *_SpiderMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SpiderMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_SpiderMgr) GetByOption(opts ...Option) (result Spider, err error) {
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
func (obj *_SpiderMgr) GetByOptions(opts ...Option) (results []*Spider, err error) {
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
func (obj *_SpiderMgr) GetFromID(id int64) (result Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromID(ids []int64) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromFrom 通过from获取内容
func (obj *_SpiderMgr) GetFromFrom(from string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from = ?", from).Find(&results).Error

	return
}

// GetBatchFromFrom 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromFrom(froms []string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from IN (?)", froms).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_SpiderMgr) GetFromTitle(title string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromTitle(titles []string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容
func (obj *_SpiderMgr) GetFromWeight(weight int) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromWeight(weights []int) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

// GetFromAuthor 通过author获取内容
func (obj *_SpiderMgr) GetFromAuthor(author string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("author = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromAuthor(authors []string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("author IN (?)", authors).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容
func (obj *_SpiderMgr) GetFromRemark(remark string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromRemark(remarks []string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_SpiderMgr) GetFromURL(url string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromURL(urls []string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 0未同步，1已同步，2同步失败
func (obj *_SpiderMgr) GetFromStatus(status int8) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 0未同步，1已同步，2同步失败
func (obj *_SpiderMgr) GetBatchFromStatus(statuss []int8) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTipMsg 通过tip_msg获取内容
func (obj *_SpiderMgr) GetFromTipMsg(tipMsg string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tip_msg = ?", tipMsg).Find(&results).Error

	return
}

// GetBatchFromTipMsg 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromTipMsg(tipMsgs []string) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tip_msg IN (?)", tipMsgs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_SpiderMgr) GetFromCreateAt(createAt time.Time) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SpiderMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找
func (obj *_SpiderMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SpiderMgr) FetchByPrimaryKey(id int64) (result Spider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
