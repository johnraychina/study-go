package main

import (
	"fmt"
	"html/template"
	"os"
)

type TaobaoShop struct {
	Name      string
	ItemPrice map[string]int
}

func main() {
	itemPrices := make(map[string]int, 0)
	itemPrices["Tent"] = 123
	shop := TaobaoShop{Name: "CampingShop", ItemPrice: itemPrices}
	tmpl, err := template.New("test").Parse("{{.Name}}:  The price of Tent is {{.ItemPrice.Tent}} ")
	if err != nil {
		fmt.Errorf("error:%v", err)
	}
	err = tmpl.Execute(os.Stdout, shop)
	if err != nil {
		fmt.Errorf("error:%v", err)
	}
}
