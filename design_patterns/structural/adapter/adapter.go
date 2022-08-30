package adapter

type Printer interface {
	Print(string) string
}

type MyPrinter struct{}

func (m MyPrinter) Print(input string) string {
	return "legacy printing: " + input
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	Printer
	Msg string
}

func (p PrinterAdapter) PrintStored() string {
	if p.Printer != nil {
		return p.Print(p.Msg)
	}
	return "modern printing: " + p.Msg
}
