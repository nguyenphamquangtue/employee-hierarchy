package dto

type Employee struct {
	ID             int         `gorm:"primaryKey" json:"_id"`
	Name           string      `gorm:"uniqueIndex" json:"name"`
	SupervisorID   *int        `gorm:"index" json:"supervisor_id"`
	SupervisorName string      `gorm:"-" json:"supervisor_name,omitempty"`
	Subordinates   []*Employee `gorm:"foreignKey:SupervisorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subordinates"`
}
