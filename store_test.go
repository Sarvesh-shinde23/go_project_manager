package main

type MockStore struct{}

// GetUserByID implements Store.
func (m *MockStore) GetUserByID(id string) (*User, error) {
	panic("unimplemented")
}

func (m *MockStore) CreateUser() error {
	return nil
}
func (m *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}
func (m *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}