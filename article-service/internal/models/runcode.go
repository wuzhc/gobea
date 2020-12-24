package models

import (
	"time"
)

// Article 数据-文章-内容
type Article struct {
	ID               uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	UserID           int       `gorm:"column:user_id;type:int(11);not null"` // 会员ID
	CategoryID       int       `gorm:"column:category_id;type:int(11);not null"`
	Title            string    `gorm:"column:title;type:varchar(100)"`                                               // 文章标题
	Type             string    `gorm:"index:idx_data_article_content_type;column:type;type:varchar(20)"`             // 文章类型(video,article,audio)
	Logo             string    `gorm:"column:logo;type:varchar(500)"`                                                // 攻略图片
	LogoFrom         bool      `gorm:"column:logo_from;type:tinyint(1)"`                                             // logo来源那里，1本地，2七牛，3阿里云，4外链
	RedirectURL      string    `gorm:"column:redirect_url;type:varchar(900)"`                                        // 跳转地址
	Keywords         string    `gorm:"column:keywords;type:varchar(100)"`                                            // seo关键字
	Remark           string    `gorm:"column:remark;type:varchar(500)"`                                              // 备注说明
	Content          string    `gorm:"column:content;type:longtext"`                                                 // 文章内容
	NumberLikes      uint64    `gorm:"column:number_likes;type:bigint(20) unsigned"`                                 // 文章点赞数
	NumberReads      int64     `gorm:"column:number_reads;type:bigint(20) unsigned"`                                 // 文章阅读数
	NumberComment    uint64    `gorm:"column:number_comment;type:bigint(20) unsigned"`                               // 文章评论数
	NumberCollection int64     `gorm:"column:number_collection;type:bigint(20) unsigned"`                            // 文章收藏数
	Sort             float32   `gorm:"column:sort;type:float unsigned"`                                              // 排序权重
	Status           int       `gorm:"index:idx_data_article_content_status;column:status;type:tinyint(1) unsigned"` // 权限状态(0审核中，1使用，2审核不通过)
	EditorType       int       `gorm:"column:editor_type;type:tinyint(1);not null"`                                  // 1markdown,2simditor
	Deleted          int       `gorm:"index:idx_data_article_content_deleted;column:deleted;type:tinyint(1)"`        // 删除状态(0未删,1已删)
	Original         int       `gorm:"column:original;type:tinyint(1);not null"`                                     // 是否原创
	Author           string    `gorm:"column:author;type:varchar(50)"`                                               // 原作者
	Link             string    `gorm:"column:link;type:varchar(255)"`                                                // 转载地址
	IsTop            int       `gorm:"column:is_top;type:tinyint(1);not null"`
	CreateAt         time.Time `gorm:"column:create_at;type:timestamp"` // 创建时间
}

// ArticleCollection 数据-文章-收藏
type ArticleCollection struct {
	ID        uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ArticleID uint64    `gorm:"index:idx_data_article_collection_cid;column:article_id;type:bigint(20) unsigned"` // 文章编号
	UserID    int64     `gorm:"index:idx_data_article_collection_mid;column:user_id;type:bigint(20) unsigned"`    // 会员MID
	Status    bool      `gorm:"column:status;type:tinyint(1);not null"`
	CreateAt  time.Time `gorm:"column:create_at;type:timestamp"` // 创建时间
}

// ArticleComment 数据-文章-评论
type ArticleComment struct {
	ID            uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ArticleID     uint64    `gorm:"index:idx_data_article_comment_cid;column:article_id;type:bigint(20) unsigned"` // 文章编号
	UserID        uint64    `gorm:"index:idx_data_article_comment_mid;column:user_id;type:bigint(20) unsigned"`    // 会员MID
	ParentID      int       `gorm:"column:parent_id;type:int(11);not null"`                                        // 父评论ID
	NumberLike    int       `gorm:"column:number_like;type:int(11);not null"`                                      // 点赞喜欢次数
	NumberComment int       `gorm:"column:number_comment;type:int(11);not null"`                                   // 评论数量
	Content       string    `gorm:"column:content;type:varchar(500)"`                                              // 评论内容
	CreateAt      time.Time `gorm:"column:create_at;type:timestamp"`                                               // 创建时间
	Status        bool      `gorm:"column:status;type:tinyint(1);not null"`                                        // 0不可用
}

// ArticleCommentLike 数据-文章-点赞
type ArticleCommentLike struct {
	ID        uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ArticleID int       `gorm:"column:article_id;type:int(11)"`
	CommentID int64     `gorm:"index:idx_data_article_like_cid;column:comment_id;type:bigint(20) unsigned;not null"` // 评论编号
	UserID    int64     `gorm:"index:idx_data_article_like_mid;column:user_id;type:bigint(20) unsigned"`             // 会员MID
	CreateAt  time.Time `gorm:"column:create_at;type:timestamp"`                                                     // 创建时间
	Status    bool      `gorm:"column:status;type:tinyint(1);not null"`
}

// ArticleHistory 数据-文章-历史
type ArticleHistory struct {
	ID        uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ArticleID uint64    `gorm:"column:article_id;type:bigint(20) unsigned"` // 文章编号
	UserID    uint64    `gorm:"column:user_id;type:bigint(20) unsigned"`    // 会员MID
	CreateAt  time.Time `gorm:"column:create_at;type:timestamp"`            // 创建时间
}

// ArticleLike 数据-文章-点赞
type ArticleLike struct {
	ID        uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ArticleID uint64    `gorm:"index:idx_data_article_like_cid;column:article_id;type:bigint(20) unsigned"` // 文章编号
	UserID    uint64    `gorm:"index:idx_data_article_like_mid;column:user_id;type:bigint(20) unsigned"`    // 会员MID
	Status    bool      `gorm:"column:status;type:tinyint(1);not null"`
	CreateAt  time.Time `gorm:"column:create_at;type:timestamp"` // 创建时间
}

// ArticlePanShare 数据-文章-网盘
type ArticlePanShare struct {
	ID         uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ArticleID  uint64    `gorm:"index:idx_data_article_like_cid;column:article_id;type:bigint(20) unsigned"`   // 文章编号
	PanShareID int64     `gorm:"index:idx_data_article_like_mid;column:pan_share_id;type:bigint(20) unsigned"` // 分享ID
	Status     bool      `gorm:"column:status;type:tinyint(1);not null"`
	CreateAt   time.Time `gorm:"column:create_at;type:timestamp"` // 创建时间
}

// ArticleTag 数据-文章-标签
type ArticleTag struct {
	ID        uint64 `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ArticleID uint64 `gorm:"index:idx_data_article_like_cid;column:article_id;type:bigint(20) unsigned;not null"` // 文章编号
	TagID     uint64 `gorm:"index:idx_data_article_like_mid;column:tag_id;type:bigint(20) unsigned"`              // 标签ID
	Status    bool   `gorm:"column:status;type:tinyint(1);not null"`
}

// Category 数据-文章-标签
type Category struct {
	ID       int64     `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	Title    string    `gorm:"column:title;type:varchar(100)"`         // 标签名称
	Remark   string    `gorm:"column:remark;type:varchar(500)"`        // 标签说明
	Style    string    `gorm:"column:style;type:varchar(50);not null"` // 样式类
	Type     string    `gorm:"column:type;type:varchar(15);not null"`
	Sort     uint64    `gorm:"column:sort;type:bigint(20) unsigned"`                                      // 排序权重
	Status   bool      `gorm:"index:idx_data_article_tags_status;column:status;type:tinyint(1) unsigned"` // 权限状态(1使用,0禁用)
	CreateAt time.Time `gorm:"column:create_at;type:timestamp"`                                           // 创建时间
}

// Content [...]
type Content struct {
	ID                 uint32 `gorm:"primaryKey;column:id;type:int(11) unsigned;not null"`
	Depth              int    `gorm:"column:depth;type:int(11)"`
	URL                string `gorm:"column:url;type:varchar(200)"`
	ArticleTitle       string `gorm:"column:article_title;type:varchar(20)"`
	ArticleHeadimg     string `gorm:"column:article_headimg;type:varchar(150)"`
	ArticleAuthor      string `gorm:"column:article_author;type:varchar(20)"`
	ArticleContent     string `gorm:"column:article_content;type:text"`
	ArticlePublishTime int    `gorm:"column:article_publish_time;type:int(10)"`
}

// Ebook [...]
type Ebook struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	Title     string `gorm:"column:title;type:varchar(255);not null"`
	ArticleID int    `gorm:"column:article_id;type:int(11);not null"`
	ParentID  int    `gorm:"column:parent_id;type:int(11);not null"`
	Level     int8   `gorm:"column:level;type:tinyint(2);not null"`
	Path      string `gorm:"column:path;type:varchar(255)"`
	Anchor    string `gorm:"column:anchor;type:varchar(50)"`
}

// EbookContent [...]
type EbookContent struct {
	ID      int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	EbookID int    `gorm:"column:ebook_id;type:int(11);not null"`
	Content string `gorm:"column:content;type:longtext;not null"`
}

// Feedback [...]
type Feedback struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Content    string    `gorm:"column:content;type:text;not null"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null"`
	UserID     int       `gorm:"column:user_id;type:int(11)"`
	Ua         string    `gorm:"column:ua;type:varchar(255);not null"`
	Token      string    `gorm:"column:token;type:char(32);not null"`
}

// PanShare [...]
type PanShare struct {
	ID             int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Title          string    `gorm:"column:title;type:varchar(255);not null"`
	ShareLink      string    `gorm:"column:share_link;type:varchar(255);not null"`
	SharePwd       string    `gorm:"column:share_pwd;type:varchar(20);not null"`
	Status         bool      `gorm:"column:status;type:tinyint(1);not null"`
	CreateAt       time.Time `gorm:"column:create_at;type:datetime;not null"`
	NumberDownload int       `gorm:"column:number_download;type:int(11);not null"`
}

// Spider [...]
type Spider struct {
	ID        int64     `gorm:"primaryKey;column:id;type:bigint(11);not null"`
	From      string    `gorm:"column:from;type:varchar(20);not null"`
	Title     string    `gorm:"column:title;type:varchar(100);not null"`
	Weight    int       `gorm:"column:weight;type:int(11);not null"`
	Author    string    `gorm:"column:author;type:varchar(50)"`
	Remark    string    `gorm:"column:remark;type:varchar(500);not null"`
	URL       string    `gorm:"column:url;type:varchar(300);not null"`
	Status    int8      `gorm:"column:status;type:tinyint(2)"` // 0未同步，1已同步，2同步失败
	TipMsg    string    `gorm:"column:tip_msg;type:varchar(500)"`
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null"`
}

// Tag 标签
type Tag struct {
	ID       uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	UserID   int       `gorm:"column:user_id;type:int(11);not null"`
	Title    string    `gorm:"column:title;type:varchar(100)"` // 标签名称
	Remark   string    `gorm:"column:remark;type:varchar(255)"`
	Thumb    string    `gorm:"column:thumb;type:varchar(255)"`
	Sort     int16     `gorm:"column:sort;type:smallint(5)"`                                              // 排序权重，值越大排越前
	Status   bool      `gorm:"index:idx_data_article_tags_status;column:status;type:tinyint(1) unsigned"` // 权限状态(1使用,0禁用)
	CreateAt time.Time `gorm:"column:create_at;type:timestamp"`                                           // 创建时间
}

// TagCategory 标签 - 分类
type TagCategory struct {
	ID         uint64 `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	CategoryID uint32 `gorm:"index:idx_data_article_like_cid;column:category_id;type:int(11) unsigned"` // 分类ID
	TagID      uint64 `gorm:"index:idx_data_article_like_mid;column:tag_id;type:bigint(20) unsigned"`   // 标签ID
	Status     bool   `gorm:"column:status;type:tinyint(1);not null"`
}

// Tool [...]
type Tool struct {
	ID       int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Type     int8      `gorm:"column:type;type:tinyint(2);not null"` // 1在线工具，2运维工具，3调试工具，4设计工具
	UserID   int       `gorm:"column:user_id;type:int(11);not null"`
	Title    string    `gorm:"column:title;type:varchar(255);not null"`
	Logo     string    `gorm:"column:logo;type:varchar(255);not null"`
	URL      string    `gorm:"column:url;type:varchar(255);not null"`
	Remark   string    `gorm:"column:remark;type:varchar(255);not null"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null"`
	Status   int8      `gorm:"column:status;type:tinyint(2);not null"`
	Scope    int8      `gorm:"column:scope;type:tinyint(2);not null"` // 是否公开，1公开，2私人
	Sort     int       `gorm:"column:sort;type:int(11);not null"`     // 排序，越大排越前
}

// ToolCollection 数据-工具-收藏
type ToolCollection struct {
	ID       uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	ToolID   uint64    `gorm:"index:idx_data_article_like_cid;column:tool_id;type:bigint(20) unsigned"` // 工具编号
	UserID   uint64    `gorm:"index:idx_data_article_like_mid;column:user_id;type:bigint(20) unsigned"` // 会员MID
	Status   bool      `gorm:"column:status;type:tinyint(1);not null"`
	CreateAt time.Time `gorm:"column:create_at;type:timestamp"` // 创建时间
}

// User 系统-用户
type User struct {
	ID          uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	Type        int8      `gorm:"column:type;type:tinyint(2);not null"` // 0普通用户，1超级用户，2虚拟用户
	Username    string    `gorm:"column:username;type:varchar(50)"`     // 用户账号
	Account     string    `gorm:"unique;column:account;type:varchar(50);not null"`
	Password    string    `gorm:"column:password;type:varchar(32)"` // 用户密码
	Headimg     string    `gorm:"column:headimg;type:varchar(255)"`
	Qq          string    `gorm:"column:qq;type:varchar(16)"`                 // 联系QQ
	Email       string    `gorm:"column:email;type:varchar(32)"`              // 联系邮箱
	Phone       string    `gorm:"column:phone;type:varchar(16)"`              // 联系手机
	ActiveScore int       `gorm:"column:active_score;type:int(11);not null"`  // 活跃度
	LoginAt     time.Time `gorm:"column:login_at;type:datetime"`              // 登录时间
	LoginIP     string    `gorm:"column:login_ip;type:varchar(255)"`          // 登录IP
	LoginNum    uint64    `gorm:"column:login_num;type:bigint(20) unsigned"`  // 登录次数
	Authorize   string    `gorm:"column:authorize;type:varchar(255)"`         // 权限授权
	Tags        string    `gorm:"column:tags;type:varchar(255)"`              // 用户标签
	Desc        string    `gorm:"column:desc;type:varchar(255)"`              // 备注说明
	Status      bool      `gorm:"column:status;type:tinyint(1) unsigned"`     // 状态(0禁用,1启用)
	IsDeleted   bool      `gorm:"column:is_deleted;type:tinyint(1) unsigned"` // 删除(1删除,0未删)
	CreateAt    time.Time `gorm:"column:create_at;type:timestamp"`            // 创建时间
}

// UserActiveScore 用户活跃度
type UserActiveScore struct {
	ID       int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	UserID   int       `gorm:"column:user_id;type:int(11);not null"`
	Scene    int8      `gorm:"column:scene;type:tinyint(2);not null"`  // 活跃度场景
	Score    int16     `gorm:"column:score;type:smallint(4);not null"` // 活跃度
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null"`
}

// UserMessage 用户 - 消息表
type UserMessage struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Type       int8      `gorm:"column:type;type:tinyint(2);not null"`      // 1系统，2留言，3评论
	FromUserID int       `gorm:"column:from_user_id;type:int(11);not null"` // 发送者ID
	ToUserID   int       `gorm:"column:to_user_id;type:int(11);not null"`   // 接受者ID
	Content    string    `gorm:"column:content;type:varchar(255);not null"` // 消息内容
	Params     string    `gorm:"column:params;type:varchar(255)"`           // 评论参数，json格式
	ObjectID   int       `gorm:"column:object_id;type:int(11);not null"`    // 对象ID，可以article_id,user_id,tool_id
	ObjectType string    `gorm:"column:object_type;type:varchar(50)"`
	URL        string    `gorm:"column:url;type:varchar(255)"`
	Status     int8      `gorm:"column:status;type:tinyint(2);not null"` // 0未读，1已读
	CreateAt   time.Time `gorm:"column:create_at;type:datetime;not null"`
}

// UserSecurity [...]
type UserSecurity struct {
	ID       int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	UserID   int    `gorm:"column:user_id;type:int(11);not null"`
	Question string `gorm:"column:question;type:varchar(255);not null"` // 密保问题
	Answer   string `gorm:"column:answer;type:varchar(255);not null"`   // 密保答案
}
