[![Build Status](https://travis-ci.org/menggggggg/lru.svg?branch=master)](https://travis-ci.org/menggggggg/lru)
[![Coverage Status](https://coveralls.io/repos/github/menggggggg/lru/badge.svg?branch=master)](https://coveralls.io/github/menggggggg/lru?branch=master)

# LRU- Least Recently Used

操作系统中进行内存管理中时采用一些页面置换算法，如LRU、LFU和FIFO等。其中LRU应用较为广泛。

大家都知道在缓存的大小是有限的，那么我们应该基于什么策略进行缓存数据呢？LRU提供的思路是将最近没有使用的数据从缓存中移除，这样的思路在实际的环境中比较符合常识。

# 原理

LRU算法的原理比较简单，数据存储的数据结构为链表。当访问数据时，如缓存中有数据，则将该数据移动至链表的顶端；没有该数据则在顶端加入该数据，并移除链表中的低端的数据。

