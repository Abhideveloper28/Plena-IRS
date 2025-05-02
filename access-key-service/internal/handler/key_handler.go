package handler

import (
	"access-key-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.AccessKeyService
}

func NewHandler(svc *service.AccessKeyService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateKey(c *gin.Context) {
	var input struct {
		RateLimit     int `json:"rate_limit"`
		ExpiryMinutes int `json:"expiry_minutes"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	key, err := h.svc.CreateKey(input.RateLimit, input.ExpiryMinutes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create key"})
		return
	}
	c.JSON(http.StatusOK, key)
}

func (h *Handler) ListKeys(c *gin.Context) {
	keys, err := h.svc.ListKeys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch keys"})
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (h *Handler) UpdateKey(c *gin.Context) {
	key := c.Param("key")
	var input struct {
		RateLimit     int `json:"rate_limit"`
		ExpiryMinutes int `json:"expiry_minutes"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	if err := h.svc.UpdateKey(key, input.RateLimit, input.ExpiryMinutes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteKey(c *gin.Context) {
	key := c.Param("key")
	if err := h.svc.DeleteKey(key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetPlanByKey(c *gin.Context) {
	key := c.Param("key")
	k, err := h.svc.GetPlan(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}
	c.JSON(http.StatusOK, k)
}

func (h *Handler) DisableKey(c *gin.Context) {
	key := c.Param("key")
	if err := h.svc.DisableKey(key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Disable failed"})
		return
	}
	c.Status(http.StatusOK)
}
