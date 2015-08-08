package matriarch

type denizen struct {
	TradeItem string
	Wish string
	HasTraded bool
}

func NewDenizen() *denizen {
	return &denizen{
		TradeItem: "flower",
		Wish: "fish",
		HasTraded: false,
	}
}
