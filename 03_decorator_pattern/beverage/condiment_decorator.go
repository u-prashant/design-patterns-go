package beverage

type Mocha struct {
	Beverage Beverage
}

func (m *Mocha) GetSize() Size {
	return m.Beverage.GetSize()
}

func (m *Mocha) SetSize(size Size) {
	m.Beverage.SetSize(size)
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", Mochas"
}

func (m *Mocha) Cost() float32 {
	switch m.Beverage.GetSize() {
	case SMALL:
		return m.Beverage.Cost() + .20
	case MEDIUM:
		return m.Beverage.Cost() + .30
	case LARGE:
		return m.Beverage.Cost() + .40
	default:
		return m.Beverage.Cost() + .40
	}
}

func NewMocha(b Beverage) Beverage {
	return &Mocha{
		Beverage: b,
	}
}

type Whip struct {
	Beverage Beverage
}

func (m *Whip) GetSize() Size {
	return m.Beverage.GetSize()
}

func (m *Whip) SetSize(size Size) {
	m.Beverage.SetSize(size)
}

func (m *Whip) GetDescription() string {
	return m.Beverage.GetDescription() + ", Whip"
}

func (m *Whip) Cost() float32 {
	switch m.Beverage.GetSize() {
	case SMALL:
		return m.Beverage.Cost() + .15
	case MEDIUM:
		return m.Beverage.Cost() + .30
	case LARGE:
		return m.Beverage.Cost() + .45
	default:
		return m.Beverage.Cost() + .45
	}
}

func NewWhip(b Beverage) Beverage {
	return &Whip{
		Beverage: b,
	}
}
