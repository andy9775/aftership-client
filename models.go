package aftership

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"time"
)

// ================================================ responses ================================================

// Checkpoint includes information about the shipments stop off location
// A checkpoint occurs when a shipment gets scanned into a location (warehouse or other)
type Checkpoint struct {
	Slug           string        `json:"slug"`
	City           interface{}   `json:"city"`
	CreatedAt      time.Time     `json:"created_at"`
	Location       interface{}   `json:"location"`
	CountryName    interface{}   `json:"country_name"`
	Message        string        `json:"message"`
	CountryIso3    interface{}   `json:"country_iso3"`
	Tag            string        `json:"tag"`
	Subtag         string        `json:"subtag"`
	SubtagMessage  string        `json:"subtag_message"`
	CheckpointTime string        `json:"checkpoint_time"`
	Coordinates    []interface{} `json:"coordinates"`
	State          interface{}   `json:"state"`
	Zip            interface{}   `json:"zip"`
}

// Tracking contains the packages tracking information
type Tracking struct {
	ID                            string        `json:"id"`
	CreatedAt                     time.Time     `json:"created_at"`
	UpdatedAt                     time.Time     `json:"updated_at"`
	LastUpdatedAt                 time.Time     `json:"last_updated_at"`
	TrackingNumber                string        `json:"tracking_number"`
	Slug                          string        `json:"slug"`
	Active                        bool          `json:"active"`
	Android                       []interface{} `json:"android"`
	CustomFields                  interface{}   `json:"custom_fields"`
	CustomerName                  interface{}   `json:"customer_name"`
	DeliveryTime                  int           `json:"delivery_time"`
	DestinationCountryIso3        string        `json:"destination_country_iso3"`
	CourierDestinationCountryIso3 string        `json:"courier_destination_country_iso3"`
	Emails                        []interface{} `json:"emails"`
	ExpectedDelivery              interface{}   `json:"expected_delivery"`
	Ios                           []interface{} `json:"ios"`
	Note                          interface{}   `json:"note"`
	OrderID                       interface{}   `json:"order_id"`
	OrderIDPath                   interface{}   `json:"order_id_path"`
	OriginCountryIso3             string        `json:"origin_country_iso3"`
	ShipmentPackageCount          int           `json:"shipment_package_count"`
	ShipmentPickupDate            string        `json:"shipment_pickup_date"`
	ShipmentDeliveryDate          string        `json:"shipment_delivery_date"`
	ShipmentType                  string        `json:"shipment_type"`
	ShipmentWeight                int           `json:"shipment_weight"`
	ShipmentWeightUnit            string        `json:"shipment_weight_unit"`
	SignedBy                      string        `json:"signed_by"`
	Smses                         []interface{} `json:"smses"`
	Source                        string        `json:"source"`
	Tag                           string        `json:"tag"`
	Subtag                        string        `json:"subtag"`
	SubtagMessage                 string        `json:"subtag_message"`
	Title                         string        `json:"title"`
	TrackedCount                  int           `json:"tracked_count"`
	LastMileTrackingSupported     interface{}   `json:"last_mile_tracking_supported"`
	Language                      interface{}   `json:"language"`
	UniqueToken                   string        `json:"unique_token"`
	Checkpoints                   []Checkpoint  `json:"checkpoints"`
	SubscribedSmses               []interface{} `json:"subscribed_smses"`
	SubscribedEmails              []interface{} `json:"subscribed_emails"`
	ReturnToSender                bool          `json:"return_to_sender"`
	TrackingAccountNumber         interface{}   `json:"tracking_account_number"`
	TrackingOriginCountry         interface{}   `json:"tracking_origin_country"`
	TrackingDestinationCountry    interface{}   `json:"tracking_destination_country"`
	TrackingKey                   interface{}   `json:"tracking_key"`
	TrackingPostalCode            interface{}   `json:"tracking_postal_code"`
	TrackingShipDate              interface{}   `json:"tracking_ship_date"`
	TrackingState                 interface{}   `json:"tracking_state"`
}

// TrackingData is the response body which includes tracking information
type TrackingData struct {
	Tracking `json:"tracking"`
}

// TrackingMeta includes metadata about the tracking information response
type TrackingMeta struct {
	Code int `json:"code"`
}

// TrackingResponse is the body returned from aftership containing tracking information
type TrackingResponse struct {
	Meta TrackingMeta `json:"meta"`
	Data TrackingData `json:"data"`
}

// ================================================= requests ================================================

// toJSON convers the tracking request into a valid json response
func (r *NewTrackingRequest) toJSON() (io.Reader, error) {
	if r.TrackingNumber == "" {
		return nil, errors.New("TrackingNumber cannot be empty")
	}

	resp, err := json.Marshal(struct {
		Tracking interface{} `json:"tracking"`
	}{Tracking: r})
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(resp), nil
}

// NewTrackingRequest is the request body for creating a new tracking record
type NewTrackingRequest struct {
	Slug string `json:"slug"`
	// TrackingNumber is required
	TrackingNumber string   `json:"tracking_number"`
	Title          string   `json:"title"`
	Smses          []string `json:"smses"`
	Emails         []string `json:"emails"`
	OrderID        string   `json:"order_id"`
	OrderIDPath    string   `json:"order_id_path"`
	CustomFields   struct {
		ProductName  string `json:"product_name"`
		ProductPrice string `json:"product_price"`
	} `json:"custom_fields"`
	Language string `json:"language"`
}
