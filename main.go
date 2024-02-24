package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

var selectedFolder string

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("STL to DRC 转换器")
	w.Resize(fyne.NewSize(700, 500))

	input := widget.NewEntry()
	input.SetText("18")

	folderLabel := widget.NewLabel("")
	outputLabel := widget.NewLabel("")

	open := widget.NewButton("Select folder...", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				selectedFolder = uri.Path()
				folderLabel.SetText("Folder selected: " + selectedFolder)
			}
		}, w)
	})

	convert := widget.NewButton("Start", func() {
		if selectedFolder != "" {
			go convert(selectedFolder, input.Text, w, outputLabel)
		} else {
			dialog.ShowInformation("Info", "Please select a folder first.", w)
		}
	})

	scrollOutput := container.NewScroll(outputLabel)
	scrollOutput.SetMinSize(fyne.NewSize(700, 300))
	w.SetContent(container.NewVBox(
		open,
		folderLabel,
		input,
		convert,
		scrollOutput,
	))

	w.ShowAndRun()
}
func convert(folder string, params string, w fyne.Window, outputLabel *widget.Label) {
	outputLabel.SetText("")
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".stl" || filepath.Ext(path) == ".STL" {
			drcPath := filepath.Join(folder, "drc")
			os.MkdirAll(drcPath, os.ModePerm)

			baseName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			output := filepath.Join(drcPath, baseName+".drc")
			cmdName := "./draco_encoder"
			if runtime.GOOS == "windows" {
				cmdName = "draco_encoder.exe"
			}

			cmd := exec.Command(cmdName, "-qp", params, "-i", path, "-o", output)

			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				return err
			}

			outputLabel.SetText(outputLabel.Text + "\nCommand: " + strings.Join(cmd.Args, " ") + "\nOutput: " + out.String())
		}

		return nil
	})

	if err != nil {
		dialog.ShowError(err, w)
	} else {
		dialog.ShowInformation("Info", "Conversion completed.", w)
	}
}
