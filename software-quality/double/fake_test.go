package double

import "testing"

type FakeSearcher struct{}

func(fs FakeSearcher) Search(people []*Person, firstname, lastname string) *Person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReurnsEmptyStringForNoPerson(t *testing.T) {
	phonebook := &Phonebook{}
	fake := &FakeSearcher{}

	phone, _ := phonebook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Wanted '', got '%s'", phone)
	}
}