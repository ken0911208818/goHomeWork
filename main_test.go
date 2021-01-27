package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ken0911208818/goHomeWork/model"
	"github.com/stretchr/testify/assert"
)

func TestAllData(t *testing.T) {
	router := setupRouter()
	result, code := Get("/role/", router)
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, string(result), "阿修羅")
	assert.Contains(t, string(result), "白塵子")
}

//測試預設資料中的ID 2
func TestOneData(t *testing.T) {
	router := setupRouter()
	result, code := Get("/role/2", router)
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, string(result), "白塵子")
}
func TestCreate(t *testing.T) {
	JsonData := model.Role{
		Name:    "伯藏主",
		Summary: "原本應該繼位成為白狐國第65代君宇，後因其弟犬若丸篡位而流離至罪惡坑，成為二惡首。",
		Skills: []model.RoleSkill{
			model.RoleSkill{
				Type: "武學",
				Name: "金風雪柳",
			},
		},
	}
	router := setupRouter()

	response, code := PostJson("/role/", JsonData, router)
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, string(response), "伯藏主")
}
func TestUpdate(t *testing.T) {
	JsonData := model.Role{
		Name:    "阿修羅",
		Summary: "測試修改Name阿修羅",
	}
	router := setupRouter()

	response, code := PutJson("/role/1", JsonData, router)
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, string(response), "測試修改Name阿修羅")
}

func TestDelete(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/role/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	// 再刪除一次　應該會無法刪除
	w1:= httptest.NewRecorder()
	req2 := httptest.NewRequest("DELETE", "/role/1", nil)
	router.ServeHTTP(w1, req2)
	assert.Equal(t, http.StatusNotFound, w1.Code)
}

func Get(uri string, router *gin.Engine) ([]byte, int) {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	return getResult(result)
}

func PostJson(uri string, param model.Role, router *gin.Engine) ([]byte, int) {
	jsonByte, _ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	return getResult(result)
}

func PutJson(uri string, param model.Role, router *gin.Engine) ([]byte, int) {
	jsonByte, _ := json.Marshal(param)
	req := httptest.NewRequest("PUT", uri, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	return getResult(result)
}

func getResult(w *http.Response) (body []byte, statusCode int) {
	defer w.Body.Close()
	body, _ = ioutil.ReadAll(w.Body)
	return body, w.StatusCode
}
