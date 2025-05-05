package dto

type LeaveRequest struct {
	EmployeeID uint   `json:"employee_id" binding:"required"`
	Reason     string `json:"reason" binding:"required"`
	Days       int    `json:"days" binding:"required,min=1"`
}
