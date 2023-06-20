package usecase

import (
	"github.com/MostajeranMohammad/static-file-server/config"
	"github.com/MostajeranMohammad/static-file-server/internal/contracts/consts"
	"github.com/h2non/bimg"
)

type imageOptimizerUseCase struct {
	cfg config.Config
}

func NewImageOptimizerUseCase(cfg config.Config) ImageOptimizer {
	return &imageOptimizerUseCase{cfg}
}

func (optimizer imageOptimizerUseCase) CheckImageNeedsOptimization(buffer []byte, imageSize int) (bool, error) {
	var width int
	switch imageSize {
	case consts.LargeImage:
		width = optimizer.cfg.ImageOptimization.LargeImageWidth
	case consts.ThumbnailImage:
		width = optimizer.cfg.ImageOptimization.ThumbnailImageWidth
	case consts.MediumImage:
		width = optimizer.cfg.ImageOptimization.MediumImageWidth
	default:
		width = optimizer.cfg.ImageOptimization.MediumImageWidth
	}

	bimgImg := bimg.NewImage(buffer)

	imageType := bimgImg.Type()

	if imageType == "svg" {
		return false, nil
	}

	imgSize, err := bimgImg.Size()
	if err != nil {
		return false, err
	}

	if imgSize.Width > width || len(buffer) > optimizer.cfg.ImageOptimization.ImageMaxSize {
		return true, nil
	} else {
		return false, nil
	}
}

func (optimizer imageOptimizerUseCase) OptimizeImage(buffer []byte, imageSize int, outputFileType int) ([]byte, error) {
	var width int
	switch imageSize {
	case consts.LargeImage:
		width = optimizer.cfg.ImageOptimization.LargeImageWidth
	case consts.ThumbnailImage:
		width = optimizer.cfg.ImageOptimization.ThumbnailImageWidth
	case consts.MediumImage:
		width = optimizer.cfg.ImageOptimization.MediumImageWidth
	default:
		width = optimizer.cfg.ImageOptimization.MediumImageWidth
	}

	var fileType bimg.ImageType
	switch outputFileType {
	case consts.JPG:
		fileType = bimg.JPEG
	case consts.PNG:
		fileType = bimg.PNG
	case consts.WEBP:
		fileType = bimg.WEBP
	default:
		fileType = bimg.WEBP
	}

	bimgImg := bimg.NewImage(buffer)

	newImage, err := bimgImg.Process(bimg.Options{
		Width:   width,
		Quality: optimizer.cfg.ImageOptimization.CompressionQuality,
		Type:    fileType,
	})
	if err != nil {
		return []byte{}, err
	}
	return newImage, nil
}
