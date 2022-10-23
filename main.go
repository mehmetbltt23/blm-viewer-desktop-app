package main

import (
	"fmt"
	"github.com/mehmetbltt23/blm-viewer-desktop-app/blm"
)

var data = [][]string{
	[]string{
		"Address1", "Address2", "Address3", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4",
	},
	[]string{
		"Address1", "Address2", "Address3", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4",
	},
	[]string{
		"Address1", "Address2", "Address3", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4",
	},
	[]string{
		"Address1", "Address2", "Address3", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4",
	},
	[]string{
		"Address1", "Address2", "Address3", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4",
	},
	[]string{
		"Address1", "Address2", "Address3", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4",
	},
	[]string{
		"Address1", "Address2", "Address3", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4", "Address4",
	},
}

func main() {
	reader, err := blm.NewReader("test_data/test.BLM")
	if err != nil {
		panic(err)
	}

	fmt.Println(reader.GetHeaders())
	/*myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")
	myWindow.Resize(fyne.Size{1000, 1000})
	myWindow.FixedSize()
	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			fmt.Println(i.Col, i.Row)
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()*/
}
