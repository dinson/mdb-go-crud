package mdbgocrud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mdb-go-crud/querybuilder"
	"time"
)

const (
	connectionTimeout = 15 * time.Second
)

type Repository[T any] interface {
	Save(ctx context.Context, entity *T, ID *string) (*primitive.ObjectID, error)
	InsertMany(ctx context.Context, docs []*T) ([]*primitive.ObjectID, error)
	FindAll(ctx context.Context, filter *querybuilder.Query) ([]*T, error)
	FindOne(ctx context.Context, filter *querybuilder.Query) (*T, error)
	DeleteOne(ctx context.Context, query *querybuilder.Query) error
	DeleteMany(ctx context.Context, query *querybuilder.Query) error
}

type repositoryImpl[T any] struct {
	collection *mongo.Collection
}

func NewRepository[T any](collection *mongo.Collection) Repository[T] {
	return &repositoryImpl[T]{
		collection: collection,
	}
}
