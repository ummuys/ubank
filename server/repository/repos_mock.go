package repository

type mockDb struct {
}

func NewMockDb() DataBase {
	return &mockDb{}
}

func (m *mockDb) Close() {}

func (m *mockDb) CheckExistsUser(login string) (bool, error) {
	return false, nil
}

func (m *mockDb) GetPassword(login string) (string, error) {
	return "", nil
}
func (m *mockDb) CreateUser(login, password string) error {
	return nil
}

func (m *mockDb) Deposite(login string, cash int64) error {
	return nil
}

func (m *mockDb) GetUserBalans(login string) (int64, error) {
	return 0, nil
}

func (m *mockDb) TransferMoney(loginFrom, loginTo string, cash int64) error {
	return nil
}
