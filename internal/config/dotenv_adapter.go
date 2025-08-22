package config

import "github.com/joho/godotenv"


func godotenvLoadImpl() error { return godotenv.Load() }