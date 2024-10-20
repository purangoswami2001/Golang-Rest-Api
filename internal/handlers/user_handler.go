package handlers

import (
	"api/internal/models"
	"api/internal/services"
	"api/internal/utils"
	"api/internal/validations"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
		return
	}

	// Validate user fields
	validationErrors := validations.ValidateUser(user, false)
	if len(validationErrors.Errors) > 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(validationErrors.Errors[0].Message))
		return
	}

	// Hash the password before saving
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	if err := h.service.Create(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse("Email Already Exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create user"))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("User created successfully", user))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid user ID"))
		return
	}

	// Retrieve the existing user by ID
	user, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
		return
	}

	// Hash the password before saving
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Bind the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to bind user data"))
		return
	}

	// Update the user fields and save it
	if err := h.service.Update(user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse("Email Already Exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update user"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("User updated successfully", user))
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid user ID"))
		return
	}

	// Retrieve the user by ID
	user, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
		return
	}

	// Return the user as a JSON response
	c.JSON(http.StatusOK, utils.SuccessResponse("User retrieved successfully", user))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid user ID"))
		return
	}

	// Delete the user by ID
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete user"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("User deleted successfully", nil))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve users"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Users retrieved successfully", users))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
