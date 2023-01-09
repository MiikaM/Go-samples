package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

var instance MongoInstance

var DB_NAME = os.Getenv("DATABASE_NAME")
var MONGOURI = os.Getenv("MONGOURI")
var MONGO_COLLECTION = os.Getenv("MONGO_COLLECTION")
var PORT = os.Getenv("PORT")

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGOURI))
	if err != nil {
		panic(err)
	}

	db := client.Database(DB_NAME)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	instance = MongoInstance{
		Client: client,
		Db:     db,
	}

	fmt.Println("Successfully connected and pinged.")
	app := fiber.New()

	app.Get("/employee", GetEmployee)

	app.Post("/employee", CreateEmployee)
	app.Put("/employee/:id", UpdateEmployee)
	app.Delete("/employee/:id", DeleteEmployee)

	log.Fatal(app.Listen(":" + PORT))
}

func GetEmployee(context *fiber.Ctx) error {
	query := bson.D{{}}

	cursor, err := instance.Db.Collection(MONGO_COLLECTION).Find(context.Context(), query)
	if err != nil {
		return context.Status(400).SendString(err.Error())
	}

	var employees []Employee = make([]Employee, 0)

	if err := cursor.All(context.Context(), &employees); err != nil {
		return context.Status(400).SendString(err.Error())

	}

	return context.JSON(employees)
}

func CreateEmployee(context *fiber.Ctx) error {
	collection := instance.Db.Collection(MONGO_COLLECTION)

	employee := new(Employee)

	if err := context.BodyParser(employee); err != nil {
		return context.Status(400).SendString(err.Error())
	}

	employee.ID = ""

	result, err := collection.InsertOne(context.Context(), employee)

	if err = context.BodyParser(employee); err != nil {
		return context.Status(400).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	createdRecord := collection.FindOne(context.Context(), filter)

	createdEmployee := &Employee{}
	createdRecord.Decode(createdEmployee)

	return context.Status(201).JSON(createdEmployee)
}

func UpdateEmployee(context *fiber.Ctx) error {
	id := context.Params("id")

	employeeId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return context.SendStatus(400)
	}

	employee := new(Employee)

	if err := context.BodyParser(employee); err != nil {
		return context.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeId}}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err = instance.Db.Collection(MONGO_COLLECTION).FindOneAndUpdate(context.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return context.SendStatus(400)
		}

		return context.SendStatus(500)
	}
	employee.ID = id

	return context.Status(200).JSON(employee)
}

func DeleteEmployee(context *fiber.Ctx) error {
	employeeId, err := primitive.ObjectIDFromHex(context.Params("id"))

	if err != nil {
		return context.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: employeeId}}

	result, err := instance.Db.Collection(MONGO_COLLECTION).DeleteOne(context.Context(), &query)

	if err != nil {
		return context.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return context.SendStatus(404)
	}

	return context.Status(200).JSON("Record deleted")
}
