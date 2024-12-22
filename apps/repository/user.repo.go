package repository

import (
	"context"
	"database/sql"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/pkg"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type RepoUser struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewRepoUser(db *sqlx.DB, client *redis.Client) *RepoUser {
	return &RepoUser{
		db:    db,
		redis: client,
	}
}

func (r *RepoUser) CreateUser(ctx context.Context, model entity.UserEntity) (err error) {
	query := `insert into users (
		email, password, role, public_id, created_at, updated_at
	) values (
	 	:email, :password, :role, :public_id, :created_at, :updated_at
	)`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	// karena setelah consume connection pool kita harus close statment tersebut
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, &model)
	return
}

func (r *RepoUser) GetUserByEmail(ctx context.Context, email string) (model entity.UserEntity, err error) {
	query := `
			SELECT 
				id, public_id, email, password, role, created_at, updated_at
			FROM users 
			where email = $1
	`

	err = r.db.GetContext(ctx, &model, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = pkg.ErrNotFound
			return
		}
		return
	}
	return
}

func (r *RepoUser) BlacklistToken(token string, expiration time.Duration) (err error) {
	return r.redis.Set(context.Background(), token, "blacklisted", expiration).Err()
}
