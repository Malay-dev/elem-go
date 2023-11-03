package elem

import (
	"testing"

	"github.com/chasefleming/elem-go/attrs"
	"github.com/stretchr/testify/assert"
)

// ========== Document Structure ==========

func TestBody(t *testing.T) {
	expected := `<body class="page-body"><p>Welcome to Elem!</p></body>`
	el := Body(Attrs{attrs.Class: "page-body"}, P(nil, Text("Welcome to Elem!")))
	assert.Equal(t, expected, el.Render())
}

func TestHead(t *testing.T) {
	// ... [Code for the Head test here]
}

func TestHtml(t *testing.T) {
	expected := `<html lang="en"><head><meta charset="UTF-8"><title>Elem Page</title></head><body><p>Welcome to Elem!</p></body></html>`
	el := Html(Attrs{attrs.Lang: "en"},
		Head(nil,
			Meta(Attrs{attrs.Charset: "UTF-8"}),
			Title(nil, Text("Elem Page")),
		),
		Body(nil, P(nil, Text("Welcome to Elem!"))),
	)
	assert.Equal(t, expected, el.Render())
}

// ========== Text Formatting and Structure ==========

func TestA(t *testing.T) {
	expected := `<a href="https://example.com">Visit Example</a>`
	el := A(Attrs{attrs.Href: "https://example.com"}, Text("Visit Example"))
	assert.Equal(t, expected, el.Render())
}

func TestBlockquote(t *testing.T) {
	expected := `<blockquote>Quote text</blockquote>`
	el := Blockquote(nil, Text("Quote text"))
	assert.Equal(t, expected, el.Render())
}

func TestBr(t *testing.T) {
	expected := `<br>`
	el := Br(nil)
	assert.Equal(t, expected, el.Render())
}

func TestCode(t *testing.T) {
	expected := `<code>Code snippet</code>`
	el := Code(nil, Text("Code snippet"))
	assert.Equal(t, expected, el.Render())
}

func TestDiv(t *testing.T) {
	expected := `<div class="container">Hello, Elem!</div>`
	el := Div(Attrs{attrs.Class: "container"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestEm(t *testing.T) {
	expected := `<em>Italic text</em>`
	el := Em(nil, Text("Italic text"))
	assert.Equal(t, expected, el.Render())
}

func TestH1(t *testing.T) {
	expected := `<h1 class="title">Hello, Elem!</h1>`
	el := H1(Attrs{attrs.Class: "title"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH2(t *testing.T) {
	expected := `<h2 class="subtitle">Hello, Elem!</h2>`
	el := H2(Attrs{attrs.Class: "subtitle"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestH3(t *testing.T) {
	expected := `<h3>Hello, Elem!</h3>`
	el := H3(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestHr(t *testing.T) {
	expected := `<hr>`
	el := Hr(nil)
	assert.Equal(t, expected, el.Render())
}

func TestI(t *testing.T) {
	expected1 := `<i>Idiomatic Text</i>`
	expected2 := `<i class="fa-regular fa-face-smile"></i>`
	el := I(nil, Text("Idiomatic Text"))
	assert.Equal(t, expected1, el.Render())
	el = I(Attrs{attrs.Class: "fa-regular fa-face-smile"})
	assert.Equal(t, expected2, el.Render())
}

func TestP(t *testing.T) {
	expected := `<p>Hello, Elem!</p>`
	el := P(nil, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestPre(t *testing.T) {
	expected := `<pre>Preformatted text</pre>`
	el := Pre(nil, Text("Preformatted text"))
	assert.Equal(t, expected, el.Render())
}

func TestSpan(t *testing.T) {
	expected := `<span class="highlight">Hello, Elem!</span>`
	el := Span(Attrs{attrs.Class: "highlight"}, Text("Hello, Elem!"))
	assert.Equal(t, expected, el.Render())
}

func TestStrong(t *testing.T) {
	expected := `<strong>Bold text</strong>`
	el := Strong(nil, Text("Bold text"))
	assert.Equal(t, expected, el.Render())
}

// ========== Lists ==========

func TestLi(t *testing.T) {
	expected := `<li>Item 1</li>`
	el := Li(nil, Text("Item 1"))
	assert.Equal(t, expected, el.Render())
}

func TestUl(t *testing.T) {
	expected := `<ul><li>Item 1</li><li>Item 2</li></ul>`
	el := Ul(nil, Li(nil, Text("Item 1")), Li(nil, Text("Item 2")))
	assert.Equal(t, expected, el.Render())
}

func TestOl(t *testing.T) {
	expected := `<ol><li>Item 1</li><li>Item 2</li></ol>`
	el := Ol(nil, Li(nil, Text("Item 1")), Li(nil, Text("Item 2")))
	assert.Equal(t, expected, el.Render())
}

func TestDl(t *testing.T) {
	expected := `<dl><dt>Term 1</dt><dd>Description 1</dd><dt>Term 2</dt><dd>Description 2</dd></dl>`
	el := Dl(nil, Dt(nil, Text("Term 1")), Dd(nil, Text("Description 1")), Dt(nil, Text("Term 2")), Dd(nil, Text("Description 2")))
	assert.Equal(t, expected, el.Render())
}

func TestDt(t *testing.T) {
	expected := `<dt>Term 1</dt>`
	el := Dt(nil, Text("Term 1"))
	assert.Equal(t, expected, el.Render())
}

func TestDd(t *testing.T) {
	expected := `<dd>Description 1</dd>`
	el := Dd(nil, Text("Description 1"))
	assert.Equal(t, expected, el.Render())
}

// ========== Forms ==========

func TestButton(t *testing.T) {
	expected := `<button class="btn">Click Me</button>`
	el := Button(Attrs{attrs.Class: "btn"}, Text("Click Me"))
	assert.Equal(t, expected, el.Render())
}

func TestForm(t *testing.T) {
	expected := `<form action="/submit" method="post"><input name="username" type="text"></form>`
	el := Form(Attrs{attrs.Action: "/submit", attrs.Method: "post"}, Input(Attrs{attrs.Type: "text", attrs.Name: "username"}))
	assert.Equal(t, expected, el.Render())
}

func TestInput(t *testing.T) {
	expected := `<input name="username" placeholder="Enter your username" type="text">`
	el := Input(Attrs{attrs.Type: "text", attrs.Name: "username", attrs.Placeholder: "Enter your username"})
	assert.Equal(t, expected, el.Render())
}

func TestLabel(t *testing.T) {
	expected := `<label for="username">Username</label>`
	el := Label(Attrs{attrs.For: "username"}, Text("Username"))
	assert.Equal(t, expected, el.Render())
}

func TestSelectAndOption(t *testing.T) {
	expected := `<select name="color"><option value="red">Red</option><option value="blue">Blue</option></select>`
	el := Select(Attrs{attrs.Name: "color"}, Option(Attrs{attrs.Value: "red"}, Text("Red")), Option(Attrs{attrs.Value: "blue"}, Text("Blue")))
	assert.Equal(t, expected, el.Render())
}

func TestTextarea(t *testing.T) {
	expected := `<textarea name="comment" rows="5">Leave a comment...</textarea>`
	el := Textarea(Attrs{attrs.Name: "comment", attrs.Rows: "5"}, Text("Leave a comment..."))
	assert.Equal(t, expected, el.Render())
}

// ========== Boolean attributes ==========
func TestCheckedTrue(t *testing.T) {
	expected := `<input checked name="allow" type="checkbox">`
	el := Input(Attrs{attrs.Type: "checkbox", attrs.Name: "allow", attrs.Checked: "true"})
	assert.Equal(t, expected, el.Render())
}

func TestCheckedFalse(t *testing.T) {
	expected := `<input name="allow" type="checkbox">`
	el := Input(Attrs{attrs.Type: "checkbox", attrs.Name: "allow", attrs.Checked: "false"})
	assert.Equal(t, expected, el.Render())
}

func TestCheckedEmpty(t *testing.T) {
	expected := `<input name="allow" type="checkbox">`
	el := Input(Attrs{attrs.Type: "checkbox", attrs.Name: "allow", attrs.Checked: ""})
	assert.Equal(t, expected, el.Render())
}

// ========== Hyperlinks and Multimedia ==========

func TestImg(t *testing.T) {
	expected := `<img alt="An image" src="image.jpg">`
	el := Img(Attrs{attrs.Src: "image.jpg", attrs.Alt: "An image"})
	assert.Equal(t, expected, el.Render())
}

// ========== Meta Elements ==========

func TestLink(t *testing.T) {
	expected := `<link href="https://example.com/styles.css" rel="stylesheet">`
	el := Link(Attrs{attrs.Rel: "stylesheet", attrs.Href: "https://example.com/styles.css"})
	assert.Equal(t, expected, el.Render())
}

func TestMeta(t *testing.T) {
	expected := `<meta charset="UTF-8">`
	el := Meta(Attrs{attrs.Charset: "UTF-8"})
	assert.Equal(t, expected, el.Render())
}

func TestScript(t *testing.T) {
	expected := `<script src="https://example.com/script.js"></script>`
	el := Script(Attrs{attrs.Src: "https://example.com/script.js"})
	assert.Equal(t, expected, el.Render())
}

// ========== Semantic Elements ==========

// --- Semantic Sectioning Elements ---

func TestArticle(t *testing.T) {
	expected := `<article><h2>Article Title</h2><p>Article content.</p></article>`
	el := Article(nil, H2(nil, Text("Article Title")), P(nil, Text("Article content.")))
	assert.Equal(t, expected, el.Render())
}

func TestAside(t *testing.T) {
	expected := `<aside><p>Sidebar content.</p></aside>`
	el := Aside(nil, P(nil, Text("Sidebar content.")))
	assert.Equal(t, expected, el.Render())
}

func TestFooter(t *testing.T) {
	expected := `<footer><p>Footer content.</p></footer>`
	el := Footer(nil, P(nil, Text("Footer content.")))
	assert.Equal(t, expected, el.Render())
}

func TestHeader(t *testing.T) {
	expected := `<header class="site-header"><h1>Welcome to Elem!</h1></header>`
	el := Header(Attrs{attrs.Class: "site-header"}, H1(nil, Text("Welcome to Elem!")))
	assert.Equal(t, expected, el.Render())
}

func TestMainElem(t *testing.T) {
	expected := `<main><p>Main content goes here.</p></main>`
	el := Main(nil, P(nil, Text("Main content goes here.")))
	assert.Equal(t, expected, el.Render())
}

func TestNav(t *testing.T) {
	expected := `<nav><a href="/home">Home</a><a href="/about">About</a></nav>`
	el := Nav(nil, A(Attrs{attrs.Href: "/home"}, Text("Home")), A(Attrs{attrs.Href: "/about"}, Text("About")))
	assert.Equal(t, expected, el.Render())
}

func TestSection(t *testing.T) {
	expected := `<section><h3>Section Title</h3><p>Section content.</p></section>`
	el := Section(nil, H3(nil, Text("Section Title")), P(nil, Text("Section content.")))
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Form Elements ---

func TestFieldset(t *testing.T) {
	expected := `<fieldset class="custom-fieldset"><legend>Personal Information</legend><input name="name" type="text"></fieldset>`
	el := Fieldset(Attrs{attrs.Class: "custom-fieldset"}, Legend(nil, Text("Personal Information")), Input(Attrs{attrs.Type: "text", attrs.Name: "name"}))
	assert.Equal(t, expected, el.Render())
}

func TestLegend(t *testing.T) {
	expected := `<legend class="custom-legend">Legend Title</legend>`
	el := Legend(Attrs{attrs.Class: "custom-legend"}, Text("Legend Title"))
	assert.Equal(t, expected, el.Render())
}

func TestDatalist(t *testing.T) {
	expected := `<datalist id="exampleList"><option value="Option1">Option 1</option><option value="Option2">Option 2</option></datalist>`
	el := Datalist(Attrs{attrs.ID: "exampleList"}, Option(Attrs{attrs.Value: "Option1"}, Text("Option 1")), Option(Attrs{attrs.Value: "Option2"}, Text("Option 2")))
	assert.Equal(t, expected, el.Render())
}

func TestMeter(t *testing.T) {
	expected := `<meter max="100" min="0" value="50">50%</meter>`
	el := Meter(Attrs{attrs.Min: "0", attrs.Max: "100", attrs.Value: "50"}, Text("50%"))
	assert.Equal(t, expected, el.Render())
}

func TestOutput(t *testing.T) {
	expected := `<output for="inputId" name="result">Output</output>`
	el := Output(Attrs{attrs.For: "inputId", attrs.Name: "result"}, Text("Output"))
	assert.Equal(t, expected, el.Render())
}

func TestProgress(t *testing.T) {
	expected := `<progress max="100" value="60"></progress>`
	el := Progress(Attrs{attrs.Max: "100", attrs.Value: "60"})
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Interactive Elements ---

func TestDialog(t *testing.T) {
	expected := `<dialog open><p>This is an open dialog window</p></dialog>`
	el := Dialog(Attrs{attrs.Open: "true"}, P(nil, Text("This is an open dialog window")))
	assert.Equal(t, expected, el.Render())
}

func TestMenu(t *testing.T) {
	expected := `<menu><li>Item One</li><li>Item Two</li></menu>`
	el := Menu(nil, Li(nil, Text("Item One")), Li(nil, Text("Item Two")))
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Script Supporting Elements ---

func TestNoScript(t *testing.T) {
	expected := `<noscript><p>JavaScript is required for this application.</p></noscript>`
	el := NoScript(nil, P(nil, Text("JavaScript is required for this application.")))
	assert.Equal(t, expected, el.Render())
}

// --- Semantic Text Content Elements ---

func TestAddress(t *testing.T) {
	expected := `<address>123 Example St.</address>`
	el := Address(nil, Text("123 Example St."))
	assert.Equal(t, expected, el.Render())
}

func TestDetails(t *testing.T) {
	expected := `<details><summary>More Info</summary><p>Details content here.</p></details>`
	el := Details(nil, Summary(nil, Text("More Info")), P(nil, Text("Details content here.")))
	assert.Equal(t, expected, el.Render())
}

func TestDetailsWithOpenFalse(t *testing.T) {
	expected := `<details><summary>More Info</summary><p>Details content here.</p></details>`
	el := Details(Attrs{attrs.Open: "false"}, Summary(nil, Text("More Info")), P(nil, Text("Details content here.")))
	assert.Equal(t, expected, el.Render())
}

func TestDetailsWithOpenTrue(t *testing.T) {
	expected := `<details open><summary>More Info</summary><p>Details content here.</p></details>`
	el := Details(Attrs{attrs.Open: "true"}, Summary(nil, Text("More Info")), P(nil, Text("Details content here.")))
	assert.Equal(t, expected, el.Render())
}

func TestFigCaption(t *testing.T) {
	expected := `<figcaption>Description of the figure.</figcaption>`
	el := FigCaption(nil, Text("Description of the figure."))
	assert.Equal(t, expected, el.Render())
}

func TestFigure(t *testing.T) {
	expected := `<figure><img alt="An image" src="image.jpg"><figcaption>An image</figcaption></figure>`
	el := Figure(nil, Img(Attrs{attrs.Src: "image.jpg", attrs.Alt: "An image"}), FigCaption(nil, Text("An image")))
	assert.Equal(t, expected, el.Render())
}

func TestMark(t *testing.T) {
	expected := `<p>You must <mark>highlight</mark> this word.</p>`
	el := P(nil, Text("You must "), Mark(nil, Text("highlight")), Text(" this word."))
	assert.Equal(t, expected, el.Render())
}

func TestSummary(t *testing.T) {
	expected := `<details><summary>Summary Title</summary></details>`
	el := Details(nil, Summary(nil, Text("Summary Title")))
	assert.Equal(t, expected, el.Render())
}

func TestTime(t *testing.T) {
	expected := `<time datetime="2023-01-01T00:00:00Z">New Year's Day</time>`
	el := Time(Attrs{attrs.DateTime: "2023-01-01T00:00:00Z"}, Text("New Year's Day"))
	assert.Equal(t, expected, el.Render())
}

// ========== Tables ==========

func TestTr(t *testing.T) {
	expected := `<tr>Row content.</tr>`
	el := Tr(nil, Text("Row content."))
	assert.Equal(t, expected, el.Render())
}

func TestTd(t *testing.T) {
	expected := `<tr><td><h1>Cell one.</h1></td><td>Cell two.</td></tr>`
	el := Tr(nil, Td(nil, H1(nil, Text("Cell one."))), Td(nil, Text("Cell two.")))
	assert.Equal(t, expected, el.Render())
}

func TestTh(t *testing.T) {
	expected := `<tr><th>First name</th><th>Last name</th><th>Age</th></tr>`
	el := Tr(nil, Th(nil, Text("First name")), Th(nil, Text("Last name")), Th(nil, Text("Age")))
	assert.Equal(t, expected, el.Render())
}

func TestTHead(t *testing.T) {
	expected := `<thead><tr><td>Text</td><td><a href="/link">Link</a></td></tr></thead>`
	el := THead(nil, Tr(nil, Td(nil, Text("Text")), Td(nil, A(Attrs{attrs.Href: "/link"}, Text("Link")))))
	assert.Equal(t, expected, el.Render())
}

func TestTBody(t *testing.T) {
	expected := `<tbody><tr><td>Table body</td></tr></tbody>`
	el := TBody(nil, Tr(nil, Td(nil, Text("Table body"))))
	assert.Equal(t, expected, el.Render())
}

func TestTFoot(t *testing.T) {
	expected := `<tfoot><tr><td><a href="/footer">Table footer</a></td></tr></tfoot>`
	el := TFoot(nil, Tr(nil, Td(nil, A(Attrs{attrs.Href: "/footer"}, Text("Table footer")))))
	assert.Equal(t, expected, el.Render())
}

func TestTable(t *testing.T) {
	expected := `<table><tr><th>Table header</th></tr><tr><td>Table content</td></tr></table>`
	el := Table(nil, Tr(nil, Th(nil, Text("Table header"))), Tr(nil, Td(nil, Text("Table content"))))
	assert.Equal(t, expected, el.Render())
}
