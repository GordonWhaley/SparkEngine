package main

import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
  "os"
  "io/ioutil"
  "log"
)

func EditorMain() {
  app := app.New()

	w := app.NewWindow("Spark Engine")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Spark Editor"),
    widget.NewButton("New Project", func() {
      textwindow := app.NewWindow("Set Project Name")
      input := widget.NewEntry()
      input.SetPlaceHolder("Name your new Project!")
      content := widget.NewVBox(input, widget.NewButton("Create", func() {
        CreateProject(input.Text)
      }))
      textwindow.SetContent(content)
      textwindow.Show()
      }),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

func CreateProject(projectName string) {
  os.Mkdir(projectName, 0777)
  projMainContents := []byte("package main\n\nfunc main() {\n\n}")
  err := ioutil.WriteFile(projectName + "/main.go", projMainContents, 0644)
  if err != nil {
    log.Fatal(err)
  }
}
