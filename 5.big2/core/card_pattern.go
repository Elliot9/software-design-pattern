package core

type CardPattern interface {
	SetNext(next CardPattern)
	Check(cards []Card, topPlay []Card) (bool, CardPattern)
	IsValid(cards []Card) bool
	IsStrongerThan(cards []Card, topPlay []Card) bool
	GetName() string
}

type BaseCardPattern struct {
	next CardPattern
	Name string
	CardPattern
}

func (b *BaseCardPattern) SetNext(next CardPattern) {
	b.next = next
}

func (b *BaseCardPattern) Check(cards []Card, topPlay []Card) (bool, CardPattern) {
	if b.CardPattern.IsValid(cards) {
		if len(topPlay) == 0 {
			return true, b.CardPattern
		}

		if b.CardPattern.IsValid(topPlay) && b.CardPattern.IsStrongerThan(cards, topPlay) {
			return true, b.CardPattern
		}
		return false, nil
	}

	if b.next == nil {
		return false, nil
	}

	return b.next.Check(cards, topPlay)
}

func (b *BaseCardPattern) IsValid(cards []Card) bool {
	panic("card pattern need to implement IsValid")
}

func (b *BaseCardPattern) IsStrongerThan(cards []Card, topPlay []Card) bool {
	panic("card pattern need to implement IsStrongerThan")
}

func (b *BaseCardPattern) GetName() string {
	return b.Name
}

var _ CardPattern = (*BaseCardPattern)(nil)
