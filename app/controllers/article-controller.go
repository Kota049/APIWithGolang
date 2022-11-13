package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全データを取得する
func Index(c echo.Context) error {
	dsn := "root:password@tcp(mysql:3306)/api_with_golang"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	result := []Article{}
	db.Find(&result)
	return c.JSON(http.StatusOK, result)
}

// 受け取った単一データをDBに書き込む
func StoreArticle(c echo.Context) error {
	dsn := "root:password@tcp(mysql:3306)/api_with_golang"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	article := Article{
		Title: c.FormValue("title"),
	}
	// 日本語がユニコードに変換されちゃう
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, article)

}

func UpdateArticle(c echo.Context) error {
	dsn := "root:password@tcp(mysql:3306)/api_with_golang"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	// パラメータを取得し、整数型に変換
	param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// モデルの定義
	article := Article{}
	article.Id = param

	// DBの書き換え
	if err := db.First(&article).Update("title", c.FormValue("title")).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, article)

}

func DestroyArticle(c echo.Context) error {
	dsn := "root:password@tcp(mysql:3306)/api_with_golang"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	// パラメータを取得し、整数型に変換
	param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	article := Article{
		Id: param,
	}

	if err := db.Delete(&article).Error; err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, article)

}

// Articleテーブルのモデル
type Article struct {
	Id    int
	Title string
}
