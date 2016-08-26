package osas

type ChooseFileNameDialog struct {
	Prompt          string
	DefaultName     string
	DefaultLocation string
}

func NewChooseFileNameDialog() *ChooseFileNameDialog {
	return &ChooseFileNameDialog{}
}

func (c *ChooseFileNameDialog) SetPrompt(text string) *ChooseFileNameDialog {
	c.Prompt = text
	return c
}
func (c *ChooseFileNameDialog) SetDefaultName(text string) *ChooseFileNameDialog {
	c.DefaultName = text
	return c
}
func (c *ChooseFileNameDialog) SetDefaultLocation(text string) *ChooseFileNameDialog {
	c.DefaultLocation = text
	return c
}

func (c *ChooseFileNameDialog) Run() (result Result, err error) {
	return RunOSAScript(chooseFileNameDialogTmpl, c)
}
