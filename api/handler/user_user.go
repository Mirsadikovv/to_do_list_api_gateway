package handler

import (
	"api_gateway/genproto/student_service"
	"api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/student/getall [GET]
// @Summary Get all studentes
// @Description API for getting all studentes
// @Tags student
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllStudent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" {

		student := &student_service.GetListStudentRequest{}

		search := c.Query("search")

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

		student.Search = search
		student.Offset = int64(page)
		student.Limit = int64(limit)

		resp, err := h.grpcClient.StudentService().GetList(c.Request.Context(), student)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating student")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and administrators can change student")
	}
}

// @Security ApiKeyAuth
// @Router /v1/student/create [POST]
// @Summary Create student
// @Description API for creating studentes
// @Tags student
// @Accept  json
// @Produce  json
// @Param		student body  student_service.CreateStudent true "student"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateStudent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" {

		student := &student_service.CreateStudent{}
		if err := c.ShouldBindJSON(&student); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(student.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(student.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidateBitrthday(student.Birthday)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		err = validator.ValidatePassword(student.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}

		resp, err := h.grpcClient.StudentService().Create(c.Request.Context(), student)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating student")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers and administrators can change student")
	}
}

// @Security ApiKeyAuth
// @Router /v1/student/update [PUT]
// @Summary Update student
// @Description API for Updating studentes
// @Tags student
// @Accept  json
// @Produce  json
// @Param		student body  student_service.UpdateStudent true "student"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateStudent(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" || authInfo.UserRole == "student" {

		student := &student_service.UpdateStudent{}
		if err := c.ShouldBindJSON(&student); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(student.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(student.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidateBitrthday(student.Birthday)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}
		resp, err := h.grpcClient.StudentService().Update(c.Request.Context(), student)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating student")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers, student and administrators can change student")
	}
}

// @Security ApiKeyAuth
// @Router /v1/student/get/{id} [GET]
// @Summary Get student
// @Description API for getting student
// @Tags student
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetStudentById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" || authInfo.UserRole == "student" {

		id := c.Param("id")
		student := &student_service.StudentPrimaryKey{Id: id}

		resp, err := h.grpcClient.StudentService().GetByID(c.Request.Context(), student)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting student")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers, student and administrators can change student")
	}
}

// @Security ApiKeyAuth
// @Router /v1/student/delete/{id} [DELETE]
// @Summary Delete student
// @Description API for deleting student
// @Tags student
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	student := &student_service.StudentPrimaryKey{Id: id}

	resp, err := h.grpcClient.StudentService().Delete(c.Request.Context(), student)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting student")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// StudentLogin godoc
// @Router       /v1/student/login [POST]
// @Summary      Student login
// @Description  Student login
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        login body student_service.StudentLoginRequest true "login"
// @Success      201  {object}  student_service.StudentLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) StudentLogin(c *gin.Context) {
	loginReq := &student_service.StudentLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.StudentService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// StudentRegister godoc
// @Router       /v1/student/register [POST]
// @Summary      Student register
// @Description  Student register
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        register body student_service.StudentRegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) StudentRegister(c *gin.Context) {
	loginReq := &student_service.StudentRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	resp, err := h.grpcClient.StudentService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr student")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// StudentRegister godoc
// @Router       /v1/student/register-confirm [POST]
// @Summary      Student register
// @Description  Student register
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        register body student_service.StudentRegisterConfRequest true "register"
// @Success      201  {object}  student_service.StudentLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *handler) StudentRegisterConfirm(c *gin.Context) {
	req := &student_service.StudentRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	//TODO: need validate login & password

	confResp, err := h.grpcClient.StudentService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Router /v1/student/change_password [PATCH]
// @Summary Update student
// @Description API for Updating studentes
// @Tags student
// @Accept  json
// @Produce  json
// @Param		student body  student_service.UpdateStudent true "student"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) StudentChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "administrator" || authInfo.UserRole == "manager" || authInfo.UserRole == "student" {

		student := &student_service.StudentChangePassword{}
		if err := c.ShouldBindJSON(&student); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		err := validator.ValidatePassword(student.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.StudentService().ChangePassword(c.Request.Context(), student)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing student's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers, student and administrators can change student")
	}
}
