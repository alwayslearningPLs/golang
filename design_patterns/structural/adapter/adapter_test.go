package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	var input = "hello"

	t.Run("adapter with old implementation", func(t *testing.T) {
		printer := PrinterAdapter{Printer: MyPrinter{}, Msg: input}
		want := "legacy printing: " + input

		assert.Equal(t, want, printer.Print(input))
		assert.Equal(t, want, printer.PrintStored())
	})

	t.Run("adapter with new implementation", func(t *testing.T) {
		printer := PrinterAdapter{Msg: input}
		want := "modern printing: " + input

		assert.Equal(t, want, printer.PrintStored())
	})
}
