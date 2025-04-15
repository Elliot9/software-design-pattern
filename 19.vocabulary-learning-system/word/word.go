package word

type Word struct {
	Name        string
	Description string
	Definitions []Definition
}

type PoS string

const (
	Noun      PoS = "noun"
	Verb      PoS = "verb"
	Adjective PoS = "adjective"
	Adverb    PoS = "adverb"
)

func (p PoS) String() string {
	switch p {
	case Noun:
		return "noun"
	case Verb:
		return "verb"
	case Adjective:
		return "adj"
	case Adverb:
		return "adv"
	default:
		return string(p)
	}
}

type Definition struct {
	Pos         PoS
	Explanation string
}

func NewDefinition(pos PoS, explanation string) Definition {
	return Definition{
		Pos:         pos,
		Explanation: explanation,
	}
}

func NewWord(name string, description string, definitions []Definition) *Word {
	return &Word{
		Name:        name,
		Description: description,
		Definitions: definitions,
	}
}
