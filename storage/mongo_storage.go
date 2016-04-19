package storage

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataStore struct {
	Session *mgo.Session
}

func (ds *DataStore) GET(country string, ecoId []string, source []string) []bson.M {
	c := ds.Session.DB("socialhub").C("socialData")
	var toReturn = []bson.M{}
	for i := range ecoId {
		for j := range source {
			result := bson.M{}
			err := c.Find(bson.M{"ecoId": ecoId[i], "source":source[j]}).One(&result)
			if err != nil {
				log.Print("Combination:%s,%s Message:%s", ecoId[i], source[j], err)
			} else {
				toReturn = append(toReturn, result)
			}
		}
	}
	return toReturn
}