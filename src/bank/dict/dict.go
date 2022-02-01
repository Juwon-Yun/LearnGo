package dict

import "errors"

var (
	errNotFound  = errors.New("not Found")
	errWordExist = errors.New("that word already exist")
	errCanUpdate = errors.New("can't update non-existing word")
)

// type에 대한 alias를 정해준다
// struct 뿐만아니라 type에서도 method를 가 수 있다.
// Dictionary type
type Dictionary map[string]string

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	// 위의 switch와 똑같이 사용 가능하다
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExist
	// }
	return nil
}

// 기존 key값을 업데이트하는 메소드
// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		// key가 존재하는 경우
		d[word] = definition
	case errNotFound:
		return errCanUpdate
	}
	return nil
}

// go map의 내장 메소드 delete()
// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
