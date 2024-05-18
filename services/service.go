package services

import (
	"time"

	"time-keeping.com/models"
)

type CreateAttendantParams struct {
	EmployeeID string    `json:"employeeID"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	BadgeID    string    `json:"badgeID"`
}

func CreateAttendant(params CreateAttendantParams) error {
	return models.AddAttendant(models.Attendant{
		EmployeeID: params.EmployeeID,
		Name:       params.Name,
		BadgeID:    params.BadgeID,
		CreatedAt:  params.CreatedAt,
	})
}

func GetAttendantByID(employeeID string) (*models.Attendant, error) {
	return models.GetAttendantByID(employeeID)
}

type AddAttendanceRecordParams struct {
	EmployeeID     string                `json:"employeeID"`
	Time           time.Time             `json:"time"`
	AttendanceType models.AttendanceType `json:"attendanceType"`
}

func AddAttendanceRecord(params AddAttendanceRecordParams) error {
	return models.AddAttendanceRecord(models.RecordAttendance{
		EmployeeID:     params.EmployeeID,
		Time:           params.Time,
		AttendanceType: params.AttendanceType,
	})
}

func GetRecordAttendancesByAttendantID(employeeID string) ([]models.RecordAttendance, error) {
	return models.GetRecordAttendancesByAttendantID(employeeID)
}

type ConfigParams struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func SetConfig(params ConfigParams) error {
	existConfig, err := models.ExistConfig(params.Field)
	if err != nil {
		return err
	}

	if existConfig {
		return models.UpdateConfig(
			params.Field,
			params.Value,
		)
	} else {
		return models.SetConfig(
			params.Field,
			params.Value,
		)
	}
}

func GetConfig(field string) (*models.Config, error) {
	return models.GetConfig(field)
}
