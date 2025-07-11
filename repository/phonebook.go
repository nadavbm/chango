package repository

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/nadavbm/chango/decorator"
	"github.com/nadavbm/chango/strategy"
)

const phoneNumberLegth = 10

type PhoneBook struct {
	logger       decorator.Logger
	contacts     map[string][]int
	lock         *sync.RWMutex
	intGenerator *strategy.Int
	strGenerator *strategy.Str
}

func newPhoneBook(logger decorator.Logger) *PhoneBook {
	return &PhoneBook{
		logger:       logger,
		contacts:     make(map[string][]int),
		lock:         &sync.RWMutex{},
		intGenerator: &strategy.Int{},
		strGenerator: &strategy.Str{},
	}
}

func (p *PhoneBook) addContact() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.contacts[p.generateContactName()] = p.generatePhoneNumber()
}

func (p *PhoneBook) updateContact(contactName string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.contacts[contactName] = p.generatePhoneNumber()
}

func (p *PhoneBook) deleteContact(contactName string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.contacts, contactName)
}

func (p *PhoneBook) showAllContacts() {
	p.lock.Lock()
	defer p.lock.Unlock()
	fmt.Println("------------------------------------------------")
	for k, v := range p.contacts {
		fmt.Printf("%s - %s\n", k, intSliceToString(v))
	}
}

func (p *PhoneBook) generatePhoneNumber() []int {
	digits := []int{}
	for i := 0; i < phoneNumberLegth; i++ {
		if i < 2 {
			digits = append(digits, i)
		} else {
			digits = append(digits, p.intGenerator.Integer(9, 0))
		}
	}
	return digits
}

func (p *PhoneBook) generateContactName() string {
	surname := p.strGenerator.String(p.intGenerator.Integer(6, 2))
	familyName := p.strGenerator.String(p.intGenerator.Integer(12, 3))
	surname = fmt.Sprintf("%s%s", strings.ToUpper(surname[0:1]), surname[1:])
	familyName = fmt.Sprintf("%s%s", strings.ToUpper(familyName[0:1]), familyName[1:])
	return fmt.Sprintf("%s, %s", familyName, surname)
}

func intSliceToString(numbers []int) string {
	str := []string{}
	for i, n := range numbers {
		if i == 2 || i == 6 {
			str = append(str, strconv.Itoa(n)+"-")
		} else {
			str = append(str, strconv.Itoa(n))
		}
	}
	return strings.Join(str, "")
}
