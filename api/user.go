package api

import (
	"net/http"
	"project_sdu/model"
	"project_sdu/service"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

// ✅ REGISTER
func (u *userAPI) Register(c *gin.Context) {
	var req model.UserRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors: map[string]string{
				"body": "Invalid JSON format",
			},
		})
		return
	}

	// Validasi data kosong
	errors := make(map[string]string)
	if req.Fullname == "" {
		errors["fullname"] = "Fullname is required"
	}
	if req.Email == "" {
		errors["email"] = "Email is required"
	}
	if req.Password == "" {
		errors["password"] = "Password is required"
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  errors,
		})
		return
	}

	// 🔒 Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to hash password",
			Errors: map[string]string{
				"bcrypt": err.Error(),
			},
		})
		return
	}

	user := model.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := u.userService.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Registration failed",
			Errors: map[string]string{
				"server": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Registration successful",
		Data: gin.H{
			"fullname": req.Fullname,
			"email":    req.Email,
		},
	})
}

// LOGIN
func (u *userAPI) Login(c *gin.Context) {
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors: map[string]string{
				"body": "Invalid JSON format",
			},
		})
		return
	}

	// Validasi input
	errors := make(map[string]string)
	if req.Email == "" {
		errors["email"] = "Email is required"
	}
	if req.Password == "" {
		errors["password"] = "Password is required"
	}
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  errors,
		})
		return
	}

	token, userID, err := u.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Message: "Invalid email or password",
			Errors: map[string]string{
				"auth": err.Error(),
			},
		})
		return
	}

	// Simpan token ke cookie
	c.SetCookie("session_token", *token, int(time.Hour.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Login successful",
		Data: gin.H{
			"user_id": userID,
			"email":   req.Email,
			"token":   token,
		},
	})
}

// LOGOUT
func (u *userAPI) Logout(c *gin.Context) {

	c.SetCookie(
		"session_token",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Logout successful",
	})
}

// ✅ GET USER TASK CATEGORY (contoh)
func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// Contoh response (belum implementasi)
	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "GetUserTaskCategory not implemented yet",
		Data:    gin.H{},
	})
}
