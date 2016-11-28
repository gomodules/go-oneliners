package oneliners

import (
	"fmt"
	"log"
	"runtime"
)

func FILE() {
	_, file, ln, ok := runtime.Caller(1)
	if ok {
		fmt.Println("__FILE__", file, "__LINE__", ln)
	} else {
		log.Fatal("Failed to detect runtime caller info.")
	}
}
