package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticleMgr struct {
	*_BaseMgr
}

// ArticleMgr open func
func ArticleMgr(db *gorm.DB) *_ArticleMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("article"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleMgr) GetTableName() string {
	return "article"
}

// Get 获取
func (obj *_ArticleMgr) Get() (result Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleMgr) Gets() (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticleMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 会员ID
func (obj *_ArticleMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCategoryID category_id获取
func (obj *_ArticleMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithTitle title获取 文章标题
func (obj *_ArticleMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithType type获取 文章类型(video,article,audio)
func (obj *_ArticleMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithLogo logo获取 攻略图片
func (obj *_ArticleMgr) WithLogo(logo string) Option {
	return optionFunc(func(o *options) { o.query["logo"] = logo })
}

// WithLogoFrom logo_from获取 logo来源那里，1本地，2七牛，3阿里云，4外链
func (obj *_ArticleMgr) WithLogoFrom(logoFrom bool) Option {
	return optionFunc(func(o *options) { o.query["logo_from"] = logoFrom })
}

// WithRedirectURL redirect_url获取 跳转地址
func (obj *_ArticleMgr) WithRedirectURL(redirectURL string) Option {
	return optionFunc(func(o *options) { o.query["redirect_url"] = redirectURL })
}

// WithKeywords keywords获取 seo关键字
func (obj *_ArticleMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

// WithRemark remark获取 备注说明
func (obj *_ArticleMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithContent content获取 文章内容
func (obj *_ArticleMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithNumberLikes number_likes获取 文章点赞数
func (obj *_ArticleMgr) WithNumberLikes(numberLikes int64) Option {
	return optionFunc(func(o *options) { o.query["number_likes"] = numberLikes })
}

// WithNumberReads number_reads获取 文章阅读数
func (obj *_ArticleMgr) WithNumberReads(numberReads uint64) Option {
	return optionFunc(func(o *options) { o.query["number_reads"] = numberReads })
}

// WithNumberComment number_comment获取 文章评论数
func (obj *_ArticleMgr) WithNumberComment(numberComment uint64) Option {
	return optionFunc(func(o *options) { o.query["number_comment"] = numberComment })
}

// WithNumberCollection number_collection获取 文章收藏数
func (obj *_ArticleMgr) WithNumberCollection(numberCollection uint64) Option {
	return optionFunc(func(o *options) { o.query["number_collection"] = numberCollection })
}

// WithSort sort获取 排序权重
func (obj *_ArticleMgr) WithSort(sort float32) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithStatus status获取 权限状态(0审核中，1使用，2审核不通过)
func (obj *_ArticleMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithEditorType editor_type获取 1markdown,2simditor
func (obj *_ArticleMgr) WithEditorType(editorType bool) Option {
	return optionFunc(func(o *options) { o.query["editor_type"] = editorType })
}

// WithDeleted deleted获取 删除状态(0未删,1已删)
func (obj *_ArticleMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithOriginal original获取 是否原创
func (obj *_ArticleMgr) WithOriginal(original bool) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithAuthor author获取 原作者
func (obj *_ArticleMgr) WithAuthor(author string) Option {
	return optionFunc(func(o *options) { o.query["author"] = author })
}

// WithLink link获取 转载地址
func (obj *_ArticleMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

// WithIsTop is_top获取
func (obj *_ArticleMgr) WithIsTop(isTop bool) Option {
	return optionFunc(func(o *options) { o.query["is_top"] = isTop })
}

// WithCreateAt create_at获取 创建时间
func (obj *_ArticleMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleMgr) GetByOption(opts ...Option) (result Article, err error) {
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
func (obj *_ArticleMgr) GetByOptions(opts ...Option) (results []*Article, err error) {
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
func (obj *_ArticleMgr) GetFromID(id uint64) (result Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticleMgr) GetBatchFromID(ids []uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 会员ID
func (obj *_ArticleMgr) GetFromUserID(userID int) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 会员ID
func (obj *_ArticleMgr) GetBatchFromUserID(userIDs []int) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容
func (obj *_ArticleMgr) GetFromCategoryID(categoryID int) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量唯一主键查找
func (obj *_ArticleMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 文章标题
func (obj *_ArticleMgr) GetFromTitle(title string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 文章标题
func (obj *_ArticleMgr) GetBatchFromTitle(titles []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 文章类型(video,article,audio)
func (obj *_ArticleMgr) GetFromType(_type string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 文章类型(video,article,audio)
func (obj *_ArticleMgr) GetBatchFromType(_types []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromLogo 通过logo获取内容 攻略图片
func (obj *_ArticleMgr) GetFromLogo(logo string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo = ?", logo).Find(&results).Error

	return
}

// GetBatchFromLogo 批量唯一主键查找 攻略图片
func (obj *_ArticleMgr) GetBatchFromLogo(logos []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo IN (?)", logos).Find(&results).Error

	return
}

// GetFromLogoFrom 通过logo_from获取内容 logo来源那里，1本地，2七牛，3阿里云，4外链
func (obj *_ArticleMgr) GetFromLogoFrom(logoFrom bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo_from = ?", logoFrom).Find(&results).Error

	return
}

// GetBatchFromLogoFrom 批量唯一主键查找 logo来源那里，1本地，2七牛，3阿里云，4外链
func (obj *_ArticleMgr) GetBatchFromLogoFrom(logoFroms []bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logo_from IN (?)", logoFroms).Find(&results).Error

	return
}

// GetFromRedirectURL 通过redirect_url获取内容 跳转地址
func (obj *_ArticleMgr) GetFromRedirectURL(redirectURL string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_url = ?", redirectURL).Find(&results).Error

	return
}

// GetBatchFromRedirectURL 批量唯一主键查找 跳转地址
func (obj *_ArticleMgr) GetBatchFromRedirectURL(redirectURLs []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_url IN (?)", redirectURLs).Find(&results).Error

	return
}

// GetFromKeywords 通过keywords获取内容 seo关键字
func (obj *_ArticleMgr) GetFromKeywords(keywords string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords = ?", keywords).Find(&results).Error

	return
}

// GetBatchFromKeywords 批量唯一主键查找 seo关键字
func (obj *_ArticleMgr) GetBatchFromKeywords(keywordss []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords IN (?)", keywordss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注说明
func (obj *_ArticleMgr) GetFromRemark(remark string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注说明
func (obj *_ArticleMgr) GetBatchFromRemark(remarks []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 文章内容
func (obj *_ArticleMgr) GetFromContent(content string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 文章内容
func (obj *_ArticleMgr) GetBatchFromContent(contents []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromNumberLikes 通过number_likes获取内容 文章点赞数
func (obj *_ArticleMgr) GetFromNumberLikes(numberLikes int64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_likes = ?", numberLikes).Find(&results).Error

	return
}

// GetBatchFromNumberLikes 批量唯一主键查找 文章点赞数
func (obj *_ArticleMgr) GetBatchFromNumberLikes(numberLikess []int64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_likes IN (?)", numberLikess).Find(&results).Error

	return
}

// GetFromNumberReads 通过number_reads获取内容 文章阅读数
func (obj *_ArticleMgr) GetFromNumberReads(numberReads uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_reads = ?", numberReads).Find(&results).Error

	return
}

// GetBatchFromNumberReads 批量唯一主键查找 文章阅读数
func (obj *_ArticleMgr) GetBatchFromNumberReads(numberReadss []uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_reads IN (?)", numberReadss).Find(&results).Error

	return
}

// GetFromNumberComment 通过number_comment获取内容 文章评论数
func (obj *_ArticleMgr) GetFromNumberComment(numberComment uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_comment = ?", numberComment).Find(&results).Error

	return
}

// GetBatchFromNumberComment 批量唯一主键查找 文章评论数
func (obj *_ArticleMgr) GetBatchFromNumberComment(numberComments []uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_comment IN (?)", numberComments).Find(&results).Error

	return
}

// GetFromNumberCollection 通过number_collection获取内容 文章收藏数
func (obj *_ArticleMgr) GetFromNumberCollection(numberCollection uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_collection = ?", numberCollection).Find(&results).Error

	return
}

// GetBatchFromNumberCollection 批量唯一主键查找 文章收藏数
func (obj *_ArticleMgr) GetBatchFromNumberCollection(numberCollections []uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number_collection IN (?)", numberCollections).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序权重
func (obj *_ArticleMgr) GetFromSort(sort float32) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序权重
func (obj *_ArticleMgr) GetBatchFromSort(sorts []float32) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 权限状态(0审核中，1使用，2审核不通过)
func (obj *_ArticleMgr) GetFromStatus(status bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 权限状态(0审核中，1使用，2审核不通过)
func (obj *_ArticleMgr) GetBatchFromStatus(statuss []bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromEditorType 通过editor_type获取内容 1markdown,2simditor
func (obj *_ArticleMgr) GetFromEditorType(editorType bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editor_type = ?", editorType).Find(&results).Error

	return
}

// GetBatchFromEditorType 批量唯一主键查找 1markdown,2simditor
func (obj *_ArticleMgr) GetBatchFromEditorType(editorTypes []bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editor_type IN (?)", editorTypes).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 删除状态(0未删,1已删)
func (obj *_ArticleMgr) GetFromDeleted(deleted bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找 删除状态(0未删,1已删)
func (obj *_ArticleMgr) GetBatchFromDeleted(deleteds []bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 是否原创
func (obj *_ArticleMgr) GetFromOriginal(original bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量唯一主键查找 是否原创
func (obj *_ArticleMgr) GetBatchFromOriginal(originals []bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original IN (?)", originals).Find(&results).Error

	return
}

// GetFromAuthor 通过author获取内容 原作者
func (obj *_ArticleMgr) GetFromAuthor(author string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("author = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量唯一主键查找 原作者
func (obj *_ArticleMgr) GetBatchFromAuthor(authors []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("author IN (?)", authors).Find(&results).Error

	return
}

// GetFromLink 通过link获取内容 转载地址
func (obj *_ArticleMgr) GetFromLink(link string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

// GetBatchFromLink 批量唯一主键查找 转载地址
func (obj *_ArticleMgr) GetBatchFromLink(links []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}

// GetFromIsTop 通过is_top获取内容
func (obj *_ArticleMgr) GetFromIsTop(isTop bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_top = ?", isTop).Find(&results).Error

	return
}

// GetBatchFromIsTop 批量唯一主键查找
func (obj *_ArticleMgr) GetBatchFromIsTop(isTops []bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_top IN (?)", isTops).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_ArticleMgr) GetFromCreateAt(createAt time.Time) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_ArticleMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticleMgr) FetchByPrimaryKey(id uint64) (result Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxDataArticleContentType  获取多个内容
func (obj *_ArticleMgr) FetchIndexByIDxDataArticleContentType(_type string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleContentStatus  获取多个内容
func (obj *_ArticleMgr) FetchIndexByIDxDataArticleContentStatus(status bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// FetchIndexByIDxDataArticleContentDeleted  获取多个内容
func (obj *_ArticleMgr) FetchIndexByIDxDataArticleContentDeleted(deleted bool) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}
