package controller

import (
	"assignment_2/database"
	"assignment_2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var data models.CreateOrder

	var data_items []models.Item

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		Customer_Name: data.Customer_Name,
		OrderedAt:     data.OrderedAt,
	}

	db.Create(&order)
	id_order := order.ID
	for _, v := range data.Item {
		item := models.Item{
			Item_Code:   v.Item_Code,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderID:     order.ID,
		}

		data_items = append(data_items, item)
	}

	result := db.Create(&data_items)
	log.Println(id_order, result.RowsAffected)

	var returnData interface{}

	returnData = models.Order{
		ID:            data_items[0].OrderID,
		Customer_Name: data.Customer_Name,
		OrderedAt:     time.Now(),
		Item:          data_items,
	}
	ctx.JSON(http.StatusOK, returnData)
}

func GetOrder(ctx *gin.Context) {
	db := database.GetDB()

	orders := []models.Order{}

	err := db.Preload("Item").Find(&orders).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func GetOrderById(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	var id_order = ctx.Param("orderID")

	err := db.Preload("Item").Find(&order, id_order).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func UpdateOrder(ctx *gin.Context) {

	db := database.GetDB()

	id := ctx.Param("orderID")
	id2, _ := strconv.Atoi(id)

	var data = models.CreateOrder{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		ID:            uint(id2),
		Customer_Name: data.Customer_Name,
		OrderedAt:     data.OrderedAt,
		Item:          data.Item,
	}
	//db.Create(&order)
	//
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)

	ctx.JSON(http.StatusOK, order)

}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()

	id := ctx.Param("orderID")
	id2, _ := strconv.Atoi(id)

	orders := models.Order{}
	items := models.Item{}

	err := db.Where("order_id = ?", id2).Delete(&items).Error
	rows := db.Where("ID = ?", id2).Delete(&orders).RowsAffected

	if rows == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	if err != nil {
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data deleted",
	})
}
