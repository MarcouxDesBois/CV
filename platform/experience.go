package experience

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	Libelle string `json:"libelle"`
	AnneeDebut  string `json:"anneeDebut"`
	AnneeFin  string `json:"anneeFin"`
	Entreprise string `json:"entreprise"`
	Poste string `json:"poste"`
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