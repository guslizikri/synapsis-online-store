package services

import (
	"context"
	"errors"
	"synapsis-online-store/apps/entity"
	"time"
)

type RepositoryReview interface {
	RepoReviewIF
	RepoProductReviewIF
}

type RepoReviewIF interface {
	CreateReview(ctx context.Context, review entity.Review) error
	GetReviewsByProductID(ctx context.Context, productID int) ([]entity.Review, error)
}

type RepoProductReviewIF interface {
	ProductExistsByID(ctx context.Context, productID int) (bool, error)
}

type ServiceReview struct {
	repo RepositoryReview
}

func NewServiceReview(repo RepositoryReview) *ServiceReview {
	return &ServiceReview{
		repo: repo,
	}
}

func (s *ServiceReview) CreateReview(ctx context.Context, review entity.Review) error {
	// Validasi apakah product_id valid
	exists, err := s.repo.ProductExistsByID(ctx, review.ProductID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("invalid product_id")
	}

	// Tambahkan timestamp
	review.CreatedAt = time.Now()

	// Simpan review ke MongoDB
	return s.repo.CreateReview(ctx, review)
}

func (s *ServiceReview) GetReviews(ctx context.Context, productID int) ([]entity.Review, error) {
	return s.repo.GetReviewsByProductID(ctx, productID)
}
