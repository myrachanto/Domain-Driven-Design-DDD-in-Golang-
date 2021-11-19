package repository

import(
	"log"
	"fmt"
	"context"
	"github.com/spf13/viper"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/myrachanto/ddd/httperors"
	// "github.com/myrachanto/asokomonolith/support"
)
// Mongorepo
var (
	Mongorepo mongorepointerface = &mongorepo{}
	ctx = context.TODO()
	Host string = "" 
)
type Open struct {
	Mongohost string `mapstructure:"Mongohost"`
	MongodbName  string `mapstructure:"MongodbName"`
}
type mongorepointerface interface{
	Mongoclient()(*mongo.Client, *httperors.HttpError)
	Mongodb()(*mongo.Database,*mongo.Client, *httperors.HttpError)
	DbClose(client *mongo.Client)
	Gethost()(string,*httperors.HttpError)
	DBPing(p *mongo.Client)(string,*httperors.HttpError)
} 
type mongorepo struct{}
func LoadConfig() (open Open, err error) {
	viper.AddConfigPath("../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&open)
	return
}
func (m *mongorepo)Mongoclient()(*mongo.Client, *httperors.HttpError){
	m.Gethost()
	clientOptions := options.Client().ApplyURI(Host)
	client, err6 := mongo.Connect(ctx, clientOptions)
		if err6 != nil {
		return nil, httperors.NewBadRequestError("Could not connect to mongodb")
	}	
	return client, nil
}
func (m *mongorepo)Mongodb()(*mongo.Database,*mongo.Client, *httperors.HttpError){
	open, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	client, err1 := m.Mongoclient()
	db := client.Database(open.MongodbName)
	return db,client, err1
}
func (m *mongorepo)DbClose(client *mongo.Client){
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
func (m *mongorepo)Gethost()(string,*httperors.HttpError){
	open, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	host := open.Mongohost
	Host = host
	return host,nil
}
func (m *mongorepo)DBPing(p *mongo.Client)(string,*httperors.HttpError){
	err8 := p.Ping(ctx, nil)
	if err8 != nil {
		return "", httperors.NewBadRequestError("Could not ping the db")
	}
	return "Db connection was succesiful", nil
}