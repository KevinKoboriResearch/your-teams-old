package Interface

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//__ INSERT __________________________________________________________________//
func InsertDB(collection string, entity interface{}) (err error) {
	err = OpenSession(collection).Insert(entity)
	return
}

//__ UPDATE SINGLE ___________________________________________________________//
func UpdateSingleDB(collection string, id string, idValue string, position string, value interface{}) (err error) {
	if id == "_id" {
		err = OpenSession(collection).Update(bson.M{id: bson.ObjectIdHex(idValue)}, bson.M{"$set": bson.M{position: value}})
		return
	}
	err = OpenSession(collection).Update(bson.M{id: idValue}, bson.M{"$set": bson.M{position: value}})
	return
}

//__ UPDATE __________________________________________________________________//
func UpdateDB(collection string, id string, idValue string, entity interface{}) (err error) {
	change := mgo.Change{
		Update:    bson.M{"$set": entity},
		ReturnNew: true,
	}
	if id == "_id" {
		_, err = OpenSession(collection).Find(bson.M{id: bson.ObjectIdHex(idValue)}).Apply(change, &entity)
		return
	}
	_, err = OpenSession(collection).Find(bson.M{id: idValue}).Apply(change, &entity)
	return
}

//__ FIND ____________________________________________________________________//
func FindDB(collection string, id string, idValue string, entity interface{}) (err error) {
	if id == "_id" {
		err = OpenSession(collection).Find(bson.M{id: bson.ObjectIdHex(idValue)}).One(entity)
		return
	}
	err = OpenSession(collection).Find(bson.M{id: idValue}).One(entity)
	return
}

//__ FIND ALL WHILE __________________________________________________________//
func FindAllWhileDB(collection string, position string, value interface{}, entity interface{}) (err error) {
	err = OpenSession(collection).Find(bson.M{position: value}).All(entity)
	return
}

//__ FIND ALL ENABLED WHILE __________________________________________________//
func FindAllEnabledWhileDB(collection string, position string, value interface{}, entity interface{}) (err error) {
	err = OpenSession(collection).Find(bson.M{position: value, "enable": true}).All(entity)
	return
}

//__ FIND ALL ENABLED ________________________________________________________//
func FindAllEnabledDB(collection string, entity interface{}) (err error) {
	err = OpenSession(collection).Find(bson.M{"enable": true}).All(entity)
	return
}

//__ FIND ALL ________________________________________________________________//
func FindAllDB(collection string, entity interface{}) (err error) {
	err = OpenSession(collection).Find(nil).All(entity)
	return
}

//__ DELETE __________________________________________________________________//
func DeleteDB(collection string, id string, idValue string, entity interface{}) (err error) {
	if id == "_id" {
		err = OpenSession(collection).Remove(bson.M{id: bson.ObjectIdHex(idValue)})
		return
	}
	err = OpenSession(collection).Remove(bson.M{id: idValue})
	return
}

//__ DELETE ALL COLLECTION __________________________________________________________________//
func DeleteAllCollection(collection string) (err error) {
	err = OpenSession(collection).Remove(nil)
	return
}
