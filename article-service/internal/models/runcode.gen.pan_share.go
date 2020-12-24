package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PanShareMgr struct {
	*_BaseMgr
}

// PanShareMgr open func
func PanShareMgr(db *gorm.DB) *_PanShareMgr {
	if db == nil {
		panic(fmt.Errorf("PanShareMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PanShareMgr{_BaseMgr: &_BaseMgr{DB: db.Table("pan_share"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PanShareMgr) GetTableName() string {
	return "pan_share"
}

// Get 获取
func (obj *_PanShareMgr) Get() (result PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PanShareMgr) Gets() (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PanShareMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取
func (obj *_PanShareMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithShareLink share_link获取
func (obj *_PanShareMgr) WithShareLink(shareLink string) Option {
	return optionFunc(func(o *options) { o.query["share_link"] = shareLink })
}

// WithSharePwd share_pwd获取
func (obj *_PanShareMgr) WithSharePwd(sharePwd string) Option {
	return optionFunc(func(o *options) { o.query["share_pwd"] = sharePwd })
}

// WithStatus status获取
func (obj *_PanShareMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取
func (obj *_PanShareMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithNumberDownload number_download获取
func (obj *_PanShareMgr) WithNumberDownload(numberDownload int) Option {
	return optionFunc(func(o *options) { o.query["number_download"] = numberDownload })
}

// GetByOption 功能选项模式获取
func (obj *_PanShareMgr) GetByOption(opts ...Option) (result PanShare, err error) {
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
func (obj *_PanShareMgr) GetByOptions(opts ...Option) (results []*PanShare, err error) {
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
func (obj *_PanShareMgr) GetFromID(id int) (result PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_PanShareMgr) GetBatchFromID(ids []int) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_PanShareMgr) GetFromTitle(title string) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_PanShareMgr) GetBatchFromTitle(titles []string) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromShareLink 通过share_link获取内容
func (obj *_PanShareMgr) GetFromShareLink(shareLink string) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_link = ?", shareLink).Find(&results).Error

	return
}

// GetBatchFromShareLink 批量唯一主键查找
func (obj *_PanShareMgr) GetBatchFromShareLink(shareLinks []string) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_link IN (?)", shareLinks).Find(&results).Error

	return
}

// GetFromSharePwd 通过share_pwd获取内容
func (obj *_PanShareMgr) GetFromSharePwd(sharePwd string) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_pwd = ?", sharePwd).Find(&results).Error

	return
}

// GetBatchFromSharePwd 批量唯一主键查找
func (obj *_PanShareMgr) GetBatchFromSharePwd(sharePwds []string) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_pwd IN (?)", sharePwds).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_PanShareMgr) GetFromStatus(status bool) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_PanShareMgr) GetBatchFromStatus(statuss []bool) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_PanShareMgr) GetFromCreateAt(createAt time.Time) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_PanShareMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromNumberDownload 通过number_download获取内容
func (obj *_PanShareMgr) GetFromNumberDownload(numberDownload int) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_download = ?", numberDownload).Find(&results).Error

	return
}

// GetBatchFromNumberDownload 批量唯一主键查找
func (obj *_PanShareMgr) GetBatchFromNumberDownload(numberDownloads []int) (results []*PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_download IN (?)", numberDownloads).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PanShareMgr) FetchByPrimaryKey(id int) (result PanShare, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
