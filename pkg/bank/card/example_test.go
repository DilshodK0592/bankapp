package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,

		},
	}))
	

	
	//Output: 
	//100000
	//300000
	//0
	//0
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Active: true,
			Balance: 100,
			PAN: "1111 xxxx xxxx 1111",
			
		},
		{
			Active: true,
			Balance: -100,
			PAN: "1111 xxxx xxxx 1112",
		},
		{
			Active: false,
			Balance: 100,
			PAN: "1111 xxxx xxxx 1113",
		},
	}
	payments := PaymentSources(cards)

	fmt.Println(payments[0].Number)
	//Output: 1111 xxxx xxxx 1111
	
}