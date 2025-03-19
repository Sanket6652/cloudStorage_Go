package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/upload", UploadFile)
	router.GET("/files", ListFiles)
	router.GET("/download/:file_id", DownloadFile)
	router.DELETE("/delete/:file_id", DeleteFile)
}

func UploadFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Upload endpoint"})
}

func ListFiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List files endpoint"})
}

func DownloadFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Download endpoint"})
}

func DeleteFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete endpoint"})
}
