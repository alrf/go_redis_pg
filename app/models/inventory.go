package models

import (
	u "main/utils"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis"
	"encoding/json"
	"log"
	"time"
)

type Inventory struct {
	gorm.Model `json:"-"`
	Department string `gorm:"not null" json:"department"`
	Section int `gorm:"not null" json:"section"`
	Equipment string `gorm:"not null" json:"equipment"`
	Description string `gorm:"not null" json:"description"`
}

func (inventory *Inventory) StructToString() string {

    res, err := json.Marshal(inventory)
    if err != nil {
        log.Println(err)
    }
    log.Println(string(res))

    return string(res)

}

func (inventory *Inventory) Validate() (map[string] interface{}, bool) {

	if len(inventory.Department) < 4 {
		return u.Message(false, "Department is required"), false
	}

	if inventory.Section == 0 {
		return u.Message(false, "Section is required"), false
	}

	if len(inventory.Equipment) < 4 {
		return u.Message(false, "Equipment is required"), false
	}

	if len(inventory.Description) < 1 {
		return u.Message(false, "Description is required"), false
	}

 	inv := inventory.StructToString()

 	//Redis cache, get full inventory JSON as Redis key
	res, err := GetRDB().Get(inv).Result()
	if err == redis.Nil {

		log.Println("Redis: inventory will be created", inv, res)
		//Add full inventory JSON as Redis key
		err = GetRDB().Set(inv, "0", 0).Err()
		if err != nil {
			return u.Message(false, "Redis: set error. Please retry"), false
		}

		//Equipment must be unique in Postgres
		temp := &Inventory{}
		err = GetDB().Table("inventories").Where("department = ? and section = ? and equipment = ?", inventory.Department, inventory.Section, inventory.Equipment).First(temp).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return u.Message(false, "Connection error. Please retry"), false
		}
		if temp.Equipment != "" {
			return u.Message(false, "Equipment already exists"), false
		}

		return u.Message(false, "Redis: requirement passed"), true

	} else if err != nil {
		return u.Message(false, "Redis: get error. Please retry"), false
	} else {
		return u.Message(false, "Redis: equipment already exists"), false
	}

}

func (inventory *Inventory) Create() (map[string] interface{}) {

	timestamp := time.Now().UTC()

	if resp, ok := inventory.Validate(); !ok {
		return resp
	}

	GetDB().Create(inventory)

	if inventory.ID <= 0 {
		return u.Message(false, "Failed to create inventory, connection error.")
	}

	response := u.Message(true, "Inventory has been created")
	response["data"] = inventory
	response["timestamp"] = timestamp
	return response

}

func List(limit uint) (map[string] interface{}) {

	inventories := make([]*Inventory, 0)
	err := GetDB().Order("id desc").Limit(limit).Find(&inventories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry")
	}

	response := u.Message(true, "success")
	response["data"] = inventories
	return response

}
