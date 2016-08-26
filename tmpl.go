package osas

import (
	"text/template"
)

var chooseDialogTmpl *template.Template
var chooseFileNameDialogTmpl *template.Template
var chooseFileDialogTmpl *template.Template
var chooseFolderDialogTmpl *template.Template
var chooseColorDialogTmpl *template.Template
var displayDialogTmpl *template.Template

var chooseDialogTmplText = `
on joinWithCR(theList)
	set argStr to ""
	repeat with v in theList
		set argStr to argStr & v & "
"
	end repeat
	return argStr
end joinWithCR

on run argv 
	tell application "System Events" to tell process "Finder"
	activate
		try
			set answer to (choose from list {{.SelectItems | ListFunc}} default items {{.DefaultItems | ListFunc}} {{.Prompt | WithPromptFunc}} {{.MultipleSelectionsAllowed | MultipleSelectionsAllowedFunc}} {{.OK_ButtonName | OK_ButtonNameFunc}} {{.CancelButtonName | CancelButtonNameFunc}} {{.EmptySelectionAllowed | EmptySelectionAllowedFunc}} )
			if (class of answer) as string = "boolean" then
				set answer to "CANCEL:"
			else
				try
					set answer to "OK:" & my joinWithCR(result)
				on error
					set answer to "OK:" & answer
				end try
			end if
		on error
			set answer to "CANCEL:"
		end try
	end tell
end run
`

var diplayDialogTmplText = `
tell application "System Events" to tell process "Finder"
	activate
	try
		display dialog {{.Text}} {{if .InputFieldFlag}} default answer {{.DefaultAnswer}} {{end}} {{.Buttons | ButtonsFunc}} {{.DefaultButton | DefaultButtonFunc}} {{.WithIcon | WithIconFunc}} {{.GivingUpAfter | GivingUpAfterFunc}} {{.Title |TitleFunc}} {{if .HiddenAnswerFlag}} hidden answer {{.HiddenAnswerFlag}} {{end}} {{.CancelButton | CancelButtonFunc}}
		try
			if gave up of result then
				try
					set answer to "TIMEOUT:" & text returned of result
				on error
					set answer to "TIMOUT:" & button returned of result
				end try
			else
				try
					set answer to "OK:" & text returned of result
				on error
					set answer to "OK:" & button returned of result
				end try
			end if
		on error
			try
				set answer to "OK:" & text returned of result
			on error
				set answer to "OK:" & button returned of result
			end try
		end try
	on error
		set answer to "CANCEL:"
	end try
end tell
`
var chooseFileNameDialogTmplText = `
try
	choose file name {{.Prompt | PromptFunc}} {{.DefaultName | DefaultNameFunc}} {{.DefaultLocation | DefaultLocationFunc}}
	set answer to "OK:" & POSIX path of result
on error
	set answer to "CANCEL:"
end try
`

var chooseFileDialogTmplText = `
on joinWithCR(theList)
	set argStr to ""
	repeat with v in theList
		set argStr to argStr & (POSIX path of v) & "
"
	end repeat
	return argStr
end joinWithCR

try
	choose file {{.Prompt | PromptFunc}} {{.DefaultLocation | DefaultLocationFunc}} {{.Types | TypesFunc}} {{.InvisiblesFlag | InvisiblesFlagFunc}} {{.MultipleSelectionsAllowed | MultipleSelectionsAllowedFunc}} {{.ShowingPackageContents | ShowingPackageContentsFunc}}
	set answer to "OK:" & POSIX path of result
on error
	try
	set answer to "OK:" & my joinWithCR(result)
	on error
		set answer to "CANCEL:"
	end try
end try
`

var chooseFolderDialogTmplText = `
on joinWithCR(theList)
	set argStr to ""
	repeat with v in theList
		set argStr to argStr & (POSIX path of v) & "
"
	end repeat
	return argStr
end joinWithCR

try
	choose folder {{.Prompt | PromptFunc}} {{.DefaultLocation | DefaultLocationFunc}} {{.InvisiblesFlag | InvisiblesFlagFunc}} {{.MultipleSelectionsAllowed | MultipleSelectionsAllowedFunc}} {{.ShowingPackageContents | ShowingPackageContentsFunc}}
	set answer to "OK:" & POSIX path of result
on error
	try
	set answer to "OK:" & my joinWithCR(result)
	on error
		set answer to "CANCEL:"
	end try
end try
`

var chooseColorDialogTmplText = `
on joinWithCR(theList)
	set argStr to ""
	repeat with v in theList
		set argStr to argStr & v & "
"
	end repeat
	return argStr
end joinWithCR

tell application "System Events" to tell process "Finder"
	activate
try
	choose color {{.DefaultColor | DefaultColorFunc}}
	set answer to "OK:" & my joinWithCR(result)
on error
	set answer to "CANCEL:"
end try
end tell
`
