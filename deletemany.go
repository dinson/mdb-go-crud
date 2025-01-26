package mongokit

import (
	"context"
	"github.com/dinson/mongokit/querybuilder"
	"go.mongodb.org/mongo-driver/bson"
)

func (r repositoryImpl[T]) DeleteMany(ctx context.Context, query *querybuilder.Query) error {
	newCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	filters := bson.D{{"$and", query.Filters}}

	_, err := r.collection.DeleteMany(newCtx, filters)
	if err != nil {
		return err
	}

	return nil
}
