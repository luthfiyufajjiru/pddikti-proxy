package Services

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/DataTransferObjects"
	"PDDiktiProxyAPI/Modules/ServerCaches"
	"strings"
)

func SearchUniversity(query string) (results []DataTransferObjects.PerguruanTinggiDTO, err error) {
	universities, err := ServerCaches.GetUniversities()
	defer func() {
		universities = nil
	}()
	if err == nil {
		for _, val := range *universities {
			query = strings.ReplaceAll(query, "%20", " ")
			if x, y := strings.ToLower(val.NamaPt), strings.ToLower(query); strings.Contains(x, y) {
				if results == nil {
					results = make([]DataTransferObjects.PerguruanTinggiDTO, 0, 5)
					results = append(results, val)
				} else if len(results) < cap(results) && val.NamaPt != "" {
					results = append(results, val)
				} else if len(results) > cap(results) {
					return
				}
			}
		}
	}
	return
}
