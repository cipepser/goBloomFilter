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

## IBLT

Supports following method:  

* Insert
* のGet
* Delete
* ListEntries

## 参考
* [C++でブルームフィルタを実装する方法](http://postd.cc/how-to-write-a-bloom-filter-cpp/)
* [Double hashing - Wikipedia](https://en.wikipedia.org/wiki/Double_hashing)
* [Bloom filter - Wikipedia](https://en.wikipedia.org/wiki/Bloom_filter)
* [ビットコインキャッシュの使用ブロック伝播速度を10倍にする新技術”Graphene”](https://bitcoiner.link/6542.html)
* [Graphene: A New Protocol for Block Propagation Using Set Reconciliation](https://people.cs.umass.edu/~gbiss/graphene.pdf)
* [Invertible Bloom Lookup Tables](https://arxiv.org/pdf/1101.2245.pdf)
* [O(1) Block Propagation](https://gist.github.com/gavinandresen/e20c3b5a1d4b97f79ac2)
* [ビットコインとは何か? 第4回:ビットコインの仕組み(取引承認と採掘) - Bitcoin日本語情報サイト](https://jpbitcoin.com/about/whatisbitcoin4)
* [BloomFilterについて調べてみた - 逆さまにした](http://cipepser.hatenablog.com/entry/2017/01/28/205316)
* [GolangでBloomFilterを実装してみた - 逆さまにした](http://cipepser.hatenablog.com/entry/2017/02/04/090629)
* [Golangでcounting BloomFilterを実装してみた - 逆さまにした](http://cipepser.hatenablog.com/entry/2017/02/24/234002)