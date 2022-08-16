package syntest

type TicketStore struct {
	ticket *uint64
	done   *uint64
	slots  []string
}
