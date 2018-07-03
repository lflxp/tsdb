package pkg

// freelist represents a list of all pages that are available for allocation.
// It also tracks pages that have been freed but are still in use by open transactions.
type freelist struct {
	ids []pgid // all free and available free page ids.
	// pending map[txid][]pgid // mapping of soon-to-be free page ids by tx.
	cache map[pgid]bool // fast lookup of all free and pending page ids.
}
