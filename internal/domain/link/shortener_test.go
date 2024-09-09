package link

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/hrvadl/link-shortener/internal/domain/link/mocks"
)

func TestNewShortener(t *testing.T) {
	t.Parallel()
	type args struct {
		links LinksSource
	}
	tests := []struct {
		name string
		args args
		want *Shortener
	}{
		{
			name: "Should create new shortener correctly",
			args: args{
				links: mocks.NewMockLinksSource(gomock.NewController(t)),
			},
			want: &Shortener{
				links: mocks.NewMockLinksSource(gomock.NewController(t)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewShortener(tt.args.links); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShortener() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenerShorten(t *testing.T) {
	t.Parallel()
	type fields struct {
		links func(c *gomock.Controller) LinksSource
	}
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should shorten URL correctly",
			fields: fields{
				links: func(c *gomock.Controller) LinksSource {
					ls := mocks.NewMockLinksSource(c)
					ls.EXPECT().
						Set(context.Background(), gomock.Any(), "http://very.long.url/and-its-reeeeeeeeeeeeeally-long").
						Times(1).
						Return(nil)
					return ls
				},
			},
			args: args{
				ctx: context.Background(),
				url: "http://very.long.url/and-its-reeeeeeeeeeeeeally-long",
			},
			wantErr: false,
		},
		{
			name: "Should return an error if saving failed for some reason",
			fields: fields{
				links: func(c *gomock.Controller) LinksSource {
					ls := mocks.NewMockLinksSource(c)
					ls.EXPECT().
						Set(context.Background(), gomock.Any(), "http://very.long.url").
						Times(1).
						Return(errors.New("failed to save"))
					return ls
				},
			},
			args: args{
				ctx: context.Background(),
				url: "http://very.long.url",
			},
			wantErr: true,
		},
		{
			name: "Should validate the empty URL",
			fields: fields{
				links: func(c *gomock.Controller) LinksSource {
					ls := mocks.NewMockLinksSource(c)
					ls.EXPECT().
						Set(context.Background(), gomock.Any(), "").
						Times(0).
						Return(nil)
					return ls
				},
			},
			args: args{
				ctx: context.Background(),
				url: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Shortener{
				links: tt.fields.links(gomock.NewController(t)),
			}

			got, err := e.Shorten(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shortener.Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if len(got) >= len(tt.args.url) {
				t.Errorf("Shortener.Shorten() = got %v, want shorter than %v", got, tt.args.url)
			}
		})
	}
}

func TestShortenerGet(t *testing.T) {
	t.Parallel()
	type fields struct {
		links func(c *gomock.Controller) LinksSource
	}
	type args struct {
		ctx  context.Context
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return an URL if got it successfully",
			fields: fields{
				links: func(c *gomock.Controller) LinksSource {
					ls := mocks.NewMockLinksSource(c)
					ls.EXPECT().
						Get(context.Background(), "hash").
						Times(1).
						Return("http://example.com", nil)
					return ls
				},
			},
			args: args{
				ctx:  context.Background(),
				hash: "hash",
			},
			wantErr: false,
			want:    "http://example.com",
		},
		{
			name: "Should return an error if failed to get",
			fields: fields{
				links: func(c *gomock.Controller) LinksSource {
					ls := mocks.NewMockLinksSource(c)
					ls.EXPECT().
						Get(context.Background(), "hash").
						Times(1).
						Return("", errors.New("failed"))
					return ls
				},
			},
			args: args{
				ctx:  context.Background(),
				hash: "hash",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := &Shortener{
				links: tt.fields.links(gomock.NewController(t)),
			}

			got, err := e.Get(tt.args.ctx, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shortener.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Shortener.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
