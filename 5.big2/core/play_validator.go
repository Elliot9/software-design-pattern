package core

type PlayValidator interface {
	SetNext(next PlayValidator)
	Check(game Big2, cards []Card) error
	IsValid(game Big2, cards []Card) error
	RemoveValidator(validator PlayValidator) PlayValidator
	ShouldRemove(game Big2, cards []Card) bool
	GetNext() PlayValidator
}

type BasePlayValidator struct {
	next PlayValidator
	PlayValidator
}

func (b *BasePlayValidator) SetNext(next PlayValidator) {
	b.next = next
}

func (b *BasePlayValidator) Check(game Big2, cards []Card) error {
	if err := b.PlayValidator.IsValid(game, cards); err != nil {
		return err
	}

	if b.next != nil {
		return b.next.Check(game, cards)
	}

	return nil
}

func (b *BasePlayValidator) IsValid(game Big2, cards []Card) error {
	panic("play validator need to implement IsValid")
}

func (b *BasePlayValidator) RemoveValidator(validator PlayValidator) PlayValidator {
	if b.PlayValidator == validator {
		return b.PlayValidator.GetNext()
	}
	if b.next != nil {
		nextValidator := b.next.RemoveValidator(validator)
		b.next = nextValidator
	}

	return b.PlayValidator
}

func (b *BasePlayValidator) ShouldRemove(game Big2, cards []Card) bool {
	return false
}

func (b *BasePlayValidator) GetNext() PlayValidator {
	return b.next
}

var _ PlayValidator = (*BasePlayValidator)(nil)
