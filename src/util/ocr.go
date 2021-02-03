package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ocrData struct {
	msg     string
	results string
}
type ocr struct {
	text       string
	confidence float64
}

func main() {

	//str := "{\"msg\":\"\",\"results\":[[{\"confidence\":0.9604636430740356,\"text\":\"口\",\"text_region\":[[1126,2],[1142,2],[1142,17],[1126,17]]},{\"confidence\":0.9423713088035583,\"text\":\"丰硕\",\"text_region\":[[26,50],[51,50],[51,62],[26,62]]},{\"confidence\":0.9944733381271362,\"text\":\"郭莉\",\"text_region\":[[472,42],[513,42],[513,68],[472,68]]},{\"confidence\":0.9666762948036194,\"text\":\"Q搜索\",\"text_region\":[[104,56],[182,56],[182,86],[104,86]]},{\"confidence\":0.9952853918075562,\"text\":\"郭莉\",\"text_region\":[[408,54],[449,54],[449,79],[408,79]]},{\"confidence\":0.9088772535324097,\"text\":\"+\",\"text_region\":[[325,63],[345,63],[345,82],[325,82]]},{\"confidence\":0.9869804382324219,\"text\":\"郭莉(上海汇航捷讯网络科技有限公司-集团-产品技术中心-技术部-测试-测试开发|资深测试工程师）\",\"text_region\":[[472,70],[1131,70],[1131,92],[472,92]]},{\"confidence\":0.6171554327011108,\"text\":\"g\",\"text_region\":[[1151,112],[1182,112],[1182,148],[1151,148]]},{\"confidence\":0.999616265296936,\"text\":\"消息\",\"text_region\":[[17,133],[57,133],[57,158],[17,158]]},{\"confidence\":0.9768012762069702,\"text\":\"@\",\"text_region\":[[99,127],[120,127],[120,145],[99,145]]},{\"confidence\":0.997240424156189,\"text\":\"39016发一下\",\"text_region\":[[911,163],[1028,163],[1028,186],[911,186]]},{\"confidence\":0.9972343444824219,\"text\":\"丰硕\",\"text_region\":[[1060,162],[1099,162],[1099,186],[1060,186]]},{\"confidence\":0.9717870950698853,\"text\":\"测试环境发.内部21:10\",\"text_region\":[[164,172],[357,170],[357,193],[164,195]]},{\"confidence\":0.5390841960906982,\"text\":\"D\",\"text_region\":[[1152,166],[1181,166],[1181,198],[1152,198]]},{\"confidence\":0.5118012428283691,\"text\":\"1c\",\"text_region\":[[23,179],[51,179],[51,203],[23,203]]},{\"confidence\":0.9341087937355042,\"text\":\"[539条]测试环境发布...\",\"text_region\":[[164,202],[336,202],[336,224],[164,224]]},{\"confidence\":0.9990017414093018,\"text\":\"文档\",\"text_region\":[[20,209],[57,209],[57,233],[20,233]]},{\"confidence\":0.9987426996231079,\"text\":\"已读\",\"text_region\":[[1015,202],[1045,202],[1045,219],[1015,219]]},{\"confidence\":0.5067358613014221,\"text\":\"口\",\"text_region\":[[1152,220],[1181,220],[1181,252],[1152,252]]},{\"confidence\":0.9986925721168518,\"text\":\"8\",\"text_region\":[[23,250],[52,250],[52,283],[23,283]]},{\"confidence\":0.9994904398918152,\"text\":\"生产发布群内部\",\"text_region\":[[165,253],[293,253],[293,275],[165,275]]},{\"confidence\":0.9994367361068726,\"text\":\"21:10\",\"text_region\":[[311,253],[356,253],[356,272],[311,272]]},{\"confidence\":0.9998178482055664,\"text\":\"30分钟前\",\"text_region\":[[716,253],[781,253],[781,272],[716,272]]},{\"confidence\":0.9998607635498047,\"text\":\"工作\",\"text_region\":[[18,281],[57,281],[57,311],[18,311]]},{\"confidence\":0.9645020961761475,\"text\":\"[44条]机器人：pro环...\",\"text_region\":[[165,283],[348,283],[348,305],[165,305]]},{\"confidence\":0.9128062129020691,\"text\":\"@\",\"text_region\":[[1151,290],[1179,290],[1179,323],[1151,323]]},{\"confidence\":0.9045066833496094,\"text\":\"告警－营销.：内部21:08\",\"text_region\":[[165,333],[357,330],[357,353],[165,356]]},{\"confidence\":0.9963915348052979,\"text\":\"郭莉\",\"text_region\":[[399,337],[435,337],[435,358],[399,358]]},{\"confidence\":0.9107789993286133,\"text\":\"[\\\"header”:(\\\"xSourceAppld\\\":\\\"60040\\\"\\\",\\\"model\\\":\",\"text_region\":[[468,336],[864,335],[864,360],[468,362]]},{\"confidence\":0.9994075894355774,\"text\":\"通讯录\",\"text_region\":[[12,359],[62,359],[62,383],[12,383]]},{\"confidence\":0.9435860514640808,\"text\":\"[18条双色球：ops错..\",\"text_region\":[[165,363],[348,363],[348,385],[165,385]]},{\"confidence\":0.9431594014167786,\"text\":\"{\\\"customer_id\\\":null,\\\"customername\\\":\\\"YQN测试\",\"text_region\":[[469,364],[888,363],[889,385],[470,386]]},{\"confidence\":0.9207372665405273,\"text\":\"11021121052\\\"\\\"customer types”:[2],\\\"contact data\\\":\",\"text_region\":[[468,387],[923,387],[923,416],[468,416]]},{\"confidence\":0.9960079193115234,\"text\":\"运去哪全员\",\"text_region\":[[164,412],[258,412],[258,435],[164,435]]},{\"confidence\":0.9989280700683594,\"text\":\"20:46\",\"text_region\":[[309,411],[359,411],[359,434],[309,434]]},{\"confidence\":0.9219264984130859,\"text\":\"{\\\"user_name\\\":\\\"用户\",\"text_region\":[[468,414],[635,414],[635,439],[468,439]]},{\"confidence\":0.9990653395652771,\"text\":\"周苏军Jack在益起动\",\"text_region\":[[163,440],[318,441],[318,464],[163,463]]},{\"confidence\":0.9957832098007202,\"text\":\"洁\",\"text_region\":[[110,441],[142,441],[142,454],[110,454]]},{\"confidence\":0.9178515076637268,\"text\":\"YQN_GUwQvb\\\",\\\"position\\\":null,\\\"urcllpone\\\":\\\"\",\"text_region\":[[468,441],[980,441],[980,467],[468,467]]},{\"confidence\":0.9915859699249268,\"text\":\"回\",\"text_region\":[[22,491],[54,491],[54,522],[22,522]]},{\"confidence\":0.9943474531173706,\"text\":\"郭莉\",\"text_region\":[[164,491],[206,491],[206,516],[164,516]]},{\"confidence\":0.9986057281494141,\"text\":\"20:40\",\"text_region\":[[309,491],[360,491],[360,514],[309,514]]},{\"confidence\":0.9241546392440796,\"text\":\"ull，\\\"shiment list\\\":f \\\"cycle time\\\":2,\\\"line id list\\\":\",\"text_region\":[[467,492],[887,491],[887,516],[467,518]]},{\"confidence\":0.9937686920166016,\"text\":\"郭莉\",\"text_region\":[[104,504],[145,504],[145,530],[104,530]]},{\"confidence\":0.9992061853408813,\"text\":\"货主有问题\",\"text_region\":[[164,521],[244,521],[244,544],[164,544]]},{\"confidence\":0.9572484493255615,\"text\":\"CIF\",\"text_region\":[[387,525],[414,525],[414,547],[387,547]]},{\"confidence\":0.9985385537147522,\"text\":\"告警-销售内部\",\"text_region\":[[165,573],[282,573],[282,595],[165,595]]},{\"confidence\":0.9997743368148804,\"text\":\"19:33\",\"text_region\":[[312,572],[357,572],[357,591],[312,591]]},{\"confidence\":0.9717372059822083,\"text\":\"[9条]春香：Slowlog....\",\"text_region\":[[163,601],[348,602],[347,625],[163,624]]},{\"confidence\":0.8267797827720642,\"text\":\"e\",\"text_region\":[[21,622],[52,622],[52,658],[21,658]]},{\"confidence\":0.9992581605911255,\"text\":\"邮箱\",\"text_region\":[[164,651],[205,651],[205,676],[164,676]]},{\"confidence\":0.9994665384292603,\"text\":\"19:26\",\"text_region\":[[311,649],[359,652],[357,677],[309,673]]},{\"confidence\":0.5195780992507935,\"text\":\"口\",\"text_region\":[[25,669],[52,669],[52,700],[25,700]]},{\"confidence\":0.989794909954071,\"text\":\"YWork@yunquna.co...\",\"text_region\":[[163,679],[327,683],[326,706],[162,702]]},{\"confidence\":0.6335545778274536,\"text\":\"心\",\"text_region\":[[23,712],[51,712],[51,745],[23,745]]},{\"confidence\":0.9043053984642029,\"text\":\"李通疯狂打码：：19:14\",\"text_region\":[[164,733],[357,730],[358,753],[164,756]]},{\"confidence\":0.9987732172012329,\"text\":\"李通\",\"text_region\":[[104,743],[145,743],[145,772],[104,772]]},{\"confidence\":0.99908047914505,\"text\":\"Enter发送，Ctrl+Enter换行\",\"text_region\":[[816,744],[1012,744],[1012,766],[816,766]]},{\"confidence\":0.999873697757721,\"text\":\"发送\",\"text_region\":[[1051,741],[1094,741],[1094,767],[1051,767]]}]],\"status\":\"0\"}\n"
	//rs := analysisResult(str)
	//pringResult(rs)
	//fmt.Print(dat)
	fmt.Println("命令行的参数有", len(os.Args))
	if len(os.Args) < 2 {
		return
	}
	for i, v := range os.Args {
		//fmt.Printf("args[%v]=%v\n", i, v)
		b, err := PathExists(v)
		if err != nil {
			fmt.Printf("PathExists(%s),err(%v)\n", v, err)
		}
		if b {
			if strings.Contains(v, "exe") {
				continue
			}
			size := imgsize(v)
			fmt.Printf("path %s 存在", v)
			fmt.Printf("文件大小 %d kb \n", size/1024)
			bts := ImagesToBase64(v, size)
			data := "{\"images\":[\"" + bts + "\"]}"
			//fmt.Printf(data)
			//ioutil.WriteFile("D://a.png.txt", []byte(data), 0667)
			fmt.Printf("正在识别第 %d 个", i)
			result := sendPost("http://139.196.38.232:18866/predict/ocr_system", data)
			list := analysisResult(result)
			fmt.Printf("识别第%d个图片结束,返回结果如下:\n", i)
			pringResult(list)
		}
	}
}
func pringResult(list [][]map[string]interface{}) {
	for _, r := range list {
		for _, v := range r {
			if nil != v {
				text := v["text"].(string)
				confidence := v["confidence"].(float64)
				if text != "" {
					fmt.Print(text + "               准确率:")
					fmt.Println(confidence)
				}
			}

		}
	}
}

func analysisResult(bs string) [][]map[string]interface{} {

	var dat map[string][][]map[string]string
	json.Unmarshal([]byte(bs), &dat)
	var data map[string][][]map[string]interface{}
	json.Unmarshal([]byte(bs), &data)
	//a := dat2["results"][0][0]["text"]
	//s := a.(string)
	//println(s)
	results := data["results"]
	var result = make([][]map[string]interface{}, len(results))
	for _, v := range data {
		vs := make([]map[string]interface{}, len(v))
		for _, s := range v {
			for _, j := range s {
				vs = append(vs, j)
			}
		}
		result = append(result, vs)
	}
	return result
}

func imgsize(str_images string) int64 {
	file, _ := os.Open(str_images)
	buf := make([]byte, 2014)
	sum := 0
	for {
		n, err := file.Read(buf)
		sum += n
		if err == io.EOF {
			break
		}
	}
	return int64(sum)
}
func sendPost(url string, post string) string {

	var jsonStr = []byte(post)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var ocrDtr ocrData
	json.Unmarshal(body, &ocrDtr)
	texts := ocrDtr.results
	for i, v := range texts {
		fmt.Print(i)
		fmt.Print(v)
	}
	result := string(body)
	//fmt.Println("response Body:",result)
	return result
}

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func ImagesToBase64(str_images string, size int64) string {
	ff, _ := os.Open(str_images)
	defer ff.Close()
	sourcebuffer := make([]byte, size)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	return sourcestring
}
