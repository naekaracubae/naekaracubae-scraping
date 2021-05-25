package etc

import (
	"encoding/json"
	_struct2 "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"strconv"
	"strings"
)

//TODO : 이런 식이면 모든 구조체에 대해 다 만들어줘야 한다. 공통 인터페이스를 만들어서 통합하기!
func StructToStr(kakaoJobs *[]_struct2.Kakao) string {
	var contents strings.Builder
	for idx, kakaoJob := range *kakaoJobs {
		jsonBytes, err := json.Marshal(kakaoJob)
		CheckErr(err)
		jsonString := string(jsonBytes)
		idxString := strconv.Itoa(idx) + ". " + jsonString
		contents.WriteString(idxString)
		contents.WriteString("</br>")
	}

	return contents.String()
}
