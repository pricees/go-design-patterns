package visitor

type Visitor interface {
	visit(Visitable) float64
}

type TaxVisitor struct{}

type Visitable interface {
	Accept(Visitor) float64
	GetPrice() float64
}

func (t TaxVisitor) visit(v Visitable) float64 {
	var tax float64

	// Type switch as a means of polymorphism
	switch v.(type) {
	case Liquor:
		tax = 0.18
	case Tobacco:
		tax = 0.32
	default: // i.e. Necessity
		tax = 0.0
	}

	return tax * v.GetPrice()
}

// Use composition to reuse the codez

type GoodForSale struct {
	price float64
}

func (p GoodForSale) Accept(v Visitor) float64 {
	return v.visit(p)
}

func (p GoodForSale) GetPrice() float64 {
	return p.price
}

type Tobacco struct {
	GoodForSale
}

func (t Tobacco) Accept(v Visitor) float64 {
	return v.visit(t)
}

func BuildTobacco(price float64) Tobacco {
	return Tobacco{GoodForSale{price: price}}
}

type Liquor struct {
	GoodForSale
}

func (l Liquor) Accept(v Visitor) float64 {
	return v.visit(l)
}

func BuildLiquor(price float64) Liquor {
	return Liquor{GoodForSale{price: price}}
}

type Necessity struct {
	GoodForSale
}

func BuildNecessity(price float64) Necessity {
	return Necessity{GoodForSale{price: price}}
}
