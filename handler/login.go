package handler

import (
	"RemoteControlServer/dao"
	"RemoteControlServer/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Login 小程序登录
// @Summary 小程序登录
// @Description 获取JSON
// @Tags Login
// @Accept  application/json
// @Product application/json
// @Param data body model.Res true "res"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /login [post]
func Login(c *gin.Context) {

	data := dao.UserDao{}
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error,
		})
		return
	}

	url := "https://api.weixin.qq.com/sns/jscode2session?" + "appid=" + os.Getenv("APP_ID") + "&secret=" + os.Getenv("APP_SECRECT") + "&js_code=" + data.Code + "&grant_type=authorization_code"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error,
		})
		return
	}

	defer resp.Body.Close()

	var result model.CheckLogin
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error,
		})
		return
	}
	fmt.Println(result)
	data.UserId = result.OpenId
	data.UserOnline()
	c.Set("openId", &result.OpenId)
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"openId": result.OpenId,
	})
}
