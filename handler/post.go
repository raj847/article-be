package handler

import (
	"article/entity"
	"article/handler/request"
	"article/handler/response"
	"article/service"
	"article/validate"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(
	postService *service.PostService,
) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (m *PostHandler) Create(c echo.Context) error {

	var postReq request.Post
	err := c.Bind(&postReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	if err = validate.Validate(&postReq); err != nil {
		errRes, ok := response.HttpValidationErrr(err)
		if !ok {
			return err
		}

		return c.JSON(http.StatusBadRequest, errRes)
	}

	post := entity.Post{
		Title:    postReq.Title,
		Content:  postReq.Content,
		Category: postReq.Category,
		Status:   entity.PostStatus(postReq.Status),
	}
	err = m.postService.AddPost(c.Request().Context(), &post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create post",
					Code:    "POST_CREATE-ERROR",
				},
			},
		})
	}
	res := response.BuildPost(post)
	return c.JSON(http.StatusCreated, res)
}

func (m *PostHandler) GetAllPagination(c echo.Context) error {
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	status := c.QueryParam("status")
	if status != "" {
		posts, err := m.postService.GetAllPostPaginationS(c.Request().Context(), entity.Pagination{
			Offset: offset,
			Limit:  limit,
		}, entity.PostStatus(status))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to read all post",
						Code:    "POST_READ-ALL-ERROR",
					},
				},
			})
		}
		res := response.BuildPosts(posts)
		return c.JSON(http.StatusOK, res)
	}
	posts, err := m.postService.GetAllPostPagination(c.Request().Context(), entity.Pagination{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all post",
					Code:    "POST_READ-ALL-ERROR",
				},
			},
		})
	}
	res := response.BuildPosts(posts)
	return c.JSON(http.StatusOK, res)
}

func (m *PostHandler) GetAll(c echo.Context) error {
	rawstatus := c.QueryParam("status")
	status := entity.PostStatus(rawstatus)
	if status == entity.Unknown {
		status = entity.Publish
	}
	posts, err := m.postService.GetAll(c.Request().Context(), status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all post",
					Code:    "POST_READ-ALL-ERROR",
				},
			},
		})
	}
	res := response.BuildPosts(posts)
	return c.JSON(http.StatusOK, res)
}

func (m *PostHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	posts, err := m.postService.GetPostByID(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to read post",
						Code:    "POST_NOT-FOUND-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read post by id",
					Code:    "POST_READ-BY-ID-ERROR",
				},
			},
		})
	}
	res := response.BuildPost(posts)
	return c.JSON(http.StatusOK, res)
}

func (m *PostHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := m.postService.DeletePost(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to delete post",
					Code:    "POST_DELETE-ERROR",
				},
			},
		})
	}
	return c.JSON(http.StatusOK, response.SuccessMessage{
		Message: "Post has been deleted",
	})
}

func (m *PostHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var postReq request.Post
	err := c.Bind(&postReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	if err = validate.Validate(&postReq); err != nil {
		errRes, ok := response.HttpValidationErrr(err)
		if !ok {
			return err
		}

		return c.JSON(http.StatusBadRequest, errRes)
	}
	res, err := m.postService.UpdatePost(c.Request().Context(), &entity.Post{
		ID:       uint(id),
		Title:    postReq.Title,
		Content:  postReq.Content,
		Category: postReq.Category,
		Status:   entity.PostStatus(postReq.Status),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to update post",
					Code:    "POST_UPDATE-ERROR",
				},
			},
		})
	}
	result := response.BuildPost(res)
	return c.JSON(http.StatusOK, result)
}

func (m *PostHandler) UpdateTrash(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	posts, err := m.postService.UpdatePostTrash(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to patch post",
						Code:    "POST_NOT-FOUND-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to patch post by id",
					Code:    "POST_PATCH-BY-ID-ERROR",
				},
			},
		})
	}
	res := response.BuildPost(posts)
	return c.JSON(http.StatusOK, res)
}
