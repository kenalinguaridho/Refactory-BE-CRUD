package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kenalinguaridho/config"
	"github.com/kenalinguaridho/entities"
	"github.com/kenalinguaridho/helper"
	"gorm.io/gorm"
)

func Create(rw http.ResponseWriter, r *http.Request) {

	var item entities.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		helper.Response(rw, 400, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&item).Error; err != nil {
		helper.Response(rw, 500, err.Error(), nil)
		return
	}

	helper.Response(rw, 201, "New item has been added", nil)

}

func GetAll(rw http.ResponseWriter, r *http.Request) {

	var items []entities.Item

	if err := config.DB.Find(&items).Error; err != nil {
		helper.Response(rw, 500, err.Error(), nil)
		return
	}

	helper.Response(rw, 200, "List of items", items)

}

func GetItemById(rw http.ResponseWriter, r *http.Request) {

	var item entities.Item

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := config.DB.First(&item, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(rw, 404, "No item was found", nil)
			return
		}

		helper.Response(rw, 500, err.Error(), nil)
		return
	}

	helper.Response(rw, 200, "Item detail", item)

}

func Update(rw http.ResponseWriter, r *http.Request) {

	var item entities.Item

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := config.DB.First(&item, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(rw, 404, "No item was found", nil)
			return
		}

		helper.Response(rw, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		helper.Response(rw, 400, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&item).Error; err != nil {
		helper.Response(rw, 500, err.Error(), nil)
		return
	}

	helper.Response(rw, 201, "Item has been updated", nil)

}

func Delete(rw http.ResponseWriter, r *http.Request) {

	var item entities.Item

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := config.DB.First(&item, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(rw, 404, "No item was found", nil)
			return
		}

		helper.Response(rw, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&item).Error; err != nil {
		helper.Response(rw, 500, err.Error(), nil)
		return
	}

	helper.Response(rw, 200, "Item has been deleted", nil)

}