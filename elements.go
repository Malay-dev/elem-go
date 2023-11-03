package elem

// ========== Document Structure ==========

func Body(props Attrs, children ...Node) *Element {
	return NewElement("body", props, children...)
}

func Head(props Attrs, children ...Node) *Element {
	return NewElement("head", props, children...)
}

func Html(props Attrs, children ...Node) *Element {
	return NewElement("html", props, children...)
}

func Title(props Attrs, children ...Node) *Element {
	return NewElement("title", props, children...)
}

// ========== Text Formatting and Structure ==========

func A(props Attrs, children ...Node) *Element {
	return NewElement("a", props, children...)
}

func Br(props Attrs) *Element {
	return NewElement("br", props)
}

func Blockquote(props Attrs, children ...Node) *Element {
	return NewElement("blockquote", props, children...)
}

func Code(props Attrs, children ...Node) *Element {
	return NewElement("code", props, children...)
}

func Div(props Attrs, children ...Node) *Element {
	return NewElement("div", props, children...)
}

func Em(props Attrs, children ...Node) *Element {
	return NewElement("em", props, children...)
}

func H1(props Attrs, children ...Node) *Element {
	return NewElement("h1", props, children...)
}

func H2(props Attrs, children ...Node) *Element {
	return NewElement("h2", props, children...)
}

func H3(props Attrs, children ...Node) *Element {
	return NewElement("h3", props, children...)
}

func Hr(props Attrs) *Element {
	return NewElement("hr", props)
}

func I(props Attrs, children ...Node) *Element {
	return NewElement("i", props, children...)
}

func P(props Attrs, children ...Node) *Element {
	return NewElement("p", props, children...)
}

func Pre(props Attrs, children ...Node) *Element {
	return NewElement("pre", props, children...)
}

func Span(props Attrs, children ...Node) *Element {
	return NewElement("span", props, children...)
}

func Strong(props Attrs, children ...Node) *Element {
	return NewElement("strong", props, children...)
}

func Text(content string) TextNode {
	return TextNode(content)
}

// ========== Lists ==========

func Li(props Attrs, children ...Node) *Element {
	return NewElement("li", props, children...)
}

func Ul(props Attrs, children ...Node) *Element {
	return NewElement("ul", props, children...)
}

func Ol(props Attrs, children ...Node) *Element {
	return NewElement("ol", props, children...)
}

func Dl(props Attrs, children ...Node) *Element {
	return NewElement("dl", props, children...)
}

func Dt(props Attrs, children ...Node) *Element {
	return NewElement("dt", props, children...)
}

func Dd(props Attrs, children ...Node) *Element {
	return NewElement("dd", props, children...)
}

// ========== Forms ==========

func Button(props Attrs, children ...Node) *Element {
	return NewElement("button", props, children...)
}

func Form(attrs Attrs, children ...Node) *Element {
	return NewElement("form", attrs, children...)
}

func Input(attrs Attrs) *Element {
	return NewElement("input", attrs)
}

func Label(attrs Attrs, children ...Node) *Element {
	return NewElement("label", attrs, children...)
}

func Option(attrs Attrs, content TextNode) *Element {
	return NewElement("option", attrs, content)
}

func Select(attrs Attrs, children ...Node) *Element {
	return NewElement("select", attrs, children...)
}

func Textarea(attrs Attrs, content TextNode) *Element {
	return NewElement("textarea", attrs, content)
}

// ========== Hyperlinks and Multimedia ==========

func Img(props Attrs) *Element {
	return NewElement("img", props)
}

// ========== Meta Elements ==========

func Link(props Attrs) *Element {
	return NewElement("link", props)
}

func Meta(props Attrs) *Element {
	return NewElement("meta", props)
}

func Script(props Attrs, children ...Node) *Element {
	return NewElement("script", props, children...)
}

// ========== Semantic Elements ==========

// --- Semantic Sectioning Elements ---

func Article(props Attrs, children ...Node) *Element {
	return NewElement("article", props, children...)
}

func Aside(props Attrs, children ...Node) *Element {
	return NewElement("aside", props, children...)
}

func Footer(props Attrs, children ...Node) *Element {
	return NewElement("footer", props, children...)
}

func Header(props Attrs, children ...Node) *Element {
	return NewElement("header", props, children...)
}

func Main(props Attrs, children ...Node) *Element {
	return NewElement("main", props, children...)
}

func Nav(props Attrs, children ...Node) *Element {
	return NewElement("nav", props, children...)
}

func Section(props Attrs, children ...Node) *Element {
	return NewElement("section", props, children...)
}

// ========== Semantic Form Elements ==========

func Fieldset(props Attrs, children ...Node) *Element {
	return NewElement("fieldset", props, children...)
}

func Legend(props Attrs, children ...Node) *Element {
	return NewElement("legend", props, children...)
}

func Datalist(props Attrs, children ...Node) *Element {
	return NewElement("datalist", props, children...)
}

func Meter(props Attrs, children ...Node) *Element {
	return NewElement("meter", props, children...)
}

func Output(props Attrs, children ...Node) *Element {
	return NewElement("output", props, children...)
}

func Progress(props Attrs, children ...Node) *Element {
	return NewElement("progress", props, children...)
}

// --- Semantic Interactive Elements ---

func Dialog(props Attrs, children ...Node) *Element {
	return NewElement("dialog", props, children...)
}

func Menu(props Attrs, children ...Node) *Element {
	return NewElement("menu", props, children...)
}

// --- Semantic Script Supporting Elements ---

func NoScript(props Attrs, children ...Node) *Element {
	return NewElement("noscript", props, children...)
}

// --- Semantic Text Content Elements ---

func Address(props Attrs, children ...Node) *Element {
	return NewElement("address", props, children...)
}

func Details(props Attrs, children ...Node) *Element {
	return NewElement("details", props, children...)
}

func FigCaption(props Attrs, children ...Node) *Element {
	return NewElement("figcaption", props, children...)
}

func Figure(props Attrs, children ...Node) *Element {
	return NewElement("figure", props, children...)
}

func Mark(props Attrs, children ...Node) *Element {
	return NewElement("mark", props, children...)
}

func Summary(props Attrs, children ...Node) *Element {
	return NewElement("summary", props, children...)
}

func Time(props Attrs, children ...Node) *Element {
	return NewElement("time", props, children...)
}

// ========== Tables ==========

func Table(props Attrs, children ...Node) *Element {
	return NewElement("table", props, children...)
}

func THead(props Attrs, children ...Node) *Element {
	return NewElement("thead", props, children...)
}

func TBody(props Attrs, children ...Node) *Element {
	return NewElement("tbody", props, children...)
}

func TFoot(props Attrs, children ...Node) *Element {
	return NewElement("tfoot", props, children...)
}

func Tr(props Attrs, children ...Node) *Element {
	return NewElement("tr", props, children...)
}

func Th(props Attrs, children ...Node) *Element {
	return NewElement("th", props, children...)
}

func Td(props Attrs, children ...Node) *Element {
	return NewElement("td", props, children...)
}
