package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/myrachanto/ddd/httperors"
)
func TestGethost(t *testing.T){
	host, err := Mongorepo.Gethost()	
	expected := "mongodb://localhost:27017"
	fmt.Println(">>>>>>>>>>>>", err)
	// assert.Equal(t, *httperors.HttpError, err, "test passed the error is equal to nill")
	assert.Equal(t, expected, host,"test passed the host is equal to Mongohost")
}
func TestDbConnectability(t *testing.T){
	client, _ := Mongorepo.Mongoclient()
    conn, _ := Mongorepo.DBPing(client)
	expected := "Db connection was succesiful"
	// assert.Equal(t, nil, err)
	assert.Equal(t, expected, conn,"test passed the connection to mongo db passed")

}
// func TestDbMongodb(t *testing.T){
// 	mongodb, client := Mongorepo.Mongodb()
// }
// func TestDbClose(t *testing.T){
// 	c, err := Mongorepo.Mongodb()
// 	 Mongorepo.DbClose(c)
// }