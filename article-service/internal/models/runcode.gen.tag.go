package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TagMgr struct {
	*_BaseMgr
}

// TagMgr open func
func TagMgr(db *gorm.DB) *_TagMgr {
	if db == nil {
		panic(fmt.Errorf("TagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TagMgr) GetTableName() string {
	return "tag"
}

// Get 获取
func (obj *_TagMgr) Get() (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagMgr) Gets() (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TagMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_TagMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithTitle title获取 标签名称
func (obj *_TagMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithRemark remark获取
func (obj *_TagMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithThumb thumb获取
func (obj *_TagMgr) WithThumb(thumb string) Option {
	return optionFunc(func(o *options) { o.query["thumb"] = thumb })
}

// WithSort sort获取 排序权重，值越大排越前
func (obj *_TagMgr) WithSort(sort int16) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithStatus status获取 权限状态(1使用,0禁用)
func (obj *_TagMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_TagMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_TagMgr) GetByOption(opts ...Option) (result Tag, err error) {
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
func (obj *_TagMgr) GetByOptions(opts ...Option) (results []*Tag, err error) {
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
func (obj *_TagMgr) GetFromID(id uint64) (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromID(ids []uint64) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_TagMgr) GetFromUserID(userID int) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromUserID(userIDs []int) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标签名称
func (obj *_TagMgr) GetFromTitle(title string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 标签名称
func (obj *_TagMgr) GetBatchFromTitle(titles []string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容
func (obj *_TagMgr) GetFromRemark(remark string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromRemark(remarks []string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromThumb 通过thumb获取内容
func (obj *_TagMgr) GetFromThumb(thumb string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("thumb = ?", thumb).Find(&results).Error

	return
}

// GetBatchFromThumb 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromThumb(thumbs []string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("thumb IN (?)", thumbs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序权重，值越大排越前
func (obj *_TagMgr) GetFromSort(sort int16) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序权重，值越大排越前
func (obj *_TagMgr) GetBatchFromSort(sorts []int16) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 权限状态(1使用,0禁用)
func (obj *_TagMgr) GetFromStatus(status bool) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 权限状态(1使用,0禁用)
func (obj *_TagMgr) GetBatchFromStatus(statuss []bool) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_TagMgr) GetFromCreateAt(createAt time.Time) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_TagMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TagMgr) FetchByPrimaryKey(id uint64) (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleTagsStatus  获取多个内容
func (obj *_TagMgr) FetchIndexByIDxDataArticleTagsStatus(status bool) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}
