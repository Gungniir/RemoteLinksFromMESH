package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

type lesson struct {
	Id          int    `json:"id"`
	SubjectName string `json:"subject_name"`
	Date        [3]int `json:"date"`
	Time        [3]int `json:"time"`
	LessonType  string `json:"lesson_type"`
	Duration    int    `json:"duration"`
}

/*
document.cookie.split("; ").filter(a => {
  let key = a.split("=")[0]
    return ["auth_token", "profile_id"].indexOf(key) !== -1
}).forEach(item => {
  let key = item.split("=")[0]
    let value = item.split("=")[1]

    console.log(key, value)
})
*/
const (
	authToken = ""
	profileId = 0
	date      = "2020-10-23"
)

func (l *lesson) isRemote() bool {
	return l.LessonType == "REMOTE"
}

func (l *lesson) getRemoteUrl() string {
	return "https://dnevnik.mos.ru/conference/?scheduled_lesson_id=" + strconv.Itoa(l.Id)
}

func (l *lesson) startTimeString() string {
	return fmt.Sprintf("%s:%s",
		func(a int) string {
			if a > 9 {
				return strconv.Itoa(a)
			}
			return "0" + strconv.Itoa(a)
		}(l.Time[0]),
		func(a int) string {
			if a > 9 {
				return strconv.Itoa(a)
			}
			return "0" + strconv.Itoa(a)
		}(l.Time[1]),
	)
}

func (l *lesson) endTimeString() string {
	return fmt.Sprintf("%s:%s",
		func(a int) string {
			if l.Time[1]+l.Duration > 60 {
				a++
			}

			if a > 9 {
				return strconv.Itoa(a)
			}
			return "0" + strconv.Itoa(a)
		}(l.Time[0]),
		func(a int) string {
			a += l.Duration
			if a > 60 {
				a %= 60
			}

			if a > 9 {
				return strconv.Itoa(a)
			}
			return "0" + strconv.Itoa(a)
		}(l.Time[1]),
	)
}

func (l *lesson) secondsFromStart() int {
	return l.Time[0]*60*60 + l.Time[1]*60 + l.Time[2]
}

func main() {
	rawJson := getJson()

	var lessons []lesson

	err := json.Unmarshal(rawJson, &lessons)
	if err != nil {
		log.Print(string(rawJson))
		log.Fatal("Failed to parse json")
	}

	sort.SliceStable(lessons, func(i, j int) bool {
		return lessons[i].secondsFromStart() < lessons[j].secondsFromStart()
	})

	for i, l := range lessons {
		fmt.Printf("%d. %s (%s-%s): %s\n", i+1, l.SubjectName, l.startTimeString(), l.endTimeString(), l.getRemoteUrl())
	}
}

func getJson() []byte {
	target, err := url.Parse("https://dnevnik.mos.ru/jersey/api/schedule_items?with_group_class_subject_info=true&with_rooms_info=true&with_course_calendar_info=true&with_lesson_info=true&from=" + date + "&to=" + date + "&group_id=5771444,5915928,5771441,5771467,5771449,5771452,5771450,5771453,5771445,5771451,5771442,5771448,5771447,5771446,5771443&student_profile_id=17568784")
	if err != nil {
		log.Fatal(err)
	}
	request := http.Request{
		URL:    target,
		Header: http.Header{},
	}
	request.AddCookie(&http.Cookie{
		Name:  "profile_id",
		Value: strconv.Itoa(profileId),
	})
	request.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: authToken,
	})

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	response, err := client.Do(&request)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	return body
}
