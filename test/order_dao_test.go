/**
 * @author  qyanzh
 * @create  2022/02/26 21:38
 */

package test

import (
	"alert/internal/dao"
	"alert/internal/model"
	"fmt"
	"testing"
	"time"
)

var orderDao dao.OrderDao

func init() {
	orderDao = *dao.NewOrderDao()
}

func TestAddOrder(t *testing.T) {
	order := model.Order{}
	order.RoomID = 0
	order.Turnover = 10
	order.OrderTime = time.Now()
	fmt.Println(orderDao.AddOrder(&order))
	fmt.Printf("%+v", orderDao.SelectOrderByID(order.ID))
}

func TestSelectOrdersByRoomID(t *testing.T) {
	fmt.Println(orderDao.SelectOrdersByRoomID(0))
}

func TestDeleteOrdersByID(t *testing.T) {
	orderDao.DeleteOrderByID(5)
}

func TestUpdateOrder(t *testing.T) {
	order := orderDao.SelectOrderByID(4)
	order.RoomID = 1
	orderDao.UpdateOrder(order)
}

func TestSelectOrderByTimeRange(t *testing.T) {
	end := time.Now()
	begin := time.Now().Add(-120 * time.Minute)
	orders := orderDao.SelectOrderByRoomIDTimeRange(0, begin, end)
	fmt.Println(orders)
}

func TestSelectValue(t *testing.T) {
	r := orderDao.SelectValue("sum(turnover)", 0, 60*30)
	fmt.Println(r)
}
