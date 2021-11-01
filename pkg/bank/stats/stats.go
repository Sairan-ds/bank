package stats

import "github.com/sairan-ds/bank/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var sum types.Money

	for _, i := range payments {
		sum += i.Amount
	}

	return sum / types.Money(len(payments))
}

