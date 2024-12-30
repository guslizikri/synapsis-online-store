package repository

import (
	"context"
	"synapsis-online-store/apps/entity"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoReview struct {
	db    *mongo.Database
	dbSql *sqlx.DB
}

func NewRepoReview(db *mongo.Database, dbSql *sqlx.DB) *RepoReview {
	return &RepoReview{
		db:    db,
		dbSql: dbSql,
	}
}

func (r *RepoReview) GetReviewsByProductID(ctx context.Context, productID int) (model []entity.Review, err error) {
	filter := bson.M{"product_id": productID}
	collection := r.db.Collection(entity.REVIEW_COLLECTION)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &model); err != nil {
		return nil, err
	}
	return
}

func (r *RepoReview) CreateReview(ctx context.Context, review entity.Review) error {
	collection := r.db.Collection(entity.REVIEW_COLLECTION)
	_, err := collection.InsertOne(ctx, review)
	return err
}

func (p *RepoReview) ProductExistsByID(ctx context.Context, productID int) (bool, error) {
	var exists bool
	err := p.dbSql.GetContext(ctx, &exists, `SELECT EXISTS (SELECT 1 FROM products WHERE id = $1)`, productID)
	return exists, err
}
