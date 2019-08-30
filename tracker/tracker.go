/*

tracker.go

MIT License

Copyright (c) 2018 rezamirz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package tracker

import "fmt"

type Tracker struct {
	size          uint64 /* Number of bits of bitmap in tracker */
	nextLowcontig uint64 /* Lowcontig id that has been tracked */
	startIndex    uint64 /* Start index of the IDs */
	bitmap        []byte
}

// NewTracker creates a tracker
func NewTracker(size uint64, nextLowcontig uint64) *Tracker {

	t := &Tracker{
		size:          size,
		nextLowcontig: nextLowcontig,
		startIndex:    nextLowcontig,
	}

	t.bitmap = make([]byte, size)
	return t
}

// Track tracks an object with specified id.
// IDs start from 0 and can grow indefinitely if the tracker has space.
// Every call to tracker_track() would shift the tracker bitmap if
// a quarter of the tracker has already been tracked. This shifting
// would make space to track new IDs.
func (tracker *Tracker) Track(id uint64) error {

	/* It is already tracked */
	if id < tracker.nextLowcontig {
		return nil
	}

	index := id - tracker.startIndex
	if index > tracker.size {
		return fmt.Errorf("Index out of range")
	}

	/* Now index becomes a byte index inside the bitmap */
	index = index / 8

	tracker.bitmap[index] = tracker.bitmap[index] | 1<<((id-tracker.startIndex)%8)

	if id == tracker.nextLowcontig {
		/* Advance the index until all bits in a byte are set */
		for tracker.bitmap[index] == 0xFF {
			index++
		}

		var bitIndex uint64
		lastByte := tracker.bitmap[index]
		for (1 << bitIndex & lastByte) != 0 {
			bitIndex++
		}

		tracker.nextLowcontig = tracker.startIndex + index*8 + bitIndex
	}

	/* If atleast a quarter of tracker has already been tracked, shift the bitmap.
	 * Exact formula: lowcontig >= size/4
	 */
	offset := tracker.nextLowcontig - tracker.startIndex
	if offset >= tracker.size/4 {
		bitmap := make([]byte, tracker.size/8+1)
		copy(bitmap, tracker.bitmap[offset/8:tracker.size/8])
		tracker.bitmap = bitmap
		tracker.startIndex += (offset / 8) * 8
	}

	return nil
}

// NextLowcontig obtains lowest contiguous id that was tracked.
func (tracker *Tracker) NextLowcontig() uint64 {
	return tracker.nextLowcontig
}

// IsTracked returns true if tracker has already tracked the object with specified id.
func (tracker *Tracker) IsTracked(id uint64) bool {
	/* It is already tracked */
	if id < tracker.nextLowcontig {
		return true
	}

	index := id - tracker.startIndex
	if index > tracker.size {
		return false
	}

	index = index / 8
	if tracker.bitmap[index]&1<<(id%8) != 0 {
		return true
	}

	return false
}

// Next obtains next_id > id that has not been tracked yet.
func (tracker *Tracker) Next(id uint64) (uint64, error) {
	/* Up to _lowcontig, all the IDs have already tracked */
	if id < tracker.nextLowcontig {
		id = tracker.nextLowcontig
	}

	if id >= tracker.size {
		return 0, fmt.Errorf("Out of range ID")
	}

	index := (id - tracker.startIndex) / 8
	for tracker.bitmap[index] == 0xFF {
		index++
	}

	var bitIndex uint64
	nextByte := tracker.bitmap[index]
	for (1 << bitIndex & nextByte) != 0 {
		bitIndex++
	}

	nextID := index*8 + bitIndex
	return nextID, nil
}
