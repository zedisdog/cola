package wxpay

type BodyMap map[string]interface{}

const (
	baseUrl        = "https://api.mch.weixin.qq.com/"            // (生产环境) 微信支付的基地址
	baseUrlSandbox = "https://api.mch.weixin.qq.com/sandboxnew/" // (沙盒环境) 微信支付的基地址

	// wxURL_DownloadFundFlow  = wxBaseUrl + "pay/downloadfundflow"            // 下载资金账单
	// wxURL_BatchQueryComment = wxBaseUrl + "billcommentsp/batchquerycomment" //
	// wxURL_SanBox_DownloadFundFlow  = wxBaseUrlSandbox + "pay/downloadfundflow"
	// wxURL_SanBox_BatchQueryComment = wxBaseUrlSandbox + "billcommentsp/batchquerycomment"

	// 服务模式
	ServiceTypeNormalDomestic      = 1 // 境内普通商户
	ServiceTypeNormalAbroad        = 2 // 境外普通商户
	ServiceTypeFacilitatorDomestic = 3 // 境内服务商
	ServiceTypeFacilitatorAbroad   = 4 // 境外服务商
	ServiceTypeBankServiceProvidor = 5 // 银行服务商

	// 支付类型
	TradeTypeApplet   = "JSAPI"    // 小程序支付
	TradeTypeJsApi    = "JSAPI"    // JSAPI支付
	TradeTypeApp      = "APP"      // APP支付
	TradeTypeH5       = "MWEB"     // H5支付
	TradeTypeNative   = "NATIVE"   // Native支付
	TradeTypeMicropay = "MICROPAY" // 付款码支付

	// 交易状态
	TradeStateSuccess    = "SUCCESS"    // 支付成功
	TradeStateRefund     = "REFUND"     // 转入退款
	TradeStateNotPay     = "NOTPAY"     // 未支付
	TradeStateClosed     = "CLOSED"     // 已关闭
	TradeStateRevoked    = "REVOKED"    // 已撤销(刷卡支付)
	TradeStateUserPaying = "USERPAYING" // 用户支付中
	TradeStatePayError   = "PAYERROR"   // 支付失败(其他原因，如银行返回失败)

	// 交易保障(MICROPAY)上报数据包的交易状态
	ReportMicropayTradeStateOk     = "OK"     // 成功
	ReportMicropayTradeStateFail   = "FAIL"   // 失败
	ReportMicropayTradeStateCancel = "CANCLE" // 取消

	// 签名方式
	SignTypeMD5        = "MD5" // 默认
	SignTypeHmacSHA256 = "HMAC-SHA256"

	// 货币类型
	FeeTypeCNY = "CNY" // 人民币

	// 指定支付方式
	LimitPayNoCredit = "no_credit" // 指定不能使用信用卡支付

	// 压缩账单
	TarTypeGzip = "GZIP"

	// 电子发票
	ReceiptEnable = "Y" // 支付成功消息和支付详情页将出现开票入口

	// 代金券类型
	CouponTypeCash   = "CASH"    // 充值代金券
	CouponTypeNoCash = "NO_CASH" // 非充值优惠券

	// 账单类型
	BillTypeAll            = "ALL"             // 返回当日所有订单信息，默认值
	BillTypeSuccess        = "SUCCESS"         // 返回当日成功支付的订单
	BillTypeRefund         = "REFUND"          // 返回当日退款订单
	BillTypeRechargeRefund = "RECHARGE_REFUND" // 返回当日充值退款订单

	// 退款渠道
	RefundChannelOriginal      = "ORIGINAL"       // 原路退款
	RefundChannelBalance       = "BALANCE"        // 退回到余额
	RefundChannelOtherBalance  = "OTHER_BALANCE"  // 原账户异常退到其他余额账户
	RefundChannelOtherBankCard = "OTHER_BANKCARD" // 原银行卡异常退到其他银行卡

	// 退款状态
	RefundStatusSuccess    = "SUCCESS"     // 退款成功
	RefundStatusClose      = "REFUNDCLOSE" // 退款关闭
	RefundStatusProcessing = "PROCESSING"  // 退款处理中
	RefundStatusChange     = "CHANGE"      // 退款异常，退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台(pay.weixin.qq.com)-交易中心，手动处理此笔退款

	// 退款资金来源
	RefundAccountRechargeFunds  = "REFUND_SOURCE_RECHARGE_FUNDS"  // 可用余额退款/基本账户
	RefundAccountUnsettledFunds = "REFUND_SOURCE_UNSETTLED_FUNDS" // 未结算资金退款

	// 退款发起来源
	RefundRequestSourceApi            = "API"             // API接口
	RefundRequestSourceVendorPlatform = "VENDOR_PLATFORM" // 商户平台

	// 找零校验用户姓名选项
	CheckNameTypeNoCheck    = "NO_CHECK"    //不校验真实姓名
	CheckNameTypeForceCheck = "FORCE_CHECK" //强校验真实姓名

	// 返回结果
	ResponseSuccess = "SUCCESS" // 成功，通信标识或业务结果
	ResponseFail    = "FAIL"    // 失败，通信标识或业务结果

	// 返回消息
	ResponseMessageOk = "OK" // 返回成功信息

	// 错误代码，包括描述、支付状态、原因、解决方案
	ErrCodeAppIdMchIdNotMatch   = "APPID_MCHID_NOT_MATCH" // appid和mch_id不匹配 支付确认失败 appid和mch_id不匹配 请确认appid和mch_id是否匹配
	ErrCodeAppIdNotExist        = "APPID_NOT_EXIST"       // APPID不存在 支付确认失败 参数中缺少APPID 请检查APPID是否正确
	ErrCodeAuthCodeError        = "AUTH_CODE_ERROR"       // 授权码参数错误 支付确认失败 请求参数未按指引进行填写 每个二维码仅限使用一次，请刷新再试
	ErrCodeAuthCodeExpire       = "AUTHCODEEXPIRE"        // 二维码已过期，请用户在微信上刷新后再试 支付确认失败 用户的条码已经过期 请收银员提示用户，请用户在微信上刷新条码，然后请收银员重新扫码。 直接将错误展示给收银员
	ErrCodeAuthCodeInvalid      = "AUTH_CODE_INVALID"     // 授权码检验错误 支付确认失败 收银员扫描的不是微信支付的条码 请扫描微信支付被扫条码/二维码
	ErrCodeBankError            = "BANKERROR"             // 银行系统异常 支付结果未知 银行端超时 请立即调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作。
	ErrCodeBuyerMismatch        = "BUYER_MISMATCH"        // 支付帐号错误 支付确认失败 暂不支持同一笔订单更换支付方 请确认支付方是否相同
	ErrCodeInvalidRequest       = "INVALID_REQUEST"       // 无效请求 支付确认失败 商户系统异常导致，商户权限异常、重复请求支付、证书错误、频率限制等 请确认商户系统是否正常，是否具有相应支付权限，确认证书是否正确，控制频率
	ErrCodeInvalidTransactionId = "INVALID_TRANSACTIONID" // 无效transaction_id 请求参数未按指引进行填写 请求参数错误，检查原交易号是否存在或发起支付交易接口返回失败
	ErrCodeLackParams           = "LACK_PARAMS"           // 缺少参数 支付确认失败 缺少必要的请求参数 请检查参数是否齐全
	ErrCodeMchIdNotExist        = "MCHID_NOT_EXIST"       // MCHID不存在 支付确认失败 参数中缺少MCHID 请检查MCHID是否正确
	ErrCodeNoAuth               = "NOAUTH"                // 商户无权限 支付确认失败 商户没有开通被扫支付权限 请开通商户号权限。请联系产品或商务申请
	ErrCodeNotEnough            = "NOTENOUGH"             // 余额不足 支付确认失败 用户的零钱余额不足 请收银员提示用户更换当前支付的卡，然后请收银员重新扫码。建议：商户系统返回给收银台的提示为“用户余额不足.提示用户换卡支付”
	ErrCodeNotSuportCard        = "NOTSUPORTCARD"         // 不支持卡类型 支付确认失败 用户使用卡种不支持当前支付形式 请用户重新选择卡种 建议：商户系统返回给收银台的提示为“该卡不支持当前支付，提示用户换卡支付或绑新卡支付”
	ErrCodeNotUtf8              = "NOT_UTF8"              // 编码格式错误 支付确认失败 未使用指定编码格式 请使用UTF-8编码格式
	ErrCodeOrderClosed          = "ORDERCLOSED"           // 订单已关闭 支付确认失败 该订单已关 商户订单号异常，请重新下单支付
	ErrCodeOrderPaid            = "ORDERPAID"             // 订单已支付 支付确认失败 订单号重复 请确认该订单号是否重复支付，如果是新单，请使用新订单号提交
	ErrCodeOrderReversed        = "ORDERREVERSED"         // 订单已撤销 支付确认失败 当前订单已经被撤销 当前订单状态为“订单已撤销”，请提示用户重新支付
	ErrCodeOutTradeNoUsed       = "OUT_TRADE_NO_USED"     // 商户订单号重复 支付确认失败 同一笔交易不能多次提交 请核实商户订单号是否重复提交
	ErrCodeParamError           = "PARAM_ERROR"           // 参数错误 支付确认失败 请求参数未按指引进行填写 请根据接口返回的详细信息检查您的程序
	ErrCodePostDataEmpty        = "POST_DATA_EMPTY"       // post数据为空 post数据不能为空 请检查post数据是否为空
	ErrCodeRefundNotExist       = "REFUNDNOTEXIST"        // 退款订单查询失败 订单号错误或订单状态不正确 请检查订单号是否有误以及订单状态是否正确，如：未支付、已支付未退款
	ErrCodeRequirePostMethod    = "REQUIRE_POST_METHOD"   // 请使用post方法 支付确认失败 未使用post传递参数 请检查请求参数是否通过post方法提交
	ErrCodeSignError            = "SIGNERROR"             // 签名错误 支付确认失败 参数签名结果不正确 请检查签名参数和方法是否都符合签名算法要求
	ErrCodeSystemError          = "SYSTEMERROR"           // 接口返回错误 支付结果未知 系统超时 请立即调用被扫订单结果查询API，查询当前订单状态，并根据订单的状态决定下一步的操作。
	ErrCodeTradeError           = "TRADE_ERROR"           // 交易错误 支付确认失败 业务错误导致交易失败、用户账号异常、风控、规则限制等 请确认帐号是否存在异常
	ErrCodeUserPaying           = "USERPAYING"            // 用户支付中，需要输入密码 支付结果未知 该笔交易因为业务规则要求，需要用户输入支付密码。 等待5秒，然后调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作。
	ErrCodeXmlFormatError       = "XML_FORMAT_ERROR"      // XML格式错误 支付确认失败 XML格式错误 请检查XML参数格式是否正确

	// 是否关注公众账号
	IsSubscribeYes = "Y" // 关注
	IsSubscribeNo  = "N" // 未关注

	// 银行类型
	BankTypeIcbcDebit    = "ICBC_DEBIT"    // 工商银行(借记卡)
	BankTypeIcbcCredit   = "ICBC_CREDIT"   // 工商银行(信用卡)
	BankTypeAbcDebit     = "ABC_DEBIT"     // 农业银行(借记卡)
	BankTypeAbcCredit    = "ABC_CREDIT"    // 农业银行(信用卡)
	BankTypePsbcDebit    = "PSBC_DEBIT"    // 邮政储蓄银行(借记卡)
	BankTypePsbcCredit   = "PSBC_CREDIT"   // 邮政储蓄银行(信用卡)
	BankTypeCcbDebit     = "CCB_DEBIT"     // 建设银行(借记卡)
	BankTypeCcbCredit    = "CCB_CREDIT"    // 建设银行(信用卡)
	BankTypeCmbDebit     = "CMB_DEBIT"     // 招商银行(借记卡)
	BankTypeCmbCredit    = "CMB_CREDIT"    // 招商银行(信用卡)
	BankTypeBocDebit     = "BOC_DEBIT"     // 中国银行(借记卡)
	BankTypeBocCredit    = "BOC_CREDIT"    // 中国银行(信用卡)
	BankTypeCommDebit    = "COMM_DEBIT"    // 交通银行(借记卡)
	BankTypeCommCredit   = "COMM_CREDIT"   // 交通银行(信用卡)
	BankTypeSpdbDebit    = "SPDB_DEBIT"    // 浦发银行(借记卡)
	BankTypeSpdbCredit   = "SPDB_CREDIT"   // 浦发银行(信用卡)
	BankTypeGdbDebit     = "GDB_DEBIT"     // 广发银行(借记卡)
	BankTypeGdbCredit    = "GDB_CREDIT"    // 广发银行(信用卡)
	BankTypeCmbcDebit    = "CMBC_DEBIT"    // 民生银行(借记卡)
	BankTypeCmbcCredit   = "CMBC_CREDIT"   // 民生银行(信用卡)
	BankTypePabDebit     = "PAB_DEBIT"     // 平安银行(借记卡)
	BankTypePabCredit    = "PAB_CREDIT"    // 平安银行(信用卡)
	BankTypeCebDebit     = "CEB_DEBIT"     // 光大银行(借记卡)
	BankTypeCebCredit    = "CEB_CREDIT"    // 光大银行(信用卡)
	BankTypeCibDebit     = "CIB_DEBIT"     // 兴业银行(借记卡)
	BankTypeCibCredit    = "CIB_CREDIT"    // 兴业银行(信用卡)
	BankTypeCiticDebit   = "CITIC_DEBIT"   // 中信银行(借记卡)
	BankTypeCiticCredit  = "CITIC_CREDIT"  // 中信银行(信用卡)
	BankTypeBoshDebit    = "BOSH_DEBIT"    // 上海银行(借记卡)
	BankTypeBoshCredit   = "BOSH_CREDIT"   // 上海银行(信用卡)
	BankTypeCrbDebit     = "CRB_DEBIT"     // 华润银行(借记卡)
	BankTypeHzbDebit     = "HZB_DEBIT"     // 杭州银行(借记卡)
	BankTypeHzbCredit    = "HZB_CREDIT"    // 杭州银行(信用卡)
	BankTypeBsbDebit     = "BSB_DEBIT"     // 包商银行(借记卡)
	BankTypeBsbCredit    = "BSB_CREDIT"    // 包商银行(信用卡)
	BankTypeCqbDebit     = "CQB_DEBIT"     // 重庆银行(借记卡)
	BankTypeSdebDebit    = "SDEB_DEBIT"    // 顺德农商行(借记卡)
	BankTypeSzrcbDebit   = "SZRCB_DEBIT"   // 深圳农商银行(借记卡)
	BankTypeSzrcbCredit  = "SZRCB_CREDIT"  // 深圳农商银行(信用卡)
	BankTypeHrbbDebit    = "HRBB_DEBIT"    // 哈尔滨银行(借记卡)
	BankTypeBocdDebit    = "BOCD_DEBIT"    // 成都银行(借记卡)
	BankTypeGdnybDebit   = "GDNYB_DEBIT"   // 南粤银行(借记卡)
	BankTypeGdnybCredit  = "GDNYB_CREDIT"  // 南粤银行(信用卡)
	BankTypeGzcbDebit    = "GZCB_DEBIT"    // 广州银行(借记卡)
	BankTypeGzcbCredit   = "GZCB_CREDIT"   // 广州银行(信用卡)
	BankTypeJsbDebit     = "JSB_DEBIT"     // 江苏银行(借记卡)
	BankTypeJsbCredit    = "JSB_CREDIT"    // 江苏银行(信用卡)
	BankTypeNbcbDebit    = "NBCB_DEBIT"    // 宁波银行(借记卡)
	BankTypeNbcbCredit   = "NBCB_CREDIT"   // 宁波银行(信用卡)
	BankTypeNjcbDebit    = "NJCB_DEBIT"    // 南京银行(借记卡)
	BankTypeQhnxDebit    = "QHNX_DEBIT"    // 青海农信(借记卡)
	BankTypeOrdosbCredit = "ORDOSB_CREDIT" // 鄂尔多斯银行(信用卡)
	BankTypeOrdosbDebit  = "ORDOSB_DEBIT"  // 鄂尔多斯银行(借记卡)
	BankTypeBjrcbCredit  = "BJRCB_CREDIT"  // 北京农商(信用卡)
	BankTypeBhbDebit     = "BHB_DEBIT"     // 河北银行(借记卡)
	BankTypeBgzbDebit    = "BGZB_DEBIT"    // 贵州银行(借记卡)
	BankTypeBeebDebit    = "BEEB_DEBIT"    // 鄞州银行(借记卡)
	BankTypePzhccbDebit  = "PZHCCB_DEBIT"  // 攀枝花银行(借记卡)
	BankTypeQdccbCredit  = "QDCCB_CREDIT"  // 青岛银行(信用卡)
	BankTypeQdccbDebit   = "QDCCB_DEBIT"   // 青岛银行(借记卡)
	BankTypeShinhanDebit = "SHINHAN_DEBIT" // 新韩银行(借记卡)
	BankTypeQlbDebit     = "QLB_DEBIT"     // 齐鲁银行(借记卡)
	BankTypeQsbDebit     = "QSB_DEBIT"     // 齐商银行(借记卡)
	BankTypeZzbDebit     = "ZZB_DEBIT"     // 郑州银行(借记卡)
	BankTypeCcabDebit    = "CCAB_DEBIT"    // 长安银行(借记卡)
	BankTypeRzbDebit     = "RZB_DEBIT"     // 日照银行(借记卡)
	BankTypeScnxDebit    = "SCNX_DEBIT"    // 四川农信(借记卡)
	BankTypeBeebCredit   = "BEEB_CREDIT"   // 鄞州银行(信用卡)
	BankTypeSdrcuDebit   = "SDRCU_DEBIT"   // 山东农信(借记卡)
	BankTypeBczDebit     = "BCZ_DEBIT"     // 沧州银行(借记卡)
	BankTypeSjbDebit     = "SJB_DEBIT"     // 盛京银行(借记卡)
	BankTypeLnnxDebit    = "LNNX_DEBIT"    // 辽宁农信(借记卡)
	BankTypeJufengbDebit = "JUFENGB_DEBIT" // 临朐聚丰村镇银行(借记卡)
	BankTypeZzbCredit    = "ZZB_CREDIT"    // 郑州银行(信用卡)
	BankTypeJxnxbDebit   = "JXNXB_DEBIT"   // 江西农信(借记卡)
	BankTypeJzbDebit     = "JZB_DEBIT"     // 晋中银行(借记卡)
	BankTypeJzcbCredit   = "JZCB_CREDIT"   // 锦州银行(信用卡)
	// BankType                 = "JZCB_DEBIT"        // 锦州银行(借记卡)
	// BankType                 = "KLB_DEBIT"         // 昆仑银行(借记卡)
	// BankType                 = "KRCB_DEBIT"        // 昆山农商(借记卡)
	// BankType                 = "KUERLECB_DEBIT"    // 库尔勒市商业银行(借记卡)
	// BankType                 = "LJB_DEBIT"         // 龙江银行(借记卡)
	// BankType                 = "NYCCB_DEBIT"       // 南阳村镇银行(借记卡)
	// BankType                 = "LSCCB_DEBIT"       // 乐山市商业银行(借记卡)
	// BankType                 = "LUZB_DEBIT"        // 柳州银行(借记卡)
	// BankType                 = "LWB_DEBIT"         // 莱商银行(借记卡)
	// BankType                 = "LYYHB_DEBIT"       // 辽阳银行(借记卡)
	// BankType                 = "LZB_DEBIT"         // 兰州银行(借记卡)
	// BankType                 = "MINTAIB_CREDIT"    // 民泰银行(信用卡)
	// BankType                 = "MINTAIB_DEBIT"     // 民泰银行(借记卡)
	// BankType                 = "NCB_DEBIT"         // 宁波通商银行(借记卡)
	// BankType                 = "NMGNX_DEBIT"       // 内蒙古农信(借记卡)
	// BankType                 = "XAB_DEBIT"         // 西安银行(借记卡)
	// BankType                 = "WFB_CREDIT"        // 潍坊银行(信用卡)
	// BankType                 = "WFB_DEBIT"         // 潍坊银行(借记卡)
	// BankType                 = "WHB_CREDIT"        // 威海商业银行(信用卡)
	// BankType                 = "WHB_DEBIT"         // 威海市商业银行(借记卡)
	// BankType                 = "WHRC_CREDIT"       // 武汉农商(信用卡)
	// BankType                 = "WHRC_DEBIT"        // 武汉农商行(借记卡)
	// BankType                 = "WJRCB_DEBIT"       // 吴江农商行(借记卡)
	// BankType                 = "WLMQB_DEBIT"       // 乌鲁木齐银行(借记卡)
	// BankType                 = "WRCB_DEBIT"        // 无锡农商(借记卡)
	// BankType                 = "WZB_DEBIT"         // 温州银行(借记卡)
	// BankType                 = "XAB_CREDIT"        // 西安银行(信用卡)
	// BankType                 = "WEB_DEBIT"         // 微众银行(借记卡)
	// BankType                 = "XIB_DEBIT"         // 厦门国际银行(借记卡)
	// BankType                 = "XJRCCB_DEBIT"      // 新疆农信银行(借记卡)
	// BankType                 = "XMCCB_DEBIT"       // 厦门银行(借记卡)
	// BankType                 = "YNRCCB_DEBIT"      // 云南农信(借记卡)
	// BankType                 = "YRRCB_CREDIT"      // 黄河农商银行(信用卡)
	// BankType                 = "YRRCB_DEBIT"       // 黄河农商银行(借记卡)
	// BankType                 = "YTB_DEBIT"         // 烟台银行(借记卡)
	// BankType                 = "ZJB_DEBIT"         // 紫金农商银行(借记卡)
	// BankType                 = "ZJLXRB_DEBIT"      // 兰溪越商银行(借记卡)
	// BankType                 = "ZJRCUB_CREDIT"     // 浙江农信(信用卡)
	// BankType                 = "AHRCUB_DEBIT"      // 安徽省农村信用社联合社(借记卡)
	// BankType                 = "BCZ_CREDIT"        // 沧州银行(信用卡)
	// BankType                 = "SRB_DEBIT"         // 上饶银行(借记卡)
	// BankType                 = "ZYB_DEBIT"         // 中原银行(借记卡)
	// BankType                 = "ZRCB_DEBIT"        // 张家港农商行(借记卡)
	// BankType                 = "SRCB_CREDIT"       // 上海农商银行(信用卡)
	// BankType                 = "SRCB_DEBIT"        // 上海农商银行(借记卡)
	// BankType                 = "ZJTLCB_DEBIT"      // 浙江泰隆银行(借记卡)
	// BankType                 = "SUZB_DEBIT"        // 苏州银行(借记卡)
	// BankType                 = "SXNX_DEBIT"        // 山西农信(借记卡)
	// BankType                 = "SXXH_DEBIT"        // 陕西信合(借记卡)
	// BankType                 = "ZJRCUB_DEBIT"      // 浙江农信(借记卡)
	// BankType                 = "AE_CREDIT"         // AE(信用卡)
	// BankType                 = "TACCB_CREDIT"      // 泰安银行(信用卡)
	// BankType                 = "TACCB_DEBIT"       // 泰安银行(借记卡)
	// BankType                 = "TCRCB_DEBIT"       // 太仓农商行(借记卡)
	// BankType                 = "TJBHB_CREDIT"      // 天津滨海农商行(信用卡)
	// BankType                 = "TJBHB_DEBIT"       // 天津滨海农商行(借记卡)
	// BankType                 = "TJB_DEBIT"         // 天津银行(借记卡)
	// BankType                 = "TRCB_DEBIT"        // 天津农商(借记卡)
	// BankType                 = "TZB_DEBIT"         // 台州银行(借记卡)
	// BankType                 = "URB_DEBIT"         // 联合村镇银行(借记卡)
	// BankType                 = "DYB_CREDIT"        // 东营银行(信用卡)
	// BankType                 = "CSRCB_DEBIT"       // 常熟农商银行(借记卡)
	// BankType                 = "CZB_CREDIT"        // 浙商银行(信用卡)
	// BankType                 = "CZB_DEBIT"         // 浙商银行(借记卡)
	// BankType                 = "CZCB_CREDIT"       // 稠州银行(信用卡)
	// BankType                 = "CZCB_DEBIT"        // 稠州银行(借记卡)
	// BankType                 = "DANDONGB_CREDIT"   // 丹东银行(信用卡)
	// BankType                 = "DANDONGB_DEBIT"    // 丹东银行(借记卡)
	// BankType                 = "DLB_CREDIT"        // 大连银行(信用卡)
	// BankType                 = "DLB_DEBIT"         // 大连银行(借记卡)
	// BankType                 = "DRCB_CREDIT"       // 东莞农商银行(信用卡)
	// BankType                 = "DRCB_DEBIT"        // 东莞农商银行(借记卡)
	// BankType                 = "CSRCB_CREDIT"      // 常熟农商银行(信用卡)
	// BankType                 = "DYB_DEBIT"         // 东营银行(借记卡)
	// BankType                 = "DYCCB_DEBIT"       // 德阳银行(借记卡)
	// BankType                 = "FBB_DEBIT"         // 富邦华一银行(借记卡)
	// BankType                 = "FDB_DEBIT"         // 富滇银行(借记卡)
	// BankType                 = "FJHXB_CREDIT"      // 福建海峡银行(信用卡)
	// BankType                 = "FJHXB_DEBIT"       // 福建海峡银行(借记卡)
	// BankType                 = "FJNX_DEBIT"        // 福建农信银行(借记卡)
	// BankType                 = "FUXINB_DEBIT"      // 阜新银行(借记卡)
	// BankType                 = "BOCDB_DEBIT"       // 承德银行(借记卡)
	// BankType                 = "JSNX_DEBIT"        // 江苏农商行(借记卡)
	// BankType                 = "BOLFB_DEBIT"       // 廊坊银行(借记卡)
	// BankType                 = "CCAB_CREDIT"       // 长安银行(信用卡)
	// BankType                 = "CBHB_DEBIT"        // 渤海银行(借记卡)
	// BankType                 = "CDRCB_DEBIT"       // 成都农商银行(借记卡)
	// BankType                 = "BYK_DEBIT"         // 营口银行(借记卡)
	// BankType                 = "BOZ_DEBIT"         // 张家口市商业银行(借记卡)
	// BankType                 = "CFT"               // 零钱
	// BankType                 = "BOTSB_DEBIT"       // 唐山银行(借记卡)
	// BankType                 = "BOSZS_DEBIT"       // 石嘴山银行(借记卡)
	// BankType                 = "BOSXB_DEBIT"       // 绍兴银行(借记卡)
	// BankType                 = "BONX_DEBIT"        // 宁夏银行(借记卡)
	// BankType                 = "BONX_CREDIT"       // 宁夏银行(信用卡)
	// BankType                 = "GDHX_DEBIT"        // 广东华兴银行(借记卡)
	// BankType                 = "BOLB_DEBIT"        // 洛阳银行(借记卡)
	// BankType                 = "BOJX_DEBIT"        // 嘉兴银行(借记卡)
	// BankType                 = "BOIMCB_DEBIT"      // 内蒙古银行(借记卡)
	// BankType                 = "BOHN_DEBIT"        // 海南银行(借记卡)
	// BankType                 = "BOD_DEBIT"         // 东莞银行(借记卡)
	// BankType                 = "CQRCB_CREDIT"      // 重庆农商银行(信用卡)
	// BankType                 = "CQRCB_DEBIT"       // 重庆农商银行(借记卡)
	// BankType                 = "CQTGB_DEBIT"       // 重庆三峡银行(借记卡)
	// BankType                 = "BOD_CREDIT"        // 东莞银行(信用卡)
	// BankType                 = "CSCB_DEBIT"        // 长沙银行(借记卡)
	// BankType                 = "BOB_CREDIT"        // 北京银行(信用卡)
	// BankType                 = "GDRCU_DEBIT"       // 广东农信银行(借记卡)
	// BankType                 = "BOB_DEBIT"         // 北京银行(借记卡)
	// BankType                 = "HRXJB_DEBIT"       // 华融湘江银行(借记卡)
	// BankType                 = "HSBC_DEBIT"        // 恒生银行(借记卡)
	// BankType                 = "HSB_CREDIT"        // 徽商银行(信用卡)
	// BankType                 = "HSB_DEBIT"         // 徽商银行(借记卡)
	// BankType                 = "HUNNX_DEBIT"       // 湖南农信(借记卡)
	// BankType                 = "HUSRB_DEBIT"       // 湖商村镇银行(借记卡)
	// BankType                 = "HXB_CREDIT"        // 华夏银行(信用卡)
	// BankType                 = "HXB_DEBIT"         // 华夏银行(借记卡)
	// BankType                 = "HNNX_DEBIT"        // 河南农信(借记卡)
	// BankType                 = "BNC_DEBIT"         // 江西银行(借记卡)
	// BankType                 = "BNC_CREDIT"        // 江西银行(信用卡)
	// BankType                 = "BJRCB_DEBIT"       // 北京农商行(借记卡)
	// BankType                 = "JCB_DEBIT"         // 晋城银行(借记卡)
	// BankType                 = "JJCCB_DEBIT"       // 九江银行(借记卡)
	// BankType                 = "JLB_DEBIT"         // 吉林银行(借记卡)
	// BankType                 = "JLNX_DEBIT"        // 吉林农信(借记卡)
	// BankType                 = "JNRCB_DEBIT"       // 江南农商(借记卡)
	// BankType                 = "JRCB_DEBIT"        // 江阴农商行(借记卡)
	// BankType                 = "JSHB_DEBIT"        // 晋商银行(借记卡)
	// BankType                 = "HAINNX_DEBIT"      // 海南农信(借记卡)
	// BankType                 = "GLB_DEBIT"         // 桂林银行(借记卡)
	// BankType                 = "GRCB_CREDIT"       // 广州农商银行(信用卡)
	// BankType                 = "GRCB_DEBIT"        // 广州农商银行(借记卡)
	// BankType                 = "GSB_DEBIT"         // 甘肃银行(借记卡)
	// BankType                 = "GSNX_DEBIT"        // 甘肃农信(借记卡)
	// BankType                 = "GXNX_DEBIT"        // 广西农信(借记卡)
	BankTypeGycbCredit       = "GYCB_CREDIT"       // 贵阳银行(信用卡)
	BankTypeGycbDebit        = "GYCB_DEBIT"        // 贵阳银行(借记卡)
	BankTypeGznxDebit        = "GZNX_DEBIT"        // 贵州农信(借记卡)
	BankTypeHainnxCredit     = "HAINNX_CREDIT"     // 海南农信(信用卡)
	BankTypeHkbDebit         = "HKB_DEBIT"         // 汉口银行(借记卡)
	BankTypeHanabDebit       = "HANAB_DEBIT"       // 韩亚银行(借记卡)
	BankTypeHbcbCredit       = "HBCB_CREDIT"       // 湖北银行(信用卡)
	BankTypeHbcbDebit        = "HBCB_DEBIT"        // 湖北银行(借记卡)
	BankTypeHbnxCredit       = "HBNX_CREDIT"       // 湖北农信(信用卡)
	BankTypeHbnxDebit        = "HBNX_DEBIT"        // 湖北农信(借记卡)
	BankTypeHdcbDebit        = "HDCB_DEBIT"        // 邯郸银行(借记卡)
	BankTypeHebnxDebit       = "HEBNX_DEBIT"       // 河北农信(借记卡)
	BankTypeHfbDebit         = "HFB_DEBIT"         // 恒丰银行(借记卡)
	BankTypeHkbeaDebit       = "HKBEA_DEBIT"       // 东亚银行(借记卡)
	BankTypeJcbCredit        = "JCB_CREDIT"        // JCB(信用卡)
	BankTypeMasterCardCredit = "MASTERCARD_CREDIT" // MASTERCARD(信用卡)
	BankTypeVisaCredit       = "VISA_CREDIT"       // VISA(信用卡)
	BankTypeLqt              = "LQT"               // 零钱通
)
