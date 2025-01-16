package handler

import (
	"api-gateway/pkg/pb"
	"api-gateway/pkg/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Client        pb.UserServiceClient
	methodService *usecase.UserService
}

func NewUserHandler(client pb.UserServiceClient, userService *usecase.UserService) *UserHandler {
	return &UserHandler{
		Client:        client,
		methodService: userService,
	}
}

type MethodRequest struct {
	Method   int `json:"method"`
	WaitTime int `json:"waitTime"`
}

func (h *UserHandler) Methods(c *gin.Context) {
	var req MethodRequest

	// Parse and validate the request body.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ctx := c.Request.Context()
	var userNames []string
	var err error

	switch req.Method {
	case 1:
		userNames, err = h.methodService.Method1(ctx, req.WaitTime)
	case 2:
		userNames, err = h.methodService.Method2(ctx, req.WaitTime)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid method"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userNames": userNames})
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req pb.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.Client.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserByID retrieves a user by their ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	req := &pb.GetUserByIDRequest{Id: parsedID}
	res, err := h.Client.GetUserByID(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUser updates an existing user's details.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req pb.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.Client.UpdateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteUser deletes a user by their ID.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	req := &pb.DeleteUserRequest{Id: parsedID}
	res, err := h.Client.DeleteUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListAllUsers lists all the users in the system.
func (h *UserHandler) ListAllUsers(c *gin.Context) {
	res, err := h.Client.ListAllUsers(c.Request.Context(), &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching users: %v", err)})
		return
	}

	c.JSON(http.StatusOK, res)
}
