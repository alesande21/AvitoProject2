package entity

const (
	Construction TenderServiceType = "Construction"
	Delivery     TenderServiceType = "Delivery"
	Manufacture  TenderServiceType = "Manufacture"
)

// Defines values for TenderStatus.
const (
	TenderStatusClosed    TenderStatus = "Closed"
	TenderStatusCreated   TenderStatus = "Created"
	TenderStatusPublished TenderStatus = "Published"
)

// OrganizationId Уникальный идентификатор организации, присвоенный сервером.
type OrganizationId = string

// Tender Информация о тендере
type Tender struct {
	// CreatedAt Серверная дата и время в момент, когда пользователь отправил тендер на создание.
	// Передается в формате RFC3339.
	CreatedAt string `json:"createdAt"`

	// Description Описание тендера
	Description TenderDescription `json:"description"`

	// Id Уникальный идентификатор тендера, присвоенный сервером.
	Id TenderId `json:"id"`

	// Name Полное название тендера
	Name TenderName `json:"name"`

	// OrganizationId Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId OrganizationId `json:"organizationId"`

	// ServiceType Вид услуги, к которой относиться тендер
	ServiceType TenderServiceType `json:"serviceType"`

	// Status Статус тендер
	Status TenderStatus `json:"status"`

	// Version Номер версии посел правок
	Version TenderVersion `json:"version"`
}

// TenderDescription Описание тендера
type TenderDescription = string

// TenderId Уникальный идентификатор тендера, присвоенный сервером.
type TenderId = string

// TenderName Полное название тендера
type TenderName = string

// TenderServiceType Вид услуги, к которой относиться тендер
type TenderServiceType string

// TenderStatus Статус тендер
type TenderStatus string

// TenderVersion Номер версии посел правок
type TenderVersion = int32

// Username Уникальный slug пользователя.
type Username = string

// PaginationLimit defines model for paginationLimit.
type PaginationLimit = int32

// PaginationOffset defines model for paginationOffset.
type PaginationOffset = int32

// GetTendersParams defines parameters for GetTenders.
type GetTendersParams struct {
	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *PaginationLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset *PaginationOffset `form:"offset,omitempty" json:"offset,omitempty"`

	// ServiceType Возвращенные тендеры должны соответствовать указанным видам услуг.
	//
	// Если список пустой, фильтры не применяются.
	ServiceType *[]TenderServiceType `form:"service_type,omitempty" json:"service_type,omitempty"`
}

// GetUserTendersParams defines parameters for GetUserTenders.
type GetUserTendersParams struct {
	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *PaginationLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset   *PaginationOffset `form:"offset,omitempty" json:"offset,omitempty"`
	Username *Username         `form:"username,omitempty" json:"username,omitempty"`
}

// CreateTenderJSONBody defines parameters for CreateTender.
type CreateTenderJSONBody struct {
	// CreatorUsername Уникальный slug пользователя.
	CreatorUsername Username `json:"creatorUsername"`

	// Description Описание тендера
	Description TenderDescription `json:"description"`

	// Name Полное название тендера
	Name TenderName `json:"name"`

	// OrganizationId Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId OrganizationId `json:"organizationId"`

	// ServiceType Вид услуги, к которой относиться тендер
	ServiceType TenderServiceType `json:"serviceType"`
}

// EditTenderJSONBody defines parameters for EditTender.
type EditTenderJSONBody struct {
	// Description Описание тендера
	Description *TenderDescription `json:"description,omitempty"`

	// Name Полное название тендера
	Name *TenderName `json:"name,omitempty"`

	// ServiceType Вид услуги, к которой относиться тендер
	ServiceType *TenderServiceType `json:"serviceType,omitempty"`
}

// EditTenderParams defines parameters for EditTender.
type EditTenderParams struct {
	Username Username `form:"username" json:"username"`
}

// RollbackTenderParams defines parameters for RollbackTender.
type RollbackTenderParams struct {
	Username Username `form:"username" json:"username"`
}

// GetTenderStatusParams defines parameters for GetTenderStatus.
type GetTenderStatusParams struct {
	Username *Username `form:"username,omitempty" json:"username,omitempty"`
}

// UpdateTenderStatusParams defines parameters for UpdateTenderStatus.
type UpdateTenderStatusParams struct {
	Status   TenderStatus `form:"status" json:"status"`
	Username Username     `form:"username" json:"username"`
}

// CreateTenderJSONRequestBody defines body for CreateTender for application/json ContentType.
type CreateTenderJSONRequestBody CreateTenderJSONBody

// EditTenderJSONRequestBody defines body for EditTender for application/json ContentType.
type EditTenderJSONRequestBody EditTenderJSONBody
