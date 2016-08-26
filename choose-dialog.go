package osas

type ChooseDialog struct {
	DefaultItems              []string
	SelectItems               []string
	Prompt                    string
	MultipleSelectionsAllowed bool
	EmptySelectionAllowed     bool
	OK_ButtonName             string
	CancelButtonName          string
}

func NewChooseDialog() *ChooseDialog {
	return &ChooseDialog{}
}

func (c *ChooseDialog) SetSelectItems(item ...string) *ChooseDialog {
	c.SelectItems = item
	return c
}

func (c *ChooseDialog) SetPrompt(text string) *ChooseDialog {
	c.Prompt = text
	return c
}

func (c *ChooseDialog) SetMultipleSelectionsAllowed(b bool) *ChooseDialog {
	c.MultipleSelectionsAllowed = b
	return c
}

func (c *ChooseDialog) Run() (result Result, err error) {
	return RunOSAScript(chooseDialogTmpl, c)
}
