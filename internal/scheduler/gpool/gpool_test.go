/**
 * @author  qyanzh
 * @create  2022/03/08 13:56
 */

package gpool

import (
	"fmt"
	"testing"
	"time"
)

func TestGPool(t *testing.T) {
	// 新的协程池
	gp := New()

	// 指定时间跳出for循环
	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	// 任务编号
	i := 0
	for {
		select {
		case <-done:
			_ = gp.Close()
			return
		default:
			t := i
			i++
			_ = gp.Enqueue(func() {
				// 每秒处理一个任务
				time.Sleep(time.Second)
				wid := fmt.Sprintf("work=%d: ", t)
				fmt.Println(wid + time.Now().String())
			})
			// 间隔一段时间发送一个任务
			time.Sleep(time.Millisecond * 10)
		}
	}
}
