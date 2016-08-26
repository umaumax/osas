package main

import (
	"fmt"

	"github.com/umaumax/osas"
)

func main() {
	displayDialog()
	return
	example()
}

func example() {
	var result osas.Result
	var err error
	result, err = osas.NewDisplayDialog("Are you ok?").Run()
	if err == nil && result.IsOK() {
		chooseDialog := osas.NewChooseDialog()
		chooseDialog.SelectItems = []string{"ONE", "TWO", "THREE"}
		result, err = chooseDialog.Run()
		if err == nil && result.IsOK() {
			osas.NewDisplayDialog(fmt.Sprintf("You choose %s", result.GetOKText())).Run()
		}
	} else {
		fmt.Println(err)
	}
}

func chooseColor() {
	result, err := osas.NewChooseColorDialog().SetDefaultColor256(255, 255, 20).Run()
	if err == nil && result.IsOK() {
		r, g, b := result.GetRGB()
		fmt.Println(r, g, b)
	}
}

func chooseDialog() {
	result, err := osas.NewChooseDialog().SetSelectItems([]string{"NANOHA", "FATE", "HAYATE"}...).SetMultipleSelectionsAllowed(true).Run()
	if err == nil && result.IsOK() {
		for i, v := range result.GetOKTexts() {
			fmt.Println(i, v)
		}
	}
}
func chooseFileDialog() {
	fmt.Println(osas.NewChooseFileDialog().SetMultipleSelectionsAllowed(true).SetTypes("txt").Run())
	return
}
func displayDialog() {
	fmt.Println(osas.NewDisplayDialog("niconico").SetIcnsFilePath("nico.icns").Run())
}
func displayDialogButton() {
	dialog := osas.NewDisplayDialog("nya-")
	dialog.Buttons = []string{"ONE", "TWO", "THREE"}
	dialog.SetCancelButton("ONE")
	fmt.Println(dialog.Run())
	return
}
