package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

type Interceptor struct {
	Headers map[string][]string    `json:"headers,omitempty"`
	Body    map[string]interface{} `json:"body,omitempty"`
	UUID    string                 `json:"UUID,omitempty"`
	APIPath string                 `json:"APIPath,omitempty"`
	EnvID   string                 `json:"envID,omitempty"`
	Status  int                    `json:"status,omitempty"`
}

var (
	debug = flag.Bool("debug", false, "Debug")
)

// CheckDebug if Debug flag is set
func CheckDebug() {
	flag.Parse()
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetOutput(os.Stdout)
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)
}

func main() {
	logrus.Infof("Starting Rancher API Interceptor")
	CheckDebug()

	//router := NewRouter()
	http.HandleFunc("/", Index)        //router
	http.HandleFunc("/secret", Secret) //router
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		logrus.Fatal("Error: ", err)
	}
	//err := http.ListenAndServe(":8000", router)
	//if err != nil {
	//	logrus.Fatal("Error: ", err)
	//}
}
