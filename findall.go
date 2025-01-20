package mdbgocrud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mdb-go-crud/querybuilder"
)

func (r repositoryImpl[T]) FindAll(ctx context.Context, filter *querybuilder.Query) ([]*T, error) {
	newCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	var filters any
	if filter.RawQuery != nil {
		filters = filter.RawQuery
	} else {
		if filter.BatchFilters != nil {
			filters = filter.BatchFilters
		} else {
			if len(filter.Filters) > 0 {
				filters = bson.D{{"$and", filter.Filters}}
			} else {
				filters = bson.D{}
			}
		}
	}

	cursor, err := r.collection.Find(newCtx, filters, filter.Options)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		if cursor != nil {
			err = cursor.Close(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}(cursor, newCtx)

	if err != nil {
		return nil, err
	}

	var resp []*T

	for cursor.Next(newCtx) {
		var item *T

		err = cursor.Decode(&item)
		if err != nil {
			return nil, err
		}

		resp = append(resp, item)
	}

	return resp, nil
}
