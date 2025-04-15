package word

type Finder interface {
	Find(name string) (Word, bool)
}
