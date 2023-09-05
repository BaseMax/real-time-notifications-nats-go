package models

const (
	TASK_INPROGRESS = "IN-PROGRESS"
	TASK_DONE       = "DONE"
	TASK_CANCELED   = "CANCELED"
	TASK_BROWSE     = "BROWSE"

	ACTION_REGISTER   = "REGISTER"
	ACTION_NEW_RECORD = "NEW RECORD"
)

type Activity struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	RecieverID uint   `gorm:"not null" json:"user_id"`
	TaskID     uint   `gorm:"not null" json:"task_id"`
	Title      string `gorm:"not null" json:"title"`
	Action     string `gorm:"not null" json:"action"`
}

func GetActivitiesByUserId(id uint) (actitivies *[]Activity, herr *DbErr) {
	err := db.Find(&actitivies, "user_id = ?", id)
	return actitivies, errGormToHttp(err)
}

func SeenActivities(id uint) *DbErr {
	err := db.Delete(&Activity{}, "user_id = ?", id)
	return errGormToHttp(err)
}
