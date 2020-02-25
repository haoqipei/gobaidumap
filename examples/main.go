package main

import (
	"fmt"
	"github.com/menduo/gobaidumap"
)

func main() {

	lat := "40.069462"
	lng := "116.346364"

	bc := gobaidumap.NewBaiduMapClient("6B5ny4pqLFLBK50HV4P4HURTD6XaYft8")

	// 从坐标到地址
	GEOToAddress, err := bc.GetAddressViaGEO(lat, lng)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("坐标到地址：", GEOToAddress)
		fmt.Println("坐标到地址 - 地址", GEOToAddress.Result.AddressComponent)
		fmt.Println("\n")
	}

	// 从地址到坐标
	address := "百度大厦"
	// address = url.QueryEscape(address)
	addressToGEO, err := bc.GetGeoViaAddress(address)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("从地址到坐标 - All", addressToGEO)
		fmt.Println("从地址到坐标 - Lat", addressToGEO.Result.Location.Lat)
		fmt.Println("从地址到坐标 - Lng", addressToGEO.Result.Location.Lng)
		fmt.Println("\n")
	}


	// 从IP到地址
	ipAddress := "202.198.16.3"
	IPToAddress, err := bc.GetAddressViaIP(ipAddress)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("从IP到地址：", IPToAddress)
		fmt.Println("从IP到地址 - 地址：", IPToAddress, IPToAddress.Content.Address)
		fmt.Println("\n")
	}

	// 从IP到地址
	ipAddress = "8.8.8.8"
	IPToAddress, err = bc.GetAddressViaIP(ipAddress)

	if err != nil {
		fmt.Println("从IP到地址，err !=nil：", err)
		fmt.Println("\n")
	} else {
		fmt.Println("从IP到地址：", IPToAddress)
		fmt.Println("从IP到地址 - 地址：", IPToAddress, IPToAddress.Content.Address)
		fmt.Println("\n")
	}
}
