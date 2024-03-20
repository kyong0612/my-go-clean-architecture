package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/kyong0612/my-go-clean-architecture/domain"
	"github.com/kyong0612/my-go-clean-architecture/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

var _ ArticleService = (*usecase.Service)(nil)

// ResponseError represent the response error struct.
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleService represent the article's usecases
//
//go:generate mockery --name ArticleService
type ArticleService interface {
	FetchArticles(ctx context.Context, cursor string, num int32) ([]*domain.Article, string, error)
	GetArticleByID(ctx context.Context, id int32) (*domain.Article, error)
	UpdateArticle(ctx context.Context, ar *domain.Article) error
	GetArticleByTitle(ctx context.Context, title string) (*domain.Article, error)
	StoreArticle(context.Context, *domain.Article) error
	DeleteArticle(ctx context.Context, id int32) error
}

// ArticleHandler  represent the httphandler for article.
type ArticleHandler struct {
	Service ArticleService
}

const defaultNum = 10

// NewArticleHandler will initialize the articles/ resources endpoint.
func NewArticleHandler(e *echo.Echo, svc ArticleService) {
	handler := &ArticleHandler{
		Service: svc,
	}
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}

// FetchArticle will fetch the article based on given params.
func (a *ArticleHandler) FetchArticle(c echo.Context) error {
	numS := c.QueryParam("num")
	num, err := strconv.Atoi(numS)
	if err != nil || num == 0 {
		num = defaultNum
	}

	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	listAr, nextCursor, err := a.Service.FetchArticles(ctx, cursor, int32(num))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

// GetByID will get article by given id.
func (a *ArticleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int32(idP)
	ctx := c.Request().Context()

	art, err := a.Service.GetArticleByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *domain.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body.
func (a *ArticleHandler) Store(c echo.Context) (err error) {
	var article domain.Article
	err = c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.Service.StoreArticle(ctx, &article)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, article)
}

// Delete will delete article by given param.
func (a *ArticleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int32(idP)
	ctx := c.Request().Context()

	err = a.Service.DeleteArticle(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
