package view

var Shared = `<!DOCTYPE html>
<html lang="en">
	<head>
		<title>golgi</title>
		<script src="https://unpkg.com/htmx.org@1.9.4" integrity="sha384-zUfuhFKKZCbHTY6aRR46gxiqszMk5tcHjsVFxnUo8VMus4kHGVdIYVbOYYNlKmHV" crossorigin="anonymous"></script>
		<link rel="stylesheet" href="/css/layout">
	</head>
<body>
	<div id="container">
		<div hx-get="/main" hx-target="#container" hx-trigger="load" > </div> 
	</div>
	</br>
	<div hx-get="/time" hx-trigger="every 1s" hx-swap="innerHTML" id="time" > </div>
</body>
</html>`

var Main = `
		<div id="main">
			<h1>Welcome to golgi</h1>
			<p>Here we have the main page.<p>

			<button hx-trigger="click"
    		hx-get="/secondPage"
			hx-target="#container"
    		hx-swap="innerHTML"
    		class="count">Click here</button>
		</div>
`
var SecondPage = `
		<div id="main">
			<h1>Welcome to golgi</h1>
			<p>This is the second page<p>

			<button hx-trigger="click"
    		hx-get="/main"
			hx-target="#container"
    		hx-swap="innerHTML"
    		class="count">Click on this one</button>
		</div>
`
