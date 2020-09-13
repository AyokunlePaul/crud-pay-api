package middleware

import (
	"bytes"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBuffer, _ := ioutil.ReadAll(c.Request.Body)
		firstReader := ioutil.NopCloser(bytes.NewBuffer(requestBuffer))
		secondReader := ioutil.NopCloser(bytes.NewBuffer(requestBuffer)) //We have to create a new Buffer, because firstReader will be read.

		logger.Info(readBody(firstReader))

		c.Request.Body = secondReader
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	readerBuffer := new(bytes.Buffer)
	_, _ = readerBuffer.ReadFrom(reader)

	content := readerBuffer.String()
	return content
}
