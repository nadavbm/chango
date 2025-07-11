package repository

import (
	"time"

	"github.com/nadavbm/chango/decorator"
	"github.com/nadavbm/chango/singleton"
)

const initialContactNumber = 11

func DynamicPhoneBook(logger decorator.Logger, cfg *singleton.Config) {
	pb := newPhoneBook(logger)
	pb.initContacts()
	logger.Info("New phone book created")
	logger.Info("----------------------")
	pb.showAllContacts()
	go func() {
		for {
			time.Sleep(time.Duration(pb.intGenerator.Integer(8, 7)) * time.Second)
			logger.Info("Current phone book print")
			logger.Info("------------------------")
			pb.showAllContacts()
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(pb.intGenerator.Integer(4, 3)) * time.Second)
			pb.addContact()
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(pb.intGenerator.Integer(2, 1)) * time.Second)
			pb.updateContact(pb.getRandomContactName())
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(pb.intGenerator.Integer(6, 5)) * time.Second)
			pb.deleteContact(pb.getRandomContactName())
		}
	}()
	time.Sleep(cfg.Duration)
	logger.Info("Last phone book update")
	logger.Info("----------------------")
	pb.showAllContacts()
}

func (p *PhoneBook) initContacts() {
	for i := 0; i < initialContactNumber; i++ {
		p.addContact()
	}
}

func (p *PhoneBook) getRandomContactName() string {
	p.lock.Lock()
	defer p.lock.Unlock()
	for k := range p.contacts {
		return k
	}
	return ""
}
