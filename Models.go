package aspace

type Accession struct {
	URI                    string                         `json:"uri,omitempty"`
	ExternalIDs            []ExternalID                   `json:"external_ids,omitempty"`
	Title                  string                         `json:"title,omitempty"`
	DisplayString          string                         `json:"display_string,omitempty"`
	Slug                   string                         `json:"slug,omitempty"`
	IsSlugAuto             bool                           `json:"is_slug_auto,omitempty"`
	ID0                    string                         `json:"id_0,omitempty"`
	ID1                    string                         `json:"id_1,omitempty"`
	ID2                    string                         `json:"id_2,omitempty"`
	ID3                    string                         `json:"id_3,omitempty"`
	ContentDescription     string                         `json:"content_description,omitempty"`
	ConditionDescription   string                         `json:"condition_description,omitempty"`
	Disposition            string                         `json:"disposition,omitempty"`
	Inventory              string                         `json:"inventory,omitempty"`
	Provenance             string                         `json:"provenance,omitempty"`
	RelatedAccessions      []AccessionSiblingRelationship `json:"related_accessions"`
	AccessionDate          string                         `json:"accession_date,omitempty"`
	Publish                bool                           `json:"publish"`
	Classifications        []map[string]string            `json:"classifications,omitempty"`
	Subjects               []Subject                      `json:"subjects,omitempty"`
	LinkedEvents           []LinkedEvent                  `json:"linked_events,omitempty"`
	Extents                []Extent                       `json:"extents,omitempty"`
	Dates                  []Date                         `json:"dates,omitempty"`
	ExternalDocuments      []ExternalDocument             `json:"external_documents"`
	RightsStatements       []RightsStatement              `json:"rights_statements"`
	CollectionManagement   CollectionManagement           `json:"collection_management,omitempty"`
	UserDefined            UserDefined                    `json:"user_defined,omitempty"`
	RelatedResources       []map[string]string            `json:"related_resources,omitempty"`
	Suppressed             bool                           `json:"suppressed"`
	AcquisitionType        string                         `json:"acquisition_type,omitempty"`
	ResourceType           string                         `json:"resource_type,omitempty"`
	RestrictionsApply      bool                           `json:"restrictions_apply,omitempty"`
	RetentionRule          string                         `json:"retention_rule,omitempty"`
	GeneralNote            string                         `json:"general_note,omitempty"`
	AccessRestrictions     bool                           `json:"access_restrictions,omitempty"`
	AccessRestrictionsNote string                         `json:"access_restrictions_note,omitempty"`
	UseRestrictions        bool                           `json:"use_restrictions,omitempty"`
	UseRestrictionsNote    string                         `json:"use_restrictions_note,omitempty"`
	LinkedAgents           []LinkedAgent                  `json:"linked_agents,omitempty"`
	Instances              []Instance                     `json:"instances,omitempty"`
	LockVersion            int                            `json:"lock_version"`
	JSONModelType          string                         `json:"json_model_type"`
	Repository             LinkedRepository               `json:"repository"`
	Parent                 map[string]string              `json:"parent,omitempty"`
}

type APIResponse struct {
	Status      string   `json:"status"`
	ID          int      `json:"id"`
	LockVersion int      `json:"lock_version"`
	Stale       bool     `json:"stale"`
	URI         string   `json:"uri"`
	Warnings    []string `json:"warnings,omitempty"`
}

type AccessionSiblingRelationship struct {
	JSONModelType string `json:"jsonmodel_type"`
	Relator       string `json:"relator"`
	RelatorType   string `json:"relator_type"`
	Ref           string `json:"ref"`
}

type AdvancedSearch struct {
	Field         string `json:"field,omitempty"`
	Value         string `json:"value,omitempty"`
	JSONModelType string `json:"jsonmodel_type,omitempty"`
	Negated       bool   `json:"negated,omitempty"`
	Literal       bool   `json:"literal,omitempty"`
}

type Agent struct {
	LockVersion                     int                `json:"lock_version"`
	Publish                         bool               `json:"publish,omitempty"`
	IsSlugAuto                      bool               `json:"is_slug_auto,omitempty"`
	JSONModelType                   string             `json:"jsonmodel_type"`
	AgentContacts                   []AgentContact     `json:"agent_contacts,omitempty"`
	LinkedAgentRoles                []string           `json:"linked_agent_roles,omitempty"`
	ExternalDocuments               []ExternalDocument `json:"external_documents,omitempty"`
	Notes                           []Note             `json:"notes,omitempty"`
	UsedWithinRepositories          []string           `json:"used_within_repositories,omitempty"`
	UsedWithinPublishedRepositories []string           `json:"used_within_published_repositories,omitempty"`
	DatesOfExistence                []interface{}      `json:"dates_of_existence,omitempty"`
	Names                           []Name             `json:"names,omitempty"`
	RelatedAgents                   []RelatedAgent     `json:"related_agents,omitempty"`
	URI                             string             `json:"uri"`
	AgentType                       string             `json:"agent_type"`
	IsLinkedToPublishedRecord       bool               `json:"is_linked_to_published_record"`
	DisplayName                     Name               `json:"display_name"`
	Title                           string             `json:"title"`
	IsUser                          string             `json:"is_user"`
}

type AgentReference struct {
	Ref   string `json:"ref,omitempty"`
	Agent Agent  `json:"agent,omitempty"`
}

type AgentContact struct {
	LockVersion   int           `json:"lock_version,omitempty"`
	Name          string        `json:"name,omitempty"`
	Address1      string        `json:"address_1,omitempty"`
	Address2      string        `json:"address_2,omitempty"`
	Address3      string        `json:"address_3,omitempty"`
	City          string        `json:"city,omitempty"`
	Region        string        `json:"region,omitempty"`
	Country       string        `json:"country,omitempty"`
	PostCode      string        `json:"post_code,omitempty"`
	Email         string        `json:"email,omitempty"`
	Note          string        `json:"note,omitempty"`
	Salutation    string        `json:"salutation,omitempty"`
	JSONModelType string        `json:"jsonmodel_type,omitempty"`
	Telephones    []interface{} `json:"telephones,omitempty"`
}

type Ancestor struct {
	Ref   string `json:"ref,omitempty"`
	Level string `json:"level,omitempty"`
}

type ArchivalObject struct {
	LockVersion             int                 `json:"lock_version"`
	Position                int                 `json:"position,omitempty"`
	Publish                 bool                `json:"publish,omitempty"`
	RefID                   string              `json:"ref_id,omitempty"`
	ComponentId             string              `json:"component_id,omitempty"`
	Title                   string              `json:"title,omitempty"`
	DisplayString           string              `json:"display_string,omitempty"`
	RestrictionsApply       bool                `json:"restrictions_apply,omitempty"`
	Suppressed              bool                `json:"suppressed,omitempty"`
	IsSlugAuto              bool                `json:"is_slug_auto,omitempty"`
	Level                   string              `json:"level,omitempty"`
	JSONModelType           string              `json:"jsonmodel_type,omitempty"`
	ExternalIDs             []ExternalID        `json:"external_ids,omitempty"`
	Subjects                []Subject           `json:"subjects,omitempty"`
	LinkedEvents            []map[string]string `json:"linked_events,omitempty"`
	Extents                 []Extent            `json:"extents,omitempty"`
	LangMaterials           []LangMaterial      `json:"lang_materials,omitempty"`
	Dates                   []Date              `json:"dates,omitempty"`
	ExternalDocuments       []ExternalDocument  `json:"external_documents,omitempty"`
	RightsStatements        []RightsStatement   `json:"rights_statements,omitempty"`
	LinkedAgents            []LinkedAgent       `json:"linked_agents,omitempty"`
	Ancestors               []Ancestor          `json:"ancestors,omitempty"`
	Instances               []Instance          `json:"instances,omitempty"`
	Notes                   []Note              `json:"notes,omitempty"`
	URI                     string              `json:"uri,omitempty"`
	Repository              Repository          `json:"repository,omitempty"`
	Parent                  map[string]string   `json:"parent,omitempty"`
	HasUnpublishedAncestors bool                `json:"has_unpublished_ancestors,omitempty"`
	Resource                map[string]string   `json:"resource"`
	RepresentativeImage     FileVersion         `json:"representative_image,omitempty"`
	ArkName                 interface{}         `json:"ark_name,omitempty"`
}

type Classification struct {
	HasClassificationTerms bool   `json:"has_classification_terms,omitempty"`
	Slug                   string `json:"slug,omitempty"`
	IsSlugAuto             bool   `json:"is_slug_auto,omitempty"`
	JSONModelType          string `json:"jsonmodel_type,omitempty"`
}

type CollectionManagement struct {
	LockVersion                    int               `json:"lock_version,omitempty"`
	JSONModelType                  string            `json:"jsonmodel_type,omitempty"`
	URI                            string            `json:"uri,omitempty"`
	ProcessingHoursPerFootEstimate string            `json:"processing_hours_per_foot_estimate,omitempty"`
	ProcessingTotalExtent          string            `json:"processing_total_extent,omitempty"`
	ProcessingTotalExtentType      string            `json:"processing_total_extent_type,omitempty"`
	ProcessingHoursTotal           string            `json:"processing_hours_total,omitempty"`
	ProcessingPlan                 string            `json:"processing_plan,omitempty"`
	ProcessingPriority             string            `json:"processing_priority,omitempty"`
	ProcessingFundingSource        string            `json:"processing_funding_source,omitempty"`
	Processors                     string            `json:"processors,omitempty"`
	RightsDetermined               bool              `json:"rights_determined,omitempty"`
	ProcessingStatus               string            `json:"processing_status,omitempty"`
	Repository                     map[string]string `json:"repository,omitempty"`
	Parent                         map[string]string `json:"parent,omitempty"`
	ExternalIDs                    []ExternalID      `json:"external_ids,omitempty"`
}

type Date struct {
	DateType      string `json:"date_type,omitempty"`
	Label         string `json:"label,omitempty"`
	Certainty     string `json:"certainty,omitempty"`
	Expression    string `json:"expression,omitempty"`
	Begin         string `json:"begin,omitempty"`
	End           string `json:"end,omitempty"`
	Era           string `json:"era,omitempty"`
	Calendar      string `json:"calendar,omitempty"`
	JSONModelType string `json:"jsonmodel_type,omitempty"`
}

type Deaccession struct {
	Scope         string   `json:"scope,omitempty"`
	Description   string   `json:"description,omitempty"`
	Date          Date     `json:"date,omitempty"`
	Reason        string   `json:"reason,omitempty"`
	Disposition   string   `json:"disposition,omitempty"`
	Notification  bool     `json:"notification,omitempty"`
	Extents       []Extent `json:"extents,omitempty"`
	JSONModelType string   `json:"jsonmodel_type,omitempty"`
}

type DigitalObject struct {
	LockVersion          int                    `json:"lock_version"`
	DigitalObjectID      string                 `json:"digital_object_id,omitempty"`
	Title                string                 `json:"title,omitempty"`
	Publish              bool                   `json:"publish"`
	Restrictions         bool                   `json:"restrictions"`
	Suppressed           bool                   `json:"suppressed"`
	IsSlugAuto           bool                   `json:"is_slug_auto"`
	JSONModelType        string                 `json:"jsonmodel_type,omitempty"`
	ExternalIds          []ExternalID           `json:"external_ids,omitempty"`
	Subjects             []SubjectReference     `json:"subjects,omitempty"`
	LinkedEvents         []LinkedEvent          `json:"linked_events,omitempty"`
	Extents              []Extent               `json:"extents,omitempty"`
	LangMaterials        []LangMaterial         `json:"lang_materials,omitempty"`
	Dates                []Date                 `json:"dates,omitempty"`
	ExternalDocuments    []ExternalDocument     `json:"external_documents,omitempty"`
	RightsStatements     []RightsStatement      `json:"rights_statements,omitempty"`
	LinkedAgents         []LinkedAgent          `json:"linked_agents,omitempty"`
	FileVersions         []FileVersion          `json:"file_versions,omitempty"`
	Notes                []interface{}          `json:"notes,omitempty"`
	LinkedInstances      []LinkedInstance       `json:"linked_instances,omitempty"`
	URI                  string                 `json:"uri,omitempty"`
	Repository           LinkedRepository       `json:"repository,omitempty"`
	Tree                 map[string]string      `json:"tree,omitempty"`
	Level                string                 `json:"level,omitempty"`
	Slug                 string                 `json:"slug,omitempty"`
	DigitalObjectType    string                 `json:"digital_object_type,omitempty"`
	UserDefined          []UserDefined          `json:"user_defined,omitempty"`
	CollectionManagement []CollectionManagement `json:"collection_management,omitempty"`
}

type ExternalDocument struct {
	Title         string `json:"title,omitempty"`
	Location      string `json:"location,omitempty"`
	Publish       bool   `json:"publish,omitempty"`
	JSONModelType string `json:"jsonmodel_type,omitempty"`
}

type ExternalID struct {
	ExternalID    string `json:"external_id,omitempty"`
	Source        string `json:"source,omitempty"`
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

type Event struct {
	URI               string             `json:"uri,omitempty"`
	Refid             string             `json:"refid,omitempty"`
	ExternalIDs       []ExternalID       `json:"external_ids,omitempty"`
	ExternalDocuments []ExternalDocument `json:"external_documents,omitempty"`
	EventType         string             `json:"event_type"`
	Date              Date               `json:"date"`
	Timestamp         string             `json:"timestamp,omitempty"`
	Outcome           string             `json:"outcome,omitempty"`
	OutcomeNote       string             `json:"outcome_note,omitempty"`
	Suppressed        bool               `json:"suppressed,omitempty"`
	LinkedAgents      []LinkedAgent      `json:"linked_agents"`
	LinkedRecords     []LinkedRecord     `json:"linked_records"`
	LockVersion       int                `json:"lock_version"`
	JSONModelType     string             `json:"jsonmodel_type"`
	Repository        LinkedRepository   `json:"repository"`
}

type FileVersion struct {
	Identifier            string `json:"identifier,omitempty"`
	LockVersion           int    `json:"lock_version,omitempty"`
	FileURI               string `json:"file_uri,omitempty"`
	Publish               bool   `json:"publish"`
	FileFormatVersion     string `json:"file_format_version,omitempty"`
	FileSizeBytes         uint64 `json:"file_size_bytes,omitempty"`
	Checksum              string `json:"checksum,omitempty"`
	ChecksumMethod        string `json:"checksum_method,omitempty"`
	Caption               string `json:"caption,omitempty"`
	UseStatement          string `json:"use_statement,omitempty"`
	XLinkActuateAttribute string `json:"xlink_actuate_attribute,omitempty"`
	XLinkShowAttribute    string `json:"xlink_show_attribute,omitempty"`
	FileFormatName        string `json:"file_format_name,omitempty"`
	JSONModelType         string `json:"jsonmodel_type,omitempty"`
	IsRepresentative      bool   `json:"is_representative"`
}

type Instance struct {
	InstanceType     string            `json:"instance_type,omitempty"`
	SubContainer     SubContainer      `json:"sub_container,omitempty"`
	DigitalObject    map[string]string `json:"digital_object,omitempty"`
	IsRepresentative bool              `json:"is_representative,omitempty"`
	JSONModelType    string            `json:"jsonmodel_type,omitempty"`
}

type Inherited struct {
	Ref    string `json:"ref,omitempty"`
	Level  string `json:"level,omitempty"`
	Direct string `json:"direct,omitempty"`
}

type LangMaterial struct {
	LanguageAndScript LanguageAndScript  `json:"language_and_script,omitempty"`
	Notes             []NoteLangMaterial `json:"notes,omitempty"`
	JSONModelType     string             `json:"jsonmodel_type,omitempty"`
}

type LanguageAndScript struct {
	Language      string `json:"language,omitempty"`
	Script        string `json:"script,omitempty"`
	JSONModelType string `json:"jsonmodel_type,omitempty"`
}

type LinkedAgent struct {
	Title         string    `json:"title,omitempty"`
	Role          string    `json:"role,omitempty"`
	Terms         []Term    `json:"terms,omitempty"`
	Ref           string    `json:"ref,omitempty"`
	JSONModelType string    `json:"jsonmodel_type,omitempty"`
	Relator       string    `json:"relator,omitempty"`
	Resolved      Agent     `json:"_resolved,omitempty"`
	Inherited     Inherited `json:"_inherited,omitempty"`
}

type LinkedEvent struct {
	Ref      string `json:"ref,omitempty"`
	Resolved Event  `json:"_resolved,omitempty"`
}

type LinkedInstance struct {
	Ref      string   `json:"ref,omitempty"`
	Resolved Instance `json:"_resolved,omitempty"`
}

type LinkedRecord struct {
	Role     string `json:"role"`
	Ref      string `json:"ref"`
	Resolved Agent  `json:"_resolved,omitempty"`
}

type LinkedRepository struct {
	Ref      string     `json:"ref"`
	Resolved Repository `json:"_resolved,omitempty"`
}

type Name struct {
	LockVersion          int           `json:"lock_version,omitempty"`
	PrimaryName          string        `json:"primary_name,omitempty"`
	Title                string        `json:"title,omitempty"`
	Prefix               string        `json:"prefix,omitempty"`
	RestOfName           string        `json:"rest_of_name,omitempty"`
	Suffix               string        `json:"suffix,omitempty"`
	FullerForm           string        `json:"fuller_form,omitempty"`
	Number               string        `json:"number,omitempty"`
	Dates                string        `json:"dates,omitempty"`
	Qualifier            string        `json:"qualifier,omitempty"`
	SortName             string        `json:"sort_name,omitempty"`
	Authorized           bool          `json:"authorized,omitempty"`
	IsDisplayName        bool          `json:"is_display_name,omitempty"`
	Source               string        `json:"source,omitempty"`
	Rules                string        `json:"rules,omitempty"`
	NameOrder            string        `json:"name_order,omitempty"`
	JSONModelType        string        `json:"jsonmodel_type,omitempty"`
	UseDates             []interface{} `json:"use_dates,omitempty"`
	AuthorityID          string        `json:"authority_id,omitempty"`
	SubordinateName1     string        `json:"subordinate_name_1,omitempty"`
	SubordinateName2     string        `json:"subordinate_name_2,omitempty"`
	SortNameAutoGenerate bool          `json:"sort_name_auto_generate"`
	FamilyName           string        `json:"family_name"`
}

type Note struct {
	JSONModelType     string                 `json:"jsonmodel_type,omitempty"`
	PersistentID      string                 `json:"persistent_id,omitempty"`
	Label             string                 `json:"label,omitempty"`
	Type              string                 `json:"type,omitempty"`
	Subnotes          []interface{}          `json:"subnotes,omitempty"`
	Content           []string               `json:"content,omitempty"`
	Publish           bool                   `json:"publish,omitempty"`
	RightsRestriction map[string]interface{} `json:"rights_restriction,omitempty"`
}

type NoteBibliography struct {
	Label         string           `json:"label,omitempty"`
	Publish       bool             `json:"publish,omitempty"`
	PersistentId  string           `json:"persistent_id,omitempty"`
	IngestProblem string           `json:"ingest_problem,omitempty"`
	LockVersion   *int             `json:"lock_version,omitempty"`
	JSONModelType string           `json:"jsonmodel_type"`
	Repository    LinkedRepository `json:"repository,omitempty"`
	Content       []string         `json:"content,omitempty"`
	Items         []string         `json:"items,omitempty"`
	Inherited     Inherited        `json:"inherited,omitempty"`
}

type NoteDigitalObject struct {
	Label         string           `json:"label,omitempty"`
	Publish       bool             `json:"publish"`
	PersistentId  string           `json:"persistent_id"`
	IngestProblem string           `json:"ingest_problem,omitempty"`
	LockVersion   *int             `json:"lock_version,omitempty"`
	JsonModelType string           `json:"jsonmodel_type"`
	Repository    LinkedRepository `json:"repository,omitempty"`
	Content       []string         `json:"content"`
	Type          string           `json:"type"`
}

type NoteLangMaterial struct {
	JSONModelType string   `json:"jsonmodel_type,omitempty"`
	PersistentID  string   `json:"persistent_id,omitempty"`
	Label         string   `json:"label,omitempty"`
	Type          string   `json:"type,omitempty"`
	Content       []string `json:"content,omitempty"`
	Publish       bool     `json:"publish,omitempty"`
}

type NoteRightsStatement struct {
	JSONModelType string   `json:"jsonmodel_type,omitempty"`
	PersistentID  string   `json:"persistent_id,omitempty"`
	Type          string   `json:"type,omitempty"`
	Content       []string `json:"content,omitempty"`
	Publish       bool     `json:"publish,omitempty"`
}

type NoteRightsStatementAct struct {
	JSONModelType string   `json:"jsonmodel_type,omitempty"`
	PersistentID  string   `json:"persistent_id,omitempty"`
	Type          string   `json:"type,omitempty"`
	Content       []string `json:"content,omitempty"`
	Publish       bool     `json:"publish,omitempty"`
}

type NoteText struct {
	JSONModelType string `json:"jsonmodel_type,omitempty"`
	Content       string `json:"content,omitempty"`
	Publish       bool   `json:"publish,omitempty"`
	Title         string `json:"title,omitempty"`
}

type RelatedAgent struct {
	Relator       string `json:"relator,omitempty"`
	JSONModelType string `json:"jsonmodel_type,omitempty"`
	Description   string `json:"description,omitempty"`
	Dates         Date   `json:"dates,omitempty"`
	Ref           string `json:"ref,omitempty"`
}

type RepositoryReference struct {
	Ref      string       `json:"ref,omitempty"`
	Resolved []Repository `json:"resolved,omitempty"`
}

type Repository struct {
	URI                   string         `json:"uri"`
	RepoCode              string         `json:"repo_code,omitempty"`
	Name                  string         `json:"name,omitempty"`
	OrgCode               string         `json:"org_code,omitempty"`
	Country               string         `json:"country,omitempty"`
	ParentInstitutionName string         `json:"parent_institution_name,omitempty"`
	Description           string         `json:"description,omitempty"`
	URL                   string         `json:",omitempty"`
	ImageURL              string         `json:"image_url,omitempty"`
	ContactPersons        string         `json:"contact_persons,omitempty"`
	Publish               bool           `json:"publish,omitempty"`
	DisplayString         string         `json:"display_string"`
	OAIIsDisabled         bool           `json:"oai_is_disabled,omitempty"`
	OAISetsAvailable      string         `json:"oai_sets_available,omitempty"`
	Slug                  string         `json:"slug,omitempty"`
	IsSlugAuto            bool           `json:"is_slug_auto"`
	AgentRepresentation   AgentReference `json:"agent_representation"`
	LockVersion           int            `json:"lock_version"`
	JSONModelType         string         `json:"jsonmodel_type,omitempty"`
}

type Resource struct {
	Classifications            []Classification    `json:"classifications,omitempty"`
	CollectionManagement       []interface{}       `json:"collection_management,omitempty"`
	Dates                      []Date              `json:"dates,omitempty"`
	Deaccessions               []Deaccession       `json:"deaccessions,omitempty"`
	EADID                      string              `json:"ead_id,omitempty"`
	EADLocation                string              `json:"ead_location,omitempty"`
	Extents                    []Extent            `json:"extents,omitempty"`
	ExternalArkURL             string              `json:"external_ark_url"`
	ExternalIDs                []ExternalID        `json:"external_ids,omitempty"`
	FindingAidAuthor           string              `json:"finding_aid_author,omitempty"`
	FindingAidDate             string              `json:"finding_aid_date,omitempty"`
	FindingAidDescriptionRules string              `json:"finding_aid_description_rules,omitempty"`
	FindingAidLanguage         string              `json:"finding_aid_language,omitempty"`
	FindingAidLanguageNote     string              `json:"finding_aid_language_note,omitempty"`
	FindingAidNote             string              `json:"finding_aid_note"`
	FindingAidScript           string              `json:"finding_aid_script,omitempty"`
	FindingAidSponsor          string              `json:"finding_aid_status,omitempty"`
	FindingAidStatus           string              `json:"finding_aid_sponsor,omitempty"`
	FindingAidTitle            string              `json:"finding_aid_title,omitempty"`
	FindingAidEditionStatement string              `json:"finding_aid_edition_statement,omitempty"`
	FindingAidSeriesStatement  string              `json:"finding_aid_series_statement,omitempty"`
	ID0                        string              `json:"id_0,omitempty"`
	ID1                        string              `json:"id_1,omitempty"`
	ID2                        string              `json:"id_2,omitempty"`
	ID3                        string              `json:"id_3,omitempty"`
	Instances                  []Instance          `json:"instances,omitempty"`
	IsSlugAuto                 bool                `json:"is_slug_auto,omitempty"`
	Json                       []byte              `json:"json,omitempty"`
	JSONModelType              string              `json:"jsonmodel_type,omitempty"`
	LangMaterials              []LangMaterial      `json:"lang_materials,omitempty"`
	Level                      string              `json:"level,omitempty"`
	LinkedAgents               []LinkedAgent       `json:"linked_agents,omitempty"`
	LinkedEvents               []map[string]string `json:"linked_events,omitempty"`
	LockVersion                int                 `json:"lock_version"`
	Notes                      []Note              `json:"notes,omitempty"`
	OtherLevel                 string              `json:"other_level,omitempty"`
	Publish                    bool                `json:"publish,omitempty"`
	RelatedAccessions          []map[string]string `json:"related_accessions,omitempty"`
	Repository                 map[string]string   `json:"repository,omitempty"`
	RepositoryProcessingNote   string              `json:"repository_processing_note,omitempty"`
	Restrictions               bool                `json:"restrictions,omitempty"`
	ResourceType               string              `json:"resource_type,omitempty"`
	RevisionStatements         []RevisionStatement `json:"revision_statements,omitempty"`
	RightsStatements           []RightsStatement   `json:"rights_statements,omitempty"`
	Slug                       string              `json:"slug,omitempty"`
	Subjects                   []map[string]string `json:"subjects,omitempty"`
	Suppressed                 bool                `json:"suppressed,omitempty"`
	Title                      string              `json:"title,omitempty"`
	Tree                       map[string]string   `json:"tree,omitempty"`
	UserDefined                UserDefined         `json:"user_defined,omitempty"`
	URI                        string              `json:"uri,omitempty"`
}

type LinkedTree struct {
	Ref      string       `json:"ref,omitempty"`
	Resolved ResourceTree `json:"_resolved,omitempty"`
}

type PrecomputedWaypoint struct {
	Nodes []Node `json:"0,omitempty"`
}

type ResourceTree struct {
	Title         string         `json:"title,omitempty"`
	Id            int            `json:"id,omitempty"`
	NodeType      string         `json:"node_type,omitempty"`
	Publish       bool           `json:"publish,omitempty"`
	Suppressed    bool           `json:"suppressed,omitempty"`
	HasChildren   bool           `json:"has_children,omitempty"`
	Children      []ResourceTree `json:"children,omitempty"`
	RecordURI     string         `json:"record_uri,omitempty"`
	Level         string         `json:"level,omitempty"`
	JSONModelType string         `json:"jsonmodel_type,omitempty"`
	InstanceTypes []string       `json:"instance_types,omitempty"`
	Containers    []string       `json:"containers,omitempty"`
}

type RevisionStatement struct {
	Date         string            `json:"date,omitempty"`
	Description  string            `json:"description,omitempty"`
	Publish      bool              `json:"publish,omitempty"`
	JSONModeType string            `json:"jsonmodel_type,omitempty"`
	URI          string            `json:"uri,omitempty"`
	Repository   map[string]string `json:"repository,omitempty"`
}

type RightsStatement struct {
	RightsType        string                `json:"rights_type,omitempty"`
	Identifier        string                `json:"identifier,omitempty"`
	Status            string                `json:"status,omitempty"`
	DeterminationDate Date                  `json:"determination_date,omitempty"`
	StartDate         string                `json:"start_date,omitempty"`
	EndDate           string                `json:"end_date,omitempty"`
	LicenseTerms      string                `json:"license_terms,omitempty"`
	StatuteCitation   string                `json:"statute_citation,omitempty"`
	Jurisdiction      string                `json:"jurisdiction,omitempty"`
	OtherRightsBasis  string                `json:"other_rights_basis,omitempty"`
	JSONModelType     string                `json:"jsonmodel_type,omitempty"`
	ExternalDocuments []ExternalDocument    `json:"external_documents,omitempty"`
	Acts              []RightsStatementsAct `json:"acts,omitempty"`
	LinkedAgents      []LinkedAgent         `json:"linked_agents,omitempty"`
	Notes             []NoteRightsStatement `json:"notes,omitempty"`
}

type RightsStatementsAct struct {
	StartDate     string                   `json:"start_date,omitempty"`
	EndDate       string                   `json:"end_date,omitempty"`
	ActType       string                   `json:"act_type,omitempty"`
	Restriction   string                   `json:"restriction,omitempty"`
	JSONModelType string                   `json:"json_model_type,omitempty"`
	Notes         []NoteRightsStatementAct `json:"notes,omitempty"`
}

type SearchResult struct {
	PageSize    int                      `json:"page_size,omitempty"`
	FirstPage   int                      `json:"first_page,omitempty"`
	LastPage    int                      `json:"last_page,omitempty"`
	ThisPage    int                      `json:"this_page,omitempty"`
	OffsetFirst int                      `json:"offset_first,omitempty"`
	OffsetLast  int                      `json:"offset_last,omitempty"`
	TotalHits   int                      `json:"total_hits,omitempty"`
	Results     []map[string]interface{} `json:"results,omitempty"`
}

type SubContainer struct {
	JSONModel     string            `json:"jsonmodel_type,omitempty"`
	TopContainer  map[string]string `json:"top_container,omitempty"`
	Type2         string            `json:"type_2,omitempty"`
	Indicator2    string            `json:"indicator_2,omitempty"`
	Barcode2      string            `json:"barcode_2,omitempty"`
	Type3         string            `json:"type_3,omitempty"`
	Indicator3    string            `json:"indicator_3,omitempty"`
	DisplayString string            `json:"display_string,omitempty"`
	Repository    interface{}       `json:"repository,omitempty"`
}

type Subject struct {
	LockVersion                     int                `json:"lock_version"`
	Title                           string             `json:"title,omitempty"`
	IsSlugAuto                      bool               `json:"is_slug_auto,omitempty"`
	Source                          string             `json:"source,omitempty"`
	JSONModelType                   string             `json:"json_model_type,omitempty"`
	ExternalIDs                     []ExternalID       `json:"external_ids,omitempty"`
	Publish                         bool               `json:"publish,omitempty"`
	UsedWithinRepositories          []interface{}      `json:"used_within_repositories,omitempty"`
	UsedWithinPublishedRepositories []interface{}      `json:"used_within_published_repositories,omitempty"`
	Terms                           []Term             `json:"terms"`
	ExternalDocuments               []ExternalDocument `json:"external_documents,omitempty"`
	URI                             string             `json:"uri,omitempty"`
	Vocabulary                      string             `json:"vocabulary"`
	IsLinkedToPublishedRecord       bool               `json:"is_linked_to_published_record,omitempty"`
}

type SubjectReference struct {
	Ref      string  `json:"ref,omitempty"`
	Resolved Subject `json:"_resolved,omitempty"`
}

type Term struct {
	ID                int
	JSONSchemaVersion int    `json:"json_schema_version,omitempty"`
	VocabID           int    `json:"vocab_id,omitempty"`
	Term              string `json:"term,omitempty"`
	TermTypeID        int    `json:"term_type_id,omitempty"`
	XForeignKeyX      int64  `json:"x_foreign_key_x,omitempty"`
	TermType          string `json:"term_type,omitempty"`
	JSONModelType     string `json:"json_model_type,omitempty"`
	URI               string `json:"uri,omitempty"`
	Vocabulary        string `json:"vocabulary,omitempty"`
}

type TopContainer struct {
	URI                       string            `json:"uri,omitempty"`
	Indicator                 string            `json:"indicator"`
	Type                      string            `json:"type"`
	Barcode                   string            `json:"barcode"`
	DisplayString             string            `json:"display_string"`
	LongDisplayString         string            `json:"long_display_string"`
	SubcontainerBarcodes      string            `json:"subcontainer_barcodes"`
	ILSHoldingID              string            `json:"ils_holding_id,omitempty"`
	ILSItemID                 string            `json:"ils_item_id,omitempty"`
	ExportToILS               string            `json:"export_to_ils,omitempty"`
	Restricted                bool              `json:"restricted"`
	CreatedForCollection      string            `json:"created_for_collection,omitempty"`
	IsLinkedToPublishedRecord bool              `json:"is_linked_to_published_record"`
	ActiveRestrictions        []interface{}     `json:"active_restrictions"`
	ContainerLocations        []interface{}     `json:"container_locations"`
	ContainerProfile          map[string]string `json:"container_profile"`
	Series                    []interface{}     `json:"series"`
	Collection                []interface{}     `json:"collection"`
	Repository                map[string]string `json:"repository"`
	JSONModelType             string            `json:"json_model_type"`
}

type Node struct {
	ChildCount           int                            `json:"child_count,omitempty"`
	Waypoints            int                            `json:"waypoints,omitempty"`
	WaypointSize         int                            `json:"waypoint_size,omitempty"`
	Title                string                         `json:"title,omitempty"`
	URI                  string                         `json:"uri,omitempty"`
	SluggedURL           string                         `json:"slugged_url,omitempty"`
	JSONModelType        string                         `json:"jsonmodel_type,omitempty"`
	ParsedTitle          string                         `json:"parsed_title,omitempty"`
	Suppressed           bool                           `json:"suppresed,omitempty"`
	HasDigitalInstance   bool                           `json:"has_digital_instanceint,omitempty"`
	Level                string                         `json:"level,omitempty"`
	Identifier           string                         `json:"identifier,omitempty"`
	PrecomputedWaypoints map[string]PrecomputedWaypoint `json:"precomputed_waypoints,omitempty"`
}

type UserDefined struct {
	Boolean1      *bool      `json:"boolean_1,omitempty"`
	Boolean2      *bool      `json:"boolean_2,omitempty"`
	Boolean3      *bool      `json:"boolean_3,omitempty"`
	Integer1      string     `json:"integer_1,omitempty"`
	Integer2      string     `json:"integer_2,omitempty"`
	Integer3      string     `json:"integer_3,omitempty"`
	Real1         string     `json:"real_1,omitempty"`
	Real2         string     `json:"real_2,omitempty"`
	Real3         string     `json:"real_3,omitempty"`
	String1       string     `json:"string_1,omitempty"`
	String2       string     `json:"string_2,omitempty"`
	String3       string     `json:"string_3,omitempty"`
	String4       string     `json:"string_4,omitempty"`
	Text1         string     `json:"text_1,omitempty"`
	Text2         string     `json:"text_2,omitempty"`
	Text3         string     `json:"text_3,omitempty"`
	Text4         string     `json:"text_4,omitempty"`
	Text5         string     `json:"text_5,omitempty"`
	Date1         Date       `json:"date_1,omitempty"`
	Date2         Date       `json:"date_2,omitempty"`
	Date3         Date       `json:"date_3,omitempty"`
	Enum1         string     `json:"enum_1,omitempty"`
	Enum2         string     `json:"enum_2,omitempty"`
	Enum3         string     `json:"enum_3,omitempty"`
	Enum4         string     `json:"enum_4,omitempty"`
	LockVersion   int        `json:"lock_version"`
	JSONModelType string     `json:"jsonmodel_type"`
	Repository    Repository `json:"repository"`
}
