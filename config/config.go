package config

import (
	"fmt"

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
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		JwtSecret string `env-required:"true" yaml:"jwt_secret" env:"JWT_SECRET"`
		Port      string `env-required:"true" yaml:"port"       env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		DSN string `env-required:"true" yaml:"pg_dsn"   env:"PG_DSN"`
	}

	// Minio -.
	Minio struct {
		Endpoint        string `env-required:"true" yaml:"endpoint"          env:"ENDPOINT"`
		AccessKeyID     string `env-required:"true" yaml:"access_key_id"     env:"ACCESS_KEY_ID"`
		SecretAccessKey string `env-required:"true" yaml:"secret_access_key" env:"SECRET_ACCESS_KEY"`
		UseSSL          bool   `yaml:"use_ssl"           env:"USE_SSL"`
	}

	// ImageOptimization -.
	ImageOptimization struct {
		LargeImageWidth     int `env-required:"true" yaml:"large_image_width"     env:"LARGE_IMAGE_WIDTH"`
		MediumImageWidth    int `env-required:"true" yaml:"medium_image_width"    env:"MEDIUM_IMAGE_WIDTH"`
		ThumbnailImageWidth int `env-required:"true" yaml:"thumbnail_image_width" env:"THUMBNAIL_IMAGE_WIDTH"`
		CompressionQuality  int `env-required:"true" yaml:"compression_quality"   env:"COMPRESSION_QUALITY"`
		ImageFormat         int `env-required:"true" yaml:"image_format"          env:"IMAGE_FORMAT"`
		ImageMaxSize        int `env-required:"true" yaml:"image_max_size"        env:"IMAGE_MAX_SIZE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
