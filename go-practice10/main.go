package main

import (
	"sort"
	"errors"
	"fmt"
)

type Company struct {
	workers []*Worker
}

type Worker struct {
	name string
	position string
	salary uint
	experience uint
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error{
	if name == "" || position == "" {
		return errors.New("Имя и позиция не должны быть пустыми")
	}
	w := &Worker{name: name, position: position, salary: salary, experience: experience}
	c.workers = append(c.workers, w)
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	if len(c.workers) == 0 {
		return nil, errors.New("Нет работников для сортировки")
	}
	positionMap := map[string]int{
		"директор" : 1,
		"зам. директора":   2,
		"начальник цеха":   3,
		"мастер":           4,
		"рабочий":          5,
	}
	sort.Slice(c.workers, func(i, j int) bool {
		incomeI := c.workers[i].salary * c.workers[i].experience
    	incomeJ := c.workers[j].salary * c.workers[j].experience
		if incomeI != incomeJ {
			return incomeI > incomeJ
		}
		return positionMap[c.workers[i].position] < positionMap[c.workers[j].position]})
	result := make([]string, len(c.workers))
	for i, w := range c.workers {
		income := w.salary * w.experience
		result[i] = fmt.Sprintf("%s — %d — %s", w.name, income, w.position)
	}
	return result, nil
}