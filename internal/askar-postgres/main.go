package askar_postgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"time"
	"workout-training-api/internal/askar-postgres/model"
	"workout-training-api/internal/config"
)

type Postgres struct {
	*model.Model
}

func New(conf *config.AskarPostgresConfig, logger *slog.Logger) (*Postgres, error) {
	db, err := NewDatabase(conf)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	return &Postgres{
		Model: model.New(conf, logger.With(slog.String("module", "model")), db),
	}, nil
}

func NewDatabase(conf *config.AskarPostgresConfig) (*sql.DB, error) {
	//psqlInfo := fmt.Sprintf(
	//	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	conf.Host, conf.Port, conf.User, conf.Password, conf.Name,
	//)
	fmt.Println("new database")
	fmt.Println(conf.DSN)
	db, err := sql.Open("postgres", conf.DSN)
	if err != nil {
		fmt.Println("errror")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
