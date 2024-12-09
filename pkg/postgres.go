package pkg

import (
	"fmt"
	"log"
	"synapsis-online-store/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Koneksi gagal:", err)
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdletimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	return
}
