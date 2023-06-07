package errorcode

import "employee-hierarchy-api/internal/response"

func Init() {
	response.Init()

	response.AddListCodes(user)

	response.AddListCodes(employee)

}
