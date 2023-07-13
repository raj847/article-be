package service

import (
	"article/entity"
	"article/repository"
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (p *PostService) GetAllPostPagination(ctx context.Context, pg entity.Pagination) ([]entity.Post, error) {
	res, err := p.postRepo.GetPagination(ctx, pg)
	if err != nil {
		return []entity.Post{}, err
	}
	return res, nil
}

func (p *PostService) GetAllPostPaginationS(ctx context.Context, pg entity.Pagination, status entity.PostStatus) ([]entity.Post, error) {
	res, err := p.postRepo.GetPaginationS(ctx, pg, status)
	if err != nil {
		return []entity.Post{}, err
	}
	return res, nil
}

func (p *PostService) GetAll(ctx context.Context, ps entity.PostStatus) ([]entity.Post, error) {
	res, err := p.postRepo.GetAllPost(ctx, ps)
	if err != nil {
		return []entity.Post{}, err
	}
	return res, nil
}

func (p *PostService) AddPost(ctx context.Context, post *entity.Post) error {
	err := p.postRepo.AddPost(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostService) GetPostByID(ctx context.Context, id uint) (entity.Post, error) {
	res, err := p.postRepo.GetPostByID(ctx, id)
	if err != nil {
		return entity.Post{}, err
	}
	if res.ID == 0 {
		return entity.Post{}, ErrNotFound
	}

	return res, nil
}

func (p *PostService) UpdatePost(ctx context.Context, post *entity.Post) (entity.Post, error) {
	err := p.postRepo.UpdatePost(ctx, post)
	if err != nil {
		return entity.Post{}, err
	}

	return *post, nil
}

func (p *PostService) UpdatePostTrash(ctx context.Context, id uint) (entity.Post, error) {
	post, err := p.postRepo.GetPostByID(ctx, id)
	if err != nil {
		return entity.Post{}, err
	}
	post.Status = entity.Trash
	err = p.postRepo.UpdatePost(ctx, &post)
	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

func (p *PostService) DeletePost(ctx context.Context, id int) error {
	err := p.postRepo.DeletePost(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
