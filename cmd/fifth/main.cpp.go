package main

import (
	"log"
	"math"
	"time"
)

type slideWindowLimiter struct {
	windowSize int64   //窗口大小，毫秒为单位
	limit      int64   //窗口内限流大小
	splitNum   int64   //切分小窗口的数目大小
	counters   []int64 //每个小窗口的计数数组
	index      int64   //当前小窗口计数器的索引
	startTime  int64   //窗口开始时间
}

func (s *slideWindowLimiter) initCounts() {
	for i := 0; i < int(s.splitNum); i++ {
		s.counters[i] = 0
	}
	return
}

func (s *slideWindowLimiter) tryAcquire() bool {
	curTime := time.Now().UnixNano() / 1000000
	dur := int64(curTime) - s.startTime
	dur = dur - s.windowSize
	if dur < 0 {
		dur = 0
	}
	//log.Printf("curTime:%d - starttime:%d - windowsize:%d = dur:%d ", curTime, s.startTime, s.windowSize, dur)

	secWindowSize := s.windowSize / s.splitNum

	windowNum := dur / secWindowSize // 计算滑动的小窗口的数量
	//log.Printf("windowNum:%d ", windowNum)

	s.slideWindow(windowNum) //滑动窗口
	//log.Printf("curr index:%d", s.index)

	var count int64 = 0
	for i := 0; i < int(s.splitNum); i++ {
		count = count + s.counters[i]
		//log.Printf("count:%d = count:%d + counters[i]:%d ", count, count, s.counters[i])
	}
	//log.Printf("count:%d ", count)

	if count >= s.limit {
		return false
	} else {
		s.counters[s.index]++
		return true
	}
}

func (s *slideWindowLimiter) slideWindow(windowsNum int64) error {
	if windowsNum == 0 {
		return nil
	}

	slideNum := s.splitNum
	if windowsNum < s.splitNum {
		slideNum = windowsNum
	}
	//log.Printf("slideNum:%d, windowsNum:%d, splitNum:%d", slideNum, windowsNum, s.splitNum)

	for i := 0; i < int(slideNum); i++ { // 这里可以想象成一个首位相连的环形数组，index可以看成永远指向最末一个小窗口
		s.index = (s.index + 1) % s.splitNum
		s.counters[s.index] = 0
		//log.Printf("calc index:%d", s.index)
	}

	s.startTime = s.startTime + (windowsNum * (s.windowSize / s.splitNum))
	//log.Printf("startTime:%d = startTime:%d + (windowsNum:%d * %d) ", s.startTime, s.startTime, windowsNum, (s.windowSize / s.splitNum))

	return nil
}

func main() {

	//每秒20个请求
	var limitNum int64 = 20
	var limiter = &slideWindowLimiter{windowSize: 1000, limit: limitNum, splitNum: 10, counters: []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, index: 0, startTime: int64(time.Now().UnixNano() / 1000000)}
	// 或者
	//var limiter = &slideWindowLimiter{windowSize: 1000, limit: limitNum, splitNum: 10, counters: []int64{0,}, index: 0, startTime: int64(time.Now().UnixNano() / 1000000)}
	//limiter.initCounts()

	var count int64 = 0

	time.Sleep(time.Millisecond * 3000)
	//计数器滑动窗口算法模拟100组间隔30ms的50次请求
	log.Println("计数器滑动窗口算法测试开始")
	log.Println("开始模拟100组间隔150ms的50次请求")
	var faliCount = 0
	for j := 0; j < 100; j++ {
		count = 0
		//for i := 0; i < 50; i++ {
		for i := 0; i < 15; i++ {
			if limiter.tryAcquire() {
				count++
				log.Printf("count1:%d", count)
			}
		}

		time.Sleep(time.Millisecond * 150)
		//模拟50次请求，看多少能通过
		for i := 0; i < 50; i++ {
			if limiter.tryAcquire() {
				count++
				log.Printf("count2:%d", count)
			}
		}

		if count > limiter.limit {
			/// 理论上不存在这种情况，//实测下来也是不存在这种情况
			log.Printf("时间窗口内放过的请求超过阈值，放过的请求数:%d, 限流:%d", count, limiter.limit)
			faliCount++
		}

		//time.Sleep(time.Millisecond * 100)
		time.Sleep(time.Duration(math.Round(10)*100) * time.Millisecond)
	}
	/// 理论上是0组 都可以限制住 //实测符合理论
	log.Println("计数器滑动窗口算法测试结束，100组间隔150ms的50次请求模拟完成，限流失败组数：", faliCount)
	log.Println("===========================================================================================")
}

