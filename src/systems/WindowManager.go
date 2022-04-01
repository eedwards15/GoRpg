package systems

import "log"

var (
	WINDOWMANAGER *WindowManager
)

type WindowManager struct {
	ScreenWidth  int
	ScreenHeight int
	XCenter      int
	YCenter      int
}

func InitWindowManger(width, height int) {
	if WINDOWMANAGER == nil {
		if width <= 0 {
			log.Fatalf("Missing Width")
		}

		if height <= 0 {
			log.Fatalf("Missing Height")
		}

		wm := &WindowManager{
			ScreenWidth:  width,
			XCenter:      width / 2,
			ScreenHeight: height,
			YCenter:      height / 2,
		}
		WINDOWMANAGER = wm
	}
}
