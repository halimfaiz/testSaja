package controller

import (
	"Praktikum/model"
	"Praktikum/repository/database"
	"Praktikum/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogController interface {
	GetBlogsController(c echo.Context) error
	GetBlogController(c echo.Context) error
	CreateBlogController(c echo.Context) error
	DeleteBlogController(c echo.Context) error
	UpdateBlogController(c echo.Context) error
}

type blogController struct {
	blogUsecase    usecase.BlogUsecase
	blogRepository database.BlogRepository
}

func NewBlogController(
	blogUsecase usecase.BlogUsecase,
	blogRepository database.BlogRepository,
) *blogController {
	return &blogController{
		blogUsecase,
		blogRepository,
	}
}

// get all blogs
func (b *blogController) GetBlogsController(c echo.Context) error {
	blogs, e := b.blogRepository.GetBlogs()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func (b *blogController) GetBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	blog, e := b.blogRepository.GetBlog(uint(id))

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get blog",
			"errorDescription": e,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blogs":   blog,
	})
}

// create new blog
func (b *blogController) CreateBlogController(c echo.Context) error {
	blog := model.Blog{}
	c.Bind(&blog)

	if err := b.blogUsecase.CreateBlog(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create blog",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete blog by id
func (b *blogController) DeleteBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.blogUsecase.DeleteBlog(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete blog",
			"errorDescription": err,
			"errorMessage":     "Blog tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Delete blog",
	})
}

// update blog by id
func (b *blogController) UpdateBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	blog := model.Blog{}
	c.Bind(&blog)
	blog.ID = uint(id)

	if err := b.blogUsecase.UpdateBlog(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update blog",
			"errorDescription": err,
			"errorMessage":     "Blog tidak dapat di ubah",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}
