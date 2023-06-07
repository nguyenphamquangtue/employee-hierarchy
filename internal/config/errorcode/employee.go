package errorcode

import "employee-hierarchy-api/internal/response"

const (
	EmployeeNotFound                   = "employee_not_found"
	EmployeeIDRequired                 = "employee_id_required"
	EmployeeNameRequired               = "employee_name_required"
	SupervisorIDRequired               = "supervisor_id_required"
	EmployeeExisted                    = "employee_existed"
	SupervisorNotFound                 = "supervisor_not_found"
	CycleDetectedInSupervisorHierarchy = "cycle_detected_in_supervisor_hierarchy"
	SubordinateCannotBeASupervisor     = "subordinate_cannot_be_a_supervisor"
)

const (
	employeeNotFoundCode = iota + 700
	employeeIDRequiredCode
	employeeNameRequired
	supervisorIDRequired
	employeeExistedCode
	supervisorNotFoundCode
	CycleDetectedInSupervisorHierarchyCode
	SubordinateCannotBeASupervisorCode
)

var employee = []response.Code{
	{
		Key:     EmployeeNotFound,
		Code:    employeeNotFoundCode,
		Message: "Employee Not Found",
	},
	{
		Key:     EmployeeIDRequired,
		Code:    employeeIDRequiredCode,
		Message: "EmployeeID Required",
	},
	{
		Key:     EmployeeNameRequired,
		Code:    employeeNameRequired,
		Message: "Employee Name Required",
	},
	{
		Key:     SupervisorIDRequired,
		Code:    supervisorIDRequired,
		Message: "Supervisor ID Required",
	},
	{
		Key:     EmployeeExisted,
		Code:    employeeExistedCode,
		Message: "Employee Existed",
	},
	{
		Key:     SupervisorNotFound,
		Code:    supervisorNotFoundCode,
		Message: "Supervisor Not Found",
	},
	{
		Key:     CycleDetectedInSupervisorHierarchy,
		Code:    CycleDetectedInSupervisorHierarchyCode,
		Message: "Cycle Detected In Supervisor Hierarchy",
	},
	{
		Key:     SubordinateCannotBeASupervisor,
		Code:    SubordinateCannotBeASupervisorCode,
		Message: "Subordinate Cannot Be A Supervisor",
	},
}
