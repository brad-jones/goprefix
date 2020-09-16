package main

import (
	"fmt"

	"github.com/brad-jones/goprefix/v2/pkg/colorchooser"
	"github.com/logrusorgru/aurora/v3"
)

func main() {
	fmt.Println(colorchooser.Sprint("foo"))
	fmt.Println(colorchooser.Sprint("foo"))
	fmt.Println(colorchooser.Sprint("bar"))

	// a custom chooser that only has 2 colors and thus will re-use a color
	cc := colorchooser.New(aurora.RedFg, aurora.GreenFg)
	fmt.Println(cc.Sprint("foo"))
	fmt.Println(cc.Sprint("bar"))
	fmt.Println(cc.Sprint("baz"))
}
