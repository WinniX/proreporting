package domain

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserToken struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserUID      uuid.UUID          `bson:"user_uid" json:"user_uid"`
	AccountCode  string             `bson:"account_code" json:"account_code"`
	AccessToken  string             `bson:"access_token" json:"access_token"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	ExpiresAt    time.Time          `bson:"expires_at" json:"expires_at"`
}

func (ut UserToken) CollectionName() string {
	return "user_tokens"
}

func (ut UserToken) Indexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "user_uid", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
