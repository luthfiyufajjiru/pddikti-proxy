package DataTransferObjects

type RasioProdiDTO struct {
	Dosen     *int64 `json:"dosen"`
	DosenNidk *int64 `json:"dosenNidk"`
	DosenNidn *int64 `json:"dosenNidn"`
	Mahasiswa *int64 `json:"mahasiswa"`
	Semester  string `json:"semester"`
}

type ProdiDTO struct {
	Akreditasi string          `json:"akreditasi"`
	IdSms      string          `json:"id_sms"`
	Jenjang    string          `json:"jenjang"`
	KodeProdi  string          `json:"kode_prodi"`
	NmLemb     string          `json:"nm_lemb"`
	RasioList  []RasioProdiDTO `json:"rasio_list"`
	StatProdi  string          `json:"stat_prodi"`
}
