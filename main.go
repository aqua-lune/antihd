package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/AllenDang/giu"
	"image"
	"strconv"
	"time"
)

//go:embed catnobrain.png
var icon []byte

var steps []string
var stepsCheck []bool

var newStep string
var shouldFocus bool
var showDialog bool

var wnd *giu.MasterWindow

var beginTime time.Time

func stepsRender() giu.Widget {
	stepsWidgets := make([]giu.Widget, len(steps))
	for i, step := range steps {
		stepsWidgets[i] = giu.Checkbox(step, &stepsCheck[i])
	}
	return giu.Column(stepsWidgets...)
}

func appendNewStep() {
	steps = append(steps, newStep)
	newStep = ""
	shouldFocus = true
}

func getCompleted() (n int) {
	nbStep := len(steps)
	for i := 0; i < nbStep; i++ {
		if stepsCheck[i] == true {
			n += 1
		}
	}
	return n
}

func FormatDuration(d time.Duration) string {
	hour := int(d.Hours())
	minute := int(d.Minutes()) % 60
	second := int(d.Seconds()) % 60
	if hour > 0 {
		return fmt.Sprintf("%02dh%02d", hour, minute)
	} else {
		return fmt.Sprintf("%02dm%02ds", minute, second)
	}

}

func loop() {
	strTasksCompleted := strconv.Itoa(getCompleted()) + "/" + strconv.Itoa(len(steps))
	elapsed := time.Since(beginTime)
	strElapsed := FormatDuration(elapsed)
	giu.SingleWindow().Layout(
		giu.PrepareMsgbox(),
		giu.Row(
			giu.Align(giu.AlignLeft).To(
				giu.Label("Liste des tâches :"),
			),
			giu.Align(giu.AlignRight).To(
				giu.Label(strTasksCompleted),
			),
		),
		giu.Child().Size(280, 360).Layout(
			stepsRender(),
		),
		giu.Custom(func() {
			if shouldFocus {
				giu.SetKeyboardFocusHere()
				shouldFocus = false
			}
		}),
		giu.Row(
			giu.InputText(&newStep).Flags(giu.InputTextFlagsEnterReturnsTrue).OnChange(appendNewStep).Size(227),
			giu.Button("Add").OnClick(appendNewStep),
		),
		giu.Align(giu.AlignRight).To(
			giu.Label("Temps écoulé : "+strElapsed),
		),
		giu.Custom(func() {
			if showDialog {
				showDialog = false
				giu.Msgbox("Attention", "Toutes les tâches n'ont pas été complétées...")
			}
		}),
	)
}

func main() {
	stepsCheck = make([]bool, 128)
	shouldFocus = true
	wnd = giu.NewMasterWindow("ANTIHD", 300, 460,
		giu.MasterWindowFlagsFloating|giu.MasterWindowFlagsNotResizable,
	)
	img, _, _ := image.Decode(bytes.NewReader(icon))
	rgba := giu.ImageToRgba(img)
	wnd.SetIcon(rgba)
	wnd.SetCloseCallback(func() bool {
		nbStep := len(steps)
		for i := 0; i < nbStep; i++ {
			if stepsCheck[i] == false {
				showDialog = true
				return false
			}
		}
		return true
	})

	go func() {
		for range time.Tick(time.Second) {
			giu.Update()
		}
	}()

	beginTime = time.Now()
	wnd.Run(loop)
}
