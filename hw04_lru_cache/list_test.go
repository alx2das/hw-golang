package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.NotNil(t, l)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("get_len", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)
		l.PushFront(20)

		require.Equal(t, 2, l.Len())
	})
}

func TestList_Add(t *testing.T) {
	t.Run("push front", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)

		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 1, l.Len())

		l.PushFront(20)

		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 2, l.Len())
	})

	t.Run("push back", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)

		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 1, l.Len())

		l.PushBack(20)

		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)
		require.Equal(t, 2, l.Len())
	})
}

func TestList_Remove(t *testing.T) {
	t.Run("remove nil", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]

		l.Remove(nil) // [10, 20]

		require.Equal(t, 2, l.Len())
	})

	t.Run("remove one", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)     // [10]
		l.Remove(l.Front()) // []

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("remove front", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)     // [10]
		l.Remove(l.Front()) // []

		require.Equal(t, 0, l.Len())

		for _, v := range [...]int{10, 20, 30, 40} {
			l.PushFront(v) // [40, 30, 20, 10]
		}
		l.Remove(l.Front()) // [30, 20, 10]

		require.Equal(t, 3, l.Len())
		require.Equal(t, 30, l.Front().Value)
	})

	t.Run("remove back", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)     // [10]
		l.Remove(l.Back()) // []

		require.Equal(t, 0, l.Len())

		for _, v := range [...]int{10, 20, 30, 40} {
			l.PushBack(v) // [10, 20, 30, 40]
		}
		l.Remove(l.Back()) // [10, 20, 30]

		require.Equal(t, 3, l.Len())
		require.Equal(t, 30, l.Back().Value)
	})
}

func TestList_MoveToFront(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		l := NewList()

		l.MoveToFront(nil)

		require.Equal(t, 0, l.Len())
	})

	t.Run("one item", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)
		l.MoveToFront(l.Back())

		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, l.Front().Value)
	})

	t.Run("center", func(t *testing.T) {
		l := NewList()

		for _, v := range [...]int{10, 20, 30} {
			l.PushBack(v) // [10, 20, 30, 40]
		}

		require.Equal(t, 3, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)

		l.MoveToFront(l.Back().Prev)

		require.Equal(t, 3, l.Len())
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)
	})

	t.Run("front", func(t *testing.T) {
		l := NewList()

		for _, v := range [...]int{10, 20, 30} {
			l.PushBack(v) // [10, 20, 30]
		}

		l.MoveToFront(l.Front())

		require.Equal(t, 3, l.Len())
		require.Equal(t, 10, l.Front().Value)
	})

	t.Run("back", func(t *testing.T) {
		l := NewList()

		for _, v := range [...]int{10, 20, 30} {
			l.PushBack(v) // [10, 20, 30]
		}

		l.MoveToFront(l.Back())

		require.Equal(t, 3, l.Len())
		require.Equal(t, 30, l.Front().Value)
	})
}
