package tavern

import "regexp"

const (
	alphaRegexString                 = "^[a-zA-Z]+$"
	alphaNumericRegexString          = "^[a-zA-Z0-9]+$"
	alphaUnicodeRegexString          = "^[\\p{L}]+$"
	alphaUnicodeNumericRegexString   = "^[\\p{L}\\p{N}]+$"
	numericRegexString               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString                = "^[0-9]+$"
	hexadecimalRegexString           = "^[0-9a-fA-F]+$"
	hexcolorRegexString              = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	rgbRegexString                   = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaRegexString                  = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslRegexString                   = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegexString                  = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	emailRegexString                 = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	e164RegexString                  = "^\\+[1-9]?[0-9]{7,14}$"
	base64RegexString                = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLRegexString             = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	iSBN10RegexString                = "^(?:[0-9]{9}X|[0-9]{10})$"
	iSBN13RegexString                = "^(?:(?:97(?:8|9))[0-9]{10})$"
	uUID3RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID4RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID5RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUIDRegexString                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID3RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uUID4RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUID5RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUIDRFC4122RegexString           = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	aSCIIRegexString                 = "^[\x00-\x7F]*$"
	printableASCIIRegexString        = "^[\x20-\x7E]*$"
	multibyteRegexString             = "[^\x00-\x7F]"
	dataURIRegexString               = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	latitudeRegexString              = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegexString             = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	sSNRegexString                   = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	hostnameRegexStringRFC952        = `^[a-zA-Z][a-zA-Z0-9\-\.]+[a-zA-Z0-9]$`                                            // https://tools.ietf.org/html/rfc952
	hostnameRegexStringRFC1123       = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$` // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	btcAddressRegexString            = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`                                                // bitcoin address
	btcAddressUpperRegexStringBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`                                                      // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	btcAddressLowerRegexStringBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`                                                      // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ethAddressRegexString            = `^0x[0-9a-fA-F]{40}$`
	ethAddressUpperRegexString       = `^0x[0-9A-F]{40}$`
	ethAddressLowerRegexString       = `^0x[0-9a-f]{40}$`
	uRLEncodedRegexString            = `(%[A-Fa-f0-9]{2})`
	hTMLEncodedRegexString           = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	hTMLRegexString                  = `<[/]?([a-zA-Z]+).*?>`
	splitParamsRegexString           = `'[^']*'|\S+`
)

var (
	regExpAlphaRegex                 = regexp.MustCompile(alphaRegexString)
	regExpAlphaNumericRegex          = regexp.MustCompile(alphaNumericRegexString)
	regExpAlphaUnicodeRegex          = regexp.MustCompile(alphaUnicodeRegexString)
	regExpAlphaUnicodeNumericRegex   = regexp.MustCompile(alphaUnicodeNumericRegexString)
	regExpNumericRegex               = regexp.MustCompile(numericRegexString)
	regExpNumberRegex                = regexp.MustCompile(numberRegexString)
	regExpNexadecimalRegex           = regexp.MustCompile(hexadecimalRegexString)
	regExpNexcolorRegex              = regexp.MustCompile(hexcolorRegexString)
	regExpRgbRegex                   = regexp.MustCompile(rgbRegexString)
	regExpRgbaRegex                  = regexp.MustCompile(rgbaRegexString)
	regExpHslRegex                   = regexp.MustCompile(hslRegexString)
	regExpHslaRegex                  = regexp.MustCompile(hslaRegexString)
	regExpE164Regex                  = regexp.MustCompile(e164RegexString)
	regExpEmailRegex                 = regexp.MustCompile(emailRegexString)
	regExpBase64Regex                = regexp.MustCompile(base64RegexString)
	regExpBase64URLRegex             = regexp.MustCompile(base64URLRegexString)
	regExpISBN10Regex                = regexp.MustCompile(iSBN10RegexString)
	regExpISBN13Regex                = regexp.MustCompile(iSBN13RegexString)
	regExpUUID3Regex                 = regexp.MustCompile(uUID3RegexString)
	regExpUUID4Regex                 = regexp.MustCompile(uUID4RegexString)
	regExpUUID5Regex                 = regexp.MustCompile(uUID5RegexString)
	regExpUUIDRegex                  = regexp.MustCompile(uUIDRegexString)
	regExpUUID3RFC4122Regex          = regexp.MustCompile(uUID3RFC4122RegexString)
	regExpUUID4RFC4122Regex          = regexp.MustCompile(uUID4RFC4122RegexString)
	regExpUUID5RFC4122Regex          = regexp.MustCompile(uUID5RFC4122RegexString)
	regExpUUIDRFC4122Regex           = regexp.MustCompile(uUIDRFC4122RegexString)
	regExpASCIIRegex                 = regexp.MustCompile(aSCIIRegexString)
	regExpASCIIPrintableRegex        = regexp.MustCompile(printableASCIIRegexString)
	regExpMultibyteRegex             = regexp.MustCompile(multibyteRegexString)
	regExpDataURIRegex               = regexp.MustCompile(dataURIRegexString)
	regExpLatitudeRegex              = regexp.MustCompile(latitudeRegexString)
	regExpLongitudeRegex             = regexp.MustCompile(longitudeRegexString)
	regExpSSNRegex                   = regexp.MustCompile(sSNRegexString)
	regExpHostnameRegexRFC952        = regexp.MustCompile(hostnameRegexStringRFC952)
	regExpHostnameRegexRFC1123       = regexp.MustCompile(hostnameRegexStringRFC1123)
	regExpBtcAddressRegex            = regexp.MustCompile(btcAddressRegexString)
	regExpBtcUpperAddressRegexBech32 = regexp.MustCompile(btcAddressUpperRegexStringBech32)
	regExpBtcLowerAddressRegexBech32 = regexp.MustCompile(btcAddressLowerRegexStringBech32)
	regExpEthAddressRegex            = regexp.MustCompile(ethAddressRegexString)
	regExpEthaddressRegexUpper       = regexp.MustCompile(ethAddressUpperRegexString)
	regExpEthAddressRegexLower       = regexp.MustCompile(ethAddressLowerRegexString)
	regExpURLEncodedRegex            = regexp.MustCompile(uRLEncodedRegexString)
	regExpHTMLEncodedRegex           = regexp.MustCompile(hTMLEncodedRegexString)
	regExpHTMLRegex                  = regexp.MustCompile(hTMLRegexString)
	regExpSplitParamsRegex           = regexp.MustCompile(splitParamsRegexString)
)
