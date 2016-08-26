package osas

type ChooseFolderDialog struct {
	Prompt                    string
	DefaultLocation           string
	InvisiblesFlag            bool
	MultipleSelectionsAllowed bool
	ShowingPackageContents    bool
}

func NewChooseFolderDialog() *ChooseFolderDialog {
	return &ChooseFolderDialog{}
}

func (c *ChooseFolderDialog) SetPrompt(text string) *ChooseFolderDialog {
	c.Prompt = text
	return c
}
func (c *ChooseFolderDialog) SetDefaultLocation(text string) *ChooseFolderDialog {
	c.DefaultLocation = text
	return c
}

func (c *ChooseFolderDialog) SetInvisiblesFlag(b bool) *ChooseFolderDialog {
	c.InvisiblesFlag = b
	return c
}
func (c *ChooseFolderDialog) SetMultipleSelectionsAllowed(b bool) *ChooseFolderDialog {
	c.MultipleSelectionsAllowed = b
	return c
}
func (c *ChooseFolderDialog) SetShowingPackageContents(b bool) *ChooseFolderDialog {
	c.ShowingPackageContents = b
	return c
}

func (c *ChooseFolderDialog) Run() (result Result, err error) {
	return RunOSAScript(chooseFolderDialogTmpl, c)
}
