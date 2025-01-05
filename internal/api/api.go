package api

import (
	"fmt"
	"gowikicli/internal/utils"
)

// -------------
//    SEARCH
// -------------

func HandleSearch(query string) {
	fmt.Printf("%s", query)
}

// -------------
//     READ
// -------------

func HandleRead(flags utils.Flags, query string) {
	fmt.Printf("%s", query)
}
