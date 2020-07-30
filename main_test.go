package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_genCleveland(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		{
			name: "empty",
			args: func(t *testing.T) args {
				return args{
					line: "",
				}
			},
			want1: "",
		},
		{
			name: "1",
			args: func(t *testing.T) args {
				return args{
					line: "1",
				}
			},
			want1: "1,1@my.clevelandcc.edu,1999CCC#1",
		},
		{
			name: "123",
			args: func(t *testing.T) args {
				return args{
					line: "123",
				}
			},
			want1: "123,123@my.clevelandcc.edu,1999CCC#123",
		},
		{
			name: "1234",
			args: func(t *testing.T) args {
				return args{
					line: "1234",
				}
			},
			want1: "1234,1234@my.clevelandcc.edu,1999CCC#234",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := genCleveland(tArgs.line)

			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_genCentral(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		{
			name: "empty",
			args: func(t *testing.T) args {
				return args{
					line: "",
				}
			},
			want1: "",
		},
		{
			name: "Thomas Jones,1249108",
			args: func(t *testing.T) args {
				return args{
					line: "Thomas Jones,1249108",
				}
			},
			want1: "Thomas Jones,1249108,TJone108,TJone108@cougarmail.cccc.edu",
		},
		{
			name: "Thomas Jones,08",
			args: func(t *testing.T) args {
				return args{
					line: "Thomas Jones,08",
				}
			},
			want1: "Thomas Jones,08,TJone08,TJone08@cougarmail.cccc.edu",
		},
		{
			name: "Thomas Jon,108",
			args: func(t *testing.T) args {
				return args{
					line: "Thomas Jon,108",
				}
			},
			want1: "Thomas Jon,108,TJon108,TJon108@cougarmail.cccc.edu",
		},
		{
			name: "Thomas Jone,108",
			args: func(t *testing.T) args {
				return args{
					line: "Thomas Jone,108",
				}
			},
			want1: "Thomas Jone,108,TJone108,TJone108@cougarmail.cccc.edu",
		},
		{
			name: "Thomas,108",
			args: func(t *testing.T) args {
				return args{
					line: "Thomas,108",
				}
			},
			want1: "",
		},
		{
			name: "Thomas",
			args: func(t *testing.T) args {
				return args{
					line: "Thomas",
				}
			},
			want1: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := genCentral(tArgs.line)

			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_genEdgecombe(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		{
			name: "empty",
			args: func(t *testing.T) args {
				return args{
					line: "",
				}
			},
			want1: "",
		},
		{
			name: "tjones",
			args: func(t *testing.T) args {
				return args{
					line: "tjones",
				}
			},
			want1: "",
		},
		{
			name: "tjones0303,471-23-5016",
			args: func(t *testing.T) args {
				return args{
					line: "tjones0303,471-23-5016",
				}
			},
			want1: "tjones0303,471-23-5016,tjones0303@myedge.cc,235016",
		},
		{
			name: "tjones0303,471235016",
			args: func(t *testing.T) args {
				return args{
					line: "tjones0303,471235016",
				}
			},
			want1: "tjones0303,471235016,tjones0303@myedge.cc,235016",
		},
		{
			name: "tjones0303,3-5016",
			args: func(t *testing.T) args {
				return args{
					line: "tjones0303,3-5016",
				}
			},
			want1: "tjones0303,3-5016,tjones0303@myedge.cc,35016",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := genEdgecombe(tArgs.line)

			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_genMayland(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		{
			name: "empty",
			args: func(t *testing.T) args {
				return args{
					line: "",
				}
			},
			want1: "",
		},
		{
			name: "tjones3137,Thomas Jones,03/03/1999",
			args: func(t *testing.T) args {
				return args{
					line: "tjones3137,Thomas Jones,03/03/1999",
				}
			},
			want1: "tjones3137,Thomas Jones,03/03/1999,tjones3137@students.mayland.edu,Jo030399",
		},
		{
			name: "tjones3137,Thomas Jones",
			args: func(t *testing.T) args {
				return args{
					line: "tjones3137,Thomas Jones",
				}
			},
			want1: "",
		},
		{
			name: "tjones3137,Thomas Jones,03/03/99",
			args: func(t *testing.T) args {
				return args{
					line: "tjones3137,Thomas Jones,03/03/99",
				}
			},
			want1: "tjones3137,Thomas Jones,03/03/99,tjones3137@students.mayland.edu,Jo030399",
		},
		{
			name: "tjones3137,Thomas Jo,03/03/99",
			args: func(t *testing.T) args {
				return args{
					line: "tjones3137,Thomas Jo,03/03/99",
				}
			},
			want1: "tjones3137,Thomas Jo,03/03/99,tjones3137@students.mayland.edu,Jo030399",
		},
		{
			name: "tjones3137,Thomas J,03/03/9",
			args: func(t *testing.T) args {
				return args{
					line: "tjones3137,Thomas J,03/03/9",
				}
			},
			want1: "tjones3137,Thomas J,03/03/9,tjones3137@students.mayland.edu,J03039",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := genMayland(tArgs.line)

			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_outputFilename(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		{
			name: "",
			args: func(t *testing.T) args {
				return args{
					filename: "",
				}
			},
			want1: "_out.csv",
		},
		{
			name: "1",
			args: func(t *testing.T) args {
				return args{
					filename: "1",
				}
			},
			want1: "1_out.csv",
		},
		{
			name: "1.txt",
			args: func(t *testing.T) args {
				return args{
					filename: "1.txt",
				}
			},
			want1: "1_out.csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := outputFilename(tArgs.filename)

			assert.Equal(t, tt.want1, got1)
		})
	}
}
