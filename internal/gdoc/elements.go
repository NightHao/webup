package gdoc

import "fmt"

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

type Paragraph struct {
	Elements []ParagraphElement `json:"elements"`
}

type ParagraphElement struct {
	Type    string
	TextRun TextRun `json:"textRun"`
}

type TextRun struct {
	Content string `json:"content"`
}

func (tr *TextRun) exists() bool {
	return len(tr.Content) > 0
}

func (pe *ParagraphElement) setType() {
	if pe.TextRun.exists() {
		pe.Type = "TextRun"
	} else {
		pe.Type = "Unknown"
	}
}

func (pe *ParagraphElement) print() {
	switch pe.Type {
	case "TextRun":
		fmt.Printf("TR: %s\n", pe.TextRun.Content)
		break
	}
}

func (paragraph *Paragraph) exists() bool {
	return len(paragraph.Elements) > 0
}

func (paragraph *Paragraph) print() {
	for _, elem := range paragraph.Elements {
		elem.print()
	}
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

func (se *StructuralElement) print() {
	switch se.Type {
	case "Paragraph":
		se.Paragraph.print()
		break
	}
}
