package application

import (
	"fmt"
	"log"

	"github.com/MostajeranMohammad/static-file-server/config"
	_ "github.com/MostajeranMohammad/static-file-server/docs"
	"github.com/MostajeranMohammad/static-file-server/internal/controller"
	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/MostajeranMohammad/static-file-server/internal/repo"
	"github.com/MostajeranMohammad/static-file-server/internal/routes"
	"github.com/MostajeranMohammad/static-file-server/internal/usecase"
	"github.com/MostajeranMohammad/static-file-server/pkg/guards"
	"github.com/MostajeranMohammad/static-file-server/pkg/logger"
	"github.com/MostajeranMohammad/static-file-server/pkg/utils"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Run(cfg config.Config) {
	l, zL := logger.New(cfg.Log.Level)

	// Repository
	dsn := cfg.PG.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("failed to extract sqlDB")
	}

	defer func() {
		err = sqlDB.Close()
		if err != nil {
			log.Fatalln("failed to close sqlDB")
		}
	}()

	// Initialize minio client object.
	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		panic(err)
	}

	// auto migrate models
	db.AutoMigrate(&entity.StaticFileMetaData{})

	// initialize repos
	fileMetaDataRepo := repo.NewStaticFileMetaRepo(db)

	// initialize usecases
	fileMetaDataUsecase := usecase.NewStaticFileMetaManagerUsecase(fileMetaDataRepo)
	imageOptimizerUseCase := usecase.NewImageOptimizerUseCase(cfg)
	objectStorageManagerUseCase := usecase.NewObjectStorageManagerUseCase(minioClient)
	err = objectStorageManagerUseCase.InitBuckets()
	if err != nil {
		l.Warn(err.Error(), err)
		err = nil
	}
	staticFileManagerUsecase := usecase.NewStaticFileManagerUsecase(objectStorageManagerUseCase, fileMetaDataUsecase, imageOptimizerUseCase)

	// initialize controllers
	fileMetaDataController := controller.NewFileMetaDataController(fileMetaDataUsecase, staticFileManagerUsecase)
	staticFileManagerController := controller.NewStaticFileController(staticFileManagerUsecase)

	// initialize guards
	jwtGuard := guards.NewJWTGuard(cfg.JwtSecret)

	// HTTP Server
	httpApp := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: utils.FiberErrorHandler(l),
	})
	httpApp.Use(fiberzerolog.New(fiberzerolog.Config{Logger: zL}))
	httpApp.Use(recover.New())
	httpApp.Get("/swagger/*", swagger.HandlerDefault)
	httpApp.Mount("/static-file", routes.NewStaticFileRouter(staticFileManagerController, jwtGuard))
	httpApp.Mount("/file-meta-data", routes.NewFileMetaDataRoutes(fileMetaDataController, jwtGuard))

	httpApp.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port))
}
