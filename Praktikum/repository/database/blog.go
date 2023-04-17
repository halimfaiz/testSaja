package database

import (
	"Praktikum/config"
	"Praktikum/model"

	"gorm.io/gorm"
)

type BlogRepository interface {
	CreateBlog(blog *model.Blog) error
	GetBlogs() (blogs []model.Blog, err error)
	GetBlog(id uint) (blog model.Blog, err error)
	GetBlogByUserId(userID uint) (blog model.Blog, err error)
	UpdateBlog(blog *model.Blog) error
	DeleteBlog(blog *model.Blog) error
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{}
}

func (b *blogRepository) NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db}
}

func (b *blogRepository) CreateBlog(blog *model.Blog) error {

	if err := config.DB.Create(blog).Error; err != nil {
		return err
	}
	return nil
}

func (b *blogRepository) GetBlogs() (blogs []model.Blog, err error) {

	if err = config.DB.Find(&blogs).Error; err != nil {
		return
	}
	return
}

func (b *blogRepository) GetBlog(id uint) (blog model.Blog, err error) {
	blog.ID = id
	if err = config.DB.First(&blog).Error; err != nil {
		return
	}
	return
}

func (b *blogRepository) GetBlogByUserId(userID uint) (blog model.Blog, err error) {
	blog.UserId = userID
	if err = config.DB.Find(&blog).Error; err != nil {
		return
	}
	return
}

func (b *blogRepository) UpdateBlog(blog *model.Blog) error {

	if err := config.DB.Updates(blog).Error; err != nil {
		return err
	}
	return nil
}

func (b *blogRepository) DeleteBlog(blog *model.Blog) error {

	if e := config.DB.Delete(blog).Error; e != nil {
		return e
	}
	return nil
}
