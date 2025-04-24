package v1

import (
	"NFTAuctionBackend/src/service/svc"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// 配置常量
const (
	UploadDir      = "./uploads"                // 文件保存目录
	MaxUploadSize  = 50 << 20                   // 最大文件大小（50MB）
	AllowedFormats = ".jpg,.png,.gif,.mp4,.mp3" // 允许的文件格式
)

// UploadResponse 文件上传成功响应
type UploadResponse struct {
	Status string `json:"status" example:"success"`
	Path   string `json:"path" example:"./uploads/123456789.jpg"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error" example:"File too large (max 50MB)"`
}

// @Summary      上传文件（NFT附件）
// @Description  接收用户上传的附件（图片、视频等）并保存到本地
// @Tags         NFT
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "上传的文件（支持JPG/PNG/GIF/MP4/MP3）"
// @Success      200  {object}  map[string]interface{}  "成功返回示例：{'status': 'success', 'path': '/uploads/123.jpg'}"
// @Failure      400  {object}  map[string]interface{}  "错误示例：{'error': 'File too large'}"
// @Failure      500  {object}  map[string]interface{}  "错误示例：{'error': 'Failed to save file'}"
// @Router       /uploadFile/uploadNftFile [post]
func UploadHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取上传的文件
		file, err := c.FormFile("file") // "file" 是前端表单的字段名
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
			return
		}

		// 检查文件大小
		if file.Size > MaxUploadSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File too large (max 50MB)"})
			return
		}

		// 检查文件格式
		ext := filepath.Ext(file.Filename)
		if !isAllowedFormat(ext) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Unsupported file format. Allowed: %s", AllowedFormats),
			})
			return
		}

		// 生成唯一文件名（防止冲突）
		newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join(UploadDir, newFilename)

		// 保存文件到本地
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"path":   "./" + filePath,
		})
	}
}

// 检查文件格式是否允许
func isAllowedFormat(ext string) bool {
	allowed := map[string]bool{
		".jpg": true,
		".png": true,
		".gif": true,
		".mp4": true,
		".mp3": true,
	}
	return allowed[ext]
}

// @Description  图片回显接口
// GetFile 获取并返回指定路径的文件内容
// @Summary 获取文件内容
// @Tags 文件操作
// @Accept json
// @Produce octet-stream
// @Produce json
// @Param path query string true "文件路径" Example("./uploads/example.jpg")
// @Success 200 {file} file "文件内容"
// @Failure 400 {object} map[string]interface{} "缺少path参数"
// @Failure 500 {object} map[string]interface{} "文件操作失败"
// @Router       /uploadFile/uploadNftFile [post]
func GetFile(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取前端传入的图片URL
		filePath := c.Query("path")
		if filePath == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "path parameter is required"})
			return
		}

		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file: " + err.Error()})
			return
		}
		defer file.Close()

		// 读取文件头信息判断文件类型
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file: " + err.Error()})
			return
		}
		contentType := http.DetectContentType(buffer)

		// 重置文件指针
		_, err = file.Seek(0, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reset file pointer: " + err.Error()})
			return
		}

		// 回显文件
		c.DataFromReader(http.StatusOK, -1, contentType, file, nil)
	}
}
