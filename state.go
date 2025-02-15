package main

import (
	"github.com/hiabhi-cpu/rssagg/internal/config"
	"github.com/hiabhi-cpu/rssagg/internal/database"
)

type state struct {
	db  *database.Queries
	con *config.Config
}
