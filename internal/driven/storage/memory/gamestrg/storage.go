package gamestrg

import (
	"context"
	"database/sql"
	"os"

	"github.com/ghazlabs/hex-mathrush/internal/core"
	//"github.com/ghazlabs/hex-mathrush/internal/driven/storage/memory/queststrg"
	_ "github.com/go-sql-driver/mysql"
)

type quest struct {
	id int
	to int
}
type Storage struct {
	gameMap map[string]core.Game
}

func New() *Storage {
	return &Storage{gameMap: map[string]core.Game{}}
}

func (s *Storage) PutGame(ctx context.Context, g core.Game) error {
	s.gameMap[g.GameID] = g
	return nil
}

func (s *Storage) GetGame(ctx context.Context, gameID string) (*core.Game, error) {
	game := s.gameMap[gameID]
	return &game, nil
}

func conn() *sql.DB {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	//db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp(db:"+db_port+")/hex_math")
	if err != nil {
		panic(err.Error())
	}
	return db
}
