package models

import (
	"database/sql"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Count int    `json:"count" binding:"required"`
	Url   string `json:"url" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}
type FBPostData struct {
	Id            string
	Author        string
	Content       string
	CreatedAt     string
	ReactionCount int
	ResponseCount int
	Title         string
	Views         int
	Topic         string
	Url           string
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(pReq CreatePostRequest) ([]FBPostData, error) {
	// Create an array to store generated posts
	var posts []FBPostData

	// Generate the specified number of posts
	for i := 0; i < pReq.Count; i++ {
		id := uuid.New().String()
		comments := gofakeit.ProductDescription()
		post := FBPostData{
			Id:            id,
			Author:        gofakeit.Name(),
			Content:       comments,
			CreatedAt:     faker.Date(),
			ReactionCount: gofakeit.Number(1, 2000),
			ResponseCount: gofakeit.Number(1, 2000),
			Title:         gofakeit.ProductName(),
			Views:         gofakeit.Number(1, 10000),
			Topic:         pReq.Topic,
			Url:           pReq.Url,
		}
		posts = append(posts, post)
	}

	stmt := `INSERT INTO fbposts (id, author, content, created_at, response_count, reaction_count, title, views, topic, url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	for _, post := range posts {
		_, err := m.DB.Exec(stmt, post.Id, post.Author, post.Content, post.CreatedAt, post.ReactionCount, post.ResponseCount, post.Title, post.Views, post.Topic, post.Url)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}
