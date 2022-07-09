package main

import (
	"log"

	"github.com/Unlites/nba_api/config"
	gameHttp "github.com/Unlites/nba_api/internal/game/handler/http"
	gameRepository "github.com/Unlites/nba_api/internal/game/repository"
	gameUseCase "github.com/Unlites/nba_api/internal/game/usecase"
	playerHttp "github.com/Unlites/nba_api/internal/player/handler/http"
	playerRepository "github.com/Unlites/nba_api/internal/player/repository"
	playerUseCase "github.com/Unlites/nba_api/internal/player/usecase"
	"github.com/Unlites/nba_api/internal/server"
	statHttp "github.com/Unlites/nba_api/internal/stat/handler/http"
	statRepository "github.com/Unlites/nba_api/internal/stat/repository"
	statUseCase "github.com/Unlites/nba_api/internal/stat/usecase"
	teamHttp "github.com/Unlites/nba_api/internal/team/handler/http"
	teamRepository "github.com/Unlites/nba_api/internal/team/repository"
	teamUseCase "github.com/Unlites/nba_api/internal/team/usecase"
	"github.com/gin-gonic/gin"

	"github.com/Unlites/nba_api/pkg/db/postgres"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("failed with init config: %v", err)
	}

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	gameRepo := gameRepository.NewGameRepo(db)
	teamRepo := teamRepository.NewTeamRepo(db)
	playerRepo := playerRepository.NewPlayerRepo(db)
	statRepo := statRepository.NewStatRepo(db)

	gameUC := gameUseCase.NewGameUseCase(gameRepo)
	teamUC := teamUseCase.NewTeamUseCase(teamRepo)
	playerUC := playerUseCase.NewPlayerUseCase(playerRepo)
	statUC := statUseCase.NewStatUseCase(statRepo)

	gameHandler := gameHttp.NewGameHandler(gameUC)
	teamHandler := teamHttp.NewTeamHandler(teamUC)
	playerHandler := playerHttp.NewPlayerHandler(playerUC)
	statHandler := statHttp.NewStatHandler(statUC)

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	gameGroup := v1.Group("/game")
	teamGroup := v1.Group("/team")
	playerGroup := v1.Group("/player")
	statGroup := v1.Group("/stat")

	gameHttp.RegisterRoutes(gameGroup, gameHandler)
	teamHttp.RegisterRoutes(teamGroup, teamHandler)
	playerHttp.RegisterRoutes(playerGroup, playerHandler)
	statHttp.RegisterRoutes(statGroup, statHandler)

	server := server.NewServer(cfg, router)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
