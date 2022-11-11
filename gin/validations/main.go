package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var animalParameter = "animalId"

type Animal struct {
	ID     uuid.UUID `json:"id" uri:"animalId"`
	Name   string    `json:"name"`
	Specie string    `json:"specie"`
}

func (animal *Animal) UnmarshalJSON(buf []byte) error {
	type Alias Animal
	aux := Alias{
		ID: uuid.Must(uuid.NewRandom()),
	}

	if err := json.Unmarshal(buf, &aux); err != nil {
		return err
	}

	*animal = Animal(aux)
	return nil
}

func (animal *Animal) BindUri(params gin.Params) error {
	if v, ok := params.Get(animalParameter); ok {
		var err error
		if animal.ID, err = uuid.Parse(v); err != nil {
			return err
		}
	}
	return nil
}

func (a Animal) Merge(target Animal) Animal {
	if a.Name == "" {
		a.Name = target.Name
	}
	if a.Specie == "" {
		a.Specie = target.Specie
	}

	return a
}

var animalsDB hashMap[uuid.UUID, Animal]

func init() {
	animalsDB = newHashMap[uuid.UUID, Animal]()
}

func main() {
	engine := gin.Default()

	animals := engine.Group("/api/animals")
	{
		animals.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, animalsDB.toArray())
		})

		animals.POST("", func(c *gin.Context) {
			var a Animal
			if err := c.BindJSON(&a); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			animalsDB.put(a.ID, a)
			c.JSON(http.StatusOK, a)
		})

		animals.GET("/:"+animalParameter, func(c *gin.Context) {
			var a Animal
			if err := a.BindUri(c.Params); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			a, ok := animalsDB.get(a.ID)
			if !ok {
				c.AbortWithError(http.StatusBadRequest, errors.New("entry not found"))
				return
			}

			c.JSON(http.StatusOK, a)
		})

		animals.PUT("/:"+animalParameter, func(c *gin.Context) {
			var a Animal
			if err := c.BindJSON(&a); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			if err := a.BindUri(c.Params); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			v, ok := animalsDB.get(a.ID)
			if !ok {
				c.AbortWithError(http.StatusBadRequest, errors.New("entry not found"))
				return
			}

			c.JSON(http.StatusOK, animalsDB.put(a.ID, a.Merge(v)))
		})
	}
	engine.Run(":8080")
}
