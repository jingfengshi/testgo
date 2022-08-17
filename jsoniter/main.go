package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

func main() {
	//extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	//output, _ := jsoniter.Marshal(struct {
	//	UserName      string
	//	FirstLanguage string
	//}{
	//	UserName:      "taowen",
	//	FirstLanguage: "Chinese",
	//})
	//fmt.Println(string(output))
	//type User struct {
	//	Email    string `json:"email"`
	//	Password string `json:"password"`
	//	// many more fields…
	//}
	//user :=User{
	//	Email: "240023810@qq.com",
	//	Password: "123",
	//}
	//userB,_:=json.Marshal(struct {
	//	*User
	//	Password bool `json:"password,omitempty"`
	//}{
	//	User: &user,
	//	Password: true,
	//})
	//fmt.Println(string(userB))

	//ab,_:=json.Marshal(struct {
	//	*User
	//	Token    string `json:"token"`
	//	Password bool `json:"password,omitempty"`
	//}{
	//	User: &user,
	//	Token: "aaaaa",
	//})

	//type BlogPost struct {
	//	URL   string `json:"url"`
	//	Title string `json:"title"`
	//}
	//
	//type Analytics struct {
	//	Visitors  int `json:"visitors"`
	//	PageViews int `json:"page_views"`
	//}
	//post :=new(BlogPost)
	////post.Title="aaa"
	////post.URL="bnnnn"
	//analytics :=new(Analytics)
	//analytics.Visitors=1
	//analytics.PageViews=2
	//ab,_:=json.Marshal(struct{
	//	*BlogPost
	//	*Analytics
	//}{post, analytics})
	//
	//fmt.Println(string(ab))

//	json.Unmarshal([]byte(`{
//  "url": "attila@attilaolah.eu",
//  "title": "Attila's Blog",
//  "visitors": 6,
//  "page_views": 14
//}`), &struct {
//		*BlogPost
//		*Analytics
//	}{post, analytics})
//
//	fmt.Printf("%+v",post,analytics)


	//type CacheItem struct {
	//	Key    string `json:"key"`
	//	MaxAge int    `json:"cacheAge"`
	//	Value  Value  `json:"cacheValue"`
	//}
	//
	//json.Marshal(struct{
	//	*CacheItem
	//
	//	// Omit bad keys
	//	OmitMaxAge omit `json:"cacheAge,omitempty"`
	//	OmitValue  omit `json:"cacheValue,omitempty"`
	//
	//	// Add nice keys
	//	MaxAge int    `json:"max_age"`
	//	Value  *Value `json:"value"`
	//}{
	//	CacheItem: item,
	//
	//	// Set the int by value:
	//	MaxAge: item.MaxAge,
	//
	//	// Set the nested struct by reference, avoid making a copy:
	//	Value: &item.Value,
	//})
	//extra.RegisterFuzzyDecoders()
	//var val map[string]interface{}
	//jsoniter.UnmarshalFromString(`[]`, &val)
	//fmt.Println(val)
	//val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	//a:=jsoniter.Get(val, "Colors", 0).ToString()
	//fmt.Println(a)
	//
	//m := map[string]interface{}{
	//	"3": 3,
	//	"1": 1,
	//	"2": 2,
	//}
	//json := jsoniter.ConfigCompatibleWithStandardLibrary
	//b, _ := json.Marshal(m)
	//fmt.Println(string(b))
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err = json_iterator.Marshal(group)
	os.Stdout.Write(b)

}
