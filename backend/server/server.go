package server

import (
	"botdetector/config"
	"net/http"
)

func Start() {
	err := http.ListenAndServe("0.0.0.0:"+config.Config.Port, nil)

	if err != nil {
		panic("Can't start the http server: " + err.Error())
	}

	print("Server started at port " + config.Config.Port)
}
