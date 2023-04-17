package route

import (
	"Praktikum/constant"
	"Praktikum/controller"
	"Praktikum/repository/database"
	"Praktikum/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {

	userRepository := database.NewUserRepository(db)
	blogRepository := database.NewBlogRepository(db)
	bookRepository := database.NewBookRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, blogRepository)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	bookUsecase := usecase.NewBookUsecase(bookRepository)

	authControlller := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase, userRepository)
	blogController := controller.NewBlogController(blogUsecase, blogRepository)
	bookController := controller.NewBookController(bookUsecase, bookRepository)

	e.Validator = &customValidator{validator: validator.New()}

	e.POST("/login", authControlller.LoginUserController)
	e.POST("/users", userController.CreateUserController)

	// Route / to handler function/ User
	user := e.Group("/users", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("", userController.GetUsersController)
	user.GET("/:id", userController.GetUserController)
	user.DELETE("/:id", userController.DeleteUserController)
	user.PUT("/:id", userController.UpdateUserController)

	//Book
	book := e.Group("/books", middleware.JWT([]byte(constant.SECRET_JWT)))
	book.GET("/:id", bookController.GetBookController)
	book.GET("", bookController.GetBooksController)
	book.POST("", bookController.CreateBookController)
	book.DELETE("/:id", bookController.DeleteBookController)
	book.PUT("/:id", bookController.UpdateBookController)

	//Blog
	blog := e.Group("/blogs", middleware.JWT([]byte(constant.SECRET_JWT)))
	blog.GET("/:id", blogController.GetBlogController)
	blog.GET("", blogController.GetBlogsController)
	blog.POST("", blogController.CreateBlogController)
	blog.DELETE("/:id", blogController.DeleteBlogController)
	blog.PUT("/:id", blogController.UpdateBlogController)

}
