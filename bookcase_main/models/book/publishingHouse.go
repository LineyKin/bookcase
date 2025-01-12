package book

type PublishingHouse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (ph PublishingHouse) IsEmpty() bool {
	if ph.Id == 0 && ph.Name == "" {
		return true
	}

	return false
}

func (ph PublishingHouse) IsNew() bool {
	if ph.Id == 0 && ph.Name != "" {
		return true
	}

	return false
}
