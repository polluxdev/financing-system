package helper

func ContainUint8(slice []uint8, item uint8) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
