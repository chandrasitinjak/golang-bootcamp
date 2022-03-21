package controller

import (
	"github.com/gin-gonic/gin"
	_struct "sesi_8/paractice/struct"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person _struct.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB
}
