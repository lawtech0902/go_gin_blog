package e

/*
错误码定制
*/

const (
	Success = 20000
	Error   = 50000
	
	InvalidParams = 40000
	AuthError     = 40001
	
	TokenCheckError   = 40002
	TokenTimeoutError = 40003
	TokenGenError     = 40004
	
	TagExistError    = 40005
	TagGetError      = 40006
	TagNotExistError = 40007
	TagGetAllError   = 40008
	TagCreateError   = 40009
	TagEditError     = 40010
	TagDeleteError   = 40011
	
	CategoryExistError    = 40012
	CategoryGetError      = 40013
	CategoryNotExistError = 40014
	CategoryGetAllError   = 40015
	CategoryCreateError   = 40016
	CategoryEditError     = 40017
	CategoryDeleteError   = 40018
	
	ArticleExistError    = 40019
	ArticleGetError      = 40020
	ArticleNotExistError = 40021
	ArticleGetAllError   = 40022
	ArticleCountError    = 40023
	ArticleCreateError   = 40024
	ArticleEditError     = 40025
	ArticleDeleteError   = 40026
	
	UserGetError      = 40027
	UserEditError     = 40028
	RestPasswordError = 40031
	
	GetUploadImageError  = 40029
	SaveUploadImageError = 40030
	
	CommentGetAllError = 40032
	SearchArticleError = 40033
	CommentGetError    = 40034
	CommentDeleteError = 40035
	
	SoupGetAllError = 40036
	SoupCreateError = 40037
	SoupGetError    = 40038
	SoupDeleteError = 40039
	SoupEditError   = 40040
	SoupGetRandError = 40041
)
