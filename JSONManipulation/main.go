package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Inventory struct {
	InventoryId int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int       `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}

type Placement struct {
	RoomId int    `json:"room_id"`
	Name   string `json:"name"`
}

// Items in the Meeting Room
func itemsInMeetingRoom(items []*Inventory) {
	var result []*Inventory
	for _, v := range items {
		if strings.ToLower(v.Placement.Name) == "meeting room" {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Items in the Meeting Room")
	fmt.Println(string(data))
}

// All electronic devices
func allElectronicDevices(items []*Inventory) {
	var result []*Inventory
	for _, v := range items {
		if strings.ToLower(v.Type) == "electronic" {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("All electronic devices")
	fmt.Println(string(data))
}

// All furniture
func allFurnitureType(items []*Inventory) {
	var result []*Inventory
	for _, v := range items {
		if strings.ToLower(v.Type) == "furniture" {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("All furniture type")
	fmt.Println(string(data))
}

// All items with brown color
func allBrownColor(items []*Inventory) {
	var result []*Inventory
	for _, v := range items {
		for _, tag := range v.Tags {
			if strings.ToLower(tag) == "brown" {
				result = append(result, v)
				break
			}
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("All items with brown color")
	fmt.Println(string(data))
}

func main() {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		panic(err)
	}

	fmt.Println("Success")

	defer jsonFile.Close()

	file, _ := ioutil.ReadAll(jsonFile)

	var inventories []*Inventory

	err = json.Unmarshal(file, &inventories)
	if err != nil {
		panic(err)
	}

	itemsInMeetingRoom(inventories)
	allElectronicDevices(inventories)
	allFurnitureType(inventories)
	allBrownColor(inventories)
}
