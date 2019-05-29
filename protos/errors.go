package ankr_default

import "errors"

// List execution errors
const (
	AuthError = "AuthError:"
	ArgumentError = "ArgumentError:"
	MarshalError = "MarshalError:"
	NotFoundError = "NotFoundError:"
	TokenError = "TokenError:"
	LogicError = "LogicError:"
	DbError = "DatabaseError:"
	HashError = "HashError:"
	PublishError = "PublishError:"
	DecodeError = "DecodeError:"
	WalletError = "WalletError:"
)
var (
	ErrDataCenterNotExist        = errors.New(NotFoundError+"DataCenter does not exist")
	ErrUserNotExist              = errors.New(TokenError+"Can not find user")
	ErrAppNotExist              = errors.New(NotFoundError+"App does not exist")
	ErrUserNotOwn                = errors.New(AuthError+"User does not own this app")
	ErrUpdateFailed              = errors.New(LogicError+"App can not be updated")
	ErrUserAlreadyExist          = errors.New(LogicError+"User already existed")
	ErrPasswordError             = errors.New(LogicError+"Password does not match")
	ErrHashPassword              = errors.New(MarshalError+"Hash password failed")
	ErrNamePasswordEmpty         = errors.New(NotFoundError+"Name or Password is empty")
	ErrStatusNotSupportOperation = errors.New(LogicError+"Current status not support operation")
	ErrAppStatusCanNotUpdate    = errors.New(LogicError+"App status is done or cancelled, can not update")
	ErrReplicaTooMany            = errors.New(ArgumentError+"Replica too many")
	ErrUnknown                   = errors.New(LogicError+"Unknown operation or code")
	ErrSyncAppInfo              = errors.New(ArgumentError+"Sync app info error")
	ErrPublish                   = errors.New(PublishError+"mq publish message error")
	ErrConnection                = errors.New(AuthError+"Connection error")
	ErrNoAvailableDataCenter     = errors.New(NotFoundError+"No available data center")
	ErrEmailFormat               = errors.New(ArgumentError+"Email invalid format")
	ErrEmailShouldNotSame        = errors.New(ArgumentError+"Email should not same as before")
	ErrPasswordFormat            = errors.New(AuthError+"Password invalid format")
	ErrUserNameFormat            = errors.New(AuthError+"User name invalid format")
	ErrPasswordLength            = errors.New(AuthError+"Password must be at least 6 characters long")
	ErrCronJobScheduleFormat     = errors.New(ArgumentError+"Cronjob schedule invalid format")
	ErrPassword                  = errors.New(AuthError+"Invalid password")
	ErrPasswordShouldNotSame     = errors.New(AuthError+"Password should not same as before")
	ErrEmailExit                 = errors.New(AuthError+"Email exist")
	ErrTokenNeedRefresh          = errors.New(TokenError+"Token is unavailable, need call refresh token")
	ErrTokenPassedMax            = errors.New(TokenError+"Tokens number reaches max limit(10)")
	ErrTokenParseFailed          = errors.New(TokenError+"Tokens parse failed")
	ErrRefreshToken              = errors.New(AuthError+"Refresh_token error, need login")
	ErrAccessTokenExpired        = errors.New(TokenError+"Access_token expired")
	ErrCanceledTwice             = errors.New(LogicError+"Can not cancel twice")
	ErrPurgedTwice               = errors.New(LogicError+"Can not purge twice")
	ErrAuthNotAllowed            = errors.New(AuthError+"Auth not allow")
	ErrUnexpectedChar            = errors.New(ArgumentError+"Unexpected char")
	ErrPasswordSame              = errors.New(ArgumentError+"Password must be not same as before")
	ErrOldPassword               = errors.New(ArgumentError+"Old password does not match")
	ErrEmailSame                 = errors.New(ArgumentError+"Email must be not same as before")
	ErrUserNotVariyEmail         = errors.New(ArgumentError+"User's email has not been varified, please verify email first")
	ErrUserDeactive              = errors.New(AuthError+"Login failed, account has been locked, please contact admin")
	ErrEmailNoExit               = errors.New(ArgumentError+"Email does not exist")
	ErrEmailNoMatch              = errors.New(ArgumentError+"Email does not match")
	ErrCanNotApplyAsProvider     = errors.New(LogicError+"User already has applied as a cluster provider")
)
