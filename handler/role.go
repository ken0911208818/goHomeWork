package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ken0911208818/goHomeWork/model"
)

var Budai []model.Role

func init() {
	Budai = model.Init()
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, Budai)
}

func Create(c *gin.Context) {
	role := model.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	Budai = append(Budai, role)
	c.JSON(http.StatusOK, role)
}

func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	for _, value := range Budai {
		if value.ID == uint(id) {
			c.JSON(http.StatusOK, value)
			return
		}
	}
	c.JSON(http.StatusNotFound, "找不到該角色")
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	role := model.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	for k, value := range Budai {
		if value.ID == uint(id) {
			Budai[k].Name = role.Name
			Budai[k].Summary = role.Summary
			c.JSON(http.StatusOK, Budai[k])
			return
		}
	}

	c.JSON(http.StatusNotFound, "找不到該角色")
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	for k, value := range Budai {
		if value.ID == uint(id) {
			Budai = remove(Budai, k)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, "找不到該資料")
}

func remove(r []model.Role, i int) []model.Role {
	r[i] = r[len(r)-1]
	return r[:len(r)-1]
}
