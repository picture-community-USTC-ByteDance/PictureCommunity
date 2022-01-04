package utils

import (
	"sync"
	"time"
)

const (
	workerBits  uint  = 6
	numberBits  uint  = 12
	flgBits     uint  = 4
	workerMax   uint  = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint  = workerBits + numberBits + flgBits
	workerShift uint  = numberBits
	numberShift uint  = flgBits
	startTime   int64 = 1640966400000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID  //2022年1月1日00:00:00
)

// id = 0 + timestamp(41) + workerID(6) + numberBits(12) + flagBits(4)

type IDGenerator struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	flag      int64
	number    int64
}

func NewIDGenerator(workerId int64, flag int64) *IDGenerator {
	return &IDGenerator{
		timestamp: 0,
		workerId:  workerId,
		flag:      flag,
		number:    0,
	}
}

func (w *IDGenerator) NewID() uint {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := (now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number << numberShift) | w.flag
	return uint(ID)
}
