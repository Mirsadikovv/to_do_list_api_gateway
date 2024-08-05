package handler

import (
	"api_gateway/genproto/admin_service"
	"api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/admin/getall [GET]
// @Summary Get all admines
// @Description API for getting all admines
// @Tags admin
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllAdmin(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" {

		admin := &admin_service.GetListAdminRequest{}

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

		admin.Search = search
		admin.Offset = int64(page)
		admin.Limit = int64(limit)

		resp, err := h.grpcClient.AdminService().GetList(c.Request.Context(), admin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating admin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change admin")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/admin/create [POST]
// @Summary Create admin
// @Description API for creating admines
// @Tags admin
// @Accept  json
// @Produce  json
// @Param		admin body  admin_service.CreateAdmin true "admin"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateAdmin(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" {

		admin := &admin_service.CreateAdmin{}
		if err := c.ShouldBindJSON(&admin); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		if !validator.ValidateGmail(admin.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(admin.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidatePassword(admin.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.AdminService().Create(c.Request.Context(), admin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating admin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins can change admin")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/admin/update [PUT]
// @Summary Update admin
// @Description API for Updating admins
// @Tags admin
// @Accept  json
// @Produce  json
// @Param		admin body  admin_service.UpdateAdmin true "admin"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateAdmin(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" {

		admin := &admin_service.UpdateAdmin{}
		if err := c.ShouldBindJSON(&admin); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(admin.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(admin.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		resp, err := h.grpcClient.AdminService().Update(c.Request.Context(), admin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating admin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and admin can change admin")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/admin/get/{id} [GET]
// @Summary Get admin
// @Description API for getting admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAdminById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" {

		id := c.Param("id")
		admin := &admin_service.AdminPrimaryKey{Id: id}

		resp, err := h.grpcClient.AdminService().GetByID(c.Request.Context(), admin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting admin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and admins can change admin")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/admin/delete/{id} [DELETE]
// @Summary Delete admin
// @Description API for deleting admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteAdmin(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" {

		id := c.Param("id")
		admin := &admin_service.AdminPrimaryKey{Id: id}

		resp, err := h.grpcClient.AdminService().Delete(c.Request.Context(), admin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting admin")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, admin can change admin")
	}
}

// AdminLogin godoc
// @Router       /v1/admin/login [POST]
// @Summary      Admin login
// @Description  Admin login
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        login body admin_service.AdminLoginRequest true "login"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) AdminLogin(c *gin.Context) {
	loginReq := &admin_service.AdminLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.AdminService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// AdminRegister godoc
// @Router       /v1/admin/register [POST]
// @Summary      Admin register
// @Description  Admin register
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        register body admin_service.AdminRegisterRequest true "register"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) AdminRegister(c *gin.Context) {
	loginReq := &admin_service.AdminRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate for (gmail.com or mail.ru) & check if email is not exists

	resp, err := h.grpcClient.AdminService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr admin")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// AdminRegister godoc
// @Router       /v1/admin/register-confirm [POST]
// @Summary      Admin register
// @Description  Admin register
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        register body admin_service.AdminRegisterConfRequest true "register"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) AdminRegisterConfirm(c *gin.Context) {
	req := &admin_service.AdminRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	if !validator.ValidateGmail(req.Admin[0].Email) {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
		return
	}

	if !validator.ValidatePhone(req.Admin[0].Phone) {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
		return
	}

	err := validator.ValidatePassword(req.Admin[0].UserPassword)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
		return
	}

	confResp, err := h.grpcClient.AdminService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/admin/change_password [PATCH]
// @Summary Update admin
// @Description API for Updating admines
// @Tags admin
// @Accept  json
// @Produce  json
// @Param		admin body  admin_service.AdminChangePassword true "admin"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) AdminChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" {

		admin := &admin_service.AdminChangePassword{}
		if err := c.ShouldBindJSON(&admin); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		err := validator.ValidatePassword(admin.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.AdminService().ChangePassword(c.Request.Context(), admin)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing admin's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, admin  can change admin")
	}
}
