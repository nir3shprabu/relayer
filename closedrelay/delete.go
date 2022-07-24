package main

func (b *ClosedRelay) DeleteEvent(id string, pubkey string) error {
	_, err := b.DB.Exec("DELETE FROM events WHERE id = $1 AND pubkey = $2")
	return err
}
