module github.com/ev2-1/mt-multiserver-signs

go 1.18

require (
	github.com/HimbeerserverDE/mt-multiserver-proxy v0.0.0-20220514170657-54c9f9bb8d42
	github.com/anon55555/mt v0.0.0-20210919124550-bcc58cb3048f
)

require github.com/HimbeerserverDE/srp v0.0.0 // indirect

replace github.com/HimbeerserverDE/mt-multiserver-proxy => ../../proxy
