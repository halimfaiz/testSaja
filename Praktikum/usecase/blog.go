package usecase

import (
	"Praktikum/model"
	"Praktikum/repository/database"
	"errors"
	"fmt"
)

type BlogUsecase interface {
	CreateBlog(blog *model.Blog) error
	GetBlog(id uint) (blog model.Blog, err error)
	GetListBlogs() (blogs []model.Blog, err error)
	UpdateBlog(blog *model.Blog) (err error)
	DeleteBlog(id uint) (err error)
}

type blogUsecase struct {
	blogRepository database.BlogRepository
}

func NewBlogUsecase(blogRepo database.BlogRepository) *blogUsecase {
	return &blogUsecase{blogRepository: blogRepo}
}

func (b *blogUsecase) CreateBlog(blog *model.Blog) error {

	if blog.Judul == "" {
		return errors.New("Judul tidak boleh kosong")
	}
	if blog.Konten == "" {
		return errors.New("Konten tidak boleh kosong")
	}

	err := b.blogRepository.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}

func (b *blogUsecase) GetBlog(id uint) (blog model.Blog, err error) {
	blog, err = b.blogRepository.GetBlog(id)
	if err != nil {
		fmt.Println("GetBlog :Error Getting blog from repository")
		return
	}
	return
}

func (b *blogUsecase) GetListBlogs() (blogs []model.Blog, err error) {
	blogs, err = b.blogRepository.GetBlogs()
	if err != nil {
		fmt.Println("GetListBlogs : Error Getting blog from repository")
		return
	}
	return
}

func (b *blogUsecase) UpdateBlog(blog *model.Blog) (err error) {
	err = b.blogRepository.UpdateBlog(blog)
	if err != nil {
		fmt.Println("UpdateBlog : Error updating Blog, err: ", err)
		return
	}
	return
}

func (b *blogUsecase) DeleteBlog(id uint) (err error) {
	blog := model.Blog{}
	blog.ID = id
	err = b.blogRepository.DeleteBlog(&blog)
	if err != nil {
		fmt.Println("DeleteBlog : Error deleting Blog, err: ", err)
		return
	}
	return
}
