package storage

import "sync"

// sync.Mutex полностью блокирует доступ и для чтения и для записи.
// RWMutex позволяет потокам читать данные одновременно, но запись осуществляется только одним потоком

// общее хранилище данных
type Store[K comparable, V any] struct {
	mutex sync.RWMutex // мьютекс для потокобезопасности
	items map[K]V      // хранилище значений по ключу
}

func NewStore[K comparable, V any](capacity int) *Store[K, V] {
	return &Store[K, V]{
		items: make(map[K]V, capacity),
	}
}

// записываем или обновляем значение
func (store *Store[K, V]) Put(key K, value V) {
	store.mutex.Lock() // блокируем доступ на запись
	defer store.mutex.Unlock()

	store.items[key] = value // сохраняем в мапу
}

// получаем значение по ключу
func (store *Store[K, V]) Get(key K) (V, bool) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	value, ok := store.items[key] // ищем элемент
	return value, ok
}

// удаляем ключ из хранилища
func (store *Store[K, V]) Delete(key K, value V) bool {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if _, ok := store.items[key]; !ok { //если ключа нет выходим
		return false
	}

	delete(store.items, key)
	return true
}

// обновляем значение по ключу с помощью функции модификатора
func (store *Store[K, V]) Update(key K, tradeQuantity float64, modif func(V) V) (V, bool) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	value, ok := store.items[key]
	if !ok {
		var zero V
		return zero, false
	}

	updatedValue := modif(value)
	return updatedValue, true
}

// индексированное хранилище с поддержкой порядка добавления элементов
type IndexedStorage[K comparable, V any] struct {
	mutex sync.RWMutex
	order []K
	byKey map[K]int //Позволяет за О(1) узнать порядок ключа в срезе
	data  map[K]V   //Хранилище значений по ключу
}

func NewIndexedStorage[K comparable, V any](capacity int) *IndexedStorage[K, V] {
	return &IndexedStorage[K, V]{
		order: make([]K, 0 , capacity),
		byKey: make(map[K]int, capacity),
		data:  make(map[K]V, capacity),
	}
}

//добавляет или обновляет значение по ключу
func (store *IndexedStorage[K, V]) Put(key K, newValue V) {
	store.mutex.Lock() //блокируем доступ на запись
	defer store.mutex.Unlock() //обещаем разблокировать

	if _,ok := store.data[key]; ok { //если ключ уже есть, то обновляем значение
		store.data[key] = newValue // обновляем значение по ключу
		// fmt.Println("значение изменено с" .value. "на")
		return
	}

	store.byKey[key] = len(store.order) //сохраняем индекс ключа в срезе
	store.order = append(store.order, key) //добавляем ключ в срез
	store.data[key] = newValue //сохраняем значение по ключу
}

func (store *IndexedStorage[K, V]) Delete(key K) bool {
	store.mutex.Lock() //блокируем доступ на запись
	defer store.mutex.Unlock() //обещаем разблокировать

	indexInSlice, ok := store.byKey[key] //если ключа нет выходим
	
	if !ok { 
		return false
	}

	lastIndex := len(store.order) - 1

	if indexInSlice != lastIndex { //если удаляемый элемент не последний, то перемещаем последний элемент в место удаляемого
		lastIndex := store.order[lastIndex] //получаем ключ последнего элемента
		store.order[indexInSlice] = lastIndex //перемещаем последний элемент в место удаляемого
		store.byKey[lastIndex] = indexInSlice //обновляем индекс последнего элемента в мапе
	}

	store.order = store.order[:lastIndex] //перемещаем последний элемент в место удаляемого
	
	delete(store.byKey, key) //удаляем ключ из мапы
	delete(store.data, key) //удаляем значение по ключу

	return true
}

// получаем все значения в порядке их добавления
func (indexStore *IndexedStorage[K, V]) Ordered() []V {
	indexStore.mutex.RLock() //блокируем только на чтение
	defer indexStore.mutex.RUnlock()

	out := make([]V, len(indexStore.order)) //выделяем память под срез
	for index, value := range indexStore.order {
		out[index] = indexStore.data[value] //заполняем по порядку
	}
	return out
}

type Storage[T comparable] struct {
	storage map[T]struct{}
} //хранение универсальных множеств

func NewStorage[T comparable](capacity int) *Storage[T] {
	return &Storage[T]{
		storage: make(map[T]struct{}, capacity),
	}
}

func (s *Storage[T]) Add(item T) {
	s.storage[item] = struct{}{}
}

func (s *Storage[T]) Set(value T) {
	s.storage[value] = struct{}{}
}

func (s *Storage[T]) Remove(value T) {
	delete(s.storage, value)
}

func (s *Storage[T]) Has(value T) bool {
	_, ok := s.storage[value]
	return ok
}

func (s *Storage[T]) Length() int {
	return len(s.storage)
}

func (s *Storage[T]) Items() []T {
	outSlice := make([]T, 0, len(s.storage))
	for value := range s.storage {
		outSlice = append(outSlice, value)
	}
	return outSlice
}
