package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 0普通用户，1超级用户，2虚拟用户
func (obj *_UserMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithUsername username获取 用户账号
func (obj *_UserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithAccount account获取
func (obj *_UserMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithPassword password获取 用户密码
func (obj *_UserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithHeadimg headimg获取
func (obj *_UserMgr) WithHeadimg(headimg string) Option {
	return optionFunc(func(o *options) { o.query["headimg"] = headimg })
}

// WithQq qq获取 联系QQ
func (obj *_UserMgr) WithQq(qq string) Option {
	return optionFunc(func(o *options) { o.query["qq"] = qq })
}

// WithEmail email获取 联系邮箱
func (obj *_UserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPhone phone获取 联系手机
func (obj *_UserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithActiveScore active_score获取 活跃度
func (obj *_UserMgr) WithActiveScore(activeScore int) Option {
	return optionFunc(func(o *options) { o.query["active_score"] = activeScore })
}

// WithLoginAt login_at获取 登录时间
func (obj *_UserMgr) WithLoginAt(loginAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["login_at"] = loginAt })
}

// WithLoginIP login_ip获取 登录IP
func (obj *_UserMgr) WithLoginIP(loginIP string) Option {
	return optionFunc(func(o *options) { o.query["login_ip"] = loginIP })
}

// WithLoginNum login_num获取 登录次数
func (obj *_UserMgr) WithLoginNum(loginNum uint64) Option {
	return optionFunc(func(o *options) { o.query["login_num"] = loginNum })
}

// WithAuthorize authorize获取 权限授权
func (obj *_UserMgr) WithAuthorize(authorize string) Option {
	return optionFunc(func(o *options) { o.query["authorize"] = authorize })
}

// WithTags tags获取 用户标签
func (obj *_UserMgr) WithTags(tags string) Option {
	return optionFunc(func(o *options) { o.query["tags"] = tags })
}

// WithDesc desc获取 备注说明
func (obj *_UserMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithStatus status获取 状态(0禁用,1启用)
func (obj *_UserMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsDeleted is_deleted获取 删除(1删除,0未删)
func (obj *_UserMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCreateAt create_at获取 创建时间
func (obj *_UserMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
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
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
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
func (obj *_UserMgr) GetFromID(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromID(ids []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 0普通用户，1超级用户，2虚拟用户
func (obj *_UserMgr) GetFromType(_type int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 0普通用户，1超级用户，2虚拟用户
func (obj *_UserMgr) GetBatchFromType(_types []int8) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户账号
func (obj *_UserMgr) GetFromUsername(username string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找 用户账号
func (obj *_UserMgr) GetBatchFromUsername(usernames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容
func (obj *_UserMgr) GetFromAccount(account string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account = ?", account).Find(&result).Error

	return
}

// GetBatchFromAccount 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromAccount(accounts []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account IN (?)", accounts).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 用户密码
func (obj *_UserMgr) GetFromPassword(password string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找 用户密码
func (obj *_UserMgr) GetBatchFromPassword(passwords []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromHeadimg 通过headimg获取内容
func (obj *_UserMgr) GetFromHeadimg(headimg string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("headimg = ?", headimg).Find(&results).Error

	return
}

// GetBatchFromHeadimg 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromHeadimg(headimgs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("headimg IN (?)", headimgs).Find(&results).Error

	return
}

// GetFromQq 通过qq获取内容 联系QQ
func (obj *_UserMgr) GetFromQq(qq string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qq = ?", qq).Find(&results).Error

	return
}

// GetBatchFromQq 批量唯一主键查找 联系QQ
func (obj *_UserMgr) GetBatchFromQq(qqs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qq IN (?)", qqs).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 联系邮箱
func (obj *_UserMgr) GetFromEmail(email string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 联系邮箱
func (obj *_UserMgr) GetBatchFromEmail(emails []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 联系手机
func (obj *_UserMgr) GetFromPhone(phone string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 联系手机
func (obj *_UserMgr) GetBatchFromPhone(phones []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromActiveScore 通过active_score获取内容 活跃度
func (obj *_UserMgr) GetFromActiveScore(activeScore int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active_score = ?", activeScore).Find(&results).Error

	return
}

// GetBatchFromActiveScore 批量唯一主键查找 活跃度
func (obj *_UserMgr) GetBatchFromActiveScore(activeScores []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active_score IN (?)", activeScores).Find(&results).Error

	return
}

// GetFromLoginAt 通过login_at获取内容 登录时间
func (obj *_UserMgr) GetFromLoginAt(loginAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_at = ?", loginAt).Find(&results).Error

	return
}

// GetBatchFromLoginAt 批量唯一主键查找 登录时间
func (obj *_UserMgr) GetBatchFromLoginAt(loginAts []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_at IN (?)", loginAts).Find(&results).Error

	return
}

// GetFromLoginIP 通过login_ip获取内容 登录IP
func (obj *_UserMgr) GetFromLoginIP(loginIP string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_ip = ?", loginIP).Find(&results).Error

	return
}

// GetBatchFromLoginIP 批量唯一主键查找 登录IP
func (obj *_UserMgr) GetBatchFromLoginIP(loginIPs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_ip IN (?)", loginIPs).Find(&results).Error

	return
}

// GetFromLoginNum 通过login_num获取内容 登录次数
func (obj *_UserMgr) GetFromLoginNum(loginNum uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_num = ?", loginNum).Find(&results).Error

	return
}

// GetBatchFromLoginNum 批量唯一主键查找 登录次数
func (obj *_UserMgr) GetBatchFromLoginNum(loginNums []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_num IN (?)", loginNums).Find(&results).Error

	return
}

// GetFromAuthorize 通过authorize获取内容 权限授权
func (obj *_UserMgr) GetFromAuthorize(authorize string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("authorize = ?", authorize).Find(&results).Error

	return
}

// GetBatchFromAuthorize 批量唯一主键查找 权限授权
func (obj *_UserMgr) GetBatchFromAuthorize(authorizes []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("authorize IN (?)", authorizes).Find(&results).Error

	return
}

// GetFromTags 通过tags获取内容 用户标签
func (obj *_UserMgr) GetFromTags(tags string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tags = ?", tags).Find(&results).Error

	return
}

// GetBatchFromTags 批量唯一主键查找 用户标签
func (obj *_UserMgr) GetBatchFromTags(tagss []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tags IN (?)", tagss).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 备注说明
func (obj *_UserMgr) GetFromDesc(desc string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 备注说明
func (obj *_UserMgr) GetBatchFromDesc(descs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态(0禁用,1启用)
func (obj *_UserMgr) GetFromStatus(status bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态(0禁用,1启用)
func (obj *_UserMgr) GetBatchFromStatus(statuss []bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容 删除(1删除,0未删)
func (obj *_UserMgr) GetFromIsDeleted(isDeleted bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找 删除(1删除,0未删)
func (obj *_UserMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_UserMgr) GetFromCreateAt(createAt time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_UserMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByAccount primay or index 获取唯一内容
func (obj *_UserMgr) FetchUniqueByAccount(account string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("account = ?", account).Find(&result).Error

	return
}
