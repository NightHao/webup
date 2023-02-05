package gdoc

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

func (paragraph *Paragraph) exists() bool {
	return len(paragraph.Elements) > 0
}
