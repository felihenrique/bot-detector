package data

import (
	"botdetector/app"
	"log"
	"sync"
	"time"
)

type asyncWriter struct {
	buffer []app.RequestLog
	locker sync.Mutex
}

var AsyncWriter = asyncWriter{
	buffer: make([]app.RequestLog, 0, 10000),
}

func (aw *asyncWriter) Start() {
	for {
		time.Sleep(time.Second * 5)
		if len(aw.buffer) == 0 {
			continue
		}
		aw.locker.Lock()
		bufferOld := aw.buffer
		aw.buffer = make([]app.RequestLog, 0, 10000)
		aw.locker.Unlock()

		err := Database.InsertRequestLogs(bufferOld)

		if err != nil {
			log.Println("Error inserting logs to database: " + err.Error())
		}
	}
}

func (aw *asyncWriter) Add(log app.RequestLog) {
	aw.locker.Lock()
	defer aw.locker.Unlock()

	aw.buffer = append(aw.buffer, log)
}
