package racing

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(u1, u2 string) (fasterUrl string) {
	start1 := time.Now()
	http.Get(u1)
	duration1 := time.Since(start1)
	fmt.Println(u1, duration1)

	start2 := time.Now()
	http.Get(u2)
	duration2 := time.Since(start2)
	fmt.Println(u2, duration2)

	if duration2 < duration1 {
		return u2
	} else {
		return u1
	}

}
