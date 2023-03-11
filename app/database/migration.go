package database

import (
	"fmt"
	class "immersiveApp/features/classes/data"
	feedback "immersiveApp/features/feedbacks/data"
	mentee "immersiveApp/features/mentees/data"
	status "immersiveApp/features/statuses/data"
	team "immersiveApp/features/teams/data"
	user "immersiveApp/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&team.Team{},
		&class.Class{},
		&status.Status{},
		&user.User{},
		&mentee.Mentee{},
		&feedback.Feedback{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
