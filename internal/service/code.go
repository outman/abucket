package service

// service option code
const (
	ServiceOptionSuccess = iota
	ServiceOptionRecordExists
	ServiceOptionRecordNameDuplicate
	ServiceOptionRecordKeyDuplicate
	ServiceOptionLayerNoSpace
	ServiceOptionDbError
	ServiceOptionSQLError
)

// CodeMessage display message text
func CodeMessage(code int) string {
	message := make(map[int]string)
	message[ServiceOptionSuccess] = "操作成功"
	message[ServiceOptionRecordExists] = "记录已经存在"
	message[ServiceOptionRecordNameDuplicate] = "记录名称重复"
	message[ServiceOptionRecordKeyDuplicate] = "记录唯一标试重复"
	message[ServiceOptionLayerNoSpace] = "层无流量可用，选择新层"
	message[ServiceOptionDbError] = "操作失败，请稍后再试"
	message[ServiceOptionSQLError] = "操作失败，请稍后再试 [SQL]"
	return message[code]
}
