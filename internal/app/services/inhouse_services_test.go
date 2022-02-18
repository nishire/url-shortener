package services

import (
	"reflect"
	"testing"
	"url-shortener/internal/app/utility"
	"url-shortener/internal/app/utility/cache"
)

func Test_inhouse_FetchShortUrl(t *testing.T) {
	type args struct {
		inputUrl interface{}
	}
	type testData struct {
		name    string
		i       inhouse
		args    args
		want    interface{}
		wantErr bool
	}

	redisKeyGen := utility.NewRedisKeyGenerator()
	redisClient := cache.GetRedisClientImp()
	base62Encoder := utility.NewBase62Encoder()

	tests := []testData{
		testData{
			name: "Test case to be passed",
			i: inhouse{
				redisClient,
				redisKeyGen,
				base62Encoder,
			},
			args:    args{inputUrl: "www.google.com"},
			want:    "/b",
			wantErr: false,
		},
		testData{
			name: "Test case to be passed",
			i: inhouse{
				redisClient,
				redisKeyGen,
				base62Encoder,
			},
			args:    args{inputUrl: "www.xyz.com"},
			want:    "/c",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.FetchShortUrl(tt.args.inputUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("inhouse.FetchShortUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inhouse.FetchShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
