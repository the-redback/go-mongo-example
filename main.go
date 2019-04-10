package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Set client options
	fmt.Println("Client Options...")
	clientOptions := options.Client().ApplyURI("mongodb://root:WkwvYVe5YayWTDle@localhost:27017")

	// Connect to MongoDB
	fmt.Println("Connect mongodb...")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo Ping...")

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// List database ============================
	fmt.Println("List database names....")
	listDB, err := client.ListDatabaseNames(context.TODO(), bsonx.Doc{}, )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DBs:", listDB)

	// Insert into document =============================
	fmt.Println("Insert into Collections....")
	collection := client.Database("test").Collection("testcoll")
	res, err := collection.InsertOne(context.TODO(), bson.M{"myfield": "123", "name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println("id:", id, "response:", res)

	//Query with cursor ===================================
	fmt.Println("Query with cursor")
	//ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("result:", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// run command on db =======================================
	fmt.Println("Ping by command.....")
	db := client.Database("test")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsl := db.RunCommand(ctx, bson.D{{"ping", "1"}})
	if rsl.Err() != nil {
		log.Fatal(rsl.Err())
	}
	raw, err := rsl.DecodeBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">> ", raw)

	// run command on db: isMaster =======================================
	fmt.Println("run command.....: isMaster ")
	db = client.Database("test")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsl = db.RunCommand(ctx, bson.D{{"isMaster", "1"}})
	if rsl.Err() != nil {
		log.Fatal(rsl.Err())
	}
	raw, err = rsl.DecodeBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">> ", raw)

	//// run command on db: listCommands =======================================
	//fmt.Println("run command.....: listCommands ")
	//db = client.Database("test")
	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//rsl = db.RunCommand(ctx, bson.D{{"listCommands", "1"}})
	//if rsl.Err() != nil {
	//	log.Fatal(rsl.Err())
	//}
	//raw, err = rsl.DecodeBytes()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(">> ", raw)

	//Query with cursor: 'is test' db is partitioned ===================================
	fmt.Println("Query config db...")
	//ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	collection = client.Database("config").Collection("databases")
	curSingle := collection.FindOne(context.TODO(), bson.D{{"_id", "test2"}})
	if curSingle.Err() != nil {
		log.Fatal(curSingle.Err())
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	//raw , err = curSingle.DecodeBytes()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(">>",raw)

	// Use map ------------------------
	dec := make(map[string]interface{})
	err = curSingle.Decode(&dec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>",dec)

	val,ok := dec["partitioned"]
	if !ok {
		log.Fatal(err)
	}
	fmt.Println(">>",val)


	// run command on db: enable sharding "test3" =======================================
	fmt.Println("run command.....: enable sharding test3")
	db = client.Database("admin")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsl = db.RunCommand(ctx, bson.D{{"enableSharding", "test3"}})
	if rsl.Err() != nil {
		log.Fatal(rsl.Err())
	}
	raw, err = rsl.DecodeBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">> ", raw)

	// Now shardCollection
	fmt.Println("run command.....: shardCollection")
	db = client.Database("admin")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsl = db.RunCommand(ctx, bson.D{{"shardCollection", "test3.testcoll"},{"key",bson.M{"myfield": 1}}})
	if rsl.Err() != nil {
		log.Fatal(rsl.Err())
	}
	raw, err = rsl.DecodeBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">> ", raw)


	//Query with cursor:: All databases state ===================================
	fmt.Println("Query config db...: All databases state ")
	//ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	collection = client.Database("config").Collection("databases")
	cur, err = collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("result:", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
