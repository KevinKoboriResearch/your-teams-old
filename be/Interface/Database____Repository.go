package Interface

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2"
	"reflect"
	"strings"
)

//-------------------------------- DB --------------------------------//
func InsertDB(entity interface{}, collection string) (err error) {
	c := OpenSession(collection)
	err = c.Insert(entity)
	if err != nil {
		log.Print("[ERROR] failed to insert entity: ", err)
		return
	}
	return
}

//-------------------------------- DB --------------------------------//
func UpdateSingleDB(position string, value string, entity interface{}, collection string, id string) (interface{}, error) {
	c := OpenSession(collection)
	var err error
	if collection == "user_entity" || collection == "user_link" {
		err = c.Update(bson.M{"username": id}, bson.M{"$set": bson.M{position: value}})
	} else if collection == "game_entity" {
		err = c.Update(bson.M{"game": id}, bson.M{"$set": bson.M{position: value}})
	} else {
		err = c.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{position: value}})
	}
	if err != nil {
		log.Fatal(err)
		return entity, err
	}
	return entity, nil
}

//-------------------------------- USR --------------------------------//
func UpdatePartialDB(entity interface{}, collection string, ids ...string) (interface{}, error) {
	c := OpenSession(collection)
	rids := make([]reflect.Value, len(ids))
	for i, a := range ids {
		rids[i] = reflect.ValueOf(a)
	}
	m := make(map[string]interface{})
	t := reflect.TypeOf(entity)
	v := reflect.ValueOf(entity)
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).String() != "" {
			if t.Field(i).Name != "Id" {
				field := strings.ToLower(t.Field(i).Name)
				m[field] = v.Field(i).Interface()
			}
		}
	}
	update := bson.M{}
	err := mapstructure.Decode(m, &update)
	if err != nil {
		log.Print("[ERROR] Can't map entity: ", err)
		return entity, err
	}
	change := mgo.Change{
		Update:    bson.M{"$set": update},
		ReturnNew: true,
	}
	if collection == "user_entity" {
		_, err = c.Find(bson.M{"username": ids[0]}).Apply(change, &entity)
	} else if collection == "threads" {
		_, err = c.Find(bson.M{"_id": bson.ObjectIdHex(rids[1].String()), "category._id": bson.ObjectIdHex(rids[0].String())}).Apply(change, &entity)
	} else {
		_, err = c.Find(bson.M{"_id": bson.ObjectIdHex(rids[0].String())}).Apply(change, &entity)
	}
	if err != nil {
		log.Print("[ERROR] failed update entity: ", err)
		return entity, err
	}
	return entity, nil
}

//-------------------------------- DB --------------------------------//
func UpdateDB(entity interface{}, collection string, id string) (interface{}, error) {
	var err error
	c := OpenSession(collection)
	if collection == "user_entity" {
		err = c.Update(bson.M{"username": id}, entity)
	} else if collection == "game_entity" {
		err = c.Update(bson.M{"name": id}, entity)
	} else {
		err = c.Update(bson.M{"_id": id}, entity)
	}
	if err != nil {
		log.Fatal(err)
		return entity, err
	}
	return entity, nil
}

//-------------------------------- DB --------------------------------//
func FindUnitDB(id string, entity interface{}, collection string) (err error) {
	c := OpenSession(collection)
	if collection == "user_entity" {
		err = c.Find(bson.M{"username": id}).One(entity)
	} else if collection == "game_entity" {
		err = c.Find(bson.M{"name": id}).One(entity)
	} else {
		err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(entity)
	}
	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}
	return
}

//-------------------------------- DB --------------------------------//
func FindAllWhileDB(position string, value string, entity interface{}, collection string) (err error) {
	c := OpenSession(collection)
	err = c.Find(bson.M{position: value}).All(entity)
	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}
	return
}

//-------------------------------- DB --------------------------------//
func FindAllEnabledWhileDB(position string, value string, entity interface{}, collection string) (err error) {
	c := OpenSession(collection)
	err = c.Find(bson.M{position: value, "enable": true}).All(entity)
	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}
	return
}

//-------------------------------- DB --------------------------------//
func FindAllEnabledDB(entity interface{}, collection string) (err error) {
	c := OpenSession(collection)
	err = c.Find(bson.M{"enable": true}).All(entity)
	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}
	return
}

//-------------------------------- DB --------------------------------//
func FindAllDB(entity interface{}, collection string) (err error) {
	c := OpenSession(collection)
	err = c.Find(nil).All(entity)
	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}
	return
}

//-------------------------------- DB --------------------------------//
func DeleteDB(id string, entity interface{}, collection string) (err error) {
	c := OpenSession(collection)
	if collection == "user_entity" || collection == "user_games" {
		err = c.Remove(bson.M{"username": id})
	} else 	if collection == "game_entity" {
		err = c.Remove(bson.M{"name": id})
	} else {
		err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	}
	if err != nil {
		log.Print("[ERROR] failed to get entity: ", err)
		return
	}
	return
}
