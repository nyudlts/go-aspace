package aspace

type AdvancedSearch struct {
	Field 	string `json:"field"`
	Value	string `json:"value"`
	JSONModelType	string `json:"jsonmodel_type"`
	Negated	bool 	`json:"negated"`
	Literal bool `json:"literal"`
}

type Ancestor struct {
	Ref   string `json:"ref"`
	Level string `json:"level"`
}

type ArchivalObject struct {
	LockVersion             int                 `json:"lock_version"`
	Position                int                 `json:"position"`
	Publish                 bool                `json:"publish"`
	RefID                   string              `json:"ref_id"`
	ComponentId             string              `json:"component_id"`
	Title                   string              `json:"title"`
	DisplayString           string              `json:"display_string"`
	RestrictionsApply       bool                `json:"restrictions_apply"`
	Supressed               bool                `json:"supressed"`
	IsSlugAuto              bool                `json:"is_slug_auto"`
	Level                   string              `json:"level"`
	JSONModelType           string              `json:"jsonmodel_type"`
	ExternalIDs             []ExternalID        `json:"external_ids"`
	Subjects                []map[string]string `json:"subjects"`
	LinkedEvents            []map[string]string `json:"linked_events"`
	Extents                 []Extent            `json:"extents"`
	LangMaterials           []*LangMaterial     `json:"lang_materials"`
	Dates                   []Date              `json:"dates"`
	ExternalDocuments       []*ExternalDocument `json:"external_documents"`
	RightsStatememts        []*Rights_Statement `json:"rights_statememts"`
	LinkedAgents            []*LinkedAgent      `json:"linked_agents"`
	Ancestors               []Ancestor          `json:"ancestors"`
	Instances               []*Instance         `json:"instances"`
	Notes                   []*Note             `json:"notes"`
	URI                     string              `json:"uri"`
	Repository              map[string]string   `json:"repository"`
	Parent                  map[string]string   `json:"parent"`
	HasUnpublishedAncestors bool                `json:"has_unpublished_ancestors"`
}

type Classification struct {
	HasClassificationTerms bool   `json:"has_classification_terms,omitempty"`
	Slug                   string `json:"slug,omitempty"`
	IsSlugAuto             bool   `json:"is_slug_auto,omitempty"`
	JSONModelType          string `json:"jsonmodel_type"`
}

type Date struct {
	DateType      string `json:"date_type"`
	Label         string `json:"label"`
	Certainty     string `json:"certainty,omitempty"`
	Expression    string `json:"expression,omitempty"`
	Begin         string `json:"begin,omitempty"`
	End           string `json:"end,omitempty"`
	Era           string `json:"era,omitempty"`
	Calendar      string `json:"calendar,omitempty"`
	JSONModelType string `json:"jsonmodel_type"`
}

type Deaccession struct {
	Scope         string   `json:"scope"`
	Description   string   `json:"description,omitempty"`
	Date          Date     `json:"date"`
	Reason        string   `json:"reason,omitempty"`
	Disposition   string   `json:"disposition,omitempty"`
	Notification  bool     `json:"notification,omitempty"`
	Extents       []Extent `json:"extents,omitempty"`
	JSONModelType string   `json:"jsonmodel_type"`
}

type DigitalObject struct {
	LockVersion int `json:"lock_version"`
	DigitalObjectID	string 	`json:"digital_object_id"`
	Title 			string 	`json:"title"`
	Publish 		bool 	`json:"publish"`
	Restrictions	bool 	`json:"restrictions"`
	Supressed		bool 	`json:"supressed"`
	IsSlugAuto 		bool 	`json:"is_slug_auto"`
	JSONModelType 	string 	`json:"jsonmodel_type"`
	ExternalIds		[]ExternalID 		`json:"external_ids"`
	Subjects        []map[string]string `json:"subjects"`
	LinkedEvents    []map[string]string `json:"linked_events"`
	Extents         []Extent          	`json:"extents"`
	LangMaterials   []*LangMaterial		`json:"lang_materials"`
	Dates        	[]*Date           	`json:"dates"`
	ExternalDocuments	[]*ExternalDocument `json:"external_documents"`
	RightsStatememts        []*Rights_Statement `json:"rights_statememts"`
	LinkedAgents            []*LinkedAgent      `json:"linked_agents"`
	FileVersions	[]*FileVersion `json:"file_versions"`
	Notes                      []*Note              `json:"notes"`
	LinkedInstances []interface{} `json:"linked_instances"`
	URI string `json:"uri"`
	Repository map[string]string `json:"repository"`
	Tree map[string]string `json:"tree"`
}

type ExternalDocument struct {
	Title         string `json:"title"`
	Location      string `json:"location"`
	Publish       bool   `json:"publish,omitempty"`
	JSONModelType string `json:"jsonmodel_type"`
}

type ExternalID struct {
	ExternalID    string `json:"external_id"`
	Source        string `json:"source"`
	JSONModelType string `json:"jsonmodel_type"`
}

type Extent struct {
	Portion          string `json:"portion,omitempty"`
	Number           string `json:"number,omitempty"`
	ExtentType       string `json:"extent_type,omitempty"`
	ContainerSummary string `json:"container_summary"`
	PhysicalDetails  string `json:"physical_details"`
	Dimensions       string `json:"dimensions"`
	JSONModelType    string `json:"jsonmodel_type"`
}

type FileVersion struct {
	Identifier string `json:"identifier"`
	Lock_Version int `json:"lock_version"`
	FileURI string `json:"file_uri"`
	Publish *bool `json:"publish,omitempty"`
	FileFormatVersion *string `json:"file_format_version"`
	FileSizeBytes *uint64 `json:"file_size_bytes"`
	Checksum *string `json:"checksum"`
	ChecksumMethod *string `json:"checksum_method"`
	Caption *string `json:"caption"`
	UseStatement string `json:"use_statement"`
	XLink_Actuate_Attribute string `json:"xlink_actuate_attribute"`
	XLink_Show_Attribute string `json:"xlink_show_attribute"`
	File_Format_Name *string `json:"file_format_name"`
	JSONModelType    string `json:"jsonmodel_type"`
	IsRepresentative bool              `json:"is_representative,omitempty"`
	
}

type Instance struct {
	InstanceType     string            `json:"instance_type"`
	SubContainer     Sub_Container     `json:"sub_container,omitempty"`
	DigitalObjects   map[string]string `json:"digital_object,omitempty"`
	IsRepresentative bool              `json:"is_representative,omitempty"`
	JSONModelType    string            `json:"jsonmodel_type"`
}

type LangMaterial struct {
	LanguageAndScript *LanguageAndScript  `json:"language_and_script,omitempty"`
	Notes             []*NoteLangmaterial `json:"notes,omitempty"`
	JSONModelType     string              `json:"jsonmodel_type"`
}

type LanguageAndScript struct {
	Language      string `json:"language,omitempty"`
	JSONModelType string `json:"jsonmodel_type,omitempty"`
}

type LinkedAgent struct {
	Title         string `json:"title"`
	Role          string `json:"role"`
	Terms         []Term `json:"terms"`
	Ref           string `json:"ref"`
	JSONModelType string `json:"jsonmodel_type"`
}

type Note struct {
	JSONModelType string     `json:"jsonmodel_type"`
	PersistentID  string     `json:"persistent_id"`
	Label         string     `json:"label"`
	Type          string     `json:"type"`
	Subnotes      []NoteText `json:"subnotes,omitempty"`
	Content       []string   `json:"content,omitempty"`
}

type NoteLangmaterial struct {
	JSONModelType string   `json:"jsonmodel_type"`
	PersistantID  string   `json:"persistant_id"`
	Label         string   `json:"label"`
	Type          string   `json:"type"`
	Content       []string `json:"content"`
	Publish       bool     `json:"publish"`
}

type NoteRightsStatement struct {
	JSONModelType string   `json:"jsonmodel_type"`
	PersistantID  string   `json:"persistant_id"`
	Type          string   `json:"type"`
	Content       []string `json:"content"`
	Publish       bool     `json:"publish"`
}

type NoteRightsStatementAct struct {
	JSONModelType string   `json:"jsonmodel_type"`
	PersistantID  string   `json:"persistant_id"`
	Type          string   `json:"type"`
	Content       []string `json:"content"`
	Publish       bool     `json:"publish"`
}

type NoteText struct {
	JSONModelType string `json:"jsonmodel_type"`
	Content       string `json:"content"`
	Publish       bool   `json:"publish"`
}

type Resource struct {
	Classifications            []*Classification    `json:"classifications,omitempty"`
	Dates                      []*Date              `json:"dates"`
	Deaccessions               []*Deaccession       `json:"deaccessions"`
	EADID                      string               `json:"ead_id"`
	EADLocation                string               `json:"ead_location"`
	Extents                    []Extent             `json:"extents"`
	ExternalIDs                []ExternalID         `json:"external_ids"`
	FindingAidAuthor           string               `json:"finding_aid_author"`
	FindingAidDate             string               `json:"finding_aid_date"`
	FindingAidDescriptionRules string               `json:"finding_aid_description_rules"`
	FindingAidLanguage         string               `json:"finding_aid_language"`
	FindingAidLanguageNote     string               `json:"finding_aid_language_note"`
	FindingAidScript           string               `json:"finding_aid_script"`
	FindingAidStatus           string               `json:"finding_aid_status"`
	FindingAidTitle            string               `json:"finding_aid_title"`
	FindingAidEditionStatement string               `json:"finding_aid_edition_statement"`
	ID0                        string               `json:"id_0"`
	ID1                        string               `json:"id_1"`
	ID2                        string               `json:"id_2"`
	ID3                        string               `json:"id_3"`
	Instances                  []*Instance          `json:"instances"`
	IsSlugAuto                 bool                 `json:"is_slug_auto"`
	JSONModelType              string               `json:"jsonmodel_type"`
	LangMaterials              []*LangMaterial      `json:"lang_materials"`
	Level                      string               `json:"level"`
	LinkedAgents               []*LinkedAgent       `json:"linked_agents"`
	LinkedEvents               []map[string]string  `json:"linked_events"`
	LockVersion                int                  `json:"lock_version"`
	Notes                      []*Note              `json:"notes"`
	Publish                    bool                 `json:"publish"`
	RelatedAccessions          []map[string]string  `json:"related_accessions"`
	Repository                 map[string]string    `json:"repository"`
	RepositoryProcessingNote   string               `json:"repository_processing_note"`
	Restrictions               bool                 `json:"restrictions"`
	RevisionStatements         []*RevisionStatement `json:"revision_statements"`
	RightsStatements           []*Rights_Statement  `json:"rights_statements"`
	Subjects                   []map[string]string  `json:"subjects"`
	Supressed                  bool                 `json:"supressed"`
	Title                      string               `json:"title"`
	Tree                       map[string]string    `json:"tree"`
	URI                        string               `json:"uri"`
}

type ResourceTree struct {
	Title         string         `json:"title"`
	Id            int            `json:"id"`
	NodeType      string         `json:"node_type"`
	Publish       bool           `json:"publish"`
	Supressed     bool           `json:"supressed"`
	HasChildren   bool           `json:"has_children"`
	Children      []ResourceTree `json:"children"`
	RecordURI     string         `json:"record_uri"`
	Level         string         `json:"level"`
	JSONModelType string         `json:"jsonmodel_type"`
	InstanceTypes []string       `json:"instance_types"`
	Containers    []string       `json:"containers"`
}

type RevisionStatement struct {
	Date         string            `json:"date"`
	Description  string            `json:"description"`
	Publish      bool              `json:"publish"`
	JSONModeType string            `json:"jsonmodel_type"`
	URI          string            `json:"uri"`
	Repository   map[string]string `json:"repository"`
}

type Rights_Statement struct {
	RightsType        string `json:"rights_type"`
	Identifier        string `json:"identifier"`
	Status            string `json:"status,omitempty"`
	DeterminationDate Date   `json:"determination_date,omitempty"`
	StartDate         string `json:"start_date,omitempty"`
	EndDate           string `json:"end_date,omitempty"`
	LicenseTerms      string `json:"license_terms,omitempty"`
	StatuteCitation   string `json:"statute_citation,omitempty"`
	Jurisdiction      string `json:"jurisdiction,omitempty"`
	OtherRightsBasis  string `json:"other_rights_basis,omitempty"`

	JSONModelType     string                 `json:"jsonmodel_type"`
	ExternalDocuments []*ExternalDocument    `json:"external_documents"`
	Acts              []*RightsStatementsAct `json:"acts"`
	LinkedAgents      []*LinkedAgent         `json:"linked_agents"`
	Notes             []*NoteRightsStatement `json:"notes"`
}

type RightsStatementsAct struct {
	StartDate     string                    `json:"start_date"`
	EndDate       string                    `json:"etart_date"`
	ActType       string                    `json:"act_type"`
	Restriction   string                    `json:"restriction"`
	JSONModelType string                    `json:"json_model_type"`
	Notes         []*NoteRightsStatementAct `json:"notes"`
}

type SearchResult struct {
	//{"page_size":10,"first_page":1,"last_page":2,"this_page":1,"offset_first":1,"offset_last":10,"total_hits":20,"results":
	PageSize	int `json:"page_size"`
	FirstPage	int `json:"first_page"`
	LastPage	int `json:"last_page"`
	ThisPage 	int `json:"this_page"`
	OffsetFirst	int `json:"offset_first"`
	OffsetLast	int `json:"offset_last"`
	TotalHits	int `json:"total_hits"`
	Results 	[]map[string]interface{} `json:"results"`
}

type Sub_Container struct {
	JSONModel    string            `json:"jsonmodel_type"`
	TopContainer map[string]interface{} `json:"top_container"`
}

type Term struct {
	ID                int
	JSONSchemaVersion int    `json:"json_schema_version"`
	VocabID           int    `json:"vocab_id"`
	Term              string `json:"term"`
	TermTypeID        int    `json:"term_type_id"`
	XForeignKeyX      int64  `json:"x_foreign_key_x"`
	TermType          string `json:"term_type"`
	JSONModelType     string `json:"json_model_type"`
	URI               string `json:"uri"`
	Vocabulary        string `json:"vocabulary"`
}
