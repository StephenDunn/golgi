package view

var Index = `<!DOCTYPE html>
<html lang="en">
	<head>
		<title>golgi</title>
		<script src="https://unpkg.com/htmx.org@1.9.4" integrity="sha384-zUfuhFKKZCbHTY6aRR46gxiqszMk5tcHjsVFxnUo8VMus4kHGVdIYVbOYYNlKmHV" crossorigin="anonymous"></script>
	</head>
<body>

	<h1>My First Heading</h1>
	<p>My first paragraph.</p>

	<div hx-trigger="click"
    hx-get="/dothing/whatwhat"
    hx-swap="outerHTML"
    class="count">Click here</div>

</body>
</html>`
