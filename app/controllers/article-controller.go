package controllers

import (
	"fmt"
	"main/db"
	"main/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// 全データを取得する
func Index(c echo.Context) error {
	db, err := db.ConnectDb()
	if err != nil {
		return err
	}
	result := []models.Article{}
	if err := db.Find(&result).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

// 受け取った単一データをDBに書き込む
func StoreArticle(c echo.Context) error {
	db, err := db.ConnectDb()
	if err != nil {
		return err
	}
	article := models.Article{
		Title: c.FormValue("title"),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, article)

}

func UpdateArticle(c echo.Context) error {
	db, err := db.ConnectDb()
	if err != nil {
		return err
	}
	// パラメータを取得し、整数型に変換
	param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// モデルの宣言
	article := models.Article{}
	article.ID = uint(param)

	// DBの書き換え
	if err := db.First(&article).Update("title", c.FormValue("title")).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, article)

}

func DestroyArticle(c echo.Context) error {
	db, err := db.ConnectDb()
	if err != nil {
		return err
	}
	// パラメータを取得し、整数型に変換
	param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	var article models.Article

	article.ID = uint(param)

	if err := db.Delete(&article).Error; err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, article)

}
