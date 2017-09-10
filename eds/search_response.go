package eds

type SearchResponseMessage struct {
	SearchRequest SearchRequest
	SearchResult  SearchResult
}
type SearchRequest struct {
	SearchCriteria            SearchCriteria
	RetrievalCriteria         RetrievalCriteria
	SearchCriteriaWithActions SearchCriteriaWithActions
}
type SearchCriteriaWithActions struct {
	QueriesWithAction      []QueryWithAction
	FacetFiltersWithAction []FacetFilterWithAction
	LimitersWithAction     []LimiterWithAction
	ExpandersWithAction    []ExpanderWithAction
	PublicationWithAction  PublicationWithAction
}
type QueryWithAction struct {
	Query        Query
	RemoveAction string
}

type FacetFilterWithAction struct {
	FilterId              string
	RemoveAction          string
	FacetValuesWithAction []FacetValueWithAction
}
type FacetValueWithAction struct {
	FacetValues  []FacetValue
	RemoveAction string
}
type LimiterWithAction struct {
	Id                      string
	LimiterValuesWithAction []LimiterValueWithAction
	RemoveAction            string
}
type LimiterValueWithAction struct {
	Value        string
	RemoveAction string
}
type ExpanderWithAction struct {
	Id           string
	RemoveAction string
}
type PublicationWithAction struct {
	Id           string
	RemoveAction string
}

type SearchResult struct {
	Statistics         Statistics
	Data               Data
	AvailableFacets    []AvailableFacet
	RelatedContent     RelatedContent
	AvailableCriteria  AvailableCriteria
	AutoSuggestedTerms []string
	AutoCorrectedTerms []string
}
type Statistics struct {
	TotalHits       int64
	TotalSearchTime int64
	Databases       []Database
}
type Database struct {
	Id     string
	Label  string
	Status string
	Hits   int64
}
type Data struct {
	RecordFormat string
	Records      []Record
}
type Record struct {
	ResultId    int64
	Header      Header
	PLink       string
	ImageInfo   []ImageInfo
	CustomLinks []CustomLink
	FullText    FullText
	Items       []Item
	RecordInfo  RecordInfo
}

type Header struct {
	DbId           string
	DbLabel        string
	An             string
	RelevancyScore string
	AccessLevel    string
	PubType        string
	PubTypeId      string
}

type ImageInfo struct {
	Size   string
	Target string
}

type CustomLink struct {
	Url           string
	Name          string
	Category      string
	Text          string
	Icon          string
	MouseOverText string
}
type FullText struct {
	Links       []Link
	Text        Text
	CustomLinks []CustomLink
}
type Link struct {
	Type string
	URL  string
}
type Text struct {
	Availability string
	Value        string
}

type Item struct {
	Label string
	Name  string
	Group string
	Data  string
}
type AvailableFacet struct{}
type RelatedContent struct{}
type AvailableCriteria struct {
	DateRange DateRange
}
type DateRange struct {
	MinDate string
	MaxDate string
}
type RecordInfo struct {
	AccessInfo AccessInfo
	BibRecord  BibRecord
	Holdings   string
}
type AccessInfo struct {
	Permissions []Permission
}
type Permission struct {
	Flag string
	Type string
}

type BibRecord struct {
	BibEntity        BibEntity
	BibRelationships BibRelationships
	FileInfo         FileInfo
	PersonRecord     PersonRecord
	RightsInfo       RightsInfo
}
type BibEntity struct {
	Classifications     []Classification
	Dates               []Date
	Identifiers         []Identifier
	Languages           []Language
	Numbering           []Number
	PhysicalDescription PhysicalDescription
	Subjects            []Subject
	Titles              []Title
	Type                string
	ItemTypes           []ItemType
	ContentDescriptions []ContentDescription
	Id                  string
}
type Classification struct {
	Code   string
	Scheme string
	Type   string
}

type Date struct {
	D    string
	M    string
	Y    string
	Text string
	Type string
}
type Identifier struct {
	Type  string
	Value string
	Scope string
}
type Language struct {
	Code string
	Text string
}
type Number struct {
	Type  string
	Value string
}
type PhysicalDescription struct {
	Pagination Pagination
}
type Pagination struct {
	PageCount string
	StartPage string
}
type Subject struct {
	Authority   string
	SubjectFull string
	Type        string
}
type Title struct {
	TitleFull string
	Type      string
}
type ItemType struct {
	Type string
	Text string
}
type ContentDescription struct {
	Type string
	Text string
}

type BibRelationships struct {
	HasContributorRelationships []ContributorRelationship
	HasPubAgentRelationships    []OrganizationEntity
	IsPartOfRelationships       []BibEntity
}
type ContributorRelationship struct {
	PersonEntity       PersonEntity
	OrganizationEntity OrganizationEntity
}
type Entity struct {
	Name Name
}
type PersonEntity struct {
	Entity
}
type OrganizationEntity struct {
	Entity
}

type Name struct {
	NameFull string
}
type FileInfo struct {
	File                File
	FileList            []File
	FilePosLinks        []FilePosLink
	FilePosLinkRefLists []FilePosLinkRefList
}
type File struct {
	IsDownloadable string
	Id             string
	FileName       string
	FileLocation   FileLocation
	ImgCategory    string
}
type FileLocation struct {
	Type               string
	LocationTemplateId string
	Path               string
}

type FilePosLink struct {
	Id     string
	FragId string
	FileId string
	Labels []Label
}
type FilePosLinkRefList struct {
	Use             string
	FilePosLinkRefs []FilePosLinkRef
}
type FilePosLinkRef struct {
	FilePosLinkId string
}
type Label struct {
	Type string
	Text string
}

type PersonRecord struct {
	Entity Entity
}
type RightsInfo struct {
	RightsStatements []RightsStatement
}

type RightsStatement struct {
	Type string
	Text string
}
