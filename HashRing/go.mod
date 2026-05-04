module github.com/sajeelwaien/consistent-hashing/hashring

go 1.26.2

replace github.com/sajeelwaien/consistent-hashing/bloomfilter => ../BloomFilter

replace github.com/sajeelwaien/consistent-hashing/node => ../Node

require (
	github.com/sajeelwaien/consistent-hashing/node v1.0.0
	github.com/spaolacci/murmur3 v1.1.0
)
