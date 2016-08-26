package osas

import ()

type ChooseFileDialog struct {
	Prompt                    string
	Types                     []string
	DefaultLocation           string
	InvisiblesFlag            bool
	MultipleSelectionsAllowed bool
	ShowingPackageContents    bool
}

func NewChooseFileDialog() *ChooseFileDialog {
	return &ChooseFileDialog{}
}

func (c *ChooseFileDialog) SetPrompt(text string) *ChooseFileDialog {
	c.Prompt = text
	return c
}
func (c *ChooseFileDialog) SetTypes(types ...string) *ChooseFileDialog {
	c.Types = types
	return c
}
func (c *ChooseFileDialog) SetDefaultLocation(text string) *ChooseFileDialog {
	c.DefaultLocation = text
	return c
}

func (c *ChooseFileDialog) SetInvisiblesFlag(b bool) *ChooseFileDialog {
	c.InvisiblesFlag = b
	return c
}
func (c *ChooseFileDialog) SetMultipleSelectionsAllowed(b bool) *ChooseFileDialog {
	c.MultipleSelectionsAllowed = b
	return c
}
func (c *ChooseFileDialog) SetShowingPackageContents(b bool) *ChooseFileDialog {
	c.ShowingPackageContents = b
	return c
}

func (c *ChooseFileDialog) Run() (result Result, err error) {
	return RunOSAScript(chooseFileDialogTmpl, c)
}
