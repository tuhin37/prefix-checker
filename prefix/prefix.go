package prefix

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/tuhin37/truecaller-prefix/prefix/util"
)

var wORDLIST_FOLDER string = "./prefix-wordlists/"
var prefixHashMap = make(map[string]string)

func init() {
	// This function finds all the wordlist files present in `prefix-wordlists` folder and loads them in a hash map

	var wg sync.WaitGroup
	// list all the files in `prefix-wordlists` folder
	files, err := ioutil.ReadDir(wORDLIST_FOLDER)
	if err != nil {
		log.Fatal(err)
	}

	// select valid files from the folder
	for _, file := range files {
		// only use files which have `.txt` extainsion and has some data
		if !file.IsDir() && file.Name()[len(file.Name())-3:] == "txt" && file.Size() > 0 {
			wg.Add(1)
			// for each file, call the function // TODO: fork a new go routine for each file
			go func(fname string) {
				defer wg.Done()
				file2hashmap(fname)
			}(file.Name())
		}
	}
	wg.Wait() // wait for the goroutines to finish before moving on
}

func file2hashmap(filename string) {
	// load all workds from one file into a map
	f, err := os.Open(wORDLIST_FOLDER + filename)
	// report if file open operation fails
	if err != nil {
		log.Fatal(err.Error())
	}

	// close file handler once done matter what
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	// iterate over each work / line
	for fileScanner.Scan() {
		// hash a word and insert it to the map
		hash := util.GetMD5(fileScanner.Text())
		// collision case detection
		oldString, isFound := prefixHashMap[hash]
		if isFound {
			// append two values in comma seperated format
			prefixHashMap[hash] = util.AppendCSV(oldString, fileScanner.Text())
			continue
		}
		prefixHashMap[hash] = fileScanner.Text()
	}
}

// ------------------------------------------------------ PUBLIC FUNCTIONS ----------------------------------------------------------
func CheckPrefix(input string) string {

	// edge case
	if input == "" {
		return ""
	}

	type MatchedPrefix struct {
		Length int
		Value  string
	}
	OutputChannel := make(chan MatchedPrefix)

	// This function takes an input string, then it will return the longest prefix that matches with the begining of the input string
	var output string
	// cumulatively add characters from the begining of the input string, the word that has formed is then hashed, and the mak is checked against the hash key,
	// if a value is found then return that.
	for i := range input {
		cumulativeSubString := input[:i] // if the input is"helloworld" then this variable grows to be "h", "he", "hel",... "helloworld" with each iterations

		go func() {
			defer close(OutputChannel)
			hashSubString := util.GetMD5(cumulativeSubString)
			valueStored, isFound := prefixHashMap[hashSubString]
			if isFound {
				var mp MatchedPrefix
				mp.Length = len(hashSubString)
				mp.Value = valueStored
				OutputChannel <- mp
			}
		}()
	}

	//

	// receive and append channel values
	var maxLength int
	for matchedPrefix := range OutputChannel {
		if matchedPrefix.Length > maxLength {
			output = matchedPrefix.Value
		}
	}

	// if output is csv (due to md5 collision), then find out which value (of all the csv values) is the actual prefix
	values, isCSV := util.Csv2List(output)
	if isCSV {
		for _, value := range values {
			if util.IsSubstring(value, input) {
				output = value
			}
		}
	}
	return output
}
