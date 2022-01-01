package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	User     int                `json:"user,omitempty" bson:"user,omitempty"`
	Reason   string             `json:"reason,omitempty" bson:"reason,omitempty"`
	Enforcer int                `json:"enforcer" bson:"enforcer"`
	Message  string             `json:"message,omitempty" bson:"message,omitempty"`
	Blacklisted   bool          `json:"blacklisted,omitempty" bson:"blacklisted,omitempty"`
}
