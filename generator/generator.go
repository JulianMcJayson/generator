package generator

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

var PASSWORD_LENGTH = 18

func Generate() (string, error) {
	if PASSWORD_LENGTH < 12 {
		return "", errors.New("Password is not secure.")
	}
	rawPassword, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}
	unpackPassword := rawPassword.String()
	limitNumberLength := len(rawPassword) - 1
	countNumber := CountInt(unpackPassword)

	for countNumber > limitNumberLength {
		newRawPassword, err := uuid.NewUUID()
		if err != nil {
			return "", err
		}
		unpackPassword = newRawPassword.String()
		countNumber = CountInt(unpackPassword)
	}

	splitPassword := strings.Split(unpackPassword, "-")
	randMiddle := rand.Intn(2) + 1
	middlePassword := splitPassword[randMiddle]
	groundPasswordSplit := splitPassword[0] + splitPassword[len(splitPassword)-1]
	trackRandom := IntDictionary{}
	trackRandomChannel := make(chan IntDictionary)
	if PASSWORD_LENGTH > len(groundPasswordSplit) {
		PASSWORD_LENGTH = len(groundPasswordSplit)
	}
	groundPassword := ""
	for range PASSWORD_LENGTH - len(middlePassword) {
		randomPickGround := rand.Intn(len(groundPasswordSplit))
		for trackRandom.Contain(randomPickGround) {
			randomPickGround = rand.Intn(len(groundPasswordSplit))
		}
		groundPassword += string(groundPasswordSplit[randomPickGround])
		go func() {
			trackRandomChannel <- trackRandom.Add(randomPickGround)
		}()
		add := <-trackRandomChannel
		trackRandom = add
	}

	groundPassword = RandomUpper(groundPassword)

	middlePlacement := rand.Intn(9)
	if middlePlacement >= 4 {
		middlePlacement = 1
	} else {
		middlePlacement = 0
	}

	arrayPlacement := StringDictionary{groundPassword}
	arrayPlacementChannel := make(chan StringDictionary)
	go func() {
		arrayPlacementChannel <- arrayPlacement.Insert(middlePassword, middlePlacement)
	}()
	insert := <-arrayPlacementChannel
	arrayPlacement = insert
	arrayToStringPassword := ""
	for _, i := range arrayPlacement {
		arrayToStringPassword += i
	}

	trackShuffle := IntDictionary{}
	trackShuffleChannel := make(chan IntDictionary)
	for range len(arrayToStringPassword) / 2 {
		begin := rand.Intn(len(arrayToStringPassword))
		target := rand.Intn(len(arrayToStringPassword))
		for trackShuffle.Contain(target) || begin == target {
			target = rand.Intn(len(arrayToStringPassword))
		}
		for trackShuffle.Contain(begin) || begin == target {
			begin = rand.Intn(len(arrayToStringPassword))
		}

		arrayToStringPassword = Swap(
			arrayToStringPassword,
			string(arrayToStringPassword[begin]),
			string(arrayToStringPassword[target]),
			begin,
			target)

		go func() {
			trackShuffleChannel <- trackShuffle.Add(begin)
			trackShuffleChannel <- trackShuffle.Add(target)
		}()

		add := <-trackShuffleChannel
		trackShuffle = add
	}

	spacialPassword := RandomSpacialChar(arrayToStringPassword)

	password := spacialPassword
	return password, nil
}
