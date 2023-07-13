package repository

import (
	"article/entity"
	"context"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (p *PostRepository) GetPagination(ctx context.Context, pg entity.Pagination) ([]entity.Post, error) {
	var posts []entity.Post

	err := p.db.
		WithContext(ctx).
		Offset(pg.Offset).Limit(pg.Limit).
		Find(&posts).Error
	if err != nil {
		return []entity.Post{}, err
	}

	return posts, nil
}

func (p *PostRepository) GetPaginationS(ctx context.Context, pg entity.Pagination, status entity.PostStatus) ([]entity.Post, error) {
	var posts []entity.Post

	err := p.db.
		WithContext(ctx).
		Offset(pg.Offset).Limit(pg.Limit).
		Where("status = ?", status).
		Find(&posts).Error
	if err != nil {
		return []entity.Post{}, err
	}

	return posts, nil
}

func (p *PostRepository) GetAllPost(ctx context.Context, status entity.PostStatus) ([]entity.Post, error) {
	var posts []entity.Post

	err := p.db.
		WithContext(ctx).
		Where("status = ?", status).
		Find(&posts).Error
	if err != nil {
		return []entity.Post{}, err
	}

	return posts, nil
}

func (p *PostRepository) AddPost(ctx context.Context, post *entity.Post) error {
	err := p.db.
		WithContext(ctx).
		Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostRepository) GetPostByID(ctx context.Context, id uint) (entity.Post, error) {
	var post entity.Post

	err := p.db.
		WithContext(ctx).
		Table("posts").
		Where("id = ?", id).
		First(&post).Error
	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

func (p *PostRepository) DeletePost(ctx context.Context, id int) error {
	err := p.db.
		WithContext(ctx).
		Delete(&entity.Post{}, id).Error
	return err
}

func (p *PostRepository) UpdatePost(ctx context.Context, post *entity.Post) error {
	err := p.db.
		WithContext(ctx).
		Table("posts").
		Where("id = ?", post.ID).
		Updates(&post).Error
	return err
}
