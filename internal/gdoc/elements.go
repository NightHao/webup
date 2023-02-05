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
	Body          Body                    `json:"body"`
	InlineObjects map[string]InlineObject `json:"inlineObjects"`
}

type Body struct {
	Content []StructuralElement `json:"content"`
}

type ImageProperties struct {
	ContentUri string `json:"contentUri"`
}

type UnitValue struct {
	Magnitude float64 `json:"magnitude"`
	Unit      string  `json:"unit"`
}

type InlineObject struct {
	InlineObjectProperties struct {
		EmbeddedObject struct {
			ImageProperties *ImageProperties `json:"imageProperties"`
			Size            struct {
				Height UnitValue `json:"height"`
				Width  UnitValue `json:"width"`
			} `json:"size"`
		}
	} `json:"inlineObjectProperties"`
}

type StructuralElement struct {
	Paragraph *Paragraph `json:"paragraph"`
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
	TextRun             *TextRun             `json:"textRun"`
	InlineObjectElement *InlineObjectElement `json:"inlineObjectElement"`
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

type InlineObjectElement struct {
	InlineObjectId string `json:"inlineObjectId"`
}

/*
How to handle bullets:
1. Each bullet point is a paragraph
2. Each sub-bullet point is also a paragraph
3. state machine approach: Iterate through SEs until a non-bullet is encountered
3. dict approach: Append result to dict of HTMLElements according to their IDs (gen pseudo ID for non-IDed element)
4. Lookup bullet styles from `lists`
*/
