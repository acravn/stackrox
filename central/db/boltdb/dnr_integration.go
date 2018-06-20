package boltdb

import (
	"errors"
	"fmt"
	"time"

	"bitbucket.org/stack-rox/apollo/central/db"
	"bitbucket.org/stack-rox/apollo/central/metrics"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/uuid"
	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
)

const dnrIntegrationBucket = "dnrintegration"

// GetDNRIntegration retrieves a DNR integration from Bolt.
func (b *BoltDB) GetDNRIntegration(id string) (integration *v1.DNRIntegration, exists bool, err error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Get", "DNRIntegration")
	err = b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dnrIntegrationBucket))
		key := []byte(id)
		bytes := b.Get(key)
		if bytes == nil {
			return nil
		}
		exists = true
		integration = new(v1.DNRIntegration)
		err := proto.Unmarshal(bytes, integration)
		if err != nil {
			return fmt.Errorf("proto unmarshalling: %s", err)
		}
		return nil
	})
	if err != nil {
		err = fmt.Errorf("DNR integration retrieval: %s", err)
	}
	return
}

// GetDNRIntegrations retrieves all D&R integrations from bolt
func (b *BoltDB) GetDNRIntegrations(req *v1.GetDNRIntegrationsRequest) (integrations []*v1.DNRIntegration, err error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "GetMany", "DNRIntegration")
	err = b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dnrIntegrationBucket))
		return b.ForEach(func(k, v []byte) error {
			var integration v1.DNRIntegration
			if err := proto.Unmarshal(v, &integration); err != nil {
				return fmt.Errorf("proto unmarshalling: %s", err)
			}
			// If a cluster id is provided, then only return integrations that match it.
			if req.GetClusterId() != "" && integration.GetClusterId() != req.GetClusterId() {
				return nil
			}
			integrations = append(integrations, &integration)
			return nil
		})
	})
	if err != nil {
		err = fmt.Errorf("DNR integration retrieval: %s", err)
	}
	return
}

// AddDNRIntegration adds a DNR integration to Bolt.
func (b *BoltDB) AddDNRIntegration(integration *v1.DNRIntegration) (string, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Add", "DNRIntegration")
	id := uuid.NewV4().String()
	integration.Id = id
	err := b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dnrIntegrationBucket))
		key := []byte(id)
		bytes, err := proto.Marshal(integration)
		if err != nil {
			return fmt.Errorf("proto marshalling: %s", err)
		}
		return b.Put(key, bytes)
	})
	if err != nil {
		return "", fmt.Errorf("DNR integration insertion: %s", err)
	}
	return id, nil
}

// UpdateDNRIntegration updates the DNR integration in Bolt.
func (b *BoltDB) UpdateDNRIntegration(integration *v1.DNRIntegration) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Update", "DNRIntegration")
	if integration.GetId() == "" {
		return errors.New("cannot update; empty id provided")
	}
	return b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dnrIntegrationBucket))
		key := []byte(integration.GetId())
		if b.Get(key) == nil {
			return db.ErrNotFound{Type: "DNRIntegration"}
		}
		bytes, err := proto.Marshal(integration)
		if err != nil {
			return fmt.Errorf("DNR integration proto marshalling: %s", err)
		}
		return b.Put(key, bytes)
	})
}

// RemoveDNRIntegration removes the DNR integration from Bolt.
func (b *BoltDB) RemoveDNRIntegration(id string) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Remove", "DNRIntegration")
	return b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dnrIntegrationBucket))
		key := []byte(id)
		if b.Get(key) == nil {
			return db.ErrNotFound{Type: "DNRIntegration"}
		}
		return b.Delete(key)
	})
}
