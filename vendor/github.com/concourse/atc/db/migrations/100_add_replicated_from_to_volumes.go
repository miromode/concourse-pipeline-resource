package migrations

import "github.com/BurntSushi/migration"

func AddReplicatedFromToVolumes(tx migration.LimitedTx) error {
	_, err := tx.Exec(`
		ALTER TABLE volumes ADD COLUMN replicated_from text DEFAULT null;
	`)
	if err != nil {
		return err
	}

	return nil
}
