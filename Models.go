package go_aspace

import (
	"time"
)

type Date struct {
	LockVersion      int       `json:"lock_version"`
	Begin            string    `json:"begin"`
	End              string    `json:"end"`
	Created_By       string    `json:"created_by"`
	Last_Modified_By string    `json:"last_modified_by"`
	Create_Time      time.Time `json:"create_time"`
	System_Mtime     time.Time `json:"system_mtime"`
	User_Mtime       time.Time `json:"user_mtime"`
	Date_Type        string    `json:"date_type"`
	Label            string    `json:"label"`
	JSONModel_Type   string    `json:"jsonmodel_type"`
}

type Deaccessions struct {
	Lock_Version     int               `json:"lock_version"`
	Description      string            `json:"description"`
	Reason           string            `json:"reason"`
	Disposition      string            `json:"disposition"`
	Notification     bool              `json:"notification"`
	Created_By       string            `json:"created_by"`
	Last_Modified_By string            `json:"last_modified_by"`
	Create_Time      time.Time         `json:"create_time"`
	System_Mtime     time.Time         `json:"system_mtime"`
	User_Mtime       time.Time         `json:"user_mtime"`
	JSONModel_Type   string            `json:"jsonmodel_type"`
	Extents          []Extent          `json:"extents"`
	Repository       map[string]string `json:"repository"`
	Date             Date              `json:"date"`
	Notes            []Note            `json:"notes"`
}

type External_ID struct {
	External_ID      string    `json:"external_id"`
	Source           string    `json:"source"`
	Created_By       string    `json:"created_by"`
	Last_Modified_By string    `json:"last_modified_by"`
	Create_Time      time.Time `json:"create_time"`
	System_Mtime     time.Time `json:"system_mtime"`
	User_Mtime       time.Time `json:"user_mtime"`
	JSONModel_Type   string    `json:"jsonmodel_type"`
}

type Extent struct {
	Lock_Version      int       `json:"lock_version"`
	Number            string    `json:"number"`
	Container_Summary string    `json:"container_summary"`
	Created_By        string    `json:"created_by"`
	Last_Modified_By  string    `json:"last_modified_by"`
	Create_Time       time.Time `json:"create_time"`
	System_Mtime      time.Time `json:"system_mtime"`
	User_Mtime        time.Time `json:"user_mtime"`
	Portion           string    `json:"portion"`
	Extent_Type       string    `json:"extent_type"`
	JSONModel_Type    string    `json:"jsonmodel_type"`
}

type Lang_Material struct {
	Lock_Version        int                 `json:"lock_version"`
	Create_Time         time.Time           `json:"create_time"`
	System_Mtime        time.Time           `json:"system_mtime"`
	User_Mtime          time.Time           `json:"user_mtime"`
	JSONModel_Type      string              `json:"json_modeltype"`
	Notes               []Note_Langmaterial `json:"notes"`
	Language_And_Script Language_And_Script `json:"language_and_script"`
}

type Language_And_Script struct {
	Lock_Version   int       `json:"lock_version"`
	Create_Time    time.Time `json:"create_time"`
	System_Mtime   time.Time `json:"system_mtime"`
	User_Mtime     time.Time `json:"user_mtime"`
	Language       string    `json:"language"`
	JSONModel_Type string    `json:"json_modeltype"`
}

type Linked_Agent struct {
	Role  string        `json:"role"`
	Terms []interface{} `json:"terms"`
	Ref   string        `json:"ref"`
}

type Note_Langmaterial struct {
	JSONModel_Type string   `json:"json_modeltype"`
	Persistant_ID  string   `json:"persistant_id"`
	Label          string   `json:"label"`
	Type           string   `json:"type"`
	Content        []string `json:"content"`
	Publish        bool     `json:"publish"`
}

type Note struct {
	JSONModel_Type string      `json:"jsonmodel_type"`
	Persistent_ID  string      `json:"persistent_id"`
	Label          string      `json:"label"`
	Type           string      `json:"type"`
	Subnotes       []Note_Text `json:"subnotes,omitempty"`
	Content        []string    `json:"content,omitempty"`
}

type Note_Text struct {
	JSONModel_Type string `json:"jsonmodel_type"`
	Content        string `json:"content"`
	Publish        bool   `json:"publish"`
}

type Resource struct {
	Classifications               []interface{}        `json:"classifications"`
	Create_Time                   time.Time            `json:"create_time"`
	Created_By                    string               `json:"created_by"`
	Dates                         []Date               `json:"dates"`
	Deaccessions                  []Deaccessions       `json:"deaccessions"`
	EAD_ID                        string               `json:"ead_id"`
	EAD_Location                  string               `json:"ead_location"`
	Extents                       []Extent             `json:"extents"`
	External_IDs                  []External_ID        `json:"external_ids"`
	Finding_Aid_Author            string               `json:"finding_aid_author"`
	Finding_Aid_Date              string               `json:"finding_aid_date"`
	Finding_Aid_Description_rules string               `json:"finding_aid_description_rules"`
	Finding_Aid_Language          string               `json:"finding_aid_language"`
	Finding_Aid_Language_Note     string               `json:"finding_aid_language_note"`
	Finding_Aid_Script            string               `json:"finding_aid_script"`
	Finding_Aid_Status            string               `json:"finding_aid_status"`
	Finding_Aid_Title             string               `json:"finding_aid_title"`
	ID_0                          string               `json:"id_0"`
	ID_1                          string               `json:"id_1"`
	ID_2                          string               `json:"id_2"`
	ID_3                          string               `json:"id_3"`
	Instances                     []interface{}        `json:"instances"`
	Is_Slug_Auto                  bool                 `json:"is_slug_auto"`
	JSONModel_Type                string               `json:"jsonmodel_type"`
	Lang_Materials                []Lang_Material      `json:"lang_materials"`
	Last_Modified_By              string               `json:"last_modified_by"`
	Level                         string               `json:"level"`
	Linked_Agents                 []Linked_Agent       `json:"linked_agents"`
	Linked_Events                 []interface{}        `json:"linked_events"`
	Lock_Version                  int                  `json:"lock_version"`
	Notes                         []Note               `json:"notes"`
	Publish                       bool                 `json:"publish"`
	Related_Accessions            []map[string]string  `json:"related_accessions"`
	Repository                    map[string]string    `json:"repository"`
	Repository_Processing_Note    string               `json:"repository_processing_note"`
	Restrictions                  bool                 `json:"restrictions"`
	Revision_Statements           []Revision_Statement `json:"revision_statements"`
	RightsStatements              []interface{}
	Subjects                      []map[string]string `json:"subjects"`
	Supressed                     bool                `json:"supressed"`
	System_Mtime                  time.Time           `json:"system_mtime"`
	Title                         string              `json:"title"`
	Tree                          map[string]string   `json:"tree"`
	URI                           string              `json:"uri"` //double check this
	User_Mtime                    time.Time           `json:"user_mtime"`
}

type Revision_Statement struct {
	Date             string            `json:"date"`
	Description      string            `json:"description"`
	Created_By       string            `json:"created_by"`
	Last_Modified_By string            `json:"last_modified_by"`
	Create_Time      time.Time         `json:"create_time"`
	System_Mtime     time.Time         `json:"system_mtime"`
	Publish          bool              `json:"publish"`
	JSONModel        string            `json:"jsonmodel_type"`
	URI              string            `json:"uri"`
	Repository       map[string]string `json:"repository"`
}
