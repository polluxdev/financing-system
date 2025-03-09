package helper

import "gorm.io/gorm"

func CommitAndRollback(tx *gorm.DB, err *error) func() {
	return func() {
		if *err != nil {
			tx.Rollback()
		} else {
			*err = tx.Commit().Error
			if err != nil {
				tx.Rollback()
			}
		}
	}
}
