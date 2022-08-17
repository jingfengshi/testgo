package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main(){
	//var json="{\"programmers\":[{\"firstName\":\"Janet\",\"lastName\":\"McLaughlin\"},{\"firstName\":\"Elliotte\",\"lastName\":\"Hunter\"},{\"firstName:\"Jason\",\"lastName\":\"Harold\"}]}"
//	var json=`{
//  "programmers": [
//    {
//      "firstName": "Janet",
//      "lastName": "McLaughlin",
//    }, {
//      "firstName": "Elliotte",
//      "lastName": "Hunter",
//    }, {
//      "firstName": "Jason",
//      "lastName": "Harold",
//    }
//  ]
//}`
	var json="111dasda"
	//result := gjson.Get(json, "programmers.#.lastName")
	//for _, name := range result.Array() {
	//	println(name.String())
	//	println(name.Raw)
	//}

	//name := gjson.Get(json, `programmers.#(lastName="Hunter").firstName`)
	//println(name.String())  // prints "Elliotte"

	//result := gjson.Get(json, "programmers")
	//result.ForEach(func(key, value gjson.Result) bool {
	//	println(value.String())
	//	return true // keep iterating
	//})
	//gjson.Parse(json).Get("programmers").ForEach(func(key, value gjson.Result) bool {
	//	fmt.Println(value.String())
	//	return true
	//})
	//fmt.Println(result.Raw)
	//m,ok:=gjson.Parse(json).Get("programmers").Value().([]interface{})
	//if !ok{
	//	fmt.Println(ok)
	//}else{
	//	fmt.Printf("%+v",m)
	//}

	//if gjson.Get(json,"programmers1").Exists() {
	//	println("has")
	//}
	//if !gjson.Valid(json) {
	//	 fmt.Println("invalid json")
	//	 return
	//}
	//value := gjson.Get(json, "programmers")
	//fmt.Printf("%+v",value)
	json =`{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

	res:=gjson.Get(json,"children.2")
	fmt.Println(res)
}