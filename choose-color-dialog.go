package osas

type ChooseColorDialog struct {
	DefaultColor []int
}

func NewChooseColorDialog() *ChooseColorDialog {
	return &ChooseColorDialog{}
}

func (c *ChooseColorDialog) SetDefaultColor256(r, g, b int) *ChooseColorDialog {
	return c.SetDefaultColor65535(r*256, g*256, b*256)
}

func (c *ChooseColorDialog) SetDefaultColor65535(r, g, b int) *ChooseColorDialog {
	c.DefaultColor = []int{r, g, b}
	return c
}

func (c *ChooseColorDialog) Run() (result Result, err error) {
	return RunOSAScript(chooseColorDialogTmpl, c)
}
