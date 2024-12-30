package entity

import "time"

const (
	REVIEW_COLLECTION = "review"
)

type Review struct {
	ID           string    `bson:"_id,omitempty" json:"-"`                         // ID dari MongoDB
	UserPublicID string    `bson:"user_public_id" json:"-"`                        // Referensi ke PostgreSQL
	Rating       int       `bson:"rating" json:"rating"`                           // Skor Rating
	ProductID    int       `bson:"product_id" json:"product_id" form:"product_id"` // Referensi ke PostgreSQL
	Review       string    `bson:"review" json:"review"`                           // Isi ulasan
	CreatedAt    time.Time `bson:"created_at" json:"-"`                            // Waktu ulasan dibuat
}

// func (r Review) GetCollection() string {
// 	return "review"
// }
