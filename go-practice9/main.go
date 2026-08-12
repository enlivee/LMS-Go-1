package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
)

type Habit struct {
	ID int
	Title string
	TargetDays int
	CompletedDays int
	IsFinished bool
}

type HabitManager struct {
	habits map[int]*Habit
	nextID int
}

func NewHabitManager() *HabitManager {
	return &HabitManager{
		habits: make(map[int]*Habit),
		nextID: 1,
	}
}

func (h *HabitManager) AddHabit(title string, targetDays int) {
	newHabit := &Habit{
		ID: h.nextID,
		Title: title,
		TargetDays: targetDays,
		CompletedDays: 0,
		IsFinished: false,
	}
	h.habits[newHabit.ID] = newHabit
	h.nextID++
	fmt.Printf("Привычка '%s' успешна добавлена с ID %d\n", title, newHabit.ID)
}

func (h *HabitManager) MarkCompleted(id int) {
	habit, exists := h.habits[id]
	if !exists {
		fmt.Printf("Привычка с ID %d не найдена\n", id)
		return
	}
	habit.CompletedDays++
	if habit.CompletedDays >= habit.TargetDays {
		habit.IsFinished = true
		fmt.Printf("Поздравляем! Вы завершили привычку '%s'!\n", h.habits[id].Title)
	} else {
		fmt.Printf("Отмечено! Прогресс: %d/%d дней\n", habit.CompletedDays, habit.TargetDays)
	}
}

func (h *HabitManager) ListHabits() {
	if len(h.habits) == 0{
		fmt.Println("Нет привычек для отображения.")
	}
	for id, habit := range h.habits {
		if habit.IsFinished {
			s := "Завершено"
			fmt.Printf("[ID: %d] %s | Прогресс: %d/%d дней | Статус: %s\n", id, habit.Title, habit.CompletedDays, habit.TargetDays, s)
		} else {
			s := "В процессе"
			fmt.Printf("[ID: %d] %s | Прогресс: %d/%d дней | Статус: %s\n", id, habit.Title, habit.CompletedDays, habit.TargetDays, s)
		}
	}
}

func main() {
	manager := NewHabitManager()
	for {
		fmt.Print("Введите команду (add, complete, list, exit): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := scanner.Text()
		switch command {
		case "add":
			var title string
			var targetDays int
			fmt.Print("Введите название привычки: ")
			fmt.Scanf("%s", &title)
			fmt.Print("Введите количество дней для завершения: ")
			fmt.Scanf("%d", &targetDays)
			manager.AddHabit(title, targetDays)
		case "complete":
			var id int
			fmt.Print("Введите ID привычки для отметки выполнения: ")
			fmt.Scanf("%d", &id)
			manager.MarkCompleted(id)
		case "list":
			manager.ListHabits()
		case "exit":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная команда. Попробуйте снова.")
		}
	}
}