package main

import (
	//"fmt"
	"io/ioutil"
)

func main() {
	s := (getTableHeader())

	s += formatRow(true, "123456789", "yourCompany/yourProduct", "1:23:45 1/2/34", "yourEmployee", "fixed a generic thing because it caused a generic issue and posted a generic error")

	s += (getTableFooter())

	writeFile("test.html", s)
}

func writeFile(name string, contents string) {
	_ = ioutil.WriteFile(name, []byte(contents), 0644)
}

//var colourOdd string =
//var colourEven string =

func getTableHeader() string {
	return `<style type="text/css">
	.tg  {border-collapse:collapse;border-spacing:0;}
	.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 20px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;}
	.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 20px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;}
	.tg .tg-li8k{font-family:"Lucida Sans Unicode", "Lucida Grande", sans-serif !important;;background-color:#1A1A1A;color:#ffffff}
	

	.tg .tg-yrsx{background-color:#313131;color:#ffffff}
	.tg .tg-uimw{background-color:#1A1A1A;color:#ffffff}


	.tg .tg-zh8g{background-color:#575757;color:#ffffff}
	.tg .tg-u7t1{background-color:#1A1A1A;color:#efefef}
	</style>
	<table class="tg">
	  <tr>
	    <th class="tg-u7t1">Change</th>
	    <th class="tg-uimw">Product</th>
	    <th class="tg-uimw">Time/Date</th>
	    <th class="tg-uimw">Developer</th>
	    <th class="tg-li8k">                              Description                             </th>
	  </tr>
	`

}

func formatRow(even bool, changeId string, product string, time string, developer string, description string) string {
	s := "	<tr>"

	class := ""
	if even {
		class = "tg-yrsx"
	} else {
		class = "tg-zh8g"
	}

	s += "<td class=\"" + class + "\">" + changeId + "</td>"
	s += "<td class=\"" + class + "\">" + product + "</td>"
	s += "<td class=\"" + class + "\">" + time + "</td>"
	s += "<td class=\"" + class + "\">" + developer + "</td>"
	s += "<td class=\"" + class + "\"; width=\"100%\">" + description + "</td>"

	s += "</tr>"
	return s
}

func getTableFooter() string {
	return "</table>"
}
