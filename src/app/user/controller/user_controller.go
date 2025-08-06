package controller

import (
	"net/http"
	"strconv"

	"github.com/faisd405/go-restapi-gin/src/app/user/model"
	"github.com/faisd405/go-restapi-gin/src/app/user/service"
	"github.com/faisd405/go-restapi-gin/src/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "User registration data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /auth/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	user, err := ctrl.userService.Register(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Registration failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", user.ToResponse())
}

// Login godoc
// @Summary Login user
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body model.LoginRequest true "User login credentials"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /auth/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	loginResponse, err := ctrl.userService.Login(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Login failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", loginResponse)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get current user's profile
// @Tags users
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /users/profile [get]
func (ctrl *UserController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated", "user ID not found")
		return
	}

	profile, err := ctrl.userService.GetProfile(userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Profile not found", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile retrieved successfully", profile)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update current user's profile
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body model.UpdateUserRequest true "User update data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /users/profile [put]
func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated", "user ID not found")
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	updatedProfile, err := ctrl.userService.UpdateProfile(userID.(uint), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Profile update failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile updated successfully", updatedProfile)
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change current user's password
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param passwords body model.ChangePasswordRequest true "Password change data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /users/change-password [put]
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated", "user ID not found")
		return
	}

	var req model.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	err := ctrl.userService.ChangePassword(userID.(uint), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Password change failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Password changed successfully", nil)
}

// GetAllUsers godoc
// @Summary Get all users (Admin only)
// @Description Get paginated list of all users
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Router /admin/users [get]
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	users, total, err := ctrl.userService.GetAllUsers(page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get users", err.Error())
		return
	}

	response := map[string]interface{}{
		"users": users,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	}

	utils.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", response)
}

// DeleteUser godoc
// @Summary Delete user (Admin only)
// @Description Delete a user by ID
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Router /admin/users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	err = ctrl.userService.DeleteUser(uint(userID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User deletion failed", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}
