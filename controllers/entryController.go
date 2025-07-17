package controllers

import (
	"log"
	"net/http"
	"strconv"
	"tryg/config"
	"tryg/models"
	"tryg/utils"

	"github.com/gin-gonic/gin"
)

type CreateEntryRequest struct {
	Key       string `json:"key" binding:"required"`
	Value     string `json:"value" binding:"required"`
	Timestamp *int64 `json:"timestamp"`
}

func CreateEntry(c *gin.Context) {
	var req CreateEntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Timestamp == nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Timestamp is required")
		return
	}

	entry := models.Entry{
		Key:       req.Key,
		Value:     req.Value,
		Timestamp: *req.Timestamp,
	}

	if err := config.DB.Create(&entry).Error; err != nil {
		log.Printf("Insert failed for key=%s: %v", req.Key, err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to store entry")
		return
	}

	utils.SuccessResponse(c, "Entry saved successfully", entry)
}

func GetEntryByTimestamp(c *gin.Context) {
	key := c.Query("key")
	timestampStr := c.Query("timestamp")

	if key == "" || timestampStr == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Key and timestamp are required")
		return
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid timestamp format")
		return
	}

	var entry models.Entry
	result := config.DB.
		Where("`key` = ? AND timestamp <= ?", key, timestamp).
		Order("timestamp DESC").
		Limit(1).
		First(&entry)

	if result.Error != nil {
		log.Printf("No entry found for key=%s at timestamp=%d", key, timestamp)
		utils.ErrorResponse(c, http.StatusNotFound, "No value found for the given key and timestamp")
		return
	}

	log.Printf("Found entry: key=%s value=%s timestamp=%d", entry.Key, entry.Value, entry.Timestamp)
	utils.SuccessResponse(c, "Value retrieved successfully", entry)
}