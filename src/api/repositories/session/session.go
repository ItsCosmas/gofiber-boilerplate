package session

import (
	"context"
	"errors"
	db "gofiber-boilerplate/api/database"
	"gofiber-boilerplate/api/services/auth"
	"time"
)

// TokenMeta represents Token Metadata passed
type TokenMeta struct {
	TokenUUID string
	UserID    string
}

// TokenDetails Represents Access and Refresh Token information
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtUUID       string
	RtUUID       string
	AtExpires    int64
	RtExpires    int64
}

// SaveToken Saves the Token to Redis
func SaveToken(userID string, rt *auth.TokenDetails) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r := time.Unix(rt.TokenExpires, 0)
	now := time.Now()

	rtCreated, err := db.RdDB.Set(ctx, rt.TokenUUID, userID, r.Sub(now)).Result()
	if err != nil {
		return err
	}
	if rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

//FetchAuth Checks the metadata saved
func FetchAuth(tokenUUID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userID, err := db.RdDB.Get(ctx, tokenUUID).Result()
	if err != nil {
		return "", err
	}
	return userID, nil
}

//DeleteTokens deletes token from specified metadata
func DeleteTokens(atkm *TokenMeta, rtkm *TokenMeta) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//delete access token
	deletedAt, err := db.RdDB.Del(ctx, atkm.TokenUUID).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := db.RdDB.Del(ctx, atkm.TokenUUID).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

// DeleteRefresh deletes refresh token
func DeleteRefresh(refreshUUID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	//delete refresh token
	deleted, err := db.RdDB.Del(ctx, refreshUUID).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
