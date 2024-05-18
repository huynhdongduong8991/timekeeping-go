package models

import (
	"time"

	"time-keeping.com/lib"
)

type AttendanceType string

var (
	AttendanceTypeCheckIn  = "CHECKIN"
	AttendanceTypeCheckOut = "CHECKOUT"
)

type RecordAttendance struct {
	ID             string         `json:"id"`
	EmployeeID     string         `json:"employeeId"`
	Time           time.Time      `json:"time"`
	AttendanceType AttendanceType `json:"attendanceType"`
}

func AddAttendanceRecord(record RecordAttendance) error {
	conn := lib.DBConn
	_, err := conn.Exec("INSERT INTO record_attendance (employee_id, time, attendance_type) VALUES ($1, $2, $3)", record.EmployeeID, record.Time, record.AttendanceType)
	return err
}

func GetRecordAttendancesByAttendantID(employeeID string) ([]RecordAttendance, error) {
	conn := lib.DBConn
	rows, err := conn.Query("SELECT * FROM record_attendance WHERE employee_id = $1", employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recordAttendances := []RecordAttendance{}
	for rows.Next() {
		var r RecordAttendance
		err := rows.Scan(&r.ID, &r.EmployeeID, &r.Time, &r.AttendanceType)
		if err != nil {
			return nil, err
		}
		recordAttendances = append(recordAttendances, r)
	}

	return recordAttendances, nil
}
