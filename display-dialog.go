package osas

import (
	"fmt"
)

type DisplayDialog struct {
	Title            string
	Text             string
	HiddenAnswerFlag bool
	InputFieldFlag   bool
	DefaultAnswer    string
	Buttons          []string
	DefaultButton    string
	CancelButton     string
	WithIcon         string
	GivingUpAfter    int
}

func NewDisplayDialog(text string) *DisplayDialog {
	return &DisplayDialog{Text: fmt.Sprintf("%q", text)}
}

func (d *DisplayDialog) SetInputField(text string) *DisplayDialog {
	d.InputFieldFlag = true
	d.DefaultAnswer = fmt.Sprintf("%q", text)
	return d
}

func (d *DisplayDialog) SetDefaultCancel() *DisplayDialog {
	d.SetDefaultButton("キャンセル")
	return d
}

func (d *DisplayDialog) SetDefaultOK() *DisplayDialog {
	d.SetDefaultButton("OK")
	return d
}
func (d *DisplayDialog) SetHiddenAnswer() *DisplayDialog {
	d.HiddenAnswerFlag = true
	d.InputFieldFlag = true
	return d
}

func (d *DisplayDialog) SetTitle(text string) *DisplayDialog {
	d.Title = text
	return d
}

func (d *DisplayDialog) SetDefaultButton(text string) *DisplayDialog {
	d.DefaultButton = text
	return d
}

func (d *DisplayDialog) SetCancelButton(text string) *DisplayDialog {
	d.CancelButton = text
	return d
}

func (d *DisplayDialog) SetIcnsFilePath(icnsFilePath string) *DisplayDialog {
	d.WithIcon = icnsFilePath
	return d
}

func (d *DisplayDialog) Run() (result Result, err error) {
	return RunOSAScript(displayDialogTmpl, d)
}
