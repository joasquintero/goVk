package src

import (
	"github.com/vulkan-go/vulkan"
	"goVk/info"
)

var appInfo = & vulkan.ApplicationInfo{
	SType:				vulkan.StructureTypeApplicationInfo,
	ApiVersion: 		vulkan.MakeVersion(1,0,0),
	ApplicationVersion: vulkan.MakeVersion(1,0,0),
	PApplicationName: 	"Go Vulkan Info\x00",
	PEngineName: 		"vulkan-compute\x00",

}
func Init()  {
	orPanic(vulkan.SetDefaultGetInstanceProcAddr())
	orPanic(vulkan.Init())
	vkDevice, err := info.NewVulkanDevice(appInfo,0)
	orPanic(err)
	info.PrintInfo(vkDevice)
	vkDevice.Destroy()
}

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case vulkan.Result:
		if err := vulkan.Error(v); err != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}