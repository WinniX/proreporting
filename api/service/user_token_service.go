package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/winnix/proreporting/api/dbconn"
	"github.com/winnix/proreporting/api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserTokenService interface {
	GetCurrent(userId uuid.UUID) (*domain.UserToken, error)
	SaveUserToken(ut *domain.UserToken) error
}

type userTokenService struct {
	coll         *mongo.Collection
	accountCode  string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewUserTokenService(accountCode string) UserTokenService {
	return &userTokenService{
		coll:         dbconn.GetDB().Collection(domain.UserToken{}.CollectionName()),
		accountCode:  accountCode,
		readTimeout:  3 * time.Second,
		writeTimeout: 5 * time.Second,
	}
}

func (s *userTokenService) GetCurrent(userId uuid.UUID) (*domain.UserToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.readTimeout)
	defer cancel()

	item := &domain.UserToken{}
	err := s.coll.FindOne(ctx, bson.M{"account_code": s.accountCode, "user_uid": userId}).Decode(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *userTokenService) SaveUserToken(ut *domain.UserToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.writeTimeout)
	defer cancel()

	if ut.ID == primitive.NilObjectID {
		_, err := s.coll.InsertOne(ctx, ut)
		if err != nil {
			return fmt.Errorf("failed inserting user token: %s", err.Error())
		}

		return nil
	}

	_, err := s.coll.ReplaceOne(
		ctx,
		bson.M{"_id": ut.ID},
		bson.M{
			"access_token":  ut.AccessToken,
			"refresh_token": ut.RefreshToken,
			"expires_at":    ut.ExpiresAt,
		})
	if err != nil {
		return fmt.Errorf("failed updating user token: %s", err.Error())
	}

	return nil
}
