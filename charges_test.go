package openpay

import (
	"reflect"
	"testing"
)

func Test_chargesClient_AddCard(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		card *Card
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
				c: testClient(Card{}, 200),
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			if err := cc.AddCard(tt.args.card); (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.AddCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_chargesClient_Get(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		txID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
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
				txID: "test_txid",
			},
			want: &Transaction{
				ID:     "test_txid",
				Status: "test_status",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				txID: "test_txid",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			got, err := cc.Get(tt.args.txID)
			if (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargesClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chargesClient_List(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		req *ChargesListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Transaction
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				c: testClient([]Transaction{{
					ID:     "test_txid",
					Status: "test_status",
				}}, 200),
			},
			args: args{
				req: &ChargesListRequest{},
			},
			want: []Transaction{{
				ID:     "test_txid",
				Status: "test_status",
			}},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				req: &ChargesListRequest{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			got, err := cc.List(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargesClient.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chargesClient_AtStore(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		charge *ChargeAtStore
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
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
				charge: &ChargeAtStore{},
			},
			want: &Transaction{
				ID:     "test_txid",
				Status: "test_status",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				charge: &ChargeAtStore{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			got, err := cc.AtStore(tt.args.charge)
			if (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.AtStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargesClient.AtStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chargesClient_AtBank(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		charge *ChargeAtBank
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
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
				charge: &ChargeAtBank{},
			},
			want: &Transaction{
				ID:     "test_txid",
				Status: "test_status",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				charge: &ChargeAtBank{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			got, err := cc.AtBank(tt.args.charge)
			if (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.AtBank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargesClient.AtBank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chargesClient_WithCard(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		charge *ChargeWithStoredCard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
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
				charge: &ChargeWithStoredCard{},
			},
			want: &Transaction{
				ID:     "test_txid",
				Status: "test_status",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				charge: &ChargeWithStoredCard{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			got, err := cc.WithCard(tt.args.charge)
			if (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.WithCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargesClient.WithCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chargesClient_Capture(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		txID   string
		amount float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
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
				txID:   "test_txid",
				amount: 100,
			},
			want: &Transaction{
				ID:     "test_txid",
				Status: "test_status",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				txID:   "test_id",
				amount: 100,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			got, err := cc.Capture(tt.args.txID, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.Capture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargesClient.Capture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chargesClient_Refund(t *testing.T) {
	type fields struct {
		c *Client
	}
	type args struct {
		txID        string
		amount      float32
		description string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Transaction
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
				txID:        "test_txid",
				amount:      100,
				description: "test_description",
			},
			want: &Transaction{
				ID:     "test_txid",
				Status: "test_status",
			},
			wantErr: false,
		},
		{
			name: "API error",
			fields: fields{
				c: testClient(nil, 400),
			},
			args: args{
				txID:        "test_id",
				amount:      100,
				description: "test_description",
			},
			want:    nil,
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &chargesClient{
				c: tt.fields.c,
			}
			got, err := cc.Refund(tt.args.txID, tt.args.amount, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("chargesClient.Refund() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargesClient.Refund() = %v, want %v", got, tt.want)
			}
		})
	}
}
