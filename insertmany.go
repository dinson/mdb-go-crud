package mdbgocrud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r repositoryImpl[T]) InsertMany(ctx context.Context, docs []*T) ([]*primitive.ObjectID, error) {
	newCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	var docInterfaceList []any

	for _, d := range docs {
		docInterfaceList = append(docInterfaceList, d)
	}

	res, err := r.collection.InsertMany(newCtx, docInterfaceList)
	if err != nil {
		return nil, err
	}

	insertedIDs := res.InsertedIDs

	var primitiveIDs []*primitive.ObjectID

	for _, i := range insertedIDs {
		if i != nil {
			primitiveOID := i.(primitive.ObjectID)
			primitiveIDs = append(primitiveIDs, &primitiveOID)
		}
	}

	return primitiveIDs, nil
}
