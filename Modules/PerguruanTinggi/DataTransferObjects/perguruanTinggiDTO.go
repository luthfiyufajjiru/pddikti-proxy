package DataTransferObjects

import (
	"strings"
)

type PerguruanTinggiDTO struct {
	IdSp   string `json:"id_sp",`
	KodePt string `json:"kode_pt"`
	NamaPt string `json:"nama_pt"`
}

func (p PerguruanTinggiDTO) GetKodePt() {
	strings.ReplaceAll(p.KodePt, " ", "")
}
