package main

import (
	"fmt"
	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
)

const (
	FONT_MY    = "微软雅黑"
	FONT_MD    = "MPBOLD"
)

var (
	largeFont = core.Font{Family: FONT_MY, Size: 15}
	headFont  = core.Font{Family: FONT_MY, Size: 12}
	textFont  = core.Font{Family: FONT_MY, Size: 10}
)

func TouchPdf() {
	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: FONT_MY,
		FileName: "./ttf/microsoft.ttf",
	}

	r.SetFonts([]*core.FontMap{&font1})
	r.SetPage("A4", "P")
	r.FisrtPageNeedHeader = true
	r.FisrtPageNeedFooter = true


	//r.RegisterExecutor(ComplexReportHeaderExecutor, core.Header)
	r.RegisterExecutor(ComplexReportExecutor, core.Detail)
	r.RegisterExecutor(ComplexReportFooterExecutor, core.Footer)

	r.Execute("./complex_report_test.pdf")
	//r.SaveAtomicCellText("./complex_report_test.txt")
}

func ComplexReportExecutor(report *core.Report) {
	lineSpace := 1.0
	lineHight := 16.0

	//分割线
	line := gopdf.NewHLine(report).SetMargin(core.Scope{Top: 3.0, Bottom: 6.8}).SetWidth(0.1)

	div := gopdf.NewDivWithWidth(1000, lineHight, lineSpace, report)

	div.SetFont(headFont).SetContent("问诊时间:  "+"2021-01-02 15:04:05").GenerateAtomicCell()
	report.SetMargin(0, 5)


	div.SetContent("病人症状").GenerateAtomicCell()
	report.SetMargin(0, 5)
	gopdf.NewDivWithWidth(415, lineHight, lineSpace, report).SetMarign(core.Scope{Left: 10}).
		SetFont(textFont).SetContent("你好").GenerateAtomicCell()
	report.SetMargin(0, 3)
	line.GenerateAtomicCell()

	div.SetContent("病情诊断").GenerateAtomicCell()
	report.SetMargin(0, 5)
	gopdf.NewDivWithWidth(415, lineHight, lineSpace, report).SetMarign(core.Scope{Left: 10}).
		SetFont(textFont).SetContent("你好").GenerateAtomicCell()
	report.SetMargin(0, 3)
	line.GenerateAtomicCell()

	div.SetContent("治疗建议").GenerateAtomicCell()
	report.SetMargin(0, 5)
	gopdf.NewDivWithWidth(415, lineHight, lineSpace, report).SetMarign(core.Scope{Left: 10}).
		SetFont(textFont).SetContent("你好").GenerateAtomicCell()
	report.SetMargin(0, 3)
	line.GenerateAtomicCell()

	return
}
func ComplexReportFooterExecutor(report *core.Report) {
	content := fmt.Sprintf("第 %v / {#TotalPage#} 页", report.GetCurrentPageNo())
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(textFont)
	footer.SetFontColor("60, 179, 113")
	footer.HorizontalCentered().SetContent(content).GenerateAtomicCell()
}
func ComplexReportHeaderExecutor(report *core.Report) {
	content := "问诊记录"
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(textFont)
	footer.SetFontColor("255,0,0")
	footer.SetBorder(core.Scope{Top: 10})
	footer.HorizontalCentered().SetContent(content).GenerateAtomicCell()
}
