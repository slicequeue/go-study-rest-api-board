package handler

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/slicequeue/go-study-rest-api-board/model"
	"github.com/slicequeue/go-study-rest-api-board/utils"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

// mapping to user model
func (r *userRegisterRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	u.Password = r.User.Password
	return nil
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *SignInRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type BoardDocumentRequest struct {
	Id      string `json:"id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content"`
}

func (r *BoardDocumentRequest) bind(c echo.Context, d *model.Document) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	boardId := utils.GetParamValue(c, "boardId", "integer", nil).(int)
	if boardId == 0 {
		return errors.New("boardId must not be null.")
	}
	userId := c.Get("user")
	if userId == nil {
		return errors.New("need siginin.")
	}
	d.BoardID = uint(boardId)
	d.Title = r.Title
	d.Content = r.Content
	d.Author.UserId = userId.(uint)

	return nil
}
