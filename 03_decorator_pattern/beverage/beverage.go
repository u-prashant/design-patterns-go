package beverage

type Size int

const (
	SMALL Size = iota + 1
	MEDIUM
	LARGE
)

type Beverage interface {
	GetDescription() string
	Cost() float32
	GetSize() Size
	SetSize(Size)
}

type BFields struct {
	Size Size
}

type Expresso struct {
	BFields
}

func (e *Expresso) GetSize() Size {
	return e.Size
}

func (e *Expresso) SetSize(size Size) {
	e.Size = size
}

func (e *Expresso) GetDescription() string {
	return "Espresso"
}

func (e *Expresso) Cost() float32 {
	return 1.99
}

func NewExpresso(size Size) Beverage {
	return &Expresso{
		BFields: BFields{
			Size: size,
		},
	}
}

type HouseBlend struct {
	BFields
}

func (h *HouseBlend) GetSize() Size {
	return h.Size
}

func (h *HouseBlend) SetSize(size Size) {
	h.Size = size
}

func (h *HouseBlend) GetDescription() string {
	return "House Blend"
}

func (h *HouseBlend) Cost() float32 {
	return 0.89
}

func NewHouseBlend(size Size) Beverage {
	return &HouseBlend{
		BFields: BFields{
			Size: size,
		},
	}
}
