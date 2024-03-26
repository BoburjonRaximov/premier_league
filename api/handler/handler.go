package handler

import (
	"github.com/BoburjonRaximov/premier_league/config"
	"github.com/BoburjonRaximov/premier_league/pkg/logger"
	"github.com/BoburjonRaximov/premier_league/storage"
)

type Handler struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
}

func NewHandler(strg storage.StorageI, cfg config.Config, log logger.LoggerI) *Handler {
	return &Handler{strg: strg, cfg: cfg, log: log}
}
