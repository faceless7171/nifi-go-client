package nifi

import (
	"context"
	"os"
	"reflect"
	"testing"
)

var flowApi *FlowApiService

func TestMain(m *testing.M) {
	cfg := NewConfiguration()
	cfg.Host = "localhost:8080"
	cfg.Scheme = "http"
	flowApi = NewAPIClient(cfg).FlowApi
	os.Exit(m.Run())
}

func TestFlowApiService_GetAboutInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		a       FlowApiService
		args    args
		want    AboutEntity
		wantErr bool
	}{
		{
			name: "get about info",
			a:    *flowApi,
			args: args{ctx: context.Background()},
			want: AboutEntity{About: AboutDto{
				Title:            "NiFi",
				Version:          "1.11.4",
				Uri:              "http://localhost:8080/nifi-api/",
				ContentViewerUrl: "../nifi-content-viewer/",
				Timezone:         "UTC",
				BuildTag:         "nifi-1.11.4-RC1",
				BuildTimestamp:   "03/18/2020 12:42:54 UTC",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.a.GetAboutInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAboutInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAboutInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}