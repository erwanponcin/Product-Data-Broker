package plugins

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
)

type MeltingData struct {
    BatchID        string `bson:"batch_id"`
    Material       string `bson:"material"`
    Temperature    int    `bson:"temperature"`
    ProductionDate string `bson:"production_date"`
}

type MongoPlugin struct {
    client     *mongo.Client
    collection *mongo.Collection
}

// Initializes a new MongoDB plugin
func NewMongoPlugin(uri, database, collection string) (*MongoPlugin, error) {
    clientOpts := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.TODO(), clientOpts)
    if err != nil {
        return nil, fmt.Errorf("could not connect to MongoDB: %v", err)
    }

    coll := client.Database(database).Collection(collection)
    return &MongoPlugin{
        client:     client,
        collection: coll,
    }, nil
}

// Fetches all data from the MongoDB collection
func (mp *MongoPlugin) FetchAllData() ([]MeltingData, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := mp.collection.Find(ctx, map[string]interface{}{})
    if err != nil {
        return nil, fmt.Errorf("could not fetch data: %v", err)
    }
    defer cursor.Close(ctx)

    var results []MeltingData
    if err = cursor.All(ctx, &results); err != nil {
        return nil, fmt.Errorf("could not decode data: %v", err)
    }

    return results, nil
}
