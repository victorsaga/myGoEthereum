package ResultCode

const (
	/// <summary>
	/// Success Code
	/// </summary>
	Success = "00000"

	/// <summary>
	/// Account - The username you requested is already being used by another Account.
	/// </summary>
	AccountNameIsAlreadyUsed = "10001"

	/// <summary>
	/// Account - Account does not exist
	/// </summary>
	AccountNameDoesNotExist = "10002"

	/// <summary>
	/// Account - Account is frozen
	/// </summary>
	AccountIsFrozen = "10003"

	/// <summary>
	/// Account - Account Name Invalid
	/// </summary>
	AccountNameInvalid = "10004"

	/// <summary>
	/// Account - Account Currency Invalid
	/// </summary>
	AccountCurrencyInvalid = "10005"

	/// <summary>
	/// Account - Account Password Invalid
	/// </summary>
	AccountPasswordInvalid = "10006"

	/// <summary>
	/// Account - Account Insert Error
	/// </summary>
	AccountInsertError = "10007"

	/// <summary>
	/// Account - Kick Failed
	/// </summary>
	AccountLogoutFailed = "10008"

	/// <summary>
	/// Account - Create Failed
	/// </summary>
	AccountCreateFailed = "10009"

	/// <summary>
	/// Account - Account No Permission
	/// </summary>
	AccountNoPermission = "10010"

	// /// <summary>
	// /// Finance - Enter amount is fail
	// /// </summary>
	AmountFail = "20001"

	// /// <summary>
	// /// Finance - Transaction Status Fail
	// /// </summary>
	// TransactionStatusFail = 20002

	// /// <summary>
	// /// External transaction ID already exists
	// /// </summary>
	// TransactionIDAlreadyExists = 20005

	// /// <summary>
	// /// External transaction ID not exists
	// /// </summary>
	// TransactionIDNotExists = 20006

	// /// <summary>
	// /// External transaction ID Create Fail
	// /// </summary>
	// TransactionIDCreateFail = 20007

	// /// <summary>
	// /// External transaction ID not exists at Thridparty
	// /// </summary>
	// TransactionIDNotExistsAtThirdparty = 20008

	// /// <summary>
	// /// Deposit - Amount is less than minimum deposit amount
	// /// </summary>
	// LessMinimumDepositAmount = 21001

	// /// <summary>
	// /// Deposit - Amount exeeds maximum deposit limit
	// /// </summary>
	// ExeedsMaximumDepositLimit = 21002

	// /// <summary>
	// /// Deposit - Account has waiting withdrawal requests
	// /// </summary>
	// WaitingWithdrawalRequest = 21003

	// /// <summary>
	// /// Deposit - kiosk admin insufficient deposit balance
	// /// </summary>
	// InsufficientDepositBalance = 21004

	// /// <summary>
	// /// Withdrawal - Amount is below minimum allowed cashout amount
	// /// </summary>
	// BelowAllowedCashoutAmount = 22001

	// /// <summary>
	// /// Withdrawal - Amount is over maximum allowed cashout amount
	// /// </summary>
	// OverAllowedCashoutAmount = 22002

	// /// <summary>
	// /// Withdrawal - Insufficient balance
	// /// </summary>
	InsufficientBalance = "22003"

	// /// <summary>
	// /// Withdrawal - Account is in game
	// /// </summary>
	// AccountIsInGame = 22004

	// /// <summary>
	// /// Withdrawal - Insufficient balance
	// /// </summary>
	// InsufficientCredit = 22005

	// /// <summary>
	// /// Withdrawal - Merchant Insufficient balance
	// /// </summary>
	// MerchantInsufficientBalance = 22006

	// /// <summary>
	// /// 超過商戶可代收額度
	// /// </summary>
	// MerchantDepositLimitExceeded = 22007

	// /// <summary>
	// /// 超過商戶可代付額度
	// /// </summary>
	// MerchantWithdrawLimitExceeded = 22008

	DataDuplicate = "30001"

	DataNotExists = "30002"

	// /// <summary>
	// /// Bet Record Not Exists
	// /// </summary>
	// OutOfRange = 30003

	/// <summary>
	/// Auth Token invalid or expire
	/// </summary>
	InvalidToken = "40001"

	// /// <summary>
	// /// Invalid Secret Key
	// /// </summary>
	// InvalidSecretKey = 40002

	// /// <summary>
	// /// Hi speed operation
	// /// </summary>
	// HiSpeedOperation = 50001

	// /// <summary>
	// /// The SSL connection could not be established
	// /// </summary>
	// HttpAuthenticationFailed = 70001

	// /// <summary>
	// /// Web server is too busy
	// /// </summary>
	WebServerIsTooBusy = "70002"

	// /// <summary>
	// /// Http Exception ex: Timeout
	// /// </summary>
	HttpException = "70003"

	// /// <summary>
	// /// Http not OK ex: 403 404 502
	// /// </summary>
	HttpNotOk = "70004"

	// /// <summary>
	// /// Invalid Sign
	// /// </summary>
	InvalidSign = "70005"

	// /// <summary>
	// /// Invalid IP
	// /// </summary>
	// InvalidIP = 70006

	/// <summary>
	/// Parameter Error
	/// </summary>
	Parameter = "90000"

	// /// <summary>
	// /// Cryptography Fail
	// /// </summary>
	// CryptographyFail = 90001

	// /// <summary>
	// /// Response Fail
	// /// </summary>
	// ResponseFail = 90002

	// /// <summary>
	// /// Maintain
	// /// </summary>
	// Maintenance = 90003

	// /// <summary>
	// /// Setting Fail
	// /// </summary>
	SettingFail = "90004"

	// /// <summary>
	// /// Not Support
	// /// </summary>
	// NotSupport = 90005

	// /// <summary>
	// /// Parameter Error
	// /// </summary>
	// ParameterBankInfo = 90006

	// /// <summary>
	// /// Thirdparty Return Fail Without Reason
	// /// </summary>
	// ThirdpartyReturnFail = 90007

	/// <summary>
	/// Unknown Error
	/// </summary>
	Unknown = "99999"
)
