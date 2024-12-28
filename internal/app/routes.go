package app

import (
	"github.com/danielgtaylor/huma/v2"

	userHTTP "github.com/gthomas08/realworld-huma/internal/domain/user/delivery/http"
	userRepository "github.com/gthomas08/realworld-huma/internal/domain/user/repository"
	userUsecase "github.com/gthomas08/realworld-huma/internal/domain/user/usecase"

	profileHTTP "github.com/gthomas08/realworld-huma/internal/domain/profile/delivery/http"
	profileRepository "github.com/gthomas08/realworld-huma/internal/domain/profile/repository"
	profileUsecase "github.com/gthomas08/realworld-huma/internal/domain/profile/usecase"

	articleHTTP "github.com/gthomas08/realworld-huma/internal/domain/article/delivery/http"
	articleRepository "github.com/gthomas08/realworld-huma/internal/domain/article/repository"
	articleUsecase "github.com/gthomas08/realworld-huma/internal/domain/article/usecase"
)

func (app *App) registerRoutes(api huma.API) {
	// User setup
	userRepo := userRepository.NewRepository(app.db)
	userUc := userUsecase.NewUsecase(app.cfg, app.logger, userRepo)
	userHandler := userHTTP.NewHandler(app.cfg, app.logger, userUc)
	userHandler.RegisterRoutes(api)

	// Profile setup
	profileRepo := profileRepository.NewRepository(app.db)
	profileUc := profileUsecase.NewUsecase(app.cfg, app.logger, profileRepo, userRepo)
	profileHandler := profileHTTP.NewHandler(app.cfg, app.logger, profileUc)
	profileHandler.RegisterRoutes(api)

	// Article setup
	articleRepo := articleRepository.NewRepository(app.db)
	articleUc := articleUsecase.NewUsecase(app.cfg, app.logger, articleRepo)
	articleHandler := articleHTTP.NewHandler(app.cfg, app.logger, articleUc)
	articleHandler.RegisterRoutes(api)
}
