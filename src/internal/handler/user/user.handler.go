package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"net/http"
	"sgcu65-backend-assignment/src/internal/domain/dto"
	"sgcu65-backend-assignment/src/internal/domain/entity"
	"sgcu65-backend-assignment/src/internal/repository"
)

type Handler struct {
	userRepo repository.IUserRepository
}

func NewHandler(userRepo repository.IUserRepository) *Handler {
	return &Handler{userRepo: userRepo}
}

func (h *Handler) FindAllUsers(c *fiber.Ctx) {
	query := &dto.FindUsersQueryParams{}

	if err := c.QueryParser(query); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid query params",
			Data:       nil,
		})
		return
	}

	users := &[]*entity.User{}
	err := h.userRepo.FindUserByKeyword(users, query.Firstname, query.Surname, query.Position)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	var usersDto []dto.UserDto
	if len(*users) != 0 {
		for _, user := range *users {
			usersDto = append(usersDto, *user.EntityToDto())
		}
	}

	c.Status(http.StatusOK).JSON(usersDto)
}

func (h *Handler) CreateUser(c *fiber.Ctx) {
	createUserDto := &dto.UserDto{}
	if err := c.BodyParser(createUserDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       nil,
		})
		return
	}

	user := &entity.User{
		Email:     createUserDto.Email,
		Firstname: createUserDto.Firstname,
		Surname:   createUserDto.Surname,
		Role:      createUserDto.Role,
		Position:  createUserDto.Position,
		Salary:    createUserDto.Salary,
	}

	err := h.userRepo.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			c.Status(http.StatusConflict).JSON(&dto.ResponseErr{
				StatusCode: http.StatusConflict,
				Message:    "Email has already been used",
				Data:       nil,
			})
			return
		}

		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	c.Status(http.StatusCreated).JSON(user)
}

func (h *Handler) FindUserById(c *fiber.Ctx) {
	userId := c.Params("userId")
	if _, err := uuid.Parse(userId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user id",
			Data:       nil,
		})
		return
	}

	user := &entity.User{}
	err := h.userRepo.FindUserById(userId, user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound).JSON(&dto.ResponseErr{
				StatusCode: http.StatusNotFound,
				Message:    fmt.Sprintf("User id %s not found", userId),
				Data:       nil,
			})
			return
		}
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	c.Status(http.StatusOK).JSON(user.EntityToDto())
}

func (h *Handler) UpdateUser(c *fiber.Ctx) {
	userId := c.Params("userId")
	if _, err := uuid.Parse(userId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user id",
			Data:       nil,
		})
		return
	}

	updateUserDto := &dto.UpdateUserDto{}
	if err := c.BodyParser(updateUserDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       nil,
		})
		return
	}

	user := &entity.User{}
	err := h.userRepo.FindUserById(userId, user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound).JSON(&dto.ResponseErr{
				StatusCode: http.StatusNotFound,
				Message:    fmt.Sprintf("User id %s not found", userId),
				Data:       nil,
			})
			return
		}
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	user.Firstname = updateUserDto.Firstname
	user.Surname = updateUserDto.Surname
	user.Position = updateUserDto.Position
	user.Salary = updateUserDto.Salary

	err = h.userRepo.UpdateUser(userId, user)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	c.Status(http.StatusOK).JSON(user.EntityToDto())
}

func (h *Handler) DeleteUser(c *fiber.Ctx) {
	userId := c.Params("userId")
	if _, err := uuid.Parse(userId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user id",
			Data:       nil,
		})
		return
	}

	err := h.userRepo.DeleteUser(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound).JSON(&dto.ResponseErr{
				StatusCode: http.StatusNotFound,
				Message:    fmt.Sprintf("User id %s not found", userId),
				Data:       nil,
			})
			return
		}
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	c.Status(http.StatusOK)
}
