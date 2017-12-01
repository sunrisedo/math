package controllers

import (
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

//
func GetRandCode(codetype, strlen int) string {
	rcode := ""
	mcode := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	scode := [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if codetype == 1 {
		for i := 0; i < strlen; i++ {
			rcode = rcode + mcode[r.Intn(10)]
		}
	} else {
		for i := 0; i < strlen; i++ {
			rcode = rcode + scode[r.Intn(36)]
		}
	}
	return rcode
}

///
func GetRandNum(min, max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min) + min
}

func GetRandNew(max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max)
}
func GetRandFloatTest() (ret float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		ret = r.Float64()
		if ret >= 0.1 {
			break
		}
	}
	return Round(ret/10, 2)
	// return Round(float64(r.Intn(100)/100), 3)
	// return float64(r.Intn(100)) / 1000
}

//wdasdas dasdda sdaaaaaaaaaaaaaaaaa
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

func GetRandFloat(min, max float64) float64 {
	nmin := int(min * 1000000)
	nmax := int(max * 1000000)
	return float64(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(nmax-nmin)+nmin) / 1000000
}

func GetProbability(probability int) bool {
	if rand.New(rand.NewSource(time.Now().UnixNano())).Intn(11) <= probability {
		return true
	}
	return false
}

func CreateFile(fileName, data string) {
	if checkFileIsExist(fileName) {
		if f, err := os.OpenFile(fileName, os.O_APPEND, 0666); err != nil {
			log.Println("open file error:", err)
		} else if _, err := io.WriteString(f, data); err != nil {
			log.Println("write file error:", err)
		}
	} else if err := ioutil.WriteFile(fileName, []byte(data), 0666); err != nil {
		log.Println("create file error:", err)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CreateDir(dirName string) {
	if checkFileIsExist(dirName) {
		return
	}
	var path string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	if dir, err := os.Getwd(); err != nil {
		log.Println("get now dir error:", err)
	} else if err := os.Mkdir(dir+path+dirName, os.ModePerm); err != nil {
		log.Println("create dir error:", err)
	}
}

// func (m *MarkClient) GetRandFloat(min, max int) int {
//  	return rand.New(mrand.NewSource(time.Now().UnixNano())).Intn(max - min) + min
// }
// var new [][]string
// for i:=0; i<value.len; i++ {
// 	contents := strings.Split(value,",")
// 	new[i] = append(new[i],contents)
// }
// log.Println(new)
