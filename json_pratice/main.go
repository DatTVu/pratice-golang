package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type hardwareInfo struct {
	vCPU   float64
	vRAM   float64
	counts float64
}

func main() {
	var instance_map = make(map[string]hardwareInfo)
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
		instance_map[instanceType.(string)] = hardwareInfo{
			vCPU:   vCPU.(float64),
			vRAM:   vRam.(float64),
			counts: counts.(float64),
		}
	}
	fmt.Println(instance_map)
}
