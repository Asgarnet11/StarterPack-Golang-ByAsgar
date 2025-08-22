package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)


func RequestLogger(next http.Handler) http.Handler {
return hlog.NewHandler(log.Logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
start := time.Now()
next.ServeHTTP(w, r)
hlog.FromRequest(r).Info().
Str("method", r.Method).
Str("path", r.URL.Path).
Dur("duration", time.Since(start)).
Msg("request")
}))
}