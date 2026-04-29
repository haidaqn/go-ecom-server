package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

const (
	logFilePath           = "logs/http.log"
	maxInspectBodyBytes   = int64(64 * 1024)
	maxResponseLogBytes   = 64 * 1024
	maxLoggedMultipartVal = 2 * 1024
	logProbeBytes         = 4 * 1024
)

type customResponseWriter struct {
	gin.ResponseWriter
	body      *bytes.Buffer
	limit     int
	truncated bool
}

func (w *customResponseWriter) Write(data []byte) (n int, err error) {
	if !w.truncated && w.body.Len() < w.limit {
		remaining := w.limit - w.body.Len()
		if len(data) > remaining {
			w.body.Write(data[:remaining])
			w.truncated = true
		} else {
			w.body.Write(data)
		}
	}

	return w.ResponseWriter.Write(data)
}

func LoggerMiddleware() gin.HandlerFunc {
	_ = os.MkdirAll("logs", os.ModePerm)
	sanitizeHTTPLogFile(logFilePath)

	logger := zerolog.New(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
	}).With().Timestamp().Logger()

	return func(ctx *gin.Context) {
		start := time.Now()
		contentType := ctx.GetHeader("Content-Type")
		requestBody := make(map[string]any)
		var formFiles []map[string]any

		if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := ctx.Request.ParseMultipartForm(32 << 20); err == nil && ctx.Request.MultipartForm != nil {
				for key, vals := range ctx.Request.MultipartForm.Value {
					if len(vals) == 1 {
						requestBody[key] = truncateString(vals[0], maxLoggedMultipartVal)
					} else {
						requestBody[key] = truncateSlice(vals, maxLoggedMultipartVal)
					}
				}

				for field, files := range ctx.Request.MultipartForm.File {
					for _, f := range files {
						formFiles = append(formFiles, map[string]any{
							"field":        field,
							"filename":     f.Filename,
							"size":         formatFileSize(f.Size),
							"content_type": f.Header.Get("Content-Type"),
						})
					}
				}

				if len(formFiles) > 0 {
					requestBody["form_files"] = formFiles
				}
			}
		} else if ctx.Request.ContentLength >= 0 && ctx.Request.ContentLength <= maxInspectBodyBytes {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				logger.Error().Err(err).Msg("failed to read request body")
			}

			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			if strings.HasPrefix(contentType, "application/json") {
				_ = json.Unmarshal(bodyBytes, &requestBody)
			} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
				values, _ := url.ParseQuery(string(bodyBytes))
				for key, vals := range values {
					if len(vals) == 1 {
						requestBody[key] = truncateString(vals[0], maxLoggedMultipartVal)
					} else {
						requestBody[key] = truncateSlice(vals, maxLoggedMultipartVal)
					}
				}
			}
		} else {
			requestBody["_skipped"] = "request body too large or unknown content-length"
			requestBody["_content_length"] = ctx.Request.ContentLength
		}

		writer := &customResponseWriter{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBufferString(""),
			limit:          maxResponseLogBytes,
		}
		ctx.Writer = writer

		ctx.Next()

		duration := time.Since(start)
		statusCode := ctx.Writer.Status()
		responseContentType := ctx.Writer.Header().Get("Content-Type")
		responseBodyRaw := writer.body.String()
		var responseBodyParsed any

		switch {
		case strings.HasPrefix(responseContentType, "image/"):
			responseBodyParsed = "[BINARY DATA]"
		case strings.HasPrefix(responseContentType, "application/json"),
			strings.HasPrefix(strings.TrimSpace(responseBodyRaw), "{"),
			strings.HasPrefix(strings.TrimSpace(responseBodyRaw), "["):
			if err := json.Unmarshal([]byte(responseBodyRaw), &responseBodyParsed); err != nil {
				responseBodyParsed = responseBodyRaw
			}
		default:
			responseBodyParsed = responseBodyRaw
		}

		logEvent := logger.Info()
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Str("client_ip", ctx.ClientIP()).
			Str("user_agent", ctx.Request.UserAgent()).
			Str("referer", ctx.Request.Referer()).
			Str("protocol", ctx.Request.Proto).
			Str("host", ctx.Request.Host).
			Str("remote_addr", ctx.Request.RemoteAddr).
			Str("request_uri", ctx.Request.RequestURI).
			Int64("content_length", ctx.Request.ContentLength).
			Interface("headers", ctx.Request.Header).
			Interface("request_body", requestBody).
			Int("status_code", statusCode).
			Interface("response_body", responseBodyParsed).
			Bool("response_body_truncated", writer.truncated).
			Int64("duration_ms", duration.Milliseconds()).
			Msg("HTTP Request Log")
	}
}

func sanitizeHTTPLogFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return
	}
	defer file.Close()

	buf := make([]byte, logProbeBytes)
	n, readErr := file.Read(buf)
	if readErr != nil && readErr != io.EOF {
		return
	}

	if bytes.IndexByte(buf[:n], 0) == -1 {
		return
	}

	_ = file.Truncate(0)
	_, _ = file.Seek(0, 0)
}

func formatFileSize(size int64) string {
	switch {
	case size >= 1<<20:
		return fmt.Sprintf("%.2f MB", float64(size)/(1<<20))
	case size >= 1<<10:
		return fmt.Sprintf("%.2f KB", float64(size)/(1<<10))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func truncateSlice(values []string, limit int) []string {
	out := make([]string, 0, len(values))
	for _, v := range values {
		out = append(out, truncateString(v, limit))
	}
	return out
}

func truncateString(value string, limit int) string {
	if len(value) <= limit {
		return value
	}
	return value[:limit] + "...(truncated)"
}
