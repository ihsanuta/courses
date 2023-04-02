package config

import (
	"os"

	"github.com/joho/godotenv"
)

var AppConfig, MysqlConfig, JwtConfig, Cloudinary map[string]interface{}

func init() {
	godotenv.Load()
	AppConfig = map[string]interface{}{
		"port": os.Getenv("SERVER_PORT"),
		"host": os.Getenv("SERVER_ADDRESS"),
	}

	MysqlConfig = map[string]interface{}{
		"username": os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASS"),
		"database": os.Getenv("DB_NAME"),
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
	}

	JwtConfig = map[string]interface{}{
		"application":   os.Getenv("APPLICATION_NAME"),
		"expiration":    os.Getenv("LOGIN_EXPIRATION_DURATION"),
		"jwt_method":    os.Getenv("JWT_SIGNING_METHOD"),
		"jwt_signature": os.Getenv("JWT_SIGNATURE_KEY"),
	}

	Cloudinary = map[string]interface{}{
		"cloud_name":    os.Getenv("CLOUDINARY_CLOUD_NAME"),
		"api_key":       os.Getenv("CLOUDINARY_API_KEY"),
		"api_secret":    os.Getenv("CLOUDINARY_API_SECRET"),
		"upload_folder": os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),
	}
}
