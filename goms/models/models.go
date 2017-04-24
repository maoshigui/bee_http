package models

import (
	"os"
	"path"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

// Topic 文章结构
// type Topic struct {
// 	Id              int64
// 	Uid             int64
// 	Title           string
// 	Content         string    `orm:"size(5000)"`
// 	Attachment      string    // 附件
// 	Created         time.Time `orm:"index"` // 创建时间
// 	Updated         time.Time `orm:"index"`
// 	Views           int64     `orm:"index"` // 浏览次数
// 	Author          string
// 	ReplyTime       time.Time `orm:"index"` // 最后评论的时间
// 	ReplyCount      int64     // 评论数量
// 	ReplyLastUserId int64     // 最后评论的用户id
// }

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Cell))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

// func AddTopic(title, content string) error {
// 	o := orm.NewOrm()

// 	topic := &Topic{
// 		Title:     title,
// 		Content:   content,
// 		Created:   time.Now(),
// 		Updated:   time.Now(),
// 		ReplyTime: time.Now(),
// 	}

// 	_, err := o.Insert(topic)
// 	return err
// }

// func GetAllTopics(isDesc bool) ([]*Topic, error) {
// 	o := orm.NewOrm()

// 	topics := make([]*Topic, 0)

// 	qs := o.QueryTable("topic")
// 	var err error
// 	if isDesc {
// 		_, err = qs.OrderBy("-created").All(&topics)
// 	} else {
// 		_, err = qs.All(&topics)

// 	}

// 	return topics, err
// }
