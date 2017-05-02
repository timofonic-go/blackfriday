package blackfriday

import (
  "bytes"
  "strings"
)

type Confluence struct {
}

func ConfluenceRenderer(flags int) Renderer {
  return &Confluence{}
}

// FIXME: Expected results are not returned
func (options *Confluence) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	// parse out the language names/classes
	count := 0
	for _, elt := range strings.Fields(lang) {
		if elt[0] == '.' {
			elt = elt[1:]
		}
		if len(elt) == 0 {
			continue
		}
		if count == 0 {
			out.WriteString("{code:language=")
		} else {
			out.WriteByte(' ')
		}
		attrEscape(out, []byte(elt))
		count++
	}

	if count == 0 {
		out.WriteString("{code}")
	} else {
		out.WriteString("}")
	}

	out.Write(text)
	out.WriteString("{code}\n")
}

func (options *Confluence) TitleBlock(out *bytes.Buffer, text []byte) {
}

func (options *Confluence) BlockQuote(out *bytes.Buffer, text []byte) {
	out.WriteString("{quote}")
	out.Write(text)
	out.WriteString("{quote}")
}

func (options *Confluence) BlockHtml(out *bytes.Buffer, text []byte) {
}

func (options *Confluence) Header(out *bytes.Buffer, text func() bool, level int, id string) {
  marker := out.Len()

	switch level {
	case 1:
		out.WriteString("h1. ")
	case 2:
		out.WriteString("h2. ")
	case 3:
		out.WriteString("h3. ")
	case 4:
		out.WriteString("h4. ")
	case 5:
		out.WriteString("h5. ")
	case 6:
		out.WriteString("h6. ")
	}
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString("\n\n")
}

func (options *Confluence) HRule(out *bytes.Buffer) {
	out.WriteString("----")
}

func (options *Confluence) List(out *bytes.Buffer, text func() bool, flags int) {
  marker := out.Len()
  if !text() {
		out.Truncate(marker)
		return
	}
  out.WriteString("\n")
}

func (options *Confluence) ListItem(out *bytes.Buffer, text []byte, flags int) {
  if flags&LIST_TYPE_ORDERED != 0 {
    out.WriteString("# ")
  } else {
    out.WriteString("* ")
  }
	out.Write(text)
  out.WriteString("\n")
}

func (options *Confluence) Paragraph(out *bytes.Buffer, text func() bool) {
  marker := out.Len()
  if !text() {
		out.Truncate(marker)
		return
	}
  out.WriteString("\n\n")
}

func (options *Confluence) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	out.Write(header)
	out.Write(body)
	out.WriteString("\n")
}

func (options *Confluence) TableRow(out *bytes.Buffer, text []byte) {

	out.Write(text)
	out.WriteString("\n")
}

func (options *Confluence) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {

	switch align {
	case TABLE_ALIGNMENT_LEFT:
		out.WriteString("||")
	case TABLE_ALIGNMENT_RIGHT:
		out.WriteString("||")
	case TABLE_ALIGNMENT_CENTER:
		out.WriteString("||")
	default:
		out.WriteString("||")
	}

	out.Write(text)
}

func (options *Confluence) TableCell(out *bytes.Buffer, text []byte, align int) {

	switch align {
	case TABLE_ALIGNMENT_LEFT:
		out.WriteString("|")
	case TABLE_ALIGNMENT_RIGHT:
		out.WriteString("|")
	case TABLE_ALIGNMENT_CENTER:
		out.WriteString("|")
	default:
		out.WriteString("|")
	}

	out.Write(text)
}

func (options *Confluence) Footnotes(out *bytes.Buffer, text func() bool) {
}

func (options *Confluence) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
}

func (options *Confluence) AutoLink(out *bytes.Buffer, link []byte, kind int) {
}

// TODO
func (options *Confluence) CodeSpan(out *bytes.Buffer, text []byte) {
}

func (options *Confluence) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	out.WriteString("*")
	out.Write(text)
	out.WriteString("*")
}

func (options *Confluence) Emphasis(out *bytes.Buffer, text []byte) {
	if len(text) == 0 {
		return
	}
	out.WriteString("_")
	out.Write(text)
	out.WriteString("_")
}

// TODO
func (options *Confluence) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
  out.WriteString("!")
	out.Write(link)
  out.WriteString("!")
}

func (options *Confluence) LineBreak(out *bytes.Buffer) {
	out.WriteString("\n")
}

func (options *Confluence) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.WriteString("[")
	out.Write(content)
	out.WriteString("|")
	out.Write(link)
	out.WriteString("]")
}

func (options *Confluence) RawHtmlTag(out *bytes.Buffer, text []byte) {
}

func (options *Confluence) TripleEmphasis(out *bytes.Buffer, text []byte) {
}

// FIXME: Expected results are not returned
func (options *Confluence) StrikeThrough(out *bytes.Buffer, text []byte) {
	out.WriteString("-")
	out.Write(text)
	out.WriteString("-")
}

func (options *Confluence) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
}

func (options *Confluence) Entity(out *bytes.Buffer, entity []byte) {
	out.Write(entity)
}

func (options *Confluence) NormalText(out *bytes.Buffer, text []byte) {
	escapeSpecialChars(out, text)
}

func (options *Confluence) DocumentHeader(out *bytes.Buffer) {
}

func (options *Confluence) DocumentFooter(out *bytes.Buffer) {
}

func (options *Confluence) GetFlags() int {
	return 0
}
