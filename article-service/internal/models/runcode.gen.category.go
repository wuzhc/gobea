package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CategoryMgr struct {
	*_BaseMgr
}

// CategoryMgr open func
func CategoryMgr(db *gorm.DB) *_CategoryMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CategoryMgr) GetTableName() string {
	return "category"
}

// Get 获取
func (obj *_CategoryMgr) Get() (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CategoryMgr) Gets() (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CategoryMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 标签名称
func (obj *_CategoryMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithRemark remark获取 标签说明
func (obj *_CategoryMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStyle style获取 样式类
func (obj *_CategoryMgr) WithStyle(style string) Option {
	return optionFunc(func(o *options) { o.query["style"] = style })
}

// WithType type获取
func (obj *_CategoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSort sort获取 排序权重
func (obj *_CategoryMgr) WithSort(sort uint64) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithStatus status获取 权限状态(1使用,0禁用)
func (obj *_CategoryMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_CategoryMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_CategoryMgr) GetByOption(opts ...Option) (result Category, err error) {
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
func (obj *_CategoryMgr) GetByOptions(opts ...Option) (results []*Category, err error) {
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
func (obj *_CategoryMgr) GetFromID(id uint64) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CategoryMgr) GetBatchFromID(ids []uint64) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标签名称
func (obj *_CategoryMgr) GetFromTitle(title string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 标签名称
func (obj *_CategoryMgr) GetBatchFromTitle(titles []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 标签说明
func (obj *_CategoryMgr) GetFromRemark(remark string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 标签说明
func (obj *_CategoryMgr) GetBatchFromRemark(remarks []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStyle 通过style获取内容 样式类
func (obj *_CategoryMgr) GetFromStyle(style string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("style = ?", style).Find(&results).Error

	return
}

// GetBatchFromStyle 批量唯一主键查找 样式类
func (obj *_CategoryMgr) GetBatchFromStyle(styles []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("style IN (?)", styles).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_CategoryMgr) GetFromType(_type string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_CategoryMgr) GetBatchFromType(_types []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序权重
func (obj *_CategoryMgr) GetFromSort(sort uint64) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序权重
func (obj *_CategoryMgr) GetBatchFromSort(sorts []uint64) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 权限状态(1使用,0禁用)
func (obj *_CategoryMgr) GetFromStatus(status bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 权限状态(1使用,0禁用)
func (obj *_CategoryMgr) GetBatchFromStatus(statuss []bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_CategoryMgr) GetFromCreateAt(createAt time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_CategoryMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CategoryMgr) FetchByPrimaryKey(id uint64) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleTagsStatus  获取多个内容
func (obj *_CategoryMgr) FetchIndexByIDxDataArticleTagsStatus(status bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}
