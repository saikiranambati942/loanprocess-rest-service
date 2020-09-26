package handlers

import "time"

var (
	l           loan
	b           balance
	m           map[time.Time]balance // map that is used to store paymentDate and newbalance
	lsd         time.Time             // loan start date
	prevPaydate string                // previous payment date
)
