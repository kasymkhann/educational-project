package DB

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"REST/internal/users"
	"REST/pkg/logging"
)

type Storage struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(ctx context.Context, collection *mongo.Collection, logger *logging.Logger) *Storage {
	return &Storage{collection: collection, logger: logger}
}

func (s *Storage) Create(ctx context.Context, user users.Users) (string, error) {
	s.logger.Debug("create user")

	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("not be created user: %v", err)
	}

	s.logger.Debug("convert insertId to objectId")

	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}
	s.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectId to hex. probably to objectId: %s ", objectId)
}

func (s *Storage) FindAll(ctx context.Context) (u []users.Users, err error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, fmt.Errorf("filed to find all users due to error: %v", err)
	}

	if err = cursor.All(ctx, &u); err != nil {
		return u, fmt.Errorf("failed to read all document from cursor error:  %v", err)
	}
	return u, nil
}

func (s *Storage) Find(ctx context.Context, id string) (u users.Users, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectId. probably to hex: %s", id)
	}

	filter := bson.M{"_id": objectId}

	result := s.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			// TODO ErrEntityNotFound
			return u, fmt.Errorf("not found")
		}
		return u, fmt.Errorf("filed to find one user by id %s due to error: %v", id, err)

	}

	// что делаеть decode
	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user(id:%s) from DB due to error: %v", id, err)
	}
	return u, nil
}

func (s *Storage) Update(ctx context.Context, users users.Users) error {
	objectId, err := primitive.ObjectIDFromHex(users.Id)
	if err != nil {
		return fmt.Errorf("failed to convert user id to objectId. id=%s", users.Id)
	}
	filter := bson.M{"_id": objectId}

	b, err := bson.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshal user. error: %v", err)
	}

	var updateUsersM bson.M
	err = bson.Unmarshal(b, &updateUsersM)
	if err != nil {
		return fmt.Errorf("failed unmarshal user bytes error: %v", err)
	}

	delete(updateUsersM, "_id")

	resultM := bson.M{
		"$set": updateUsersM,
	}

	fuResult, err := s.collection.UpdateOne(ctx, filter, resultM)
	if err != nil {
		return fmt.Errorf("failed to execute update user query error: %v", err)
	}

	if fuResult.MatchedCount == 0 {
		//TODO: 404
		return fmt.Errorf("not found")
	}

	s.logger.Tracef("Mathed %s documents and Modified %s documents", fuResult.MatchedCount, fuResult.ModifiedCount)

	return nil

}

func (s *Storage) Delete(ctx context.Context, users users.Users) error {
	objectId, err := primitive.ObjectIDFromHex(users.Id)
	if err != nil {
		return fmt.Errorf("failed to convert user id to objectId. id=%s", users.Id)
	}
	filter := bson.M{"_id": objectId}

	result, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query error: %v", err)
	}
	if result.DeletedCount == 0 {
		//    TODO 404
		return fmt.Errorf("not found")
	}
	s.logger.Tracef("Delete %s documents", result.DeletedCount)

	return nil

}
