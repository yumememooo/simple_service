package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"simple_service/internal/model"
)

//TODO
// func FindPet(animal_kind string, page string, size string) ([]contract.Pet, string, error) {
//}

func FindPet(animal_kind string) ([]model.Pet, error) {
	pets := []model.Pet{}
	jsonFile, err := os.Open("./sample/mock/json/pets.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return pets, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return pets, err
	}

	json.Unmarshal(byteValue, &pets)
	if animal_kind == "" {
		return pets, nil
	}
	subPets := []model.Pet{}

	for _, v := range pets {
		if v.Animal_kind == animal_kind {
			subPets = append(subPets, v)
		}

	}
	return subPets, nil

}
