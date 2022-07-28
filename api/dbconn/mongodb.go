package dbconn

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/winnix/proreporting/api/config"
	"github.com/winnix/proreporting/api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbConn *mongo.Database

const (
	connectTimeout = 5 * time.Second
	writeTimeout   = 5 * time.Second
	readTimeout    = 3 * time.Second
)

func GetDB() *mongo.Database {
	return dbConn
}

func InitDB(cfg *config.Config, logger *logrus.Logger) (client *mongo.Client, ctx context.Context, cancel context.CancelFunc, err error) {
	ctx, cancel = context.WithTimeout(context.Background(), connectTimeout)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.DBURI))
	if err != nil {
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}

	dbConn = client.Database(cfg.DBName)

	err = initCollections()
	if err != nil {
		return
	}

	return
}

func initCollections() error {
	entities := []domain.Base{
		domain.UserToken{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), readTimeout)
	defer cancel()
	existing, err := dbConn.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return err
	}
	for _, entity := range entities {
		if checkIfExists(existing, entity.CollectionName()) {
			break
		}

		if err := createCollection(entity.CollectionName()); err != nil {
			return err
		}

		if err := createIndexes(entity); err != nil {
			return err
		}
	}

	return nil
}

func createCollection(cName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), writeTimeout)
	defer cancel()
	if err := dbConn.CreateCollection(ctx, cName); err != nil {
		return err
	}

	return nil
}

func createIndexes(model domain.Base) error {
	ctx, cancel := context.WithTimeout(context.Background(), writeTimeout)
	defer cancel()
	if _, err := dbConn.Collection(model.CollectionName()).Indexes().CreateMany(ctx, model.Indexes()); err != nil {
		return err
	}

	return nil
}

func checkIfExists(items []string, search string) bool {
	for _, item := range items {
		if item == search {
			return true
		}
	}

	return false
}
