// Package controller has actions to handle
// the correct logic
package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/felipehfs/blog/model"
	"github.com/gorilla/mux"
)

// BlogContext represents the database type
type BlogContext string

// CreatePost handler
func CreatePost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post model.Post
		db, ok := r.Context().Value(BlogContext("database")).(*model.Database)
		if !ok {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			log.Println("Database not found")
			return
		}

		dao := model.NewPostRepository(db)
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		post.ID = bson.NewObjectIdWithTime(time.Now())
		post.CreatedAt = time.Now()

		if err := dao.Create(post); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	})
}

// ReadPost returns all post data
func ReadPost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, ok := r.Context().Value(BlogContext("database")).(*model.Database)
		if !ok {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			log.Println("Database not found")
			return
		}
		dao := model.NewPostRepository(db)
		posts, err := dao.Read()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	})
}

// UpdatePost updates the data
func UpdatePost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post model.Post
		params := mux.Vars(r)
		db, ok := r.Context().Value(BlogContext("database")).(*model.Database)

		if !ok {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			log.Println("Database not found")
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		dao := model.NewPostRepository(db)
		if err := dao.Update(params["id"], post); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	})
}

// RemovePost will remove the post by id
func RemovePost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, ok := r.Context().Value(BlogContext("database")).(*model.Database)
		if !ok {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			log.Println("Database not found")
			return
		}
		params := mux.Vars(r)
		dao := model.NewPostRepository(db)
		if err := dao.Remove(params["id"]); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

// FindByIDPost retrieves a post
func FindByIDPost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, ok := r.Context().Value(BlogContext("database")).(*model.Database)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInsufficientStorage), http.StatusInternalServerError)
			log.Println("database not found")
			return
		}

		params := mux.Vars(r)
		dao := model.NewPostRepository(db)
		post, err := dao.FindByID(params["id"])

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	})
}
