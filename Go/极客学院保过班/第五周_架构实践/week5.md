# 第五周作业



**使用滑动窗口实现计数器限流**

```go
package main

import (
	"time"
)

type counterLimiter struct {
	timestamp int64
	windows []int
	windowCount int
	windowSize int64
	start int
	count int64
	limit int64
}

func (ct *counterLimiter) init(windowCnt int, windowsize int64) *counterLimiter {
	arr := make([]int, 0) // 咋改啊6

	for i:=0;i < int(windowsize);i++ {
		arr = append(arr, 0)
	}
	return &counterLimiter{
		windowCount: windowCnt,
		windowSize:  windowsize,
		limit: int64(windowCnt * windowCnt),
		start:       0,
		timestamp: time.Now().Unix(),
		windows: arr,
	}
}

func (ct *counterLimiter) tryAcquire()  bool{
	now := time.Now().Unix()
	time := now-ct.timestamp
	if time < ct.limit{
		if ct.count < ct.limit{
			ct.count = ct.count+1
			offset := int(ct.start+(int(time/ct.windowSize)%ct.windowCount))
			ct.windows[offset]++
			return true
		}else {
			return false
		}
	}else {
		diffWindow := time/ct.windowSize
		if diffWindow < int64(ct.windowCount*2){
			i := int64(0)
			for ; i < diffWindow-int64(ct.windowCount); i++ {
				idx := ct.start+1
				if idx > ct.windowCount{
					idx %= ct.windowCount
				}
				ct.count += int64((-1)*ct.windows[idx])
				ct.windows[idx] = 0
			}
			if i >= int64(ct.windowCount){
				i %= int64(ct.windowCount)
			}
			ct.windows[i]++
			return true
		}else {
			for i := 0; i < int(ct.windowSize); i++ {
				ct.windows[i] = 0
			}
			ct.count = 0
			return true
		}
	}


}


```

