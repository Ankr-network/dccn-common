package ankr_default

import "errors"

// List execution errors
const (
	AuthError          = "AuthError:"
	ArgumentError      = "ArgumentError:"
	MarshalError       = "MarshalError:"
	NotFoundError      = "NotFoundError:"
	TokenError         = "TokenError:"
	LogicError         = "LogicError:"
	DbError            = "DatabaseError:"
	HashError          = "HashError:"
	PublishError       = "PublishError:"
	DecodeError        = "DecodeError:"
	WalletError        = "WalletError:"
	CertificationError = "CertificationError:"
	DialError          = "DialError:"
	LoadError          = "LoadError:"
	ReCaptchaError     = "ReCaptchaError:"
	PasswordCharError  = "PasswordCharError:"
	LogMgrError        = "LogMgrError:"
	ClusterError       = "DcMgrError:"
)

var (
	ErrPasswordForbiddenChar     = errors.New(PasswordCharError + "There are not allowed characters in the password. Please only use the combination of letters, numbers and symbols.")
	ErrTimeoutorDuplicate        = errors.New(ReCaptchaError + "There is duplcate or time-out recaptha code. Please refresh the page and try again.")
	ErrNoChartReadme             = errors.New(NotFoundError + "Cannot find readme in chart tarball")
	ErrCreateRequest             = errors.New(DialError + "Cannot create request")
	ErrCannotReadDownload        = errors.New(ArgumentError + "Cannot read chart tarball")
	ErrUnMarshalChartDetail      = errors.New(MarshalError + "Cannot unmarshal chart details")
	ErrCannotReadChartDetails    = errors.New(ArgumentError + "Cannot read chart details response body")
	ErrChartDetailGet            = errors.New(DialError + "Cannot http get chart details")
	ErrUnMarshalChartList        = errors.New(MarshalError + "Cannot unmarshal chart list")
	ErrCannotReadChartList       = errors.New(ArgumentError + "Cannot read chart list response body")
	ErrCannotGetChartList        = errors.New(DialError + "Cannot http get chart list")
	ErrEmptyChartProperties      = errors.New(ArgumentError + "Empty chart properties not accepted")
	ErrCannotUploadChartTar      = errors.New(LoadError + "Cannot upload chart tar file")
	ErrCannotGetChartTar         = errors.New(LoadError + "Cannot open chart tar file")
	ErrCannotGetChartOutdir      = errors.New(LoadError + "Cannot get chart outdir")
	ErrCannotLoadChart           = errors.New(LoadError + "Cannot load chart from tar file")
	ErrNoApp                     = errors.New(NotFoundError + "Null app provided")
	ErrNoAppname                 = errors.New(NotFoundError + "Null app name provided")
	ErrInvalidInput              = errors.New(ArgumentError + "Invalid input, create failed")
	ErrChartAlreadyExist         = errors.New(ArgumentError + "Chart already exist, create failed")
	ErrSaveChartAlreadyExist     = errors.New(ArgumentError + "Save Chart already exist, save failed")
	ErrChartDetailEmpty          = errors.New(NotFoundError + "ChartDetail is empty/nil")
	ErrChartNotExist             = errors.New(NotFoundError + "App chart not exist")
	ErrOriginalChartNotExist     = errors.New(NotFoundError + "Original chart not exist")
	ErrChartMuseumGet            = errors.New(NotFoundError + "Cannot get app chart from chart museum")
	ErrNsEmpty                   = errors.New(NotFoundError + "Namespace is empty/nil")
	ErrNsClEmpty                 = errors.New(NotFoundError + "Namespace or ClusterId is empty/nil")
	ErrDataCenterNotExist        = errors.New(NotFoundError + "DataCenter does not exist")
	ErrUserNotExist              = errors.New(TokenError + "Can not find user")
	ErrAppNotExist               = errors.New(NotFoundError + "App does not exist")
	ErrUserNotOwn                = errors.New(AuthError + "User does not own this app")
	ErrUpdateFailed              = errors.New(LogicError + "App can not be updated")
	ErrUserAlreadyExist          = errors.New(LogicError + "User already existed")
	ErrPasswordError             = errors.New(LogicError + "Password does not match")
	ErrHashPassword              = errors.New(MarshalError + "Hash password failed")
	ErrNamePasswordEmpty         = errors.New(NotFoundError + "Name or Password is empty")
	ErrStatusNotSupportOperation = errors.New(LogicError + "Current status not support operation")
	ErrAppStatusCanNotUpdate     = errors.New(LogicError + "App status is done or cancelled, can not update")
	ErrNSStatusCanNotUpdate      = errors.New(LogicError + "Namespace status is not running, can not update")
	ErrReplicaTooMany            = errors.New(ArgumentError + "Replica too many")
	ErrUnknown                   = errors.New(LogicError + "Unknown operation or code")
	ErrSyncAppInfo               = errors.New(ArgumentError + "Sync app info error")
	ErrPublish                   = errors.New(PublishError + "mq publish message error")
	ErrConnection                = errors.New(AuthError + "Connection error")
	ErrNoAvailableDataCenter     = errors.New(NotFoundError + "No available data center")
	ErrEmailFormat               = errors.New(ArgumentError + "Email invalid format")
	ErrEmailShouldNotSame        = errors.New(ArgumentError + "Email should not same as before")
	ErrPasswordFormat            = errors.New(AuthError + "Password invalid format")
	ErrUserNameFormat            = errors.New(AuthError + "User name invalid format")
	ErrPasswordLength            = errors.New(AuthError + "Password must be at least 6 characters long")
	ErrCronJobScheduleFormat     = errors.New(ArgumentError + "Cronjob schedule invalid format")
	ErrPassword                  = errors.New(AuthError + "Invalid password")
	ErrPasswordShouldNotSame     = errors.New(AuthError + "Password should not same as before")
	ErrEmailExit                 = errors.New(AuthError + "This email address is already associated with an existing account")
	ErrTokenNeedRefresh          = errors.New(TokenError + "Token is unavailable, need call refresh token")
	ErrTokenPassedMax            = errors.New(TokenError + "Tokens number reaches max limit(10)")
	ErrTokenParseFailed          = errors.New(TokenError + "Tokens parse failed")
	ErrRefreshToken              = errors.New(AuthError + "Refresh_token error, need login")
	ErrAccessTokenExpired        = errors.New(TokenError + "Access_token expired")
	ErrCanceledTwice             = errors.New(LogicError + "Can not cancel twice")
	ErrPurgedTwice               = errors.New(LogicError + "Can not purge twice")
	ErrAlreadyPurged             = errors.New(LogicError + "App already purged")
	ErrAuthNotAllowed            = errors.New(AuthError + "Auth not allow")
	ErrUnexpectedChar            = errors.New(ArgumentError + "Unexpected char")
	ErrPasswordSame              = errors.New(ArgumentError + "Password must be not same as before")
	ErrOldPassword               = errors.New(ArgumentError + "Old password does not match")
	ErrEmailSame                 = errors.New(ArgumentError + "Email must be not same as before")
	ErrUserNotVariyEmail         = errors.New(ArgumentError + "User's email has not been varified, please verify email first")
	ErrUserDeactive              = errors.New(AuthError + "Login failed, account has been locked, please contact admin")
	ErrEmailNoExit               = errors.New(ArgumentError + "Email does not exist")
	ErrEmailNoMatch              = errors.New(ArgumentError + "Email does not match")
	ErrCanNotApplyAsProvider     = errors.New(LogicError + "User already has applied as a cluster provider")
	ErrElasticsearchPing         = errors.New(LogMgrError + "Elasticsearch can not connect")
	ErrElasticsearchSearchAfter  = errors.New(LogMgrError + "Elasticsearch search after failed")
	ErrElasticsearchCount        = errors.New(LogMgrError + "Elasticsearch count failed")
	ErrElasticsearchQuery        = errors.New(LogMgrError + "Elasticsearch query syntax wrong")
	ErrUserStillPending          = errors.New(ClusterError + "User is still pending for a becoming cluster provider, can not register now")
	ErrUserRegisteredCluster     = errors.New(ClusterError + "User has already registered one cluster, can not register again")
)
