package adminV1Controller

import (
	"SMT/repository"
	requestTypes "SMT/types/requests"
	stringTypes "SMT/types/strings"
	"SMT/utility"
	"SMT/validations"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddAssignment(c *gin.Context) {

	file, _ := c.FormFile("file")
	if file != nil {
		fmt.Println("file.Filename", file.Filename)
		fmt.Println("file.Header", file.Header)
		fmt.Println("file.Size", file.Size)
	}

	var requestBody requestTypes.AddAssignment
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, err.Error())
		return
	}

	// fmt.Println("requestBody", c.PostForm("name"))
	// fmt.Println("requestBody", c.PostForm("semester"))
	// requestBody, err := c.MultipartForm()
	// fmt.Println(err)
	// fmt.Println(requestBody.Value["deadline"][0])
	// fmt.Println(requestBody.Value["department_id"][0])
	// fmt.Println(requestBody.Value["semester"][0])
	// fmt.Println(requestBody.Value["name"][0])

	// jsonEncodedRequestBody, _ := json.Marshal(requestBody.Value)
	// json.Unmarshal(jsonEncodedRequestBody, &request)
	// fmt.Println("request", request)

	// c.String(200, "done")

	// return

	if err := validations.DateValidation(requestBody.Deadline); err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, stringTypes.INVALID_DATE)
		return
	}
	if !repository.IsValidDepartmentID(int(requestBody.DepartmentID)) {
		utility.ErrorResponse(c, stringTypes.INVALID_DEPARTMENT_ID)
		return
	}
	if err := repository.AddAssignment(requestBody); err != nil {
		log.Default().Println(err)
		utility.ErrorResponse(c, stringTypes.ASSIGNMENT_INSERT_FAIL)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.ASSIGNMENT_INSERTED)
}

func DeleteAssignment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utility.ErrorResponse(c, stringTypes.INVALID_DEPARTMENT_ID)
		return
	}
	if !repository.DeleteAssignment(id) {
		utility.ErrorResponse(c, stringTypes.INVALID_ASSIGNMENT_ID)
		return
	}
	utility.SuccessResponseWithoutData(c, stringTypes.ASSIGNMENT_DELETED)
}

func GetAssignments(c *gin.Context) {
	var departmentID, semester string
	departmentID, exists := c.GetQuery("department-id")
	if !exists {
		utility.ErrorResponse(c, stringTypes.DEPARTMENT_ID_NOT_FOUND)
		return
	}
	semester, exists = c.GetQuery("semester")
	if !exists {
		utility.ErrorResponse(c, stringTypes.SEMESTER_NOT_FOUND)
		return
	}
	if semester < "1" || semester > "8" {
		utility.ErrorResponse(c, stringTypes.INVALID_SEMESTER)
		return
	}
	assignments := repository.GetAssignments(departmentID, semester)
	if len(assignments) > 0 {
		utility.SuccessResponseWithData(c, stringTypes.ASSIGNMENTS_FETCHED, assignments, len(assignments))
		return
	}
	utility.ErrorResponse(c, stringTypes.ASSIGNMENTS_NOT_FOUND)
}
