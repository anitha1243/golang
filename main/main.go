package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	compID     string
	userID     string
	appID      string
	systemType string
}

var users []int64
var assessedUsers []int64

func main() {

	records, err := readData("sample-large.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		user := User{
			compID:     record[0],
			userID:     record[1],
			appID:      record[2],
			systemType: record[3],
		}

		userID, err := strconv.ParseInt(user.userID, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		appID, err := strconv.ParseInt(user.appID, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if appID == 374 {

			userIDExists := isValueInList(userID, users)
			userIDAssessed := isValueInList(userID, assessedUsers)
			if userIDExists == true && userIDAssessed == false {

				for _, record := range records {
					userIDExisting, err := strconv.ParseInt(record[1], 10, 64)
					if err != nil {
						log.Fatal(err)
					}
					if userIDExisting == userID {
						var compIDs []int64
						var sysTypes []string
						compIDExistingUser, err := strconv.ParseInt(record[0], 10, 64)
						if err != nil {
							log.Fatal(err)
						}
						compIDs = append(compIDs, compIDExistingUser)
						sysTypes = append(sysTypes, strings.ToLower(record[3]))
						var laptopCount int
						var desktopCount int
						compIDs = removeDuplicates(compIDs)

						for _, record := range records {
							compIDExistingUser, err := strconv.ParseInt(record[0], 10, 64)
							if err != nil {
								log.Fatal(err)
							}
							for _, compID := range compIDs {
								if compID == compIDExistingUser {
									if record[3] == "laptop" {
										laptopCount = laptopCount + 1
									}
									if record[3] == "desktop" {
										desktopCount = desktopCount + 1
									}
								}
							}
						}
						var totalSystemTypeCount int
						if laptopCount == desktopCount {
							totalSystemTypeCount = laptopCount
						} else {
							if laptopCount > desktopCount {
								totalSystemTypeCount = (laptopCount - desktopCount) + ((laptopCount + desktopCount) / 2)
							} else {
								totalSystemTypeCount = (desktopCount - laptopCount) + ((laptopCount + desktopCount) / 2)
							}
						}

						finalCountExistingUser := totalSystemTypeCount - 1
						for i := 1; i < finalCountExistingUser; i++ {
							users = append(users, userID)
						}
						assessedUsers = append(assessedUsers, userID)
					}
				}

				continue
			}
			users = append(users, userID)
		}
	}

	fmt.Printf("Minimum number of copies for application 374 is %v ", len(users))
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func isValueInList(value int64, list []int64) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func removeDuplicates(elements []int64) []int64 {
	// Use map to record duplicates as we find them.
	encountered := map[int64]bool{}
	result := []int64{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
