/**
 * @author  qyanzh
 * @create  2022/03/08 1:33
 */

package gpool

import (
	"errors"
	"log"
	"sync"
)

type GPool struct {
	goroutines chan mGoroutine
	count      uint
	countLock  sync.Mutex
	total      uint
	done       chan struct{}
}

type mGoroutine struct {
	gid   uint
	works chan func()
}

func New() *GPool {
	return &GPool{
		goroutines: make(chan mGoroutine),
		done:       make(chan struct{}),
	}
}

func (gp *GPool) Enqueue(work func()) error {
	if gp.Closed() {
		return errors.New("GPool has been closed")
	}
	select {
	case g := <-gp.goroutines:
		g.works <- work
	default:
		gp.newGoroutine(work)
	}
	return nil
}

func (gp *GPool) newGoroutine(work func()) {
	gp.countLock.Lock()
	gp.count++
	gp.total++
	gp.countLock.Unlock()
	go func(gid uint) {
		log.Printf("created a new goroutine, gid=%d\n", gid)
		this := mGoroutine{gid: gid, works: make(chan func(), 1)}
		this.works <- work
		for {
			select {
			case w := <-this.works:
				log.Printf("gid=%d received a work.\n", this.gid)
				w()
				gp.goroutines <- this
			case <-gp.done:
				log.Printf("goroutine terminated, gid=%d\n", this.gid)
				return
			}
		}
	}(gp.count)
}

func (gp *GPool) Close() error {
	if gp.Closed() {
		return errors.New("close a closed GPool")
	}
	close(gp.done)
	for {
		gp.countLock.Lock()
		if gp.count == 0 {
			gp.countLock.Unlock()
			break
		} else {
			gp.count--
			gp.countLock.Unlock()
			_ = <-gp.goroutines
		}
	}
	close(gp.goroutines)
	log.Printf("GPool closed, %d goroutines in total.\n", gp.total)
	return nil
}

func (gp *GPool) Closed() bool {
	select {
	case <-gp.done:
		return true
	default:
		return false
	}
}
