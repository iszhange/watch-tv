/*
请求、响应日志
*/
package middleware

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"watchtv/utils"

	"github.com/go-chi/chi/v5/middleware"
)

type responseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(r.Context())
		var bodyBuffer bytes.Buffer
		teeReader := io.TeeReader(r.Body, &bodyBuffer)
		body, err := io.ReadAll(teeReader)
		if err != nil {
			body = []byte{}
		}
		bodyStr := string(body)
		r.Body = io.NopCloser(&bodyBuffer)
		rw := &responseWriter{ResponseWriter: w, body: &bytes.Buffer{}}

		utils.InfoMessage(fmt.Sprintf("HTTP Serve %s %s %s %v", reqID, r.Method, r.URL, bodyStr))

		next.ServeHTTP(rw, r)

		rBody := ""
		if isGzipResponse(rw) {
			rBody, _ = decompressData(rw.body.Bytes())
		} else {
			rBody = rw.body.String()
		}
		utils.InfoMessage(fmt.Sprintf("HTTP Serve %s %s", reqID, rBody))
	})
}

// 判断响应内容是否被gzip压缩
func isGzipResponse(w http.ResponseWriter) bool {
	// 获取响应头中的Content-Encoding字段
	contentEncoding := w.Header().Get("Content-Encoding")
	// 检查Content-Encoding是否为gzip
	return contentEncoding == "gzip"
}

// 解压数据
func decompressData(compressedData []byte) (string, error) {
	gr, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return "", err
	}
	defer gr.Close()

	uncompressed, err := io.ReadAll(gr)
	if err != nil {
		return "", err
	}

	return string(uncompressed), nil
}
