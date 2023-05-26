package task

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"sgcu65-backend-assignment/src/internal/domain/dto"
	"sgcu65-backend-assignment/src/internal/domain/entity"
	"sgcu65-backend-assignment/src/internal/repository"
	"time"
)

type Handler struct {
	taskRepo     repository.ITaskRepository
	userTaskRepo repository.IUserTaskRepository
}

func NewHandler(taskRepo repository.ITaskRepository, userTaskRepo repository.IUserTaskRepository) *Handler {
	return &Handler{taskRepo: taskRepo, userTaskRepo: userTaskRepo}
}

func (h *Handler) CreateTask(c *fiber.Ctx) {
	taskDto := &dto.TaskDto{}
	if err := c.BodyParser(taskDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
		})
		return
	}
	deadlineTime, _ := time.Parse(time.RFC3339, taskDto.Deadline)
	task := &entity.Task{
		Name:     taskDto.Name,
		Content:  taskDto.Content,
		Status:   taskDto.Status,
		Deadline: deadlineTime,
	}

	err := h.taskRepo.CreateTask(task)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	c.Status(http.StatusCreated).JSON(task.EntityToDto())
}

func (h *Handler) FindAllTask(c *fiber.Ctx) {
	query := &dto.FindTaskQueryParams{}
	if err := c.QueryParser(query); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid input",
			Data:       nil,
		})
		return
	}

	tasks := &[]*entity.Task{}
	err := h.taskRepo.FindTaskByName(tasks, query.Name)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	var taskDto []dto.TaskDto
	for _, task := range *tasks {
		taskDto = append(taskDto, *task.EntityToDto())
	}
	c.Status(http.StatusOK).JSON(taskDto)
}

func (h *Handler) FindTaskById(c *fiber.Ctx) {
	taskId := c.Params("taskId")
	if _, err := uuid.Parse(taskId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user id",
			Data:       nil,
		})
		return
	}

	task := &entity.Task{}
	err := h.taskRepo.FindTaskById(taskId, task)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound).JSON(&dto.ResponseErr{
				StatusCode: http.StatusNotFound,
				Message:    fmt.Sprintf("Task id %s not found", taskId),
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

	c.Status(http.StatusOK).JSON(task.EntityToDto())
}

func (h *Handler) UpdateTask(c *fiber.Ctx) {
	taskId := c.Params("taskId")
	if _, err := uuid.Parse(taskId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user id",
			Data:       nil,
		})
		return
	}

	updateTaskDto := &dto.UpdateTaskDto{}
	if err := c.BodyParser(updateTaskDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       nil,
		})
		return
	}

	task := &entity.Task{}
	err := h.taskRepo.FindTaskById(taskId, task)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound).JSON(&dto.ResponseErr{
				StatusCode: http.StatusNotFound,
				Message:    fmt.Sprintf("Task id %s not found", taskId),
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

	if updateTaskDto.Name != "" {
		task.Name = updateTaskDto.Name
	}
	if updateTaskDto.Content != "" {
		task.Content = updateTaskDto.Content
	}
	if updateTaskDto.Status != "" {
		task.Status = updateTaskDto.Status
	}
	if updateTaskDto.Deadline != "" {
		deadlineTime, _ := time.Parse(time.RFC3339, updateTaskDto.Deadline)
		task.Deadline = deadlineTime
	}

	err = h.taskRepo.UpdateTask(taskId, task)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	c.Status(http.StatusOK).JSON(task.EntityToDto())
}

func (h *Handler) DeleteTask(c *fiber.Ctx) {
	taskId := c.Params("taskId")
	if _, err := uuid.Parse(taskId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user id",
			Data:       nil,
		})
		return
	}

	err := h.taskRepo.DeleteTask(taskId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound).JSON(&dto.ResponseErr{
				StatusCode: http.StatusNotFound,
				Message:    fmt.Sprintf("Task id %s not found", taskId),
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

func (h *Handler) AssignUser(c *fiber.Ctx) {
	assignTask := &dto.AssignTaskDto{}
	if err := c.BodyParser(assignTask); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
		})
		return
	}

	userId, _ := uuid.Parse(assignTask.UserId)
	taskId, _ := uuid.Parse(assignTask.TaskId)
	userTask := &entity.UserTask{
		UserID: &userId,
		TaskID: &taskId,
	}
	err := h.userTaskRepo.AssignTask(userTask)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) RemoveUser(c *fiber.Ctx) {
	taskId := c.Params("taskId")
	userId := c.Params("userId")
	if _, err := uuid.Parse(taskId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid task id",
			Data:       nil,
		})
		return
	}
	if _, err := uuid.Parse(userId); err != nil {
		c.Status(http.StatusBadRequest).JSON(&dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user id",
			Data:       nil,
		})
		return
	}

	err := h.userTaskRepo.RemoveUser(userId, taskId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(http.StatusNotFound).JSON(&dto.ResponseErr{
				StatusCode: http.StatusNotFound,
				Message:    "Not found",
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
