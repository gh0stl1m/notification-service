package mongodb

import (
	"context"
	"time"

	"github.com/gh0stl1m/notification-service/configs"
	"github.com/gh0stl1m/notification-service/domains"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const EVENTS_COLLECTION = "events"

var env *configs.MongoDBConfig = configs.ReadMongoDBConfig()

type EventsRepository struct {
	conn *mongo.Client
}

func New(db *mongo.Client) domains.EventModel {
	return &EventsRepository{
		conn: db,
	}
}

func (er *EventsRepository) Create(event domains.Event) error {

	ctxOperation, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := er.conn.Database(env.Database).Collection(EVENTS_COLLECTION)

	res, err := collection.InsertOne(ctxOperation, event)

	if err != nil {

		log.Errorf("Error inserting document in %s with error %s", EVENTS_COLLECTION, err)

		return err
	}

	log.Debugf("Inserted document in %s with id %s", EVENTS_COLLECTION, res.InsertedID)

	return nil
}

func (er *EventsRepository) GetById(id int64) (domains.Event, error) {

	ctxOperation, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := er.conn.Database(env.Database).Collection(EVENTS_COLLECTION)

	var result domains.Event
	err := collection.FindOne(ctxOperation, bson.M{"_id": id}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Errorf("No documents found with the given id: %v", err)
			return domains.Event{}, err
		}
		log.Errorf("Error finding event: %v", err)
		return domains.Event{}, err
	}

	return result, nil
}

func (er *EventsRepository) UpdateStateById(id int64, state string) error {

	ctxOperation, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := er.conn.Database(env.Database).Collection(EVENTS_COLLECTION)

	newEventState := domains.EventsHistory{
		State:     state,
		CreatedAt: time.Now(),
	}

	_, err := collection.UpdateByID(ctxOperation, id, bson.M{"$set": newEventState})

	if err != nil {
		log.Errorf("Error updating event: %s with err %v", id, err)
		return err
	}

	return nil
}
