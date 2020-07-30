package form

type RequestFormGroup struct {
	UUID    string `form:"uuid" binding:"required"`
	UniqKey string `form:"key" binding:"required"`
}

type ResponseFormGroup struct {
	Key      string
	HitGroup string
	Bucket   uint
}
