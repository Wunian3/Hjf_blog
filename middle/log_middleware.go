package middle

import (
	log_stash "blog_server/plugin/log_stash_v2"
	"bytes"
	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	byteData *bytes.Buffer
}

func (rw responseWriter) Write(buf []byte) (int, error) {
	rw.byteData.Write(buf)
	return rw.ResponseWriter.Write(buf)
}

func LogMiddleWare() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		r := responseWriter{
			ResponseWriter: c.Writer,
			byteData:       bytes.NewBuffer([]byte{}),
		}
		c.Writer = r
		c.Next()
		_action, ok := c.Get("action")
		if !ok {
			return
		}
		action, ok := _action.(*log_stash.Action)
		if !ok {
			return
		}
		action.SetResponseContent(r.byteData.String())
		action.SetFlush()
	}
}
