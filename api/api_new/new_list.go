package api_new

import (
	"blog_server/models/res"
	"blog_server/service/ser_redis"
	"blog_server/utils/requests"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"time"
)

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    []ser_redis.NewData `json:"data"`
}

const timeout = 2 * time.Minute

func (ApiNew) NewList(c *gin.Context) {
	var headers header
	err := c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 获取 source 和 size 参数
	source := c.Query("source")
	sizeStr := c.Query("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		size = 10 // 默认返回 10 条数据
	}

	// 定义新闻源和对应的 API 地址
	apiMap := map[string]string{
		"baidu":    "https://api.vvhan.com/api/hotlist/baiduRD",
		"bilibili": "https://api.vvhan.com/api/hotlist/bili",
		"zhihu":    "https://api.vvhan.com/api/hotlist/zhihuHot",
		"weibo":    "https://api.vvhan.com/api/hotlist/wbHot",
		"toutiao":  "https://api.vvhan.com/api/hotlist/toutiao",
	}

	// 检查 source 是否有效
	apiURL, ok := apiMap[source]
	if !ok {
		res.FailWithMessage("无效的新闻源参数", c)
		return
	}

	// 生成缓存键，包含 source 和 size
	key := fmt.Sprintf("%s-hotlist-%d", source, size)
	newsData, _ := ser_redis.GetNews(key)
	if len(newsData) != 0 {
		res.OkWithData(newsData, c)
		return
	}

	// 发送 GET 请求获取数据
	httpResponse, err := requests.Get(apiURL, structs.Map(headers), timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	var response NewResponse
	byteData, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	if !response.Success {
		res.FailWithMessage(response.Message, c)
		return
	}

	// 截取前 size 条数据
	limitedData := response.Data
	if size < len(response.Data) {
		limitedData = response.Data[:size]
	}

	res.OkWithData(limitedData, c)
	ser_redis.SetNews(key, limitedData)
	return
}
