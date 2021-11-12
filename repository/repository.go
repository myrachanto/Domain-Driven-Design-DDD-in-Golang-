package repository

import(
	"os"
	"log"
	"fmt"
	"context"
	"github.com/joho/godotenv"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/myrachanto/ddd/httperors"
	// "github.com/myrachanto/asokomonolith/support"
)
var (
	Mongorepo mongorepointerface = &mongorepo{}
	ctx = context.TODO()
	Host string = ""
)
type mongorepointerface interface{
	Mongoclient()(*mongo.Client, *httperors.HttpError)
	Mongodb()(*mongo.Database,*mongo.Client, *httperors.HttpError)
	DbClose(client *mongo.Client)
	gethost()*httperors.HttpError
	DBPing(p *mongo.Client)(string,*httperors.HttpError)
}
type mongorepo struct{}

func (m *mongorepo)Mongoclient()(*mongo.Client, *httperors.HttpError){
	m.gethost()
	clientOptions := options.Client().ApplyURI(Host)
	client, err6 := mongo.Connect(ctx, clientOptions)
		if err6 != nil {
		return nil, httperors.NewBadRequestError("Could not connect to mongodb")
	}	
	return client, nil
}
func (m *mongorepo)Mongodb()(*mongo.Database,*mongo.Client, *httperors.HttpError){
	err7 := godotenv.Load()
	if err7 != nil {
		return nil, nil, httperors.NewBadRequestError("error loading env file")
	}
	mongodb := os.Getenv("MongodbName")
	client, err1 := m.Mongoclient()
	db := client.Database(mongodb)
	return db,client, err1
}
func (m *mongorepo)DbClose(client *mongo.Client){
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
func (m *mongorepo)gethost()*httperors.HttpError{
	err := godotenv.Load()
	if err != nil {
		return  httperors.NewBadRequestError("error loading env file")
	}
	host := os.Getenv("Mongohost")
	Host = host
	return nil
}
func (m *mongorepo)DBPing(p *mongo.Client)(string,*httperors.HttpError){
	err8 := p.Ping(ctx, nil)
	if err8 != nil {
		return "", httperors.NewBadRequestError("Could not ping the db")
	}
	return "Db connection was succesiful", nil
}