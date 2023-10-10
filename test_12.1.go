// Проверка коллизий с использованием "Цепочки обязанностей" 
// Проверяем правильно ли определяется окрестность для каждого игрового объекта
// Проверяем удаляются ли объекты из старых окрестностей и добавляются ли в новые окрестности
 
package main

import (
  "testing"
)

func TestDetermineNeighborhood(t *testing.T) {
  // создайте игровой объект и определите его окрестность
  obj := &GameObject{ID: 1, Position: [2]float64{5, 5}}
  neighborhood := DetermineNeighborhood(obj) // функция определяет окрестность

  if neighborhood != "ExpectedNeighborhood" {
    t.Errorf("Expected %s but got %s", "ExpectedNeighborhood", neighborhood)
  }
}

func TestSwitchNeighborhood(t *testing.T) {
  // Проверка - перемещается ли объект из одной окрестности в другую
  // Предположим что есть два списка для двух разных окрестностей
  oldNeighborhood := make([]*GameObject, 0)
  newNeighborhood := make([]*GameObject, 0)

  obj := &GameObject{ID: 1, Position: [2]float64{5, 5}}
  oldNeighborhood = append(oldNeighborhood, obj)

  // Функция, которая должна удалять объект из старой и добавлять в новую окрестность
  SwitchNeighborhood(obj, &oldNeighborhood, &newNeighborhood)

  if len(oldNeighborhood) != 0 {
    t.Errorf("Object was not removed from the old neighborhood")
  }

  if len(newNeighborhood) != 1 {
    t.Errorf("Object was not added to the new neighborhood")
  }
}

