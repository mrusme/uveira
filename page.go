package main

// Formatting of text within setences
type Formatting struct {
  Bold []string `json:"bold,omitempty" bson:"bold"`
  Italic []string `json:"italic,omitempty" bson:"italic"`
}

// Link within sentence
type Link struct {
  Text string `json:"text,omitempty" bson:"text"`
  Type string `json:"type,omitempty" bson:"type"`
  Page string `json:"page,omitempty" bson:"page"`
}

// Sentence inside a paragraph
type Sentence struct {
  Text string `json:"text,omitempty" bson:"text"`
  Links []Link `json:"links,omitempty" bson:"links"`
  Formatting Formatting `json:"formatting,omitempty" bson:"formatting"`
}

// Paragraph made of sentences
type Paragraph struct {
  Sentences []Sentence `json:"sentences,omitempty" bson:"sentences"`
}

// Section made of paragraphs
type Section struct {
  Title string `json:"title,omitempty" bson:"title"`
  Depth int `json:"depth,omitempty" bson:"depth"`
  Paragraphs []Paragraph `json:"paragraphs,omitempty" bson:"paragraphs"`
  // Infoboxes []Infobox `json:"infoboxes,omitempty" bson:"infoboxes"`
}

// Coordinate set of page (e.g. for cities, countries, etc.)
type Coordinate struct {
  Display string `json:"display,omitempty" bson:"display"`
  Template string `json:"template,omitempty" bson:"template"`
  Lat float64 `json:"lat,omitempty" bson:"lat"`
  Lon float64 `json:"lon,omitempty" bson:"lon"`
}

// Image inside page
type Image struct {
  File string `json:"file,omitempty" bson:"file"`
  Thumb string `json:"thumb,omitempty" bson:"thumb"`
  URL string `json:"url,omitempty" bson:"url"`
  Caption string `json:"caption,omitempty" bson:"caption"`
}

// Page structure
type Page struct {
  ID string `json:"_id,omitempty" bson:"_id"`
  Title string `json:"title,omitempty" bson:"title"`
  Categories []string `json:"categories,omitempty" bson:"categories"`
  Sections []Section `json:"sections,omitempty" bson:"sections"`
  Coordinates []Coordinate `json:"coordinates,omitempty" bson:"coordinates"`
  Images []Image `json:"images,omitempty" bson:"images"`
}

// RenderParagraph renders one paragraph
func (paragraph *Paragraph) RenderParagraph() (string) {
  var rendered string

  rendered = ""

  for _, sentence := range paragraph.Sentences {
    rendered += sentence.Text + " "
  }

  return rendered
}

// RenderSection renders one section
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

// RenderPage renders one page
func (page *Page) RenderPage(titleOnly bool) (string) {
  var rendered string

  if titleOnly == true {
    return "- " + page.Title + "\n"
  }

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

    rendered +=  renderedSectionTitle + renderedSection + "\n"
  }

  return rendered
}
