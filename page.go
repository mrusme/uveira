package main

type Formatting struct {
  Bold []string `json:"bold,omitempty" bson:"bold"`
  Italic []string `json:"italic,omitempty" bson:"italic"`
}

type Link struct {
  Text string `json:"text,omitempty" bson:"text"`
  Type string `json:"type,omitempty" bson:"type"`
  Page string `json:"page,omitempty" bson:"page"`
}

type Sentence struct {
  Text string `json:"text,omitempty" bson:"text"`
  Links []Link `json:"links,omitempty" bson:"links"`
  Formatting Formatting `json:"formatting,omitempty" bson:"formatting"`
}

type Paragraph struct {
  Sentences []Sentence `json:"sentences,omitempty" bson:"sentences"`
}

type Section struct {
  Title string `json:"title,omitempty" bson:"title"`
  Depth int `json:"depth,omitempty" bson:"depth"`
  Paragraphs []Paragraph `json:"paragraphs,omitempty" bson:"paragraphs"`
  // Infoboxes []Infobox `json:"infoboxes,omitempty" bson:"infoboxes"`
}

type Coordinate struct {
  Display string `json:"display,omitempty" bson:"display"`
  Template string `json:"template,omitempty" bson:"template"`
  Lat float64 `json:"lat,omitempty" bson:"lat"`
  Lon float64 `json:"lon,omitempty" bson:"lon"`
}

type Image struct {
  File string `json:"file,omitempty" bson:"file"`
  Thumb string `json:"thumb,omitempty" bson:"thumb"`
  URL string `json:"url,omitempty" bson:"url"`
  Caption string `json:"caption,omitempty" bson:"caption"`
}

type Page struct {
  Id string `json:"_id,omitempty" bson:"_id"`
  Title string `json:"title,omitempty" bson:"title"`
  Categories []string `json:"categories,omitempty" bson:"categories"`
  Sections []Section `json:"sections,omitempty" bson:"sections"`
  Coordinates []Coordinate `json:"coordinates,omitempty" bson:"coordinates"`
  Images []Image `json:"images,omitempty" bson:"images"`
}

func (paragraph *Paragraph) RenderParagraph() (string) {
  var rendered string

  rendered = ""

  for _, sentence := range paragraph.Sentences {
    rendered += sentence.Text + " "
  }

  return rendered
}

func (section *Section) RenderSection() (string) {
  var rendered string

  rendered = ""

  for _, paragraph := range section.Paragraphs {
    renderedParagraph := paragraph.RenderParagraph()
    if renderedParagraph == "" {
      continue
    }

    rendered += renderedParagraph + "\n\n"
  }

  return rendered
}

func (page *Page) RenderPage() (string) {
  var rendered string

  rendered = "# " + page.Title + "\n\n"

  for _, section := range page.Sections {
    renderedSection := section.RenderSection()
    if renderedSection == "" {
      continue
    }

    renderedSectionTitle := ""
    if section.Title != "" {
      renderedSectionTitle = "## " + section.Title + "\n\n"
    }

    rendered +=  renderedSectionTitle + renderedSection + "\n\n"
  }

  return rendered
}
