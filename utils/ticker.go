package utils

import "time"

// SetInterval 设置一个定时器，类似 js 的 SetInterval，返回一个用于关闭该定时器的通道
func SetInterval(second int, callback func(), onClose func()) chan bool {
	closeChan := make(chan bool, 1)
	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(second))
		for {
			select {
			case <-ticker.C:
				callback()
			case <-closeChan:
				onClose()
				return
			}
		}
	}()
	return closeChan
}
