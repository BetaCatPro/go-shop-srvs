package main

import (
	"context"
	"go-shop-srvs/goods_srv/global"
	"go-shop-srvs/goods_srv/model"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func initDB(dsn string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	return db, err
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goods_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := initDB(dsn)
	_ = db.AutoMigrate(&model.Category{},
		&model.Brands{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})
	//Mysql2Es()
}

func Mysql2Es() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_shop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := initDB(dsn)

	host := "http://169.254.14.87:9200"
	logger := log.New(os.Stdout, "", log.LstdFlags)
	var err error
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	var goods []model.Goods
	db.Find(&goods)
	for _, g := range goods {
		esModel := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}

		_, err = global.EsClient.Index().Index(esModel.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
