package main

import (
	"fmt"
	zl "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"k8s.io/klog"
	"log"
	"math/rand"
	"time"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	for {
		for i := 0; i < 10; i++ {
			zl.Error().Int("user_id", 15858).Msg("user try to break our system")
			zl.Error().Int("user_id", rand.Intn(1000000)).Msg("user try to break our system")
			zl.Info().Int("user_id", rand.Intn(1000000)).Msg("user logined")
			zl.Print("Test 123 my test")

			log.Print("standart log library")
			klog.Info("hello")
			klog.Error(fmt.Errorf("klog server error"))
			klog.V(0).Info("yaaa ", map[string]string{"ke": "value", "vehicle": "car"})

			logrus.Debug("My car ", map[string]string{"type": "people", "entry": "decline"})
			logrus.Print("Bad choice")
			logrus.Error("badly happened")

			log.Println("wanna ride?")
			logrus.Trace(rand.Intn(1000000))
		}
		time.Sleep(time.Second)
	}
}
