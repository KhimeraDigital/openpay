package openpay

import (
	"reflect"
	"testing"
)

func Test_webhooksClient_Create(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		wh *Webhook
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				c: testClient(Transaction{
					ID:     "test_txid",
					Status: "test_status",
				}, 200),
			},
			args: args{
				wh: &Webhook{},
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				wh: &Webhook{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &webhooksClient{
				c: tt.fields.c,
			}
			if err := wc.Create(tt.args.wh); (err != nil) != tt.wantErr {
				t.Errorf("webhooksClient.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_webhooksClient_Get(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		whID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Webhook
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				c: testClient(Webhook{
					ID: "test_w_id",
				}, 200),
			},
			args: args{
				whID: "test_whid",
			},
			want: &Webhook{
				ID: "test_w_id",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				whID: "test_whid",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &webhooksClient{
				c: tt.fields.c,
			}
			got, err := wc.Get(tt.args.whID)
			if (err != nil) != tt.wantErr {
				t.Errorf("webhooksClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("webhooksClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_webhooksClient_List(t *testing.T) {
	type fields struct {
		c *Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Webhook
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				c: testClient([]Webhook{{
					ID: "test_w_id",
				}}, 200),
			},
			want: []Webhook{{
				ID: "test_w_id",
			}},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &webhooksClient{
				c: tt.fields.c,
			}
			got, err := wc.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("webhooksClient.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("webhooksClient.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_webhooksClient_Delete(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		whID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				c: testClient(Webhook{
					ID: "test_w_id",
				}, 200),
			},
			args: args{
				whID: "test_whid",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				whID: "test_whid",
			},
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &webhooksClient{
				c: tt.fields.c,
			}
			if err := wc.Delete(tt.args.whID); (err != nil) != tt.wantErr {
				t.Errorf("webhooksClient.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
