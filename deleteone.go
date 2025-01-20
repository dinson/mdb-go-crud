package mdbgocrud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"mdb-go-crud/querybuilder"
)

func (r repositoryImpl[T]) DeleteOne(ctx context.Context, query *querybuilder.Query) error {
	newCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	filters := bson.D{{"$and", query.Filters}}

	_, err := r.collection.DeleteOne(newCtx, filters)
	if err != nil {
		return err
	}

	return nil
}
