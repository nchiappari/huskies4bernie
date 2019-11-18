package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Page name: ")
	pName, _ := reader.ReadString('\n')
	pName = strings.TrimSuffix(pName, "\n")

	fmt.Print("Redirect URL: ")
	rURL, _ := reader.ReadString('\n')
	rURL = strings.TrimSuffix(rURL, "\n")

	f, err := os.Create(pName + ".html")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<!-- Global site tag (gtag.js) - Google Analytics -->
			<script async src="https://www.googletagmanager.com/gtag/js?id=UA-101321756-2"></script>
			<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());

			gtag('config', 'UA-101321756-2');
			</script>

		</head>


		<script>
			let url = "` + rURL + `"
			window.location = url;
		</script>


		redirecting...

		<meta http-equiv="Refresh" content="0; url=` + rURL + `" />


		<a href="` + rURL + `">` + rURL + `</a>


		</html>
	`)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
