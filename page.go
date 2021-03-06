package main

import "net/http"
import "html/template"

const pageContent string = `
<html>
  <head>
    <meta http-equiv="X-UA-Compatible" content="IE=Edge">
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="copyright" content="Z Carioca">
		<meta name="robots" content="index,follow">
		<meta name="author" content="Rafael Pacheco Chargel">
    <title>{{.Title}}</title>
    <link rel="stylesheet" type="text/css" href="main.css" media="screen"/>
    <!-- Google Analytics -->
		<script type="text/javascript">
		(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
		(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
		m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
		})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
		ga('create', 'UA-39324990-1', 'auto');
		ga('send', 'pageview');
		</script>
		<!-- End Google Analytics -->
  </head>
  <body>
    {{.Body}}
  </body>
</html>
`

var pageTemplate *template.Template

// Page a simple html page
type Page struct {
	Title string
	Body  template.HTML
}

// Render writes the content of the page using the template.
func (p *Page) Render(w http.ResponseWriter) {
	if pageTemplate == nil {
		pageTemplate, _ = template.New("PAGE").Parse(pageContent)
	}
	pageTemplate.Execute(w, p)
}
