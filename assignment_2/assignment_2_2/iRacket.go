package assignment_2_2

type IRacket interface {
	setLogo(logo string)
	getLogo() string
	setMaterial(material string)
	getMaterial() string
}

type Racket struct {
	logo     string
	material string
}

func (r *Racket) setLogo(logo string) {
	r.logo = logo
}
func (r *Racket) getLogo() string {
	return r.logo
}

func (r *Racket) setMaterial(material string) {
	r.material = material
}

func (r *Racket) getMaterial() string {
	return r.material
}
