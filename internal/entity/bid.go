package entity

// Defines values for BidAuthorType.
const (
	Organization BidAuthorType = "Organization"
	User         BidAuthorType = "User"
)

// Defines values for BidDecision.
const (
	Approved BidDecision = "Approved"
	Rejected BidDecision = "Rejected"
)

// Defines values for BidStatus.
const (
	BidStatusCanceled  BidStatus = "Canceled"
	BidStatusCreated   BidStatus = "Created"
	BidStatusPublished BidStatus = "Published"
)

// Bid Информация о предложении
type Bid struct {
	// AuthorId Уникальный идентификатор автора предложения, присвоенный сервером.
	AuthorId BidAuthorId `json:"authorId"`

	// AuthorType Тип автора
	AuthorType BidAuthorType `json:"authorType"`

	// CreatedAt Серверная дата и время в момент, когда пользователь отправил предложение на создание.
	// Передается в формате RFC3339.
	CreatedAt string `json:"createdAt"`

	// Description Описание предложения
	Description BidDescription `json:"description"`

	// Id Уникальный идентификатор предложения, присвоенный сервером.
	Id BidId `json:"id"`

	// Name Полное название предложения
	Name BidName `json:"name"`

	// Status Статус предложения
	Status BidStatus `json:"status"`

	// TenderId Уникальный идентификатор тендера, присвоенный сервером.
	TenderId TenderId `json:"tenderId"`

	// Version Номер версии посел правок
	Version BidVersion `json:"version"`
}

// BidAuthorId Уникальный идентификатор автора предложения, присвоенный сервером.
type BidAuthorId = string

// BidAuthorType Тип автора
type BidAuthorType string

// BidDecision Решение по предложению
type BidDecision string

// BidDescription Описание предложения
type BidDescription = string

// BidFeedback Отзыв на предложение
type BidFeedback = string

// BidId Уникальный идентификатор предложения, присвоенный сервером.
type BidId = string

// BidName Полное название предложения
type BidName = string

// BidReview Отзыв о предложении
type BidReview struct {
	// CreatedAt Серверная дата и время в момент, когда пользователь отправил отзыв на предложение.
	// Передается в формате RFC3339.
	CreatedAt string `json:"createdAt"`

	// Description Описание предложения
	Description BidReviewDescription `json:"description"`

	// Id Уникальный идентификатор отзыва, присвоенный сервером.
	Id BidReviewId `json:"id"`
}

// BidReviewDescription Описание предложения
type BidReviewDescription = string

// BidReviewId Уникальный идентификатор отзыва, присвоенный сервером.
type BidReviewId = string

// BidStatus Статус предложения
type BidStatus string

// BidVersion Номер версии посел правок
type BidVersion = int32

// ErrorResponse Используется для возвращения ошибки пользователю
type ErrorResponse struct {
	// Reason Описание ошибки в свободной форме
	Reason string `json:"reason"`
}

// GetUserBidsParams defines parameters for GetUserBids.
type GetUserBidsParams struct {
	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *PaginationLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset   *PaginationOffset `form:"offset,omitempty" json:"offset,omitempty"`
	Username *Username         `form:"username,omitempty" json:"username,omitempty"`
}

// CreateBidJSONBody defines parameters for CreateBid.
type CreateBidJSONBody struct {
	// AuthorId Уникальный идентификатор автора предложения, присвоенный сервером.
	AuthorId BidAuthorId `json:"authorId"`

	// AuthorType Тип автора
	AuthorType BidAuthorType `json:"authorType"`

	// Description Описание предложения
	Description BidDescription `json:"description"`

	// Name Полное название предложения
	Name BidName `json:"name"`

	// TenderId Уникальный идентификатор тендера, присвоенный сервером.
	TenderId TenderId `json:"tenderId"`
}

// EditBidJSONBody defines parameters for EditBid.
type EditBidJSONBody struct {
	// Description Описание предложения
	Description *BidDescription `json:"description,omitempty"`

	// Name Полное название предложения
	Name *BidName `json:"name,omitempty"`
}

// EditBidParams defines parameters for EditBid.
type EditBidParams struct {
	Username Username `form:"username" json:"username"`
}

// SubmitBidFeedbackParams defines parameters for SubmitBidFeedback.
type SubmitBidFeedbackParams struct {
	BidFeedback BidFeedback `form:"bidFeedback" json:"bidFeedback"`
	Username    Username    `form:"username" json:"username"`
}

// RollbackBidParams defines parameters for RollbackBid.
type RollbackBidParams struct {
	Username Username `form:"username" json:"username"`
}

// GetBidStatusParams defines parameters for GetBidStatus.
type GetBidStatusParams struct {
	Username Username `form:"username" json:"username"`
}

// UpdateBidStatusParams defines parameters for UpdateBidStatus.
type UpdateBidStatusParams struct {
	Status   BidStatus `form:"status" json:"status"`
	Username Username  `form:"username" json:"username"`
}

// SubmitBidDecisionParams defines parameters for SubmitBidDecision.
type SubmitBidDecisionParams struct {
	Decision BidDecision `form:"decision" json:"decision"`
	Username Username    `form:"username" json:"username"`
}

// GetBidsForTenderParams defines parameters for GetBidsForTender.
type GetBidsForTenderParams struct {
	Username Username `form:"username" json:"username"`

	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *PaginationLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset *PaginationOffset `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetBidReviewsParams defines parameters for GetBidReviews.
type GetBidReviewsParams struct {
	// AuthorUsername Имя пользователя автора предложений, отзывы на которые нужно просмотреть.
	AuthorUsername Username `form:"authorUsername" json:"authorUsername"`

	// RequesterUsername Имя пользователя, который запрашивает отзывы.
	RequesterUsername Username `form:"requesterUsername" json:"requesterUsername"`

	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *PaginationLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset *PaginationOffset `form:"offset,omitempty" json:"offset,omitempty"`
}

// CreateBidJSONRequestBody defines body for CreateBid for application/json ContentType.
//type CreateBidJSONRequestBody struct {
//
//	// Name Полное название предложения
//	Name BidName `json:"name"`
//
//	// Description Описание предложения
//	Description BidDescription `json:"description"`
//
//	// Description Описание предложения
//	Status BidStatus `json:"status"`
//
//	// TenderId Уникальный идентификатор тендера, присвоенный сервером.
//	TenderId TenderId `json:"tenderId"`
//
//	// OrganizationId Уникальный идентификатор тендера, присвоенный сервером.
//	OrgId OrganizationId `json:"organizationId"`
//
//	CreatorUsername Username `json:"creatorUsername"`
//
//	AuthorType BidAuthorType `json:"authorType"`
//}

type CreateBidJSONRequestBody CreateBidJSONBody

// EditBidJSONRequestBody defines body for EditBid for application/json ContentType.
type EditBidJSONRequestBody EditBidJSONBody
