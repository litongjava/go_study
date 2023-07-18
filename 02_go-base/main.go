package main

import (
  "encoding/json"
  "fmt"
  "math"
  "math/big"
  "math/rand"
  "os"
  "sort"
  "strconv"
  "strings"
  "time"
  "unicode/utf8"
)

func main() {
  test042201()
}

type xyzn struct {
  name string
  lat  float64
  long float64
}

type coordinate struct {
  d, m, s float64
  h       rune
}

func newLocation(lat, long coordinate) location {
  return location{lat.decimal(), long.decimal()}
}

type world struct {
  radius float64
}

//type report struct {
//  sol         int
//  temperature temperature
//  location    location
//}
//struct嵌入,嵌入之后,temperature的方法和location的方法可以被report使用
type report struct {
  sol
  temperature
  location
}

//方法转发
func (r report) average() celsius {
  return r.temperature.average()
}

func (r report) days(s2 sol) int {
  return r.sol.days(s2)
}

type temperature struct {
  high, low celsius
}
type location struct {
  lat, long float64
}

func (l location) days(l2 location) int {
  return 5
}

type celsius float64

func (t temperature) average() celsius {
  return (t.high + t.low) / 2
}

type sol int

func (s sol) days(s2 sol) int {
  days := int(s2 - s)
  if days < 0 {
    days = -days
  }
  return days
}

var t interface {
  talk() string
}

type martian struct{}

func (m martian) talk() string {
  return "nack nack"
}

type laser int

func (l laser) talk() string {
  return strings.Repeat("pew ", int(l))
}
func test042201() {
  t = martian{}
  fmt.Println(t.talk())

  t = laser(3)
  fmt.Println(t.talk())
}
func test032106() {
  report := report{sol: 15}
  fmt.Println(report.sol.days(1446))
  //语句报错
  //fmt.Println(report.days())
}
func test032103() {
  bradbury := location{-4.5895, 137.4417}
  t := temperature{high: -1.0, low: -78.0}
  report := report{
    sol:         15,
    temperature: t,
    location:    bradbury,
  }
  fmt.Printf("%+v\n", report)
  fmt.Printf("a balmy %v ℃\n", report.temperature.high)
  fmt.Println(report.temperature.average())
  fmt.Println(report.average())
}
func (w world) distance(p1, p2 location) float64 {
  s1, c1 := math.Sincos(rad(p1.lat))
  s2, c2 := math.Sincos(rad(p2.lat))
  clong := math.Cos(rad(p1.long - p2.long))
  return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
  return deg * math.Pi / 180
}

func test032007() {
  var mars = world{radius: 3389.5}
  spirit := location{-14.5684, 175.472636}
  opportunity := location{-1.9462, 354.4734}

  dist := mars.distance(spirit, opportunity)
  fmt.Printf("%.2f km\n", dist)
}
func test032004() {
  //curiosity := location{-23.0, -22.0}
  lat := coordinate{4, 35, 22.2, 'S'}
  long := coordinate{137, 26, 30.12, 'E'}
  curiosity := newLocation(lat, long)
  fmt.Println(curiosity)
}
func (c coordinate) decimal() float64 {
  sign := 1.0
  switch c.h {
  case 'S', 'W', 's', 'w':
    sign = -1
  }
  return sign * (c.d + c.m/60 + c.s/3600)
}
func test032002() {
  lat := coordinate{4, 35, 22.2, 'S'}
  long := coordinate{137, 26, 30.12, 'E'}
  fmt.Println(lat.decimal(), long.decimal())
}
func test0190301() {
  type xy struct {
    Lat  float64 `json:"latitude"`
    Long float64 `json:"longitude"`
  }
  curiosity := xy{-4.5895, 127.4417}
  bytes, err := json.Marshal(curiosity)
  exitOnError(err)
  fmt.Println(string(bytes))
}
func test03191101() {
  curiosity := location{-4.5895, 137.4417}
  bytes, err := json.Marshal(curiosity)
  exitOnError(err)
  fmt.Println(string(bytes))
}

func exitOnError(err error) {
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
func test03190901() {
  lats := []float64{-4.5895, -14.5684, -1.9462}
  longs := []float64{137.4417, 175.472636, 354.4734}

  locations := []xyzn{
    {name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
    {name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
    {name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
  }
  fmt.Println(lats, longs, locations)
}

func test03190801() {
  brandbury := location{-4.5895, 137.4417}
  form(brandbury)
  fmt.Println(brandbury)
}

func form(brandbury location) {
  brandbury.long += 0.10
}
func test03190601() {
  brandbury := location{-4.5895, 137.4417}
  curiousity := brandbury

  curiousity.long += 0.0106
  fmt.Println(brandbury, curiousity)
}
func test03190502() {
  curiosity := location{-4.4895, 174.4417}
  fmt.Printf("%v\n", curiosity)
  fmt.Printf("%+v\n", curiosity)
}
func test03190501() {
  opportunity := location{lat: -1.9462, long: 354.4734}
  fmt.Println(opportunity)

  insight := location{lat: 4.5, long: 135.9}
  fmt.Println(insight)

  spirit := location{-14.5684, 175.472636}
  fmt.Println(spirit)
}

type dis location

func distance(lat1, long1 location) dis {
  return dis{0.0, 0.0}
}
func test03190201() {
  var spirit location
  spirit.lat = -14.5684
  spirit.long = 175.472636

  var opportunity location
  opportunity.lat = -1.9462
  opportunity.long = 354.4734

  fmt.Println(spirit, opportunity)
}
func test03190101() {
  var curiousity struct {
    lat  float64
    long float64
  }
  curiousity.lat = -4.5895
  curiousity.long = 137.4417

  fmt.Println(curiousity.lat, curiousity.long)
  fmt.Println(curiousity)
}
func homework031814() {
  text := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
  fields := strings.Fields(text)
  //fmt.Println(fields)
  fmt.Printf("%v \n", len(fields))
  //stringSlice := strings.Split(text, " ")
  //fmt.Printf("%T\n", stringSlice)
  //fmt.Println(stringSlice)

  countMap := make(map[string]int)
  for _, e := range fields {
    e = strings.Trim(e, ".")
    e = strings.Trim(e, ",")
    e = strings.Trim(e, ";")
    countMap[strings.ToLower(e)]++
  }

  for k, v := range countMap {
    fmt.Printf("%v %v\n", k, v)
  }
}
func test031812() {
  temperatures := []float64{
    -28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
  }
  set := make(map[float64]bool)
  for _, t := range temperatures {
    set[t] = true
  }

  if set[-28.0] {
    fmt.Println("set member")
  }
  fmt.Println(set)
  //有序输出
  unique := make([]float64, 0, len(set))

  for k := range set {
    unique = append(unique, k)
  }
  sort.Float64s(unique)
  fmt.Println(unique)
}
func test031810() {
  temperatures := []float64{
    -28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
  }
  //Map<float,[]float64>
  groups := make(map[float64][]float64)

  for _, v := range temperatures {
    //对温度进行分组,以10位跨度
    //math.Trunc将浮点类型截断为整数
    g := math.Trunc(v/10) * 10
    //附加元素并赋值
    groups[g] = append(groups[g], v)
  }

  for k, v := range groups {
    fmt.Printf("%v: %v\n", k, v)
  }
}
func test031808() {
  temperatures := []float64{
    -28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
  }
  frequcency := make(map[float64]int)

  for _, t := range temperatures {
    //添加元素,默认值是t类型的默认值,如果没有默认值则加+1
    frequcency[t]++
  }

  for t, num := range frequcency {
    fmt.Printf("%+.2f occurs %d times\n", t, num)
  }
}
func test031807() {
  temperature := make(map[float64]int, 8)
  fmt.Println(len(temperature))
}
func test031805() {
  planets := map[string]string{
    "Earth": "Sector ZZ9",
    "Mars":  "Sector ZZ9",
  }
  planetsMarkII := planets
  planets["Earth"] = "whoops"

  fmt.Println(planets)
  fmt.Println(planetsMarkII)

  delete(planets, "Earth")
  fmt.Println(planetsMarkII)
}
func test031803() {
  temperature := map[string]int{
    "Earth": 15,
    "Mars":  -65,
  }
  temp := temperature["Earth"]
  fmt.Printf("On average the Earth is %v ℃.\n", temp)

  temperature["Earth"] = 16
  temperature["Venus"] = 464

  fmt.Println(temperature)

  moon := temperature["Moon"]
  fmt.Println(moon)

  //moon是具体的值,ok是布尔类型,如果对应的key有value,ok的值为true
  if moon, ok := temperature["Moon"]; ok {
    fmt.Printf("On average the moon is %v ℃.\n", moon)
  } else {
    fmt.Println("Where is the moon?")
  }
}
func test031802() {
  temperature := map[string]int{
    "Earth": 15,
    "Mars":  -65,
  }
  temp := temperature["Earth"]
  fmt.Printf("On average the Earth is %v ℃.\n", temp)

  temperature["Earth"] = 16
  temperature["Venus"] = 464

  fmt.Println(temperature)

  moon := temperature["Moon"]
  fmt.Println(moon)
}
func homework031713() {
  slice := make([]string, 0, 4)
  for i := 0; i < 10; i++ {
    slice = append(slice, "Ping")
    dump("slice", slice)
  }
}
func test03171101() {
  twoWorlds := terraformString("New", "Venus", "Mars")
  fmt.Println(twoWorlds)

  planets := []string{"Venus", "Mars", "Jupiter"}

  newPlanets := terraformString("New", planets...)
  fmt.Println(newPlanets)
}
func terraformString(prefix string, worlds ...string) []string {
  newWorlds := make([]string, len(worlds))

  for i := range worlds {
    newWorlds[i] = prefix + "  " + worlds[i]
  }
  return newWorlds
}
func dump(label string, slice []string) {
  fmt.Printf("%v: length %v,capacity %v %v\n", label, len(slice), cap(slice), slice)
}
func text031709() {
  dwarfs := make([]string, 0, 10)
  dump("dwarfs", dwarfs)

  dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
  dump("dwarfs", dwarfs)
}

func test03170701() {
  planets := []string{
    "Mercury", "Venus", "Earth", "Mars",
    "Jupiter", "Saturn", "Uranus", "Neptune",
  }
  //分配了一个新的内存空间
  terrestrial := planets[0:4:4]
  //追加元素,进行扩容
  worlds := append(terrestrial, "Cerer")
  dump("plants", planets)
  dump("terrestrial", terrestrial)
  dump("worlds", worlds)
}
func test03170501() {
  dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
  dwarfs2 := append(dwarfs1, "Orcus")
  dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")

  dump("dwarfs1", dwarfs1)
  dump("dwarfs2", dwarfs2)
  dump("dwarfs3", dwarfs3)
}
func test03170301() {
  dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
  dump("dwarfs", dwarfs)
  dump("dwarfs[1:2]", dwarfs[1:2])
  dump("dwarfs[1:3]", dwarfs[1:3])
  dump("dwarfs[1:4]", dwarfs[1:4])
  dump("dwarfs[1:5]", dwarfs[1:5])
  dump("dwarfs[0:5]", dwarfs[0:5])

  dump("dwarfs[2:3]", dwarfs[2:3])
  dump("dwarfs[4:5]", dwarfs[4:5])
}

func test03170101() {
  dwarfs := []string{"Cerer", "Pluto", "Haumea", "Makemake", "Eris"}
  dwarfs = append(dwarfs, "Orcus")
  dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
  fmt.Println(dwarfs)
  fmt.Println(len(dwarfs))
}

type Planets []string

func (p Planets) terraform() {
  for i, v := range p {
    //fmt.Printf("%v %v \n", i, v)
    p[i] = "New " + v
  }
}
func homework031611() {
  planets := []string{
    "Mars", "Uranus", "Neptune",
  }
  fmt.Printf("%T\n", planets)
  p := Planets(planets)
  fmt.Printf("%T\n", p)
  p.terraform()
  fmt.Println(p)
}
func test03160901() {
  planets := []string{
    "Mercury", "Venus", "Earth", "Mars",
    "Juipter", "Saturn", "Uranus", "Neptune",
  }
  fmt.Printf("%T\n", planets)
  slice := sort.StringSlice(planets)
  fmt.Printf("%T\n", slice)
  slice.Sort()
  fmt.Println(planets)
}
func test03160701() {
  plants := []string{" Venus  ", "Earth  ", " Mars"}
  hyperspace(plants)
  fmt.Println(strings.Join(plants, " "))
}

func hyperspace(worlds []string) {
  for i := range worlds {
    worlds[i] = strings.TrimSpace(worlds[i])
  }
}
func test03160501() {
  dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
  dwarfSlice := dwarfArray[:]
  dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
  fmt.Println(dwarfSlice)
  fmt.Println(dwarfs)
  fmt.Printf("%T %T\n", dwarfArray, dwarfs)
}
func homework0315() {
  //放弃
}
func test03160203() {
  question := "如何学好golang"
  fmt.Println(question[:6])
}
func test03160202() {
  neptune := "Neptune"
  tune := neptune[3:]
  fmt.Println(tune)

  neptune = "Poseidon"
  fmt.Println(tune)
}
func test03160201() {
  plants := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
  }
  terrestrial := plants[:4]
  gasGiants := plants[4:6]
  iceGiants := plants[6:]
  fmt.Println(terrestrial, gasGiants, iceGiants)

  allPlants := plants[:]
  fmt.Println(allPlants)

}
func test03151001() {
  var board [8][8]string
  board[0][0] = "r"
  board[0][7] = "r"

  for column := range board[1] {
    board[1][column] = "p"
  }
  fmt.Println(board)
}
func test03150902() {
  plants := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
  }
  terraform(plants)
  fmt.Println(plants)
}
func terraform(planets [8]string) {
  for i := range planets {
    planets[i] = "New " + planets[i]
  }
}
func test03150901() {
  plants := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
  }
  //此处发生了copy
  plantsMarkII := plants
  plants[2] = "whoops"
  fmt.Println(plants)
  fmt.Println(plantsMarkII)
}
func test03150702() {
  dwarfs := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
  for i, dwarf := range dwarfs {
    fmt.Println(i, dwarf)
  }
}
func test03150701() {
  dwarfs := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
  for i := 0; i < len(dwarfs); i++ {
    dwarf := dwarfs[i]
    fmt.Println(i, dwarf)
  }
}
func test031506() {
  //dwarfs := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
  plants := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
  }
  fmt.Println(plants)
}
func test031503() {
  var plants [8]string
  i := 8
  plants[i] = "Pluto"
  pluto := plants[i]
  fmt.Println(pluto)
}
func test031502() {
  var planets [8]string
  planets[0] = "Mercury"
  planets[1] = "Venus"
  planets[2] = "Earth"

  earth := planets[2]
  fmt.Println(earth)
  fmt.Println(len(planets))
  fmt.Println(planets[3] == "")
}
func test03130803() {
  func() {
    fmt.Println("Funcitons anonymous")
  }()
}
func test03130802() {
  f := func(message string) {
    fmt.Println(message)
  }
  f("Go to the party")
}

var f = func() {
  fmt.Print("Dress up for the masquerade")
}

func test031308() {
  f()
}

type kelvin float64

func test031304() {
  measureTemperature(3, fakeSensor)
}
func measureTemperature(samples int, sensor func() kelvin) {
  for i := 0; i < samples; i++ {
    k := sensor()
    fmt.Printf("%v °K\n", k)
    time.Sleep(time.Second)
  }
}
func fakeSensor() kelvin {
  return kelvin(rand.Intn(151) + 150)
}

func test031302() {
  //将函数赋值给变量,但是函数没有执行
  sensor := fakeSensor
  fmt.Println(sensor())

  sensor = realSensor
  fmt.Print(sensor())
}

func realSensor() kelvin {
  return 0
}
func test021204() {
  var k kelvin = 294.0
  var c celsius

  c = kelvinToCelsius(k)
  fmt.Println(c)
  c = k.celsius()
  fmt.Println(c)
}

func kelvinToCelsius(k kelvin) celsius {
  return celsius(k - 273.15)
}

//声明celsius是kelvin的一个方法,但是这种关系在go中称为关联
func (k kelvin) celsius() celsius {
  return celsius(k - 273.15)
}
func test03120201() {
  const degrees = 20
  var temperature celsius = degrees
  temperature += 10
  //var warmUp float64=10
  //temperature+warmUp
  fmt.Println(temperature)
}
func test031107() {
  kelvin := 294.0
  celsius := KelvinToCelsius(kelvin)
  fmt.Println(kelvin, "ºK is", celsius, "ºC")
}
func KelvinToCelsius(k float64) float64 {
  return k - 273.15
}
func homework0211() {
  text := "true"
  if text == "false" || text == "yes" || text == "1" {
    fmt.Println(false)
  } else if text == "true" || text == "no" || text == "0" {
    fmt.Println(true)
  } else {
    fmt.Println("不处理")
  }
}
func test021111() {
  lunched := false
  var lunchedInt int
  if lunched {
    lunchedInt = 1
  } else {
    lunchedInt = 0
  }
  fmt.Println(lunchedInt)
}
func test021110() {
  lunched := false
  lunchedText := fmt.Sprintf("%v", lunched)
  fmt.Println(lunchedText)
}
func test0211075() {
  fmt.Println(strconv.Atoi("1"))
  fmt.Println(strconv.Atoi("10"))
  str, err := strconv.Atoi("1000000000000000000000000000000000")
  if err != nil {
    fmt.Println("something went to wrong")
  } else {
    fmt.Println(str)
  }
}
func test0211074() {
  countdown := 10
  str := fmt.Sprintf("Launch In %v seconds", countdown)
  fmt.Println(str)
}
func test0211073() {
  countdown := 10
  str := "Launch In " + strconv.Itoa(countdown) + " seconds"
  fmt.Println(str)
}
func test0211072() {
  fmt.Println(string(65))
  fmt.Println(string(42000000000000))
  fmt.Println(string(42000000000001))
}
func test0211071() {
  var pi rune = 960
  var alpha rune = 940
  var omega rune = 969
  var bang byte = 33
  fmt.Println(string(pi), string(alpha), string(omega), string(bang))
}
func test021106() {
  v := 42
  if v >= 0 && v <= math.MaxUint8 {
    v8 := uint8(v)
    fmt.Println("converted", v8)
  }
}
func test021105() {
  //var bh float64 = 32768
  //var b = int16(bh)
  //fmt.Println(b)

  var bh float64 = 32768
  if bh < math.MinInt16 || bh > math.MaxInt16 {
    fmt.Println("handle out of range value")
  } else {
    fmt.Println("OK")
  }
}
func test021104() {
  //age := 41
  //marsAge := float64(41)
  earthDay := 365.2425
  fmt.Println(int(earthDay))
}
func test0210182() {
  message := "Hola Estación Espacial Internacional"
  for _, c := range message {
    if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
      c += 13
    }
    fmt.Printf("%c", c)
  }
}
func test0210181() {
  message := "L fdph, L vdz, L frqtxhuhg"
  for _, c := range message {
    //只处理字母
    if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
      c -= 3
    }
    fmt.Printf("%c", c)
  }
}
func test021017() {
  question := "美国 United States"
  for _, c := range question {
    fmt.Printf("%c\n", c)
  }
}
func test0210161() {
  question := "美国 United States"
  //输出20,但是上面的字符串实际上并不是二十个字符
  fmt.Println(len(question), "bytes")
  //统计字符数
  fmt.Println(utf8.RuneCountInString(question), "runes")

  c, size := utf8.DecodeRuneInString(question)
  fmt.Printf("First rune:%c,%v bytes\n", c, size)
}
func test021016() {
  message := "uv vagreangvbany fcnpr fgngvba"
  for i := 0; i < len(message); i++ {
    c := message[i]
    if c >= 'a' && c <= 'z' {
      c += 13
      if c > 'z' {
        c -= 26
      }
    }
    fmt.Printf("%c", c)
  }
}
func test021013() {
  message := "I'am Ping E Lee"
  fmt.Println(len(message))
}
func test021012() {
  c := 'c'
  c = c - 'a' + 'A' //=>C=67=99-97+65
  fmt.Printf("%c ", c)
  g := 'g'
  g = g - 'a' + 'A'
  fmt.Printf("%c ", g)

  d := 'd'
  //d = d - 'a' + 'A'
  d = d - 32
  fmt.Printf("%c ", d)
}
func test021011() {
  c := 'a'
  c = c + 3
  fmt.Printf("%c ", c)

  if c > 'z' {
    c = c - 26
  }
  fmt.Printf("%c\n", c)
}
func test021010() {
  str := "shalom"
  for i := 0; i < len(str); i++ {
    fmt.Println(str[i])
    fmt.Printf("%c\n", str[i])
  }
}
func test02105() {
  var pi rune = 960
  var alpha rune = 940
  var omega rune = 969
  var bang byte = 33
  //打印出code porint的值
  fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
  //打印出字符
  fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
}
func test02103() {
  fmt.Println("C:\\go")
  //fmt.Println("C:\go")//报错
  fmt.Println(`C:\go`)
}
func homework0295() {
  const lightPerYear = 9460730472581
  const distance = 236000000000000000
  fmt.Print(distance / lightPerYear)
}
func test02102() {
  fmt.Println("Andromeda Galaxy is", 24000000000000000000/299792/85400, "light days away")
}
func test02101() {
  //  const distance uint64 =24000000000000000000
  const distance = 24000000000000000000
  //fmt.Println(distance)

}
func test02922() {
  distance := new(big.Int)
  distance.SetString("24000000000000000", 10)
  fmt.Println(distance)
}
func test02921() {
  lightSpeed := big.NewInt(299792)
  secondPerDay := big.NewInt(86400)
  fmt.Println(lightSpeed, secondPerDay)
}
func test0291() {
  distance := 24e18
  fmt.Println(distance)
}
func homework028() {
  total := 0.0
  for {
    intn := rand.Intn(3)
    if intn == 0 {
      total += 0.05
    } else if intn == 1 {
      total += 0.1
    } else {
      total += 0.25
    }
    fmt.Printf("$%5.2f\n", total)
    if total > 20 {
      break
    }
  }
}
func test02811() {
  unix := time.Unix(100000000000, 0)
  fmt.Println(unix)
}
func test0288() {
  fmt.Println(math.MaxInt16)
  fmt.Println(math.MinInt16)
}
func test0287() {
  var green uint8 = 3
  fmt.Printf("%08b\n", green)
  green++
  fmt.Printf("%08b\n", green)
}
func test0286() {
  var red uint8 = 255
  red++
  fmt.Println(red)

  var number int8 = 127
  number++
  fmt.Println(number)
}
func test0285() {
  var red, green, blue uint8 = 0, 141, 123
  fmt.Printf("%x %x %x\n", red, green, blue)
  //也可以指定最小宽度和填充：
  fmt.Printf("color: #%02x%02x%02x;\n", red, green, blue)
}

var era = "AD"

func test0283() {
  year := 2018
  fmt.Printf("Type %T for %v\n", year, year)

  a := "text"
  fmt.Printf("Type %T for %[1]v\n", a, a)

  b := 42
  fmt.Printf("Type %T for %[1]v\n", b, b)

  c := 3.14
  fmt.Printf("Type %T for %[1]v\n", c, c)

  d := true
  fmt.Printf("Type %T for %[1]v\n", d, d)
}
func test0281() {
  //year:=2022
  //var year=2022
  //var year int=2022
  //var month uint=12
  //var red, green, blue uint8 = 0, 141, 123
  //var red, green, blue uint8 = 0x00, 0x8d, 0xd5
}
func homework27() {
  total := 0.0
  for {
    intn := rand.Intn(3)
    if intn == 0 {
      total += 0.05
    } else if intn == 1 {
      total += 0.1
    } else {
      total += 0.25
    }
    fmt.Printf("%5.2f\n", total)
    if total > 20 {
      break
    }
  }
}

/*
rand.Intn()返回的是整数还是小数
返回的值整数
*/
func testRandIntn() {
  for i := 0; i < 10; i++ {
    //生成0,1,2
    num := rand.Intn(3)
    fmt.Println(num)
  }
}

func test0278() {
  piggyBank := 0.1
  piggyBank += 0.2
  fmt.Println(piggyBank == 0.3)

  fmt.Println(math.Abs(piggyBank)-0.3 < 0.0001)
}
func test0277() {
  third := 1.0 / 3
  fmt.Println(third + third + third)

  piggyBank := 0.1
  piggyBank += 0.2
  fmt.Println(piggyBank)

  celsius := 21.0
  fmt.Print(celsius/5.0*9.0+32, "℉\n")
  fmt.Print(9.0/5.0*celsius+32, "℉\n")
  fmt.Print(celsius*9.0/5.0+32, "℉\n")
}
func test0276() {
  third := 1.0 / 3
  fmt.Printf("%05.2f\n", third)
}
func test0275() {
  third := 1.0 / 3
  //小数点后16位
  fmt.Println(third)
  //小数点后16位
  fmt.Printf("%v\n", third)
  //小数点后6位
  fmt.Printf("%f\n", third)
  //小数点后3为
  fmt.Printf("%.3f\n", third)
  //小数点后2位,总长度4位
  fmt.Printf("%4.2f\n", third)
}
func test0274() {
  var price float64
  fmt.Println(price)
}
func test0272() {
  var pi64 = math.Pi
  var pi32 float32 = math.Pi

  fmt.Println(pi64)
  fmt.Println(pi32)
}
func test0271() {
  // days := 365.2425
  // var days = 365.2425
  // var days float64 = 365.2425
  // var answer float64 = 42
  // dog := 42.0

}
func homework03() {

  daysInMounth := 31
  for i := 0; i < 10; i++ {
    year := rand.Intn(10000)
    month := rand.Intn(12) + 1
    switch month {
    case 2:
      /*
      	1.能被4整除,但是不能被100整除的年份 比如2020年
      	2.能被400整除的年份 比如2000年
      */
      if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
        daysInMounth = 29
      } else {
        daysInMounth = 29
      }
    case 4, 6, 9, 11:
      daysInMounth = 30
    }

    day := rand.Intn(daysInMounth) + 1
    fmt.Println(era, year, month, day)
  }
}
func test0253() {
  year := 2018
  month := rand.Intn(12) + 1
  daysInMounth := 31

  switch month {
  case 2:
    daysInMounth = 28
  case 4, 6, 9, 11:
    daysInMounth = 30
  }

  day := rand.Intn(daysInMounth) + 1
  fmt.Println(era, year, month, day)
}
func test0252() {
  // var count=0
  // for count=10; count>0; count--{
  // 	fmt.Println(count)
  // }

  for count := 10; count > 0; count-- {
    fmt.Println(count)
  }

  if num := rand.Intn(3); num == 0 {
    fmt.Println("Space Advenrures")
  } else if num == 1 {
    fmt.Println("SpaceX")
  } else {
    fmt.Println("Virgin Galactic")
  }
  switch num := rand.Intn(3); num {
  case 0:
    fmt.Println("Space Advenrures")
  case 1:
    fmt.Println("SpaceX")
  case 2:
    fmt.Println("Virgin Galactic")
  default:
    fmt.Println("Random spaceline #", num)
  }
}
func test0251() {
  var count = 0
  for count < 10 {
    var num = rand.Intn(10) + 1
    fmt.Println(num)
    count++
  }
}
func homework02() {
  var num = rand.Intn(100)
  fmt.Println(num)

  var rnd = rand.Intn(100)
  for {
    if rnd > num {
      fmt.Printf("%v 大于 %v\n", rnd, num)
    } else if rnd < num {
      fmt.Printf("%v 小于 %v\n", rnd, num)
    } else {
      fmt.Println("等于")
      break
    }
    rnd = rand.Intn(100)
  }
}
func test38() {
  var count = 10
  for count > 0 {
    fmt.Println(count)
    time.Sleep(time.Second)
    count--
  }
  fmt.Println("liftoff")
}
func test37() {
  var room = "lake"
  switch {
  case room == "cave":
    fmt.Println("You find yourself in a dimly lit cavern.")
  case room == "lake":
    fmt.Println("The ice seems solid enough.")
    fallthrough
  case room == "underwater":
    fmt.Println("The water is freezing cold.")
  }
}
func test36() {
  fmt.Println("There is a cavern entrance here and a path to the east")
  var command = "go inside"

  switch command {
  case "go east":
    fmt.Println("You head further up the mountain")
  case "enter cave", "go inside":
    fmt.Println("You find yourself in a dimly lit cavern")
  case "read sign":
    fmt.Println("The sign reads 'No Minors'.")
  default:
    fmt.Println("Didn't quite get that")
  }
}
func test35() {
  var haveTorch = true
  var litTorch = false
  if !haveTorch || !litTorch {
    fmt.Println("Nothing to see here")
  }
}
func test34() {
  fmt.Println("This year is 2100,should you leap")
  var year = 2100
  var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
  if leap {
    fmt.Println("Look before you leap")
  } else {
    fmt.Println("Keep you feet on the ground")
  }
}
func test33() {
  var command = "go east"
  if command == "go east" {
    fmt.Println("You head further up the mountain.")
  } else if command == "go inside" {
    fmt.Println("You enter the cave where you live out the rest of you life")
  } else {
    fmt.Println("Didn't quit get that.")
  }
}
func test32() {
  fmt.Println("There is a sign near the entrance that reads 'No Minors'")
  var age = 41
  var minor = age < 18
  fmt.Printf("at age %v, am I minor? %v", age, minor)
}
func test31() {
  fmt.Println("You find yourself in a dimly lit cavern")

  var command = "walk outside"
  var exit = strings.Contains(command, "outside")
  fmt.Println("You leave the cave", exit)
}
func homework01() {
  var distance = 56000000
  var day = 28
  var totalHour = day * 24
  fmt.Println(distance / totalHour)
}
func test08() {
  var num = rand.Intn(10) + 1
  fmt.Println(num)

  num = rand.Intn(10) + 1
  fmt.Print(num)
}
func test07() {
  var age = 41
  age = age + 1
  age += 1
  age++
  fmt.Println(age)
}
func test06() {
  var weight = 149.0
  weight = weight * 0.3783
  fmt.Println(weight)
  weight *= 0.3783
  fmt.Print(weight)
}
func test05() {
  // var distance = 56000000
  // var speed = 100800

  var (
    distance = 56000000
    speed    = 100800
  )

  // var distance, sppend = 56000000, 100800
  const hoursPerDay, minutesPerHour = 24, 60

  fmt.Println(distance, speed)
}
func test04() {
  const lightSpeed = 299792 //km/s
  var distance = 560000000  //km
  fmt.Println(distance/lightSpeed, "seconds")

  distance = 401000000
  fmt.Println(distance/lightSpeed, "seconds")
}

func test03() {
  fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
  fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
func test02() {
  fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.783)
  fmt.Printf("and I would be %v years old\n", 41*365/687)
  fmt.Printf("My weight on the surface of Mars is %v lbs and I would be %v years old\n", 149.0*0.783, 41*365/687)
}

func test01() {
  fmt.Print("My weight on the surface of Mars is ")
  fmt.Print(149.0 * 0.783)
  fmt.Print(" lbs and I would be ")
  fmt.Print(41 * 365 / 687)
  fmt.Println(" years old")
  fmt.Println("My weight on the surface of Mars is", 149.0*0.783, "lbs and I would be", 41*365/687, "years old")
}
