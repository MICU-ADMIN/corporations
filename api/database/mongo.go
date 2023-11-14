package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Establishment represents an establishment document in the collection
type Establishment struct {
	Name     string
	Address  string
	Geo      string
	Amenities []string
	Capacity  int `json:"capacity" bson:"capacity"`
}

// ConnectToAtlasCluster connects to the MongoDB Atlas cluster
func ConnectToAtlasCluster() (*mongo.Client, error) {
	var mongoURI = "mongodb+srv://your_username:your_password@your_cluster.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!\n")
	return client, nil
}

// InsertDocuments inserts a list of establishments into the specified collection
func InsertDocuments(ctx context.Context, collection *mongo.Collection, establishments []Establishment) error {
	insertManyResult, err := collection.InsertMany(ctx, establishments)
	if err != nil {
		return err
	}

	fmt.Println(len(insertManyResult.InsertedIDs), "documents successfully inserted.\n")
	return nil
}

// FindDocuments retrieves establishments from the collection based on the provided filter and options
func FindDocuments(ctx context.Context, collection *mongo.Collection, filter bson.M, options *options.FindOptions) error {
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		establishment := Establishment{}
		err := cursor.Decode(&establishment)
		if err != nil {
			return err
		} else {
			fmt.Println(establishment.Name, "has", len(establishment.Amenities), "amenities, and takes", establishment.Capacity, "minutes to make.\n")
		}
	}

	return nil
}

// FindOneDocument retrieves a single establishment document from the collection based on the provided filter
func FindOneDocument(ctx context.Context, collection *mongo.Collection, filter bson.D) (Establishment, error) {
	var result Establishment
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return Establishment{}, err
	}

	fmt.Println("Found a document with the ingredient potato", result, "\n")
	return result, nil
}

// UpdateDocument updates a single document in the collection based on the provided filter and update document
func UpdateDocument(ctx context.Context, collection *mongo.Collection, filter bson.D, updateDoc bson.D) error {
	myRes := collection.FindOneAndUpdate(ctx, filter, updateDoc, nil)
	if myRes.Err() != nil {
		return myRes.Err()
	}

	_establishment := Establishment{}
	decodeErr := myRes.Decode(&_establishment)
	if decodeErr != nil {
		return decodeErr
	}

	updatedEstablishment, _ := json.MarshalIndent(_establishment, "", "\t")
	fmt.Println("The following document has been updated: \n", string(updatedEstablishment), "\n")

	return nil
}

// DeleteDocuments deletes documents from the collection based on the provided filter
func DeleteDocuments(ctx context.Context, collection *mongo.Collection, deleteQuery bson.M) error {
	deleteResult, err := collection.DeleteMany(ctx, deleteQuery)
	if err != nil {
		return err
	}

	fmt.Println("Deleted", deleteResult.DeletedCount, "documents in the establishments collection\n")
	return nil
}

func main() {
	// TODO: Replace the placeholder connection string below with your
	// Atlas cluster specifics. Be sure it includes
	// a valid username and password! Note that in a production environment,
	// you do not want to store your password in plain-text here.
	

	// CONNECT TO YOUR ATLAS CLUSTER:
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := ConnectToAtlasCluster(ctx, mongoURI)
	if err != nil {
		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Provide the name of the database and collection you want to use.
	// If they don't already exist, the driver and Atlas will create them
	// automatically when you first write data.
	var dbName = "myDatabase"
	var collectionName = "establishments"
	collection := client.Database(dbName).Collection(collectionName)

	// Create establishments
	eloteEstablishment := Establishment{
		Name:      "elote",
		Amenities: []string{"corn", "mayonnaise", "cotija cheese", "sour cream", "lime"},
		Capacity:  35,
	}

	locoMocoEstablishment := Establishment{
		Name:      "loco moco",
		Amenities: []string{"ground beef", "butter", "onion", "egg", "bread bun", "mushrooms"},
		Capacity:  54,
	}

	patatasBravasEstablishment := Establishment{
		Name:      "patatas bravas",
		Amenities: []string{"potato", "tomato", "olive oil", "onion", "garlic", "paprika"},
		Capacity:  80,
	}

	friedRiceEstablishment := Establishment{
		Name:      "fried rice",
		Amenities: []string{"rice", "soy sauce", "egg", "onion", "pea", "carrot", "sesame oil"},
		Capacity:  40,
	}

	// Create an interface of all the created establishments
	establishments := []interface{}{eloteEstablishment, locoMocoEstablishment, patatasBravasEstablishment, friedRiceEstablishment}

	// Insert documents
	err = InsertDocuments(ctx, collection, establishments)
	if err != nil {
		fmt.Println("Something went wrong trying to insert the new documents:")
		panic(err)
	}

	// Find documents
	var filter = bson.M{"capacity": bson.M{"$lt": 45}}
	options := options.Find()

	// Sort by `name` field ascending
	options.SetSort(bson.D{{"name", 1}})

	err = FindDocuments(ctx, collection, filter, options)
	if err != nil {
		fmt.Println("Something went wrong trying to find the documents:")
		panic(err)
	}

	// Find one document
	var myFilter = bson.D{{"amenities", "potato"}}
	result, e := FindOneDocument(ctx, collection, myFilter)
	if e != nil {
		fmt.Println("Something went wrong trying to find one document:")
		panic(e)
	}

	// Update a document
	var updateDoc = bson.D{{"$set", bson.D{{"capacity", 72}}}}
	err = UpdateDocument(ctx, collection, myFilter, updateDoc)
	if err != nil {
		fmt.Println("Something went wrong trying to update one document:")
		panic(err)
	}

	// Delete documents
	deletedEstablishmentNameList := [...]string{"elote", "fried rice"}

	var deleteQuery = bson.M{"name": bson.M{"$in": deletedEstablishmentNameList}}
	err = DeleteDocuments(ctx, collection, deleteQuery)
	if err != nil {
		fmt.Println("Something went wrong trying to delete documents:")
		panic(err)
	}
	fmt.Println("Deleted", len(deletedEstablishmentNameList), "documents in the establishments collection\n")
}