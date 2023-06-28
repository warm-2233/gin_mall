package model

import (
	"gin_mall/cache"
	"strconv"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	Category      uint
	Title         string
	Info          string
	ImageUrl      string
	Price         float64
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
	Views         int64
}

// views
func (p *Product) View() int64 {
	countStr, err := cache.RedisClient.Get(cache.ProductViewKey(p.ID)).Result()
	if err != nil {
		if err == redis.Nil {
			cache.RedisClient.IncrBy(cache.ProductViewKey(p.ID), p.Views)
		}
		return 0
	}

	count, _ := strconv.ParseInt(countStr, 10, 64)
	return count
}

// Add view
func (p *Product) AddView() {
	cache.RedisClient.Incr(cache.ProductViewKey(p.ID))
}
