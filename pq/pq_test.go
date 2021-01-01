package pq

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	charset2 = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	MINPQ_N = 1000
)

var seededRand2 = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand2.Intn(len(charset))]
	}
	return string(b)
}

func TestMinPQ(t *testing.T) {
	minpq := NewHashedPQ(MinPQ, 100)

	assert.Equal(t, true, minpq.IsEmpty())
	assert.Equal(t, 0, minpq.Size())

	/* Insert "key1" priority 10 */
	_, err := minpq.Put("key1", 10)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 1, minpq.Size())

	priority, err := minpq.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)

	/* Insert "key2" priority 11 */
	_, err = minpq.Put("key2", 11)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 2, minpq.Size())

	priority, err = minpq.Get("key2")
	assert.NoError(t, err)
	assert.Equal(t, int64(11), priority)

	key, priority, err := minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	/* Insert "key3" priotity 19 */
	_, err = minpq.Put("key3", 19)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 3, minpq.Size())

	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	/* Insert "key4" priority 9 */
	_, err = minpq.Put("key4", 9)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 4, minpq.Size())

	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(9), priority)
	assert.Equal(t, "key4", key)

	/* Now update "key4" priority 9 to 15 */
	_, err = minpq.Put("key4", 15)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 4, minpq.Size())

	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	/* Insert "key5" priority 21 */
	_, err = minpq.Put("key5", 21)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 5, minpq.Size())

	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	/* Update "key1" priority 10 to 8 */
	_, err = minpq.Put("key1", 8)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 5, minpq.Size())

	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(8), priority)
	assert.Equal(t, "key1", key)

}

func TestMinPQ2(t *testing.T) {
	N := 1000
	minpq := NewHashedPQ(MinPQ, 100)

	keys := make(map[string]bool)
	for i := 1; i <= N; i++ {
		key := randomString(32, charset2)
		keys[key] = true
		contained, err := minpq.Put(key, int64(i+2000))
		assert.NoError(t, err)
		assert.Equal(t, false, minpq.IsEmpty())
		assert.Equal(t, i, minpq.Size())
		assert.Equal(t, false, contained)
	}

	/* Insert "key1" priority 10 */
	contained, err := minpq.Put("key1", 10)
	assert.NoError(t, err)
	assert.Equal(t, false, contained)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, N+1, minpq.Size())

	key, priority, err := minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	i := 0
	for key, _ := range keys {
		minKey, priority, err := minpq.First()
		assert.NoError(t, err)
		assert.Equal(t, int64(10), priority)
		assert.Equal(t, "key1", minKey)

		err = minpq.Delete(key)
		i++
		assert.NoError(t, err)
		assert.Equal(t, N-i+1, minpq.Size())
	}

}

func TestMinPQ3(t *testing.T) {
	N := 1000
	minpq := NewHashedPQ(MinPQ, 100)

	keys := make(map[string]bool)
	/* Insert N keys priority 2001..2000+N */
	for i := 1; i <= N; i++ {
		key := randomString(32, charset2)
		keys[key] = true
		contained, err := minpq.Put(key, int64(i+2000))
		assert.NoError(t, err)
		assert.Equal(t, false, minpq.IsEmpty())
		assert.Equal(t, i, minpq.Size())
		assert.Equal(t, false, contained)
	}

	/* Insert "key1" priority 10 */
	contained, err := minpq.Put("key1", 10)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, N+1, minpq.Size())
	assert.Equal(t, false, contained)

	key, priority, err := minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	/* Update "key1" priority to 100 */
	contained, err = minpq.Put("key1", 100)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, N+1, minpq.Size())
	assert.Equal(t, true, contained)

	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), priority)
	assert.Equal(t, "key1", key)

	i := 0
	for key, _ := range keys {
		minKey, priority, err := minpq.First()
		assert.NoError(t, err)
		assert.Equal(t, int64(100), priority)
		assert.Equal(t, "key1", minKey)

		err = minpq.Delete(key)
		i++
		assert.NoError(t, err)
		assert.Equal(t, N-i+1, minpq.Size())
	}

	/* At this point minpq has only one element, which is "key1" */
	assert.Equal(t, 1, minpq.Size())
	minKey, priority, err := minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), priority)
	assert.Equal(t, "key1", minKey)

	/* Update "key1" priority to 2000 */
	contained, err = minpq.Put("key1", 2000)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 1, minpq.Size())
	assert.Equal(t, true, contained)

	keys2 := make(map[int]string)
	for i := 1; i <= N; i++ {
		key := randomString(32, charset2)
		keys2[i] = key
		contained, err := minpq.Put(key, int64(i))
		assert.NoError(t, err)
		assert.Equal(t, false, minpq.IsEmpty())
		assert.Equal(t, i+1, minpq.Size())
		assert.Equal(t, false, contained)
	}

	/* minpq has to have N+1 elements */
	assert.Equal(t, N+1, minpq.Size())

	for i := 1; i <= N; i++ {
		key2 := keys2[i]
		key1, priority1, err := minpq.First()
		assert.NoError(t, err)
		assert.Equal(t, int64(i), priority1)
		assert.Equal(t, key2, key1)

		key, priority, err := minpq.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, int64(i), priority)
		assert.Equal(t, key2, key)
		assert.Equal(t, N-i+1, minpq.Size())
	}

	assert.Equal(t, 1, minpq.Size())
	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, "key1", key)
	assert.Equal(t, int64(2000), priority)
}

func TestMinPQ4(t *testing.T) {
	N := 3
	minpq := NewHashedPQ(MinPQ, N)

	assert.Equal(t, true, minpq.IsEmpty())
	assert.Equal(t, 0, minpq.Size())

	/* Insert "key1" priority 10 */
	_, err := minpq.Put("key1", 10)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 1, minpq.Size())

	priority, err := minpq.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)

	/* Insert "key2" priority 11 */
	_, err = minpq.Put("key2", 11)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 2, minpq.Size())

	priority, err = minpq.Get("key2")
	assert.NoError(t, err)
	assert.Equal(t, int64(11), priority)

	key, priority, err := minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	/* Insert "key3" priotity 19 */
	_, err = minpq.Put("key3", 19)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 3, minpq.Size())

	key, priority, err = minpq.First()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)

	/* Insert "key4" priority 9 */
	_, err = minpq.Put("key4", 9)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 4, minpq.Size())

	key, priority, err = minpq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, int64(9), priority)
	assert.Equal(t, "key4", key)
	assert.Equal(t, 3, minpq.Size())

	key, priority, err = minpq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), priority)
	assert.Equal(t, "key1", key)
	assert.Equal(t, 2, minpq.Size())

	key, priority, err = minpq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, int64(11), priority)
	assert.Equal(t, "key2", key)
	assert.Equal(t, 1, minpq.Size())

	key, priority, err = minpq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, int64(19), priority)
	assert.Equal(t, "key3", key)
	assert.Equal(t, 0, minpq.Size())
}

func TestIntMinPQ(t *testing.T) {
	minpq := NewHashedPQ(MinPQ, 100)

	assert.Equal(t, true, minpq.IsEmpty())
	assert.Equal(t, 0, minpq.Size())

	/* Insert key=10 priority 20 */
	_, err := minpq.Put(10, 20)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 1, minpq.Size())

	priority, err := minpq.Get(10)
	assert.NoError(t, err)
	assert.Equal(t, int64(20), priority)

	/* Insert key=11 priority 22 */
	_, err = minpq.Put(11, 22)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 2, minpq.Size())

	priority, err = minpq.Get(11)
	assert.NoError(t, err)
	assert.Equal(t, int64(22), priority)

}

func TestDeleteMinPQ(t *testing.T) {
	minpq := NewHashedPQ(MinPQ, 100)

	assert.Equal(t, true, minpq.IsEmpty())
	assert.Equal(t, 0, minpq.Size())

	/* Insert key=10 priority 20 */
	_, err := minpq.Put(10, 20)
	assert.NoError(t, err)
	assert.Equal(t, false, minpq.IsEmpty())
	assert.Equal(t, 1, minpq.Size())

	err = minpq.Delete(10)
	assert.NoError(t, err)

	err = minpq.Delete(10)
	assert.Error(t, err)


}