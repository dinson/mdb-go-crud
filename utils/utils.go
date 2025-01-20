package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

// StringToObjectID ... convert a string to mongo ObjectID
var StringToObjectID = func(objectIDStr string) (*primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(objectIDStr)
	if err != nil && !objectId.IsZero() {
		return nil, err
	}
	return &objectId, nil
}
