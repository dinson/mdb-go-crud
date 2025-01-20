package mdbgocrud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mdb-go-crud/utils"
)

func (r repositoryImpl[T]) Save(ctx context.Context, entity *T, ID *string) (*primitive.ObjectID, error) {
	newCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	objectID := primitive.NewObjectID()
	if ID != nil {
		oID, err := utils.StringToObjectID(*ID)
		if err != nil {
			return nil, err
		}
		objectID = *oID
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", objectID}}

	res, err := r.collection.UpdateOne(newCtx, filter, bson.D{{"$set", entity}}, opts)
	if err != nil {
		return nil, err
	}

	oid := objectID
	if res.UpsertedID != nil {
		oid = res.UpsertedID.(primitive.ObjectID)
	}

	return &oid, nil
}
