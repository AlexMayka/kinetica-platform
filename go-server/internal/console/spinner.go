package console

import (
	"fmt"
	"time"
)

func PrintStatusBar(message string) chan bool {
	done := make(chan bool)
	go func(msg string) {
		chars := []rune{'|', '/', '-', '\\'}
		i := 0
		for {
			select {
			case result := <-done:
				clear := "\r\033[K"
				if result {
					fmt.Printf("%s%s ✅ \n", clear, msg)
				} else {
					fmt.Printf("%s%s ❌ \n", clear, msg)
				}
				close(done)
				return
			default:
				fmt.Printf("\r%s %c", msg, chars[i%len(chars)])
				time.Sleep(100 * time.Millisecond)
				i++
			}
		}
	}(message)
	return done
}
