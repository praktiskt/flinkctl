package tools

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
)

func TablePrint(t interface{}) {
	tableprinter.Print(os.Stdout, t)
}

func JSONPrint(t interface{}) {
	b, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
