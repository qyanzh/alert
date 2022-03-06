/**
 * @author  qyanzh
 * @create  2022/02/26 21:38
 */

package order

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
	order := &model.Order{}
	order.RoomID = 0
	order.Turnover = 10
	order.OrderTime = time.Now()
	fmt.Println(orderDao.AddOrder(order))
	order, _ = orderDao.SelectOrderByID(order.ID)
	fmt.Printf("%+v", order)
}

func TestSelectOrdersByRoomID(t *testing.T) {
	fmt.Println(orderDao.SelectOrdersByRoomID(0))
}

func TestDeleteOrdersByID(t *testing.T) {
	orderDao.DeleteOrderByID(5)
}

func TestUpdateOrder(t *testing.T) {
	order, _ := orderDao.SelectOrderByID(4)
	order.RoomID = 1
	orderDao.UpdateOrder(order)
}

func TestSelectOrderByTimeRange(t *testing.T) {
	end := time.Now()
	begin := time.Now().Add(-120 * time.Minute)
	orders, _ := orderDao.SelectOrderByRoomIDTimeRange(0, begin, end)
	fmt.Println(orders)
}

func TestSelectValue(t *testing.T) {
	r, _ := orderDao.SelectValue("sum[turnover]", 0, 60*30)
	fmt.Println(r)
}
