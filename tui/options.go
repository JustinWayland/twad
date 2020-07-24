package tui

import (
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/zmnpl/twad/cfg"
)

const (
	optsOkButtonLabel            = "Save"
	optsHeader                   = "Options"
	optsPathLabel                = "Base Path"
	optsPathDoesntExist          = " [red](doesn't exist)"
	optsDontDOOMWADDIR           = "Do NOT shadow DOOMWADDIR (use your shell's default)"
	optsWriteBasePathToEngineCFG = "Write the path into DOOM engines *.ini files"
	optsDontWarn                 = "Do NOT warn before deletion"
	optsSourcePortLabel          = "Source Ports"
	optsIwadsLabel               = "IWADs"
	optsNextTimeFirstStart       = "Show path selection dialog on next start"
	optsDefaultSaveDirs          = "Use default save dir"
	optsHideHeader               = "UI - Hide big DOOM logo"
	optsGamesListRelativeWitdh   = "UI - Game list relative width (1-100%)"
	optsDetailPaneVertical       = "UI - Split right side detail pane vertically"
)

func makeOptions() *tview.Flex {
	o := tview.NewForm()

	path := tview.NewInputField().SetLabel(optsPathLabel).SetLabelColor(tview.Styles.SecondaryTextColor).SetText(cfg.GetInstance().BasePath)
	o.AddFormItem(path)
	path.SetDoneFunc(func(key tcell.Key) {
		if _, err := os.Stat(path.GetText()); os.IsNotExist(err) {
			path.SetLabel(optsPathLabel + optsPathDoesntExist)
		} else {
			path.SetLabel(optsPathLabel)
		}
	})

	firstStart := tview.NewCheckbox().SetLabel(optsNextTimeFirstStart).SetLabelColor(tview.Styles.SecondaryTextColor).SetChecked(!cfg.GetInstance().Configured)
	o.AddFormItem(firstStart)

	sourcePorts := tview.NewInputField().SetLabel(optsSourcePortLabel).SetLabelColor(tview.Styles.SecondaryTextColor).SetText(strings.Join(cfg.GetInstance().SourcePorts, ","))
	o.AddFormItem(sourcePorts)

	iwads := tview.NewInputField().SetLabel(optsIwadsLabel).SetLabelColor(tview.Styles.SecondaryTextColor).SetText(strings.Join(cfg.GetInstance().IWADs, ","))
	o.AddFormItem(iwads)

	defaultSaveDirs := tview.NewCheckbox().SetLabel(optsDefaultSaveDirs).SetLabelColor(tview.Styles.SecondaryTextColor).SetChecked(cfg.GetInstance().DefaultSaveDir)
	o.AddFormItem(defaultSaveDirs)

	dontWarn := tview.NewCheckbox().SetLabel(optsDontWarn).SetLabelColor(tview.Styles.SecondaryTextColor).SetChecked(cfg.GetInstance().DeleteWithoutWarning)
	o.AddFormItem(dontWarn)

	printHeader := tview.NewCheckbox().SetLabel(optsHideHeader).SetLabelColor(tview.Styles.SecondaryTextColor).SetChecked(cfg.GetInstance().HideHeader)
	o.AddFormItem(printHeader)

	gameListRelWidth := tview.NewInputField().SetLabel(optsGamesListRelativeWitdh).SetLabelColor(tview.Styles.SecondaryTextColor).SetAcceptanceFunc(func(text string, char rune) bool {
		if text == "-" {
			return false
		}
		i, err := strconv.Atoi(text)
		return err == nil && i > 0 && i <= 100
	})
	gameListRelWidth.SetText(strconv.Itoa(cfg.GetInstance().GameListRelativeWidth))
	o.AddFormItem(gameListRelWidth)

	detailPaneVertical := tview.NewCheckbox().SetLabel(optsDetailPaneVertical).SetLabelColor(tview.Styles.SecondaryTextColor).SetChecked(cfg.GetInstance().DetailPaneSplitVertical)
	o.AddFormItem(detailPaneVertical)

	o.AddButton(optsOkButtonLabel, func() {
		c := cfg.GetInstance()

		c.BasePath = path.GetText()

		sps := strings.Split(sourcePorts.GetText(), ",")
		for i := range sps {
			sps[i] = strings.TrimSpace(sps[i])
		}
		c.SourcePorts = sps

		iwds := strings.Split(iwads.GetText(), ",")
		for i := range iwds {
			iwds[i] = strings.TrimSpace(iwds[i])
		}
		c.IWADs = iwds

		c.HideHeader = printHeader.IsChecked()
		c.DeleteWithoutWarning = dontWarn.IsChecked()
		c.DefaultSaveDir = defaultSaveDirs.IsChecked()
		c.GameListRelativeWidth, _ = strconv.Atoi(gameListRelWidth.GetText())
		c.DetailPaneSplitVertical = detailPaneVertical.IsChecked()
		c.Configured = !firstStart.IsChecked()

		cfg.Persist()
		appModeNormal()

	})

	// layout
	settingsPage := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(o, 90, 0, true).
		AddItem(tview.NewBox().SetBorder(false), 0, 1, false)

	return settingsPage
}
