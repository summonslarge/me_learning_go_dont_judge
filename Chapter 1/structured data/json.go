package main

import (
	"encoding/xml"
	"fmt"
)

type Thing struct {
	Belt string `xml:"id,attr"`
	Door string `xml:"parent>child"`
}

func main() {
	f := Thing{"Joe Junior", "Roshambo"}
	b, _ := xml.Marshal(f)
	fmt.Println(string(b))
	xml.Unmarshal(b, &f)
}
