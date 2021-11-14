package business

type BusinessType string

const (
	OTHER   BusinessType = "OTHER"
	INSERT  BusinessType = "INSERT"
	UPDATE  BusinessType = "UPDATE"
	DELETE  BusinessType = "DELETE"
	GRANT   BusinessType = "GRANT"
	EXPORT  BusinessType = "EXPORT"
	IMPORT  BusinessType = "IMPORT"
	FORCE   BusinessType = "FORCE"
	GENCODE BusinessType = "GENCODE"
	CLEAN   BusinessType = "CLEAN"
)

var businessTypeMap = map[BusinessType]int8{
	OTHER:   0,
	INSERT:  1,
	UPDATE:  2,
	DELETE:  3,
	GRANT:   4,
	EXPORT:  5,
	IMPORT:  6,
	FORCE:   7,
	GENCODE: 8,
	CLEAN:   9,
}

func (c BusinessType) Msg() int8 {
	msg, ok := businessTypeMap[c]
	if !ok {
		msg = businessTypeMap[OTHER]
	}
	return msg
}
