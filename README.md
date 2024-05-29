### Format files

Usage: FormatFiles("./data/", "suite")

```
func FormatFiles(root, filename string) {
	files, _ := file.FindFiles(root, filename)
	for _, f := range files {
		ctxs := []ctx.Context{

			// removes lines contains key word
			// clears passed tests
			&actions.RemoveSubLine{
				RemoveLine: &actions.RemoveLine{
					Line: "passed",
				},
			},

			// clears skipped tests
			&actions.RemoveSubLine{
				RemoveLine: &actions.RemoveLine{
					Line: "skipped",
				},
			},

			// replaces csv default deliminer with a pipe token for convenience
			&actions.ReplaceWord{
				ReplaceLine: &actions.ReplaceLine{
					Old: "\",\"",
					New: "|",
				},
			},

			// covers csv edge cases with double double quotes in names
			&actions.ReplaceWord{
				ReplaceLine: &actions.ReplaceLine{
					Old: "\"\"",
					New: "'",
				},
			},

			// removes columns with redundant data
			&actions.RemoveColumn{
				Deliminer: "|",
				RemoveLine: &actions.RemoveLine{
					Line: "1,2,3,4,6,7,8,10",
				},
			},

			// replaces back pipe token to csv deliminer
			&actions.ReplaceWord{
				ReplaceLine: &actions.ReplaceLine{
					Old: "|",
					New: "\",\"",
				},
			},
			&actions.AddEndLine{
				Line: "\"",
			},
		}

		for _, ct := range ctxs {
			ctx.Exec(ct, f.FilePath)
		}

		ctx.Sort(f.FilePath)
	}
}
```

Use case example:
Allure csv report file of tests results
```
"Status","Start Time","Stop Time","Duration in ms","Parent Suite","Suite","Sub Suite","Test Class","Test Method","Name","Description"
...
"passed","Sat May 04 01:53:02 GMT 2024","Sat May 04 01:53:47 GMT 2024","44931","","353 Quality Output","","353 Quality Output","13 User sees relevant Output quality setting applied to the asset","13 User sees relevant Output quality setting applied to the asset",""
"passed","Sat May 04 02:27:22 GMT 2024","Sat May 04 02:27:22 GMT 2024","369","","097 Asset Picker - List attributes, Edit All","","097 Asset Picker - List attributes, Edit All","08 Author navigates to 'Quick Logic' tab for 'Asset Picker' properties panel","08 Author navigates to 'Quick Logic' tab for 'Asset Picker' properties panel",""
"passed","Sat May 04 00:57:55 GMT 2024","Sat May 04 00:57:57 GMT 2024","2001","","039 Manage Attributes - Copy, Edit","","039 Manage Attributes - Copy, Edit","09 User review fields on 'Edit Attribute' page","09 User review fields on 'Edit Attribute' page",""
"skipped","Sat May 04 02:40:49 GMT 2024","Sat May 04 02:40:50 GMT 2024","40","","236 Configure Template - Multi Select Question","","236 Configure Template - Multi Select Question","29 User views multi-select picker to appear as a standard asset picker","29 User views multi-select picker to appear as a standard asset picker",""
"passed","Sat May 04 01:35:01 GMT 2024","Sat May 04 01:35:13 GMT 2024","12175","","244 [Approval Enhancement] User details","","244 [Approval Enhancement] User details","03 User views Approval Chain on the Asset Information page without Approver","03 User views Approval Chain on the Asset Information page without Approver",""
"skipped","Sat May 04 02:39:10 GMT 2024","Sat May 04 02:39:10 GMT 2024","6","","208 Approver can edit template - Edit & Approve","","208 Approver can edit template - Edit & Approve","40 User returns to the Approval details page and sees the original ordered design if answer - no","40 User returns to the Approval details page and sees the original ordered design if answer - no",""
"skipped","Sat May 04 02:36:38 GMT 2024","Sat May 04 02:36:38 GMT 2024","10","","190 End user can edit text inline","","190 End user can edit text inline","08 User is not able to edit text inline if text element has 'Inline Edit' permission locked","08 User is not able to edit text inline if text element has 'Inline Edit' permission locked",""
"failed","Sat May 04 02:47:39 GMT 2024","Sat May 04 02:47:52 GMT 2024","13017","","341 Edit Template - Undo/redo - align, distribute, inline edit of text","","341 Edit Template - Undo/redo - align, distribute, inline edit of text","01 User can login to system via saml-idp","01 User can login to system via saml-idp",""
"passed","Sat May 04 01:06:13 GMT 2024","Sat May 04 01:06:13 GMT 2024","380","","089 Edit Template - Placeholders 1-2","","089 Edit Template - Placeholders 1-2","12 Author adds 'Asset Picker' question to the template","12 Author adds 'Asset Picker' question to the template",""
"passed","Sat May 04 01:08:14 GMT 2024","Sat May 04 01:08:14 GMT 2024","85","","095 Edit Template - Focus Question","","095 Edit Template - Focus Question","45 User views not a focus questions to be visible when no focus elements selected","45 User views not a focus questions to be visible when no focus elements selected",""
"broken","Sat May 04 01:54:44 GMT 2024","Sat May 04 01:54:46 GMT 2024","1941","","050 SSO - Create SSO User","","050 SSO - Create SSO User","04 User navigates to My Profile page","04 User navigates to My Profile page",""
...
```

Let's say we are interested only in failed or broken tests: "Status", their corresponding scope, i.e. "Test Method" and "Name". 
So, `format.FormatFiles(pathToFiles, fileNamePattern string)` with specified above contexes `[]ctx.Context` will result in:
```
"Status","Suite","Name"
...
"failed","341 Edit Template - Undo/redo - align, distribute, inline edit of text","01 User can login to system via saml-idp"
"broken","050 SSO - Create SSO User","04 User navigates to My Profile page"
...
```
