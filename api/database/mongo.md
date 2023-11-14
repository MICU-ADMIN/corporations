package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// A Location Struct allows you to insert location documents into your
// collection

type Location struct {
	Name              string
	Coordinates       GeoJSON `json:"coordinates" bson:"coordinates"`
}

type GeoJSON struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}


func main() {

	// TODO:
	// Replace the placeholder connection string below with your
	// Atlas cluster specifics. Be sure it includes
	// a valid username and password! Note that in a production environment,
	// you do not want to store your password in plain-text here.
	var mongoUri = "mongodb+srv://admin:ea58.CPVW33pf7r@cluster0.aehzac9.mongodb.net/?retryWrites=true&w=majority"

	// CONNECT TO YOUR ATLAS CLUSTER:
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri,
	))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
	}

	fmt.Println("Connected to MongoDB!\n")

	// Provide the name of the database and collection you want to use.
	// If they don't already exist, the driver and Atlas will create them
	// automatically when you first write data.
var dbName = "myDatabase"
	var collectionName = "locations"
	collection := client.Database(dbName).Collection(collectionName)

	// Create a 2dsphere index on the "coordinates" field for geospatial queries
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"coordinates", "2dsphere"}},
	}

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		fmt.Println("Error creating index:", err)
		return
	}
	/*      *** INSERT DOCUMENTS ***
	 *
	 * You can insert individual documents using collection.Insert().
	 * In this example, we're going to create 4 documents and then
	 * insert them all in one call with InsertMany().
	 */

newYorkLocation := Location{
		Name:              "New York",
		Coordinates: GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-74.0060, 40.7128}, // New York City coordinates
		},
	}
	
losAngelesLocation := Location{
		Name:              "Los Angeles",
		Coordinates: GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-118.2437, 34.0522}, // Los Angeles coordinates
		},
	}

chicagoLocation := Location{
		Name:              "Chicago",
		Coordinates: GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-87.6298, 41.8781}, // Chicago coordinates
		},
	}

houstonLocation := Location{
		Name:              "Houston",
		Coordinates: GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-95.3698, 29.7604}, // Houston coordinates
		},
	}

	// Create an interface of all the created locations
	locations := []interface{}{newYorkLocation, losAngelesLocation, chicagoLocation, houstonLocation}
	insertManyResult, err := collection.InsertMany(context.TODO(), locations)
	if err != nil {
		fmt.Println("Something went wrong trying to insert the new documents:")
		panic(err)
	}

	fmt.Println(len(insertManyResult.InsertedIDs), "documents successfully inserted.\n")

	/*
	 * *** FIND DOCUMENTS ***
	 *
	 * Now that we have data in Atlas, we can read it. To retrieve all of
	 * the data in a collection, we create a filter for locations that are
	 * within a certain distance from a given point.
	 */

	 // Define the center point coordinates (replace with the desired coordinates)
centerPoint := GeoJSON{
	Type:        "Point",
	Coordinates: []float64{-74.006, 40.7128}, // New York City coordinates as an example
}

// Set up a geospatial query for locations within 1000 kilometers of the center point
geoQuery := bson.M{
	"coordinates": bson.M{
		"$nearSphere": bson.M{
			"$geometry":    centerPoint,
			"$maxDistance": 1000000, // 1000 kilometers in meters
		},
	},
}

// Perform the geospatial query
geoCursor, err := collection.Find(context.TODO(), geoQuery)
if err != nil {
	fmt.Println("Something went wrong trying to find locations near the center point:")
	panic(err)
}

defer func() {
	geoCursor.Close(context.Background())
}()

fmt.Println("Locations near the center point:")
for geoCursor.Next(ctx) {
	geoLocation := Location{}
	err := geoCursor.Decode(&geoLocation)
	if err != nil {
		fmt.Println("Error decoding geospatial query result:")
		panic(err)
	} else {
		fmt.Println(geoLocation.Name, "is near the specified location.")
	}
}

	/*      *** DELETE DOCUMENTS ***
	 *
	 *      As with other CRUD methods, you can delete a single document
	 *      or all documents that match a specified filter. To delete all
	 *      of the documents in a collection, pass an empty filter to
	 *      the DeleteMany() method. In this example, we'll delete two of
	 *      the locations.
	 */

	deletedLocationNameList := [...]string{"New York", "Los Angeles"}

	var deleteQuery = bson.M{"name": bson.M{"$in": deletedLocationNameList}}
	deleteResult, err := collection.DeleteMany(context.TODO(), deleteQuery)
	if err != nil {
		fmt.Println("Something went wrong trying to delete documents:")
		panic(err)
	}
	fmt.Println("Deleted", deleteResult.DeletedCount, "documents in the locations collection\n")

}
