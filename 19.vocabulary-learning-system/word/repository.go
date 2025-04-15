package word

type Repository interface {
	Add(word Word)
	Delete(word Word) bool
	Get(name string) (Word, bool)
	GetAll() []Word
}
