package main

import (
	"bufio"
	"net"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

var cache = &Cache{m: make(map[string]string)}

type Cache struct {
	mx sync.RWMutex
	m  map[string]string
}

func (c *Cache) set(key string, val string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = val
}

func (c *Cache) get(key string) (string, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Cache) delete(key string) (ok bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	_, ok = c.m[key]
	delete(c.m, key)
	return ok
}

func (c *Cache) exists(key string) (ok bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	_, ok = c.m[key]
	return ok
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Connected. Wait for commands\n\r"))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		textArray := regexp.MustCompile(`[^\s"]+|"([^"]*)"`).FindAllString(text, -1)
		command := strings.ToLower(textArray[0])
		if command == "exit" {
			conn.Write([]byte("Bye\n\r"))
			break
		} else if command == "set" {
			if len(textArray) == 3 {
				cache.set(textArray[1], textArray[2])
				conn.Write([]byte("OK\n\r"))
			} else {
				conn.Write([]byte("Wrong parameters in set\n\r"))
			}
		} else if command == "get" {
			if len(textArray) == 2 {
				val, ok := cache.get(textArray[1])
				if ok {
					conn.Write([]byte("1\n\r"))
					conn.Write([]byte(val + "\n\r"))
				} else {
					conn.Write([]byte("0\n\r"))
				}
			} else {
				conn.Write([]byte("Wrong parameters in get\n\r"))
			}
		} else if command == "del" {
			if len(textArray) == 2 {
				ok := cache.delete(textArray[1])
				if ok {
					conn.Write([]byte("1\n\r"))
				} else {
					conn.Write([]byte("0\n\r"))
				}
			} else {
				conn.Write([]byte("Wrong parameters in del\n\r"))
			}
		} else if command == "exists" {
			if len(textArray) == 2 {
				ok := cache.exists(textArray[1])
				if ok {
					conn.Write([]byte("1\n\r"))
				} else {
					conn.Write([]byte("0\n\r"))
				}
			} else {
				conn.Write([]byte("Wrong parameters in exist\n\r"))
			}
		}
	}
}

func MemControl(sizeKb uint64) {
	for {
		memStats := &runtime.MemStats{}
		runtime.ReadMemStats(memStats)
		if memStats.Alloc/1024 > sizeKb {
			panic("Memory critical threshold exceeded")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listner, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	//Ограничение памяти 1000 килобайт
	go MemControl(1000)

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

//go build -o tcp_listner.exe . && tcp_listner.exe
