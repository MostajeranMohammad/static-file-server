package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/MostajeranMohammad/static-file-server/pkg/utils"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App               `yaml:"app"`
		HTTP              `yaml:"http"`
		Log               `yaml:"logger"`
		PG                `yaml:"postgres"`
		Minio             `yaml:"minio"`
		ImageOptimization `yaml:"image_optimization"`
	}

	// App -.
	App struct {
		Name    string `validate:"required" yaml:"name"    env:"APP_NAME"`
		Version string `validate:"required" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		JwtSecret string `validate:"required" yaml:"jwt_secret" env:"JWT_SECRET"`
		Port      string `validate:"required" yaml:"port"       env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `validate:"required" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		DSN string `validate:"required" yaml:"pg_dsn"   env:"PG_DSN"`
	}

	// Minio -.
	Minio struct {
		Endpoint        string `validate:"required" yaml:"endpoint"          env:"ENDPOINT"`
		AccessKeyID     string `validate:"required" yaml:"access_key_id"     env:"ACCESS_KEY_ID"`
		SecretAccessKey string `validate:"required" yaml:"secret_access_key" env:"SECRET_ACCESS_KEY"`
		UseSSL          bool   `yaml:"use_ssl"           env:"USE_SSL"`
	}

	// ImageOptimization -.
	ImageOptimization struct {
		LargeImageWidth     int `validate:"required" yaml:"large_image_width"     env:"LARGE_IMAGE_WIDTH"`
		MediumImageWidth    int `validate:"required" yaml:"medium_image_width"    env:"MEDIUM_IMAGE_WIDTH"`
		ThumbnailImageWidth int `validate:"required" yaml:"thumbnail_image_width" env:"THUMBNAIL_IMAGE_WIDTH"`
		CompressionQuality  int `validate:"required" yaml:"compression_quality"   env:"COMPRESSION_QUALITY"`
		ImageFormat         int `validate:"required" yaml:"image_format"          env:"IMAGE_FORMAT"`
		ImageMaxSize        int `validate:"required" yaml:"image_max_size"        env:"IMAGE_MAX_SIZE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := parseConfigFiles([]string{"./config/config.yml", "./.env"}, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateDto(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseConfigFiles(files []string, cfg *Config) error {
	for _, path := range files {
		if path == "./.env" {
			if _, err := os.Stat("/path/to/whatever"); errors.Is(err, os.ErrNotExist) {
				return nil
			}
		}
		err := cleanenv.ReadConfig(path, cfg)
		if err != nil {
			return fmt.Errorf("config error: %w", err)
		}
	}
	return nil
}
