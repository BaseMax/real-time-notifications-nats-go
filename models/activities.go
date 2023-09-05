package models

const (
	TASK_INPROGRESS = "IN-PROGRESS"
	TASK_DONE       = "DONE"
	TASK_CANCELED   = "CANCELED"
	TASK_BROWSE     = "BROWSE"

	ACTION_REGISTER  = "REGISTER"
	ACTION_NEW_ORDER = "NEW ORDER"
)

type Activity struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"not null" json:"user_id"`
	Title  string `gorm:"not null" json:"title"`
	Action string `gorm:"not null" json:"action"`
}

func GetActivitiesByUserId(id uint) (actitivies *[]Activity, herr *DbErr) {
	err := db.Find(&actitivies, "user_id = ?", id)
	return actitivies, errGormToHttp(err)
}

func SeenActivities(id uint) *DbErr {
	err := db.Delete(&Activity{}, "user_id = ?", id)
	return errGormToHttp(err)
}
