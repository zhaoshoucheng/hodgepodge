package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/i18n/gi18n"
)

var cities = []string{
	"City_SuZhou",
	"City_Peking",
}
func main() {
	t := gi18n.New()
	err := t.SetPath("/Users/didi/go/src/hodgepodge/i8n")
	if err != nil {
		fmt.Println(err)
		return
	}
	language := "en"
	t.SetLanguage(language)
	for _, city := range cities {
		fmt.Println(t.Translate(context.Background(), city))
	}
}
