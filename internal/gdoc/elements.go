package gdoc

import "strings"

/*
TODOs:
1. Bullet
2. Table
3. Coloring
4. Image
5. Font size styling
6. URL
*/

type Document struct {
	Body Body `json:"body"`
}

type Body struct {
	Content []StructuralElement `json:"content"`
}

type StructuralElement struct {
	Type      string
	Paragraph Paragraph `json:"paragraph"`
}

type ParagraphStyle struct {
	NamedStyleType string `json:"namedStyleType"`
}

type Bullet struct {
	ListId string `json:"listId"`
	// TextStyle TextStyle `json:"textStyle"` // IDK what's the use case yet
}

type Paragraph struct {
	Type           string
	Elements       []ParagraphElement `json:"elements"`
	ParagraphStyle ParagraphStyle     `json:"paragraphStyle"`
	Bullet         Bullet             `json:"bullet"`
}

type ParagraphElement struct {
	Type    string
	TextRun TextRun `json:"textRun"`
}

type TextStyle struct {
	Bold      bool `json:"bold"`
	Italic    bool `json:"italic"`
	Underline bool `json:"underline"`
}

type TextRun struct {
	Content   string    `json:"content"`
	TextStyle TextStyle `json:"textStyle"`
}

/*
How to handle bullets:
1. Each bullet point is a paragraph
2. Each sub-bullet point is also a paragraph
3. state machine approach: Iterate through SEs until a non-bullet is encountered
3. dict approach: Append result to dict of HTMLElements according to their IDs (gen pseudo ID for non-IDed element)
4. Lookup bullet styles from `lists`

func (b *Bullet) exists() bool {
	return b.ListId != ""
}

*/

func (tr *TextRun) exists() bool {
	return len(tr.Content) > 0
}

func (tr *TextRun) toHTML() string {
	result := tr.Content
	result = strings.Replace(result, "\n", "<br />", -1)
	result = strings.Replace(result, "\u000b", "<br />", -1)
	if tr.TextStyle.Bold {
		result = "<b>" + result + "</b>"
	}
	if tr.TextStyle.Italic {
		result = "<i>" + result + "</i>"
	}
	if tr.TextStyle.Underline {
		result = "<u>" + result + "</u>"
	}
	return "<span>" + result + "</span>"
}

func (pe *ParagraphElement) setType() {
	if pe.TextRun.exists() {
		pe.Type = "TextRun"
	} else {
		pe.Type = "Unknown"
	}
}

func (pe *ParagraphElement) toHTML() string {
	switch pe.Type {
	case "TextRun":
		return pe.TextRun.toHTML()
	}
	return ""
}

func (paragraph *Paragraph) exists() bool {
	return len(paragraph.Elements) > 0
}

func (paragraph *Paragraph) toHTML() string {
	result := ""
	for _, elem := range paragraph.Elements {
		result += elem.toHTML()
	}
	switch paragraph.ParagraphStyle.NamedStyleType {
	case "HEADING_1":
		return "<h1>" + result + "</h1>"
	case "HEADING_2":
		return "<h2>" + result + "</h2>"
	case "HEADING_3":
		return "<h3>" + result + "</h3>"
	}
	return "<p>" + result + "</p>"
}

func (se *StructuralElement) setType() {
	if se.Paragraph.exists() {
		se.Type = "Paragraph"
		for i, _ := range se.Paragraph.Elements {
			se.Paragraph.Elements[i].setType()
		}
	} else {
		se.Type = "Unknown"
	}
}

func (se *StructuralElement) toHTML() string {
	switch se.Type {
	case "Paragraph":
		return se.Paragraph.toHTML()
	}
	return ""
}
