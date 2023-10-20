package data

import (
	"botdetector/domain"
	"log"
	"sync"
	"time"
)

type asyncWriter struct {
	buffer []domain.RequestLog
	locker sync.Mutex
}

var AsyncWriter = asyncWriter{
	buffer: make([]domain.RequestLog, 0, 15000),
}

func (aw *asyncWriter) Start() {
	for {
		time.Sleep(time.Second * 5)
		if len(aw.buffer) == 0 {
			continue
		}
		aw.locker.Lock()
		bufferOld := aw.buffer
		aw.buffer = make([]domain.RequestLog, 0, 15000)
		aw.locker.Unlock()

		err := Database.InsertLogBatch(bufferOld)

		if err != nil {
			log.Println("Error inserting logs to database: " + err.Error())
		}
	}
}

func (aw *asyncWriter) Add(log domain.RequestLog) {
	aw.locker.Lock()
	defer aw.locker.Unlock()

	aw.buffer = append(aw.buffer, log)
}
