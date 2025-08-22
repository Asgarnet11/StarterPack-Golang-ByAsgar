package config

import (
	"log"
	"os"
	"strconv"
)


type Config struct {
AppName string
Env string
HTTPPort int
LogLevel string
// Tambahkan koneksi DB jika perlu, contoh:
DatabaseURL string
}


func getEnv(key, def string) string {
if v := os.Getenv(key); v != "" { return v }
return def
}


func getEnvInt(key string, def int) int {
if v := os.Getenv(key); v != "" {
if i, err := strconv.Atoi(v); err == nil { return i }
}
return def
}


func Load() *Config {
// Muat .env kalau ada
_ = loadDotEnvOnce()


cfg := &Config{
AppName: getEnv("APP_NAME", "go-starter"),
Env: getEnv("APP_ENV", "development"),
HTTPPort: getEnvInt("HTTP_PORT", 8080),
LogLevel: getEnv("LOG_LEVEL", "info"),
DatabaseURL: getEnv("DATABASE_URL", ""),
}
return cfg
}


var dotenvLoaded bool


func loadDotEnvOnce() error {
if dotenvLoaded { return nil }
dotenvLoaded = true
if _, err := os.Stat(".env"); err == nil {
if err := godotenvLoad(); err != nil { // dipisah agar tidak jadi dep direkt
log.Println("warn: failed loading .env:", err)
}
}
return nil
}


// kecilkan import agar tetap rapih
func godotenvLoad() error { return godotenvLoadImpl() }