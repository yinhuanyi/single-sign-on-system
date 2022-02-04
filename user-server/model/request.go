/**
 * @Author：Robby
 * @Date：2021/11/22 10:37 上午
 * @Function：
 **/

package model

type GoodsCreateInput struct {
	GoodsSN int64 `json:"goods_sn" db:"goods_sn"`
	Name string `json:"name" binding:"required" db:"goods_sn"`
	MarketPrice float64 `json:"market_price" binding:"required" db:"market_price"`
}

