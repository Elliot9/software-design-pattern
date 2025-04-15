package word

type WordService struct {
	finder      Finder
	repository  Repository
	lastQuiz    *Quiz
	currentQuiz *Quiz

	// for the testing
	randomizer Randomizer
}

func NewWordService(finder Finder, repository Repository, randomizer Randomizer) *WordService {
	return &WordService{
		finder:      finder,
		repository:  repository,
		lastQuiz:    nil,
		currentQuiz: nil,
		randomizer:  randomizer,
	}
}

func (s *WordService) Search(name string) (Word, bool) {
	return s.finder.Find(name)
}

func (s *WordService) AddWord(word Word) {
	s.repository.Add(word)
}

func (s *WordService) DeleteWord(word Word) bool {
	return s.repository.Delete(word)
}

func (s *WordService) GetAllWords() []Word {
	return s.repository.GetAll()
}

func (s *WordService) StartNewQuiz(number int) {
	words := s.GetAllWords()

	if len(words) < number {
		number = len(words)
	}

	indices := make([]int, number)
	for i := range indices {
		indices[i] = i
	}

	s.randomizer.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	selected := make([]Word, number)
	for i := 0; i < number; i++ {
		selected[i] = words[indices[i]]
	}

	s.currentQuiz = NewQuiz(selected, s.randomizer)
}

func (s *WordService) GetCurrentQuestion() (wordBlank, definition string) {
	return s.currentQuiz.GetCurrentQuestion()
}

func (s *WordService) CheckAnswer(input string) (bool, string) {
	result, answer := s.currentQuiz.CheckAnswer(input)

	// finish quiz
	if s.currentQuiz.GetTotalWords() == s.currentQuiz.GetCurrentIndex() {
		s.lastQuiz = s.currentQuiz
		s.currentQuiz = nil
	}

	return result, answer
}

func (s *WordService) GetCurrentQuiz() *Quiz {
	return s.currentQuiz
}

func (s *WordService) GetLastQuiz() *Quiz {
	return s.lastQuiz
}
