package repository

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/myrachanto/ddd/httperors"
)
func TestGethost(t *testing.T){
	host, err := Mongorepo.Gethost()	
	// expected := "mongodb://localhost:27017"
	// fmt.Println(">>>>>>>>>>>>", err)
	assert.Nil(t, err, "the error is nil ")
	assert.NotNil(t, host,"test passed the host is equal to Mongohost")
}
func TestDbConnectability(t *testing.T){
	client, _ := Mongorepo.Mongoclient()
    conn, _ := Mongorepo.DBPing(client)
	// expected := "Db connection was succesiful"
	// assert.Nil(t, err, "the error is nil ")
	assert.NotNil(t, conn,"test passed the connection to mongo db passed")

}

// /more testing
// func TestDbMongodb(t *testing.T){
// 	mongodb, client := Mongorepo.Mongodb()
// }
// func TestDbClose(t *testing.T){
// 	c, err := Mongorepo.Mongodb()
// 	 Mongorepo.DbClose(c)
// }