package postgres

// nolint
import (
	"context"
	"fmt"
	"github.com/ikuyotagan/movier/internal/config"
	"github.com/ikuyotagan/movier/pkg/sre/log"
	"github.com/ra9dev/shutdown"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

const (
	checkCount   = 120
	maxIdleConns = 2
	maxOpenConns = 20
)

func newDB(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	// Создаем подключение к базе
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = maxOpenConns
	cfg.MaxConnIdleTime = time.Duration(maxOpenConns) * time.Second

	db, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	dbAvailable := false
	for i := 0; i < checkCount; i++ {
		if err := db.Ping(ctx); err != nil {
			log.Warnf(ctx, "Database not available ( %v ), connection try %d times", err, i+1)
			<-time.After(time.Second)
		} else {
			dbAvailable = true
			break
		}
	}

	if !dbAvailable {
		err := fmt.Errorf("database not available in %d seconds after run app", checkCount)
		return nil, err
	}

	log.Info(ctx, "Connected to DB!")
	return db, nil
}

// NewClient ...
func NewClient(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	db, err := newDB(ctx, cfg.DataStore.URL)
	if err != nil {
		return nil, err
	}

	// Закрываем соединения с базой при завершении работы приложения
	return db, shutdown.Add("database", func(ctx context.Context) {
		db.Close()
	})
}
