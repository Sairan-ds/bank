package stats

import (
	"fmt"

	"github.com/sairan-ds/bank/pkg/bank/types"
)

func ExampleAvg() {

	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10000,
			
		},
		{
			ID: 2,
			Amount: 20000,
		},
		{
			ID: 3,
			Amount: 30000,
		},
		
	}
	
	fmt.Println(Avg(payments))


	// Output: 
	// 20_000

}
