package competence

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	Libelle string `json:"libelle"`
	Description  string `json:"description"`
	Categorie string `json:"categorie"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}