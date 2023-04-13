package models

type Repo struct {
	sku []Sku
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(sku Sku) {
	r.sku = append(r.sku, sku)
}

func (r *Repo) GetAll() []Sku {
	return r.sku
}
