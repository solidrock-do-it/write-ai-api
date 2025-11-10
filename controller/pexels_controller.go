package controller

import (
	"net/http"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/gin-gonic/gin"
)

// GetPexelsKeys 返回 Pexels keys 列表（受保护，需要鉴权）
func GetPexelsKeys(c *gin.Context) {
	// 从数据库读取
	provider, err := model.GetMaterialProvider("pexels")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    []string{},
		})
		return
	}
	keys, err := provider.GetKeys()
	if err != nil {
		common.ApiError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    keys,
	})
}

// UpdatePexelsKeys 更新 Pexels keys
func UpdatePexelsKeys(c *gin.Context) {
	var req struct {
		Keys []string `json:"keys"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	err = model.UpdateMaterialProvider("pexels", req.Keys)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}