package main

import (
	"fmt"
	"net/http"
	internalConfig "starter-golang/internal/config"
	internalHTTP "starter-golang/internal/http"
	"starter-golang/internal/logger"
)


func main() {
cfg := internalConfig.Load()
log := logger.New(cfg.LogLevel)
router := internalHTTP.NewRouter()


addr := fmt.Sprintf(":%d", cfg.HTTPPort)
log.Info().Str("addr", addr).Msg("server starting")
if err := http.ListenAndServe(addr, router); err != nil {
log.Fatal().Err(err).Msg("server failed")
}
}