package res

//响应封装包
import (
	"blog_server/utils"
	valid "blog_server/utils/vaild"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = 0
	Error   = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r Response) Json() string {
	byteData, _ := json.Marshal(r)
	return string(byteData)
}

type ListRes[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

// 调用响应
func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithDataSSE(data any, c *gin.Context) {
	content := Response{
		Code: Success,
		Data: data,
		Msg:  "成功啦",
	}.Json()
	c.SSEvent("", content)
}
func OkWithSSE(data any, msg string, c *gin.Context) {
	content := Response{
		Code: Success,
		Data: data,
		Msg:  msg,
	}.Json()
	c.SSEvent("", content)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功啦", c)

}
func OkWithList(List any, count int64, c *gin.Context) {
	OkWithData(ListRes[any]{
		List:  List,
		Count: count,
	}, c)
}
func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}
func OkWith(c *gin.Context) {
	Result(Success, map[string]any{}, "成功啦", c)
}
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithMessageSSE(msg string, c *gin.Context) {
	data := Response{
		Code: Error,
		Data: map[string]any{},
		Msg:  msg,
	}.Json()
	c.SSEvent("", data)
}
func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
func FailWithValidError(err error, c *gin.Context) {
	msg := valid.Error(err)
	FailWithMessage(msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
