package main

import "fmt"

// 用多态来模拟实现，将移动硬盘或者U盘插到电脑上进行读写数据

type Storager interface {
	Read()
	Write()
}

type MobileDisk struct {
	id string
}

func (m *MobileDisk) Read() {
	fmt.Println("移动硬盘读取数据")
}

func (m *MobileDisk) Write() {
	fmt.Println("移动硬盘写入数据")
}

type UsbDisk struct {
	id string
}

func (u *UsbDisk) Read() {
	fmt.Println("U盘读取数据")
}

func (u *UsbDisk) Write() {
	fmt.Println("U盘写入数据")
}

func Computer(c Storager) {
	c.Read()
	c.Write()
}

func main() {
	var m MobileDisk
	var u UsbDisk

	Computer(&m)
	Computer(&u)
}
