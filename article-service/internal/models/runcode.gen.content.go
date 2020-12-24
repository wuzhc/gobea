package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ContentMgr struct {
	*_BaseMgr
}

// ContentMgr open func
func ContentMgr(db *gorm.DB) *_ContentMgr {
	if db == nil {
		panic(fmt.Errorf("ContentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ContentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("content"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ContentMgr) GetTableName() string {
	return "content"
}

// Get 获取
func (obj *_ContentMgr) Get() (result Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ContentMgr) Gets() (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ContentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDepth depth获取
func (obj *_ContentMgr) WithDepth(depth int) Option {
	return optionFunc(func(o *options) { o.query["depth"] = depth })
}

// WithURL url获取
func (obj *_ContentMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithArticleTitle article_title获取
func (obj *_ContentMgr) WithArticleTitle(articleTitle string) Option {
	return optionFunc(func(o *options) { o.query["article_title"] = articleTitle })
}

// WithArticleHeadimg article_headimg获取
func (obj *_ContentMgr) WithArticleHeadimg(articleHeadimg string) Option {
	return optionFunc(func(o *options) { o.query["article_headimg"] = articleHeadimg })
}

// WithArticleAuthor article_author获取
func (obj *_ContentMgr) WithArticleAuthor(articleAuthor string) Option {
	return optionFunc(func(o *options) { o.query["article_author"] = articleAuthor })
}

// WithArticleContent article_content获取
func (obj *_ContentMgr) WithArticleContent(articleContent string) Option {
	return optionFunc(func(o *options) { o.query["article_content"] = articleContent })
}

// WithArticlePublishTime article_publish_time获取
func (obj *_ContentMgr) WithArticlePublishTime(articlePublishTime int) Option {
	return optionFunc(func(o *options) { o.query["article_publish_time"] = articlePublishTime })
}

// GetByOption 功能选项模式获取
func (obj *_ContentMgr) GetByOption(opts ...Option) (result Content, err error) {
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
func (obj *_ContentMgr) GetByOptions(opts ...Option) (results []*Content, err error) {
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
func (obj *_ContentMgr) GetFromID(id int) (result Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromID(ids []int) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromDepth 通过depth获取内容
func (obj *_ContentMgr) GetFromDepth(depth int) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depth = ?", depth).Find(&results).Error

	return
}

// GetBatchFromDepth 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromDepth(depths []int) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depth IN (?)", depths).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_ContentMgr) GetFromURL(url string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromURL(urls []string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromArticleTitle 通过article_title获取内容
func (obj *_ContentMgr) GetFromArticleTitle(articleTitle string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_title = ?", articleTitle).Find(&results).Error

	return
}

// GetBatchFromArticleTitle 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromArticleTitle(articleTitles []string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_title IN (?)", articleTitles).Find(&results).Error

	return
}

// GetFromArticleHeadimg 通过article_headimg获取内容
func (obj *_ContentMgr) GetFromArticleHeadimg(articleHeadimg string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_headimg = ?", articleHeadimg).Find(&results).Error

	return
}

// GetBatchFromArticleHeadimg 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromArticleHeadimg(articleHeadimgs []string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_headimg IN (?)", articleHeadimgs).Find(&results).Error

	return
}

// GetFromArticleAuthor 通过article_author获取内容
func (obj *_ContentMgr) GetFromArticleAuthor(articleAuthor string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_author = ?", articleAuthor).Find(&results).Error

	return
}

// GetBatchFromArticleAuthor 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromArticleAuthor(articleAuthors []string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_author IN (?)", articleAuthors).Find(&results).Error

	return
}

// GetFromArticleContent 通过article_content获取内容
func (obj *_ContentMgr) GetFromArticleContent(articleContent string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_content = ?", articleContent).Find(&results).Error

	return
}

// GetBatchFromArticleContent 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromArticleContent(articleContents []string) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_content IN (?)", articleContents).Find(&results).Error

	return
}

// GetFromArticlePublishTime 通过article_publish_time获取内容
func (obj *_ContentMgr) GetFromArticlePublishTime(articlePublishTime int) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_publish_time = ?", articlePublishTime).Find(&results).Error

	return
}

// GetBatchFromArticlePublishTime 批量唯一主键查找
func (obj *_ContentMgr) GetBatchFromArticlePublishTime(articlePublishTimes []int) (results []*Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_publish_time IN (?)", articlePublishTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ContentMgr) FetchByPrimaryKey(id int) (result Content, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
