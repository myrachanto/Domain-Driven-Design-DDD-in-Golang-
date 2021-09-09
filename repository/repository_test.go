package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestGethost(t *testing.T){
	Mongorepo.gethost()	
	expected := "Mongohost"
	assert.Equal(t, expected, Host)
}
func TestDbConnectability(t *testing.T){
	client, _ := Mongorepo.Mongoclient()
    conn, err := Mongorepo.DBPing(client)
	expected := "Db connection was succesiful"
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, conn)

}
// func TestDbMongodb(t *testing.T){
// 	mongodb, client := Mongorepo.Mongodb()
// }
// func TestDbClose(t *testing.T){
// 	c, err := Mongorepo.Mongodb()
// 	 Mongorepo.DbClose(c)
// }