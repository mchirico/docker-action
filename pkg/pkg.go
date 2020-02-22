package pkg

import (
	"fmt"
	"time"
)

func Speak(args []string) string {
	current := time.Now()
	return fmt.Sprintf("::set-output name=time::%v %s\n",args,current)
}
