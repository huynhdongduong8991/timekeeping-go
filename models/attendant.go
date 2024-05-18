package models

import (
	"time"

	"time-keeping.com/lib"
)

type Attendant struct {
	ID         string    `json:"id"`
	EmployeeID string    `json:"employeeId"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	BadgeID    string    `json:"badgeId"`
}

func AddAttendant(record Attendant) error {
	conn := lib.DBConn
	_, err := conn.Exec("INSERT INTO attendant (employee_id, name, badge_id, created_at) VALUES ($1, $2, $3, $4)", record.EmployeeID, record.Name, record.BadgeID, record.CreatedAt)
	return err
}

func GetAllAttendants() ([]Attendant, error) {
	conn := lib.DBConn
	rows, err := conn.Query("SELECT * FROM attendant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	attendants := []Attendant{}
	for rows.Next() {
		var a Attendant
		err := rows.Scan(&a.ID, &a.EmployeeID, &a.Name, &a.BadgeID, &a.CreatedAt)
		if err != nil {
			return nil, err
		}
		attendants = append(attendants, a)
	}

	return attendants, nil
}

func GetAttendantByID(employeeID string) (*Attendant, error) {
	conn := lib.DBConn
	row := conn.QueryRow("SELECT * FROM attendant WHERE employee_id = $1", employeeID)

	var a Attendant
	err := row.Scan(&a.ID, &a.EmployeeID, &a.Name, &a.BadgeID, &a.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
