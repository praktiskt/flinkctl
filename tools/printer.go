package tools

import (
	"os"

	"github.com/landoop/tableprinter"
)

func TablePrint(t interface{}) {
	tableprinter.Print(os.Stdout, t)
}
