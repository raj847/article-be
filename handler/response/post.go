package response

import "article/entity"

type Post struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

func BuildPost(post entity.Post) Post {
	return Post{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Status:   string(post.Status),
	}
}

func BuildPosts(posts []entity.Post) []Post {
	res := []Post{}
	for _, v := range posts {
		post := BuildPost(v)
		res = append(res, post)
	}
	return res
}
