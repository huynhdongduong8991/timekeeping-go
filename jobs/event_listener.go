package jobs

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"time-keeping.com/lib"
	"time-keeping.com/models"
	"time-keeping.com/services"
)

type SetEmployee struct {
	EmployeeID common.Address `json:"employeeID"`
	Name       string         `json:"name"`
	BadgeID    string         `json:"badgeID"`
	CreatedAt  *big.Int       `json:"createdAt"`
}

type RecordAttendance struct {
	EmployeeID     common.Address `json:"employeeID"`
	Time           *big.Int       `json:"time"`
	AttendanceType uint8          `json:"attendanceType"`
}

func handleEvents(_log types.Log) {
	abi, err := lib.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	switch _log.Topics[0].Hex() {
	case abi.Events["SetEmployee"].ID.String():
		var setEmployee SetEmployee
		err := abi.UnpackIntoInterface(&setEmployee, "SetEmployee", _log.Data)
		if err != nil {
			log.Fatal(err)
		}

		err = services.CreateAttendant(services.CreateAttendantParams{
			EmployeeID: setEmployee.EmployeeID.String(),
			Name:       setEmployee.Name,
			BadgeID:    setEmployee.BadgeID,
			CreatedAt:  time.Unix(setEmployee.CreatedAt.Int64(), 0),
		})
		if err != nil {
			log.Fatal(err)
		}
	case abi.Events["RecordAttendance"].ID.String():
		var recordAttendance RecordAttendance
		abi.UnpackIntoInterface(&recordAttendance, "RecordAttendance", _log.Data)
		if err != nil {
			log.Fatal(err)
		}

		var attendanceType models.AttendanceType
		if recordAttendance.AttendanceType == 0 {
			attendanceType = models.AttendanceType(models.AttendanceTypeCheckIn)
		} else {
			attendanceType = models.AttendanceType(models.AttendanceTypeCheckOut)
		}

		err = services.AddAttendanceRecord(services.AddAttendanceRecordParams{
			EmployeeID:     recordAttendance.EmployeeID.String(),
			Time:           time.Unix(recordAttendance.Time.Int64(), 0),
			AttendanceType: attendanceType,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func afterHandleEvents(lastLog *types.Log) {
	latestBlock, err := lib.GetLastetBlock()
	if err != nil {
		log.Fatal(err)
	}

	err = services.SetConfig(services.ConfigParams{
		Field: "latestBlock",
		Value: latestBlock.Number.String(),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("latestBlock", latestBlock.Number.String())
}

func handleListenToNewBlocks(log types.Log) {
	handleEvents(log)
	afterHandleEvents(&log)
}

func ListenToEvents() {
	envConfig, err := lib.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	config, err := services.GetConfig("latestBlock")
	if err != nil {
		if err == sql.ErrNoRows {
			config = nil
		} else {
			log.Fatal(err)
		}
	}

	var startBlock int64
	if config == nil {
		startBlock = envConfig.Contract.FROM_BLOCK
	} else {
		latestBlock, err := strconv.Atoi(config.Value)
		if err != nil {
			log.Fatal(err)
		}
		startBlock = int64(latestBlock)
	}

	fmt.Println("startBlock", startBlock)

	lib.GetAllEvents(startBlock, handleEvents, afterHandleEvents)
	go lib.ListenToNewBlocks(handleListenToNewBlocks)
}
