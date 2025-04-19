// templates are part of
// text package and html package

package main

import "text/template"

func main() {
	tmpl := template.New("example")
	tmpl.Parse()
}