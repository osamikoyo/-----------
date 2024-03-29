package main

import(
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)
type An struct{
	Name string
	Symp string
}
type disease struct{
	Name string
    Id int
	Claim string
}
type str struct{
    sending string
}
func server (){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		strf := "Приветствуем на главной странице!"
		str := str{
			sending: strf,
		}
		maintmpl, _:= template.ParseFiles("./static/index.html")
		maintmpl.Execute(w,str)
	})
	http.HandleFunc("/clien", func(w http.ResponseWriter, r *http.Request) {
	strf := "Приветствуем на главной странице!"
	str := str{
		sending: strf,
	}
	maintmpl, _:= template.ParseFiles("./static/main.html")
	maintmpl.Execute(w,str)
})
http.HandleFunc("/claim", func(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	addDB(r.Form.Get("name"),r.Form.Get("claim"), rand.Intn(1000000))

	fmt.Println(r.Form.Get("name"),r.Form.Get("claim"), "djj")


	fromtmpl,_ := template.ParseFiles("./static/form.html")
	complaint := disease{
		Name: r.Form.Get("claim"),
		Claim: r.Form.Get("name"),
		Id: rand.Intn(1000000),
	}
	fromtmpl.Execute(w,"efd")
	fmt.Println(complaint)

})
http.HandleFunc("/docktor", func(w http.ResponseWriter, r *http.Request) {
	nm, symp := get()
	datAnimal :=  An{
		Name: nm,
		Symp: symp,
	}
	tmpls, _ := template.ParseFiles("./satic/dok.html")
	tmpls.Execute(w,datAnimal)
})
fmt.Println("Server is listening...")
http.ListenAndServe(":8181", nil)
}