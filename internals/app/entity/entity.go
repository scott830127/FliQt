package entity

func AutoMigrate() []any {
	return []any{
		new(Employee),
		new(LeaveRecord),
	}
}
