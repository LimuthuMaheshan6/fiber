package get

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
	"project/graph"
	"project/graph/model"

)
type queryResolver struct{ *graph.Resolver }

func Users(r *queryResolver, ctx context.Context) ([]*model.User, error) {
	if r.Mongo == nil {
		 fmt.Errorf("mongo client is nil — server misconfigured")
	}

	db := r.Mongo.Database("sample_mflix")
	collection := db.Collection("users")

	var users []*model.User

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursorA, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Mongo Find error:", err)
		fmt.Errorf("failed to fetch users: %w", err)
	}

	// Check cursorA is not nil
	if cursorA == nil {
		 fmt.Errorf("mongo cursor is nil")
	}

	defer func() {
		if err := cursorA.Close(ctx); err != nil {
			log.Println("Failed to close cursor:", err)
		}
	}()

	if err := cursorA.All(ctx, &users); err != nil {
		log.Println("Failed to decode users:", err)
		 fmt.Errorf("failed to decode users: %w", err)
	}

	return users, nil
}



