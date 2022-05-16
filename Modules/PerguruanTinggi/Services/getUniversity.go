package Services

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/DataTransferObjects"
	"PDDiktiProxyAPI/Modules/ServerCaches"
	"strings"
)

func GetUniversityByKode(input string) (result DataTransferObjects.PerguruanTinggiDTO, err error) {
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

func GetUniversityByName(input string) (result DataTransferObjects.PerguruanTinggiDTO, err error) {
	universities, err := ServerCaches.GetUniversities()
	defer func() {
		universities = nil
	}()
	if err == nil {
		for _, val := range *universities {
			input = strings.ReplaceAll(input, "%20", " ")
			if x, y := strings.ToLower(val.NamaPt), strings.ToLower(input); x == y {
				result = val
				return
			}
		}
	}
	return
}
