package data

const (
	TYPE_NAME    = "name"
	TYPE_DATE    = "date"
	TYPE_ADDRESS = "address"
	TYPE_PHONE   = "phone"

	SUBTYPE_STREET = "street"
	SUBTYPE_CITY   = "city"
)

var Supported = map[string]bool{
	TYPE_NAME:    true,
	TYPE_DATE:    true,
	TYPE_ADDRESS: true,
	TYPE_PHONE:   true,
}
