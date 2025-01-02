package addresses

import "strings"

// AddressType verify if address have a valid type and returns
func AddressType(address string) string {
	validTypes := []string{"street", "road", "avenue", "highway"}

	lowerAddress := strings.ToLower(address)

	addressFirstWord := strings.Split(lowerAddress, " ")[0]

	isValidType := false
	for _, validType := range validTypes {
		if validType == addressFirstWord {
			isValidType = true
		}
	}

	if isValidType {
		return strings.Title(addressFirstWord)
	}

	return "Invalid Type!"

}
