package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GZip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isNeedToDecompressRequest(*r) {
			decompressRequest(w, r)
		}

		if isNeedToCompressResponse(*r) {
			compressResponse(w, r, next)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isNeedToDecompressRequest(r http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Encoding"), "gzip")
}

func decompressRequest(w http.ResponseWriter, r *http.Request) {
	result, err := gzip.NewReader(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Body = result
}

func isNeedToCompressResponse(r http.Request) bool {
	return strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
}

func compressResponse(w http.ResponseWriter, r *http.Request, next http.Handler) {
	gz, err := gzip.NewWriterLevel(w, gzip.BestCompression)

	if err != nil {
		io.WriteString(w, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer gz.Close()

	w.Header().Set("Content-Encoding", "gzip")
	next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
}
