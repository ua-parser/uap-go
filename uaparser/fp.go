package uaparser

import "fmt"

var fpDict = map[string]*Client{
	// Windows, Chrome
	"MozillaWindowsNTAppleWebKitKHTMLlikeGeckoChromeSafari":              {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Chrome", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTWinxAppleWebKitKHTMLlikeGeckoChromeSafari":          {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Chrome", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoChromeSafari":           {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Chrome", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoChromeTaoBrowserSafari": {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Chrome", "", "", ""}, Device: &Device{"Other", "", ""}},
	// Windows, Firefox
	"MozillaWindowsNTWinxrvGeckoFirefox":  {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Firefox Beta", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTWOWrvbGeckoFirefoxb": {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Firefox Beta", "", "", ""}, Device: &Device{"Other", "", ""}},
	// Windows, Edge/IE
	"MozillaWindowsNTWOWTridentrvlikeGecko":                                               {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"IE", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillacompatibleMSIEWindowsNTTrident":                                               {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"IE", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillacompatibleMSIEWindowsNTNETCLRNETCLR":                                          {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"IE", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillacompatibleMSIEWindowsNTSVNETCLR":                                              {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"IE", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillacompatibleMSIEWindowsNT":                                                      {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"IE", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillacompatibleMSIEWindowsNTTencentTraveler":                                       {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"IE", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTWinxAppleWebKitKHTMLlikeGeckoChromeSafariEdg":                        {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Edge", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTWinxAppleWebKitKHTMLlikeGeckoChromeSafariEdge":                       {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Edge", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillacompatibleMSIEWindowsNTWOWTridentSLCCNETCLRNETCLRNETCLRMediaCenterPCNETCNETE": {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"IE", "", "", ""}, Device: &Device{"Other", "", ""}},
	// Windows, Sogou
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoChromeSafariSEXMetaSr": {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Sogou Explorer", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTAppleWebKitKHTMLlikeGeckoChromeSafariSEXMetaSr":    {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Sogou Explorer", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoMaxthonChromeSafari":   {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Maxthon", "", "", ""}, Device: &Device{"Other", "", ""}},
	// Mac OS X, Chrome
	"MozillaMacintoshIntelMacOSXAppleWebKitKHTMLlikeGeckoChromeSafari": {Os: &Os{"Mac OS X", "", "", "", ""}, UserAgent: &UserAgent{"Chrome", "", "", ""}, Device: &Device{"Mac", "", ""}},
	// Mac OS X, Firefox
	"MozillaMacintoshIntelMacOSXrvGeckoFirefox": {Os: &Os{"Mac OS X", "", "", "", ""}, UserAgent: &UserAgent{"Firefox", "", "", ""}, Device: &Device{"Mac", "", ""}},
	// Mobile
	"MozillaLinuxUAndroidzhcnMIBuildPKQAppleWebKitKHTMLlikeGeckoVersionChromeMobileSafariXiaoMiMiuiBrowser":               {Os: &Os{"Android", "", "", "", ""}, UserAgent: &UserAgent{"MiuiBrowser", "", "", ""}, Device: &Device{"XiaoMi MI 6", "", ""}},
	"MozillaLinuxAndroidNexusBuildMRANAppleWebKitKHTMLlikeGeckoChromeMobileSafari":                                        {Os: &Os{"Android", "", "", "", ""}, UserAgent: &UserAgent{"Chrome Mobile", "", "", ""}, Device: &Device{"Nexus 5", "", ""}},
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoHeadlessChromeSafari":                                                    {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"HeadlessChrome", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoMobileE":                                                  {Os: &Os{"iOS", "", "", "", ""}, UserAgent: &UserAgent{"Mobile Safari UI/WKWebView", "", "", ""}, Device: &Device{"iPhone", "", ""}},
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoVersionMobileESafari":                                     {Os: &Os{"iOS", "", "", "", ""}, UserAgent: &UserAgent{"Mobile Safari", "", "", ""}, Device: &Device{"iPhone", "", ""}},
	"MozillaWindowsNTWOWAppleWebKitKHTMLlikeGeckoChromeSafariMicroMessengerNetTypeWIFIMiniProgramEnvWindowsWindowsWechat": {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Chrome", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoMobileEMicroMessengerxNetTypeWIFILanguagezhCN":            {Os: &Os{"iOS", "", "", "", ""}, UserAgent: &UserAgent{"Mobile Safari UI/WKWebView", "", "", ""}, Device: &Device{"iPhone", "", ""}},
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoMobileEMicroMessengerxNetTypeGLanguagezhCN":               {Os: &Os{"iOS", "", "", "", ""}, UserAgent: &UserAgent{"Mobile Safari UI/WKWebView", "", "", ""}, Device: &Device{"iPhone", "", ""}},
	// customer special
	"MozillaiPhoneCPUiPhoneOSlikeMacOSXAppleWebKitKHTMLlikeGeckoMobileEszairApp": {Os: &Os{"iOS", "", "", "", ""}, UserAgent: &UserAgent{"Mobile Safari UI/WKWebView", "", "", ""}, Device: &Device{"iPhone", "", ""}},
	"MozillaWindowsNTWinxAppleWebKitKHTMLlikeGeckoChromeSafariHutool":            {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Chrome", "", "", ""}, Device: &Device{"Other", "", ""}},
	// tools
	"Java":                    {Os: &Os{"Other", "", "", "", ""}, UserAgent: &UserAgent{"Java", "", "", ""}, Device: &Device{"Spider", "", ""}},
	"ApacheHttpClientjava":    {Os: &Os{"Other", "", "", "", ""}, UserAgent: &UserAgent{"Apache-HttpClient", "", "", ""}, Device: &Device{"Other", "", ""}},
	"pyspiderhttppyspiderorg": {Os: &Os{"Other", "", "", "", ""}, UserAgent: &UserAgent{"pyspider", "", "", ""}, Device: &Device{"Spider", "", ""}},
	"axios":                   {Os: &Os{"Other", "", "", "", ""}, UserAgent: &UserAgent{"axios", "", "", ""}, Device: &Device{"Other", "", ""}},
	// test cases 10
	"MozillaWindowsNTrvbpreGeckoFirefoxbpre":                   {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Firefox Beta", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsNTrvGeckoFirefox":                           {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Firefox", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaWindowsUWindowsNTfrrvGeckoFirefoxGTB":              {Os: &Os{"Windows", "", "", "", ""}, UserAgent: &UserAgent{"Other", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaAndroidLinuxarmvlrvbpreGeckoFirefoxbpreFennecbpre": {Os: &Os{"Android", "", "", "", ""}, UserAgent: &UserAgent{"Firefox Mobile", "", "", ""}, Device: &Device{"Generic Smartphone", "Generic Smartphone", ""}},
	"MozillaXULinuxxenUSrvGeckoThunderbird":                    {Os: &Os{"Linux", "", "", "", ""}, UserAgent: &UserAgent{"Thunderbird", "", "", ""}, Device: &Device{"Other", "", ""}},
	"MozillaMacintoshIntelMacOSXrvGeckoFirefoxSeaMonkey":       {Os: &Os{"Mac OS X", "", "", "", ""}, UserAgent: &UserAgent{"SeaMonkey", "", "", ""}, Device: &Device{"Mac", "Apple", "Mac"}},
	"ACSNFACSNFNECc":                         {Os: &Os{"Other", "", "", "", ""}, UserAgent: &UserAgent{"Other", "", "", ""}, Device: &Device{"Generic Feature Phone", "Generic Feature Phone", ""}},
	"BlackBerryBlackBerryUPLink":             {Os: &Os{"BlackBerry OS", "", "", "", ""}, UserAgent: &UserAgent{"BlackBerry", "", "", ""}, Device: &Device{"BlackBerry 7730", "BlackBerry 7730", ""}},
	"GaleonMozillaXULinuxienUSrvGeckoGaleon": {Os: &Os{"Linux", "", "", "", ""}, UserAgent: &UserAgent{"Galeon", "", "", ""}, Device: &Device{"Other", "", ""}},
	"OperaJMEMIDPOperaMiniUidPrestoVersion":  {Os: &Os{"Other", "", "", "", ""}, UserAgent: &UserAgent{"Opera Mini", "", "", ""}, Device: &Device{"Generic Feature Phone", "Generic Feature Phone", ""}},
}

func (parser *Parser) GetFp(line string) string {
	fp := make([]byte, 0)
	for i, v := range line {
		if v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' {
			fp = append(fp, line[i])
		}
	}
	return string(fp)
}

func (parser *Parser) ParseFp(line string) *Client {
	fp := parser.GetFp(line)
	if v, ok := fpDict[fp]; ok {
		return v
	}
	client := parser.Parse(line)
	fmt.Printf("MISS %s: { %s, %s, %s } -> %s\n", fp, client.Os.Family, client.UserAgent.Family, client.Device.Family, line)
	return client
}
