package main

import "fmt"

type Bid struct {
  ImageMetaData *ImageMetaData
  UserMetaData *UserMetaData
  Contract *Contract
}

type Ask struct {
  ClusterMetaData *ClusterMetaData
  UserMetaData *UserMetaData
  Contract *Contract
}

type ImageMetaData struct {
  RegistryUrl string
  RepositoryName string
  ImageName string
  ImageTag string
}

type UserMetaData struct {
  Username string
}

type ClusterMetaData struct {
  KubeConfigPath string
}

type Contract struct {
  Duration int
  Rate int
}

func main() {
  fmt.Println("It's running at least")
  testImage := ImageMetaData{"testUrl", "testRep", "testImageName", "testImageTag"}
  testBidUser := UserMetaData{"testUserName"}
  testContract := Contract{1,1}
  testBid := Bid{&testImage, &testBidUser, &testContract}
  fmt.Println(testBid.ImageMetaData.RegistryUrl)
}




// type AskContract interface {
//   FulfilledBy(Bid) bool
// }
//
// type BidContract interface {
//   FulfilledBy(Ask) bool
// }
//
// type BasicContract interface {
//   FulfilledBy()
// }
//
//
// type ImageMetaData interface {
//   // tbd
// }
//
// type ClusterMetaData interface {
//   // tbd
// }
//
// type SimpleContractMetaData struct {
//   Price int
//   MaxCpu int
//   MinCpu int
//   MaxRam int
//   MinRam int
//   Guaranteed bool
// }
//
//
//
// type CompatibilityChecker interface{
//   Compatible(b Bid, a Ask) (result bool, err Error)
// }
//
//
//
// type SimpleCompatibilityChecker struct{}

// func (checker SimpleCompatibilityChecker) Compatible(b Bid, a Ask) bool {
//   if b.Price() > a.Price() && b.MaxCpu() == a.MaxCpu() && b.MinCpu() == a.MinCpu() && b.MaxRam() == a.MaxRam() && b.MinRam() == a.MinRam() {
//     return true
//   }
//   return false
// }
//
// type Fulfiller interface {
//   Fulfill(b Bid, a Ask) (success bool, err error)
// }
//
// type BidAskHandler interface {
//   CompatibilityChecker
//   Fulfiller
// }
//
// type BidGetter interface {
//   GetBids() // should return a collection of bids
// }
//
// type AskGetter interface {
//   GetAsks() // should return a collection of Asks
// }
//
// type Exchange interface {
//   BidAskHandler
//   BidGetter
//   AskGetter
// }
