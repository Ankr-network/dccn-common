package ankr_default

import "errors"

// List execution errors
const (
	AuthError = "AuthError:"
	MarshalError = "MarshalError:"
	NotFoundError = "NotFoundError:"
	TokenError = "TokenError:"
	LogicError = "LogicError:"
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
	ErrReplicaTooMany            = errors.New("replica too many")
	ErrUnknown                   = errors.New("unknown operation or code")
	ErrSyncAppInfo              = errors.New("sync app info error")
	ErrPublish                   = errors.New("mq publish message error")
	ErrConnection                = errors.New("connection error")
	ErrNoAvailableDataCenter     = errors.New("no available data center")
	ErrEmailFormat               = errors.New("email invalid format")
	ErrEmailShouldNotSame        = errors.New("email should not same as before")
	ErrPasswordFormat            = errors.New("password invalid format")
	ErrUserNameFormat            = errors.New("user name invalid format")
	ErrPasswordLength            = errors.New("password must be at least 6 characters long")
	ErrCronJobScheduleFormat     = errors.New("cronjob schedule invalid format")
	ErrPassword                  = errors.New("invalid password")
	ErrPasswordShouldNotSame     = errors.New("password should not same as before")
	ErrEmailExit                 = errors.New("email exist")
	ErrTokenNeedRefresh          = errors.New("token is unavailable, need call refresh token")
	ErrTokenPassedMax            = errors.New("tokens number reaches max limit(10)")
	ErrTokenParseFailed          = errors.New("tokens parse failed")
	ErrRefreshToken              = errors.New("refresh_token error, need login")
	ErrAccessTokenExpired        = errors.New("access_token expired")
	ErrCanceledTwice             = errors.New("can not cancel twice")
	ErrPurgedTwice               = errors.New("can not purge twice")
	ErrAuthNotAllowed            = errors.New("AuthError:Auth not allow")
	ErrUnexpectedChar            = errors.New("unexpected char")
	ErrPasswordSame              = errors.New("Password must be not same as before")
	ErrOldPassword               = errors.New("old password does not match")
	ErrEmailSame                 = errors.New("email must be not same as before")
	ErrUserNotVariyEmail         = errors.New("user's email has not been varified, please verify email first")
	ErrUserDeactive              = errors.New("login failed, account has been locked, please contact admin")
	ErrEmailNoExit               = errors.New("email does not exist")
	ErrEmailNoMatch              = errors.New("email does not match")
	ErrCanNotApplyAsProvider     = errors.New("user already has applied as a cluster provider")
)
