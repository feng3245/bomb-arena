package userManagement

import (
	"log"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct{
	UserName []byte
	Password []byte

}

func Save(u *User) (*mgo.Session, error){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
        panic(err)
    }
    defer session.Close()
session.SetMode(mgo.Monotonic, true)	
	c := session.DB("test").C("users")
        return session, c.Insert(u)
}
func Load(uname []byte) (*User){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Printf("Connection failure")
        panic(err)
    }
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	result := User{}
	c := session.DB("test").C("users")
	count, err :=c.Find(bson.M{"username": uname}).Count()
	log.Printf("%d",count)
	err = c.Find(bson.M{"username": uname}).One(&result)
        if err != nil {
		log.Printf(err.Error())
		log.Printf("Couldn't get a result for username")
				return nil
        }
	return &result
}