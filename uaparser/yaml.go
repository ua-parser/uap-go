package uaparser

var DefinitionYaml = []byte(`user_agent_parsers:
- regex: (Chromium|Chrome)/(\d+)\.(\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)/(\d+)\.(\d+)(pre|[ab]\d+[a-z]*|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Version)/(\d+)\.(\d+)(?:\.(\d+)|).*Safari/
  regex_flag: ""
  family_replacement: Safari
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Trident)/(7|8)\.(0)
  regex_flag: ""
  family_replacement: IE
  v1_replacement: "11"
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(.*)-iPad\/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)(?:\.(\d+)|) CFNetwork
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(.*)-iPhone/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)(?:\.(\d+)|) CFNetwork
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(.*)/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)(?:\.(\d+)|) CFNetwork
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (espn\.go)
  regex_flag: ""
  family_replacement: ESPN
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (espnradio\.com)
  regex_flag: ""
  family_replacement: ESPN
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ESPN APP$
  regex_flag: ""
  family_replacement: ESPN
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (audioboom\.com)
  regex_flag: ""
  family_replacement: AudioBoom
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ' (Rivo) RHYTHM'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (CFNetwork)(?:/(\d+)\.(\d+)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: CFNetwork
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Pingdom\.com_bot_version_)(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: PingdomBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PingdomTMS)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: PingdomBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ' (PTST)/(\d+)(?:\.(\d+)|)$'
  regex_flag: ""
  family_replacement: WebPageTest.org bot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: X11; (Datanyze); Linux
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (NewRelicPinger)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: NewRelicPingerBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Tableau)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Tableau
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Salesforce)(?:.)\/(\d+)\.(\d?)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (\(StatusCake\))
  regex_flag: ""
  family_replacement: StatusCakeBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (facebookexternalhit)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: FacebookBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Google.*/\+/web/snippet
  regex_flag: ""
  family_replacement: GooglePlusBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: via ggpht\.com GoogleImageProxy
  regex_flag: ""
  family_replacement: GmailImageProxy
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: YahooMailProxy; https://help\.yahoo\.com/kb/yahoo-mail-proxy-SLN28749\.html
  regex_flag: ""
  family_replacement: YahooMailProxy
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Twitterbot)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: TwitterBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: /((?:Ant-|)Nutch|[A-z]+[Bb]ot|[A-z]+[Ss]pider|Axtaris|fetchurl|Isara|ShopSalad|Tailsweep)[
    \-](\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \b(008|Altresium|Argus|BaiduMobaider|BoardReader|DNSGroup|DataparkSearch|EDI|Goodzer|Grub|INGRID|Infohelfer|LinkedInBot|LOOQ|Nutch|PathDefender|Peew|PostPost|Steeler|Twitterbot|VSE|WebCrunch|WebZIP|Y!J-BR[A-Z]|YahooSeeker|envolk|sproose|wminer)/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MSIE) (\d+)\.(\d+)([a-z]\d|[a-z]|);.* MSIECrawler
  regex_flag: ""
  family_replacement: MSIECrawler
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (DAVdroid)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Google-HTTP-Java-Client|Apache-HttpClient|Go-http-client|scalaj-http|http%20client|Python-urllib|HttpMonitor|TLSProber|WinHTTP|JNLP|okhttp|aihttp|reqwest)(?:[
    /](\d+)(?:\.(\d+)|)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Pinterest(?:bot|))/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)[;\s(]+\+https://www.pinterest.com/bot.html
  regex_flag: ""
  family_replacement: Pinterestbot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '(CSimpleSpider|Cityreview Robot|CrawlDaddy|CrawlFire|Finderbots|Index crawler|Job
    Roboter|KiwiStatus Spider|Lijit Crawler|QuerySeekerSpider|ScollSpider|Trends Crawler|USyd-NLP-Spider|SiteCat
    Webbot|BotName\/\$BotVersion|123metaspider-Bot|1470\.net crawler|50\.nu|8bo Crawler
    Bot|Aboundex|Accoona-[A-z]{1,30}-Agent|AdsBot-Google(?:-[a-z]{1,30}|)|altavista|AppEngine-Google|archive.{0,30}\.org_bot|archiver|Ask
    Jeeves|[Bb]ai[Dd]u[Ss]pider(?:-[A-Za-z]{1,30})(?:-[A-Za-z]{1,30}|)|bingbot|BingPreview|blitzbot|BlogBridge|Bloglovin|BoardReader
    Blog Indexer|BoardReader Favicon Fetcher|boitho.com-dc|BotSeer|BUbiNG|\b\w{0,30}favicon\w{0,30}\b|\bYeti(?:-[a-z]{1,30}|)|Catchpoint(?:
    bot|)|[Cc]harlotte|Checklinks|clumboot|Comodo HTTP\(S\) Crawler|Comodo-Webinspector-Crawler|ConveraCrawler|CRAWL-E|CrawlConvera|Daumoa(?:-feedfetcher|)|Feed
    Seeker Bot|Feedbin|findlinks|Flamingo_SearchEngine|FollowSite Bot|furlbot|Genieo|gigabot|GomezAgent|gonzo1|(?:[a-zA-Z]{1,30}-|)Googlebot(?:-[a-zA-Z]{1,30}|)|Google
    SketchUp|grub-client|gsa-crawler|heritrix|HiddenMarket|holmes|HooWWWer|htdig|ia_archiver|ICC-Crawler|Icarus6j|ichiro(?:/mobile|)|IconSurf|IlTrovatore(?:-Setaccio|)|InfuzApp|Innovazion
    Crawler|InternetArchive|IP2[a-z]{1,30}Bot|jbot\b|KaloogaBot|Kraken|Kurzor|larbin|LEIA|LesnikBot|Linguee
    Bot|LinkAider|LinkedInBot|Lite Bot|Llaut|lycos|Mail\.RU_Bot|masscan|masidani_bot|Mediapartners-Google|Microsoft
    .{0,30} Bot|mogimogi|mozDex|MJ12bot|msnbot(?:-media {0,2}|)|msrbot|Mtps Feed Aggregation
    System|netresearch|Netvibes|NewsGator[^/]{0,30}|^NING|Nutch[^/]{0,30}|Nymesis|ObjectsSearch|Orbiter|OOZBOT|PagePeeker|PagesInventory|PaxleFramework|Peeplo
    Screenshot Bot|PlantyNet_WebRobot|Pompos|Qwantify|Read%20Later|Reaper|RedCarpet|Retreiver|Riddler|Rival
    IQ|scooter|Scrapy|Scrubby|searchsight|seekbot|semanticdiscovery|SemrushBot|Simpy|SimplePie|SEOstats|SimpleRSS|SiteCon|Slackbot-LinkExpanding|Slack-ImgProxy|Slurp|snappy|Speedy
    Spider|Squrl Java|Stringer|TheUsefulbot|ThumbShotsBot|Thumbshots\.ru|Tiny Tiny
    RSS|TwitterBot|WhatsApp|URL2PNG|Vagabondo|VoilaBot|^vortex|Votay bot|^voyager|WASALive.Bot|Web-sniffer|WebThumb|WeSEE:[A-z]{1,30}|WhatWeb|WIRE|WordPress|Wotbox|www\.almaden\.ibm\.com|Xenu(?:.s|)
    Link Sleuth|Xerka [A-z]{1,30}Bot|yacy(?:bot|)|YahooSeeker|Yahoo! Slurp|Yandex\w{1,30}|YodaoBot(?:-[A-z]{1,30}|)|YottaaMonitor|Yowedo|^Zao|^Zao-Crawler|ZeBot_www\.ze\.bz|ZooShot|ZyBorg)(?:[
    /]v?(\d+)(?:\.(\d+)(?:\.(\d+)|)|)|)'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \b(Boto3?|JetS3t|aws-(?:cli|sdk-(?:cpp|go|java|nodejs|ruby2?))|s3fs)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \[(FBAN/MessengerForiOS|FB_IAB/MESSENGER);FBAV/(\d+)(?:\.(\d+)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: Facebook Messenger
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \[FB.*;(FBAV)/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Facebook
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \[FB.*;
  regex_flag: ""
  family_replacement: Facebook
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (?:\/[A-Za-z0-9\.]+|) {0,5}([A-Za-z0-9 \-_\!\[\]:]{0,50}(?:[Aa]rchiver|[Ii]ndexer|[Ss]craper|[Bb]ot|[Ss]pider|[Cc]rawl[a-z]{0,50}))[/
    ](\d+)(?:\.(\d+)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ((?:[A-Za-z][A-Za-z0-9 -]{0,50}|)[^C][^Uu][Bb]ot)\b(?:(?:[ /]| v)(\d+)(?:\.(\d+)|)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '((?:[A-z0-9]{1,50}|[A-z\-]{1,50} ?|)(?: the |)(?:[Ss][Pp][Ii][Dd][Ee][Rr]|[Ss]crape|[Cc][Rr][Aa][Ww][Ll])[A-z0-9]{0,50})(?:(?:[
    /]| v)(\d+)(?:\.(\d+)|)(?:\.(\d+)|)|)'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (HbbTV)/(\d+)\.(\d+)\.(\d+) \(
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Chimera|SeaMonkey|Camino|Waterfox)/(\d+)\.(\d+)\.?([ab]?\d+[a-z]*|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (SailfishBrowser)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Sailfish Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \[(Pinterest)/[^\]]+\]
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '(Pinterest)(?: for Android(?: Tablet|)|)/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Mozilla.*Mobile.*(Instagram).(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Mozilla.*Mobile.*(Flipboard).(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Mozilla.*Mobile.*(Flipboard-Briefing).(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Mozilla.*Mobile.*(Onefootball)\/Android.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Snapchat)\/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)/(\d+)\.(\d+) Basilisk/(\d+)
  regex_flag: ""
  family_replacement: Basilisk
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PaleMoon)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Pale Moon
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Fennec)/(\d+)\.(\d+)\.?([ab]?\d+[a-z]*)
  regex_flag: ""
  family_replacement: Firefox Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Fennec)/(\d+)\.(\d+)(pre)
  regex_flag: ""
  family_replacement: Firefox Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Fennec)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Firefox Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (?:Mobile|Tablet);.*(Firefox)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Firefox Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Namoroka|Shiretoko|Minefield)/(\d+)\.(\d+)\.(\d+(?:pre|))
  regex_flag: ""
  family_replacement: Firefox ($1)
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)/(\d+)\.(\d+)(a\d+[a-z]*)
  regex_flag: ""
  family_replacement: Firefox Alpha
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)/(\d+)\.(\d+)(b\d+[a-z]*)
  regex_flag: ""
  family_replacement: Firefox Beta
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)-(?:\d+\.\d+|)/(\d+)\.(\d+)(a\d+[a-z]*)
  regex_flag: ""
  family_replacement: Firefox Alpha
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)-(?:\d+\.\d+|)/(\d+)\.(\d+)(b\d+[a-z]*)
  regex_flag: ""
  family_replacement: Firefox Beta
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Namoroka|Shiretoko|Minefield)/(\d+)\.(\d+)([ab]\d+[a-z]*|)
  regex_flag: ""
  family_replacement: Firefox ($1)
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox).*Tablet browser (\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: MicroB
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MozillaDeveloperPreview)/(\d+)\.(\d+)([ab]\d+[a-z]*|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (FxiOS)/(\d+)\.(\d+)(\.(\d+)|)(\.(\d+)|)
  regex_flag: ""
  family_replacement: Firefox iOS
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Flock)/(\d+)\.(\d+)(b\d+?)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (RockMelt)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Navigator)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Netscape
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Navigator)/(\d+)\.(\d+)([ab]\d+)
  regex_flag: ""
  family_replacement: Netscape
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Netscape6)/(\d+)\.(\d+)\.?([ab]?\d+|)
  regex_flag: ""
  family_replacement: Netscape
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MyIBrow)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: My Internet Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (UC? ?Browser|UCWEB|U3)[ /]?(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: UC Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Opera Tablet).*Version/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Opera Mini)(?:/att|)/?(\d+|)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Opera)/.+Opera Mobi.+Version/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Opera Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Opera)/(\d+)\.(\d+).+Opera Mobi
  regex_flag: ""
  family_replacement: Opera Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Opera Mobi.+(Opera)(?:/|\s+)(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Opera Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Opera Mobi
  regex_flag: ""
  family_replacement: Opera Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Opera)/9.80.*Version/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (?:Mobile Safari).*(OPR)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Opera Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (?:Chrome).*(OPR)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Opera
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Coast)/(\d+).(\d+).(\d+)
  regex_flag: ""
  family_replacement: Opera Coast
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (OPiOS)/(\d+).(\d+).(\d+)
  regex_flag: ""
  family_replacement: Opera Mini
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Chrome/.+( MMS)/(\d+).(\d+).(\d+)
  regex_flag: ""
  family_replacement: Opera Neon
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (hpw|web)OS/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: webOS Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (luakit)
  regex_flag: ""
  family_replacement: LuaKit
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Snowshoe)/(\d+)\.(\d+).(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Gecko/\d+ (Lightning)/(\d+)\.(\d+)\.?((?:[ab]?\d+[a-z]*)|(?:\d*))
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)/(\d+)\.(\d+)\.(\d+(?:pre|)) \(Swiftfox\)
  regex_flag: ""
  family_replacement: Swiftfox
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)/(\d+)\.(\d+)([ab]\d+[a-z]*|) \(Swiftfox\)
  regex_flag: ""
  family_replacement: Swiftfox
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (rekonq)/(\d+)\.(\d+)(?:\.(\d+)|) Safari
  regex_flag: ""
  family_replacement: Rekonq
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: rekonq
  regex_flag: ""
  family_replacement: Rekonq
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (conkeror|Conkeror)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Conkeror
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (konqueror)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Konqueror
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (WeTab)-Browser
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Comodo_Dragon)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Comodo Dragon
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Symphony) (\d+).(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: PLAYSTATION 3.+WebKit
  regex_flag: ""
  family_replacement: NetFront NX
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: PLAYSTATION 3
  regex_flag: ""
  family_replacement: NetFront
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PlayStation Portable)
  regex_flag: ""
  family_replacement: NetFront
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PlayStation Vita)
  regex_flag: ""
  family_replacement: NetFront NX
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: AppleWebKit.+ (NX)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: NetFront NX
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Nintendo 3DS)
  regex_flag: ""
  family_replacement: NetFront NX
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Silk)/(\d+)\.(\d+)(?:\.([0-9\-]+)|)
  regex_flag: ""
  family_replacement: Amazon Silk
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Puffin)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Windows Phone .*(Edge)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Edge Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (SamsungBrowser)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Samsung Internet
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (SznProhlizec)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Seznam prohlížeč
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (coc_coc_browser)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Coc Coc
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (baidubrowser)[/\s](\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Baidu Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (FlyFlow)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Baidu Explorer
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MxBrowser)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Maxthon
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Crosswalk)/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Line)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: LINE
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MiuiBrowser)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: MiuiBrowser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Mint Browser)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Mint Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Mozilla.+Android.+(GSA)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Google
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Version/.+(Chrome)/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Chrome Mobile WebView
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ; wv\).+(Chrome)/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Chrome Mobile WebView
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (CrMo)/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Chrome Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (CriOS)/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Chrome Mobile iOS
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Chrome)/(\d+)\.(\d+)\.(\d+)\.(\d+) Mobile(?:[ /]|$)
  regex_flag: ""
  family_replacement: Chrome Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ' Mobile .*(Chrome)/(\d+)\.(\d+)\.(\d+)\.(\d+)'
  regex_flag: ""
  family_replacement: Chrome Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (chromeframe)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Chrome Frame
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (SLP Browser)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Tizen Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (SE 2\.X) MetaSr (\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Sogou Explorer
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MQQBrowser/Mini)(?:(\d+)(?:\.(\d+)|)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: QQ Browser Mini
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MQQBrowser)(?:/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: QQ Browser Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (QQBrowser)(?:/(\d+)(?:\.(\d+)\.(\d+)(?:\.(\d+)|)|)|)
  regex_flag: ""
  family_replacement: QQ Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Rackspace Monitoring)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: RackspaceBot
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PyAMF)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (YaBrowser)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Yandex Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Chrome)/(\d+)\.(\d+)\.(\d+).* MRCHROME
  regex_flag: ""
  family_replacement: Mail.ru Chromium Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (AOL) (\d+)\.(\d+); AOLBuild (\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PodCruncher|Downcast)[ /]?(\d+)(?:\.(\d+)|)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ' (BoxNotes)/(\d+)\.(\d+)\.(\d+)'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Whale)/(\d+)\.(\d+)\.(\d+)\.(\d+) Mobile(?:[ /]|$)
  regex_flag: ""
  family_replacement: Whale
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Whale)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Whale
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Ghost)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Slack_SSB)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Slack Desktop Client
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (HipChat)/?(\d+|)
  regex_flag: ""
  family_replacement: HipChat Desktop Client
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \b(MobileIron|FireWeb|Jasmine|ANTGalio|Midori|Fresco|Lobo|PaleMoon|Maxthon|Lynx|OmniWeb|Dillo|Camino|Demeter|Fluid|Fennec|Epiphany|Shiira|Sunrise|Spotify|Flock|Netscape|Lunascape|WebPilot|NetFront|Netfront|Konqueror|SeaMonkey|Kazehakase|Vienna|Iceape|Iceweasel|IceWeasel|Iron|K-Meleon|Sleipnir|Galeon|GranParadiso|Opera
    Mini|iCab|NetNewsWire|ThunderBrowse|Iris|UP\.Browser|Bunjalloo|Google Earth|Raven
    for Mac|Openwave|MacOutlook|Electron|OktaMobile)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Microsoft Office Outlook 12\.\d+\.\d+|MSOffice 12
  regex_flag: ""
  family_replacement: Outlook
  v1_replacement: "2007"
  v2_replacement: $3
  v3_replacement: $4
- regex: Microsoft Outlook 14\.\d+\.\d+|MSOffice 14
  regex_flag: ""
  family_replacement: Outlook
  v1_replacement: "2010"
  v2_replacement: $3
  v3_replacement: $4
- regex: Microsoft Outlook 15\.\d+\.\d+
  regex_flag: ""
  family_replacement: Outlook
  v1_replacement: "2013"
  v2_replacement: $3
  v3_replacement: $4
- regex: Microsoft Outlook (?:Mail )?16\.\d+\.\d+
  regex_flag: ""
  family_replacement: Outlook
  v1_replacement: "2016"
  v2_replacement: $3
  v3_replacement: $4
- regex: Microsoft Office (Word) 2014
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: Outlook-Express\/7\.0.*
  regex_flag: ""
  family_replacement: Windows Live Mail
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Airmail) (\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Thunderbird)/(\d+)\.(\d+)(?:\.(\d+(?:pre|))|)
  regex_flag: ""
  family_replacement: Thunderbird
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Postbox)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Postbox
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Barca(?:Pro)?)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Barca
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Lotus-Notes)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Lotus Notes
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Vivaldi)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Edge?)/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Edge
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (brave)/(\d+)\.(\d+)\.(\d+) Chrome
  regex_flag: ""
  family_replacement: Brave
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Chrome)/(\d+)\.(\d+)\.(\d+)[\d.]* Iron[^/]
  regex_flag: ""
  family_replacement: Iron
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '\b(Dolphin)(?: |HDCN/|/INT\-)(\d+)\.(\d+)(?:\.(\d+)|)'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (HeadlessChrome)(?:/(\d+)\.(\d+)\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Evolution)/(\d+)\.(\d+)\.(\d+\.\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (RCM CardDAV plugin)/(\d+)\.(\d+)\.(\d+(?:-dev|))
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (bingbot|Bolt|AdobeAIR|Jasmine|IceCat|Skyfire|Midori|Maxthon|Lynx|Arora|IBrowse|Dillo|Camino|Shiira|Fennec|Phoenix|Flock|Netscape|Lunascape|Epiphany|WebPilot|Opera
    Mini|Opera|NetFront|Netfront|Konqueror|Googlebot|SeaMonkey|Kazehakase|Vienna|Iceape|Iceweasel|IceWeasel|Iron|K-Meleon|Sleipnir|Galeon|GranParadiso|iCab|iTunes|MacAppStore|NetNewsWire|Space
    Bison|Stainless|Orca|Dolfin|BOLT|Minimo|Tizen Browser|Polaris|Abrowser|Planetweb|ICE
    Browser|mDolphin|qutebrowser|Otter|QupZilla|MailBar|kmail2|YahooMobileMail|ExchangeWebServices|ExchangeServicesClient|Dragon|Outlook-iOS-Android)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (ESPN)[%20| ]+Radio/(\d+)\.(\d+)\.(\d+) CFNetwork
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (IEMobile)[ /](\d+)\.(\d+)
  regex_flag: ""
  family_replacement: IE Mobile
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (BacaBerita App)\/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(bPod|Pocket Casts|Player FM)$
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(AlexaMediaPlayer|VLC)/(\d+)\.(\d+)\.([^.\s]+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(AntennaPod|WMPlayer|Zune|Podkicker|Radio|ExoPlayerDemo|Overcast|PocketTunes|NSPlayer|okhttp|DoggCatcher|QuickNews|QuickTime|Peapod|Podcasts|GoldenPod|VLC|Spotify|Miro|MediaGo|Juice|iPodder|gPodder|Banshee)/(\d+)\.(\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Peapod|Liferea)/([^.\s]+)\.([^.\s]+|)\.?([^.\s]+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(bPod|Player FM) BMID/(\S+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '^(Podcast ?Addict)/v(\d+) '
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '^(Podcast ?Addict) '
  regex_flag: ""
  family_replacement: PodcastAddict
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Replay) AV
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (VOX) Music Player
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (CITA) RSS Aggregator/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Pocket Casts)$
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Player FM)$
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (LG Player|Doppler|FancyMusic|MediaMonkey|Clementine) (\d+)\.(\d+)\.?([^.\s]+|)\.?([^.\s]+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (philpodder)/(\d+)\.(\d+)\.?([^.\s]+|)\.?([^.\s]+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Player FM|Pocket Casts|DoggCatcher|Spotify|MediaMonkey|MediaGo|BashPodder)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (QuickTime)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Kinoma)(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Fancy) Cloud Music (\d+)\.(\d+)
  regex_flag: ""
  family_replacement: FancyMusic
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: EspnDownloadManager
  regex_flag: ""
  family_replacement: ESPN
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '(ESPN) Radio (\d+)\.(\d+)(?:\.(\d+)|) ?(?:rv:(\d+)|) '
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (podracer|jPodder) v ?(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (ZDM)/(\d+)\.(\d+)[; ]?
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Zune|BeyondPod) (\d+)(?:\.(\d+)|)[\);]
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (WMPlayer)/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Lavf)
  regex_flag: ""
  family_replacement: WMPlayer
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(RSSRadio)[ /]?(\d+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (RSS_Radio) (\d+)\.(\d+)
  regex_flag: ""
  family_replacement: RSSRadio
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Podkicker) \S+/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Podkicker
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(HTC) Streaming Player \S+ / \S+ / \S+ / (\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Stitcher)/iOS
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Stitcher)/Android
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(VLC) .*version (\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ' (VLC) for'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (vlc)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: VLC
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(foobar)\S+/([^.\s]+)\.([^.\s]+|)\.?([^.\s]+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Clementine)\S+ ([^.\s]+)\.([^.\s]+|)\.?([^.\s]+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (amarok)/([^.\s]+)\.([^.\s]+|)\.?([^.\s]+|)
  regex_flag: ""
  family_replacement: Amarok
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Custom)-Feed Reader
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iRider|Crazy Browser|SkipStone|iCab|Lunascape|Sleipnir|Maemo Browser) (\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iCab|Lunascape|Opera|Android|Jasmine|Polaris|Microsoft SkyDriveSync|The
    Bat!) (\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Kindle)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Android) Donut
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "1"
  v2_replacement: "2"
  v3_replacement: $4
- regex: (Android) Eclair
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "2"
  v2_replacement: "1"
  v3_replacement: $4
- regex: (Android) Froyo
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "2"
  v2_replacement: "2"
  v3_replacement: $4
- regex: (Android) Gingerbread
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "2"
  v2_replacement: "3"
  v3_replacement: $4
- regex: (Android) Honeycomb
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "3"
  v2_replacement: $3
  v3_replacement: $4
- regex: (MSIE) (\d+)\.(\d+).*XBLWP7
  regex_flag: ""
  family_replacement: IE Large Screen
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Nextcloud)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (mirall)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (ownCloud-android)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Owncloud
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (OC)/(\d+)\.(\d+)\.(\d+)\.(\d+) \(Skype for Business\)
  regex_flag: ""
  family_replacement: Skype
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Obigo)InternetBrowser
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Obigo)\-Browser
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Obigo|OBIGO)[^\d]*(\d+)(?:.(\d+)|)
  regex_flag: ""
  family_replacement: Obigo
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MAXTHON|Maxthon) (\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Maxthon
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Maxthon|MyIE2|Uzbl|Shiira)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "0"
  v2_replacement: $3
  v3_replacement: $4
- regex: (BrowseX) \((\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (NCSA_Mosaic)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: NCSA Mosaic
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (POLARIS)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Polaris
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Embider)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Polaris
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (BonEcho)/(\d+)\.(\d+)\.?([ab]?\d+|)
  regex_flag: ""
  family_replacement: Bon Echo
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPhone|iPad).+GSA/(\d+)\.(\d+)\.(\d+) Mobile
  regex_flag: ""
  family_replacement: Google
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPhone|iPad).+Version/(\d+)\.(\d+)(?:\.(\d+)|).*[ +]Safari
  regex_flag: ""
  family_replacement: Mobile Safari
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPod touch|iPhone|iPad);.*CPU.*OS[ +](\d+)_(\d+)(?:_(\d+)|).* AppleNews\/\d+\.\d+\.\d+?
  regex_flag: ""
  family_replacement: Mobile Safari UI/WKWebView
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPhone|iPad).+Version/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: Mobile Safari UI/WKWebView
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPod touch|iPhone|iPad);.*CPU.*OS[ +](\d+)_(\d+)(?:_(\d+)|).*Mobile.*[
    +]Safari
  regex_flag: ""
  family_replacement: Mobile Safari
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPod touch|iPhone|iPad);.*CPU.*OS[ +](\d+)_(\d+)(?:_(\d+)|).*Mobile
  regex_flag: ""
  family_replacement: Mobile Safari UI/WKWebView
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPhone|iPad).* Safari
  regex_flag: ""
  family_replacement: Mobile Safari
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (iPod|iPhone|iPad)
  regex_flag: ""
  family_replacement: Mobile Safari UI/WKWebView
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Watch)(\d+),(\d+)
  regex_flag: ""
  family_replacement: Apple $1 App
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Outlook-iOS)/\d+\.\d+\.prod\.iphone \((\d+)\.(\d+)\.(\d+)\)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (AvantGo) (\d+).(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (OneBrowser)/(\d+).(\d+)
  regex_flag: ""
  family_replacement: ONE Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Avant)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "1"
  v2_replacement: $3
  v3_replacement: $4
- regex: (QtCarBrowser)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "1"
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(iBrowser/Mini)(\d+).(\d+)
  regex_flag: ""
  family_replacement: iBrowser Mini
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(iBrowser|iRAPP)/(\d+).(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Nokia)
  regex_flag: ""
  family_replacement: Nokia Services (WAP) Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (NokiaBrowser)/(\d+)\.(\d+).(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Nokia Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (NokiaBrowser)/(\d+)\.(\d+).(\d+)
  regex_flag: ""
  family_replacement: Nokia Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (NokiaBrowser)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Nokia Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (BrowserNG)/(\d+)\.(\d+).(\d+)
  regex_flag: ""
  family_replacement: Nokia Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Series60)/5\.0
  regex_flag: ""
  family_replacement: Nokia Browser
  v1_replacement: "7"
  v2_replacement: "0"
  v3_replacement: $4
- regex: (Series60)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Nokia OSS Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (S40OviBrowser)/(\d+)\.(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Ovi Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Nokia)[EN]?(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PlayBook).+RIM Tablet OS (\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: BlackBerry WebKit
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Black[bB]erry|BB10).+Version/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: BlackBerry WebKit
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Black[bB]erry)\s?(\d+)
  regex_flag: ""
  family_replacement: BlackBerry
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (OmniWeb)/v(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Blazer)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Palm Blazer
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Pre)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Palm Pre
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (ELinks)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (ELinks) \((\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Links) \((\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (QtWeb) Internet Browser/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (PhantomJS)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (AppleWebKit)/(\d+)(?:\.(\d+)|)\+ .* Safari
  regex_flag: ""
  family_replacement: WebKit Nightly
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (TopPodcasts)Pro/(\d+) CFNetwork
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Safari)/\d+
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (OLPC)/Update(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (OLPC)/Update()\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: "0"
  v2_replacement: $3
  v3_replacement: $4
- regex: (SEMC\-Browser)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Teleca)
  regex_flag: ""
  family_replacement: Teleca Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Phantom)/V(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Phantom Browser
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (MusicDownloader)Lite/(\d+)\.(\d+)\.(\d+) CFNetwork
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Trident)/(6)\.(0)
  regex_flag: ""
  family_replacement: IE
  v1_replacement: "10"
  v2_replacement: $3
  v3_replacement: $4
- regex: (Trident)/(5)\.(0)
  regex_flag: ""
  family_replacement: IE
  v1_replacement: "9"
  v2_replacement: $3
  v3_replacement: $4
- regex: (Trident)/(4)\.(0)
  regex_flag: ""
  family_replacement: IE
  v1_replacement: "8"
  v2_replacement: $3
  v3_replacement: $4
- regex: (Espial)/(\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (AppleWebKit)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Apple Mail
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Firefox)/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Antenna)/(\d+) CFNetwork
  regex_flag: ""
  family_replacement: AntennaPod
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ([MS]?IE) (\d+)\.(\d+)
  regex_flag: ""
  family_replacement: IE
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (python-requests)/(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Python Requests
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: \b(Windows-Update-Agent|Microsoft-CryptoAPI|SophosUpdateManager|SophosAgent|Debian
    APT-HTTP|Ubuntu APT-HTTP|libcurl-agent|libwww-perl|urlgrabber|curl|PycURL|Wget|aria2|Axel|OpenBSD
    ftp|lftp|jupdate|insomnia)(?:[ /](\d+)(?:\.(\d+)|)(?:\.(\d+)|)|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Java)[/ ]{0,1}\d+\.(\d+)\.(\d+)[_-]*([a-zA-Z0-9]+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Cyberduck)/(\d+)\.(\d+)\.(\d+)(?:\.\d+|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(S3 Browser) (\d+)-(\d+)-(\d+)(?:\s*http://s3browser\.com|)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(rclone)/v(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(Roku)/DVP-(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: (Kurio)\/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: Kurio App
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: '^(Box(?: Sync)?)/(\d+)\.(\d+)\.(\d+)'
  regex_flag: ""
  family_replacement: $1
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
- regex: ^(ViaFree|Viafree)-(?:tvOS-)?[A-Z]{2}/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  family_replacement: ViaFree
  v1_replacement: $2
  v2_replacement: $3
  v3_replacement: $4
os_parsers:
- reg: {}
  regex: HbbTV/\d+\.\d+\.\d+ \( ;(LG)E ;NetCast 4.0
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2013"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/\d+\.\d+\.\d+ \( ;(LG)E ;NetCast 3.0
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2012"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/1.1.1 \(;;;;;\) Maple_2011
  regex_flag: ""
  os_replacement: Samsung
  os_v1_replacement: "2011"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/\d+\.\d+\.\d+ \(;(Samsung);SmartTV([0-9]{4});.*FXPDEUC
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: UE40F7000
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/\d+\.\d+\.\d+ \(;(Samsung);SmartTV([0-9]{4});.*MST12DEUC
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: UE32F4500
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/1\.1\.1 \(; (Philips);.*NETTV/4
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2013"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/1\.1\.1 \(; (Philips);.*NETTV/3
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2012"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/1\.1\.1 \(; (Philips);.*NETTV/2
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2011"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/\d+\.\d+\.\d+.*(firetv)-firefox-plugin (\d+).(\d+).(\d+)
  regex_flag: ""
  os_replacement: FireHbbTV
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: HbbTV/\d+\.\d+\.\d+ \(.*; ?([a-zA-Z]+) ?;.*(201[1-9]).*\)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows Phone) (?:OS[ /])?(\d+)\.(\d+)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CPU[ +]OS|iPhone[ +]OS|CPU[ +]iPhone)[ +]+(\d+)[_\.](\d+)(?:[_\.](\d+)|).*Outlook-iOS-Android
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Android)[ \-/](\d+)(?:\.(\d+)|)(?:[.\-]([a-z0-9]+)|)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Android) Donut
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "1"
  os_v2_replacement: "2"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Android) Eclair
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2"
  os_v2_replacement: "1"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Android) Froyo
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2"
  os_v2_replacement: "2"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Android) Gingerbread
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "2"
  os_v2_replacement: "3"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Android) Honeycomb
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: "3"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ^UCWEB.*; (Adr) (\d+)\.(\d+)(?:[.\-]([a-z0-9]+)|);
  regex_flag: ""
  os_replacement: Android
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ^UCWEB.*; (iPad|iPh|iPd) OS (\d+)_(\d+)(?:_(\d+)|);
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ^UCWEB.*; (wds) (\d+)\.(\d+)(?:\.(\d+)|);
  regex_flag: ""
  os_replacement: Windows Phone
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ^(JUC).*; ?U; ?(?:Android|)(\d+)\.(\d+)(?:[\.\-]([a-z0-9]+)|)
  regex_flag: ""
  os_replacement: Android
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (android)\s(?:mobile\/)(\d+)(?:\.(\d+)(?:\.(\d+)|)|)
  regex_flag: ""
  os_replacement: Android
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Silk-Accelerated=[a-z]{4,5})
  regex_flag: ""
  os_replacement: Android
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (x86_64|aarch64)\ (\d+)\.(\d+)\.(\d+).*Chrome.*(?:CitrixChromeApp)$
  regex_flag: ""
  os_replacement: Chrome OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (XBLWP7)
  regex_flag: ""
  os_replacement: Windows Phone
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows ?Mobile)
  regex_flag: ""
  os_replacement: Windows Mobile
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows 10)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "10"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows (?:NT 5\.2|NT 5\.1))
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: XP
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 6\.1)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "7"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 6\.0)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: Vista
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Win 9x 4\.90)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: ME
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 6\.2; ARM;)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: RT
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 6\.2)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "8"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 6\.3; ARM;)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: RT 8
  os_v2_replacement: "1"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 6\.3)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "8"
  os_v2_replacement: "1"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 6\.4)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "10"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 10\.0)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "10"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows NT 5\.0)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "2000"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (WinNT4.0)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: NT 4.0
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows ?CE)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: CE
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Win(?:dows)? ?(95|98|3.1|NT|ME|2000|XP|Vista|7|CE)
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: $1
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Win16
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "3.1"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Win32
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: "95"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ^Box.*Windows/([\d.]+);
  regex_flag: ""
  os_replacement: Windows
  os_v1_replacement: $1
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Tizen)[/ ](\d+)\.(\d+)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ((?:Mac[ +]?|; )OS[ +]X)[\s+/](?:(\d+)[_.](\d+)(?:[_.](\d+)|)|Mach-O)
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \w+\s+Mac OS X\s+\w+\s+(\d+).(\d+).(\d+).*
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: $1
  os_v2_replacement: $2
  os_v3_replacement: $3
  os_v4_replacement: $5
- reg: {}
  regex: ' (Dar)(win)/(9).(\d+).*\((?:i386|x86_64|Power Macintosh)\)'
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "5"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ' (Dar)(win)/(10).(\d+).*\((?:i386|x86_64)\)'
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "6"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ' (Dar)(win)/(11).(\d+).*\((?:i386|x86_64)\)'
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "7"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ' (Dar)(win)/(12).(\d+).*\((?:i386|x86_64)\)'
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "8"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ' (Dar)(win)/(13).(\d+).*\((?:i386|x86_64)\)'
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "9"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Mac_PowerPC
  regex_flag: ""
  os_replacement: Mac OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (?:PPC|Intel) (Mac OS X)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ^Box.*;(Darwin)/(10)\.(1\d)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Apple\s?TV)(?:/(\d+)\.(\d+)|)
  regex_flag: ""
  os_replacement: ATV OS X
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CPU[ +]OS|iPhone[ +]OS|CPU[ +]iPhone|CPU IPhone OS)[ +]+(\d+)[_\.](\d+)(?:[_\.](\d+)|)
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (iPhone|iPad|iPod); Opera
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (iPhone|iPad|iPod).*Mac OS X.*Version/(\d+)\.(\d+)
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/(5)48\.0\.3.* Darwin/11\.0\.0
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/(5)48\.(0)\.4.* Darwin/(1)1\.0\.0
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/(5)48\.(1)\.4
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/(4)85\.1(3)\.9
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/(6)09\.(1)\.4
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/(6)(0)9
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/6(7)2\.(1)\.13
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/6(7)2\.(1)\.(1)4
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CF)(Network)/6(7)(2)\.1\.15
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "7"
  os_v2_replacement: "1"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/6(7)2\.(0)\.(?:2|8)
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CFNetwork)/709\.1
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "8"
  os_v2_replacement: 0.b5
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CF)(Network)/711\.(\d)
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "8"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CF)(Network)/(720)\.(\d)
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "10"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CF)(Network)/(760)\.(\d)
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "11"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/7.* Darwin/15\.4\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "9"
  os_v2_replacement: "3"
  os_v3_replacement: "1"
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/7.* Darwin/15\.5\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "9"
  os_v2_replacement: "3"
  os_v3_replacement: "2"
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/7.* Darwin/15\.6\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "9"
  os_v2_replacement: "3"
  os_v3_replacement: "5"
  os_v4_replacement: $5
- reg: {}
  regex: (CF)(Network)/758\.(\d)
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "9"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/808\.3 Darwin/16\.3\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "10"
  os_v2_replacement: "2"
  os_v3_replacement: "1"
  os_v4_replacement: $5
- reg: {}
  regex: (CF)(Network)/808\.(\d)
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "10"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/.* Darwin/17\.\d+.*\(x86_64\)
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "13"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/.* Darwin/16\.\d+.*\(x86_64\)
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "12"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/15\.\d+.*\(x86_64\)
  regex_flag: ""
  os_replacement: Mac OS X
  os_v1_replacement: "10"
  os_v2_replacement: "11"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/.* Darwin/(9)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "1"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/.* Darwin/(10)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "4"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/.* Darwin/(11)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "5"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/.* Darwin/(13)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "6"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/6.* Darwin/(14)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "7"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/7.* Darwin/(14)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "8"
  os_v2_replacement: "0"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/7.* Darwin/(15)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "9"
  os_v2_replacement: "0"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/16\.5\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "10"
  os_v2_replacement: "3"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/16\.6\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "10"
  os_v2_replacement: "3"
  os_v3_replacement: "2"
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/16\.7\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "10"
  os_v2_replacement: "3"
  os_v3_replacement: "3"
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/(16)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "10"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/17\.0\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: "0"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/17\.2\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: "1"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/17\.3\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: "2"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/17\.4\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: "2"
  os_v3_replacement: "6"
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/17\.5\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: "3"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/9.* Darwin/17\.6\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: "4"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/9.* Darwin/17\.7\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: "4"
  os_v3_replacement: "1"
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/8.* Darwin/(17)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "11"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/9.* Darwin/18\.0\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "12"
  os_v2_replacement: "0"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/9.* Darwin/(18)\.\d+
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: "12"
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: CFNetwork/.* Darwin/
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: '\b(iOS[ /]|iOS; |iPhone(?:/| v|[ _]OS[/,]|; | OS : |\d,\d/|\d,\d; )|iPad/)(\d{1,2})[_\.](\d{1,2})(?:[_\.](\d+)|)'
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((iOS);
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (watchOS)/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: WatchOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Outlook-(iOS)/\d+\.\d+\.prod\.iphone
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (iPod|iPhone|iPad)
  regex_flag: ""
  os_replacement: iOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (tvOS)[/ ](\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: tvOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CrOS) [a-z0-9_]+ (\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: Chrome OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ([Dd]ebian)
  regex_flag: ""
  os_replacement: Debian
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Linux Mint)(?:/(\d+)|)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: '(Mandriva)(?: Linux|)/(?:[\d.-]+m[a-z]{2}(\d+).(\d)|)'
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Symbian[Oo][Ss])[/ ](\d+)\.(\d+)
  regex_flag: ""
  os_replacement: Symbian OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Symbian/3).+NokiaBrowser/7\.3
  regex_flag: ""
  os_replacement: Symbian^3 Anna
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Symbian/3).+NokiaBrowser/7\.4
  regex_flag: ""
  os_replacement: Symbian^3 Belle
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Symbian/3)
  regex_flag: ""
  os_replacement: Symbian^3
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \b(Series 60|SymbOS|S60Version|S60V\d|S60\b)
  regex_flag: ""
  os_replacement: Symbian OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (MeeGo)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Symbian [Oo][Ss]
  regex_flag: ""
  os_replacement: Symbian OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Series40;
  regex_flag: ""
  os_replacement: Nokia Series 40
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: Series30Plus;
  regex_flag: ""
  os_replacement: Nokia Series 30 Plus
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (BB10);.+Version/(\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  os_replacement: BlackBerry OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Black[Bb]erry)[0-9a-z]+/(\d+)\.(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: BlackBerry OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Black[Bb]erry).+Version/(\d+)\.(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: BlackBerry OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (RIM Tablet OS) (\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  os_replacement: BlackBerry Tablet OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Play[Bb]ook)
  regex_flag: ""
  os_replacement: BlackBerry Tablet OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Black[Bb]erry)
  regex_flag: ""
  os_replacement: BlackBerry OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Gecko/18.0 Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: "1"
  os_v2_replacement: "0"
  os_v3_replacement: "1"
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Gecko/18.1 Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: "1"
  os_v2_replacement: "1"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Gecko/26.0 Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: "1"
  os_v2_replacement: "2"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Gecko/28.0 Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: "1"
  os_v2_replacement: "3"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Gecko/30.0 Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: "1"
  os_v2_replacement: "4"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Gecko/32.0 Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: "2"
  os_v2_replacement: "0"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Gecko/34.0 Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: "2"
  os_v2_replacement: "1"
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((?:Mobile|Tablet);.+Firefox/\d+\.\d+
  regex_flag: ""
  os_replacement: Firefox OS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (BREW)[ /](\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (BREW);
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Brew MP|BMP)[ /](\d+)\.(\d+)\.(\d+)
  regex_flag: ""
  os_replacement: Brew MP
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: BMP;
  regex_flag: ""
  os_replacement: Brew MP
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: '(GoogleTV)(?: (\d+)\.(\d+)(?:\.(\d+)|)|/[\da-z]+)'
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (WebTV)/(\d+).(\d+)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (CrKey)(?:[/](\d+)\.(\d+)(?:\.(\d+)|)|)
  regex_flag: ""
  os_replacement: Chromecast
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (hpw|web)OS/(\d+)\.(\d+)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: webOS
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (VRE);
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Fedora|Red Hat|PCLinuxOS|Puppy|Ubuntu|Kindle|Bada|Sailfish|Lubuntu|BackTrack|Slackware|(?:Free|Open|Net|\b)BSD)[/
    ](\d+)\.(\d+)(?:\.(\d+)|)(?:\.(\d+)|)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Linux)[ /](\d+)\.(\d+)(?:\.(\d+)|).*gentoo
  regex_flag: ""
  os_replacement: Gentoo
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((Bada);
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Windows|Android|WeTab|Maemo|Web0S)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Ubuntu|Kubuntu|Arch Linux|CentOS|Slackware|Gentoo|openSUSE|SUSE|Red Hat|Fedora|PCLinuxOS|Mageia|(?:Free|Open|Net|\b)BSD)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: (Linux)(?:[ /](\d+)\.(\d+)(?:\.(\d+)|)|)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: SunOS
  regex_flag: ""
  os_replacement: Solaris
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \(linux-gnu\)
  regex_flag: ""
  os_replacement: Linux
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \(x86_64-redhat-linux-gnu\)
  regex_flag: ""
  os_replacement: Red Hat
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: \((freebsd)(\d+)\.(\d+)\)
  regex_flag: ""
  os_replacement: FreeBSD
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
- reg: {}
  regex: ^(Roku)/DVP-(\d+)\.(\d+)
  regex_flag: ""
  os_replacement: $1
  os_v1_replacement: $2
  os_v2_replacement: $3
  os_v3_replacement: $4
  os_v4_replacement: $5
device_parsers:
- reg: {}
  regex: (?:(?:iPhone|Windows CE|Windows Phone|Android).*(?:(?:Bot|Yeti)-Mobile|YRSpider|BingPreview|bots?/\d|(?:bot|spider)\.html)|AdsBot-Google-Mobile.*iPhone)
  regex_flag: i
  device_replacement: Spider
  brand_replacement: Spider
  model_replacement: Smartphone
- reg: {}
  regex: (?:DoCoMo|\bMOT\b|\bLG\b|Nokia|Samsung|SonyEricsson).*(?:(?:Bot|Yeti)-Mobile|bots?/\d|(?:bot|crawler)\.html|(?:jump|google|Wukong)bot|ichiro/mobile|/spider|YahooSeeker)
  regex_flag: i
  device_replacement: Spider
  brand_replacement: Spider
  model_replacement: Feature Phone
- reg: {}
  regex: ' PTST/\d+(?:\.)?\d+$'
  regex_flag: ""
  device_replacement: Spider
  brand_replacement: Spider
  model_replacement: $1
- reg: {}
  regex: X11; Datanyze; Linux
  regex_flag: ""
  device_replacement: Spider
  brand_replacement: Spider
  model_replacement: $1
- reg: {}
  regex: \bSmartWatch *\( *([^;]+) *; *([^;]+) *;
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: 'Android Application[^\-]+ - (Sony) ?(Ericsson|) (.+) \w+ - '
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1$2
  model_replacement: $3
- reg: {}
  regex: 'Android Application[^\-]+ - (?:HTC|HUAWEI|LGE|LENOVO|MEDION|TCT) (HTC|HUAWEI|LG|LENOVO|MEDION|ALCATEL)[
    _\-](.+) \w+ - '
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: 'Android Application[^\-]+ - ([^ ]+) (.+) \w+ - '
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *([BLRQ]C\d{4}[A-Z]+) +Build/
  regex_flag: ""
  device_replacement: 3Q $1
  brand_replacement: 3Q
  model_replacement: $1
- reg: {}
  regex: ; *(?:3Q_)([^;/]+) +Build
  regex_flag: ""
  device_replacement: 3Q $1
  brand_replacement: 3Q
  model_replacement: $1
- reg: {}
  regex: 'Android [34].*; *(A100|A101|A110|A200|A210|A211|A500|A501|A510|A511|A700(?:
    Lite| 3G|)|A701|B1-A71|A1-\d{3}|B1-\d{3}|V360|V370|W500|W500P|W501|W501P|W510|W511|W700|Slider
    SL101|DA22[^;/]+) Build'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Acer
  model_replacement: $1
- reg: {}
  regex: ; *Acer Iconia Tab ([^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Acer
  model_replacement: $1
- reg: {}
  regex: ; *(Z1[1235]0|E320[^/]*|S500|S510|Liquid[^;/]*|Iconia A\d+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Acer
  model_replacement: $1
- reg: {}
  regex: ; *(Acer |ACER )([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Acer
  model_replacement: $2
- reg: {}
  regex: ; *(Advent |)(Vega(?:Bean|Comb|)).* Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Advent
  model_replacement: $2
- reg: {}
  regex: ; *(Ainol |)((?:NOVO|[Nn]ovo)[^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Ainol
  model_replacement: $2
- reg: {}
  regex: ; *AIRIS[ _\-]?([^/;\)]+) *(?:;|\)|Build)
  regex_flag: i
  device_replacement: $1
  brand_replacement: Airis
  model_replacement: $1
- reg: {}
  regex: ; *(OnePAD[^;/]+) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: Airis
  model_replacement: $1
- reg: {}
  regex: ; *Airpad[ \-]([^;/]+) Build
  regex_flag: ""
  device_replacement: Airpad $1
  brand_replacement: Airpad
  model_replacement: $1
- reg: {}
  regex: ; *(one ?touch) (EVO7|T10|T20) Build
  regex_flag: ""
  device_replacement: Alcatel One Touch $2
  brand_replacement: Alcatel
  model_replacement: One Touch $2
- reg: {}
  regex: ; *(?:alcatel[ _]|)(?:(?:one[ _]?touch[ _])|ot[ \-])([^;/]+);? Build
  regex_flag: i
  device_replacement: Alcatel One Touch $1
  brand_replacement: Alcatel
  model_replacement: One Touch $1
- reg: {}
  regex: ; *(TCL)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(Vodafone Smart II|Optimus_Madrid) Build
  regex_flag: ""
  device_replacement: Alcatel $1
  brand_replacement: Alcatel
  model_replacement: $1
- reg: {}
  regex: ; *BASE_Lutea_3 Build
  regex_flag: ""
  device_replacement: Alcatel One Touch 998
  brand_replacement: Alcatel
  model_replacement: One Touch 998
- reg: {}
  regex: ; *BASE_Varia Build
  regex_flag: ""
  device_replacement: Alcatel One Touch 918D
  brand_replacement: Alcatel
  model_replacement: One Touch 918D
- reg: {}
  regex: ; *((?:FINE|Fine)\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Allfine
  model_replacement: $1
- reg: {}
  regex: ; *(ALLVIEW[ _]?|Allview[ _]?)((?:Speed|SPEED).*) Build/
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Allview
  model_replacement: $2
- reg: {}
  regex: ; *(ALLVIEW[ _]?|Allview[ _]?|)(AX1_Shine|AX2_Frenzy) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Allview
  model_replacement: $2
- reg: {}
  regex: ; *(ALLVIEW[ _]?|Allview[ _]?)([^;/]*) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Allview
  model_replacement: $2
- reg: {}
  regex: ; *(A13-MID) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Allwinner
  model_replacement: $1
- reg: {}
  regex: ; *(Allwinner)[ _\-]?([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Allwinner
  model_replacement: $1
- reg: {}
  regex: ; *(A651|A701B?|A702|A703|A705|A706|A707|A711|A712|A713|A717|A722|A785|A801|A802|A803|A901|A902|A1002|A1003|A1006|A1007|A9701|A9703|Q710|Q80)
    Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Amaway
  model_replacement: $1
- reg: {}
  regex: ; *(?:AMOI|Amoi)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: Amoi $1
  brand_replacement: Amoi
  model_replacement: $1
- reg: {}
  regex: ^(?:AMOI|Amoi)[ _]([^;/]+) Linux
  regex_flag: ""
  device_replacement: Amoi $1
  brand_replacement: Amoi
  model_replacement: $1
- reg: {}
  regex: ; *(MW(?:0[789]|10)[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Aoc
  model_replacement: $1
- reg: {}
  regex: ; *(G7|M1013|M1015G|M11[CG]?|M-?12[B]?|M15|M19[G]?|M30[ACQ]?|M31[GQ]|M32|M33[GQ]|M36|M37|M38|M701T|M710|M712B|M713|M715G|M716G|M71(?:G|GS|T|)|M72[T]?|M73[T]?|M75[GT]?|M77G|M79T|M7L|M7LN|M81|M810|M81T|M82|M92|M92KS|M92S|M717G|M721|M722G|M723|M725G|M739|M785|M791|M92SK|M93D)
    Build
  regex_flag: ""
  device_replacement: Aoson $1
  brand_replacement: Aoson
  model_replacement: $1
- reg: {}
  regex: ; *Aoson ([^;/]+) Build
  regex_flag: i
  device_replacement: Aoson $1
  brand_replacement: Aoson
  model_replacement: $1
- reg: {}
  regex: ; *[Aa]panda[ _\-]([^;/]+) Build
  regex_flag: ""
  device_replacement: Apanda $1
  brand_replacement: Apanda
  model_replacement: $1
- reg: {}
  regex: '; *(?:ARCHOS|Archos) ?(GAMEPAD.*?)(?: Build|[;/\(\)\-])'
  regex_flag: ""
  device_replacement: Archos $1
  brand_replacement: Archos
  model_replacement: $1
- reg: {}
  regex: ARCHOS; GOGI; ([^;]+);
  regex_flag: ""
  device_replacement: Archos $1
  brand_replacement: Archos
  model_replacement: $1
- reg: {}
  regex: '(?:ARCHOS|Archos)[ _]?(.*?)(?: Build|[;/\(\)\-]|$)'
  regex_flag: ""
  device_replacement: Archos $1
  brand_replacement: Archos
  model_replacement: $1
- reg: {}
  regex: ; *(AN(?:7|8|9|10|13)[A-Z0-9]{1,4}) Build
  regex_flag: ""
  device_replacement: Archos $1
  brand_replacement: Archos
  model_replacement: $1
- reg: {}
  regex: ; *(A28|A32|A43|A70(?:BHT|CHT|HB|S|X)|A101(?:B|C|IT)|A7EB|A7EB-WK|101G9|80G9)
    Build
  regex_flag: ""
  device_replacement: Archos $1
  brand_replacement: Archos
  model_replacement: $1
- reg: {}
  regex: ; *(PAD-FMD[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Arival
  model_replacement: $1
- reg: {}
  regex: ; *(BioniQ) ?([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Arival
  model_replacement: $1 $2
- reg: {}
  regex: ; *(AN\d[^;/]+|ARCHM\d+) Build
  regex_flag: ""
  device_replacement: Arnova $1
  brand_replacement: Arnova
  model_replacement: $1
- reg: {}
  regex: ; *(?:ARNOVA|Arnova) ?([^;/]+) Build
  regex_flag: ""
  device_replacement: Arnova $1
  brand_replacement: Arnova
  model_replacement: $1
- reg: {}
  regex: ; *(?:ASSISTANT |)(AP)-?([1789]\d{2}[A-Z]{0,2}|80104) Build
  regex_flag: ""
  device_replacement: Assistant $1-$2
  brand_replacement: Assistant
  model_replacement: $1-$2
- reg: {}
  regex: '; *(ME17\d[^;/]*|ME3\d{2}[^;/]+|K00[A-Z]|Nexus 10|Nexus 7(?: 2013|)|PadFone[^;/]*|Transformer[^;/]*|TF\d{3}[^;/]*|eeepc)
    Build'
  regex_flag: ""
  device_replacement: Asus $1
  brand_replacement: Asus
  model_replacement: $1
- reg: {}
  regex: ; *ASUS[ _]*([^;/]+) Build
  regex_flag: ""
  device_replacement: Asus $1
  brand_replacement: Asus
  model_replacement: $1
- reg: {}
  regex: ; *Garmin-Asus ([^;/]+) Build
  regex_flag: ""
  device_replacement: Garmin-Asus $1
  brand_replacement: Garmin-Asus
  model_replacement: $1
- reg: {}
  regex: ; *(Garminfone) Build
  regex_flag: ""
  device_replacement: Garmin $1
  brand_replacement: Garmin-Asus
  model_replacement: $1
- reg: {}
  regex: ; (@TAB-[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Attab
  model_replacement: $1
- reg: {}
  regex: ; *(T-(?:07|[^0]\d)[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Audiosonic
  model_replacement: $1
- reg: {}
  regex: ; *(?:Axioo[ _\-]([^;/]+)|(picopad)[ _\-]([^;/]+)) Build
  regex_flag: i
  device_replacement: Axioo $1$2 $3
  brand_replacement: Axioo
  model_replacement: $1$2 $3
- reg: {}
  regex: ; *(V(?:100|700|800)[^;/]*) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Azend
  model_replacement: $1
- reg: {}
  regex: ; *(IBAK\-[^;/]*) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: Bak
  model_replacement: $1
- reg: {}
  regex: ; *(HY5001|HY6501|X12|X21|I5) Build
  regex_flag: ""
  device_replacement: Bedove $1
  brand_replacement: Bedove
  model_replacement: $1
- reg: {}
  regex: ; *(JC-[^;/]*) Build
  regex_flag: ""
  device_replacement: Benss $1
  brand_replacement: Benss
  model_replacement: $1
- reg: {}
  regex: ; *(BB) ([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Blackberry
  model_replacement: $2
- reg: {}
  regex: ; *(BlackBird)[ _](I8.*) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(BlackBird)[ _](.*) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *([0-9]+BP[EM][^;/]*|Endeavour[^;/]+) Build
  regex_flag: ""
  device_replacement: Blaupunkt $1
  brand_replacement: Blaupunkt
  model_replacement: $1
- reg: {}
  regex: ; *((?:BLU|Blu)[ _\-])([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Blu
  model_replacement: $2
- reg: {}
  regex: ; *(?:BMOBILE )?(Blu|BLU|DASH [^;/]+|VIVO 4\.3|TANK 4\.5) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Blu
  model_replacement: $1
- reg: {}
  regex: ; *(TOUCH\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Blusens
  model_replacement: $1
- reg: {}
  regex: ; *(AX5\d+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Bmobile
  model_replacement: $1
- reg: {}
  regex: ; *([Bb]q) ([^;/]+);? Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: bq
  model_replacement: $2
- reg: {}
  regex: ; *(Maxwell [^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: bq
  model_replacement: $1
- reg: {}
  regex: ; *((?:B-Tab|B-TAB) ?\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Braun
  model_replacement: $1
- reg: {}
  regex: ; *(Broncho) ([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *CAPTIVA ([^;/]+) Build
  regex_flag: ""
  device_replacement: Captiva $1
  brand_replacement: Captiva
  model_replacement: $1
- reg: {}
  regex: ; *(C771|CAL21|IS11CA) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Casio
  model_replacement: $1
- reg: {}
  regex: ; *(?:Cat|CAT) ([^;/]+) Build
  regex_flag: ""
  device_replacement: Cat $1
  brand_replacement: Cat
  model_replacement: $1
- reg: {}
  regex: ; *(?:Cat)(Nova.*) Build
  regex_flag: ""
  device_replacement: Cat $1
  brand_replacement: Cat
  model_replacement: $1
- reg: {}
  regex: ; *(INM8002KP|ADM8000KP_[AB]) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Cat
  model_replacement: Tablet PHOENIX 8.1J0
- reg: {}
  regex: ; *(?:[Cc]elkon[ _\*]|CELKON[ _\*])([^;/\)]+) ?(?:Build|;|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Celkon
  model_replacement: $1
- reg: {}
  regex: Build/(?:[Cc]elkon)+_?([^;/_\)]+)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Celkon
  model_replacement: $1
- reg: {}
  regex: ; *(CT)-?(\d+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Celkon
  model_replacement: $1$2
- reg: {}
  regex: ; *(A19|A19Q|A105|A107[^;/\)]*) ?(?:Build|;|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Celkon
  model_replacement: $1
- reg: {}
  regex: ; *(TPC[0-9]{4,5}) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: ChangJia
  model_replacement: $1
- reg: {}
  regex: ; *(Cloudfone)[ _](Excite)([^ ][^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2 $3
  brand_replacement: Cloudfone
  model_replacement: $1 $2 $3
- reg: {}
  regex: ; *(Excite|ICE)[ _](\d+[^;/]+) Build
  regex_flag: ""
  device_replacement: Cloudfone $1 $2
  brand_replacement: Cloudfone
  model_replacement: Cloudfone $1 $2
- reg: {}
  regex: ; *(Cloudfone|CloudPad)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Cloudfone
  model_replacement: $1 $2
- reg: {}
  regex: ; *((?:Aquila|Clanga|Rapax)[^;/]+) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: Cmx
  model_replacement: $1
- reg: {}
  regex: ; *(?:CFW-|Kyros )?(MID[0-9]{4}(?:[ABC]|SR|TV)?)(\(3G\)-4G| GB 8K| 3G| 8K|
    GB)? *(?:Build|[;\)])
  regex_flag: ""
  device_replacement: CobyKyros $1$2
  brand_replacement: CobyKyros
  model_replacement: $1$2
- reg: {}
  regex: ; *([^;/]*)Coolpad[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Coolpad
  model_replacement: $1$2
- reg: {}
  regex: ; *(CUBE[ _])?([KU][0-9]+ ?GT.*|A5300) Build
  regex_flag: i
  device_replacement: $1$2
  brand_replacement: Cube
  model_replacement: $2
- reg: {}
  regex: ; *CUBOT ([^;/]+) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: Cubot
  model_replacement: $1
- reg: {}
  regex: ; *(BOBBY) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: Cubot
  model_replacement: $1
- reg: {}
  regex: ; *(Dslide [^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Danew
  model_replacement: $1
- reg: {}
  regex: ; *(XCD)[ _]?(28|35) Build
  regex_flag: ""
  device_replacement: Dell $1$2
  brand_replacement: Dell
  model_replacement: $1$2
- reg: {}
  regex: ; *(001DL) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: Streak
- reg: {}
  regex: ; *(?:Dell|DELL) (Streak) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: Streak
- reg: {}
  regex: ; *(101DL|GS01|Streak Pro[^;/]*) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: Streak Pro
- reg: {}
  regex: ; *([Ss]treak ?7) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: Streak 7
- reg: {}
  regex: ; *(Mini-3iX) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: $1
- reg: {}
  regex: ; *(?:Dell|DELL)[ _](Aero|Venue|Thunder|Mini.*|Streak[ _]Pro) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: $1
- reg: {}
  regex: ; *Dell[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: $1
- reg: {}
  regex: ; *Dell ([^;/]+) Build
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: $1
- reg: {}
  regex: ; *(TA[CD]-\d+[^;/]*) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Denver
  model_replacement: $1
- reg: {}
  regex: ; *(iP[789]\d{2}(?:-3G)?|IP10\d{2}(?:-8GB)?) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Dex
  model_replacement: $1
- reg: {}
  regex: ; *(AirTab)[ _\-]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: DNS
  model_replacement: $1 $2
- reg: {}
  regex: ; *(F\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Fujitsu
  model_replacement: $1
- reg: {}
  regex: ; *(HT-03A) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: Magic
- reg: {}
  regex: ; *(HT\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: ; *(L\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: ; *(N\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Nec
  model_replacement: $1
- reg: {}
  regex: ; *(P\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Panasonic
  model_replacement: $1
- reg: {}
  regex: ; *(SC\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ; *(SH\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sharp
  model_replacement: $1
- reg: {}
  regex: ; *(SO\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: SonyEricsson
  model_replacement: $1
- reg: {}
  regex: ; *(T\-0[12][^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Toshiba
  model_replacement: $1
- reg: {}
  regex: ; *(DOOV)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: DOOV
  model_replacement: $2
- reg: {}
  regex: ; *(Enot|ENOT)[ -]?([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Enot
  model_replacement: $2
- reg: {}
  regex: ; *[^;/]+ Build/(?:CROSS|Cross)+[ _\-]([^\)]+)
  regex_flag: ""
  device_replacement: CROSS $1
  brand_replacement: Evercoss
  model_replacement: Cross $1
- reg: {}
  regex: ; *(CROSS|Cross)[ _\-]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Evercoss
  model_replacement: Cross $2
- reg: {}
  regex: ; *Explay[_ ](.+?)(?:[\)]| Build)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Explay
  model_replacement: $1
- reg: {}
  regex: ; *(IQ.*) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Fly
  model_replacement: $1
- reg: {}
  regex: ; *(Fly|FLY)[ _](IQ[^;]+|F[34]\d+[^;]*);? Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Fly
  model_replacement: $2
- reg: {}
  regex: ; *(M532|Q572|FJL21) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Fujitsu
  model_replacement: $1
- reg: {}
  regex: ; *(G1) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Galapad
  model_replacement: $1
- reg: {}
  regex: ; *(Geeksphone) ([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(G[^F]?FIVE) ([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Gfive
  model_replacement: $2
- reg: {}
  regex: ; *(Gionee)[ _\-]([^;/]+)(?:/[^;/]+|) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: Gionee
  model_replacement: $2
- reg: {}
  regex: ; *(GN\d+[A-Z]?|INFINITY_PASSION|Ctrl_V1) Build
  regex_flag: ""
  device_replacement: Gionee $1
  brand_replacement: Gionee
  model_replacement: $1
- reg: {}
  regex: ; *(E3) Build/JOP40D
  regex_flag: ""
  device_replacement: Gionee $1
  brand_replacement: Gionee
  model_replacement: $1
- reg: {}
  regex: \sGIONEE[-\s_](\w*)
  regex_flag: i
  device_replacement: Gionee $1
  brand_replacement: Gionee
  model_replacement: $1
- reg: {}
  regex: ; *((?:FONE|QUANTUM|INSIGNIA) \d+[^;/]*|PLAYTAB) Build
  regex_flag: ""
  device_replacement: GoClever $1
  brand_replacement: GoClever
  model_replacement: $1
- reg: {}
  regex: ; *GOCLEVER ([^;/]+) Build
  regex_flag: ""
  device_replacement: GoClever $1
  brand_replacement: GoClever
  model_replacement: $1
- reg: {}
  regex: ; *(Glass \d+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Google
  model_replacement: $1
- reg: {}
  regex: ; *(Pixel.*) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Google
  model_replacement: $1
- reg: {}
  regex: ; *(GSmart)[ -]([^/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Gigabyte
  model_replacement: $1 $2
- reg: {}
  regex: ; *(imx5[13]_[^/]+) Build
  regex_flag: ""
  device_replacement: Freescale $1
  brand_replacement: Freescale
  model_replacement: $1
- reg: {}
  regex: ; *Haier[ _\-]([^/]+) Build
  regex_flag: ""
  device_replacement: Haier $1
  brand_replacement: Haier
  model_replacement: $1
- reg: {}
  regex: ; *(PAD1016) Build
  regex_flag: ""
  device_replacement: Haipad $1
  brand_replacement: Haipad
  model_replacement: $1
- reg: {}
  regex: ; *(M701|M7|M8|M9) Build
  regex_flag: ""
  device_replacement: Haipad $1
  brand_replacement: Haipad
  model_replacement: $1
- reg: {}
  regex: '; *(SN\d+T[^;\)/]*)(?: Build|[;\)])'
  regex_flag: ""
  device_replacement: Hannspree $1
  brand_replacement: Hannspree
  model_replacement: $1
- reg: {}
  regex: Build/HCL ME Tablet ([^;\)]+)[\);]
  regex_flag: ""
  device_replacement: HCLme $1
  brand_replacement: HCLme
  model_replacement: $1
- reg: {}
  regex: ; *([^;\/]+) Build/HCL
  regex_flag: ""
  device_replacement: HCLme $1
  brand_replacement: HCLme
  model_replacement: $1
- reg: {}
  regex: ; *(MID-?\d{4}C[EM]) Build
  regex_flag: ""
  device_replacement: Hena $1
  brand_replacement: Hena
  model_replacement: $1
- reg: {}
  regex: ; *(EG\d{2,}|HS-[^;/]+|MIRA[^;/]+) Build
  regex_flag: ""
  device_replacement: Hisense $1
  brand_replacement: Hisense
  model_replacement: $1
- reg: {}
  regex: ; *(andromax[^;/]+) Build
  regex_flag: i
  device_replacement: Hisense $1
  brand_replacement: Hisense
  model_replacement: $1
- reg: {}
  regex: ; *(?:AMAZE[ _](S\d+)|(S\d+)[ _]AMAZE) Build
  regex_flag: ""
  device_replacement: AMAZE $1$2
  brand_replacement: hitech
  model_replacement: AMAZE $1$2
- reg: {}
  regex: ; *(PlayBook) Build
  regex_flag: ""
  device_replacement: HP $1
  brand_replacement: HP
  model_replacement: $1
- reg: {}
  regex: ; *HP ([^/]+) Build
  regex_flag: ""
  device_replacement: HP $1
  brand_replacement: HP
  model_replacement: $1
- reg: {}
  regex: ; *([^/]+_tenderloin) Build
  regex_flag: ""
  device_replacement: HP TouchPad
  brand_replacement: HP
  model_replacement: TouchPad
- reg: {}
  regex: ; *(HUAWEI |Huawei-|)([UY][^;/]+) Build/(?:Huawei|HUAWEI)([UY][^\);]+)\)
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Huawei
  model_replacement: $2
- reg: {}
  regex: ; *([^;/]+) Build[/ ]Huawei(MT1-U06|[A-Z]+\d+[^\);]+)[^\);]*\)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Huawei
  model_replacement: $2
- reg: {}
  regex: ; *(S7|M860) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: ; *((?:HUAWEI|Huawei)[ \-]?)(MediaPad) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Huawei
  model_replacement: $2
- reg: {}
  regex: ; *((?:HUAWEI[ _]?|Huawei[ _]|)Ascend[ _])([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Huawei
  model_replacement: $2
- reg: {}
  regex: ; *((?:HUAWEI|Huawei)[ _\-]?)((?:G700-|MT-)[^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Huawei
  model_replacement: $2
- reg: {}
  regex: ; *((?:HUAWEI|Huawei)[ _\-]?)([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Huawei
  model_replacement: $2
- reg: {}
  regex: ; *(MediaPad[^;]+|SpringBoard) Build/Huawei
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: ; *([^;]+) Build/(?:Huawei|HUAWEI)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: ; *([Uu])([89]\d{3}) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Huawei
  model_replacement: U$2
- reg: {}
  regex: ; *(?:Ideos |IDEOS )(S7) Build
  regex_flag: ""
  device_replacement: Huawei Ideos$1
  brand_replacement: Huawei
  model_replacement: Ideos$1
- reg: {}
  regex: ; *(?:Ideos |IDEOS )([^;/]+\s*|\s*)Build
  regex_flag: ""
  device_replacement: Huawei Ideos$1
  brand_replacement: Huawei
  model_replacement: Ideos$1
- reg: {}
  regex: ; *(Orange Daytona|Pulse|Pulse Mini|Vodafone 858|C8500|C8600|C8650|C8660|Nexus
    6P|ATH-.+?) Build[/ ]
  regex_flag: ""
  device_replacement: Huawei $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: ; *((?:[A-Z]{3})\-L[A-Za0-9]{2})[\)]
  regex_flag: ""
  device_replacement: Huawei $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: ; *HTC[ _]([^;]+); Windows Phone
  regex_flag: ""
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: '; *(?:HTC[ _/])+([^ _/]+)(?:[/\\]1\.0 | V|/| +)\d+\.\d[\d\.]*(?: *Build|\))'
  regex_flag: ""
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: '; *(?:HTC[ _/])+([^ _/]+)(?:[ _/]([^ _/]+)|)(?:[/\\]1\.0 | V|/| +)\d+\.\d[\d\.]*(?:
    *Build|\))'
  regex_flag: ""
  device_replacement: HTC $1 $2
  brand_replacement: HTC
  model_replacement: $1 $2
- reg: {}
  regex: '; *(?:HTC[ _/])+([^ _/]+)(?:[ _/]([^ _/]+)(?:[ _/]([^ _/]+)|)|)(?:[/\\]1\.0
    | V|/| +)\d+\.\d[\d\.]*(?: *Build|\))'
  regex_flag: ""
  device_replacement: HTC $1 $2 $3
  brand_replacement: HTC
  model_replacement: $1 $2 $3
- reg: {}
  regex: '; *(?:HTC[ _/])+([^ _/]+)(?:[ _/]([^ _/]+)(?:[ _/]([^ _/]+)(?:[ _/]([^ _/]+)|)|)|)(?:[/\\]1\.0
    | V|/| +)\d+\.\d[\d\.]*(?: *Build|\))'
  regex_flag: ""
  device_replacement: HTC $1 $2 $3 $4
  brand_replacement: HTC
  model_replacement: $1 $2 $3 $4
- reg: {}
  regex: '; *(?:(?:HTC|htc)(?:_blocked|)[ _/])+([^ _/;]+)(?: *Build|[;\)]| - )'
  regex_flag: ""
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: '; *(?:(?:HTC|htc)(?:_blocked|)[ _/])+([^ _/]+)(?:[ _/]([^ _/;\)]+)|)(?:
    *Build|[;\)]| - )'
  regex_flag: ""
  device_replacement: HTC $1 $2
  brand_replacement: HTC
  model_replacement: $1 $2
- reg: {}
  regex: '; *(?:(?:HTC|htc)(?:_blocked|)[ _/])+([^ _/]+)(?:[ _/]([^ _/]+)(?:[ _/]([^
    _/;\)]+)|)|)(?: *Build|[;\)]| - )'
  regex_flag: ""
  device_replacement: HTC $1 $2 $3
  brand_replacement: HTC
  model_replacement: $1 $2 $3
- reg: {}
  regex: '; *(?:(?:HTC|htc)(?:_blocked|)[ _/])+([^ _/]+)(?:[ _/]([^ _/]+)(?:[ _/]([^
    _/]+)(?:[ _/]([^ /;]+)|)|)|)(?: *Build|[;\)]| - )'
  regex_flag: ""
  device_replacement: HTC $1 $2 $3 $4
  brand_replacement: HTC
  model_replacement: $1 $2 $3 $4
- reg: {}
  regex: HTC Streaming Player [^\/]*/[^\/]*/ htc_([^/]+) /
  regex_flag: ""
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: '(?:[;,] *|^)(?:htccn_chs-|)HTC[ _-]?([^;]+?)(?: *Build|clay|Android|-?Mozilla|
    Opera| Profile| UNTRUSTED|[;/\(\)]|$)'
  regex_flag: i
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: '; *(A6277|ADR6200|ADR6300|ADR6350|ADR6400[A-Z]*|ADR6425[A-Z]*|APX515CKT|ARIA|Desire[^_
    ]*|Dream|EndeavorU|Eris|Evo|Flyer|HD2|Hero|HERO200|Hero CDMA|HTL21|Incredible|Inspire[A-Z0-9]*|Legend|Liberty|Nexus
    ?(?:One|HD2)|One|One S C2|One[ _]?(?:S|V|X\+?)\w*|PC36100|PG06100|PG86100|S31HT|Sensation|Wildfire)(?:
    Build|[/;\(\)])'
  regex_flag: i
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: ; *(ADR6200|ADR6400L|ADR6425LVW|Amaze|DesireS?|EndeavorU|Eris|EVO|Evo\d[A-Z]+|HD2|IncredibleS?|Inspire[A-Z0-9]*|Inspire[A-Z0-9]*|Sensation[A-Z0-9]*|Wildfire)[
    _-](.+?)(?:[/;\)]|Build|MIUI|1\.0)
  regex_flag: i
  device_replacement: HTC $1 $2
  brand_replacement: HTC
  model_replacement: $1 $2
- reg: {}
  regex: ; *HYUNDAI (T\d[^/]*) Build
  regex_flag: ""
  device_replacement: Hyundai $1
  brand_replacement: Hyundai
  model_replacement: $1
- reg: {}
  regex: ; *HYUNDAI ([^;/]+) Build
  regex_flag: ""
  device_replacement: Hyundai $1
  brand_replacement: Hyundai
  model_replacement: $1
- reg: {}
  regex: ; *(X700|Hold X|MB-6900) Build
  regex_flag: ""
  device_replacement: Hyundai $1
  brand_replacement: Hyundai
  model_replacement: $1
- reg: {}
  regex: ; *(?:iBall[ _\-]|)(Andi)[ _]?(\d[^;/]*) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: iBall
  model_replacement: $1 $2
- reg: {}
  regex: ; *(IBall)(?:[ _]([^;/]+)|) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: iBall
  model_replacement: $2
- reg: {}
  regex: '; *(NT-\d+[^ ;/]*|Net[Tt]AB [^;/]+|Mercury [A-Z]+|iconBIT)(?: S/N:[^;/]+|)
    Build'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: IconBIT
  model_replacement: $1
- reg: {}
  regex: ; *(IMO)[ _]([^;/]+) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: IMO
  model_replacement: $2
- reg: {}
  regex: ; *i-?mobile[ _]([^/]+) Build/
  regex_flag: i
  device_replacement: i-mobile $1
  brand_replacement: imobile
  model_replacement: $1
- reg: {}
  regex: ; *(i-(?:style|note)[^/]*) Build/
  regex_flag: i
  device_replacement: i-mobile $1
  brand_replacement: imobile
  model_replacement: $1
- reg: {}
  regex: ; *(ImPAD) ?(\d+(?:.)*) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Impression
  model_replacement: $1 $2
- reg: {}
  regex: ; *(Infinix)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Infinix
  model_replacement: $2
- reg: {}
  regex: ; *(Informer)[ \-]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Informer
  model_replacement: $2
- reg: {}
  regex: ; *(TAB) ?([78][12]4) Build
  regex_flag: ""
  device_replacement: Intenso $1
  brand_replacement: Intenso
  model_replacement: $1 $2
- reg: {}
  regex: ; *(?:Intex[ _]|)(AQUA|Aqua)([ _\.\-])([^;/]+) *(?:Build|;)
  regex_flag: ""
  device_replacement: $1$2$3
  brand_replacement: Intex
  model_replacement: $1 $3
- reg: {}
  regex: ; *(?:INTEX|Intex)(?:[_ ]([^\ _;/]+))(?:[_ ]([^\ _;/]+)|) *(?:Build|;)
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Intex
  model_replacement: $1 $2
- reg: {}
  regex: ; *([iI]Buddy)[ _]?(Connect)(?:_|\?_| |)([^;/]*) *(?:Build|;)
  regex_flag: ""
  device_replacement: $1 $2 $3
  brand_replacement: Intex
  model_replacement: iBuddy $2 $3
- reg: {}
  regex: ; *(I-Buddy)[ _]([^;/]+) *(?:Build|;)
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Intex
  model_replacement: iBuddy $2
- reg: {}
  regex: ; *(iOCEAN) ([^/]+) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: iOCEAN
  model_replacement: $2
- reg: {}
  regex: ; *(TP\d+(?:\.\d+|)\-\d[^;/]+) Build
  regex_flag: ""
  device_replacement: ionik $1
  brand_replacement: ionik
  model_replacement: $1
- reg: {}
  regex: ; *(M702pro) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Iru
  model_replacement: $1
- reg: {}
  regex: ; *(DE88Plus|MD70) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Ivio
  model_replacement: $1
- reg: {}
  regex: ; *IVIO[_\-]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Ivio
  model_replacement: $1
- reg: {}
  regex: ; *(TPC-\d+|JAY-TECH) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Jaytech
  model_replacement: $1
- reg: {}
  regex: ; *(JY-[^;/]+|G[234]S?) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Jiayu
  model_replacement: $1
- reg: {}
  regex: ; *(JXD)[ _\-]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: JXD
  model_replacement: $2
- reg: {}
  regex: ; *Karbonn[ _]?([^;/]+) *(?:Build|;)
  regex_flag: i
  device_replacement: $1
  brand_replacement: Karbonn
  model_replacement: $1
- reg: {}
  regex: ; *([^;]+) Build/Karbonn
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Karbonn
  model_replacement: $1
- reg: {}
  regex: ; *(A11|A39|A37|A34|ST8|ST10|ST7|Smart Tab3|Smart Tab2|Titanium S\d) +Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Karbonn
  model_replacement: $1
- reg: {}
  regex: ; *(IS01|IS03|IS05|IS\d{2}SH) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sharp
  model_replacement: $1
- reg: {}
  regex: ; *(IS04) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Regza
  model_replacement: $1
- reg: {}
  regex: ; *(IS06|IS\d{2}PT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Pantech
  model_replacement: $1
- reg: {}
  regex: ; *(IS11S) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: SonyEricsson
  model_replacement: Xperia Acro
- reg: {}
  regex: ; *(IS11CA) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Casio
  model_replacement: GzOne $1
- reg: {}
  regex: ; *(IS11LG) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: LG
  model_replacement: Optimus X
- reg: {}
  regex: ; *(IS11N) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Medias
  model_replacement: $1
- reg: {}
  regex: ; *(IS11PT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Pantech
  model_replacement: MIRACH
- reg: {}
  regex: ; *(IS12F) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Fujitsu
  model_replacement: Arrows ES
- reg: {}
  regex: ; *(IS12M) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Motorola
  model_replacement: XT909
- reg: {}
  regex: ; *(IS12S) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: SonyEricsson
  model_replacement: Xperia Acro HD
- reg: {}
  regex: ; *(ISW11F) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Fujitsu
  model_replacement: Arrowz Z
- reg: {}
  regex: ; *(ISW11HT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: EVO
- reg: {}
  regex: ; *(ISW11K) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Kyocera
  model_replacement: DIGNO
- reg: {}
  regex: ; *(ISW11M) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Motorola
  model_replacement: Photon
- reg: {}
  regex: ; *(ISW11SC) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Samsung
  model_replacement: GALAXY S II WiMAX
- reg: {}
  regex: ; *(ISW12HT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: EVO 3D
- reg: {}
  regex: ; *(ISW13HT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: J
- reg: {}
  regex: ; *(ISW?[0-9]{2}[A-Z]{0,2}) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: KDDI
  model_replacement: $1
- reg: {}
  regex: ; *(INFOBAR [^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: KDDI
  model_replacement: $1
- reg: {}
  regex: ; *(JOYPAD|Joypad)[ _]([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Kingcom
  model_replacement: $1 $2
- reg: {}
  regex: ; *(Vox|VOX|Arc|K080) Build/
  regex_flag: i
  device_replacement: $1
  brand_replacement: Kobo
  model_replacement: $1
- reg: {}
  regex: \b(Kobo Touch)\b
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Kobo
  model_replacement: $1
- reg: {}
  regex: ; *(K-Touch)[ _]([^;/]+) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: Ktouch
  model_replacement: $2
- reg: {}
  regex: ; *((?:EV|KM)-S\d+[A-Z]?) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: KTtech
  model_replacement: $1
- reg: {}
  regex: ; *(Zio|Hydro|Torque|Event|EVENT|Echo|Milano|Rise|URBANO PROGRESSO|WX04K|WX06K|WX10K|KYL21|101K|C5[12]\d{2})
    Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Kyocera
  model_replacement: $1
- reg: {}
  regex: ; *(?:LAVA[ _]|)IRIS[ _\-]?([^/;\)]+) *(?:;|\)|Build)
  regex_flag: i
  device_replacement: Iris $1
  brand_replacement: Lava
  model_replacement: Iris $1
- reg: {}
  regex: ; *LAVA[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Lava
  model_replacement: $1
- reg: {}
  regex: ; *(?:(Aspire A1)|(?:LEMON|Lemon)[ _]([^;/]+))_? Build
  regex_flag: ""
  device_replacement: Lemon $1$2
  brand_replacement: Lemon
  model_replacement: $1$2
- reg: {}
  regex: ; *(TAB-1012) Build/
  regex_flag: ""
  device_replacement: Lenco $1
  brand_replacement: Lenco
  model_replacement: $1
- reg: {}
  regex: ; Lenco ([^;/]+) Build/
  regex_flag: ""
  device_replacement: Lenco $1
  brand_replacement: Lenco
  model_replacement: $1
- reg: {}
  regex: ; *(A1_07|A2107A-H|S2005A-H|S1-37AH0) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Lenovo
  model_replacement: $1
- reg: {}
  regex: ; *(Idea[Tp]ab)[ _]([^;/]+);? Build
  regex_flag: ""
  device_replacement: Lenovo $1 $2
  brand_replacement: Lenovo
  model_replacement: $1 $2
- reg: {}
  regex: ; *(Idea(?:Tab|pad)) ?([^;/]+) Build
  regex_flag: ""
  device_replacement: Lenovo $1 $2
  brand_replacement: Lenovo
  model_replacement: $1 $2
- reg: {}
  regex: ; *(ThinkPad) ?(Tablet) Build/
  regex_flag: ""
  device_replacement: Lenovo $1 $2
  brand_replacement: Lenovo
  model_replacement: $1 $2
- reg: {}
  regex: ; *(?:LNV-|)(?:=?[Ll]enovo[ _\-]?|LENOVO[ _])(.+?)(?:Build|[;/\)])
  regex_flag: ""
  device_replacement: Lenovo $1
  brand_replacement: Lenovo
  model_replacement: $1
- reg: {}
  regex: '[;,] (?:Vodafone |)(SmartTab) ?(II) ?(\d+) Build/'
  regex_flag: ""
  device_replacement: Lenovo $1 $2 $3
  brand_replacement: Lenovo
  model_replacement: $1 $2 $3
- reg: {}
  regex: ; *(?:Ideapad |)K1 Build/
  regex_flag: ""
  device_replacement: Lenovo Ideapad K1
  brand_replacement: Lenovo
  model_replacement: Ideapad K1
- reg: {}
  regex: ; *(3GC101|3GW10[01]|A390) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Lenovo
  model_replacement: $1
- reg: {}
  regex: \b(?:Lenovo|LENOVO)+[ _\-]?([^,;:/ ]+)
  regex_flag: ""
  device_replacement: Lenovo $1
  brand_replacement: Lenovo
  model_replacement: $1
- reg: {}
  regex: ; *(MFC\d+)[A-Z]{2}([^;,/]*),? Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Lexibook
  model_replacement: $1$2
- reg: {}
  regex: ; *(E[34][0-9]{2}|LS[6-8][0-9]{2}|VS[6-9][0-9]+[^;/]+|Nexus 4|Nexus 5X?|GT540f?|Optimus
    (?:2X|G|4X HD)|OptimusX4HD) *(?:Build|;)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: '[;:] *(L-\d+[A-Z]|LGL\d+[A-Z]?)(?:/V\d+|) *(?:Build|[;\)])'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: ; *(LG-)([A-Z]{1,2}\d{2,}[^,;/\)\(]*?)(?:Build| V\d+|[,;/\)\(]|$)
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: LG
  model_replacement: $2
- reg: {}
  regex: ; *(LG[ \-]|LG)([^;/]+)[;/]? Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: LG
  model_replacement: $2
- reg: {}
  regex: ^(LG)-([^;/]+)/ Mozilla/.*; Android
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: LG
  model_replacement: $2
- reg: {}
  regex: (Web0S); Linux/(SmartTV)
  regex_flag: ""
  device_replacement: LG $1 $2
  brand_replacement: LG
  model_replacement: $1 $2
- reg: {}
  regex: ; *((?:SMB|smb)[^;/]+) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Malata
  model_replacement: $1
- reg: {}
  regex: ; *(?:Malata|MALATA) ([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Malata
  model_replacement: $1
- reg: {}
  regex: ; *(MS[45][0-9]{3}|MID0[568][NS]?|MID[1-9]|MID[78]0[1-9]|MID970[1-9]|MID100[1-9])
    Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Manta
  model_replacement: $1
- reg: {}
  regex: ; *(M1052|M806|M9000|M9100|M9701|MID100|MID120|MID125|MID130|MID135|MID140|MID701|MID710|MID713|MID727|MID728|MID731|MID732|MID733|MID735|MID736|MID737|MID760|MID800|MID810|MID820|MID830|MID833|MID835|MID860|MID900|MID930|MID933|MID960|MID980)
    Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Match
  model_replacement: $1
- reg: {}
  regex: ; *(GenxDroid7|MSD7.*|AX\d.*|Tab 701|Tab 722) Build/
  regex_flag: ""
  device_replacement: Maxx $1
  brand_replacement: Maxx
  model_replacement: $1
- reg: {}
  regex: ; *(M-PP[^;/]+|PhonePad ?\d{2,}[^;/]+) Build
  regex_flag: ""
  device_replacement: Mediacom $1
  brand_replacement: Mediacom
  model_replacement: $1
- reg: {}
  regex: ; *(M-MP[^;/]+|SmartPad ?\d{2,}[^;/]+) Build
  regex_flag: ""
  device_replacement: Mediacom $1
  brand_replacement: Mediacom
  model_replacement: $1
- reg: {}
  regex: ; *(?:MD_|)LIFETAB[ _]([^;/]+) Build
  regex_flag: i
  device_replacement: Medion Lifetab $1
  brand_replacement: Medion
  model_replacement: Lifetab $1
- reg: {}
  regex: ; *MEDION ([^;/]+) Build
  regex_flag: ""
  device_replacement: Medion $1
  brand_replacement: Medion
  model_replacement: $1
- reg: {}
  regex: ; *(M030|M031|M035|M040|M065|m9) Build
  regex_flag: ""
  device_replacement: Meizu $1
  brand_replacement: Meizu
  model_replacement: $1
- reg: {}
  regex: ; *(?:meizu_|MEIZU )(.+?) *(?:Build|[;\)])
  regex_flag: ""
  device_replacement: Meizu $1
  brand_replacement: Meizu
  model_replacement: $1
- reg: {}
  regex: ; *(?:Micromax[ _](A111|A240)|(A111|A240)) Build
  regex_flag: i
  device_replacement: Micromax $1$2
  brand_replacement: Micromax
  model_replacement: $1$2
- reg: {}
  regex: ; *Micromax[ _](A\d{2,3}[^;/]*) Build
  regex_flag: i
  device_replacement: Micromax $1
  brand_replacement: Micromax
  model_replacement: $1
- reg: {}
  regex: ; *(A\d{2}|A[12]\d{2}|A90S|A110Q) Build
  regex_flag: i
  device_replacement: Micromax $1
  brand_replacement: Micromax
  model_replacement: $1
- reg: {}
  regex: ; *Micromax[ _](P\d{3}[^;/]*) Build
  regex_flag: i
  device_replacement: Micromax $1
  brand_replacement: Micromax
  model_replacement: $1
- reg: {}
  regex: ; *(P\d{3}|P\d{3}\(Funbook\)) Build
  regex_flag: i
  device_replacement: Micromax $1
  brand_replacement: Micromax
  model_replacement: $1
- reg: {}
  regex: ; *(MITO)[ _\-]?([^;/]+) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: Mito
  model_replacement: $2
- reg: {}
  regex: ; *(Cynus)[ _](F5|T\d|.+?) *(?:Build|[;/\)])
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: Mobistel
  model_replacement: $1 $2
- reg: {}
  regex: ; *(MODECOM |)(FreeTab) ?([^;/]+) Build
  regex_flag: i
  device_replacement: $1$2 $3
  brand_replacement: Modecom
  model_replacement: $2 $3
- reg: {}
  regex: ; *(MODECOM )([^;/]+) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: Modecom
  model_replacement: $2
- reg: {}
  regex: ; *(MZ\d{3}\+?|MZ\d{3} 4G|Xoom|XOOM[^;/]*) Build
  regex_flag: ""
  device_replacement: Motorola $1
  brand_replacement: Motorola
  model_replacement: $1
- reg: {}
  regex: ; *(Milestone )(XT[^;/]*) Build
  regex_flag: ""
  device_replacement: Motorola $1$2
  brand_replacement: Motorola
  model_replacement: $2
- reg: {}
  regex: ; *(Motoroi ?x|Droid X|DROIDX) Build
  regex_flag: i
  device_replacement: Motorola $1
  brand_replacement: Motorola
  model_replacement: DROID X
- reg: {}
  regex: ; *(Droid[^;/]*|DROID[^;/]*|Milestone[^;/]*|Photon|Triumph|Devour|Titanium)
    Build
  regex_flag: ""
  device_replacement: Motorola $1
  brand_replacement: Motorola
  model_replacement: $1
- reg: {}
  regex: ; *(A555|A85[34][^;/]*|A95[356]|ME[58]\d{2}\+?|ME600|ME632|ME722|MB\d{3}\+?|MT680|MT710|MT870|MT887|MT917|WX435|WX453|WX44[25]|XT\d{3,4}[A-Z\+]*|CL[iI]Q|CL[iI]Q
    XT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Motorola
  model_replacement: $1
- reg: {}
  regex: ; *(Motorola MOT-|Motorola[ _\-]|MOT\-?)([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Motorola
  model_replacement: $2
- reg: {}
  regex: ; *(Moto[_ ]?|MOT\-)([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Motorola
  model_replacement: $2
- reg: {}
  regex: ; *((?:MP[DQ]C|MPG\d{1,4}|MP\d{3,4}|MID(?:(?:10[234]|114|43|7[247]|8[24]|7)C|8[01]1))[^;/]*)
    Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Mpman
  model_replacement: $1
- reg: {}
  regex: ; *(?:MSI[ _]|)(Primo\d+|Enjoy[ _\-][^;/]+) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: Msi
  model_replacement: $1
- reg: {}
  regex: ; *Multilaser[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Multilaser
  model_replacement: $1
- reg: {}
  regex: ; *(My)[_]?(Pad)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2 $3
  brand_replacement: MyPhone
  model_replacement: $1$2 $3
- reg: {}
  regex: ; *(My)\|?(Phone)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2 $3
  brand_replacement: MyPhone
  model_replacement: $3
- reg: {}
  regex: ; *(A\d+)[ _](Duo|) Build
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: MyPhone
  model_replacement: $1 $2
- reg: {}
  regex: ; *(myTab[^;/]*) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Mytab
  model_replacement: $1
- reg: {}
  regex: ; *(NABI2?-)([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Nabi
  model_replacement: $2
- reg: {}
  regex: ; *(N-\d+[CDE]) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Nec
  model_replacement: $1
- reg: {}
  regex: ; ?(NEC-)(.*) Build/
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Nec
  model_replacement: $2
- reg: {}
  regex: ; *(LT-NA7) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Nec
  model_replacement: Lifetouch Note
- reg: {}
  regex: '; *(NXM\d+[A-z0-9_]*|Next\d[A-z0-9_ \-]*|NEXT\d[A-z0-9_ \-]*|Nextbook [A-z0-9_
    ]*|DATAM803HC|M805)(?: Build|[\);])'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Nextbook
  model_replacement: $1
- reg: {}
  regex: ; *(Nokia)([ _\-]*)([^;/]*) Build
  regex_flag: i
  device_replacement: $1$2$3
  brand_replacement: Nokia
  model_replacement: $3
- reg: {}
  regex: ; *(Nook ?|Barnes & Noble Nook |BN )([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Nook
  model_replacement: $2
- reg: {}
  regex: ; *(NOOK |)(BNRV200|BNRV200A|BNTV250|BNTV250A|BNTV400|BNTV600|LogicPD Zoom2)
    Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Nook
  model_replacement: $2
- reg: {}
  regex: ; Build/(Nook)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Nook
  model_replacement: Tablet
- reg: {}
  regex: ; *(OP110|OliPad[^;/]+) Build
  regex_flag: ""
  device_replacement: Olivetti $1
  brand_replacement: Olivetti
  model_replacement: $1
- reg: {}
  regex: ; *OMEGA[ _\-](MID[^;/]+) Build
  regex_flag: ""
  device_replacement: Omega $1
  brand_replacement: Omega
  model_replacement: $1
- reg: {}
  regex: ^(MID7500|MID\d+) Mozilla/5\.0 \(iPad;
  regex_flag: ""
  device_replacement: Omega $1
  brand_replacement: Omega
  model_replacement: $1
- reg: {}
  regex: ; *((?:CIUS|cius)[^;/]*) Build
  regex_flag: ""
  device_replacement: Openpeak $1
  brand_replacement: Openpeak
  model_replacement: $1
- reg: {}
  regex: ; *(Find ?(?:5|7a)|R8[012]\d{1,2}|T703\d{0,1}|U70\d{1,2}T?|X90\d{1,2}) Build
  regex_flag: ""
  device_replacement: Oppo $1
  brand_replacement: Oppo
  model_replacement: $1
- reg: {}
  regex: ; *OPPO ?([^;/]+) Build/
  regex_flag: ""
  device_replacement: Oppo $1
  brand_replacement: Oppo
  model_replacement: $1
- reg: {}
  regex: ; *(?:Odys\-|ODYS\-|ODYS )([^;/]+) Build
  regex_flag: ""
  device_replacement: Odys $1
  brand_replacement: Odys
  model_replacement: $1
- reg: {}
  regex: ; *(SELECT) ?(7) Build
  regex_flag: ""
  device_replacement: Odys $1 $2
  brand_replacement: Odys
  model_replacement: $1 $2
- reg: {}
  regex: ; *(PEDI)_(PLUS)_(W) Build
  regex_flag: ""
  device_replacement: Odys $1 $2 $3
  brand_replacement: Odys
  model_replacement: $1 $2 $3
- reg: {}
  regex: ; *(AEON|BRAVIO|FUSION|FUSION2IN1|Genio|EOS10|IEOS[^;/]*|IRON|Loox|LOOX|LOOX
    Plus|Motion|NOON|NOON_PRO|NEXT|OPOS|PEDI[^;/]*|PRIME[^;/]*|STUDYTAB|TABLO|Tablet-PC-4|UNO_X8|XELIO[^;/]*|Xelio
    ?\d+ ?[Pp]ro|XENO10|XPRESS PRO) Build
  regex_flag: ""
  device_replacement: Odys $1
  brand_replacement: Odys
  model_replacement: $1
- reg: {}
  regex: ; (ONE [a-zA-Z]\d+) Build/
  regex_flag: ""
  device_replacement: OnePlus $1
  brand_replacement: OnePlus
  model_replacement: $1
- reg: {}
  regex: '; (ONEPLUS [a-zA-Z]\d+)(?: Build/|)'
  regex_flag: ""
  device_replacement: OnePlus $1
  brand_replacement: OnePlus
  model_replacement: $1
- reg: {}
  regex: ; *(TP-\d+) Build/
  regex_flag: ""
  device_replacement: Orion $1
  brand_replacement: Orion
  model_replacement: $1
- reg: {}
  regex: ; *(G100W?) Build/
  regex_flag: ""
  device_replacement: PackardBell $1
  brand_replacement: PackardBell
  model_replacement: $1
- reg: {}
  regex: ; *(Panasonic)[_ ]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(FZ-A1B|JT-B1) Build
  regex_flag: ""
  device_replacement: Panasonic $1
  brand_replacement: Panasonic
  model_replacement: $1
- reg: {}
  regex: ; *(dL1|DL1) Build
  regex_flag: ""
  device_replacement: Panasonic $1
  brand_replacement: Panasonic
  model_replacement: $1
- reg: {}
  regex: ; *(SKY[ _]|)(IM\-[AT]\d{3}[^;/]+).* Build/
  regex_flag: ""
  device_replacement: Pantech $1$2
  brand_replacement: Pantech
  model_replacement: $1$2
- reg: {}
  regex: '; *((?:ADR8995|ADR910L|ADR930L|ADR930VW|PTL21|P8000)(?: 4G|)) Build/'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Pantech
  model_replacement: $1
- reg: {}
  regex: ; *Pantech([^;/]+).* Build/
  regex_flag: ""
  device_replacement: Pantech $1
  brand_replacement: Pantech
  model_replacement: $1
- reg: {}
  regex: ; *(papyre)[ _\-]([^;/]+) Build/
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: Papyre
  model_replacement: $2
- reg: {}
  regex: ; *(?:Touchlet )?(X10\.[^;/]+) Build/
  regex_flag: ""
  device_replacement: Pearl $1
  brand_replacement: Pearl
  model_replacement: $1
- reg: {}
  regex: ; PHICOMM (i800) Build/
  regex_flag: ""
  device_replacement: Phicomm $1
  brand_replacement: Phicomm
  model_replacement: $1
- reg: {}
  regex: ; PHICOMM ([^;/]+) Build/
  regex_flag: ""
  device_replacement: Phicomm $1
  brand_replacement: Phicomm
  model_replacement: $1
- reg: {}
  regex: ; *(FWS\d{3}[^;/]+) Build/
  regex_flag: ""
  device_replacement: Phicomm $1
  brand_replacement: Phicomm
  model_replacement: $1
- reg: {}
  regex: ; *(D633|D822|D833|T539|T939|V726|W335|W336|W337|W3568|W536|W5510|W626|W632|W6350|W6360|W6500|W732|W736|W737|W7376|W820|W832|W8355|W8500|W8510|W930)
    Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Philips
  model_replacement: $1
- reg: {}
  regex: ; *(?:Philips|PHILIPS)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: Philips $1
  brand_replacement: Philips
  model_replacement: $1
- reg: {}
  regex: Android 4\..*; *(M[12356789]|U[12368]|S[123])\ ?(pro)? Build
  regex_flag: ""
  device_replacement: Pipo $1$2
  brand_replacement: Pipo
  model_replacement: $1$2
- reg: {}
  regex: ; *(MOMO[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Ployer
  model_replacement: $1
- reg: {}
  regex: ; *(?:Polaroid[ _]|)((?:MIDC\d{3,}|PMID\d{2,}|PTAB\d{3,})[^;/]*)(\/[^;/]*|)
    Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Polaroid
  model_replacement: $1
- reg: {}
  regex: ; *(?:Polaroid )(Tablet) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Polaroid
  model_replacement: $1
- reg: {}
  regex: ; *(POMP)[ _\-](.+?) *(?:Build|[;/\)])
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Pomp
  model_replacement: $2
- reg: {}
  regex: ; *(TB07STA|TB10STA|TB07FTA|TB10FTA) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Positivo
  model_replacement: $1
- reg: {}
  regex: ; *(?:Positivo |)((?:YPY|Ypy)[^;/]+) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Positivo
  model_replacement: $1
- reg: {}
  regex: ; *(MOB-[^;/]+) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: POV
  model_replacement: $1
- reg: {}
  regex: ; *POV[ _\-]([^;/]+) Build/
  regex_flag: ""
  device_replacement: POV $1
  brand_replacement: POV
  model_replacement: $1
- reg: {}
  regex: ; *((?:TAB-PLAYTAB|TAB-PROTAB|PROTAB|PlayTabPro|Mobii[ _\-]|TAB-P)[^;/]*)
    Build/
  regex_flag: ""
  device_replacement: POV $1
  brand_replacement: POV
  model_replacement: $1
- reg: {}
  regex: ; *(?:Prestigio |)((?:PAP|PMP)\d[^;/]+) Build/
  regex_flag: ""
  device_replacement: Prestigio $1
  brand_replacement: Prestigio
  model_replacement: $1
- reg: {}
  regex: ; *(PLT[0-9]{4}.*) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Proscan
  model_replacement: $1
- reg: {}
  regex: ; *(A2|A5|A8|A900)_?(Classic|) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Qmobile
  model_replacement: $1 $2
- reg: {}
  regex: ; *(Q[Mm]obile)_([^_]+)_([^_]+) Build
  regex_flag: ""
  device_replacement: Qmobile $2 $3
  brand_replacement: Qmobile
  model_replacement: $2 $3
- reg: {}
  regex: ; *(Q\-?[Mm]obile)[_ ](A[^;/]+) Build
  regex_flag: ""
  device_replacement: Qmobile $2
  brand_replacement: Qmobile
  model_replacement: $2
- reg: {}
  regex: ; *(Q\-Smart)[ _]([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Qmobilevn
  model_replacement: $2
- reg: {}
  regex: ; *(Q\-?[Mm]obile)[ _\-](S[^;/]+) Build/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Qmobilevn
  model_replacement: $2
- reg: {}
  regex: ; *(TA1013) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Quanta
  model_replacement: $1
- reg: {}
  regex: ; (RCT\w+) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: RCA
  model_replacement: $1
- reg: {}
  regex: ; RCA (\w+) Build/
  regex_flag: ""
  device_replacement: RCA $1
  brand_replacement: RCA
  model_replacement: $1
- reg: {}
  regex: ; *(RK\d+),? Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Rockchip
  model_replacement: $1
- reg: {}
  regex: ' Build/(RK\d+)'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Rockchip
  model_replacement: $1
- reg: {}
  regex: ; *(SAMSUNG |Samsung |)((?:Galaxy (?:Note II|S\d)|GT-I9082|GT-I9205|GT-N7\d{3}|SM-N9005)[^;/]*)\/?[^;/]*
    Build/
  regex_flag: ""
  device_replacement: Samsung $1$2
  brand_replacement: Samsung
  model_replacement: $2
- reg: {}
  regex: '; *(Google |)(Nexus [Ss](?: 4G|)) Build/'
  regex_flag: ""
  device_replacement: Samsung $1$2
  brand_replacement: Samsung
  model_replacement: $2
- reg: {}
  regex: ; *(SAMSUNG |Samsung )([^\/]*)\/[^ ]* Build/
  regex_flag: ""
  device_replacement: Samsung $2
  brand_replacement: Samsung
  model_replacement: $2
- reg: {}
  regex: '; *(Galaxy(?: Ace| Nexus| S ?II+|Nexus S| with MCR 1.2| Mini Plus 4G|))
    Build/'
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ; *(SAMSUNG[ _\-]|)(?:SAMSUNG[ _\-])([^;/]+) Build
  regex_flag: ""
  device_replacement: Samsung $2
  brand_replacement: Samsung
  model_replacement: $2
- reg: {}
  regex: ; *(SAMSUNG-|)(GT\-[BINPS]\d{4}[^\/]*)(\/[^ ]*) Build
  regex_flag: ""
  device_replacement: Samsung $1$2$3
  brand_replacement: Samsung
  model_replacement: $2
- reg: {}
  regex: (?:; *|^)((?:GT\-[BIiNPS]\d{4}|I9\d{2}0[A-Za-z\+]?\b)[^;/\)]*?)(?:Build|Linux|MIUI|[;/\)])
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ; (SAMSUNG-)([A-Za-z0-9\-]+).* Build/
  regex_flag: ""
  device_replacement: Samsung $1$2
  brand_replacement: Samsung
  model_replacement: $2
- reg: {}
  regex: ; *((?:SCH|SGH|SHV|SHW|SPH|SC|SM)\-[A-Za-z0-9 ]+)(/?[^ ]*|) Build
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ; *((?:SC)\-[A-Za-z0-9 ]+)(/?[^ ]*|)\)
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ' ((?:SCH)\-[A-Za-z0-9 ]+)(/?[^ ]*|) Build'
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ; *(Behold ?(?:2|II)|YP\-G[^;/]+|EK-GC100|SCL21|I9300) Build
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ; *((?:SCH|SGH|SHV|SHW|SPH|SC|SM)\-[A-Za-z0-9]{5,6})[\)]
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: ; *(SH\-?\d\d[^;/]+|SBM\d[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sharp
  model_replacement: $1
- reg: {}
  regex: ; *(SHARP[ -])([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Sharp
  model_replacement: $2
- reg: {}
  regex: ; *(SPX[_\-]\d[^;/]*) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Simvalley
  model_replacement: $1
- reg: {}
  regex: ; *(SX7\-PEARL\.GmbH) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Simvalley
  model_replacement: $1
- reg: {}
  regex: ; *(SP[T]?\-\d{2}[^;/]*) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Simvalley
  model_replacement: $1
- reg: {}
  regex: ; *(SK\-.*) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: SKtelesys
  model_replacement: $1
- reg: {}
  regex: ; *(?:SKYTEX|SX)-([^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Skytex
  model_replacement: $1
- reg: {}
  regex: ; *(IMAGINE [^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Skytex
  model_replacement: $1
- reg: {}
  regex: ; *(SmartQ) ?([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(WF7C|WF10C|SBT[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Smartbitt
  model_replacement: $1
- reg: {}
  regex: ; *(SBM(?:003SH|005SH|006SH|007SH|102SH)) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sharp
  model_replacement: $1
- reg: {}
  regex: ; *(003P|101P|101P11C|102P) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Panasonic
  model_replacement: $1
- reg: {}
  regex: ; *(00\dZ) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: ZTE
  model_replacement: $1
- reg: {}
  regex: ; HTC(X06HT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: ; *(001HT|X06HT) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: ; *(201M) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Motorola
  model_replacement: XT902
- reg: {}
  regex: ; *(ST\d{4}.*)Build/ST
  regex_flag: ""
  device_replacement: Trekstor $1
  brand_replacement: Trekstor
  model_replacement: $1
- reg: {}
  regex: ; *(ST\d{4}.*) Build/
  regex_flag: ""
  device_replacement: Trekstor $1
  brand_replacement: Trekstor
  model_replacement: $1
- reg: {}
  regex: ; *(Sony ?Ericsson ?)([^;/]+) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: SonyEricsson
  model_replacement: $2
- reg: {}
  regex: ; *((?:SK|ST|E|X|LT|MK|MT|WT)\d{2}[a-z0-9]*(?:-o|)|R800i|U20i) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: SonyEricsson
  model_replacement: $1
- reg: {}
  regex: ; *(Xperia (?:A8|Arc|Acro|Active|Live with Walkman|Mini|Neo|Play|Pro|Ray|X\d+)[^;/]*)
    Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: SonyEricsson
  model_replacement: $1
- reg: {}
  regex: ; Sony (Tablet[^;/]+) Build
  regex_flag: ""
  device_replacement: Sony $1
  brand_replacement: Sony
  model_replacement: $1
- reg: {}
  regex: ; Sony ([^;/]+) Build
  regex_flag: ""
  device_replacement: Sony $1
  brand_replacement: Sony
  model_replacement: $1
- reg: {}
  regex: ; *(Sony)([A-Za-z0-9\-]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(Xperia [^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sony
  model_replacement: $1
- reg: {}
  regex: ; *(C(?:1[0-9]|2[0-9]|53|55|6[0-9])[0-9]{2}|D[25]\d{3}|D6[56]\d{2}) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sony
  model_replacement: $1
- reg: {}
  regex: ; *(SGP\d{3}|SGPT\d{2}) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sony
  model_replacement: $1
- reg: {}
  regex: ; *(NW-Z1000Series) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sony
  model_replacement: $1
- reg: {}
  regex: PLAYSTATION 3
  regex_flag: ""
  device_replacement: PlayStation 3
  brand_replacement: Sony
  model_replacement: PlayStation 3
- reg: {}
  regex: (PlayStation (?:Portable|Vita|\d+))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sony
  model_replacement: $1
- reg: {}
  regex: ; *((?:CSL_Spice|Spice|SPICE|CSL)[ _\-]?|)([Mm][Ii])([ _\-]|)(\d{3}[^;/]*)
    Build/
  regex_flag: ""
  device_replacement: $1$2$3$4
  brand_replacement: Spice
  model_replacement: Mi$4
- reg: {}
  regex: ; *(Sprint )(.+?) *(?:Build|[;/])
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Sprint
  model_replacement: $2
- reg: {}
  regex: '\b(Sprint)[: ]([^;,/ ]+)'
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Sprint
  model_replacement: $2
- reg: {}
  regex: ; *(TAGI[ ]?)(MID) ?([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1$2$3
  brand_replacement: Tagi
  model_replacement: $2$3
- reg: {}
  regex: ; *(Oyster500|Opal 800) Build
  regex_flag: ""
  device_replacement: Tecmobile $1
  brand_replacement: Tecmobile
  model_replacement: $1
- reg: {}
  regex: ; *(TECNO[ _])([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Tecno
  model_replacement: $2
- reg: {}
  regex: '; *Android for (Telechips|Techvision) ([^ ]+) '
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(T-Hub2) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Telstra
  model_replacement: $1
- reg: {}
  regex: ; *(PAD) ?(100[12]) Build/
  regex_flag: ""
  device_replacement: Terra $1$2
  brand_replacement: Terra
  model_replacement: $1$2
- reg: {}
  regex: ; *(T[BM]-\d{3}[^;/]+) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Texet
  model_replacement: $1
- reg: {}
  regex: ; *(tolino [^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Thalia
  model_replacement: $1
- reg: {}
  regex: ; *Build/.* (TOLINO_BROWSER)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Thalia
  model_replacement: Tolino Shine
- reg: {}
  regex: ; *(?:CJ[ -])?(ThL|THL)[ -]([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Thl
  model_replacement: $2
- reg: {}
  regex: ; *(T100|T200|T5|W100|W200|W8s) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Thl
  model_replacement: $1
- reg: {}
  regex: ; *(T-Mobile[ _]G2[ _]Touch) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: Hero
- reg: {}
  regex: ; *(T-Mobile[ _]G2) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: Desire Z
- reg: {}
  regex: ; *(T-Mobile myTouch Q) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Huawei
  model_replacement: U8730
- reg: {}
  regex: ; *(T-Mobile myTouch) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Huawei
  model_replacement: U8680
- reg: {}
  regex: ; *(T-Mobile_Espresso) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: Espresso
- reg: {}
  regex: ; *(T-Mobile G1) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: HTC
  model_replacement: Dream
- reg: {}
  regex: \b(T-Mobile ?|)(myTouch)[ _]?([34]G)[ _]?([^\/]*) (?:Mozilla|Build)
  regex_flag: ""
  device_replacement: $1$2 $3 $4
  brand_replacement: HTC
  model_replacement: $2 $3 $4
- reg: {}
  regex: \b(T-Mobile)_([^_]+)_(.*) Build
  regex_flag: ""
  device_replacement: $1 $2 $3
  brand_replacement: Tmobile
  model_replacement: $2 $3
- reg: {}
  regex: \b(T-Mobile)[_ ]?(.*?)Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Tmobile
  model_replacement: $2
- reg: {}
  regex: ' (ATP[0-9]{4}) Build'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Tomtec
  model_replacement: $1
- reg: {}
  regex: ' *(TOOKY)[ _\-]([^;/]+) ?(?:Build|;)'
  regex_flag: i
  device_replacement: $1 $2
  brand_replacement: Tooky
  model_replacement: $2
- reg: {}
  regex: \b(TOSHIBA_AC_AND_AZ|TOSHIBA_FOLIO_AND_A|FOLIO_AND_A)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Toshiba
  model_replacement: Folio 100
- reg: {}
  regex: ; *([Ff]olio ?100) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Toshiba
  model_replacement: Folio 100
- reg: {}
  regex: ; *(AT[0-9]{2,3}(?:\-A|LE\-A|PE\-A|SE|a|)|AT7-A|AT1S0|Hikari-iFrame/WDPF-[^;/]+|THRiVE|Thrive)
    Build/
  regex_flag: ""
  device_replacement: Toshiba $1
  brand_replacement: Toshiba
  model_replacement: $1
- reg: {}
  regex: ; *(TM-MID\d+[^;/]+|TOUCHMATE|MID-750) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Touchmate
  model_replacement: $1
- reg: {}
  regex: ; *(TM-SM\d+[^;/]+) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Touchmate
  model_replacement: $1
- reg: {}
  regex: ; *(A10 [Bb]asic2?) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Treq
  model_replacement: $1
- reg: {}
  regex: ; *(TREQ[ _\-])([^;/]+) Build
  regex_flag: i
  device_replacement: $1$2
  brand_replacement: Treq
  model_replacement: $2
- reg: {}
  regex: ; *(X-?5|X-?3) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Umeox
  model_replacement: $1
- reg: {}
  regex: ; *(A502\+?|A936|A603|X1|X2) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Umeox
  model_replacement: $1
- reg: {}
  regex: (TOUCH(?:TAB|PAD).+?) Build/
  regex_flag: i
  device_replacement: Versus $1
  brand_replacement: Versus
  model_replacement: $1
- reg: {}
  regex: (VERTU) ([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Vertu
  model_replacement: $2
- reg: {}
  regex: ; *(Videocon)[ _\-]([^;/]+) *(?:Build|;)
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: Videocon
  model_replacement: $2
- reg: {}
  regex: ' (VT\d{2}[A-Za-z]*) Build'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Videocon
  model_replacement: $1
- reg: {}
  regex: ; *((?:ViewPad|ViewPhone|VSD)[^;/]+) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Viewsonic
  model_replacement: $1
- reg: {}
  regex: ; *(ViewSonic-)([^;/]+) Build/
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: Viewsonic
  model_replacement: $2
- reg: {}
  regex: ; *(GTablet.*) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Viewsonic
  model_replacement: $1
- reg: {}
  regex: ; *([Vv]ivo)[ _]([^;/]+) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: vivo
  model_replacement: $2
- reg: {}
  regex: (Vodafone) (.*) Build/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(?:Walton[ _\-]|)(Primo[ _\-][^;/]+) Build
  regex_flag: i
  device_replacement: Walton $1
  brand_replacement: Walton
  model_replacement: $1
- reg: {}
  regex: ; *(?:WIKO[ \-]|)(CINK\+?|BARRY|BLOOM|DARKFULL|DARKMOON|DARKNIGHT|DARKSIDE|FIZZ|HIGHWAY|IGGY|OZZY|RAINBOW|STAIRWAY|SUBLIM|WAX|CINK
    [^;/]+) Build/
  regex_flag: i
  device_replacement: Wiko $1
  brand_replacement: Wiko
  model_replacement: $1
- reg: {}
  regex: ; *WellcoM-([^;/]+) Build
  regex_flag: ""
  device_replacement: Wellcom $1
  brand_replacement: Wellcom
  model_replacement: $1
- reg: {}
  regex: (?:(WeTab)-Browser|; (wetab) Build)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: WeTab
  model_replacement: WeTab
- reg: {}
  regex: ; *(AT-AS[^;/]+) Build
  regex_flag: ""
  device_replacement: Wolfgang $1
  brand_replacement: Wolfgang
  model_replacement: $1
- reg: {}
  regex: ; *(?:Woxter|Wxt) ([^;/]+) Build
  regex_flag: ""
  device_replacement: Woxter $1
  brand_replacement: Woxter
  model_replacement: $1
- reg: {}
  regex: ; *(?:Xenta |Luna |)(TAB[234][0-9]{2}|TAB0[78]-\d{3}|TAB0?9-\d{3}|TAB1[03]-\d{3}|SMP\d{2}-\d{3})
    Build/
  regex_flag: ""
  device_replacement: Yarvik $1
  brand_replacement: Yarvik
  model_replacement: $1
- reg: {}
  regex: '; *([A-Z]{2,4})(M\d{3,}[A-Z]{2})([^;\)\/]*)(?: Build|[;\)])'
  regex_flag: ""
  device_replacement: Yifang $1$2$3
  brand_replacement: Yifang
  model_replacement: $2
- reg: {}
  regex: ; *((Mi|MI|HM|MI-ONE|Redmi)[ -](NOTE |Note |)[^;/]*) (Build|MIUI)/
  regex_flag: ""
  device_replacement: XiaoMi $1
  brand_replacement: XiaoMi
  model_replacement: $1
- reg: {}
  regex: ; *((Mi|MI|HM|MI-ONE|Redmi)[ -](NOTE |Note |)[^;/\)]*)
  regex_flag: ""
  device_replacement: XiaoMi $1
  brand_replacement: XiaoMi
  model_replacement: $1
- reg: {}
  regex: ; *(MIX) (Build|MIUI)/
  regex_flag: ""
  device_replacement: XiaoMi $1
  brand_replacement: XiaoMi
  model_replacement: $1
- reg: {}
  regex: ; *((MIX) ([^;/]*)) (Build|MIUI)/
  regex_flag: ""
  device_replacement: XiaoMi $1
  brand_replacement: XiaoMi
  model_replacement: $1
- reg: {}
  regex: ; *XOLO[ _]([^;/]*tab.*) Build
  regex_flag: i
  device_replacement: Xolo $1
  brand_replacement: Xolo
  model_replacement: $1
- reg: {}
  regex: ; *XOLO[ _]([^;/]+) Build
  regex_flag: i
  device_replacement: Xolo $1
  brand_replacement: Xolo
  model_replacement: $1
- reg: {}
  regex: ; *(q\d0{2,3}[a-z]?) Build
  regex_flag: i
  device_replacement: Xolo $1
  brand_replacement: Xolo
  model_replacement: $1
- reg: {}
  regex: ; *(PAD ?[79]\d+[^;/]*|TelePAD\d+[^;/]) Build
  regex_flag: ""
  device_replacement: Xoro $1
  brand_replacement: Xoro
  model_replacement: $1
- reg: {}
  regex: ; *(?:(?:ZOPO|Zopo)[ _]([^;/]+)|(ZP ?(?:\d{2}[^;/]+|C2))|(C[2379])) Build
  regex_flag: ""
  device_replacement: $1$2$3
  brand_replacement: Zopo
  model_replacement: $1$2$3
- reg: {}
  regex: ; *(ZiiLABS) (Zii[^;/]*) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: ZiiLabs
  model_replacement: $2
- reg: {}
  regex: ; *(Zii)_([^;/]*) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: ZiiLabs
  model_replacement: $2
- reg: {}
  regex: ; *(ARIZONA|(?:ATLAS|Atlas) W|D930|Grand (?:[SX][^;]*|Era|Memo[^;]*)|JOE|(?:Kis|KIS)\b[^;]*|Libra|Light
    [^;]*|N8[056][01]|N850L|N8000|N9[15]\d{2}|N9810|NX501|Optik|(?:Vip )Racer[^;]*|RacerII|RACERII|San
    Francisco[^;]*|V9[AC]|V55|V881|Z[679][0-9]{2}[A-z]?) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: ZTE
  model_replacement: $1
- reg: {}
  regex: ; *([A-Z]\d+)_USA_[^;]* Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: ZTE
  model_replacement: $1
- reg: {}
  regex: ; *(SmartTab\d+)[^;]* Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: ZTE
  model_replacement: $1
- reg: {}
  regex: ; *(?:Blade|BLADE|ZTE-BLADE)([^;/]*) Build
  regex_flag: ""
  device_replacement: ZTE Blade$1
  brand_replacement: ZTE
  model_replacement: Blade$1
- reg: {}
  regex: ; *(?:Skate|SKATE|ZTE-SKATE)([^;/]*) Build
  regex_flag: ""
  device_replacement: ZTE Skate$1
  brand_replacement: ZTE
  model_replacement: Skate$1
- reg: {}
  regex: ; *(Orange |Optimus )(Monte Carlo|San Francisco) Build
  regex_flag: ""
  device_replacement: $1$2
  brand_replacement: ZTE
  model_replacement: $1$2
- reg: {}
  regex: ; *(?:ZXY-ZTE_|ZTE\-U |ZTE[\- _]|ZTE-C[_ ])([^;/]+) Build
  regex_flag: ""
  device_replacement: ZTE $1
  brand_replacement: ZTE
  model_replacement: $1
- reg: {}
  regex: ; (BASE) (lutea|Lutea 2|Tab[^;]*) Build
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: ZTE
  model_replacement: $1 $2
- reg: {}
  regex: ; (Avea inTouch 2|soft stone|tmn smart a7|Movistar[ _]Link) Build
  regex_flag: i
  device_replacement: $1
  brand_replacement: ZTE
  model_replacement: $1
- reg: {}
  regex: ; *(vp9plus)\)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: ZTE
  model_replacement: $1
- reg: {}
  regex: ; ?(Cloud[ _]Z5|z1000|Z99 2G|z99|z930|z999|z990|z909|Z919|z900) Build/
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Zync
  model_replacement: $1
- reg: {}
  regex: ; ?(KFOT|Kindle Fire) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire
  brand_replacement: Amazon
  model_replacement: Kindle Fire
- reg: {}
  regex: ; ?(KFOTE|Amazon Kindle Fire2) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire 2
  brand_replacement: Amazon
  model_replacement: Kindle Fire 2
- reg: {}
  regex: ; ?(KFTT) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HD
  brand_replacement: Amazon
  model_replacement: Kindle Fire HD 7"
- reg: {}
  regex: ; ?(KFJWI) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HD 8.9" WiFi
  brand_replacement: Amazon
  model_replacement: Kindle Fire HD 8.9" WiFi
- reg: {}
  regex: ; ?(KFJWA) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HD 8.9" 4G
  brand_replacement: Amazon
  model_replacement: Kindle Fire HD 8.9" 4G
- reg: {}
  regex: ; ?(KFSOWI) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HD 7" WiFi
  brand_replacement: Amazon
  model_replacement: Kindle Fire HD 7" WiFi
- reg: {}
  regex: ; ?(KFTHWI) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HDX 7" WiFi
  brand_replacement: Amazon
  model_replacement: Kindle Fire HDX 7" WiFi
- reg: {}
  regex: ; ?(KFTHWA) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HDX 7" 4G
  brand_replacement: Amazon
  model_replacement: Kindle Fire HDX 7" 4G
- reg: {}
  regex: ; ?(KFAPWI) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HDX 8.9" WiFi
  brand_replacement: Amazon
  model_replacement: Kindle Fire HDX 8.9" WiFi
- reg: {}
  regex: ; ?(KFAPWA) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire HDX 8.9" 4G
  brand_replacement: Amazon
  model_replacement: Kindle Fire HDX 8.9" 4G
- reg: {}
  regex: ; ?Amazon ([^;/]+) Build\b
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Amazon
  model_replacement: $1
- reg: {}
  regex: ; ?(Kindle) Build\b
  regex_flag: ""
  device_replacement: Kindle
  brand_replacement: Amazon
  model_replacement: Kindle
- reg: {}
  regex: ; ?(Silk)/(\d+)\.(\d+)(?:\.([0-9\-]+)|) Build\b
  regex_flag: ""
  device_replacement: Kindle Fire
  brand_replacement: Amazon
  model_replacement: Kindle Fire$2
- reg: {}
  regex: ' (Kindle)/(\d+\.\d+)'
  regex_flag: ""
  device_replacement: Kindle
  brand_replacement: Amazon
  model_replacement: $1 $2
- reg: {}
  regex: ' (Silk|Kindle)/(\d+)\.'
  regex_flag: ""
  device_replacement: Kindle
  brand_replacement: Amazon
  model_replacement: Kindle
- reg: {}
  regex: (sprd)\-([^/]+)/
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: ; *(H\d{2}00\+?) Build
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Hero
  model_replacement: $1
- reg: {}
  regex: ; *(iphone|iPhone5) Build/
  regex_flag: ""
  device_replacement: Xianghe $1
  brand_replacement: Xianghe
  model_replacement: $1
- reg: {}
  regex: ; *(e\d{4}[a-z]?_?v\d+|v89_[^;/]+)[^;/]+ Build/
  regex_flag: ""
  device_replacement: Xianghe $1
  brand_replacement: Xianghe
  model_replacement: $1
- reg: {}
  regex: \bUSCC[_\-]?([^ ;/\)]+)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Cellular
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|)(?:ALCATEL)[^;]*;
    *([^;,\)]+)
  regex_flag: ""
  device_replacement: Alcatel $1
  brand_replacement: Alcatel
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|WpsLondonTest;
    ?|)(?:ASUS|Asus)[^;]*; *([^;,\)]+)
  regex_flag: ""
  device_replacement: Asus $1
  brand_replacement: Asus
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|)(?:DELL|Dell)[^;]*;
    *([^;,\)]+)
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|WpsLondonTest;
    ?|)(?:HTC|Htc|HTC_blocked[^;]*)[^;]*; *(?:HTC|)([^;,\)]+)
  regex_flag: ""
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|)(?:HUAWEI)[^;]*;
    *(?:HUAWEI |)([^;,\)]+)
  regex_flag: ""
  device_replacement: Huawei $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|)(?:LG|Lg)[^;]*;
    *(?:LG[ \-]|)([^;,\)]+)
  regex_flag: ""
  device_replacement: LG $1
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|)(?:rv:11;
    |)(?:NOKIA|Nokia)[^;]*; *(?:NOKIA ?|Nokia ?|LUMIA ?|[Ll]umia ?|)(\d{3,10}[^;\)]*)
  regex_flag: ""
  device_replacement: Lumia $1
  brand_replacement: Nokia
  model_replacement: Lumia $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|)(?:NOKIA|Nokia)[^;]*;
    *(RM-\d{3,})
  regex_flag: ""
  device_replacement: Nokia $1
  brand_replacement: Nokia
  model_replacement: $1
- reg: {}
  regex: (?:Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)]|WPDesktop;) ?(?:ARM; ?Touch;
    ?|Touch; ?|)(?:NOKIA|Nokia)[^;]*; *(?:NOKIA ?|Nokia ?|LUMIA ?|[Ll]umia ?|)([^;\)]+)
  regex_flag: ""
  device_replacement: Nokia $1
  brand_replacement: Nokia
  model_replacement: $1
- reg: {}
  regex: 'Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch;
    ?|)(?:Microsoft(?: Corporation|))[^;]*; *([^;,\)]+)'
  regex_flag: ""
  device_replacement: Microsoft $1
  brand_replacement: Microsoft
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|WpsLondonTest;
    ?|)(?:SAMSUNG)[^;]*; *(?:SAMSUNG |)([^;,\.\)]+)
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|WpsLondonTest;
    ?|)(?:TOSHIBA|FujitsuToshibaMobileCommun)[^;]*; *([^;,\)]+)
  regex_flag: ""
  device_replacement: Toshiba $1
  brand_replacement: Toshiba
  model_replacement: $1
- reg: {}
  regex: Windows Phone [^;]+; .*?IEMobile/[^;\)]+[;\)] ?(?:ARM; ?Touch; ?|Touch; ?|WpsLondonTest;
    ?|)([^;]+); *([^;,\)]+)
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: (?:^|; )SAMSUNG\-([A-Za-z0-9\-]+).* Bada/
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: \(Mobile; ALCATEL ?(One|ONE) ?(Touch|TOUCH) ?([^;/]+)(?:/[^;]+|); rv:[^\)]+\)
    Gecko/[^\/]+ Firefox/
  regex_flag: ""
  device_replacement: Alcatel $1 $2 $3
  brand_replacement: Alcatel
  model_replacement: One Touch $3
- reg: {}
  regex: \(Mobile; (?:ZTE([^;]+)|(OpenC)); rv:[^\)]+\) Gecko/[^\/]+ Firefox/
  regex_flag: ""
  device_replacement: ZTE $1$2
  brand_replacement: ZTE
  model_replacement: $1$2
- reg: {}
  regex: Nokia(N[0-9]+)([A-z_\-][A-z0-9_\-]*)
  regex_flag: ""
  device_replacement: Nokia $1
  brand_replacement: Nokia
  model_replacement: $1$2
- reg: {}
  regex: (?:NOKIA|Nokia)(?:\-| *)(?:([A-Za-z0-9]+)\-[0-9a-f]{32}|([A-Za-z0-9\-]+)(?:UCBrowser)|([A-Za-z0-9\-]+))
  regex_flag: ""
  device_replacement: Nokia $1$2$3
  brand_replacement: Nokia
  model_replacement: $1$2$3
- reg: {}
  regex: Lumia ([A-Za-z0-9\-]+)
  regex_flag: ""
  device_replacement: Lumia $1
  brand_replacement: Nokia
  model_replacement: Lumia $1
- reg: {}
  regex: \(Symbian; U; S60 V5; [A-z]{2}\-[A-z]{2}; (SonyEricsson|Samsung|Nokia|LG)([^;/]+)\)
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: \(Symbian(?:/3|); U; ([^;]+);
  regex_flag: ""
  device_replacement: Nokia $1
  brand_replacement: Nokia
  model_replacement: $1
- reg: {}
  regex: BB10; ([A-Za-z0-9\- ]+)\)
  regex_flag: ""
  device_replacement: BlackBerry $1
  brand_replacement: BlackBerry
  model_replacement: $1
- reg: {}
  regex: Play[Bb]ook.+RIM Tablet OS
  regex_flag: ""
  device_replacement: BlackBerry Playbook
  brand_replacement: BlackBerry
  model_replacement: Playbook
- reg: {}
  regex: Black[Bb]erry ([0-9]+);
  regex_flag: ""
  device_replacement: BlackBerry $1
  brand_replacement: BlackBerry
  model_replacement: $1
- reg: {}
  regex: Black[Bb]erry([0-9]+)
  regex_flag: ""
  device_replacement: BlackBerry $1
  brand_replacement: BlackBerry
  model_replacement: $1
- reg: {}
  regex: Black[Bb]erry;
  regex_flag: ""
  device_replacement: BlackBerry
  brand_replacement: BlackBerry
  model_replacement: $1
- reg: {}
  regex: (Pre|Pixi)/\d+\.\d+
  regex_flag: ""
  device_replacement: Palm $1
  brand_replacement: Palm
  model_replacement: $1
- reg: {}
  regex: Palm([0-9]+)
  regex_flag: ""
  device_replacement: Palm $1
  brand_replacement: Palm
  model_replacement: $1
- reg: {}
  regex: Treo([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Palm Treo $1
  brand_replacement: Palm
  model_replacement: Treo $1
- reg: {}
  regex: webOS.*(P160U(?:NA|))/(\d+).(\d+)
  regex_flag: ""
  device_replacement: HP Veer
  brand_replacement: HP
  model_replacement: Veer
- reg: {}
  regex: (Touch[Pp]ad)/\d+\.\d+
  regex_flag: ""
  device_replacement: HP TouchPad
  brand_replacement: HP
  model_replacement: TouchPad
- reg: {}
  regex: HPiPAQ([A-Za-z0-9]+)/\d+.\d+
  regex_flag: ""
  device_replacement: HP iPAQ $1
  brand_replacement: HP
  model_replacement: iPAQ $1
- reg: {}
  regex: PDA; (PalmOS)/sony/model ([a-z]+)/Revision
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Sony
  model_replacement: $1 $2
- reg: {}
  regex: (Apple\s?TV)
  regex_flag: ""
  device_replacement: AppleTV
  brand_replacement: Apple
  model_replacement: AppleTV
- reg: {}
  regex: (QtCarBrowser)
  regex_flag: ""
  device_replacement: Tesla Model S
  brand_replacement: Tesla
  model_replacement: Model S
- reg: {}
  regex: (iPhone|iPad|iPod)(\d+,\d+)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Apple
  model_replacement: $1$2
- reg: {}
  regex: (iPad)(?:;| Simulator;)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Apple
  model_replacement: $1
- reg: {}
  regex: (iPod)(?:;| touch;| Simulator;)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Apple
  model_replacement: $1
- reg: {}
  regex: (iPhone)(?:;| Simulator;)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Apple
  model_replacement: $1
- reg: {}
  regex: (Watch)(\d+,\d+)
  regex_flag: ""
  device_replacement: Apple $1
  brand_replacement: Apple
  model_replacement: Apple $1 $2
- reg: {}
  regex: (Apple Watch)(?:;| Simulator;)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Apple
  model_replacement: $1
- reg: {}
  regex: (HomePod)(?:;| Simulator;)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Apple
  model_replacement: $1
- reg: {}
  regex: iPhone
  regex_flag: ""
  device_replacement: iPhone
  brand_replacement: Apple
  model_replacement: iPhone
- reg: {}
  regex: CFNetwork/.* Darwin/\d.*\(((?:Mac|iMac|PowerMac|PowerBook)[^\d]*)(\d+)(?:,|%2C)(\d+)
  regex_flag: ""
  device_replacement: $1$2,$3
  brand_replacement: Apple
  model_replacement: $1$2,$3
- reg: {}
  regex: CFNetwork/.* Darwin/\d+\.\d+\.\d+ \(x86_64\)
  regex_flag: ""
  device_replacement: Mac
  brand_replacement: Apple
  model_replacement: Mac
- reg: {}
  regex: CFNetwork/.* Darwin/\d
  regex_flag: ""
  device_replacement: iOS-Device
  brand_replacement: Apple
  model_replacement: iOS-Device
- reg: {}
  regex: Outlook-(iOS)/\d+\.\d+\.prod\.iphone
  regex_flag: ""
  device_replacement: iPhone
  brand_replacement: Apple
  model_replacement: iPhone
- reg: {}
  regex: acer_([A-Za-z0-9]+)_
  regex_flag: ""
  device_replacement: Acer $1
  brand_replacement: Acer
  model_replacement: $1
- reg: {}
  regex: (?:ALCATEL|Alcatel)-([A-Za-z0-9\-]+)
  regex_flag: ""
  device_replacement: Alcatel $1
  brand_replacement: Alcatel
  model_replacement: $1
- reg: {}
  regex: (?:Amoi|AMOI)\-([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Amoi $1
  brand_replacement: Amoi
  model_replacement: $1
- reg: {}
  regex: (?:; |\/|^)((?:Transformer (?:Pad|Prime) |Transformer |PadFone[ _]?)[A-Za-z0-9]*)
  regex_flag: ""
  device_replacement: Asus $1
  brand_replacement: Asus
  model_replacement: $1
- reg: {}
  regex: (?:asus.*?ASUS|Asus|ASUS|asus)[\- ;]*((?:Transformer (?:Pad|Prime) |Transformer
    |Padfone |Nexus[ _]|)[A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Asus $1
  brand_replacement: Asus
  model_replacement: $1
- reg: {}
  regex: (?:ASUS)_([A-Za-z0-9\-]+)
  regex_flag: ""
  device_replacement: Asus $1
  brand_replacement: Asus
  model_replacement: $1
- reg: {}
  regex: \bBIRD[ \-\.]([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Bird $1
  brand_replacement: Bird
  model_replacement: $1
- reg: {}
  regex: \bDell ([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Dell $1
  brand_replacement: Dell
  model_replacement: $1
- reg: {}
  regex: DoCoMo/2\.0 ([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: DoCoMo $1
  brand_replacement: DoCoMo
  model_replacement: $1
- reg: {}
  regex: ([A-Za-z0-9]+)_W;FOMA
  regex_flag: ""
  device_replacement: DoCoMo $1
  brand_replacement: DoCoMo
  model_replacement: $1
- reg: {}
  regex: ([A-Za-z0-9]+);FOMA
  regex_flag: ""
  device_replacement: DoCoMo $1
  brand_replacement: DoCoMo
  model_replacement: $1
- reg: {}
  regex: \b(?:HTC/|HTC/[a-z0-9]+/|)HTC[ _\-;]? *(.*?)(?:-?Mozilla|fingerPrint|[;/\(\)]|$)
  regex_flag: ""
  device_replacement: HTC $1
  brand_replacement: HTC
  model_replacement: $1
- reg: {}
  regex: Huawei([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Huawei $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: HUAWEI-([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Huawei $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: HUAWEI ([A-Za-z0-9\-]+)
  regex_flag: ""
  device_replacement: Huawei $1
  brand_replacement: Huawei
  model_replacement: $1
- reg: {}
  regex: vodafone([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Huawei Vodafone $1
  brand_replacement: Huawei
  model_replacement: Vodafone $1
- reg: {}
  regex: i\-mate ([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: i-mate $1
  brand_replacement: i-mate
  model_replacement: $1
- reg: {}
  regex: Kyocera\-([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Kyocera $1
  brand_replacement: Kyocera
  model_replacement: $1
- reg: {}
  regex: KWC\-([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Kyocera $1
  brand_replacement: Kyocera
  model_replacement: $1
- reg: {}
  regex: Lenovo[_\-]([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Lenovo $1
  brand_replacement: Lenovo
  model_replacement: $1
- reg: {}
  regex: (HbbTV)/[0-9]+\.[0-9]+\.[0-9]+ \([^;]*; *(LG)E *; *([^;]*) *;[^;]*;[^;]*;\)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: $2
  model_replacement: $3
- reg: {}
  regex: (HbbTV)/1\.1\.1.*CE-HTML/1\.\d;(Vendor/|)(THOM[^;]*?)[;\s].{0,30}(LF[^;]+);?
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Thomson
  model_replacement: $4
- reg: {}
  regex: '(HbbTV)(?:/1\.1\.1|) ?(?: \(;;;;;\)|); *CE-HTML(?:/1\.\d|); *([^ ]+) ([^;]+);'
  regex_flag: ""
  device_replacement: $1
  brand_replacement: $2
  model_replacement: $3
- reg: {}
  regex: (HbbTV)/1\.1\.1 \(;;;;;\) Maple_2011
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: (HbbTV)/[0-9]+\.[0-9]+\.[0-9]+ \([^;]*; *(?:CUS:([^;]*)|([^;]+)) *; *([^;]*)
    *;.*;
  regex_flag: ""
  device_replacement: $1
  brand_replacement: $2$3
  model_replacement: $4
- reg: {}
  regex: (HbbTV)/[0-9]+\.[0-9]+\.[0-9]+
  regex_flag: ""
  device_replacement: $1
  brand_replacement: ""
  model_replacement: $1
- reg: {}
  regex: LGE; (?:Media\/|)([^;]*);[^;]*;[^;]*;?\); "?LG NetCast(\.TV|\.Media|)-\d+
  regex_flag: ""
  device_replacement: NetCast$2
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: InettvBrowser/[0-9]+\.[0-9A-Z]+ \([^;]*;(Sony)([^;]*);[^;]*;[^\)]*\)
  regex_flag: ""
  device_replacement: Inettv
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: InettvBrowser/[0-9]+\.[0-9A-Z]+ \([^;]*;([^;]*);[^;]*;[^\)]*\)
  regex_flag: ""
  device_replacement: Inettv
  brand_replacement: Generic_Inettv
  model_replacement: $1
- reg: {}
  regex: (?:InettvBrowser|TSBNetTV|NETTV|HBBTV)
  regex_flag: ""
  device_replacement: Inettv
  brand_replacement: Generic_Inettv
  model_replacement: $1
- reg: {}
  regex: Series60/\d\.\d (LG)[\-]?([A-Za-z0-9 \-]+)
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: \b(?:LGE[ \-]LG\-(?:AX|)|LGE |LGE?-LG|LGE?[ \-]|LG[ /\-]|lg[\-])([A-Za-z0-9]+)\b
  regex_flag: ""
  device_replacement: LG $1
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: (?:^LG[\-]?|^LGE[\-/]?)([A-Za-z]+[0-9]+[A-Za-z]*)
  regex_flag: ""
  device_replacement: LG $1
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: ^LG([0-9]+[A-Za-z]*)
  regex_flag: ""
  device_replacement: LG $1
  brand_replacement: LG
  model_replacement: $1
- reg: {}
  regex: (KIN\.[^ ]+) (\d+)\.(\d+)
  regex_flag: ""
  device_replacement: Microsoft $1
  brand_replacement: Microsoft
  model_replacement: $1
- reg: {}
  regex: (?:MSIE|XBMC).*\b(Xbox)\b
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Microsoft
  model_replacement: $1
- reg: {}
  regex: ; ARM; Trident/6\.0; Touch[\);]
  regex_flag: ""
  device_replacement: Microsoft Surface RT
  brand_replacement: Microsoft
  model_replacement: Surface RT
- reg: {}
  regex: Motorola\-([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Motorola $1
  brand_replacement: Motorola
  model_replacement: $1
- reg: {}
  regex: MOTO\-([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Motorola $1
  brand_replacement: Motorola
  model_replacement: $1
- reg: {}
  regex: MOT\-([A-z0-9][A-z0-9\-]*)
  regex_flag: ""
  device_replacement: Motorola $1
  brand_replacement: Motorola
  model_replacement: $1
- reg: {}
  regex: Nintendo WiiU
  regex_flag: ""
  device_replacement: Nintendo Wii U
  brand_replacement: Nintendo
  model_replacement: Wii U
- reg: {}
  regex: Nintendo (DS|3DS|DSi|Wii);
  regex_flag: ""
  device_replacement: Nintendo $1
  brand_replacement: Nintendo
  model_replacement: $1
- reg: {}
  regex: (?:Pantech|PANTECH)[ _-]?([A-Za-z0-9\-]+)
  regex_flag: ""
  device_replacement: Pantech $1
  brand_replacement: Pantech
  model_replacement: $1
- reg: {}
  regex: Philips([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Philips $1
  brand_replacement: Philips
  model_replacement: $1
- reg: {}
  regex: Philips ([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Philips $1
  brand_replacement: Philips
  model_replacement: $1
- reg: {}
  regex: '(SMART-TV); .* Tizen '
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: SymbianOS/9\.\d.* Samsung[/\-]([A-Za-z0-9 \-]+)
  regex_flag: ""
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: (Samsung)(SGH)(i[0-9]+)
  regex_flag: ""
  device_replacement: $1 $2$3
  brand_replacement: $1
  model_replacement: $2-$3
- reg: {}
  regex: SAMSUNG-ANDROID-MMS/([^;/]+)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: SAMSUNG(?:; |[ -/])([A-Za-z0-9\-]+)
  regex_flag: i
  device_replacement: Samsung $1
  brand_replacement: Samsung
  model_replacement: $1
- reg: {}
  regex: (Dreamcast)
  regex_flag: ""
  device_replacement: Sega $1
  brand_replacement: Sega
  model_replacement: $1
- reg: {}
  regex: ^SIE-([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Siemens $1
  brand_replacement: Siemens
  model_replacement: $1
- reg: {}
  regex: Softbank/[12]\.0/([A-Za-z0-9]+)
  regex_flag: ""
  device_replacement: Softbank $1
  brand_replacement: Softbank
  model_replacement: $1
- reg: {}
  regex: SonyEricsson ?([A-Za-z0-9\-]+)
  regex_flag: ""
  device_replacement: Ericsson $1
  brand_replacement: SonyEricsson
  model_replacement: $1
- reg: {}
  regex: Android [^;]+; ([^ ]+) (Sony)/
  regex_flag: ""
  device_replacement: $2 $1
  brand_replacement: $2
  model_replacement: $1
- reg: {}
  regex: (Sony)(?:BDP\/|\/|)([^ /;\)]+)[ /;\)]
  regex_flag: ""
  device_replacement: $1 $2
  brand_replacement: $1
  model_replacement: $2
- reg: {}
  regex: Puffin/[\d\.]+IT
  regex_flag: ""
  device_replacement: iPad
  brand_replacement: Apple
  model_replacement: iPad
- reg: {}
  regex: Puffin/[\d\.]+IP
  regex_flag: ""
  device_replacement: iPhone
  brand_replacement: Apple
  model_replacement: iPhone
- reg: {}
  regex: Puffin/[\d\.]+AT
  regex_flag: ""
  device_replacement: Generic Tablet
  brand_replacement: Generic
  model_replacement: Tablet
- reg: {}
  regex: Puffin/[\d\.]+AP
  regex_flag: ""
  device_replacement: Generic Smartphone
  brand_replacement: Generic
  model_replacement: Smartphone
- reg: {}
  regex: Android[\- ][\d]+\.[\d]+; [A-Za-z]{2}\-[A-Za-z]{0,2}; WOWMobile (.+)( Build[/
    ]|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Android
  model_replacement: $1
- reg: {}
  regex: Android[\- ][\d]+\.[\d]+\-update1; [A-Za-z]{2}\-[A-Za-z]{0,2} *; *(.+?)(
    Build[/ ]|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Android
  model_replacement: $1
- reg: {}
  regex: Android[\- ][\d]+(?:\.[\d]+)(?:\.[\d]+|); *[A-Za-z]{2}[_\-][A-Za-z]{0,2}\-?
    *; *(.+?)( Build[/ ]|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Android
  model_replacement: $1
- reg: {}
  regex: Android[\- ][\d]+(?:\.[\d]+)(?:\.[\d]+|); *[A-Za-z]{0,2}\- *; *(.+?)( Build[/
    ]|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Android
  model_replacement: $1
- reg: {}
  regex: Android[\- ][\d]+(?:\.[\d]+)(?:\.[\d]+|); *[a-z]{0,2}[_\-]?[A-Za-z]{0,2};?(
    Build[/ ]|\))
  regex_flag: ""
  device_replacement: Generic Smartphone
  brand_replacement: Generic
  model_replacement: Smartphone
- reg: {}
  regex: Android[\- ][\d]+(?:\.[\d]+)(?:\.[\d]+|); *\-?[A-Za-z]{2}; *(.+?)( Build[/
    ]|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Android
  model_replacement: $1
- reg: {}
  regex: Android[\- ][\d]+(?:\.[\d]+)(?:\.[\d]+|)(?:;.*|); *(.+?)( Build[/ ]|\))
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Android
  model_replacement: $1
- reg: {}
  regex: (GoogleTV)
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Inettv
  model_replacement: $1
- reg: {}
  regex: (WebTV)/\d+.\d+
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Inettv
  model_replacement: $1
- reg: {}
  regex: ^(Roku)/DVP-\d+\.\d+
  regex_flag: ""
  device_replacement: $1
  brand_replacement: Generic_Inettv
  model_replacement: $1
- reg: {}
  regex: (Android 3\.\d|Opera Tablet|Tablet; .+Firefox/|Android.*(?:Tab|Pad))
  regex_flag: i
  device_replacement: Generic Tablet
  brand_replacement: Generic
  model_replacement: Tablet
- reg: {}
  regex: (Symbian|\bS60(Version|V\d)|\bS60\b|\((Series 60|Windows Mobile|Palm OS|Bada);
    Opera Mini|Windows CE|Opera Mobi|BREW|Brew|Mobile; .+Firefox/|iPhone OS|Android|MobileSafari|Windows
    *Phone|\(webOS/|PalmOS)
  regex_flag: ""
  device_replacement: Generic Smartphone
  brand_replacement: Generic
  model_replacement: Smartphone
- reg: {}
  regex: (hiptop|avantgo|plucker|xiino|blazer|elaine)
  regex_flag: i
  device_replacement: Generic Smartphone
  brand_replacement: Generic
  model_replacement: Smartphone
- reg: {}
  regex: (bot|BUbiNG|zao|borg|DBot|oegp|silk|Xenu|zeal|^NING|CCBot|crawl|htdig|lycos|slurp|teoma|voila|yahoo|Sogou|CiBra|Nutch|^Java/|^JNLP/|Daumoa|Daum|Genieo|ichiro|larbin|pompos|Scrapy|snappy|speedy|spider|msnbot|msrbot|vortex|^vortex|crawler|favicon|indexer|Riddler|scooter|scraper|scrubby|WhatWeb|WinHTTP|bingbot|BingPreview|openbot|gigabot|furlbot|polybot|seekbot|^voyager|archiver|Icarus6j|mogimogi|Netvibes|blitzbot|altavista|charlotte|findlinks|Retreiver|TLSProber|WordPress|SeznamBot|ProoXiBot|wsr\-agent|Squrl
    Java|EtaoSpider|PaperLiBot|SputnikBot|A6\-Indexer|netresearch|searchsight|baiduspider|YisouSpider|ICC\-Crawler|http%20client|Python-urllib|dataparksearch|converacrawler|Screaming
    Frog|AppEngine-Google|YahooCacheSystem|fast\-webcrawler|Sogou Pic Spider|semanticdiscovery|Innovazion
    Crawler|facebookexternalhit|Google.*/\+/web/snippet|Google-HTTP-Java-Client|BlogBridge|IlTrovatore-Setaccio|InternetArchive|GomezAgent|WebThumbnail|heritrix|NewsGator|PagePeeker|Reaper|ZooShot|holmes|NL-Crawler|Pingdom|StatusCake|WhatsApp|masscan|Google
    Web Preview|Qwantify|Yeti)
  regex_flag: i
  device_replacement: Spider
  brand_replacement: Spider
  model_replacement: Desktop
- reg: {}
  regex: ^(1207|3gso|4thp|501i|502i|503i|504i|505i|506i|6310|6590|770s|802s|a wa|acer|acs\-|airn|alav|asus|attw|au\-m|aur
    |aus |abac|acoo|aiko|alco|alca|amoi|anex|anny|anyw|aptu|arch|argo|bmobile|bell|bird|bw\-n|bw\-u|beck|benq|bilb|blac|c55/|cdm\-|chtm|capi|comp|cond|dall|dbte|dc\-s|dica|ds\-d|ds12|dait|devi|dmob|doco|dopo|dorado|el(?:38|39|48|49|50|55|58|68)|el[3456]\d{2}dual|erk0|esl8|ex300|ez40|ez60|ez70|ezos|ezze|elai|emul|eric|ezwa|fake|fly\-|fly_|g\-mo|g1
    u|g560|gf\-5|grun|gene|go.w|good|grad|hcit|hd\-m|hd\-p|hd\-t|hei\-|hp i|hpip|hs\-c|htc
    |htc\-|htca|htcg)
  regex_flag: i
  device_replacement: Generic Feature Phone
  brand_replacement: Generic
  model_replacement: Feature Phone
- reg: {}
  regex: ^(htcp|htcs|htct|htc_|haie|hita|huaw|hutc|i\-20|i\-go|i\-ma|i\-mobile|i230|iac|iac\-|iac/|ig01|im1k|inno|iris|jata|kddi|kgt|kgt/|kpt
    |kwc\-|klon|lexi|lg g|lg\-a|lg\-b|lg\-c|lg\-d|lg\-f|lg\-g|lg\-k|lg\-l|lg\-m|lg\-o|lg\-p|lg\-s|lg\-t|lg\-u|lg\-w|lg/k|lg/l|lg/u|lg50|lg54|lge\-|lge/|leno|m1\-w|m3ga|m50/|maui|mc01|mc21|mcca|medi|meri|mio8|mioa|mo01|mo02|mode|modo|mot
    |mot\-|mt50|mtp1|mtv |mate|maxo|merc|mits|mobi|motv|mozz|n100|n101|n102|n202|n203|n300|n302|n500|n502|n505|n700|n701|n710|nec\-|nem\-|newg|neon)
  regex_flag: i
  device_replacement: Generic Feature Phone
  brand_replacement: Generic
  model_replacement: Feature Phone
- reg: {}
  regex: ^(netf|noki|nzph|o2 x|o2\-x|opwv|owg1|opti|oran|ot\-s|p800|pand|pg\-1|pg\-2|pg\-3|pg\-6|pg\-8|pg\-c|pg13|phil|pn\-2|pt\-g|palm|pana|pire|pock|pose|psio|qa\-a|qc\-2|qc\-3|qc\-5|qc\-7|qc07|qc12|qc21|qc32|qc60|qci\-|qwap|qtek|r380|r600|raks|rim9|rove|s55/|sage|sams|sc01|sch\-|scp\-|sdk/|se47|sec\-|sec0|sec1|semc|sgh\-|shar|sie\-|sk\-0|sl45|slid|smb3|smt5|sp01|sph\-|spv
    |spv\-|sy01|samm|sany|sava|scoo|send|siem|smar|smit|soft|sony|t\-mo|t218|t250|t600|t610|t618|tcl\-|tdg\-|telm|tim\-|ts70|tsm\-|tsm3|tsm5|tx\-9|tagt)
  regex_flag: i
  device_replacement: Generic Feature Phone
  brand_replacement: Generic
  model_replacement: Feature Phone
- reg: {}
  regex: ^(talk|teli|topl|tosh|up.b|upg1|utst|v400|v750|veri|vk\-v|vk40|vk50|vk52|vk53|vm40|vx98|virg|vertu|vite|voda|vulc|w3c
    |w3c\-|wapj|wapp|wapu|wapm|wig |wapi|wapr|wapv|wapy|wapa|waps|wapt|winc|winw|wonu|x700|xda2|xdag|yas\-|your|zte\-|zeto|aste|audi|avan|blaz|brew|brvw|bumb|ccwa|cell|cldc|cmd\-|dang|eml2|fetc|hipt|http|ibro|idea|ikom|ipaq|jbro|jemu|jigs|keji|kyoc|kyok|libw|m\-cr|midp|mmef|moto|mwbp|mywa|newt|nok6|o2im|pant|pdxg|play|pluc|port|prox|rozo|sama|seri|smal|symb|treo|upsi|vx52|vx53|vx60|vx61|vx70|vx80|vx81|vx83|vx85|wap\-|webc|whit|wmlb|xda\-|xda_)
  regex_flag: i
  device_replacement: Generic Feature Phone
  brand_replacement: Generic
  model_replacement: Feature Phone
- reg: {}
  regex: ^(Ice)$
  regex_flag: ""
  device_replacement: Generic Feature Phone
  brand_replacement: Generic
  model_replacement: Feature Phone
- reg: {}
  regex: (wap[\-\ ]browser|maui|netfront|obigo|teleca|up\.browser|midp|Opera Mini)
  regex_flag: i
  device_replacement: Generic Feature Phone
  brand_replacement: Generic
  model_replacement: Feature Phone`)
