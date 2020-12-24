package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticleCommentMgr struct {
	*_BaseMgr
}

// ArticleCommentMgr open func
func ArticleCommentMgr(db *gorm.DB) *_ArticleCommentMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleCommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleCommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article_comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleCommentMgr) GetTableName() string {
	return "article_comment"
}

// Get 获取
func (obj *_ArticleCommentMgr) Get() (result ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleCommentMgr) Gets() (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticleCommentMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章编号
func (obj *_ArticleCommentMgr) WithArticleID(articleID uint64) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithUserID user_id获取 会员MID
func (obj *_ArticleCommentMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithParentID parent_id获取 父评论ID
func (obj *_ArticleCommentMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithNumberLike number_like获取 点赞喜欢次数
func (obj *_ArticleCommentMgr) WithNumberLike(numberLike int) Option {
	return optionFunc(func(o *options) { o.query["number_like"] = numberLike })
}

// WithNumberComment number_comment获取 评论数量
func (obj *_ArticleCommentMgr) WithNumberComment(numberComment int) Option {
	return optionFunc(func(o *options) { o.query["number_comment"] = numberComment })
}

// WithContent content获取 评论内容
func (obj *_ArticleCommentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ArticleCommentMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithStatus status获取 0不可用
func (obj *_ArticleCommentMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleCommentMgr) GetByOption(opts ...Option) (result ArticleComment, err error) {
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
func (obj *_ArticleCommentMgr) GetByOptions(opts ...Option) (results []*ArticleComment, err error) {
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
func (obj *_ArticleCommentMgr) GetFromID(id uint64) (result ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticleCommentMgr) GetBatchFromID(ids []uint64) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章编号
func (obj *_ArticleCommentMgr) GetFromArticleID(articleID uint64) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找 文章编号
func (obj *_ArticleCommentMgr) GetBatchFromArticleID(articleIDs []uint64) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 会员MID
func (obj *_ArticleCommentMgr) GetFromUserID(userID uint64) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 会员MID
func (obj *_ArticleCommentMgr) GetBatchFromUserID(userIDs []uint64) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父评论ID
func (obj *_ArticleCommentMgr) GetFromParentID(parentID int) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量唯一主键查找 父评论ID
func (obj *_ArticleCommentMgr) GetBatchFromParentID(parentIDs []int) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromNumberLike 通过number_like获取内容 点赞喜欢次数
func (obj *_ArticleCommentMgr) GetFromNumberLike(numberLike int) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_like = ?", numberLike).Find(&results).Error

	return
}

// GetBatchFromNumberLike 批量唯一主键查找 点赞喜欢次数
func (obj *_ArticleCommentMgr) GetBatchFromNumberLike(numberLikes []int) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_like IN (?)", numberLikes).Find(&results).Error

	return
}

// GetFromNumberComment 通过number_comment获取内容 评论数量
func (obj *_ArticleCommentMgr) GetFromNumberComment(numberComment int) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_comment = ?", numberComment).Find(&results).Error

	return
}

// GetBatchFromNumberComment 批量唯一主键查找 评论数量
func (obj *_ArticleCommentMgr) GetBatchFromNumberComment(numberComments []int) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_comment IN (?)", numberComments).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 评论内容
func (obj *_ArticleCommentMgr) GetFromContent(content string) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 评论内容
func (obj *_ArticleCommentMgr) GetBatchFromContent(contents []string) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ArticleCommentMgr) GetFromCreateAt(createAt time.Time) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ArticleCommentMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 0不可用
func (obj *_ArticleCommentMgr) GetFromStatus(status bool) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 0不可用
func (obj *_ArticleCommentMgr) GetBatchFromStatus(statuss []bool) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticleCommentMgr) FetchByPrimaryKey(id uint64) (result ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleCommentCid  获取多个内容
func (obj *_ArticleCommentMgr) FetchIndexByIDxDataArticleCommentCid(articleID uint64) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleCommentMid  获取多个内容
func (obj *_ArticleCommentMgr) FetchIndexByIDxDataArticleCommentMid(userID uint64) (results []*ArticleComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
