package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type HardwareInfo struct {
	VCPU   float64
	VRAM   float64
	Counts float64
}

func main() {
	var instance_map = make(map[string]HardwareInfo)
	var result map[string]interface{}
	content, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(content, &result)
	strs := result["Instances"].([]interface{})
	for _, val := range strs {
		instanceType := val.(map[string]interface{})["type"]
		vCPU := val.(map[string]interface{})["vCPU"]
		vRam := val.(map[string]interface{})["vRam"]
		counts := val.(map[string]interface{})["counts"]
		instance_map[instanceType.(string)] = HardwareInfo{
			VCPU:   vCPU.(float64),
			VRAM:   vRam.(float64),
			Counts: counts.(float64),
		}
	}
	fmt.Println(instance_map)
}
