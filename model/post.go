// Package model contains all bussiness logic
package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Post represents each content of the blog
type Post struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	CreatedAt   time.Time     `json:"created,omitempty"`
}

type postRepository struct {
	DB *Database
}

// NewPostRepository returns a instance for manage the post
func NewPostRepository(c *Database) *postRepository {
	return &postRepository{DB: c}
}

func (pr postRepository) getCollection() *mgo.Collection {
	return pr.DB.GetCollection("blog", "posts")
}

// Create inserts the new post
func (pr postRepository) Create(p Post) error {
	return pr.getCollection().Insert(p)
}

// Read retrieves all posts saved
func (pr postRepository) Read() ([]Post, error) {
	var posts []Post
	if err := pr.getCollection().Find(bson.M{}).All(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}

// FindByID search the post by ID
func (pr postRepository) FindByID(id string) (*Post, error) {
	var post Post
	err := pr.getCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// Update change the data into database
func (pr postRepository) Update(id string, p Post) error {
	return pr.getCollection().Update(bson.M{"_id": bson.ObjectIdHex(id)}, p)
}

// Remove drop the posts
func (pr postRepository) Remove(id string) error {
	return pr.getCollection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
