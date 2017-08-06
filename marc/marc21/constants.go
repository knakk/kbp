package marc21

// Leader fields
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

type ControlFieldPos struct {
	Pos    int
	Length int
}

// Control fields
//
// 008 - Fixed-Length Data Elements-General Information (NR)
// Field has no indicators or subfield codes; the data elements are positionally defined by type of material.
// Descriptions of the elements defined for field 008 positions 18-34 are in seven separate sections
// corresponding to the following type of material configurations: Books (BK), Computer Files (CF),
// Maps (MP), Music (MU), Continuing Resources (CR), Visual Materials (VM), and Mixed Materials (MX).
// In this general section, validity of a particular field 008 data element is indicated by V in
// the composite field 008 table.
var (
	// All materials
	C008DateEnteredOnFile                       = ControlFieldPos{Pos: 0, Length: 4}
	C008TypeOfDateOrPublicationStatus           = ControlFieldPos{Pos: 5, Length: 1}
	C008Date1                                   = ControlFieldPos{Pos: 7, Length: 4}
	C008Date2                                   = ControlFieldPos{Pos: 11, Length: 4}
	C008PlaceOfPublicationProductionOrExecution = ControlFieldPos{Pos: 15, Length: 3}
	C008Language                                = ControlFieldPos{Pos: 35, Length: 3}
	C008ModifiedRecord                          = ControlFieldPos{Pos: 38, Length: 1}
	C008CatalogingSource                        = ControlFieldPos{Pos: 39, Length: 1}
	// Books
	// 18-34 - [See one of the seven separate 008/18-34 configuration sections for these elements.]
)

const (
	// Literary forms (C008 pos 33)
	LiteraryFormNotFiction   = '0'
	LiteraryFormFiction      = '1'
	LiteraryFormDramas       = 'd'
	LiteraryFormEssays       = 'e'
	LiteraryFormNovels       = 'f'
	LiteraryFormHumorSatires = 'h'
	LiteraryFormLetters      = 'i'
	LiteraryFormShortStories = 'j'
	LiteraryFormMixedForms   = 'm'
	LiteraryFormPoetry       = 'p'
	LiteraryFormSpeeches     = 's'
	LiteraryFormUnknown      = 'u'
)
