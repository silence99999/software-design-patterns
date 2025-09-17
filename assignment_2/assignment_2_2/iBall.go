package assignment_2_2

type IBall interface {
	setLogo(logo string)
	getLogo() string
	setMaterial(material string)
	getMaterial() string
}

type Ball struct {
	logo     string
	material string
}

func (b *Ball) setLogo(logo string) {
	b.logo = logo
}

func (b *Ball) getLogo() string {
	return b.logo
}

func (b *Ball) setMaterial(material string) {
	b.material = material
}

func (b *Ball) getMaterial() string {
	return b.material
}
