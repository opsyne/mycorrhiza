package markup

import (
	"fmt"
	"html"
	"strings"
)

// HyphaExists holds function that checks that a hypha is present.
var HyphaExists func(string) bool

// HyphaAccess holds function that accesses a hypha by its name.
var HyphaAccess func(string) (rawText, binaryHtml string, err error)

// GemLexerState is used by markup parser to remember what is going on.
type GemLexerState struct {
	// Name of hypha being parsed
	name  string
	where string // "", "list", "pre"
	// Line id
	id  int
	buf string
	// Temporaries
	img Img
}

type Line struct {
	id int
	// interface{} may be bad. What I need is a sum of string and Transclusion
	contents interface{}
}

func lex(name, content string) (ast []Line) {
	var state = GemLexerState{name: name}

	for _, line := range append(strings.Split(content, "\n"), "") {
		geminiLineToAST(line, &state, &ast)
	}
	return ast
}

// Lex `line` in markup and save it to `ast` using `state`.
func geminiLineToAST(line string, state *GemLexerState, ast *[]Line) {
	addLine := func(text interface{}) {
		*ast = append(*ast, Line{id: state.id, contents: text})
	}

	if "" == strings.TrimSpace(line) {
		if state.where == "list" {
			state.where = ""
			addLine(state.buf + "</ul>")
		} else if state.where == "number" {
			state.where = ""
			addLine(state.buf + "</ol>")
		}
		return
	}

	startsWith := func(token string) bool {
		return strings.HasPrefix(line, token)
	}

	// Beware! Usage of goto. Some may say it is considered evil but in this case it helped to make a better-structured code.
	switch state.where {
	case "img":
		goto imgState
	case "pre":
		goto preformattedState
	case "list":
		goto listState
	case "number":
		goto numberState
	default:
		goto normalState
	}

imgState:
	if shouldGoBackToNormal := state.img.Process(line); shouldGoBackToNormal {
		state.where = ""
		addLine(state.img)
	}
	return

preformattedState:
	switch {
	case startsWith("```"):
		state.where = ""
		state.buf = strings.TrimSuffix(state.buf, "\n")
		addLine(state.buf + "</code></pre>")
		state.buf = ""
	default:
		state.buf += html.EscapeString(line) + "\n"
	}
	return

listState:
	switch {
	case startsWith("* "):
		state.buf += fmt.Sprintf("\t<li>%s</li>\n", ParagraphToHtml(state.name, line[2:]))
	case startsWith("```"):
		state.where = "pre"
		addLine(state.buf + "</ul>")
		state.id++
		state.buf = fmt.Sprintf("<pre id='%d' alt='%s' class='codeblock'><code>", state.id, strings.TrimPrefix(line, "```"))
	default:
		state.where = ""
		addLine(state.buf + "</ul>")
		goto normalState
	}
	return

numberState:
	switch {
	case startsWith("*. "):
		state.buf += fmt.Sprintf("\t<li>%s</li>\n", ParagraphToHtml(state.name, line[3:]))
	case startsWith("```"):
		state.where = "pre"
		addLine(state.buf + "</ol>")
		state.id++
		state.buf = fmt.Sprintf("<pre id='%d' alt='%s' class='codeblock'><code>", state.id, strings.TrimPrefix(line, "```"))
	default:
		state.where = ""
		addLine(state.buf + "</ol>")
		goto normalState
	}
	return

normalState:
	state.id++
	switch {

	case startsWith("```"):
		state.where = "pre"
		state.buf = fmt.Sprintf("<pre id='%d' alt='%s' class='codeblock'><code>", state.id, strings.TrimPrefix(line, "```"))
	case startsWith("* "):
		state.where = "list"
		state.buf = fmt.Sprintf("<ul id='%d'>\n", state.id)
		goto listState
	case startsWith("*. "):
		state.where = "number"
		state.buf = fmt.Sprintf("<ol id='%d'>\n", state.id)
		goto numberState

	case startsWith("###### "):
		addLine(fmt.Sprintf(
			"<h6 id='%d'>%s</h6>", state.id, line[7:]))
	case startsWith("##### "):
		addLine(fmt.Sprintf(
			"<h5 id='%d'>%s</h5>", state.id, line[6:]))
	case startsWith("#### "):
		addLine(fmt.Sprintf(
			"<h4 id='%d'>%s</h4>", state.id, line[5:]))
	case startsWith("### "):
		addLine(fmt.Sprintf(
			"<h3 id='%d'>%s</h3>", state.id, line[4:]))
	case startsWith("## "):
		addLine(fmt.Sprintf(
			"<h2 id='%d'>%s</h2>", state.id, line[3:]))
	case startsWith("# "):
		addLine(fmt.Sprintf(
			"<h1 id='%d'>%s</h1>", state.id, line[2:]))

	case startsWith(">"):
		addLine(fmt.Sprintf(
			"<blockquote id='%d'>%s</blockquote>", state.id, remover(">")(line)))
	case startsWith("=>"):
		href, text, class := Rocketlink(line, state.name)
		addLine(fmt.Sprintf(
			`<p><a id='%d' class='rocketlink %s' href="%s">%s</a></p>`, state.id, class, href, text))

	case startsWith("<="):
		addLine(parseTransclusion(line, state.name))
	case line == "----":
		*ast = append(*ast, Line{id: -1, contents: "<hr/>"})
	case MatchesImg(line):
		state.where = "img"
		state.img = ImgFromFirstLine(line, state.name)
	default:
		addLine(fmt.Sprintf("<p id='%d'>%s</p>", state.id, ParagraphToHtml(state.name, line)))
	}
}