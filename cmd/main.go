package main

import (
	"context"
	"fmt"

	"github.com/BoburjonRaximov/premier_league/api"
	_ "github.com/BoburjonRaximov/premier_league/api/docs"
	"github.com/BoburjonRaximov/premier_league/api/handler"
	"github.com/BoburjonRaximov/premier_league/config"
	"github.com/BoburjonRaximov/premier_league/pkg/logger"
	"github.com/BoburjonRaximov/premier_league/storage/memory"
)

func main() {

	cfg := config.LoadP()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := memory.NewStorage(context.Background(), cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	h := handler.NewHandler(strg, *config.Load(), log)

	r := api.NewServer(h)
	r.Run(":8080")
}
