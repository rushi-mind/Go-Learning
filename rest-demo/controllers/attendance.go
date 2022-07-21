package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-demo/config"
	"rest-demo/structs"
)

func GetAttendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getAttendanceResult, err := config.DB.Query(`select * from attendance`)
	if err != nil {
		log.Fatal(err)
	}

	var data []structs.Attendance
	for getAttendanceResult.Next() {
		var attendance structs.Attendance
		_ = getAttendanceResult.Scan(&attendance.Sr, &attendance.Id, &attendance.Date, &attendance.CurrentStatus, &attendance.Intime, &attendance.Outtime, &attendance.EffectiveHours)
		data = append(data, attendance)
	}

	json.NewEncoder(w).Encode(data)
}
