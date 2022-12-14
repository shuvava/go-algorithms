package linkedlist_test

import (
	"fmt"
	"github.com/shuvava/go-algorithms/pkg/list"
	"strconv"
	"strings"
	"testing"

	"github.com/shuvava/go-algorithms/pkg/linkedlist"
)

func TestDoubleLinkedList(t *testing.T) {
	t.Run("New DoubleLinkedList should return 0 Len", func(t *testing.T) {
		dll := linkedlist.NewDoubleLinkedList[int]()
		l := dll.Len()
		if l != 0 {
			t.Errorf("Length of empty list should be 0")
		}
	})
	t.Run("add should increase size of DoubleLinkedList", func(t *testing.T) {
		a := []int{1, 3, 5, 6}
		dll := linkedlist.NewDoubleLinkedList[int]()
		for k := range a {
			dll.Add(k)
		}
		if len(a) != dll.Len() {
			t.Errorf("Length of DoubleLinkedList is not equal to test array: len(arr)=%d len(dll)=%d", len(a), dll.Len())
		}
	})
	t.Run("DoubleLinkedList be iterable", func(t *testing.T) {
		a := []int{1, 3, 5, 6}
		dll := linkedlist.NewDoubleLinkedList[int]()
		for _, k := range a {
			dll.Add(k)
		}
		it := dll.Iterator()
		i := 0
		for it.HasNext() {
			k := it.Next()
			if a[i] != k {
				t.Errorf("DoubleLinkedList on index %d snot equal to the input %d != %d", i, a[i], k)
			}
			i++
		}
		if i != len(a) {
			t.Errorf("DoubleLinkedList itterator number of elements is not equal to the input %d != %d", i, len(a))
		}
	})
	t.Run("New DoubleLinkedList should return false for HasNext", func(t *testing.T) {
		dll := linkedlist.NewDoubleLinkedList[int]()
		iter := dll.Iterator()
		if iter.HasNext() {
			t.Errorf("HasNext should return false")
		}
	})
	t.Run("Del method of DoubleLinkedList should decrease size of the list", func(t *testing.T) {
		a := []int{1, 3, 5, 6}
		check := make([]*linkedlist.Item[int], 0)
		dll := linkedlist.NewDoubleLinkedList[int]()
		for k := range a {
			item := dll.Add(k)
			check = append(check, item)
		}
		l := len(a) - 1
		var item *linkedlist.Item[int]
		for i := 0; i < l; i++ {
			m := len(check) >> 1
			item, check = list.Pop(check, m)
			dll.Del(item)
			if dll.Len() != l {
				t.Errorf("Del method break DoubleLinkedList")
			}
			l--
		}
	})
	t.Run("Get should return items preserve order of adding", func(t *testing.T) {
		a := []int{1, 3, 5, 6}
		check := make([]*linkedlist.Item[int], 0)
		dll := linkedlist.NewDoubleLinkedList[int]()
		for _, k := range a {
			item := dll.Add(k)
			check = append(check, item)
		}
		l := len(a) - 1
		first, found1 := dll.Get(0)
		if !found1 || first == nil {
			t.Errorf("the first item should be found")
		}
		if first.Value != a[0] {
			t.Errorf("Get return incorrect the first item")
		}
		last, found2 := dll.Get(l)
		if !found2 || last == nil {
			t.Errorf("the first item should be found")
		}
		if last.Value != a[l] {
			t.Errorf("Get return incorrect the last item")
		}
	})
}

func ExampleNewDoubleLinkedList() {
	dll := linkedlist.NewDoubleLinkedList[int]()
	for _, v := range []int{1, 3, 5, 6} {
		dll.Add(v) // add a new val
	}
	iter := dll.Iterator() //get DoubleLinkedList iterator
	result := make([]string, 0)
	for iter.HasNext() {
		v := iter.Next()
		result = append(result, strconv.FormatInt(int64(v), 10))
	}
	fmt.Println(strings.Join(result[:], ", "))
	// Output: 1, 3, 5, 6
}
