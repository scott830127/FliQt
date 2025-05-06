package enum

type LeaveRecordStatus uint8 // 貼文審核狀態

const (
	_                         LeaveRecordStatus = iota
	LeaveRecordStatusPending                    // 待審核
	LeaveRecordStatusAccepted                   // 上架
	LeaveRecordStatusRejected                   // 下架
)
