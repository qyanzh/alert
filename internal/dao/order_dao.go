/**
 * @author  qyanzh
 * @create  2022/02/26 21:29
 */

package dao

import (
	"alert/internal/db"
	"alert/internal/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type OrderDao struct {
	db *gorm.DB
}

func NewOrderDao() *OrderDao {
	return &OrderDao{db: db.DbClient}
}

func (dao *OrderDao) AddOrder(order *model.Order) (int64, error) {
	result := dao.db.Create(order)
	if result.Error != nil {
		return 0, fmt.Errorf("adding order %v: %v", *order, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *OrderDao) DeleteOrderByID(ID uint) (int64, error) {
	result := dao.db.Delete(&model.Order{}, ID)
	if result.Error != nil {
		return 0, fmt.Errorf("deleting order by id=%d: %v", ID, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *OrderDao) UpdateOrder(order *model.Order) (int64, error) {
	result := dao.db.Save(order)
	if result.Error != nil {
		return 0, fmt.Errorf("updating order %v: %v", *order, result.Error)
	}
	return result.RowsAffected, nil
}

func (dao *OrderDao) SelectOrderByID(ID uint) (*model.Order, error) {
	order := model.Order{}
	result := dao.db.First(&order, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting order by id=%d: %v", ID, result.Error)
	}
	return &order, nil
}

func (dao *OrderDao) SelectOrdersByRoomIDTimeRange(roomID uint, begin, end time.Time) (*[]model.Order, error) {
	var orders []model.Order
	result := dao.db.Where("room_id = ? AND order_time BETWEEN ? AND ?", roomID, begin, end).Find(&orders)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting orders by roomID=%d and time between %v-%v: %v",
			roomID, begin, end, result.Error)
	}
	return &orders, nil
}

func (dao *OrderDao) SelectOrdersByRoomID(roomID uint) (*[]model.Order, error) {
	var orders []model.Order
	result := dao.db.Where(&model.Order{RoomID: roomID}).Find(&orders)
	if result.Error != nil {
		return nil, fmt.Errorf("selecting orders by roomID=%d: %v", roomID, result.Error)
	}
	return &orders, nil
}

func (dao *OrderDao) SelectValue(expr string, roomID uint, timeRange uint) (float64, error) {
	r := struct{ Result float64 }{}
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
	result := dao.db.Table(model.Order{}.TableName()).
		Select("("+expr+") Result").
		Where(where, param...).
		Scan(&r)
	if result.Error != nil {
		return 0, fmt.Errorf("selecting value(expr=%s): %v", expr, result.Error)
	}
	return r.Result, nil
}
