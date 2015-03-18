package mediator

import "fmt"

type StockOffer struct {
	stockShares   int
	stockSymbol   string
	colleagueCode int
}

func BuildStockOffer(shares int, stock string, code int) StockOffer {
	return StockOffer{shares, stock, code}
}

func (s StockOffer) getStockShares() int {
	return s.stockShares
}

func (s StockOffer) getStockSymbol() string {
	return s.stockSymbol
}

func (s StockOffer) getCollCode() int {
	return s.colleagueCode
}

///

type Colleague struct {
	mediator      Mediator
	colleagueCode int
}

func BuildColleague(mediator Mediator) Colleague {
	c := Colleague{mediator, 0}
	mediator.addColleague(c)
	return c
}

func (c Colleague) saleOffer(stock string, shares int) {
	c.mediator.saleOffer(stock, shares, c.colleagueCode)
}

func (c Colleague) buyOffer(stock string, shares int) {
	c.mediator.buyOffer(stock, shares, c.colleagueCode)
}

func (c *Colleague) setCollCode(collCode int) {
	c.colleagueCode = collCode
}

//

type GormanSlacks struct {
	Colleague
}

func BuildGormanSlacks(mediator Mediator) GormanSlacks {
	fmt.Println("%p", &mediator)
	fmt.Println("GormanSlacks signed up")
	return GormanSlacks{Colleague: BuildColleague(mediator)}
}

//

type JTPoorman struct {
	Colleague
}

func BuildJTPoorman(mediator Mediator) JTPoorman {
	fmt.Println("JTPoorman signed up")
	fmt.Println("%p", &mediator)
	return JTPoorman{Colleague: BuildColleague(mediator)}
}

//

type Mediator interface {
	saleOffer(string, int, int)
	buyOffer(string, int, int)
	addColleague(Colleague)
}

//

type StockMediator struct {
	colleagues      []Colleague
	stockBuyOffers  []StockOffer
	stockSaleOffers []StockOffer

	colleagueCodes int
}

func BuildStockMediator() *StockMediator {
	return &StockMediator{make([]Colleague, 0), make([]StockOffer, 0),
		make([]StockOffer, 0), 0}
}

func (s StockMediator) addColleague(newColleague Colleague) {
	s.colleagues = append(s.colleagues, newColleague)
	s.colleagueCodes = len(s.colleagues)
	newColleague.setCollCode(s.colleagueCodes)
}

func (s *StockMediator) saleOffer(stock string, shares int, collCode int) {
	var stockSold bool

	for i, offer := range s.stockBuyOffers {
		if offer.getStockSymbol() == stock &&
			offer.getStockShares() == shares {
			fmt.Println(shares, " shares of ", stock, " sold to colleague code", offer.getCollCode())
			copy(s.stockBuyOffers[i:], s.stockBuyOffers[i+1:])
			stockSold = true
		}
		if stockSold {
			break
		}
	}
	if !stockSold {
		fmt.Println(shares, " shares of ", stock, " added to inventory")
		newOffering := BuildStockOffer(shares, stock, collCode)

		s.stockSaleOffers = append(s.stockSaleOffers, newOffering)
		fmt.Println(s.stockSaleOffers)
	}
}

func (s *StockMediator) buyOffer(stock string, shares int, collCode int) {
	var stockBought bool

	for i, offer := range s.stockSaleOffers {
		if offer.getStockSymbol() == stock &&
			offer.getStockShares() == shares {
			fmt.Println(shares, " shares of ", stock, " bought by colleague code", offer.getCollCode())
			copy(s.stockSaleOffers[i:], s.stockSaleOffers[i+1:])
			stockBought = true
		}
		if stockBought {
			break
		}
	}
	if !stockBought {
		fmt.Println(shares, " shares of ", stock, " added to inventory")
		newOffering := BuildStockOffer(shares, stock, collCode)
		s.stockBuyOffers = append(s.stockBuyOffers, newOffering)
	}
}

func (s StockMediator) getStockOfferings() {
	for _, offer := range s.stockBuyOffers {
		if offer.getStockSymbol() != "" {
			fmt.Println(offer.getStockShares(), " of ", offer.getStockSymbol())
		}
	}
	for _, offer := range s.stockSaleOffers {
		if offer.getStockSymbol() != "" {
			fmt.Println(offer.getStockShares(), " of ", offer.getStockSymbol())
		}
	}
}
