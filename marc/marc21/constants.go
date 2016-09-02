package marc21

const (
	Unspecified byte = ' '

	// Record status
	IncreaseEncoding                   = 'a'
	StatusCorrected                    = 'c'
	StatusDeleted                      = 'd'
	StatusNew                          = 'n'
	IncreaseEncodingFromPrepublication = 'p'

	// Type of record
	TypeMonographPart       = 'a'
	TypeSerialPart          = 'b'
	TypeCollection          = 'c'
	TypeSubunit             = 'd'
	TypeIntegratingResource = 'i'
	TypeMonographItem       = 'm'
	TypeSerial              = 's'

	// Bibliographic level

	// Type of control

	// Character encoding scheme
)
