package model

import (
	"time"

	"go-scraper/src/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	POSTS_TN = "posts"
)

type (
	Posts struct {
		Id   primitive.ObjectID `bson:"_id"    json:"_id"`
		Post `bson:",inline"`
	}

	Post struct {
		PId        int64     `bson:"pid"`        // pid
		Title      string    `bson:"title"`      // title
		Content    string    `bson:"content"`    // content
		Categories []string  `bson:"categories"` // categories
		Topics     []string  `bson:"topics"`     // topics
		MediaURL   []string  `bson:"media_url"`  // media url
		CreatedAt  time.Time `bson:"created_at"` // create at
	}
)

/*
	new post
*/
func NewPost(pid int64, title string, created_at time.Time) *Post {
	p := new(Post)
	p.PId = pid
	p.Title = title
	p.Content = ""
	p.Categories = make([]string, 0)
	p.Topics = make([]string, 0)
	p.MediaURL = make([]string, 0)
	p.CreatedAt = created_at

	return p
}

/*
	check post is existed by pid
*/
func IsPostExistedByPId(pid int64) (exist bool, code int, err error) {
	dao := dao.GetDao()
	t := POSTS_TN
	f := bson.M{
		"pid": pid,
	}

	count := int64(0)
	if count, code, err = dao.QueryCount(t, f); err != nil {
		return
	}
	exist = count > 0

	return
}

/*
	create posts
*/
func SavePosts(rs []interface{}) (_ids []string, code int, err error) {
	dao := dao.GetDao()
	t := POSTS_TN
	_ids, code, err = dao.Insert(t, rs)
	if err != nil {
		return
	}

	return
}

/*
	delete posts
*/
func DeletePosts() (code int, err error) {
	dao := dao.GetDao()
	t := POSTS_TN
	f := bson.M{}
	if _, code, err = dao.Delete(t, f); err != nil {
		return
	}

	return
}
