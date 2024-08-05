package handler

import (
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/user/getall [GET]
// @Summary Get all useres
// @Description API for getting all useres
// @Tags user
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllUser(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "admin" {

		user := &user_service.GetListUserRequest{}

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

		user.Search = search
		user.Offset = int64(page)
		user.Limit = int64(limit)

		resp, err := h.grpcClient.UserService().GetList(c.Request.Context(), user)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating user")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and admins can change user")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/user/create [POST]
// @Summary Create user
// @Description API for creating useres
// @Tags user
// @Accept  json
// @Produce  json
// @Param		user body  user_service.CreateUser true "user"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateUser(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		user := &user_service.CreateUser{}
		if err := c.ShouldBindJSON(&user); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(user.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(user.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidateBitrthday(user.Birthday)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong birthday"), "error while validating birthday")
			return
		}

		err = validator.ValidatePassword(user.UserPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}

		resp, err := h.grpcClient.UserService().Create(c.Request.Context(), user)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while creating user")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins and admins can change user")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/user/update [PUT]
// @Summary Update user
// @Description API for Updating useres
// @Tags user
// @Accept  json
// @Produce  json
// @Param		user body  user_service.UpdateUser true "user"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateUser(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		user := &user_service.UpdateUser{}
		if err := c.ShouldBindJSON(&user); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}
		if !validator.ValidateGmail(user.Email) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}

		if !validator.ValidatePhone(user.Phone) {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
			return
		}

		err := validator.ValidateBitrthday(user.Birthday)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
			return
		}
		resp, err := h.grpcClient.UserService().Update(c.Request.Context(), user)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while updating user")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, user and admins can change user")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/user/get/{id} [GET]
// @Summary Get user
// @Description API for getting user
// @Tags user
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetUserById(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		id := c.Param("id")
		user := &user_service.UserPrimaryKey{Id: id}

		resp, err := h.grpcClient.UserService().GetByID(c.Request.Context(), user)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while getting user")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, user and admins can change user")
	}
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/user/delete/{id} [DELETE]
// @Summary Delete user
// @Description API for deleting user
// @Tags user
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteUser(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		id := c.Param("id")
		user := &user_service.UserPrimaryKey{Id: id}

		resp, err := h.grpcClient.UserService().Delete(c.Request.Context(), user)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while deleting user")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, user and admins can change user")
	}
}

// UserLogin godoc
// @Router       /v1/user/login [POST]
// @Summary      User login
// @Description  User login
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        login body user_service.UserLoginRequest true "login"
// @Success		 200  {object}  models.ResponseSuccess
// @Failure		 400  {object}  models.ResponseError
// @Failure		 404  {object}  models.ResponseError
// @Failure		 500  {object}  models.ResponseError
func (h *handler) UserLogin(c *gin.Context) {
	loginReq := &user_service.UserLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	//TODO: need validate login & password

	loginResp, err := h.grpcClient.UserService().Login(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "unauthorized")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, loginResp)

}

// UserRegister godoc
// @Router       /v1/user/register [POST]
// @Summary      User register
// @Description  User register
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        register body user_service.UserRegisterRequest true "register"
// @Success	   	 200  {object}  models.ResponseSuccess
// @Failure		 400  {object}  models.ResponseError
// @Failure	     404  {object}  models.ResponseError
// @Failure		 500  {object}  models.ResponseError
func (h *handler) UserRegister(c *gin.Context) {
	loginReq := &user_service.UserRegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("loginReq: ", loginReq)

	resp, err := h.grpcClient.UserService().Register(c.Request.Context(), loginReq)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while registr user")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Otp sent successfull")
	c.JSON(http.StatusOK, resp)
}

// UserRegister godoc
// @Router       /v1/user/register-confirm [POST]
// @Summary      User register
// @Description  User register
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        register body user_service.UserRegisterConfRequest true "register"
// @Success		 200  {object}  models.ResponseSuccess
// @Failure		 400  {object}  models.ResponseError
// @Failure		 404  {object}  models.ResponseError
// @Failure		 500  {object}  models.ResponseError
func (h *handler) UserRegisterConfirm(c *gin.Context) {
	req := &user_service.UserRegisterConfRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}
	fmt.Println("req: ", req)

	if !validator.ValidateGmail(req.User[0].Email) {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
		return
	}

	if !validator.ValidatePhone(req.User[0].Phone) {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong phone"), "error while validating phone")
		return
	}

	err := validator.ValidatePassword(req.User[0].UserPassword)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
		return
	}
	confResp, err := h.grpcClient.UserService().RegisterConfirm(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while confirming")
		return
	}

	handleGrpcErrWithDescription(c, h.log, nil, "Succes")
	c.JSON(http.StatusOK, confResp)
}

// @Security ApiKeyAuth
// @Param   Authorization  header  string  true  "Authorization token"
// @Router /v1/user/change_password [PATCH]
// @Summary Update user
// @Description API for Updating useres
// @Tags user
// @Accept  json
// @Produce  json
// @Param		user body  user_service.UserChangePassword true "user"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UserChangePassword(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "Unauthorized")

	}
	if authInfo.UserRole == "superadmin" || authInfo.UserRole == "admin" || authInfo.UserRole == "user" {

		user := &user_service.UserChangePassword{}
		if err := c.ShouldBindJSON(&user); err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
			return
		}

		err := validator.ValidatePassword(user.NewPassword)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, errors.New("wrong password"), "error while validating password")
			return
		}
		resp, err := h.grpcClient.UserService().ChangePassword(c.Request.Context(), user)
		if err != nil {
			handleGrpcErrWithDescription(c, h.log, err, "error while changing user's password")
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		handleGrpcErrWithDescription(c, h.log, errors.New("Forbidden"), "Only superadmins, managers, user and admins can change user")
	}
}
