package test

import (
	"os"
	"io/ioutil"
	"bufio"
	// "bytes"
	"time"
	// "path"
	"path/filepath"
	"net/url"
	"encoding/json"
	"encoding/base64"
	"crypto/md5"
	// "encoding/hex"
	"strings"
	"strconv"
	"fmt"
	"log"
	"flag"
	// "github.com/spf13/cobra"
	"go-necro/config"
)

var (
	f *os.File
	fi *os.File
	fo *os.File
	err error
)

// Flag Names
const (
	nmConfig           = "config"
)

var (
	help         = flag.Bool("help", false, "show the usage")
	// version      = flagBoolean(nmVersion, false, "print version information and exit")
	configPath   = flag.String(nmConfig, "", "config file path")
)

const (
	ASSET_ROOT = "https://cdn-r18.necro-sm.dmmgames.com"
	// ASSET_URI = "/secure/data/webgl/resources/"
	// ASSET_URI = "/secure/data/android/resources/"
	APP_KEY = "B3u5C3Y6kqw5jfUwH4SurNcjXV44rHyr"
	MANIFEST_DIR = "manifest/"
	DOWNLOAD_DIR = "download/"
	ARIA2_FILE_EXT = ".aria2.list"
)

var (
	// MANIFEST_ASSET_NAME_LIST_FILE = filepath.Join(MANIFEST_DIR, "asset_name_test.txt")
	MANIFEST_ASSET_NAME_LIST_FILE = filepath.Join(MANIFEST_DIR, "asset_name_test_9.txt")
)

var (
	ASSET_URI_MAP = map[string]string{
		"webgl": "/secure/data/webgl/resources/",
		"android": "/secure/data/android/resources/",
	}
	ASSET_ROOT_PATH_PLATFORM_MAP = map[string]string{
		"webgl": "/secure/data/webgl/resources/",
		"android": "/secure/data/android/resources/",
	}
)

type Manifest struct {
	Data	[]ManifestData	`json:"d"`
}

type ManifestData struct {
	Name	string	`json:"n"`
	V	string	`json:"v"`
	O	string	`json:"o"`
	S	string	`json:"s"`
	// Size	string	`json:"s"`
	H	string	`json:"h"`
	// Hash	string	`json:"h"`
	C	string	`json:"c"`
	D	string	`json:"d"`
	I	string	`json:"i"`
	F	string	`json:"f"`

}

type AssetBundle struct {
	Id		int
	Key		string
	Ver		string
	FileName	string
	Hash		string
	Url		string
	AssetOutPut	string
}

func inLocTimestamp(inLoc time.Time) int64 {
	_, offset := inLoc.Zone()
	ts := inLoc.Add(time.Duration(offset) * time.Second)
	return ts.Unix()
}

func inLocTimestampMilli(inLoc time.Time) int64 {
	_, offset := inLoc.Zone()
	ts := inLoc.Add(time.Duration(offset) * time.Second)
	return ts.UnixMilli()
}

func getTimeStampString() string {
	// testLocTimestamp()
	timeZone, _ := time.LoadLocation("Asia/Tokyo")

	now := time.Now().Round(0)
	nowInLoc := now.In(timeZone)
	timeStamp := inLocTimestamp(nowInLoc)
	timeStampStr := strconv.FormatInt(timeStamp, 10)
	return timeStampStr
}

func testLocTimestamp() {
	timeZone, _ := time.LoadLocation("Asia/Shanghai")
	timeZoneJP, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().Round(0)
	nowInLoc := now.In(timeZone)
	nowInLocJP := now.In(timeZoneJP)
	fmt.Println(now, nowInLoc, nowInLocJP)
	fmt.Printf("time		| now: %v | inLoc: %v | inLocJP: %v\n", now, nowInLoc, nowInLocJP)
	fmt.Printf("unix		| now: %v | inLoc: %v | inLocJP: %v\n", now.Unix(), nowInLoc.Unix(), nowInLocJP.Unix())
	fmt.Printf("timestamp	| now: %v | inLoc: %v | inLocJP: %v\n", now.Unix(), inLocTimestamp(nowInLoc), inLocTimestamp(nowInLocJP))
	fmt.Printf("difference | inLoc: timestamp - unix = %v\n", inLocTimestamp(nowInLoc)-nowInLoc.Unix())
	fmt.Printf("difference | inLoc: timestamp - unix = %v\n", inLocTimestamp(nowInLocJP)-nowInLocJP.Unix())
}

func manifestSave (name string, url string, path string) {
	file := "./manifest/" + "aria2c.list"

	os.MkdirAll("./manifest/", os.ModePerm)
	// err := os.Remove(file)
	// if err != nil {
	// 	log.Println(err)
	// }
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(url + "\n  out=" + path + name + "\n"); err != nil {
		log.Println(err)
	}
}

func ManifestDL(path string) {
	manifestSave("master.json",
		"https://cdn-r18.necro-sm.dmmgames.com/files/manifest/webgl/r18/master.json?v=134292&2726217",
		"./webgl/" + path + "/")
	manifestSave("assetbundle.json",
		"https://cdn-r18.necro-sm.dmmgames.com/files/manifest/webgl/r18/assetbundle.json?v=134292&5955371",
		"./webgl/" + path + "/")
	manifestSave("advvoice.json",
		"https://cdn-r18.necro-sm.dmmgames.com/files/manifest/webgl/r18/advvoice.json?v=134292&2125604",
		"./webgl/" + path + "/")

	manifestSave("master.json",
		"https://cdn-r18.necro-sm.dmmgames.com/files/manifest/android/r18/master.json?v=134292&2726217",
		"./android/" + path + "/")
	manifestSave("assetbundle.json",
		"https://cdn-r18.necro-sm.dmmgames.com/files/manifest/android/r18/assetbundle.json?v=134292&5955371",
		"./android/" + path + "/")
	manifestSave("advvoice.json",
		"https://cdn-r18.necro-sm.dmmgames.com/files/manifest/android/r18/advvoice.json?v=134292&2125604",
		"./android/" + path + "/")
}

func DecryptString (message string, timeStamp string) (string, error) {
	data := []byte(APP_KEY + message + timeStamp)
	hasher := md5.New()
	hasher.Write(data)
	// str := strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
	// md5sum := md5.Sum(data)
	encoded := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(hasher.Sum(nil))
	// fmt.Printf("%x", md5.Sum(data))
	base64Replace := strings.NewReplacer(
		"+", "-",
		"=", "_",
		"/", "_",
	)
	s := base64Replace.Replace(encoded)
	// log.Println(string(encoded), s)
	return s, nil
}

func getAssetURLPath (base string, resource string) *url.URL {
	urlBase, err := url.Parse(base)
	if err != nil {
		log.Fatal(err)
	}
	urlResource, err := url.Parse(resource)
	if err != nil {
		log.Fatal(err)
	}
	return urlBase.ResolveReference(urlResource)
}

func getAssetURL (domain string, base string, resource string) string {
	urlDomain, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
	}
	urlPath := getAssetURLPath(base, resource)
	t := getTimeStampString()
	s, err := DecryptString(urlPath.String(), t)
	if err != nil {
		log.Fatal(err)
	}
	url := formatString(
		"%v?s=%v&t=%v",
		urlDomain.ResolveReference(urlPath).String(), s, t)
	return url
}

func assetURLHandler (domain string, name string, hash string) string {
	// testLocTimestamp()
	timeZone, _ := time.LoadLocation("Asia/Tokyo")

	AppKey := "B3u5C3Y6kqw5jfUwH4SurNcjXV44rHyr"

	URI := name

	now := time.Now().Round(0)
	nowInLoc := now.In(timeZone)
	timeStamp := inLocTimestamp(nowInLoc)
	timeStampStr := strconv.FormatInt(timeStamp, 10)

	data := []byte(AppKey + URI + timeStampStr)
	hasher := md5.New()
	hasher.Write(data)
	// str := strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
	// md5sum := md5.Sum(data)
	encoded := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(hasher.Sum(nil))
	// fmt.Printf("%x", md5.Sum(data))
	base64Replace := strings.NewReplacer(
		"+", "-",
		"=", "_",
		"/", "_",
	)
	s := base64Replace.Replace(encoded)
	// log.Println(string(encoded), s)
	// log.Printf("%v%v?s=%v&t=%v&h=%v", domain, URI, s, timeStampStr, hash)
	return fmt.Sprintf("%v%v?s=%v&t=%v&h=%v", domain, URI, s, timeStampStr, hash)
}

func test (ASSET_ROOT string, ASSET_URI string, assetName string, assetHash string) {
	assetURLHandler(ASSET_ROOT, ASSET_URI + assetName, assetHash)
}

func fixAssetName (assetName string, input string, output string) string {
	// "notinit/image_event_r18_tw"
	// "notinit/image_event_tw_r18"
	if strings.Contains(assetName, input) {
		// log.Printf("fix assetName [%v] from %v to %v", assetName, input, output)
		return strings.Replace(assetName, input, output, -1)
	}
	return assetName
}

func genAria2FromManifest(manifestPath string, platform string, timeStr string) {
	fileExt := filepath.Ext(manifestPath)
	fileName := strings.Replace(filepath.Base(manifestPath), fileExt, "", -1)

	outputDIR := fmt.Sprintf("./download/%v/%v/", platform, timeStr)
	outputAria2File := outputDIR + fileName + ".list"

	os.MkdirAll(outputDIR, os.ModePerm)
	err := os.Remove(outputAria2File)
	if err != nil {
		log.Println(err)
	}
	f, err := os.OpenFile(outputAria2File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	// assetNameFile := "./manifest/asset_name.txt"

	// assetNameFn, err := os.OpenFile(assetNameFile,
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer assetNameFn.Close()

	byteValue, _ := ioutil.ReadFile(manifestPath)

	var manifest Manifest
	json.Unmarshal([]byte(byteValue), &manifest)

	// var assetBundleList []AssetBundle

	count := len(manifest.Data)
	log.Printf("Data count: %v", count)

	for i := 0; i < len(manifest.Data); i++ {
		assetName := manifest.Data[i].Name
		// assetName := fixAssetName(manifest.Data[i].Name, "notinit/image_event_r18_tw", "notinit/image_event_tw_r18")
		assetO := manifest.Data[i].O
		// assetO := fixAssetName(manifest.Data[i].O, "notinit/image_event_r18_tw", "notinit/image_event_tw_r18")
		assetH := manifest.Data[i].H
		assetURL := assetURLHandler(ASSET_ROOT, ASSET_URI_MAP[platform] + assetName, assetH)

		// assetBundleList = append(assetBundleList, AssetBundle{
		// 	Key:		assetName,
		// 	FileName:	assetName,
		// 	Hash:		assetH,
		// 	Url:		assetURL,
		// 	AssetOutPut:	ASSET_URI + assetName,
		// })

		// if _, err := assetNameFn.WriteString(assetName + "\n"); err != nil {
		// 	log.Println(err)
		// }

		if _, err := f.WriteString(assetURL + "\n  out=." + ASSET_URI_MAP[platform] + assetName + "\n"); err != nil {
			log.Println(err)
		}
		if assetO != "" {
			// if _, err := assetNameFn.WriteString(assetO + "\n"); err != nil {
			// 	log.Println(err)
			// }

			assetURLO := assetURLHandler(ASSET_ROOT, ASSET_URI_MAP[platform] + assetO, assetH)
			if _, err := f.WriteString(assetURLO + "\n  out=./O" + ASSET_URI_MAP[platform] + assetO + "\n"); err != nil {
				log.Println(err)
			}
		}
	}
}

func formatString(format string, v ...any) string {
	var b strings.Builder
	_, err := fmt.Fprintf(&b, format, v...)
	if err != nil {
		log.Fatal(err)
	}
	return b.String()
}

func formatAria2(url string, out string) string {
	return formatString("%v\n  out=%v\n", url, out)
}

func genAria2FromAssetName(assetNamePath string, domain string, platform string, timeStr string) {
	fileName := strings.TrimSuffix(filepath.Base(assetNamePath), filepath.Ext(assetNamePath))

	outputDir := filepath.Join(DOWNLOAD_DIR, "tmp", platform, timeStr)
	outputAria2File := filepath.Join(outputDir, fileName + ARIA2_FILE_EXT)

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.Remove(outputAria2File)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	fo, err := os.OpenFile(outputAria2File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	const BufSize int = 4 * 1024 * 1024
	foBufferedWriter := bufio.NewWriterSize(fo, BufSize)

	fi, err := os.Open(assetNamePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	fiScanner := bufio.NewScanner(fi)
	const MaxScanTokenSize int = 4 * 1024
	buf := make([]byte, MaxScanTokenSize)
	fiScanner.Buffer(buf, MaxScanTokenSize)

	for fiScanner.Scan() {
		line := fiScanner.Text()
		url := getAssetURL(domain, ASSET_ROOT_PATH_PLATFORM_MAP[platform], line)
		out := getAssetURLPath(ASSET_ROOT_PATH_PLATFORM_MAP[platform], line).String()
		// var b bytes.Buffer
		var b strings.Builder
		b.WriteString(formatAria2(url, out))
		_, err = foBufferedWriter.WriteString(b.String())
		if err != nil {
			log.Fatal(err)
		}
	}
	foBufferedWriter.Flush()
	if err := fiScanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func _replacePlaceholder (inList []string, outList []string) []string {
	var tmpList []string
	if len(inList) == 0 {
		return outList
	}
	for _, s := range inList {
		if strings.Contains(s, "(.(dmmpf|store))?") {
			tmpList = append(tmpList,
				strings.Replace(s, "(.(dmmpf|store))?", "", 1),
				strings.Replace(s, "(.(dmmpf|store))?", ".dmmpf", 1),
				strings.Replace(s, "(.(dmmpf|store))?", ".store", 1))
		} else if strings.Contains(s, "(__(dmmpf|store))?") {
			tmpList = append(tmpList,
				strings.Replace(s, "(__(dmmpf|store))?", "", 1),
				strings.Replace(s, "(__(dmmpf|store))?", "__dmmpf", 1),
				strings.Replace(s, "(__(dmmpf|store))?", "__store", 1))
		} else if strings.Contains(s, "(.(en|ja|tw))?") {
			tmpList = append(tmpList,
				strings.Replace(s, "(.(en|ja|tw))?", "", 1),
				strings.Replace(s, "(.(en|ja|tw))?", "_en", 1),
				strings.Replace(s, "(.(en|ja|tw))?", "_ja", 1),
				strings.Replace(s, "(.(en|ja|tw))?", "_tw", 1))
		} else if strings.Contains(s, "(_(en|ja|tw))?") {
			tmpList = append(tmpList,
				strings.Replace(s, "(_(en|ja|tw))?", "", 1),
				strings.Replace(s, "(_(en|ja|tw))?", "_en", 1),
				strings.Replace(s, "(_(en|ja|tw))?", "_ja", 1),
				strings.Replace(s, "(_(en|ja|tw))?", "_tw", 1))
		} else if strings.Contains(s, "(_r18)?") {
			tmpList = append(tmpList,
				strings.Replace(s, "(_r18)?", "", 1),
				strings.Replace(s, "(_r18)?", "_r18", 1))
		} else if strings.Contains(s, "(_voice)?") {
			tmpList = append(tmpList,
				strings.Replace(s, "(_voice)?", "", 1),
				strings.Replace(s, "(_voice)?", "_voice", 1))
		} else if strings.Contains(s, "(_voice|_image)?") {
			tmpList = append(tmpList,
				strings.Replace(s, "(_voice|_image)?", "", 1),
				strings.Replace(s, "(_voice|_image)?", "_voice", 1),
				strings.Replace(s, "(_voice|_image)?", "_image", 1))
		} else if strings.Contains(s, " ?") {
			tmpList = append(tmpList,
				strings.Replace(s, " ?", "", 1),
				strings.Replace(s, " ?", " ", 1))
		} else {
			outList = append(outList, s)
		}
	}
	return _replacePlaceholder(tmpList, outList)
}

func replacePlaceholder (s string) []string {
	return _replacePlaceholder([]string{s}, []string{})
}

func genAssetName () {
	var assetNameList []string

	outputFile := MANIFEST_ASSET_NAME_LIST_FILE

	err = os.Remove(outputFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	fo, err := os.OpenFile(outputFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	const BufSize int = 4 * 1024 * 1024
	fobw := bufio.NewWriterSize(fo, BufSize)

	// ^adv/chara_icon_image/fc\d{5}[a-z]?\.dmm(\.(dmmpf|store))?
	// for i := 0; i <= 99999; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("adv/chara_icon_image/fc%05d.dmm(.(dmmpf|store))?", i))...)
	// 	for j := 'a'; j <= 'z'; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("adv/chara_icon_image/fc%05d%c.dmm(.(dmmpf|store))?", i, j))...)
	// 	}
	// }

	// ^image_banner(_(en|ja|tw))?/bnh(00|10|20|21|40|41|50|51|60)_0\d{4}\.dmm(\.(dmmpf|store))?
	// for _, i := range []string{"00", "10", "20", "21", "40", "41", "50", "51", "60"} {
	// 	for j := 0; j <= 9999; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("image_banner(_(en|ja|tw))?/bnh%s_%05d.dmm(.(dmmpf|store))?", i, j))...)
	// 	}
	// }
	// ^image_banner(_(en|ja|tw))?/bnh(00|10|20|21|40|41|50|51|60)_0[0-2]\d{4}\.dmm(\.(dmmpf|store))?
	// for _, i := range []string{"00", "10", "20", "21", "40", "41", "50", "51", "60"} {
	// 	for j := 0; j <= 20000; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("image_banner(_(en|ja|tw))?/bnh%s_%06d.dmm(.(dmmpf|store))?", i, j))...)
	// 	}
	// }

	// ^image_enemy/et\d{5}\.dmm(\.(dmmpf|store))?
	// for i := 0; i <= 99999; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_enemy/et%05d.dmm(.(dmmpf|store))?", i))...)
	// }

	// ^image_equipment_icon/et[01]\d{3}_[01]\d\.dmm
	// for i := 0; i <= 1500; i++ {
	// 	for j := 0; j <= 20; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("image_equipment_icon/et%03d_%02d.dmm", i, j))...)
	// 	}
	// }

	// ^image_equipment_small_icon/ets[01]\d{3}_[01]\d\.dmm
	// for i := 0; i <= 1500; i++ {
	// 	for j := 0; j <= 20; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("image_equipment_small_icon/ets%03d_%02d.dmm", i, j))...)
	// 	}
	// }

	// ^image_gacha(_(en|ja|tw))?/bns(20|21)_[09]\d{4}\.dmm(\.(dmmpf|store))?
	// for _, i := range []string{"20", "21"} {
	// 	for j := 0; j <= 9999; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("image_gacha(_(en|ja|tw))?/bns%s_%05d.dmm(.(dmmpf|store))?", i, j))...)
	// 	}
	// 	for j := 90000; j <= 99999; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("image_gacha(_(en|ja|tw))?/bns%s_%05d.dmm(.(dmmpf|store))?", i, j))...)
	// 	}
	// }

	// ^image_home_banner/home_bnh_dungeon(_(en|ja|tw))?\.dmm
	// assetNameList = append(assetNameList,
	// 	replacePlaceholder("image_home_banner/home_bnh_dungeon(_(en|ja|tw))?.dmm")...)

	// ^image_home_banner/home_bnh_pvp(_(en|ja|tw))?\.dmm
	// assetNameList = append(assetNameList,
	// 	replacePlaceholder("image_home_banner/home_bnh_pvp(_(en|ja|tw))?.dmm")...)

	// ^image_home_banner/quest_btn_limit_\d{1,3}\.dmm
	// for i := 0; i <= 400; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_home_banner/quest_btn_limit_%d.dmm", i))...)
	// }

	// ^image_login_bonus_item_icon/lit[01]\d{4}\.dmm(\.(en|ja|tw))?
	// for i := 0; i <= 15000; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_login_bonus_item_icon/lit%05d.dmm(.(en|ja|tw))?", i))...)
	// }

	// ^image_quest_place/qp000\d{2}\.dmm
	// for i := 0; i <= 200; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_quest_place/qp00%03d.dmm", i))...)
	// }

	// ^image_shop(_(en|ja|tw))?/sin(00|40|41)_0[0-1]\d{3}\.dmm(\.(dmmpf|store))?
	// for _, i := range []string{"00", "40", "41"} {
	// 	for j := 0; j <= 2000; j++ {
	// 		assetNameList = append(assetNameList,
	// 			replacePlaceholder(fmt.Sprintf("image_shop(_(en|ja|tw))?/sin%s_%05d.dmm(.(dmmpf|store))?", i, j))...)
	// 	}
	// }

	// // ^image_shop_icon/sit\d{5}\.dmm
	// for i := 0; i <= 99999; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_shop_icon/sit%05d.dmm", i))...)
	// }

	// // ^image_unit_harem(_r18)?/harem_00\d{3}\.dmm(\.(dmmpf|store))?
	// for i := 0; i <= 500; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_unit_harem(_r18)?/harem_00%03d.dmm(.(dmmpf|store))?", i))...)
	// }

	// // ^image_unit_icon/ut[01]0\d{3}\.dmm(\.(dmmpf|store))?
	// for i := 0; i <= 500; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_unit_icon/ut00%03d.dmm(.(dmmpf|store))?", i))...)
	// }
	// for i := 0; i <= 100; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_unit_icon/ut10%03d.dmm(.(dmmpf|store))?", i))...)
	// }

	// // ^image_unit_top/up[01]0\d{3}\.dmm(\.(dmmpf|store))?
	// for i := 0; i <= 500; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_unit_top/up00%03d.dmm(.(dmmpf|store))?", i))...)
	// }
	// for i := 0; i <= 100; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("image_unit_top/up10%03d.dmm(.(dmmpf|store))?", i))...)
	// }

	// // ^notinit/_adv/adv_\d{5}(_voice)?\.dmm(\.(en|ja|tw))?
	// for i := 0; i <= 99999; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("notinit/_adv/adv_%05d(_voice)?.dmm(.(en|ja|tw))?", i))...)
	// }
	// // ^notinit/_adv/adv_0101\d{4}(_voice)?\.dmm(\.(en|ja|tw))?
	// for i := 0; i <= 1111; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("notinit/_adv/adv_0101%04d(_voice)?.dmm(.(en|ja|tw))?", i))...)
	// }
	// // ^notinit/_adv/adv_0900\d{4}(_voice)?\.dmm(\.(en|ja|tw))?
	// for i := 0; i <= 100; i++ {
	// 	assetNameList = append(assetNameList,
	// 		replacePlaceholder(fmt.Sprintf("notinit/_adv/adv_0900%04d(_voice)?.dmm(.(en|ja|tw))?", i))...)
	// }

	// ^notinit/adv_texture/back_image/bg\d{3}_[1-4]\d\.dmm(\.(en|ja|tw))?
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 50; j++ {
			assetNameList = append(assetNameList,
				replacePlaceholder(fmt.Sprintf("notinit/adv_texture/back_image/bg%03d_%02d.dmm(.(en|ja|tw))?", i, j))...)
		}
	}

	// ^notinit/adv_texture/chara_top_image/ob[0-5]0\d{3}[a-z]?\.dmm
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 100; j++ {
			for k := 'a'; k <= 'z'; k++ {
				assetNameList = append(assetNameList,
					replacePlaceholder(fmt.Sprintf("notinit/adv_texture/chara_top_image/ob%d0%03d%c.dmm", i, j, k))...)
			}
		}
	}

	// ^notinit/adv_texture/chara_top_image/st[0-5]0\d{3}[a-z]?\.dmm
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 999; j++ {
			for k := 'a'; k <= 'z'; k++ {
				assetNameList = append(assetNameList,
					replacePlaceholder(fmt.Sprintf("notinit/adv_texture/chara_top_image/st%d0%03d%c.dmm", i, j, k))...)
			}
		}
	}

	// ^notinit/advchara(_r18)?/har_00\d{3}(_voice|_image)?\.dmm(\.(en|ja|tw))?
	for i := 0; i <= 500; i++ {
		assetNameList = append(assetNameList,
			replacePlaceholder(fmt.Sprintf("notinit/advchara(_r18)?/har_00%03d(_voice|_image)?.dmm(.(en|ja|tw))?", i))...)
	}

	// ^notinit/battle/event/advquest[0-4][0-5]\d{3}(_voice|_image)?\.dmm(\.(en|ja|tw))?
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 5; j++ {
			for k := 0; k <= 700; k++ {
				assetNameList = append(assetNameList,
					replacePlaceholder(fmt.Sprintf("notinit/battle/event/advquest%d%d%03d(_voice|_image)?.dmm(.(en|ja|tw))?", i, j, k))...)
			}
		}
	}

	for _, i := range assetNameList {
		var b strings.Builder
		b.WriteString(i)
		b.WriteString("\n")
		if _, err = fobw.WriteString(b.String()); err != nil {
			log.Fatal(err)
		}
	}
	fobw.Flush()
}

func init() {
}

func main() {
	config.InitializeConfig(*configPath)

	// timeStr := time.Now().Format(time.RFC3339)

	// ManifestDL(timeStr)

	// genAssetName()

	// genAria2FromAssetName(MANIFEST_ASSET_NAME_LIST_FILE, ASSET_ROOT, "webgl", timeStr)

	// fmt.Println(assetURLHandler("", "/secure/data/webgl/resources/adv/chara_icon_image/fc00000.dmm", ""))

	// genAria2FromManifest("./manifest/webgl/2022-06-30T15:55:12+08:00/master.json", "webgl", timeStr)
	// genAria2FromManifest("./manifest/webgl/2022-06-30T15:55:12+08:00/assetbundle.json", "webgl", timeStr)
	// genAria2FromManifest("./manifest/webgl/2022-06-30T15:55:12+08:00/advvoice.json", "webgl", timeStr)

	// genAria2FromManifest("./manifest/android/2022-06-30T16:18:35+08:00/master.json", "android", timeStr)
	// genAria2FromManifest("./manifest/android/2022-06-30T16:18:35+08:00/assetbundle.json", "android", timeStr)
	// genAria2FromManifest("./manifest/android/2022-06-30T16:18:35+08:00/advvoice.json", "android", timeStr)
}
