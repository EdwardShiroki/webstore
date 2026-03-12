package handler

import (
	"net/http"
	"time"

	"github.com/EdwardShiroki/webstore/internal/domain/item"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemHandler struct {
	itemRepo item.Repository
}

func NewItemHandler(itemRepo item.Repository) *ItemHandler {
	return &ItemHandler{itemRepo: itemRepo}
}

func (h *ItemHandler) Create(c *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	newItem := &item.Item{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}
	err := h.itemRepo.Create(c.Request.Context(), newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Item created successfully")
}

func (h *ItemHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}
	item, err := h.itemRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) List(c *gin.Context) {
	items, err := h.itemRepo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
