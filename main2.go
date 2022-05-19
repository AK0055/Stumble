package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// declaring a struct
type User struct {
	Id       int
	Name     string
	Gender   string
	Email    string
	Location int
	Like     string
}

// main function
func main() {
	fmt.Println("STUMBLE")
	// defining a struct instance
	var User []User

	// JSON array to be decoded
	// to an array in golang
	Data := []byte(`
    [{"id":1,"name":"Auguste","location":-19,"gender":"Female"},
	{"id":2,"name":"Ned","location":-25,"gender":"Male"},
	{"id":3,"name":"Saree","location":11,"gender":"Female"},
	{"id":4,"name":"Katheryn","location":-5,"gender":"Female"},
	{"id":5,"name":"Brice","location":-14},
	{"id":6,"name":"Kristofer","location":-19,"gender":"Male"},
	{"id":7,"name":"Crawford","location":-15,"email":"ccrosbie6@psu.edu"},
	{"id":8,"name":"Parker","location":-7,"gender":"Male","email":"plisciandri7@unesco.org"},
	{"id":9,"name":"Bevin","location":28,"email":"bgalier8@google.ru"},
	{"id":10,"name":"Christoffer","location":-25,"gender":"Male"},
	{"id":11,"name":"Marje","location":-9,"gender":"Female","email":"mfoldera@bbc.co.uk"},
	{"id":12,"name":"Lorine","location":7,"gender":"Female","email":"lthomtsonb@studiopress.com"},
	{"id":13,"name":"Monica","location":-22,"gender":"Female","email":"mbarabischc@exblog.jp"},
	{"id":14,"name":"Kattie","location":0,"gender":"Female","email":"kshoppeed@uiuc.edu"},
	{"id":15,"name":"Lindy","location":10,"gender":"Female","email":"lfrymane@elegantthemes.com"},
	{"id":16,"name":"Amity","location":-1,"gender":"Female"},
	{"id":17,"name":"Dorie","location":-9,"gender":"Male"},
	{"id":18,"name":"Shelby","location":-7,"gender":"Male"},
	{"id":19,"name":"Mame","location":-18,"gender":"Female","email":"mcommingsi@wikispaces.com"},
	{"id":20,"name":"Arv","location":-4,"gender":"Male"},
	{"id":21,"name":"Berkley","location":-10,"gender":"Male"},
	{"id":22,"name":"Huntley","location":-5,"gender":"Male"},
	{"id":23,"name":"Iago","location":-1,"gender":"Non-binary","email":"igrishankovm@amazonaws.com"},
	{"id":24,"name":"Ollie","location":-25,"gender":"Female","email":"odeverelln@artisteer.com"},
	{"id":25,"name":"Aliza","location":24},
	{"id":26,"name":"Natassia","location":-19,"gender":"Female"},
	{"id":27,"name":"Harv","location":20},
	{"id":28,"name":"Kimberli","location":15,"gender":"Female"},
	{"id":29,"name":"Amaleta","location":28,"gender":"Female","email":"abambridges@google.ru"},
	{"id":30,"name":"Marita","location":12,"gender":"Female"},
	{"id":31,"name":"Gilbertina","location":-22},
	{"id":32,"name":"Claudian","location":9,"gender":"Male","email":"chammerbergv@japanpost.jp"},
	{"id":33,"name":"Anya","location":6,"email":"akelletw@rediff.com"},
	{"id":34,"name":"Mildrid","location":-5,"gender":"Female","email":"mpersehousex@symantec.com"},
	{"id":35,"name":"Tomasina","location":18,"gender":"Polygender","email":"twelshy@cocolog-nifty.com"},
	{"id":36,"name":"Augustus","location":-30,"gender":"Male"},
	{"id":37,"name":"Anya","location":5,"gender":"Genderfluid","email":"abath10@eventbrite.com"},
	{"id":38,"name":"Wilfrid","location":21,"gender":"Male","email":"wgoodwill11@twitpic.com"},
	{"id":39,"name":"Tess","location":-28,"gender":"Female","email":"tjambrozek12@businessinsider.com"},
	{"id":40,"name":"Dickie","location":-29,"gender":"Male","email":"dcorbet13@etsy.com"},
	{"id":41,"name":"Cirilo","location":26},
	{"id":42,"name":"Gilli","location":21,"gender":"Female"},
	{"id":43,"name":"Franz","location":-25,"gender":"Male"},
	{"id":44,"name":"Eddie","location":18,"gender":"Male"},
	{"id":45,"name":"Rex","location":-20,"gender":"Agender","email":"rmytton18@liveinternet.ru"},
	{"id":46,"name":"Yank","location":-5,"gender":"Male"},
	{"id":47,"name":"Catrina","location":-30,"email":"cpiner1a@networksolutions.com"},
	{"id":48,"name":"Luca","location":2,"gender":"Agender","email":"lmarland1b@smugmug.com"},
	{"id":49,"name":"Natka","location":-15,"email":"nantonetti1c@reverbnation.com"},
	{"id":50,"name":"Normy","location":-5,"gender":"Male","email":"ntruswell1d@google.ca"},
	{"id":51,"name":"Cookie","location":3,"gender":"Female"},
	{"id":52,"name":"Otha","location":-13,"gender":"Female","email":"odance1f@ning.com"},
	{"id":53,"name":"Gina","location":-2,"gender":"Female"},
	{"id":54,"name":"Denise","location":-1,"gender":"Female"},
	{"id":55,"name":"Karney","location":-2},
	{"id":56,"name":"Barty","location":4,"gender":"Agender","email":"bbenbow1j@behance.net"},
	{"id":57,"name":"Pietro","location":11,"gender":"Male","email":"pdesouza1k@rambler.ru"},
	{"id":58,"name":"Bogey","location":-3},
	{"id":59,"name":"Aldis","location":-3,"gender":"Male"},
	{"id":60,"name":"Lenora","location":26,"gender":"Female","email":"lsautter1n@creativecommons.org"},
	{"id":61,"name":"Ainslee","location":-16,"gender":"Female"},
	{"id":62,"name":"Sigvard","location":-4,"gender":"Male"},
	{"id":63,"name":"Nicholle","location":-3,"gender":"Female"},
	{"id":64,"name":"Kelly","location":-19,"gender":"Male","email":"kpretsell1r@jimdo.com"},
	{"id":65,"name":"Daniella","location":18,"gender":"Female"},
	{"id":66,"name":"Lorne","location":10,"gender":"Male","email":"lborrett1t@tuttocitta.it"},
	{"id":67,"name":"Lona","location":-7,"gender":"Female","email":"lgreetham1u@hostgator.com"},
	{"id":68,"name":"Selia","location":-12,"gender":"Female","email":"scradick1v@ebay.com"},
	{"id":69,"name":"Danella","location":25,"gender":"Female"},
	{"id":70,"name":"Sanderson","location":-4,"email":"smourant1x@csmonitor.com"},
	{"id":71,"name":"Idaline","location":-10,"email":"ishepherdson1y@sitemeter.com"},
	{"id":72,"name":"Jemmy","location":-28,"gender":"Female"},
	{"id":73,"name":"Max","location":-3,"email":"mholsey20@disqus.com"},
	{"id":74,"name":"Paulina","location":-15,"gender":"Non-binary"},
	{"id":75,"name":"Ilka","location":23,"gender":"Female"},
	{"id":76,"name":"Roy","location":13,"gender":"Male","email":"ramner23@alibaba.com"},
	{"id":77,"name":"Devy","location":28,"gender":"Male","email":"dbrearton24@cafepress.com"},
	{"id":78,"name":"Port","location":-30,"gender":"Male","email":"pgaskin25@illinois.edu"},
	{"id":79,"name":"Elene","location":16,"gender":"Female"},
	{"id":80,"name":"Hashim","location":-5,"gender":"Male","email":"hdunbabin27@foxnews.com"},
	{"id":81,"name":"Claresta","location":-11,"gender":"Female"},
	{"id":82,"name":"Isidora","location":-9,"gender":"Female"},
	{"id":83,"name":"Cornie","location":-27,"gender":"Female","email":"chousden2a@mashable.com"},
	{"id":84,"name":"Abdel","location":3,"gender":"Male","email":"agreenshields2b@netscape.com"},
	{"id":85,"name":"Gwenore","location":-30,"gender":"Female"},
	{"id":86,"name":"Nona","location":4,"email":"ntuckley2d@cam.ac.uk"},
	{"id":87,"name":"Lorri","location":-8,"gender":"Female","email":"ldeignan2e@stanford.edu"},
	{"id":88,"name":"Rayna","location":-7,"gender":"Female"},
	{"id":89,"name":"Hillyer","location":20,"gender":"Male","email":"hpalmer2g@cnbc.com"},
	{"id":90,"name":"Ruthe","location":-9},
	{"id":91,"name":"Sterne","location":22,"email":"sivanishin2i@theatlantic.com"},
	{"id":92,"name":"Carlie","location":-6,"email":"cbridgens2j@google.ca"},
	{"id":93,"name":"Avril","location":-26,"gender":"Female"},
	{"id":94,"name":"Izabel","location":16,"gender":"Female","email":"iheymann2l@china.com.cn"},
	{"id":95,"name":"Adolphus","location":6,"gender":"Male"},
	{"id":96,"name":"Bonni","location":26,"gender":"Female"},
	{"id":97,"name":"Tobi","location":0,"gender":"Female","email":"tchampney2o@va.gov"},
	{"id":98,"name":"Giacomo","location":2,"gender":"Male"},
	{"id":99,"name":"Federica","location":-24,"gender":"Female","email":"fbunworth2q@mayoclinic.com"},
	{"id":100,"name":"Jacques","location":-11,"gender":"Male","email":"jespie2r@cisco.com"}]`)

	// decoding JSON array to
	// the User array
	err := json.Unmarshal(Data, &User)

	if err != nil {
		fmt.Println(err)
	}

	// printing decoded array
	// values one by one
	for i := range User {
		fmt.Println("Name:" + User[i].Name + " - Gender:" + User[i].Gender + " - Email:" + User[i].Email)
		fmt.Println("Location: " + strconv.Itoa(User[i].Location))
	}
	fmt.Printf("Enter number of people you want to like: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Printf("Which users do you want to Like? (format: <id,id,..id>)")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your like list is: ", arr)
	for i := range User {
		for j := 0; j < size; j++ {
			if User[i].Id == arr[j] && User[i].Like != "💗" {
				User[i].Like = "💗"
			}
		}

	}
	for i := range User {
		fmt.Println(User[i].Name + " - " + User[i].Like)

	}
	for i := range User {
		if User[i].Like == "💗" {
			fmt.Println("You like:" + User[i].Name + User[i].Like)
		} else {
			continue
		}
	}
}
