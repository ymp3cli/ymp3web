package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CurrentSong struct {
	By    string
	Url   string
	Img   string
	Title string
}

func Ping(url string) map[string]string {
	client := http.Client{
		Timeout: 5 * 1e6, // 5 ms timeout
	}
	req, err := client.Get(url)

	if err != nil {
		return map[string]string{"status": "offline"}
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	var song CurrentSong
	json.Unmarshal(body, &song)
	// return a json
	return map[string]string{"status": "online", "by": song.By, "url": song.Url, "img": song.Img, "title": song.Title}

}

func ScanPorts(c echo.Context) error {
	/*
		Class A: 10.0. 0.0 to 10.255. 255.255.
		Class B: 172.16. 0.0 to 172.31. 255.255.
		Class C: 192.168. 0.0 to 192.168. 255.255.
	*/
	ip, err := LocalIp()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
	ip = ip[:len(ip)-3]
	for i := 1; i <= 255; i++ {
		port := strconv.Itoa(i)
		url := "http://" + ip + "." + port + ":8888" + "/currentSong"
		fmt.Println(url)
		status := Ping(url)
		if status["status"] == "online" {
			fmt.Println("ymp3cli-ports, port " + port + " is online")
			fmt.Println(status)
			songs, err := http.Get("http://" + ip + "." + port + ":8888/songs")
			if err != nil {
				fmt.Println(err)
			}
			defer songs.Body.Close()
			body, _ := ioutil.ReadAll(songs.Body)
			json.NewEncoder(c.Response()).Encode(map[string]string{"status": "online", "by": status["by"], "url": status["url"], "img": status["img"], "title": status["title"], "ip": "http://" + ip + "." + port + ":8888", "songs": string(body)})
			break
		}

	}
	return nil

}

func LocalIp() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
