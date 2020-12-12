package redisUtil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
	_ "strings"
	"unicode"
)

type DataObj struct {
	Key   string
	Value string
	List  []string
}

var ctx = context.Background()

func test() {

	//fmt.Println(buildDbStr("139.196.38.232:6379", "adminfeng@.", 0))
}
func GetRedisDb(address, password string, db int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	return rdb
}

func GetDbSize(rdb *redis.Client) int {

	str := rdb.ConfigGet(ctx, "databases").String()
	ls := strings.Fields(str)
	for s := range ls {
		sr := ls[s]
		if strings.Contains(sr, "]") {
			new := strings.ReplaceAll(sr, "]", "")
			num, _ := strconv.Atoi(new)
			return num
		}

	}
	return 0
}
func KeyList(rdb *redis.Client) []string {
	//str:= rdb.ConfigGet(ctx, "databases").String()
	//获取key的数量
	keysize := rdb.DBSize(ctx)
	//获取所有key的值，游标设置0
	val, _ := rdb.Scan(ctx, 0, "*", keysize.Val()).Val()
	return val
}

func GetDbList(rdb *redis.Client) int64 {
	keysize := rdb.DBSize(ctx)

	return keysize.Val()
}

func GetRedisValue(key string, rdb *redis.Client) string {
	//获取key对应值的的类型
	valuetype := rdb.Type(ctx, key)
	ts, _ := valuetype.Result()
	resultStr := "{"
	ss := DataObj{
		Key: key,
	}
	switch ts {

	case "list": //list类型
		valueLen := rdb.LLen(ctx, key).Val()
		res := rdb.LRange(ctx, key, 0, valueLen).Val()
		slice := make([]string, valueLen)
		var listStr = "["
		for _, i := range res {
			slice = append(slice, i)
			listStr += "\"" + i + "\","
		}
		ss.List = slice
		listStr = listStr[0 : len(listStr)-1]
		listStr += "],"
		resultStr += "\"" + key + "\"" + ":" + listStr
		break
	case "set": //set类型
		setLen := rdb.LLen(ctx, key).Val()
		setList := rdb.SMembers(ctx, key).Val()
		setSlice := make([]string, setLen)
		var str = "["
		for _, i := range setList {
			setSlice = append(setSlice, i)
			str += "\"" + i + "\"" + ","
		}
		str = str[0 : len(str)-1]
		str += "],"
		resultStr += "\"" + key + "\"" + ":" + str
		ss.List = setSlice
		break
	case "hash": //hash类型
		hashStr := ""
		hashLen := rdb.LLen(ctx, key).Val()
		hashKeys := rdb.HKeys(ctx, key).Val()
		hashSlice := make([]string, hashLen)
		for _, i := range hashKeys {
			//fmt.Println(i)
			hashValues := rdb.HGetAll(ctx, key).Val()
			hashStr += "\"" + i + "\":["
			for _, j := range hashValues {
				hashStr += "\"" + j + "\","
				//fmt.Println( j)
				hashSlice = append(hashSlice, i)
			}
			hashStr = hashStr[0 : len(hashStr)-1]
			hashStr += "],"
		}

		resultStr += hashStr
		ss.List = hashSlice
		break
	case "zset":
		zsetStr := "\"" + key + "\":["
		zsetlen := rdb.LLen(ctx, key).Val()
		zsetValue := rdb.ZRange(ctx, key, 0, zsetlen).Val()
		zsetSlice := make([]string, zsetlen)
		for i := range zsetValue {
			zsetStr += "\"" + zsetValue[i] + "\","
			zsetSlice = append(zsetSlice, zsetValue[i])
		}
		zsetStr = zsetStr[0 : len(zsetStr)-1]
		zsetStr += "],"
		resultStr += zsetStr
		ss.List = zsetSlice
		break
	default:
		value := rdb.Get(ctx, key).Val()
		s2 := []rune(value)
		resultStr += "\"" + key + "\"" + ":" + "\"" + string(s2) + "\","
		ss.Value = string(s2)
	}
	resultStr = resultStr[0:len(resultStr)-1] + "}"
	return resultStr
}
func chr(r rune) rune {
	if r >= 0 && r <= unicode.MaxASCII {
		return r
	}
	return unicode.ReplacementChar
}
func BuildDbStr(address, password string, db int) map[string]DataObj {

	mp := make(map[string]DataObj)
	//获取redis连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	rdb.ConfigGet(ctx, "databases").Val()
	//获取key的数量
	keysize := rdb.DBSize(ctx)
	//获取所有key的值，游标设置0
	val, _ := rdb.Scan(ctx, 0, "*", keysize.Val()).Val()
	var resultStr = "{"
	for i := 0; i < len(val); i++ {
		//获取key对应值的的类型
		valuetype := rdb.Type(ctx, val[i])
		ts, _ := valuetype.Result()
		key := val[i]

		ss := DataObj{
			Key: key,
		}
		switch ts {

		case "list": //list类型
			valueLen := rdb.LLen(ctx, key).Val()
			res := rdb.LRange(ctx, key, 0, valueLen).Val()
			slice := make([]string, valueLen)
			var listStr = "["
			for _, i := range res {
				slice = append(slice, i)
				listStr += "\"" + i + "\","
			}
			ss.List = slice
			listStr = listStr[0 : len(listStr)-1]
			listStr += "],"
			resultStr += "\"" + key + "\"" + ":" + listStr
			break
		case "set": //set类型
			setLen := rdb.LLen(ctx, key).Val()
			setList := rdb.SMembers(ctx, key).Val()
			setSlice := make([]string, setLen)
			var str = "["
			for _, i := range setList {
				setSlice = append(setSlice, i)
				str += "\"" + i + "\"" + ","
			}
			str = str[0 : len(str)-1]
			str += "],"
			resultStr += "\"" + key + "\"" + ":" + str
			ss.List = setSlice
			break
		case "hash": //hash类型
			hashStr := ""
			hashLen := rdb.LLen(ctx, key).Val()
			hashKeys := rdb.HKeys(ctx, key).Val()
			hashSlice := make([]string, hashLen)
			for _, i := range hashKeys {
				//fmt.Println(i)
				hashValues := rdb.HGetAll(ctx, key).Val()
				hashStr += "\"" + i + "\":["
				for _, j := range hashValues {
					hashStr += "\"" + j + "\","
					//fmt.Println( j)
					hashSlice = append(hashSlice, i)
				}
				hashStr = hashStr[0 : len(hashStr)-1]
				hashStr += "],"
			}

			resultStr += hashStr
			ss.List = hashSlice
			break
		case "zset":
			zsetStr := "\"" + key + "\":["
			zsetlen := rdb.LLen(ctx, key).Val()
			zsetValue := rdb.ZRange(ctx, key, 0, zsetlen).Val()
			zsetSlice := make([]string, zsetlen)
			for i := range zsetValue {
				zsetStr += "\"" + zsetValue[i] + "\","
				zsetSlice = append(zsetSlice, zsetValue[i])
			}
			zsetStr = zsetStr[0 : len(zsetStr)-1]
			zsetStr += "],"
			resultStr += zsetStr
			ss.List = zsetSlice
			break
		default:

			value := rdb.Get(ctx, key).Val()
			resultStr += "\"" + key + "\"" + ":" + "\"" + value + "\","
			ss.Value = value
		}
		mp[key] = ss
	}

	resultStr = resultStr[0:len(resultStr)-1] + "}"
	return mp
}
