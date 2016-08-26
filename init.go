package osas

import (
	"fmt"
	"path/filepath"
	"text/template"
)

func init() {
	var t *template.Template
	t = template.New("choose-dialog")
	t.Funcs(template.FuncMap{
		"ListFunc": func(items []string) string {
			return StringsToListStrings(items)
		},
		"WithPromptFunc":                MakeAddPrefixFunc("with prompt"),
		"OK_ButtonNameFunc":             MakeAddPrefixFunc("OK button name"),
		"CancelButtonNameFunc":          MakeAddPrefixFunc("cancel button name"),
		"MultipleSelectionsAllowedFunc": MakeBoolAddPrefixFunc("multiple selections allowed"),
		"EmptySelectionAllowedFunc":     MakeBoolAddPrefixFunc("empty selection allowed"),
	})
	template.Must(t.Parse(chooseDialogTmplText))
	chooseDialogTmpl = t

	t = template.New("display-dialog")
	t.Funcs(template.FuncMap{
		"TitleFunc":         MakeAddPrefixFunc("with title"),
		"ButtonsFunc":       MakeListAddPrefixFunc("buttons"),
		"DefaultButtonFunc": MakeAddPrefixFunc("default button"),
		"CancelButtonFunc":  MakeAddPrefixFunc("cancel button"),
		"WithIconFunc": func(input string) string {
			if input == "" {
				return ""
			}
			//	Mac従来のパス形式	コロン区切り
			//	通常のパス形式		POSIX 形式
			if filepath.Ext(input) == ".icns" {
				//	絶対パスに変換
				if !filepath.IsAbs(input) {
					dir, err := filepath.Abs(filepath.Dir(input))
					if err != nil {
						return ""
					}
					input = filepath.Join(dir, input)
				}
				return "with icon " + fmt.Sprintf("%q", input) + " as POSIX file"
			}
			return "with icon " + input
		},
		"GivingUpAfterFunc": MakeIntAddPrefixFunc("giving up after"),
	})
	template.Must(t.Parse(diplayDialogTmplText))
	displayDialogTmpl = t

	t = template.New("choose-file-name-dialog")
	t.Funcs(template.FuncMap{
		"PromptFunc":          MakeAddPrefixFunc("with prompt"),
		"DefaultNameFunc":     MakeAddPrefixFunc("default name"),
		"DefaultLocationFunc": MakeAddPrefixFunc("default location"),
	})
	template.Must(t.Parse(chooseFileNameDialogTmplText))
	chooseFileNameDialogTmpl = t

	t = template.New("choose-file-dialog")
	t.Funcs(template.FuncMap{
		"TypesFunc":                     MakeListAddPrefixFunc("of type"),
		"PromptFunc":                    MakeAddPrefixFunc("with prompt"),
		"DefaultLocationFunc":           MakeAddPrefixFunc("default location"),
		"MultipleSelectionsAllowedFunc": MakeBoolAddPrefixFunc("multiple selections allowed"),
		"ShowingPackageContentsFunc":    MakeBoolAddPrefixFunc("showing package contents"),
		"InvisiblesFlagFunc":            MakeBoolAddPrefixFunc("invisibles"),
	})
	template.Must(t.Parse(chooseFileDialogTmplText))
	chooseFileDialogTmpl = t

	t = template.New("choose-folder-dialog")
	t.Funcs(template.FuncMap{
		"PromptFunc":                    MakeAddPrefixFunc("with prompt"),
		"DefaultLocationFunc":           MakeAddPrefixFunc("default location"),
		"MultipleSelectionsAllowedFunc": MakeBoolAddPrefixFunc("multiple selections allowed"),
		"ShowingPackageContentsFunc":    MakeBoolAddPrefixFunc("showing package contents"),
		"InvisiblesFlagFunc":            MakeBoolAddPrefixFunc("invisibles"),
	})
	template.Must(t.Parse(chooseFolderDialogTmplText))
	chooseFolderDialogTmpl = t

	t = template.New("choose-color-dialog")
	t.Funcs(template.FuncMap{
		"DefaultColorFunc": MakeIntListAddPrefixFunc("default color"),
	})
	template.Must(t.Parse(chooseColorDialogTmplText))
	chooseColorDialogTmpl = t
}
