package handler

import (
	"api_gateway/genproto/task_service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/task/getall [GET]
// @Summary Get all taskes
// @Description API for getting all taskes
// @Tags task
// @Accept  json
// @Produce  json
// @Param		user_id query string true "user_id"
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllTask(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		task := &task_service.GetListTaskRequest{}

		search := c.Query("search")
		user_id := c.Query("user_id")

		page, err := ParsePageQueryParam(c)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while parsing page")
			return
		}
		limit, err := ParseLimitQueryParam(c)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while parsing limit")
			return
		}

		task.Search = search
		task.Offset = int64(page)
		task.Limit = int64(limit)
		task.OwnerId = user_id

		resp, err := h.grpcClient.TaskService().GetList(c.Request.Context(), task)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating task")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change task")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/task/create [POST]
// @Summary Create task
// @Description API for creating taskes
// @Tags task
// @Accept  json
// @Produce  json
// @Param		task body  task_service.CreateTask true "task"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateTask(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		task := &task_service.CreateTask{}
		if err := c.ShouldBindJSON(&task); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.TaskService().Create(c.Request.Context(), task)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating task")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change task")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/task/update [PUT]
// @Summary Update task
// @Description API for Updating tasks
// @Tags task
// @Accept  json
// @Produce  json
// @Param		task body  task_service.UpdateTask true "task"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateTask(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		task := &task_service.UpdateTask{}
		if err := c.ShouldBindJSON(&task); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.TaskService().Update(c.Request.Context(), task)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating task")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and task can change task")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/task/get/{id} [GET]
// @Summary Get task
// @Description API for getting task
// @Tags task
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetTaskById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		id := c.Param("id")
		task := &task_service.TaskPrimaryKey{Id: id}

		resp, err := h.grpcClient.TaskService().GetByID(c.Request.Context(), task)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting task")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and tasks can change task")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/task/get_by_task_id/{id} [GET]
// @Summary Get task
// @Description API for getting task
// @Tags task
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByExternalId(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		id := c.Param("id")
		task := &task_service.TaskPrimaryKey{Id: id}

		resp, err := h.grpcClient.TaskService().GetByExternalId(c.Request.Context(), task)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting task")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and tasks can change task")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/task/delete/{id} [DELETE]
// @Summary Delete task
// @Description API for deleting task
// @Tags task
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteTask(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		id := c.Param("id")
		task := &task_service.TaskPrimaryKey{Id: id}

		resp, err := h.grpcClient.TaskService().Delete(c.Request.Context(), task)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting task")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, task can change task")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/task/change_status [PATCH]
// @Summary Update task
// @Description API for Updating taskes
// @Tags task
// @Accept  json
// @Produce  json
// @Param		task body  task_service.TaskChangeStatus true "task"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) TaskChangeStatus(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		task := &task_service.TaskChangeStatus{}
		if err := c.ShouldBindJSON(&task); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		resp, err := h.grpcClient.TaskService().ChangeStatus(c.Request.Context(), task)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing task's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, task  can change task")
	}
}
