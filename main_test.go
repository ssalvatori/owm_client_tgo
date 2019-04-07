package main

import "testing"

func Test_getCity(t *testing.T) {
	type args struct {
		cities []CitiesMapKeyValue
		args   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "city found",
			args: args{
				cities: []CitiesMapKeyValue{
					CitiesMapKeyValue{
						Key:   "scl",
						Value: "1234",
					},
					CitiesMapKeyValue{
						Key:   "ccp",
						Value: "7890",
					},
				},
				args: "ccp",
			},
			want: "7890",
		},
		{
			name: "city not found",
			args: args{
				cities: []CitiesMapKeyValue{
					CitiesMapKeyValue{
						Key:   "scl",
						Value: "1234",
					},
					CitiesMapKeyValue{
						Key:   "ccp",
						Value: "7890",
					},
				},
				args: "tll",
			},
			want: "1234",
		},
		{
			name: "no cities",
			args: args{
				cities: []CitiesMapKeyValue{},
				args:   "tll",
			},
			want: "3873544",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCity(tt.args.cities, tt.args.args); got != tt.want {
				t.Errorf("getCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildURL(t *testing.T) {

	got := buildURL("api12", "23", "es")
	want := "https://api.openweathermap.org/data/2.5/weather?id=23&appid=api12&units=metric&lang=es"

	if got != want {
		t.Errorf("buildURL() = %v, want %v", got, want)
	}
}
