package api_new

import (
	"blog_server/models/res"
	"blog_server/service/ser_redis"
	"blog_server/utils/requests"
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewResponse struct {
	//Code int                 `json:"code"`
	//Data []ser_redis.NewData `json:"data"`
	//Msg  string              `json:"msg"`   //0.1版本存档
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    []ser_redis.NewData `json:"data"`
}
type OutputResponse struct {
	Code int                 `json:"code"`
	Data []ser_redis.NewData `json:"data"`
}

const newAPI = "https://api.vvhan.com/api/hotlist/baiduRD"
const timeout = 2 * time.Minute

// old v0.1
//
//	func (ApiNew) NewList(c *gin.Context) {
//		var cr params
//		var headers header
//		err := c.ShouldBindJSON(&cr)
//		err = c.ShouldBindHeader(&headers)
//		if err != nil {
//			res.FailWithCode(res.ArgumentError, c)
//			return
//		}
//		if cr.Size == 0 {
//			cr.Size = 1
//		}
//
//		key := fmt.Sprintf("%s-%d", cr.ID, cr.Size)
//		newsData, _ := ser_redis.GetNews(key)
//		if len(newsData) != 0 {
//			res.OkWithData(newsData, c)
//			return
//		}
//
//		httpResponse, err := requests.Post(newAPI, cr, structs.Map(headers), timeout)
//		if err != nil {
//			res.FailWithMessage(err.Error(), c)
//			return
//		}
//
//		var response NewResponse
//		byteData, err := io.ReadAll(httpResponse.Body)
//		err = json.Unmarshal(byteData, &response)
//		if err != nil {
//			res.FailWithMessage(err.Error(), c)
//			return
//		}
//		if response.Code != 200 {
//			res.FailWithMessage(response.Msg, c)
//			return
//		}
//		res.OkWithData(response.Data, c)
//		ser_redis.SetNews(key, response.Data)
//		return
//	}
func (ApiNew) NewList(c *gin.Context) {
	var headers header
	err := c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	key := "baidu-hotlist"
	newsData, _ := ser_redis.GetNews(key)
	if len(newsData) != 0 {
		res.OkWithData(newsData, c)
		return
	}

	httpResponse, err := requests.Get(newAPI, structs.Map(headers), timeout)
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
	res.OkWithData(response.Data, c)
	ser_redis.SetNews(key, response.Data)
	return
}
