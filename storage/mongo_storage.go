package storage

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SocialData struct {
	ecoId       string
	source      string
	externalId  []string
	data        interface{}
	autoMapping bool
	version     int64
}

var glb_session mgo.Session

func init() {
	glb_session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer glb_session.Close()

	// Optional. Switch the session to a monotonic behavior.
	glb_session.SetMode(mgo.Monotonic, true)

	c := glb_session.DB("socialhub").C("socialData")
	//err = c.Insert(&SocialData{"Ale", "+55 53 8116 9639"},
	//	&SocialData{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//log.Fatal(err)
	//}

	result := SocialData{}
	err = c.Find(bson.M{"ecoId": "8231330"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result)
}