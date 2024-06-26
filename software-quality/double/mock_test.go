package double

import "testing"

type MockSearcher struct {
	phone string
	methodsToCall map[string]bool
}

func(ms *MockSearcher) Search(people []*Person, firstname, lastname string) *Person {
	ms.methodsToCall["Search"] = true
	return &Person {
		Firstname: firstname,
		Lastname: lastname,
		Phone: ms.phone,
	}
}

func(ms *MockSearcher) Create(people []*Person, firstname, lastname string) *Person {
	ms.methodsToCall["Create"] = true
	return &Person {
		Firstname: firstname,
		Lastname: lastname,
		Phone: ms.phone,
	}
}

func(ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}
	ms.methodsToCall[methodName] = false
}

func(ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReurnsPersonUsingMock(t *testing.T) {
	fakephone := "0123456789"
	Phonebook := &Phonebook{}
	mock := &MockSearcher{phone : fakephone}
	mock.ExpectToCall("Search")

	phone, _ := Phonebook.Find(mock, "Jane", "Doe")

	if phone != fakephone {
		t.Errorf("Want '%s', got '%s'", fakephone, phone)
	}

	mock.Verify(t)
} 