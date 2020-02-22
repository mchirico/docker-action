package pkg

import (
	"fmt"
	"testing"
)

func TestSpeak(t *testing.T) {

	s := []string{"one","two"}
	fmt.Printf("here: %v\n",Speak(s))
}
