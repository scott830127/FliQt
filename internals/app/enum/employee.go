package enum

type EmployeePosition uint8 // 貼文審核狀態

const (
	_                  EmployeePosition = iota
	EmployeePositionJr                  // 待審核
	EmployeePositionSr                  // 上架
	EmployeePositionTl                  // 下架
)
