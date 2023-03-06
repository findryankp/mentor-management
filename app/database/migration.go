package database

import (
	class "clean-arch/features/classes/data"
	feedback "clean-arch/features/feedbacks/data"
	mentee "clean-arch/features/mentees/data"
	status "clean-arch/features/statuses/data"
	team "clean-arch/features/teams/data"
	user "clean-arch/features/users/data"
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&team.Team{},
		&user.User{},
		&class.Class{},
		&status.Status{},
		&mentee.Mentee{},
		&feedback.Feedback{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
