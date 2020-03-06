package go_aspace

import (
	"time"
)

type Date struct {
	LockVersion    int       `json:"lock_version"`
	Begin          string    `json:"begin"`
	End            string    `json:"end"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	CreateTime     time.Time `json:"create_time"`
	SystemMtime    time.Time `json:"system_mtime"`
	UserMtime      time.Time `json:"user_mtime"`
	DateType       string    `json:"date_type"`
	Label          string    `json:"label"`
	JSONModelType  string    `json:"jsonmodel_type"`
}

type Deaccessions struct {
	LockVersion    int               `json:"lock_version"`
	Description    string            `json:"description"`
	Reason         string            `json:"reason"`
	Disposition    string            `json:"disposition"`
	Notification   bool              `json:"notification"`
	CreatedBy      string            `json:"created_by"`
	LastModifiedBy string            `json:"last_modified_by"`
	CreateTime     time.Time         `json:"create_time"`
	SystemMtime    time.Time         `json:"system_mtime"`
	UserMtime      time.Time         `json:"user_mtime"`
	JSONModelType  string            `json:"jsonmodel_type"`
	Extents        []Extent          `json:"extents"`
	Repository     map[string]string `json:"repository"`
	Date           Date              `json:"date"`
	Notes          []Note            `json:"notes"`
}

type ExternalDocument struct {
	LockVersion    int       `json:"lock_version"`
	Title          string    `json:"title"`
	Location       string    `json:"location"`
	Publish        bool      `json:"publish"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	CreateTime     time.Time `json:"create_time"`
	SystemMtime    time.Time `json:"system_mtime"`
	UserMtime      time.Time `json:"user_mtime"`
	JSONModelType  string    `json:"jsonmodel_type"`
}

type ExternalID struct {
	ExternalID     string    `json:"external_id"`
	Source         string    `json:"source"`
	CreatedBy      string    `json:"created_by"`
	LastModifiedBy string    `json:"last_modified_by"`
	CreateTime     time.Time `json:"create_time"`
	SystemMtime    time.Time `json:"system_mtime"`
	UserMtime      time.Time `json:"user_mtime"`
	JSONModelType  string    `json:"jsonmodel_type"`
}

type Extent struct {
	LockVersion      int       `json:"lock_version"`
	Number           string    `json:"number"`
	ContainerSummary string    `json:"container_summary"`
	CreatedBy        string    `json:"created_by"`
	LastModifiedBy   string    `json:"last_modified_by"`
	CreateTime       time.Time `json:"create_time"`
	SystemMtime      time.Time `json:"system_mtime"`
	UserMtime        time.Time `json:"user_mtime"`
	Portion          string    `json:"portion"`
	ExtentType       string    `json:"extent_type"`
	JSONModelType    string    `json:"jsonmodel_type"`
}

type Instance struct {
	CreateTime       time.Time     `json:"create_time"`
	CreatedBy        string        `json:"created_by"`
	InstanceType     string        `json:"instance_type"`
	IsRepresentative bool          `json:"is_representative"`
	JSONModelType    string        `json:"jsonmodel_type"`
	LastModifiedBy   string        `json:"last_modified_by"`
	LockVersion      int           `json:"lock_version"`
	SubContainer     Sub_Container `json:"sub_container"`
	SystemMtime      time.Time     `json:"system_mtime"`
	UserMtime        time.Time     `json:"user_mtime"`
}

type LangMaterial struct {
	LockVersion       int                `json:"lock_version"`
	CreateTime        time.Time          `json:"create_time"`
	SystemMtime       time.Time          `json:"system_mtime"`
	UserMtime         time.Time          `json:"user_mtime"`
	JSONModelType     string             `json:"jsonmodel_type"`
	Notes             []NoteLangmaterial `json:"notes"`
	LanguageAndScript *LanguageAndScript `json:"language_and_script,omitempty"`
}

type LanguageAndScript struct {
	LockVersion   int       `json:"lock_version,omitempty"`
	CreateTime    time.Time `json:"create_time,omitempty"`
	SystemMtime   time.Time `json:"system_mtime,omitempty"`
	UserMtime     time.Time `json:"user_mtime,omitempty"`
	Language      string    `json:"language,omitempty"`
	JSONModelType string    `json:"jsonmodel_type,omitempty"`
}

type LinkedAgent struct {
	Title string `json:"title"`
	Role  string `json:"role"`
	Terms []Term `json:"terms"`
	Ref   string `json:"ref"`
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
	Classifications             []map[string]string `json:"classifications,omitempty"`
	CreateTime                  time.Time           `json:"create_time"`
	Created_By                  string              `json:"created_by"`
	Dates                       []Date              `json:"dates"`
	Deaccessions                []Deaccessions      `json:"deaccessions"`
	EADID                       string              `json:"ead_id"`
	EADLocation                 string              `json:"ead_location"`
	Extents                     []Extent            `json:"extents"`
	ExternalIDs                 []ExternalID        `json:"external_ids"`
	FindingAidAuthor            string              `json:"finding_aid_author"`
	FindingAidDate              string              `json:"finding_aid_date"`
	FindingAidDescription_rules string              `json:"finding_aid_description_rules"`
	FindingAidLanguage          string              `json:"finding_aid_language"`
	FindingAidLanguage_Note     string              `json:"finding_aid_language_note"`
	FindingAidScript            string              `json:"finding_aid_script"`
	FindingAidStatus            string              `json:"finding_aid_status"`
	FindingAidTitle             string              `json:"finding_aid_title"`
	ID_0                        string              `json:"id_0"`
	ID_1                        string              `json:"id_1"`
	ID_2                        string              `json:"id_2"`
	ID_3                        string              `json:"id_3"`
	Instances                   []Instance          `json:"instances"`
	IsSlugAuto                  bool                `json:"is_slug_auto"`
	JSONModelType               string              `json:"jsonmodel_type"`
	LangMaterials               []LangMaterial      `json:"lang_materials"`
	LastModifiedBy              string              `json:"last_modified_by"`
	Level                       string              `json:"level"`
	LinkedAgents                []LinkedAgent       `json:"linked_agents"`
	LinkedEvents                []map[string]string `json:"linked_events"`
	LockVersion                 int                 `json:"lock_version"`
	Notes                       []Note              `json:"notes"`
	Publish                     bool                `json:"publish"`
	RelatedAccessions           []map[string]string `json:"related_accessions"`
	Repository                  map[string]string   `json:"repository"`
	RepositoryProcessingNote    string              `json:"repository_processing_note"`
	Restrictions                bool                `json:"restrictions"`
	RevisionStatements          []RevisionStatement `json:"revision_statements"`
	RightsStatements            []Rights_Statement  `json:"rights_statements"`
	Subjects                    []map[string]string `json:"subjects"`
	Supressed                   bool                `json:"supressed"`
	SystemMtime                 time.Time           `json:"system_mtime"`
	Title                       string              `json:"title"`
	Tree                        map[string]string   `json:"tree"`
	URI                         string              `json:"uri"`
	UserMtime                   time.Time           `json:"user_mtime"`
}

type RevisionStatement struct {
	Date           string            `json:"date"`
	Description    string            `json:"description"`
	Created_By     string            `json:"created_by"`
	LastModifiedBy string            `json:"last_modified_by"`
	CreateTime     time.Time         `json:"create_time"`
	SystemMtime    time.Time         `json:"system_mtime"`
	Publish        bool              `json:"publish"`
	JSONModel_type string            `json:"jsonmodel_type"`
	URI            string            `json:"uri"`
	Repository     map[string]string `json:"repository"`
}

type Rights_Statement struct {
	Lock_Version      int                   `json:"lock_version"`
	Identifier        string                `json:"identifier"`
	CreatedBy         string                `json:"created_by"`
	LastModifiedBy    string                `json:"last_modified_by"`
	CreateTime        time.Time             `json:"create_time"`
	SystemMtime       time.Time             `json:"system_mtime"`
	UserMtime         time.Time             `json:"user_mtime"`
	LicenseTerms      string                `json:"license_terms"`
	RightsType        string                `json:"rights_type"`
	JSONModelType     string                `json:"jsonmodel_type"`
	ExternalDocuments []ExternalDocument    `json:"external_documents"`
	Acts              []RightsStatementsAct `json:"acts"`
	LinkedAgents      []LinkedAgent         `json:"linked_agents"`
	Notes             []NoteRightsStatement `json:"notes"`
}

type RightsStatementsAct struct {
	StartDate      string                   `json:"start_date"`
	EndDate        string                   `json:"etart_date"`
	CreatedBy      string                   `json:"created_by"`
	LastModifiedBy string                   `json:"last_modified_by"`
	CreateTime     time.Time                `json:"create_time"`
	SystemMtime    time.Time                `json:"system_mtime"`
	UserMtime      time.Time                `json:"user_mtime"`
	ActType        string                   `json:"act_type"`
	Restriction    string                   `json:"restriction"`
	JSONModelType  string                   `json:"json_model_type"`
	Notes          []NoteRightsStatementAct `json:"notes"`
}

type Sub_Container struct {
	CreateTime      time.Time         `json:"create_time"`
	CreatedBy       string            `json:"created_by"`
	JSONModel       string            `json:"jsonmodel_type"`
	LastModified_By string            `json:"last_modified_by"`
	LockVersion     int               `json:"lock_version"`
	SystemMtime     time.Time         `json:"system_mtime"`
	Top_Container   map[string]string `json:"top_container"`
	UserMtime       time.Time         `json:"user_mtime"`
}

type Term struct {
	ID                int
	LockVersion       int       `json:"lock_version"`
	JSONSchemaVersion int       `json:"json_schema_version"`
	VocabID           int       `json:"vocab_id"`
	Term              string    `json:"term"`
	TermTypeID        int       `json:"term_type_id"`
	CreatedBy         string    `json:"created_by"`
	LastModifiedBy    string    `json:"last_modified_by"`
	CreateTime        time.Time `json:"create_time"`
	SystemMtime       time.Time `json:"system_mtime"`
	UserMtime         time.Time `json:"user_mtime"`
	XForeignKeyX      int64     `json:"x_foreign_key_x"`
	TermType          string    `json:"term_type"`
	JSONModelType     string    `json:"json_model_type"`
	URI               string    `json:"uri"`
	Vocabulary        string    `json:"vocabulary"`
}
