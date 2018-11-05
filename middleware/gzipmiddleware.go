package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// GZip Middleware
func GZipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodings := r.Header.Get("Accept-Encoding")
		if !strings.Contains(encodings, "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Add("Content-Encoding", "gzip")
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()
		var rw http.ResponseWriter
		if pusher, ok := w.(http.Pusher); ok {
			rw = gzipPusherResponseWriter{
				gzipResponseWriter: gzipResponseWriter{
					ResponseWriter: w,
					Writer:         gzipWriter,
				},
				Pusher: pusher,
			}
		} else {
			rw = gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gzipWriter,
			}
		}
		next.ServeHTTP(rw, r)
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

type gzipPusherResponseWriter struct {
	gzipResponseWriter
	http.Pusher
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}
