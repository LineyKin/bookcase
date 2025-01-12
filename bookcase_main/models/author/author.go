package author

type Author struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FatherName string `json:"fatherName,omitempty"`
	LastName   string `json:"lastName"`
}

func (a Author) GetName() string {
	if a.LastName == "" && a.FatherName == "" {
		return a.Name
	}

	if a.FatherName == "" {
		return a.Name + " " + a.LastName
	}

	return a.Name + " " + a.FatherName + " " + a.LastName
}
