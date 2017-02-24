package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

type Interceptor struct {
	Headers   map[string][]string    `json:"headers,omitempty"`
	Body      map[string]interface{} `json:"body,omitempty"`
	UUID      string                 `json:"UUID,omitempty"`
	APIPath   string                 `json:"APIPath,omitempty"`
	APIMethod string                 `json:"APIMethod,omitempty"`
	EnvID     string                 `json:"envID,omitempty"`
	Status    int                    `json:"status,omitempty"`
	Message   string                 `json:"message,omitempty"`
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

	http.HandleFunc("/", Index)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/authtokenvalidator", Auth)
	http.HandleFunc("/modifystackname", ModifyBody)
	http.HandleFunc("/finaldestination/", Destination)
	http.HandleFunc("/secret1", ChainedSecret1)
	http.HandleFunc("/secret2", ChainedSecret2)
	http.HandleFunc("/blockuser", BlockLDAPUser)
	http.HandleFunc("/sleep", Sleepy)
	http.HandleFunc("/unhandled", Unhandled)
	http.HandleFunc("/failure", Failure)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		logrus.Fatal("Error: ", err)
	}
}
