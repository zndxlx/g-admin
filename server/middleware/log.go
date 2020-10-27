package middleware

import (
    "github.com/gin-gonic/gin"
    "time"
    "g-admin/core/log"
    "go.uber.org/zap"
    "bytes"
    "net/http"
    "io/ioutil"
    "strings"
)

func GinCommonLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        query := c.Request.URL.RawQuery
        c.Next()
        
        cost := time.Since(start)
        log.Info(path,
            zap.Int("status", c.Writer.Status()),
            zap.String("method", c.Request.Method),
            zap.String("path", path),
            zap.String("query", query),
            zap.String("ip", c.ClientIP()),
            zap.String("user-agent", c.Request.UserAgent()),
            zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
            zap.Duration("cost", cost),
        )
    }
}



func GinDetailLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        var body []byte
        //var userId int
        if c.Request.Method != http.MethodGet {
            var err error
            body, err = ioutil.ReadAll(c.Request.Body)
            if err != nil {
                log.Error("read body from request error:", zap.Any("err", err))
            } else {
                c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
            }
        }
        // if claims, ok := c.Get("claims"); ok {
        //     waitUse := claims.(*request.CustomClaims)
        //     userId = int(waitUse.ID)
        // }else {
        //     id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
        //     if err != nil {
        //         userId = 0
        //     }
        //     userId = id
        // }

        values := c.Request.Header["content-type"]
        if len(values) >0 && strings.Contains(values[0], "boundary") {
            body = []byte("file")
        }
        writer := responseBodyWriter{
            ResponseWriter: c.Writer,
            body:           &bytes.Buffer{},
        }
        c.Writer = writer
        now := time.Now()
        
        c.Next()
        
        latency := time.Now().Sub(now)
        log.Info("",
            zap.Int("status", c.Writer.Status()),
            zap.String("method", c.Request.Method),
            zap.String("path", c.Request.URL.Path),
            zap.String("query", c.Request.URL.RawQuery),
            zap.String("ip", c.ClientIP()),
            zap.String("req", string(body)),
            zap.String("Resp", writer.body.String()),
            zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
            zap.Duration("latency", latency),
        )

    }
}

type responseBodyWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
    r.body.Write(b)
    return r.ResponseWriter.Write(b)
}