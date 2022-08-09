package cloudstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io/ioutil"
	"time"
)

type getObject struct {
	bucket string
}

func NewGetObjectRepository(bucket string) *getObject {
	return &getObject{
		bucket: bucket,
	}
}

func (o *getObject) GetObject(ctx context.Context, object string) ([]byte, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()

	rc, err := client.Bucket(o.bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	return data, nil
}
