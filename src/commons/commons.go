package commons

import (
	"fmt"
	"qdb/src/kv"
)

const DATA_BASE_PATH = "./data/"

func PrintPrompt() {
	fmt.Printf("%v> ", kv.STORAGE_FILE_PREFIX)
}
