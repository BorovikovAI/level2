package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

type Shop struct {
	Name     string
	Products []Product
}

func NewShop(name string) *Shop {
	return &Shop{Name: name, Products: []Product{}}
}

func (s *Shop) AddProduct(prod Product) {
	s.Products = append(s.Products, prod)
}

func (s *Shop) Sell(user User, prod Product) (err error) {
	fmt.Printf("Shop %s selling %s for user: %s\n", s.Name, prod.Name, user.ID)
	err = user.CheckBalance()
	if err != nil {
		return err
	}
	for _, p := range s.Products {
		if prod != p {
			err = errors.New("Shop: Sell: no such product!")
			continue
		}
		if user.GetBalance() >= p.Price {
			fmt.Printf("Товар %s успешно продан клиенту %s в магазине %s\n", p.Name, user.ID, s.Name)
			err = nil
		} else {
			err = errors.New("Shop: Sell: not enough money to buy this product!")
		}
	}
	return err
}

type Bank struct {
	Name  string
	Cards []Card
}

func NewBank(name string) *Bank {
	return &Bank{Name: name, Cards: []Card{}}
}

func (b *Bank) AddCard(card Card) {
	b.Cards = append(b.Cards, card)
}

func (b *Bank) CheckBalance(card Card) (err error) {
	for _, c := range b.Cards {
		if card != c {
			err = errors.New("Банк - Нет такой карты!")
			continue
		}
		if c.Balance <= 0 {
			err = errors.New("Банк - Недостаточно средств!")
		}
		if c.Balance > 0 {
			fmt.Println("Банк - Баланс больше 0!")
			err = nil
		}
	}
	return err
}

type Card struct {
	Number  int
	Balance float64
	Bank    *Bank
}

func NewCard(number int, balance float64, bank *Bank) *Card {
	return &Card{Number: number, Balance: balance, Bank: bank}
}

func (c *Card) CheckBalance() error {
	fmt.Println("Card: CheckBalance.")
	err := c.Bank.CheckBalance(*c)
	return err
}

type User struct {
	ID   string
	Card *Card
}

func NewUser(id string, card *Card) *User {
	return &User{ID: id, Card: card}
}

func (u *User) CheckBalance() error {
	fmt.Println("User: CheckBalance.")
	err := u.Card.CheckBalance()
	return err
}

func (u *User) GetBalance() float64 {
	return u.Card.Balance
}

func main() {
	shop01 := NewShop("MAGNIT")
	prod01 := NewProduct("cheese", 456.17)
	shop01.AddProduct(*prod01)

	bank01 := NewBank("VTB")
	bank02 := NewBank("SBER")
	card01 := NewCard(12343456, 1234.67, bank01)
	card02 := NewCard(98760987, 347.43, bank02)
	//card03 := NewCard("debit", 46.87, bank02)
	bank01.AddCard(*card01)
	bank02.AddCard(*card02)

	user01 := NewUser("0001", card01)
	user02 := NewUser("0002", card02)
	err := shop01.Sell(*user01, *prod01)
	fmt.Println(err)
	err = shop01.Sell(*user02, *prod01)
	fmt.Println(err)
}
