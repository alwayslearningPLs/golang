package main

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	queryFoodFirst = `SELECT * FROM "foods" WHERE "foods"."deleted_at" IS NULL ORDER BY "foods"."id" LIMIT 1`
	queryFoodTake  = `SELECT * FROM "foods" WHERE "foods"."deleted_at" IS NULL LIMIT 1`
	queryFoodLast  = `SELECT * FROM "foods" WHERE "foods"."deleted_at" IS NULL ORDER BY "foods"."id" DESC LIMIT 1`
)

func TestMain(m *testing.M) {
	configureDB("host=localhost user=postgres password=abc123. dbname=learning port=5555 sslmode=disable TimeZone=Europe/Madrid")

	os.Exit(m.Run())
}

// first and last works with: struct and gorm.Model
// take works with: struct, gorm.Model and table
// When no primary key is inside the structure, it will order by the first field inside the struct (not tested)
//
// ErrRecordNotFound is triggered when using first/last method and no record is found. A possible solution, would be using .Limit(1).Find(&result);
// ErrModelValueRequired is triggered when using first/last method using table method.
func TestQueryString(t *testing.T) {
	var food Food

	t.Run("query food first", func(t *testing.T) {
		assert.Equal(t, queryFoodFirst, getInstanceDry().First(&food).Statement.SQL.String())
	})

	t.Run("query food take", func(t *testing.T) {
		assert.Equal(t, queryFoodTake, getInstanceDry().Take(&food).Statement.SQL.String())
	})

	t.Run("query food last", func(t *testing.T) {
		assert.Equal(t, queryFoodLast, getInstanceDry().Last(&food).Statement.SQL.String())
	})
}

func TestQueryErrors(t *testing.T) {
	var food Food

	t.Run("trigger ErrRecordNotFound", func(t *testing.T) {
		tx := getInstanceWithCtx(context.Background()).Where("id=20000").First(&food)

		assert.Equal(t, tx.RowsAffected, int64(0))
		assert.ErrorIs(t, tx.Error, gorm.ErrRecordNotFound)
	})

	t.Run("avoid ErrRecordNotFound", func(t *testing.T) {
		// remember that find method accepts a struct or a slice.
		tx := getInstanceWithCtx(context.Background()).Where("id=20000").Limit(1).Find(&food)

		assert.Equal(t, tx.RowsAffected, int64(0))
		assert.Nil(t, tx.Error)
	})

	t.Run("first method using table doesn't work", func(t *testing.T) {
		var food map[string]any
		tx := getInstanceWithCtx(context.Background()).Table("foods").First(&food)

		assert.ErrorIs(t, tx.Error, gorm.ErrModelValueRequired)
	})

	t.Run("last method using table doesn't work", func(t *testing.T) {
		var food map[string]any
		tx := getInstanceWithCtx(context.Background()).Table("foods").Last(&food)

		assert.ErrorIs(t, tx.Error, gorm.ErrModelValueRequired)
	})
}

func TestQueryStruct(t *testing.T) {
	t.Run("first method using struct", func(t *testing.T) {
		var food Food
		tx := getInstanceWithCtx(context.Background()).First(&food)

		assert.Nil(t, tx.Error)
		assert.NotEmpty(t, food.Name)
	})

	t.Run("last method using struct", func(t *testing.T) {
		var food Food
		tx := getInstanceWithCtx(context.Background()).Last(&food)

		assert.Nil(t, tx.Error)
		assert.NotEmpty(t, food.Name)
	})

	t.Run("take method using struct", func(t *testing.T) {
		var food Food
		tx := getInstanceWithCtx(context.Background()).Take(&food)

		assert.Nil(t, tx.Error)
		assert.NotEmpty(t, food.Name)
	})
}

func TestQueryModel(t *testing.T) {
	t.Run("first method using gorm.Model", func(t *testing.T) {
		var food map[string]any
		tx := getInstanceWithCtx(context.Background()).Model(&Food{}).First(&food)

		assert.Nil(t, tx.Error)
		assert.NotEmpty(t, food["name"])
	})

	t.Run("last method using gorm.Model", func(t *testing.T) {
		var food map[string]any
		tx := getInstanceWithCtx(context.Background()).Model(&Food{}).Last(&food)

		assert.Nil(t, tx.Error)
		assert.NotEmpty(t, food["name"])
	})

	t.Run("take method using gorm.Model", func(t *testing.T) {
		var food map[string]any
		tx := getInstanceWithCtx(context.Background()).Model(&Food{}).Take(&food)

		assert.Nil(t, tx.Error)
		assert.NotEmpty(t, food["name"])
	})
}

func TestQueryTable(t *testing.T) {
	t.Run("take method using table", func(t *testing.T) {
		var food map[string]any
		tx := getInstanceWithCtx(context.Background()).Table("foods").Take(&food)

		assert.Nil(t, tx.Error)
		assert.NotEmpty(t, food["name"])
	})
}

func TestQueryByPrimaryKey(t *testing.T) {
	t.Run("first by primary key", func(t *testing.T) {
		var food Food
		tx := getInstanceWithCtx(context.Background()).First(&food, 1)

		assert.Nil(t, tx.Error)
		assert.Equal(t, uint(1), food.ID)
	})

	t.Run("first by primary key by input struct", func(t *testing.T) {
		var food Food
		food.ID = 1
		tx := getInstanceWithCtx(context.Background()).First(&food)

		assert.Nil(t, tx.Error)
		assert.Equal(t, uint(1), food.ID)
		assert.NotEmpty(t, food.Name)
	})

	t.Run("first by primary key by input gorm.Model", func(t *testing.T) {
		var (
			food   Food
			result Food // It won't work using map, we need to use the Model food
		)
		food.ID = 1
		tx := getInstanceWithCtx(context.Background()).Model(&food).First(&result)

		assert.Nil(t, tx.Error)
		assert.Equal(t, uint(1), result.ID)
	})

	t.Run("last by primary key", func(t *testing.T) {
		var food Food
		tx := getInstanceWithCtx(context.Background()).Last(&food, 1)

		assert.Nil(t, tx.Error)
		assert.Equal(t, uint(1), food.ID)
	})

	t.Run("last by primary key by input struct", func(t *testing.T) {
		var food Food
		food.ID = 1
		tx := getInstanceWithCtx(context.Background()).Last(&food)

		assert.Nil(t, tx.Error)
		assert.Equal(t, uint(1), food.ID)
		assert.NotEmpty(t, food.Name)
	})

	t.Run("find by primary key with more than one value", func(t *testing.T) {
		var result []Food
		tx := getInstanceWithCtx(context.Background()).Find(&result, []int{1, 2, 3})

		assert.Nil(t, tx.Error)
		for _, i := range []uint{1, 2, 3} {
			assert.Equal(t, i, result[i-1].ID)
		}
	})
}

func TestQueryJoins(t *testing.T) {
	t.Run("Simple join", func(t *testing.T) {
		var result []Warehouse
		tx := getInstanceWithCtx(context.Background()).Joins("Food").Find(&result)

		assert.Nil(t, tx.Error)
		assert.Greater(t, tx.RowsAffected, int64(5))
	})

	t.Run("Simple join with query", func(t *testing.T) {
		var (
			result []Warehouse
			food   Food
		)
		food.Name = "name_2"
		tx := getInstanceWithCtx(context.Background()).Joins("Food", &food).Find(&result)
		// t.Fatal(getInstanceDry().Joins("Food", &food).Find(&result).Statement.SQL.String())
		t.Fatal(getInstanceDry().Joins("LEFT JOIN foods on foods.food_id = warehouses.food_id and foods.food_name like ?", food.Name).Find(&result).Statement.SQL.String())

		assert.Nil(t, tx.Error)
		assert.Equal(t, int64(1), tx.RowsAffected)
	})
}

func populate(DB *gorm.DB) {
	var foods []Food
	for i := 0; i < 10; i++ {
		tmp := Food{
			Name:  "name_" + strconv.Itoa(i),
			Value: uint(i),
		}
		DB.Create(&tmp)
		foods = append(foods, tmp)
	}

	for i := 0; i < 10; i++ {
		DB.Create(&Warehouse{
			Name:     "warehouse_" + strconv.Itoa(i),
			Location: "location_" + strconv.Itoa(i),
			Food:     foods[i],
		})
	}
}
