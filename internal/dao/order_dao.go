/**
 * @author  qyanzh
 * @create  2022/02/26 21:29
 */

package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"gorm.io/gorm"
	"time"
)

type OrderDao struct {
	db *gorm.DB
}

func NewOrderDao() *OrderDao {
	return &OrderDao{db: db.DbClient}
}

func (dao *OrderDao) AddOrder(order *model.Order) int64 {
	result := dao.db.Create(&order)
	return result.RowsAffected
}

func (dao *OrderDao) DeleteOrderByID(ID uint) int64 {
	result := dao.db.Delete(&model.Order{}, ID)
	return result.RowsAffected
}

func (dao *OrderDao) UpdateOrder(order *model.Order) int64 {
	result := dao.db.Save(&order)
	return result.RowsAffected
}

func (dao *OrderDao) SelectOrderByID(ID uint) *model.Order {
	order := model.Order{}
	dao.db.First(&order, ID)
	return &order
}

func (dao *OrderDao) SelectOrderByRoomIDTimeRange(roomID uint, begin, end time.Time) *[]model.Order {
	var orders []model.Order
	dao.db.Where("room_id = ? AND order_time BETWEEN ? AND ?", roomID, begin, end).Find(&orders)
	return &orders
}

func (dao *OrderDao) SelectOrdersByRoomID(roomID uint) *[]model.Order {
	var orders []model.Order
	dao.db.Where(&model.Order{RoomID: roomID}).Find(&orders)
	return &orders
}

type Result struct {
	Result float64
}

func (dao *OrderDao) SelectValue(expr string, roomID uint, timeRange uint) float64 {
	r := Result{}
	where := "`orders`.`deleted_at` IS NULL AND room_id = ?"
	param := make([]interface{}, 0)
	param = append(param, roomID)
	if timeRange != 0 {
		where = where + " AND order_time BETWEEN ? AND ?"
		end := time.Now()
		begin := end.Add(-time.Duration(timeRange) * time.Second)
		param = append(param, begin)
		param = append(param, end)
	}
	dao.db.Table(model.Order{}.TableName()).
		Select(expr+" result").
		Where(where, param...).
		Find(&r)
	return r.Result
}
