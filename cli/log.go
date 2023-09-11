package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
)

func logError(err error) {
	boldRed := color.New(color.FgRed, color.Bold)
	boldRed.Fprintf(os.Stderr, "\nerror: ")

	fmt.Fprintf(os.Stderr, "%s\n", color.RedString(err.Error()))
}

func logJSON(iList ...interface{}) {
	for _, i := range iList {
		m, err := json.Marshal(i)
		if err != nil {
			logError(err)
			
			return
		}

		pj, err := prettyjson.Format(m)
		if err != nil {
			logError(err)

			return
		}

		fmt.Fprintf(os.Stdout, "\n%s\n", string(pj))
	}
}
