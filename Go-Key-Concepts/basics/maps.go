package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

// Maps
// maps is referenced data type
// like slice
// they can be changed in function
// unlike primitive data types like variable, string, int and etc.

// Map Example 1
type user struct {
	name   string
	number int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes of names/phoneNumber slice")
	}

	userMap := make(map[string]user, len(names))

	//for _, name := range names {
	//	for _, phoneNumber := range phoneNumbers {
	//		userMap[name] = user{
	//			name:        name,
	//			number: phoneNumber,
	//		}
	//	}
	//}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:   name,
			number: phoneNumber,
		}
	}

	return userMap, nil
}

func testMap1(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("==================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value (user struct):\n", name)
		fmt.Println(" - name: ", users[name].name)
		fmt.Println(" - phone number: ", users[name].number)
	}
}

// Map Example 2
type userSchedule struct {
	user
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]userSchedule, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	//var canBeUserDeleted bool = existingUser.scheduledForDeletion

	if !ok {
		return false, errors.New("not found")
	}
	if !existingUser.scheduledForDeletion {
		return false, errors.New("user cannot be deleted")
	}

	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}

	return false, errors.New("something went wrong")
}

func testMap2(users map[string]userSchedule, name string) {
	fmt.Printf("Attempting to delete %v...\n", name)
	defer fmt.Println("==================")

	isUserDeleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if isUserDeleted {
		fmt.Println("Deleted: ", name)
		return
	}
	fmt.Println("Did not delete: ", name)
}

// Count Instances
// Map Example 3
func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}

func testMap3(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))
	defer fmt.Println("==================")

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, key := range ids {
		val := counts[key]
		fmt.Printf(" - %s: %d\n", key, val)
	}
}

func generateRandomUserIDs() []string {
	const generatedIDsNumber = 10000
	const IDsDigitLimit = 2

	userIDs := make([]string, 0, generatedIDsNumber)

	for i := 0; i < generatedIDsNumber; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:IDsDigitLimit])
	}

	return userIDs
}

// Map Example 4
// nested maps
// rune is character type
func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)

	//for i := 0; i < len(names); i++ {
	//	name := names[i]
	//	iFirstLetter := name[0]

	//	count := nameCounts[iFirstLetter][name]
	//	count++
	//	nameCounts[iFirstLetter][name] = count
	//}

	for _, name := range names {
		if name == "" {
			continue
		}
		firstLetter := rune(name[0])

		_, ok := nameCounts[firstLetter]
		if !ok {
			nameCounts[firstLetter] = make(map[string]int)
		}

		nameCounts[firstLetter][name]++
	}

	return nameCounts
}

func testMap4(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))
	defer fmt.Println("==================")

	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
}

func getNames(length int) []string {
	names := []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker", "Parker", "Parker", "Collin", "Hayden", "George", "Bradley", "Mitchell", "Devon", "Ricardo", "Shawn", "Taylor", "Nicolas", "Gregory", "Francisco", "Liam", "Kaleb", "Preston", "Erik", "Alexis", "Owen", "Omar", "Diego", "Dustin", "Corey", "Fernando", "Clayton", "Carter", "Ivan", "Jaden", "Javier", "Alec", "Johnathan", "Scott", "Manuel", "Cristian", "Alan", "Raymond", "Brett", "Max", "Drew", "Andres", "Gage", "Mario", "Dawson", "Dillon", "Cesar", "Wesley", "Levi", "Jakob", "Chandler", "Martin", "Malik", "Edgar", "Sergio", "Trenton", "Josiah", "Nolan", "Marco", "Drew", "Peyton", "Harrison", "Drew", "Hector", "Micah", "Roberto", "Drew", "Brady", "Erick", "Conner", "Jonah", "Casey", "Jayden", "Edwin", "Emmanuel", "Andre", "Phillip", "Brayden", "Landon", "Giovanni", "Bailey", "Ronald", "Braden", "Damian", "Donovan", "Ruben", "Frank", "Gerardo", "Pedro", "Andy", "Chance", "Abraham", "Calvin", "Trey", "Cade", "Donald", "Derrick", "Payton", "Darius", "Enrique", "Keith", "Raul", "Jaylen", "Troy", "Jonathon", "Cory", "Marc", "Eli", "Skyler", "Rafael", "Trent", "Griffin", "Colby", "Johnny", "Chad", "Armando", "Kobe", "Caden", "Marcos", "Cooper", "Elias", "Brenden", "Israel", "Avery", "Zane", "Zane", "Zane", "Zane", "Dante", "Josue", "Zackary", "Allen", "Philip", "Mathew", "Dennis", "Leonardo", "Ashton", "Philip", "Philip", "Philip", "Julio", "Miles", "Damien", "Ty", "Gustavo", "Drake", "Jaime", "Simon", "Jerry", "Curtis", "Kameron", "Lance", "Brock", "Bryson", "Alberto", "Dominick", "Jimmy", "Kaden", "Douglas", "Gary", "Brennan", "Zachery", "Randy", "Louis", "Larry", "Nickolas", "Albert", "Tony", "Fabian", "Keegan", "Saul", "Danny", "Tucker", "Myles", "Damon", "Arturo", "Corbin", "Deandre", "Ricky", "Kristopher", "Lane", "Pablo", "Darren", "Jarrett", "Zion", "Alfredo", "Micheal", "Angelo", "Carl", "Oliver", "Kyler", "Tommy", "Walter", "Dallas", "Jace", "Quinn", "Theodore", "Grayson", "Lorenzo", "Joe", "Arthur", "Bryant", "Roman", "Brent", "Russell", "Ramon", "Lawrence", "Moises", "Aiden", "Quentin", "Jay", "Tyrese", "Tristen", "Emanuel", "Salvador", "Terry", "Morgan", "Jeffery", "Esteban", "Tyson", "Braxton", "Branden", "Marvin", "Brody", "Craig", "Ismael", "Rodney", "Isiah", "Marshall", "Maurice", "Ernesto", "Emilio", "Brendon", "Kody", "Eddie", "Malachi", "Abel", "Keaton", "Jon", "Shaun", "Skylar", "Ezekiel", "Nikolas", "Santiago", "Kendall", "Axel", "Camden", "Trevon", "Bobby", "Conor", "Jamal", "Lukas", "Malcolm", "Zackery", "Jayson", "Javon", "Roger", "Reginald", "Zachariah", "Desmond", "Felix", "Johnathon", "Dean", "Quinton", "Ali", "Davis", "Gerald", "Rodrigo", "Demetrius", "Billy", "Rene", "Reece", "Kelvin", "Leo", "Justice", "Chris", "Guillermo", "Matthew", "Matthew", "Matthew", "Kevon", "Steve", "Frederick", "Clay", "Weston", "Dorian", "Hugo", "Roy", "Orlando", "Terrance", "Kai", "Khalil", "Khalil", "Khalil", "Graham", "Noel", "Willie", "Nathanael", "Terrell", "Tyrone",
	}
	if length > len(names) {
		length = len(names)
	}
	return names[:length]
}

func main() {
	// Maps
	// maps in go is like dictionary in python
	// assosiate key with value and key represents value
	// like array index:value structure
	// but with stirng as index key:value
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Marry"] = 24
	ages["Marry"] = 21 // overwrites 24 and key 'Marry' value

	declaredAges := map[string]int{
		"John":  37,
		"Marry": 21,
	}

	fmt.Println("map: ", declaredAges)
	fmt.Println("length map: ", len(declaredAges))

	fmt.Println("---")

	// Map Example 1
	testMap1(
		[]string{"John", "Bob", "Jill"},
		[]int{14443331111, 12223334455, 21113330010},
	)
	testMap1(
		[]string{"John", "Bob"},
		[]int{14443331111, 12223334455, 21113330010},
	)

	fmt.Println("---")

	// Mutations

	// INSERT AND ELEMENT
	// map[key] = elem

	// GET AN ELEMENT
	// elem = map[key]

	// DELETE AND ELEMENT
	// delete(map, key)

	// CHECK IF A KEY EXISTS
	// elem, ok := m[key]
	// if 'key' is in 'map', then 'ok' is 'true'
	// if not 'ok' is 'false'
	// 'ok' variable is boolean
	// if 'key' is not in 'map', then 'elem' is zero value for map's element type

	// Map Example 2
	usersMap := map[string]userSchedule{
		"John": {
			user:                 user{name: "john", number: 14445556666},
			scheduledForDeletion: true,
		},
		"Alice": {
			user: user{
				name:   "Alice",
				number: 15556667777,
			},
			scheduledForDeletion: false,
		},
		"Bob": {
			user: user{
				name:   "Bob",
				number: 16667778888,
			},
		},
	}
	// John can be deleted
	testMap2(usersMap, "John")

	// Must is not found
	testMap2(usersMap, "Musk")

	// Alice cannot be deleted
	testMap2(usersMap, "Alice")

	// Bob's scheduledForDeletion bool field
	// doesn't exists
	testMap2(usersMap, "Bob")

	fmt.Println("---")

	// slices, maps and functions are not be compared

	// Nestad maps
	//nestedMap := make(map[string]map[string]int)
	//n := hits["/doc/"]["au"]
	//func add(m map[string]map[string]int, path, country string) {
	//	mm, ok := m[path]
	//	if !ok {
	//		mm = make(map[string]int)
	//		m[path] = mm
	//	}
	//	mm[country]++
	//}
	//add(hits, "/doc/", "au")

	// instead of nested maps
	// you can use structs that comparable
	//type Key struct{ Path, Country string }
	//hits := make(map[Key]int)
	//hits[Key{"/", "vn"}]++

	// Count Instances
	// Map Example 3
	var userIDs []string = generateRandomUserIDs()

	testMap3(userIDs, []string{"00", "ff", "dd"})
	testMap3(userIDs, []string{"aa", "12", "32"})
	testMap3(userIDs, []string{"bb", "33"})

	fmt.Println("---")

	// Map Literals
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"MST": -7 * 60 * 60,
	}
	fmt.Println(timeZone)
	//var secondsValue int
	//var ok bool
	//secondsValue, ok = timeZone[tz]
	// 'ok' returns boolean that indicates whether key exists
	// and 'secondsValue' is setted 0 if doesn't exists

	fmt.Println("---")

	// Map Example 4
	// nested maps
	// rune data type is used
	// for represent character rather that whole string
	// like char data type in c#
	testMap4(getNames(50), 'M', "Matthew")
	testMap4(getNames(100), 'G', "George")
	testMap4(getNames(150), 'D', "Drew")
	testMap4(getNames(200), 'P', "Philip")
	testMap4(getNames(250), 'B', "Bryant")
	testMap4(getNames(300), 'M', "Matthew")
}
