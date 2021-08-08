package main

import "fmt"

// Car Adapter Pattern Exercise

// client: CarFeatures
// service: DriveService
// unknown: FlyService
// Adaptor: FlyDriveServiceAdaptor

type CarFeatures struct {}
type DriveService struct {}
type FlyService struct {}
type FlyDriveServiceAdaptor struct {
	flyService *FlyService
}

type DriveFeature interface {
	drive()
}

func (cf *CarFeatures) useFeatures(df DriveFeature) {
	df.drive()
}

func (ds *DriveService) drive() {
	fmt.Println("Car is Driving and using wheels")
}

func (fs *FlyService) fly() {
	fmt.Println("is now flying")
}

func (fds *FlyDriveServiceAdaptor) drive() {
	fmt.Println("Car has now spreading it wings")
	fds.flyService.fly()
}

func main() {
	cf := &CarFeatures{}
	ds := &DriveService{}
	fs := &FlyService{}
	fds := &FlyDriveServiceAdaptor{fs}

	cf.useFeatures(ds)
	// cf.useFeatures(fs)
	fs.fly()
	fmt.Println("________")
	cf.useFeatures(fds)
}