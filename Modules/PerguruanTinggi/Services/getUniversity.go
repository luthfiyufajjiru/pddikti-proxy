package Services

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/DataTransferObjects"
	"PDDiktiProxyAPI/Modules/ServerCaches"
	"strings"
)

func GetUniversity(input string) (result DataTransferObjects.PerguruanTinggiDTO, err error) {
	universities, err := ServerCaches.GetUniversities()
	defer func() {
		universities = nil
	}()
	if err == nil {
		for _, val := range *universities {
			input = strings.ReplaceAll(input, " ", "")
			if x, y := strings.ToLower(val.KodePt), strings.ToLower(input); strings.Contains(x, y) {
				result = val
				return
			}
		}
	}
	return
}
