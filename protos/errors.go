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
	ErrStatusNotSupportOperation = errors.New("Current status not support operation")
	ErrAppStatusCanNotUpdate    = errors.New("app status is done or cancelled, can not update")
	ErrReplicaTooMany            = errors.New(ArgumentError+"replica too many")
	ErrUnknown                   = errors.New("unknown operation or code")
	ErrSyncAppInfo              = errors.New("sync app info error")
	ErrPublish                   = errors.New(PublishError+"mq publish message error")
	ErrConnection                = errors.New("connection error")
	ErrNoAvailableDataCenter     = errors.New("no available data center")
	ErrEmailFormat               = errors.New(ArgumentError+"email invalid format")
	ErrEmailShouldNotSame        = errors.New(ArgumentError+"email should not same as before")
	ErrPasswordFormat            = errors.New(AuthError+"password invalid format")
	ErrUserNameFormat            = errors.New(AuthError+"user name invalid format")
	ErrPasswordLength            = errors.New(AuthError+"password must be at least 6 characters long")
	ErrCronJobScheduleFormat     = errors.New("cronjob schedule invalid format")
	ErrPassword                  = errors.New(AuthError+"invalid password")
	ErrPasswordShouldNotSame     = errors.New(AuthError+"password should not same as before")
	ErrEmailExit                 = errors.New(AuthError+"email exist")
	ErrTokenNeedRefresh          = errors.New(TokenError+"token is unavailable, need call refresh token")
	ErrTokenPassedMax            = errors.New(TokenError+"tokens number reaches max limit(10)")
	ErrTokenParseFailed          = errors.New(TokenError+"tokens parse failed")
	ErrRefreshToken              = errors.New(AuthError+"refresh_token error, need login")
	ErrAccessTokenExpired        = errors.New(TokenError+"access_token expired")
	ErrCanceledTwice             = errors.New(LogicError+"can not cancel twice")
	ErrPurgedTwice               = errors.New(LogicError+"can not purge twice")
	ErrAuthNotAllowed            = errors.New(AuthError+"Auth not allow")
	ErrUnexpectedChar            = errors.New(ArgumentError+"unexpected char")
	ErrPasswordSame              = errors.New(ArgumentError+"Password must be not same as before")
	ErrOldPassword               = errors.New(ArgumentError+"old password does not match")
	ErrEmailSame                 = errors.New(ArgumentError+"email must be not same as before")
	ErrUserNotVariyEmail         = errors.New(ArgumentError+"user's email has not been varified, please verify email first")
	ErrUserDeactive              = errors.New(AuthError+"login failed, account has been locked, please contact admin")
	ErrEmailNoExit               = errors.New(ArgumentError+"email does not exist")
	ErrEmailNoMatch              = errors.New(ArgumentError+"email does not match")
	ErrCanNotApplyAsProvider     = errors.New(LogicError+"user already has applied as a cluster provider")
)
