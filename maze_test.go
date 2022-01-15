package maze

import (
	"testing"
)

func TestFindPath(t *testing.T) {

	tests := []struct {
		matrix             [][]int
		pathLen            int
		err                error
		startcol, startrow int
	}{
		{
			matrix: [][]int{
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 1, 1, 0, 1},
				{1, 0, 0, 1, 0, 0, 1},
				{1, 0, 0, 1, 0, 1, 1},
				{1, 0, 1, 1, 0, 0, 1},
				{1, 0, 0, 1, 1, 0, 1},
				{1, 1, 2, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			pathLen:  7,
			err:      nil,
			startcol: 3,
			startrow: 7,
		},
		{
			matrix: [][]int{
				{1, 1, 1, 1},
				{1, 1, 2, 1},
				{1, 1, 1, 1},
			},
			pathLen:  0,
			err:      ErrNoExit,
			startcol: 3,
			startrow: 2,
		},
		{
			matrix: [][]int{
				{1, 1, 1, 1},
				{1, 1, 2, 1},
				{1, 1, 1, 1},
			},
			pathLen:  0,
			err:      ErrInvalidInput,
			startcol: 2,
			startrow: 1,
		},
	}
	for _, level := range tests {
		mz, err := NewMaze(level.matrix, []int{2, 0})
		if err != nil {
			t.Fatalf("expected error: %v, got: %v", level.err, err)
		} else {
			pathLen, err := mz.FindLenExit(level.startcol, level.startrow)
			//error
			if err != nil {
				if level.err != err {
					t.Fatalf("expected error: %v, got: %v", level.err, err)
				}
				//success
			} else {
				if pathLen != level.pathLen {
					t.Fatalf("expected pathLen: %d, got: %d", level.pathLen, pathLen)
				}
			}
		}

	}
}

func TestNewMaze(t *testing.T) {
	tests := []struct {
		matrix [][]int
		err    error
	}{
		{
			matrix: [][]int{
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 1, 1, 0, 1},
			},
			err: nil,
		},
		{
			matrix: [][]int{
				{1, 1, 1, 1},
				{1, 1, 0},
				{1, 1, 1, 1},
			},
			err: ErrInvalidMatrix,
		},
		{
			matrix: [][]int{},
			err:    ErrInvalidMatrix,
		},
	}
	for _, m := range tests {
		_, err := NewMaze(m.matrix, []int{2, 0})
		if err != m.err {
			t.Fatalf("expected error: %v, got: %v", m.err, err)
		}
	}
}
