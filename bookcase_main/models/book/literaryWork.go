package book

type LiteraryWork struct {
	Id   int
	Name string
}

func (lw LiteraryWork) IsEmpty() bool {
	if lw.Id == 0 && lw.Name == "" {
		return true
	}

	return false
}

func (lw LiteraryWork) IsNew() bool {
	if lw.Id == 0 && lw.Name != "" {
		return true
	}

	return false
}
