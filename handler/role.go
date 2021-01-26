package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ken0911208818/goHomeWork/lib/middleware"
	"github.com/ken0911208818/goHomeWork/model"
)

var Budai []model.Role

func init() {
	Budai = model.Init()
}

func Index(c *gin.Context) {
	middleware.SendResponse(c, http.StatusOK, Budai)
}

func Create(c *gin.Context) {
	role := model.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		middleware.SendErrorResponse(c, http.StatusNotFound, nil)
		return
	}
	Budai = append(Budai, role)
	middleware.SendResponse(c, http.StatusOK, role)
}

func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	for _, value := range Budai {
		if value.ID == uint(id) {
			middleware.SendResponse(c, http.StatusOK, value)
			return
		}
	}
	middleware.SendErrorResponse(c, http.StatusNotFound, errors.New("找不到該角色"))
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	role := model.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		middleware.SendErrorResponse(c, http.StatusNotFound, nil)
		return
	}

	for k, value := range Budai {
		if value.ID == uint(id) {
			Budai[k].Name = role.Name
			Budai[k].Summary = role.Summary
			middleware.SendResponse(c, http.StatusOK, Budai[k])
			return
		}
	}
	middleware.SendErrorResponse(c, http.StatusNotFound, errors.New("找不到該角色"))
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		middleware.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	for k, value := range Budai {
		if value.ID == uint(id) {
			Budai = remove(Budai, k)
			middleware.SendResponse(c, http.StatusNoContent, nil)
			return
		}
	}
	middleware.SendErrorResponse(c, http.StatusNotFound, errors.New("找不到該資料"))
}

func remove(r []model.Role, i int) []model.Role {
	r[i] = r[len(r)-1]
	return r[:len(r)-1]
}
