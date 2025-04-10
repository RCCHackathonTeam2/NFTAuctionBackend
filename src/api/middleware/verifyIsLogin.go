package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware 验证用户登录状态的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头/Cookie中获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			cookie, err := c.Cookie("auth_token")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    401,
					"message": "未提供认证令牌",
				})
				return
			}
			token = cookie
		}

		// 2. 验证token有效性（示例：JWT验证）
		//claims, err := ValidateToken(token) // 替换为你的token验证逻辑
		//if err != nil {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		//		"code":    401,
		//		"message": "无效或过期的令牌",
		//	})
		//	return
		//}

		// 3. 将用户信息存入上下文（可选）
		//c.Set("userID", claims.UserID)
		//c.Set("role", claims.Role)

		// 4. 继续处理后续请求
		c.Next()
	}
}

// ValidateToken 示例：JWT验证逻辑（需根据实际项目实现）
//func ValidateToken(token string) (*YourClaimsStruct, error) {
//	// 实现你的JWT解析和验证逻辑
//	// 例如使用 github.com/golang-jwt/jwt
//}
