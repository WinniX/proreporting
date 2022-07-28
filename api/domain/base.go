package domain

import "go.mongodb.org/mongo-driver/mongo"

type Base interface {
	CollectionName() string
	Indexes() []mongo.IndexModel
}
