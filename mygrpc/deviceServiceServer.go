package mygrpc

import (
	"context"
	"errors"
	"github.com/onezerobinary/db-box/repository"
	pb_mobile "github.com/onezerobinary/db-box/proto/device"
)

type DeviceServiceServer struct {

}

func (s *DeviceServiceServer) AddDevice (ctx context.Context, device *pb_mobile.Device) (*pb_mobile.Response, error) {

	response, _ := repository.AddDevice(device)

	return &response, nil
}

// Get Device
func (s *DeviceServiceServer) GetDeviceByExpoToken (ctx context.Context, expoPushToken *pb_mobile.ExpoPushToken) (*pb_mobile.Device, error) {

	device, err := repository.GetDeviceByExpoToken(expoPushToken)

	if err != nil {
		errorText := "It was not possible to get the Device and a Fake one is return " + err.Error()
		newError := errors.New(errorText)
		return &device, newError
	}

	return &device, nil
}

// Update Device Status
func (s *DeviceServiceServer) UpdateStatus (ctx context.Context, status *pb_mobile.Status) (*pb_mobile.Response, error) {

	response, _ := repository.UpdateStatus(status)

	return &response, nil
}

// Update Device Position
func (s *DeviceServiceServer) UpdatePosition (ctx context.Context, position *pb_mobile.Position) (*pb_mobile.Response, error) {

	response, _ := repository.UpdatePosition(position)

	return &response, nil
}

// Update Device MobileNumber
func (s *DeviceServiceServer) UpdateMobileNumber (ctx context.Context, mobileNumber *pb_mobile.MobileNumber) (*pb_mobile.Response, error) {

	response, _ := repository.UpdateMobileNumber(mobileNumber)

	return &response, nil
}