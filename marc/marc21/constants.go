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
	C008DateEnteredOnFile                       = ControlFieldPos{Pos: 0, Length: 6}
	C008TypeOfDateOrPublicationStatus           = ControlFieldPos{Pos: 6, Length: 1}
	C008Date1                                   = ControlFieldPos{Pos: 7, Length: 4}
	C008Date2                                   = ControlFieldPos{Pos: 11, Length: 4}
	C008PlaceOfPublicationProductionOrExecution = ControlFieldPos{Pos: 15, Length: 3}
	C008Language                                = ControlFieldPos{Pos: 35, Length: 3}
	C008ModifiedRecord                          = ControlFieldPos{Pos: 38, Length: 1}
	C008CatalogingSource                        = ControlFieldPos{Pos: 39, Length: 1}
	// Books
	// 18-34 - [See one of the seven separate 008/18-34 configuration sections for these elements.]
)

// Literary forms (C008 pos 33)
const (
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

// Relator code list
// https://www.loc.gov/marc/relators/relaterm.html
const (
	RelatorAbridger                            = "abr"
	RelatorActor                               = "act"
	RelatorAdapter                             = "adp"
	RelatorAddressee                           = "rcp"
	RelatorAnalyst                             = "anl"
	RelatorAnimator                            = "anm"
	RelatorAnnotator                           = "ann"
	RelatorAppellant                           = "apl"
	RelatorAppellee                            = "ape"
	RelatorApplicant                           = "app"
	RelatorArchitect                           = "arc"
	RelatorArranger                            = "arr"
	RelatorArtCopyist                          = "acp"
	RelatorArtDirector                         = "adi"
	RelatorArtist                              = "art"
	RelatorArtisticDirector                    = "ard"
	RelatorAssignee                            = "asg"
	RelatorAssociatedName                      = "asn"
	RelatorAttributedName                      = "att"
	RelatorAuctioneer                          = "auc"
	RelatorAuthor                              = "aut"
	RelatorAuthorInQuotationsOrTextAbstracts   = "aqt"
	RelatorAuthorOfAfterwordOrClophon          = "aft"
	RelatorAuthorOfDialog                      = "aud"
	RelatorAuthorOfIntroduction                = "aui"
	RelatorAutographer                         = "ato"
	RelatorBibliographicAntecedent             = "ant"
	RelatorBinder                              = "bnd"
	RelatorBindingDesigner                     = "bdd"
	RelatorBlurbWriter                         = "blw"
	RelatorBookDesigner                        = "bkd"
	RelatorBookjacketDesigner                  = "bjd"
	RelatorBookplateDesigner                   = "bpd"
	RelatorBookProducer                        = "bkp"
	RelatorBookseller                          = "bsl"
	RelatorBrailleEmbosser                     = "brl"
	RelatorBroadcaster                         = "brd"
	RelatorCalligrapher                        = "cll"
	RelatorCartographer                        = "ctg"
	RelatorCaster                              = "cas"
	RelatorCensor                              = "cns"
	RelatorChoreographer                       = "chr"
	RelatorCinematographer                     = "cng"
	RelatorClient                              = "cli"
	RelatorCollaborator                        = "-clb" // Discontinued
	RelatorCollectionRegistrar                 = "cor"
	RelatorCollector                           = "col"
	RelatorCollotyper                          = "clt"
	RelatorColorist                            = "clr"
	RelatorCommentator                         = "cmm"
	RelatorCommentatorForWrittenText           = "cwt"
	RelatorCompiler                            = "com"
	RelatorComplainant                         = "cpl"
	RelatorComplainantAppellant                = "cpt"
	RelatorComplainantAppellee                 = "cpe"
	RelatorComposer                            = "cmp"
	RelatorCompositor                          = "cmt"
	RelatorConceptor                           = "ccp"
	RelatorConductor                           = "cnd"
	RelatorConservator                         = "con"
	RelatorConsultant                          = "csl"
	RelatorConsultantToProject                 = "csp"
	RelatorContestant                          = "cos"
	RelatorContestantAppellant                 = "cot"
	RelatorContestantAppellee                  = "coe"
	RelatorContestee                           = "cts"
	RelatorContesteeAppellant                  = "ctt"
	RelatorContesteeAppellee                   = "cte"
	RelatorContractor                          = "ctr"
	RelatorContributor                         = "ctb"
	RelatorCopyrightClaimant                   = "cpc"
	RelatorCopyrightHolder                     = "cph"
	RelatorCorrector                           = "crr"
	RelatorCorrespondent                       = "crp"
	RelatorCostumeDesigner                     = "cst"
	RelatorCourtGoverned                       = "cou"
	RelatorCourtReporter                       = "crt"
	RelatorCoverDesigner                       = "cov"
	RelatorCreator                             = "cre"
	RelatorCurator                             = "cur"
	RelatorDancer                              = "dnc"
	RelatorDataContributor                     = "dtc"
	RelatorDataManager                         = "dtm"
	RelatorDedicatee                           = "dte"
	RelatorDedicator                           = "dto"
	RelatorDefendant                           = "dfd"
	RelatorDefendantAppellant                  = "dft"
	RelatorDefendantAppellee                   = "dfe"
	RelatorDegreeGrantingInstitution           = "dgg"
	RelatorDegreeSupervisor                    = "dgs"
	RelatorDelineator                          = "dln"
	RelatorDepicted                            = "dpc"
	RelatorDepositor                           = "dpt"
	RelatorDesigner                            = "dsr"
	RelatorDirector                            = "drt"
	RelatorDissertant                          = "dis"
	RelatorDistributionPlace                   = "dbp"
	RelatorDistributor                         = "dst"
	RelatorDonor                               = "dnr"
	RelatorDraftsman                           = "drm"
	RelatorDubiousAuthor                       = "dub"
	RelatorEditor                              = "edt"
	RelatorEditorOfCompilation                 = "edc"
	RelatorEditorOfMovingImageWork             = "edm"
	RelatorElectrician                         = "elg"
	RelatorElectrotyper                        = "elt"
	RelatorEnactingJurisdiction                = "enj"
	RelatorEngineer                            = "eng"
	RelatorEngraver                            = "egr"
	RelatorEtcher                              = "etr"
	RelatorEvenPlace                           = "evp"
	RelatorExpert                              = "exp"
	RelatorFacsimilist                         = "fac"
	RelatorFieldDirector                       = "fld"
	RelatorFilmDirector                        = "fmd"
	RelatorFilmDistributor                     = "fds"
	RelatorFilmEditor                          = "flm"
	RelatorFilmmaker                           = "fmk"
	RelatorFilmProducer                        = "fmp"
	RelatorFirstParty                          = "fpy"
	RelatorForger                              = "frg"
	RelatorFormerOwner                         = "fmo"
	RelatorFunder                              = "fnd"
	RelatorGeographicInformationSpecialist     = "gis"
	RelatorGraphicTechnician                   = "-grt" // Discontinued
	RelatorHonoree                             = "hnr"
	RelatorHost                                = "hst"
	RelatorHostInstitution                     = "his"
	RelatorIlluminator                         = "ilu"
	RelatorIllustrator                         = "ill"
	RelatorInscriber                           = "ins"
	RelatorInstrumentalist                     = "itr"
	RelatorInterviewee                         = "ive"
	RelatorInterviewer                         = "ivr"
	RelatorInventor                            = "inv"
	RelatorIssuingBody                         = "isb"
	RelatorJudge                               = "jud"
	RelatorJurisdictionGoverned                = "jug"
	RelatorLaboratory                          = "lbr"
	RelatorLaboratoryDirector                  = "ldr"
	RelatorLandscapeArchitect                  = "lsa"
	RelatorLead                                = "led"
	RelatorLender                              = "len"
	RelatorLibelant                            = "lil"
	RelatorLibelantAppellant                   = "lit"
	RelatorLibelantAppellee                    = "lie"
	RelatorLibelee                             = "lel"
	RelatorLibeleeAppellant                    = "let"
	RelatorLibeleeAppellee                     = "lee"
	RelatorLibrettist                          = "lbt"
	RelatorLicensee                            = "lse"
	RelatorLicensor                            = "lso"
	RelatorLightingDesigner                    = "lgd"
	RelatorLithographer                        = "ltg"
	RelatorLyricist                            = "lyr"
	RelatorManufacturePlace                    = "mfp"
	RelatorManufacturer                        = "mfr"
	RelatorMarbler                             = "mrb"
	RelatorMarkupEditor                        = "mrk"
	RelatorMedium                              = "med"
	RelatorMetadataContact                     = "mdc"
	RelatorMetalEngraver                       = "mte"
	RelatorMinuteTaker                         = "mtk"
	RelatorModerator                           = "mod"
	RelatorMonitor                             = "mon"
	RelatorMusicalDirector                     = "msd"
	RelatorMusicCopyist                        = "mcp"
	RelatorMusician                            = "mus"
	RelatorNarrator                            = "nrt"
	RelatorOnscreenPresenter                   = "osp"
	RelatorOpponent                            = "opn"
	RelatorOrganizer                           = "orm"
	RelatorOriginator                          = "org"
	RelatorOther                               = "oth"
	RelatorOwner                               = "own"
	RelatorPanelist                            = "pan"
	RelatorPapermaker                          = "ppm"
	RelatorPatentApplicant                     = "pta"
	RelatorPatentHolder                        = "pth"
	RelatorPatron                              = "pat"
	RelatorPerformer                           = "prf"
	RelatorPermittingAgency                    = "pma"
	RelatorPhotographer                        = "pht"
	RelatorPlaintiff                           = "ptf"
	RelatorPlaintiffAppellant                  = "ptt"
	RelatorPlaintiffAppellee                   = "pte"
	RelatorPlatemaker                          = "plt"
	RelatorPraeses                             = "pra"
	RelatorPresenter                           = "pre"
	RelatorPrinter                             = "prt"
	RelatorPrinterOfPlates                     = "pop"
	RelatorPrintmaker                          = "prm"
	RelatorProcessContact                      = "prc"
	RelatorProducer                            = "pro"
	RelatorProductionCompany                   = "prn"
	RelatorProductionDesigner                  = "prs"
	RelatorProductionManager                   = "pmn"
	RelatorProductionPersonnel                 = "prd"
	RelatorProductionPlace                     = "prp"
	RelatorProgrammer                          = "prg"
	RelatorProjectDirector                     = "pdr"
	RelatorProofreader                         = "pfr"
	RelatorProvider                            = "prv"
	RelatorPublicationPlace                    = "pup"
	RelatorPublisher                           = "pbl"
	RelatorPublishingDirector                  = "pbd"
	RelatorPuppeteer                           = "ppt"
	RelatorRadioDirector                       = "rdd"
	RelatorRadioProducer                       = "rpc"
	RelatorRecordingEngineer                   = "rce"
	RelatorRecordist                           = "rcd"
	RelatorRedaktor                            = "red"
	RelatorRenderer                            = "ren"
	RelatorReporter                            = "rpt"
	RelatorRepository                          = "rps"
	RelatorResearcher                          = "res"
	RelatorResearchTeamHead                    = "rth"
	RelatorResearchTeamMember                  = "rtm"
	RelatorRespondent                          = "rsp"
	RelatorRespondentAppellant                 = "rst"
	RelatorRespondentAppellee                  = "rse"
	RelatorResponsibleParty                    = "rpy"
	RelatorRestager                            = "rsg"
	RelatorRestorationist                      = "rsr"
	RelatorReviewer                            = "rev"
	RelatorRubricator                          = "rbr"
	RelatorScenarist                           = "sce"
	RelatorScientificAdvisor                   = "sad"
	RelatorScreenwriter                        = "aus"
	RelatorScribe                              = "scr"
	RelatorSculptor                            = "scl"
	RelatorSecondParty                         = "spy"
	RelatorSecretary                           = "sec"
	RelatorSeller                              = "sll"
	RelatorSetDesigner                         = "std"
	RelatorSetting                             = "stg"
	RelatorSigner                              = "sgn"
	RelatorSinger                              = "sng"
	RelatorSoundDesigner                       = "sds"
	RelatorSpeaker                             = "spk"
	RelatorSponsor                             = "spn"
	RelatorStageDirector                       = "sgd"
	RelatorStageManager                        = "stm"
	RelatorStandardsBody                       = "stn"
	RelatorStereotyper                         = "str"
	RelatorStoryteller                         = "stl"
	RelatorSupportingHost                      = "sht"
	RelatorSurveyor                            = "srv"
	RelatorTeacher                             = "tch"
	RelatorTechnicalDirector                   = "tcd"
	RelatorTelevisionDirector                  = "tld"
	RelatorTelevisionProducer                  = "tlp"
	RelatorThesisAdvisor                       = "ths"
	RelatorTranscriber                         = "trc"
	RelatorTranslator                          = "trl"
	RelatorTypeDesigner                        = "tyd"
	RelatorTypographer                         = "tyg"
	RelatorUniversityPlace                     = "uvp"
	RelatorVideographer                        = "vdg"
	RelatorVocalist                            = "-voc" // Discontinued
	RelatorVoiceActor                          = "vac"
	RelatorWitness                             = "wit"
	RelatorWoodcutter                          = "wdc"
	RelatorWoodEngraver                        = "wde"
	RelatorWriterOfAccompanyingMaterial        = "wam"
	RelatorWriterOfAddedCommentary             = "wac"
	RelatorWriterOfAddedLyrics                 = "wal"
	RelatorWriterOfAddedText                   = "wat"
	RelatorWriterOfIntroduction                = "win"
	RelatorWriterOfPreface                     = "wpr"
	RelatorWriterOfSupplementaryTextualContent = "wst"
)
