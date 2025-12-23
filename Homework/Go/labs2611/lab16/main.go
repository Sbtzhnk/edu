// for reverse engineering
// Создайте базовую структуру для персонажа из полей имени, уровня, опыта, класса
// Создайте структуру для хранения уникальных атрибутов Воина, Мага, Лучника с включением базовой структуры.
// Определите общие методы работы с персонажами в интерфейсе
// Реализуйте методы нанесения повреждений, воcстанновления здоровья, супер удара.
// Организуйте бой двух персонажей с нанесением удара и вычислением оставшегося здоровья.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Базовая структура персонажа
type Character struct {
	Name       string
	Level      int
	Experience int
	Health     int
	MaxHealth  int
	Attack     int
	Defense    int
	Class      string
}

// Интерфейс для общих методов персонажей
type CharacterInterface interface {
	AttackTarget(target CharacterInterface) int
	TakeDamage(damage int)
	Heal(amount int)
	SpecialAttack(target CharacterInterface) int
	IsAlive() bool
	GetHealth() int
	GetMaxHealth() int
	GetInfo() string
}

// Структура Воина
type Warrior struct {
	Character
	Rage int // Уникальный атрибут для воина
}

// Структура Мага
type Mage struct {
	Character
	Mana int // Уникальный атрибут для мага
}

// Структура Лучника
type Archer struct {
	Character
	Arrows int // Уникальный атрибут для лучника
}

// Конструкторы для создания персонажей
func NewWarrior(name string, level int) *Warrior {
	return &Warrior{
		Character: Character{
			Name:      name,
			Level:     level,
			Health:    120 + level*10,
			MaxHealth: 120 + level*10,
			Attack:    15 + level*2,
			Defense:   12 + level,
			Class:     "Warrior",
		},
		Rage: 100,
	}
}

func NewMage(name string, level int) *Mage {
	return &Mage{
		Character: Character{
			Name:      name,
			Level:     level,
			Health:    80 + level*5,
			MaxHealth: 80 + level*5,
			Attack:    20 + level*3,
			Defense:   8 + level,
			Class:     "Mage",
		},
		Mana: 150 + level*10,
	}
}

func NewArcher(name string, level int) *Archer {
	return &Archer{
		Character: Character{
			Name:      name,
			Level:     level,
			Health:    90 + level*7,
			MaxHealth: 90 + level*7,
			Attack:    18 + level*2,
			Defense:   10 + level,
			Class:     "Archer",
		},
		Arrows: 20 + level*2,
	}
}

// Реализация методов для Warrior
func (w *Warrior) AttackTarget(target CharacterInterface) int {
	damage := w.Attack + rand.Intn(10)
	target.TakeDamage(damage)
	w.Rage += 10
	fmt.Printf("%s наносит удар мечом на %d урона!\n", w.Name, damage)
	return damage
}

func (w *Warrior) TakeDamage(damage int) {
	actualDamage := damage - w.Defense/2
	if actualDamage < 1 {
		actualDamage = 1
	}
	w.Health -= actualDamage
	fmt.Printf("%s получает %d урона. Здоровье: %d/%d\n",
		w.Name, actualDamage, w.Health, w.MaxHealth)
}

func (w *Warrior) Heal(amount int) {
	w.Health += amount
	if w.Health > w.MaxHealth {
		w.Health = w.MaxHealth
	}
	fmt.Printf("%s восстанавливает %d здоровья. Здоровье: %d/%d\n",
		w.Name, amount, w.Health, w.MaxHealth)
}

func (w *Warrior) SpecialAttack(target CharacterInterface) int {
	if w.Rage >= 30 {
		damage := w.Attack*2 + rand.Intn(20)
		target.TakeDamage(damage)
		w.Rage -= 30
		fmt.Printf("%s использует СМЕРТЕЛЬНЫЙ УДАР на %d урона!\n", w.Name, damage)
		return damage
	}
	fmt.Printf("%s не хватает ярости для супер удара!\n", w.Name)
	return 0
}

func (w *Warrior) IsAlive() bool {
	return w.Health > 0
}

func (w *Warrior) GetHealth() int {
	return w.Health
}

func (w *Warrior) GetMaxHealth() int {
	return w.MaxHealth
}

func (w *Warrior) GetInfo() string {
	return fmt.Sprintf("[%s] %s (Ур.%d) HP: %d/%d | Атака: %d | Защита: %d | Ярость: %d",
		w.Class, w.Name, w.Level, w.Health, w.MaxHealth, w.Attack, w.Defense, w.Rage)
}

// Реализация методов для Mage
func (m *Mage) AttackTarget(target CharacterInterface) int {
	if m.Mana >= 10 {
		damage := m.Attack + rand.Intn(15)
		target.TakeDamage(damage)
		m.Mana -= 10
		fmt.Printf("%s бросает огненный шар на %d урона!\n", m.Name, damage)
		return damage
	}
	// Если мана закончилась, слабая атака
	damage := m.Attack/2 + rand.Intn(5)
	target.TakeDamage(damage)
	fmt.Printf("%s атакует посохом на %d урона!\n", m.Name, damage)
	return damage
}

func (m *Mage) TakeDamage(damage int) {
	actualDamage := damage - m.Defense/3
	if actualDamage < 1 {
		actualDamage = 1
	}
	m.Health -= actualDamage
	fmt.Printf("%s получает %d урона. Здоровье: %d/%d\n",
		m.Name, actualDamage, m.Health, m.MaxHealth)
}

func (m *Mage) Heal(amount int) {
	if m.Mana >= 20 {
		m.Health += amount
		if m.Health > m.MaxHealth {
			m.Health = m.MaxHealth
		}
		m.Mana -= 20
		fmt.Printf("%s использует исцеление на %d здоровья. Здоровье: %d/%d\n",
			m.Name, amount, m.Health, m.MaxHealth)
	} else {
		fmt.Printf("%s не хватает маны для исцеления!\n", m.Name)
	}
}

func (m *Mage) SpecialAttack(target CharacterInterface) int {
	if m.Mana >= 50 {
		damage := m.Attack*3 + rand.Intn(30)
		target.TakeDamage(damage)
		m.Mana -= 50
		fmt.Printf("%s призывает МОЛНИЮ на %d урона!\n", m.Name, damage)
		return damage
	}
	fmt.Printf("%s не хватает маны для супер заклинания!\n", m.Name)
	return 0
}

func (m *Mage) IsAlive() bool {
	return m.Health > 0
}

func (m *Mage) GetHealth() int {
	return m.Health
}

func (m *Mage) GetMaxHealth() int {
	return m.MaxHealth
}

func (m *Mage) GetInfo() string {
	return fmt.Sprintf("[%s] %s (Ур.%d) HP: %d/%d | Атака: %d | Защита: %d | Мана: %d",
		m.Class, m.Name, m.Level, m.Health, m.MaxHealth, m.Attack, m.Defense, m.Mana)
}

// Реализация методов для Archer
func (a *Archer) AttackTarget(target CharacterInterface) int {
	if a.Arrows > 0 {
		damage := a.Attack + rand.Intn(12)
		target.TakeDamage(damage)
		a.Arrows--
		fmt.Printf("%s выпускает стрелу на %d урона! Осталось стрел: %d\n",
			a.Name, damage, a.Arrows)
		return damage
	}
	// Если стрелы закончились, атака кинжалом
	damage := a.Attack/3 + rand.Intn(3)
	target.TakeDamage(damage)
	fmt.Printf("%s атакует кинжалом на %d урона!\n", a.Name, damage)
	return damage
}

func (a *Archer) TakeDamage(damage int) {
	actualDamage := damage - a.Defense/4
	if actualDamage < 1 {
		actualDamage = 1
	}
	a.Health -= actualDamage
	fmt.Printf("%s получает %d урона. Здоровье: %d/%d\n",
		a.Name, actualDamage, a.Health, a.MaxHealth)
}

func (a *Archer) Heal(amount int) {
	a.Health += amount
	if a.Health > a.MaxHealth {
		a.Health = a.MaxHealth
	}
	fmt.Printf("%s использует бинты на %d здоровья. Здоровье: %d/%d\n",
		a.Name, amount, a.Health, a.MaxHealth)
}

func (a *Archer) SpecialAttack(target CharacterInterface) int {
	if a.Arrows >= 3 {
		damage := a.Attack*2 + rand.Intn(25)
		target.TakeDamage(damage)
		a.Arrows -= 3
		fmt.Printf("%s использует ШТОРМ СТРЕЛ на %d урона!\n", a.Name, damage)
		return damage
	}
	fmt.Printf("%s не хватает стрел для супер атаки!\n", a.Name)
	return 0
}

func (a *Archer) IsAlive() bool {
	return a.Health > 0
}

func (a *Archer) GetHealth() int {
	return a.Health
}

func (a *Archer) GetMaxHealth() int {
	return a.MaxHealth
}

func (a *Archer) GetInfo() string {
	return fmt.Sprintf("[%s] %s (Ур.%d) HP: %d/%d | Атака: %d | Защита: %d | Стрелы: %d",
		a.Class, a.Name, a.Level, a.Health, a.MaxHealth, a.Attack, a.Defense, a.Arrows)
}

// Функция для проведения боя
func Battle(player1, player2 CharacterInterface) {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== НАЧАЛО БОЯ ===")
	fmt.Println(player1.GetInfo())
	fmt.Println(player2.GetInfo())
	fmt.Println()

	round := 1
	for player1.IsAlive() && player2.IsAlive() && round <= 20 {
		fmt.Printf("--- Раунд %d ---\n", round)

		// Первый игрок атакует
		if player1.IsAlive() {
			// 30% шанс использовать супер атаку
			if rand.Intn(100) < 30 {
				player1.SpecialAttack(player2)
			} else {
				player1.AttackTarget(player2)
			}
		}

		// Второй игрок атакует
		if player2.IsAlive() {
			// 30% шанс использовать супер атаку
			if rand.Intn(100) < 30 {
				player2.SpecialAttack(player1)
			} else {
				player2.AttackTarget(player1)
			}
		}

		// 20% шанс исцеления
		if player1.IsAlive() && rand.Intn(100) < 20 {
			player1.Heal(15 + rand.Intn(20))
		}
		if player2.IsAlive() && rand.Intn(100) < 20 {
			player2.Heal(15 + rand.Intn(20))
		}

		fmt.Println()
		round++
	}

	// Определяем победителя
	fmt.Println("=== РЕЗУЛЬТАТ БОЯ ===")
	if player1.IsAlive() && !player2.IsAlive() {
		fmt.Printf("ПОБЕДИТЕЛЬ: %s\n", player1.GetInfo())
	} else if !player1.IsAlive() && player2.IsAlive() {
		fmt.Printf("ПОБЕДИТЕЛЬ: %s\n", player2.GetInfo())
	} else if player1.IsAlive() && player2.IsAlive() {
		fmt.Println("НИЧЬЯ! Оба персонажа выжили.")
	} else {
		fmt.Println("ОБА ПЕРСОНАЖА ПОГИБЛИ!")
	}
}

func main() {
	// Создаем персонажей
	warrior := NewWarrior("Гром", 5)
	mage := NewMage("Мерлин", 5)
	archer := NewArcher("Леголас", 5)

	// Демонстрация информации о персонажах
	fmt.Println("=== ПЕРСОНАЖИ ===")
	fmt.Println(warrior.GetInfo())
	fmt.Println(mage.GetInfo())
	fmt.Println(archer.GetInfo())
	fmt.Println()

	// Проводим бои
	Battle(warrior, mage)
	fmt.Println()

	// Создаем новых персонажей для следующего боя
	warrior2 := NewWarrior("Артур", 3)
	archer2 := NewArcher("Робин", 4)

	Battle(warrior2, archer2)
}
