package common

import (
	"bufio"
	"empresa/settings"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetLastID(key string) (int, error) {
	file, err := os.Open(settings.IDTrackerFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 && parts[0] == key {
			return strconv.Atoi(parts[1])
		}
	}
	return 0, nil
}

func UpdateLastID(key string, newID int) error {
	ids := make(map[string]int)

	file, _ := os.Open(settings.IDTrackerFilename)
	if file != nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				val, _ := strconv.Atoi(parts[1])
				ids[parts[0]] = val
			}
		}
		file.Close()
	}

	ids[key] = newID

	fileW, err := os.Create(settings.IDTrackerFilename)
	if err != nil {
		return err
	}
	defer fileW.Close()

	for k, v := range ids {
		fmt.Fprintf(fileW, "%s:%d\n", k, v)
	}
	return nil
}
