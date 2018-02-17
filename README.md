# goBloomFilter

[![Build Status](https://travis-ci.org/cipepser/goBloomFilter.svg?branch=master)](https://travis-ci.org/cipepser/goBloomFilter)
[![Coverage Status](https://coveralls.io/repos/github/cipepser/goBloomFilter/badge.svg?branch=master)](https://coveralls.io/github/cipepser/goBloomFilter?branch=master)

いろいろなBloomFilterのGo実装

## bf(general Bloom filter)

Supports following method:  
* Add
* Exists

bf don't support Remove.

## cbf(counting Bloom filter)

Supports following method:  
* Add
* Exists
* Remove

## 参考
* [C++でブルームフィルタを実装する方法](http://postd.cc/how-to-write-a-bloom-filter-cpp/)
* [Double hashing - Wikipedia](https://en.wikipedia.org/wiki/Double_hashing)
* [Bloom filter - Wikipedia](https://en.wikipedia.org/wiki/Bloom_filter)