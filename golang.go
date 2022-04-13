package main

//helo world
// func main() {
// 	fmt.Println("hello world")
// }

//number
// func main() {
// 	fmt.Println("satu = ", 1)
// 	fmt.Println("satu = ", 2)
// 	fmt.Println("tiga koma lima  = ", 3.5)
// }

//string
// func main() {
// 	fmt.Println("Eko")
// 	fmt.Println("Eko kurniawan")
// 	fmt.Println("Eko kurniawan perdana")
// len itu ngitung karakter dalam string
// klo [0 ngambil data string ke 1 tp ga lgsg string tp ke byte hars convert dlu ke string
// 	//lanjut di file strings.go
// 	fmt.Println(len("Eko"))
// 	fmt.Println("Eko Kurniawan"[0])
// 	fmt.Println("Eko Kurniawan perdana"[1])
// }

//bool
// func main() {
// 	fmt.Println("benar = ", true)
// 	fmt.Println("salah = ", false)
// }

//variabel
// func main() {
// 	//di golang 1 variabel 1 data,jd ga bsa d pke berulang.1 data unique
// 	// 1 variabel beda value bsa
// 	//var gawajib d tulis bsa ke gni :=
// 	var name string
// 	name = "Eko Kurniawan"
// 	fmt.Println(name)

// 	name = "Eko Khannedy"
// 	fmt.Println(name)

// 	var friendName = "Budi"
// 	fmt.Println(friendName)

// 	var age = 30
// 	fmt.Println(age)

// 	country := "Indonesia"
// 	fmt.Println(country)

// 	//bkin multiple variabel
// 	var (
// 		firstName = "Eko"
// 		lastName  = "Khannedy"
// 	)

// 	fmt.Println(firstName)
// 	fmt.Println(lastName)
// }

//constant
//multple constant
// func main() {
// 	const (
// 		firstName string = "Eko"
// 		lastName         = "Khannedy"
// 		value            = 1000
// 	)

// 	fmt.Println(firstName)
// 	fmt.Println(lastName)
// 	fmt.Println(value)
// }

//konversi tipe data
//tipe data angka dr 32 ke 64 dll

// func main() {
// 	var nilai32 int32 = 129
// 	var nilai64 int64 = int64(nilai32)
// 	var nilai8 int8 = int8(nilai32)

// 	fmt.Println(nilai32)
// 	fmt.Println(nilai64)
// 	fmt.Println(nilai8)

// 	var name = "Eko"
// 	//konversi e jadi byte
// 	var e byte = name[0]
// 	//konversi byte balik lg jd string
// 	var eString string = string(e)

// 	fmt.Println(name)
// 	fmt.Println(eString)
// }

//type deklarasi
//bkin alias
// func main() {
// 	type noKtp string
// 	var noKtpeko noKtp = "12312312312312"
// 	var merriedstatus = true
// 	fmt.Println(noKtpeko)
// 	fmt.Println(merriedstatus)
// }

//math operation
// func main() {

// 	var result = 10 + 10
// 	fmt.Println(result)

// 	var a = 10
// 	var b = 10
// 	var c = a * b
// 	fmt.Println(c)

// 	var i = 10
// 	i += 10 // i = i + 10
// 	fmt.Println(i)

// 	i++ // i = i + 1
// 	fmt.Println(i)

// 	var negative = -100
// 	var positive = +100
// 	fmt.Println(negative)
// 	fmt.Println(positive)
// }

//operation comparation
//hasilnya true false
// func main() {
// 	var name1 = "Eko"
// 	var name2 = "Budi"
// 	var result bool = name1 == name2
// 	fmt.Println(result)
// 	var value1 = 100
// 	var value2 = 200
// 	fmt.Println(value1 > value2)
// 	fmt.Println(value1 < value2)
// 	fmt.Println(value1 == value2)
// 	fmt.Println(value1 != value2)
// }
//operasi BOOLEAN
//cuma ada 3 && || !
// func main() {
// 	var ujian = 80
// 	var absensi = 75
// 	var lulusUjian = ujian >= 80
// 	var lulusAbsensi = absensi >= 80
// 	fmt.Println(lulusUjian)
// 	fmt.Println(lulusAbsensi)
// 	var lulus = lulusAbsensi && lulusUjian
// 	fmt.Println(lulus)
// 	fmt.Println(ujian >= 80 && absensi >= 80)
// }
//TIPE DATA ARRAY
//kumpulan data dgn tipe yang sama
// func main() {

// 	var names [3]string

// 	names[0] = "Eko"
// 	names[1] = "Kurniawan"
// 	names[2] = "Khannedy"

// 	fmt.Println(names[0])
// 	fmt.Println(names[1])
// 	fmt.Println(names[2])

// 	var values = [3]int{
// 		90,
// 		95,
// 		80,
// 	}

// 	fmt.Println(values)
// 	fmt.Println(values[0])
// 	fmt.Println(values[1])
// 	fmt.Println(values[2])

// 	fmt.Println(len(names))
// 	fmt.Println(len(values))

// 	var lagi [10]string
// 	fmt.Println(len(lagi))
// }

//tipe data slice mirip array
//bisa ambil smua atau sebagian arry
//slice pointer length dan capacity
//klo misal gatau kpasitas data pke [...]
// func main() {
// 	var months = [...]string{
// 		"Januari",
// 		"Februari",
// 		"Maret",
// 		"April",
// 		"Mei",
// 		"Juni",
// 		"Juli",
// 		"Agustus",
// 		"September",
// 		"Oktober",
// 		"November",
// 		"Desember",
// 	}

// 	var slice1 = months[4:7]
// 	fmt.Println(slice1)
// 	fmt.Println(len(slice1))
// 	fmt.Println(cap(slice1))

// 	//ubah isi arrray
// 	//months[5] = "Diubah"
// 	//fmt.Println(slice1)

//ubah isi slice array jga brubah
// 	//slice1[0] = "Mei Lagi"
// 	//fmt.Println(months)

// 	var slice2 = months[11:]
// 	fmt.Println(slice2)

// 	var slice3 = append(slice2, "Eko")
// 	fmt.Println(slice3)
// 	slice3[1] = "Bukan Desember"
// 	fmt.Println(slice3)

// 	fmt.Println(slice2)
// 	fmt.Println(months)

// 	newSlice := make([]string, 2, 5)

// 	newSlice[0] = "Eko"
// 	newSlice[1] = "Kurniawan"

// 	fmt.Println(newSlice)
// 	fmt.Println(len(newSlice))
// 	fmt.Println(cap(newSlice))

// 	copySlice := make([]string, len(newSlice), cap(newSlice))
// 	copy(copySlice, newSlice)
// 	fmt.Println(copySlice)

// 	iniArray := [5]int{1, 2, 3, 4, 5}
// 	iniSlice := []int{1, 2, 3, 4, 5}

// 	fmt.Println(iniArray)
// 	fmt.Println(iniSlice)
// }

//tipe data map
//func main() {

// person := map[string]string{
// 	"name":    "Eko",
// 	"address": "Subang",
// }
// //nambah data
// person["title"] = "Programmer"

// fmt.Println(person)
// fmt.Println(person["name"])
// fmt.Println(person["address"])

// var book map[string]string = make(map[string]string)
// book["title"] = "Belajar Go-Lang"
// book["author"] = "Eko"
// book["ups"] = "Salah"
// fmt.Println(book)
// fmt.Println(len(book))

// delete(book, "ups")

// fmt.Println(book)
// fmt.Println(len(book))
//}

//if expresion
// func main() {
// 	var name = "Kurniawan"

// 	if name == "Eko" {
// 		fmt.Println("Hello Eko")
// 	} else if name == "Joko" {
// 		fmt.Println("Hello Joko")
// 	} else if name == "Budi" {
// 		fmt.Println("Hello Budi")
// 	} else {
// 		fmt.Println("Hi, kenalan donk")
// 	}

// 	if length := len(name); length > 5 {
// 		fmt.Println("Terlalu Panjang")
// 	} else {
// 		fmt.Println("Nama sudah benar")
// 	}
// }

//switch expression
// func main() {

// 	name := "Kurniawan"

// 	switch name {
// 	case "Eko":
// 		fmt.Println("Hello Eko")
// 		fmt.Println("Hello Eko")
// 	case "Joko":
// 		fmt.Println("Hello Joko")
// 		fmt.Println("Hello Joko")
// 	default:
// 		fmt.Println("Kenalan Donk")
// 		fmt.Println("Kenalan Donk")
// 	}

//bolean
//switch length := len(name); length > 5 {
//case true:
//	fmt.Println("Nama terlalu panjang")
//case false:
//	fmt.Println("Nama sudah benar")
//}
//mirip if tp versi sederhana pake switch
// 	length := len(name)
// 	switch {
// 	case length > 10:
// 		fmt.Println("Nama terlalu panjang")
// 	case length > 5:
// 		fmt.Println("Nama lumayan panjang")
// 	default:
// 		fmt.Println("Nama sudah benar")
// 	}
// }

//loops
// func main() {
// 	// counter := 1

// 	// for counter <= 10 {
// 	// 	fmt.Println("perulangan ke", counter)
// 	// 	counter++
// 	// }

// 	// for counter := 1; counter <= 10; counter++ {
// 	// 	fmt.Println(counter)
// 	// }

// 	//skrng coba dari array
// 	slice := []string{"didi", "reja", "nugi"}

// 	for i := 0; i < len(slice); i++ {
// 		fmt.Println(slice[i])
// 	}
// 	//cara cepat ambil array data pake range
// 	for i, value := range slice {
// 		fmt.Println("Index", i, "=", value)
// 	}

// 	//loop bsa jga pake map
// 	person := make(map[string]string)
// 	person["name"] = "eko"
// 	person["title"] = "programmer"

// 	for key, value := range person {
// 		fmt.Println(key, "=", value)
// 	}
// }

//break dan continue golang
// func main() {
// 	// for i := 0; i < 10; i++ {
// 	// 	if i == 5 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println("Perulangan ke ", i)
// 	// }

// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			//menemukan angka genap ttp continue
// 			continue
// 		}
// 		//pas ktmu ganjil baru di d taro print
// 		fmt.Println("Perulangan ke", i)
// 	}
// }

//membuat function
// func hello() {
// 	fmt.Println("hello")
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		hello()
// 	}
//}
// pake parameter
// func sayHelloTo(firstName string, lastName string) {
// 	fmt.Println("hello", firstName, lastName)
// }

// func main() {
// 	//atau ke gni
// 	firstName := "didi"
// 	lastName := "perdana"
// 	//sayHelloTo("didi", "perdana")
// 	sayHelloTo(firstName, lastName)
// }

//function return value
// func getHello(nama string) string {
// 	// result := "hello"
// 	// // return "hello" + nama
// 	// return result + nama
// 	// //klo d bawah return ada fungsi atau variabel lg auto eror

// 	if nama == "" {
// 		return "hallo bro"
// 	} else {
// 		return "halo eko"
// 	}

// }

// func main() {
// 	result := getHello("Eko")
// 	fmt.Println(result)
// 	fmt.Println(getHello(""))
// }

// multiple return value
// func getFullName() (string, string, string) {
// 	return "Eko", "Kurniawan", "Khannedy"
// }

//kalo pake _ hiraukan return value jd cuma ambil1 aja ga smua
// func main() {
// 	firstName, _, _ := getFullName()
// 	fmt.Println(firstName)
// 	// fmt.Println(middleName)
// 	// fmt.Println(lastName)
// }

//return name value
// func getFullName2() (firstName string, middleName string, lastName string) {
// 	firstName = "Eko"
// 	middleName = "Kurniawan"
// 	lastName = "Khannedy"

// 	//return aja boleh
// 	return
// 	//return firstname,middlename,lastname juga boleh
// }

// func main() {
// 	a, b, c := getFullName2()
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

//veriadic function
// func sumAll(numbers ...int) int {
// 	total := 0
// 	for _, value := range numbers {
// 		total += value
// 	}
// 	return total
// }

// func main() {
// 	total := sumAll(10, 90, 30, 50, 40)
// 	fmt.Println(total)

// 	//contoh klo d variadic ada data slice atau array gmna mau nambahinnya
// 	slice := []int{10, 20, 30, 40, 50}
// 	//hrs ... krna buat masukin slice d argument function sumall
// 	total = sumAll(slice...)
// 	fmt.Println(total)
// }

//function value
//bkin variabel tp ambil datanya dr function
// func getGoodBye(name string) string {
// 	return "Good Bye " + name
// }

// func main() {

// 	sayGoodBye := getGoodBye

// 	result := sayGoodBye("Eko")
// 	fmt.Println(result)
// 	fmt.Println(getGoodBye("Eko"))

// }

//parameter as a function
// type Filter func(string) string

// func sayHelloWithFilter(name string, filter Filter) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	}else {
// 		return name
// 	}
// }

// func main() {
// 	sayHelloWithFilter("Eko", spamFilter)
// 	sayHelloWithFilter("Anjing", spamFilter)
// }

// anonymous function

// type Blacklist func(string) bool

// func registerUser(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("You are blocked", name)
// 	} else {
// 		fmt.Println("Welcome", name)
// 	}
// }

// //func blacklistAdmin(name string) bool {
// //	return name == "admin"
// //}
// //
// //func blacklistRoot(name string) bool {
// //	return name == "root"
// //}

// func main() {
// 	blacklist := func(name string) bool {
// 		return name == "admin"
// 	}

// 	registerUser("admin", blacklist)
// 	registerUser("eko", blacklist)

// 	//anonimous function bsa juga ke gni bsa juga ke atas
// 	registerUser("root", func(name string) bool {
// 		return name == "root"
// 	})
// 	registerUser("eko", func(name string) bool {
// 		return name == "root"
// 	})
// }

//reskursif
// func factorialLoop(value int) int {
// 	result := 1
// 	for i := value; i > 0; i-- {
// 		result *= i // result = result * i
// 	}
// 	return result
// }

// func factorialRecursive(value int) int {
// 	if value == 1 {
// 		return 1
// 	} else {
// 		return value * factorialRecursive(value-1)
// 	}
// }

// func main() {
// 	loop := factorialLoop(10)
// 	fmt.Println(loop)
// 	// fmt.Println(5 * 4 * 3 * 2 * 1)

// 	recursive := factorialRecursive(10)
// 	fmt.Println(recursive)
// }

//closure funsion
// func main() {
// 	name := "Eko"
// 	counter := 0

// 	increment := func() {
// 		name := "Budi"
// 		fmt.Println("Increment")
// 		counter++
// 		fmt.Println(name)
// 	}

// 	increment()
// 	increment()

// 	fmt.Println(counter)
// 	fmt.Println(name)
// }

//defer panic recover
//defer funsi yg d jadwalin untuk d eksekusi stlh fungsi laen d eksekusi
